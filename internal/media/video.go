/*
   GoToSocial
   Copyright (C) 2021-2023 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package media

import (
	"fmt"
	"io"
	"os"

	"github.com/abema/go-mp4"
)

type gtsVideo struct {
	frame     *gtsImage
	duration  float32 // in seconds
	bitrate   uint64
	framerate float32
}

// decodeVideoFrame decodes and returns an image from a single frame in the given video stream.
// (note: currently this only returns a blank image resized to fit video dimensions).
func decodeVideoFrame(r io.Reader) (*gtsVideo, error) {
	// We'll need a readseeker to decode the video. We can get a readseeker
	// without burning too much mem by first copying the reader into a temp file.
	// First create the file in the temporary directory...
	tmp, err := os.CreateTemp(os.TempDir(), "gotosocial-")
	if err != nil {
		return nil, err
	}

	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	// Now copy the entire reader we've been provided into the
	// temporary file; we won't use the reader again after this.
	if _, err := io.Copy(tmp, r); err != nil {
		return nil, err
	}

	// probe the video file to extract useful metadata from it; for methodology, see:
	// https://github.com/abema/go-mp4/blob/7d8e5a7c5e644e0394261b0cf72fef79ce246d31/mp4tool/probe/probe.go#L85-L154
	info, err := mp4.Probe(tmp)
	if err != nil {
		return nil, fmt.Errorf("error probing tmp file %s: %w", tmp.Name(), err)
	}

	var (
		width  int
		height int
		video  gtsVideo
	)

	for _, tr := range info.Tracks {
		if tr.AVC == nil {
			continue
		}

		if w := int(tr.AVC.Width); w > width {
			width = w
		}

		if h := int(tr.AVC.Height); h > height {
			height = h
		}

		if br := tr.Samples.GetBitrate(tr.Timescale); br > video.bitrate {
			video.bitrate = br
		} else if br := info.Segments.GetBitrate(tr.TrackID, tr.Timescale); br > video.bitrate {
			video.bitrate = br
		}

		if d := float64(tr.Duration) / float64(tr.Timescale); d > float64(video.duration) {
			video.framerate = float32(len(tr.Samples)) / float32(d)
			video.duration = float32(d)
		}
	}

	// Check for empty video metadata.
	var empty []string
	if width == 0 {
		empty = append(empty, "width")
	}
	if height == 0 {
		empty = append(empty, "height")
	}
	if video.duration == 0 {
		empty = append(empty, "duration")
	}
	if video.framerate == 0 {
		empty = append(empty, "framerate")
	}
	if video.bitrate == 0 {
		empty = append(empty, "bitrate")
	}
	if len(empty) > 0 {
		return nil, fmt.Errorf("error determining video metadata: %v", empty)
	}

	// Create new empty "frame" image.
	// TODO: decode frame from video file.
	video.frame = blankImage(width, height)

	return &video, nil
}
