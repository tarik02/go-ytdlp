// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Video Selection Option Group

package ytdlp

import (
	"strconv"
)

// PlaylistStart maps to cli flags: --playlist-start=NUMBER (hidden).
func (c *Command) PlaylistStart(number int) *Command {
	c.addFlag(&Flag{
		ID:   "playliststart",
		Flag: "--playlist-start",
		Args: []string{
			strconv.Itoa(number),
		},
	})
	return c
}

// UnsetPlaylistStart unsets any flags that were previously set by
// PlaylistStart().
func (c *Command) UnsetPlaylistStart() *Command {
	c.removeFlagByID("playliststart")
	return c
}

// PlaylistEnd maps to cli flags: --playlist-end=NUMBER (hidden).
func (c *Command) PlaylistEnd(number int) *Command {
	c.addFlag(&Flag{
		ID:   "playlistend",
		Flag: "--playlist-end",
		Args: []string{
			strconv.Itoa(number),
		},
	})
	return c
}

// UnsetPlaylistEnd unsets any flags that were previously set by
// PlaylistEnd().
func (c *Command) UnsetPlaylistEnd() *Command {
	c.removeFlagByID("playlistend")
	return c
}

// Comma separated playlist_index of the items to download. You can specify a range
// using "[START]:[STOP][:STEP]". For backward compatibility, START-STOP is also
// supported. Use negative indices to count from the right and negative STEP to
// download in reverse order. E.g. "-I 1:3,7,-5::2" used on a playlist of size 15
// will download the items at index 1,2,3,7,11,13,15
//
// PlaylistItems maps to cli flags: -I/--playlist-items=ITEM_SPEC.
func (c *Command) PlaylistItems(itemSpec string) *Command {
	c.addFlag(&Flag{
		ID:   "playlist_items",
		Flag: "--playlist-items",
		Args: []string{itemSpec},
	})
	return c
}

// UnsetPlaylistItems unsets any flags that were previously set by
// PlaylistItems().
func (c *Command) UnsetPlaylistItems() *Command {
	c.removeFlagByID("playlist_items")
	return c
}

// MatchTitle maps to cli flags: --match-title=REGEX (hidden).
func (c *Command) MatchTitle(regex string) *Command {
	c.addFlag(&Flag{
		ID:   "matchtitle",
		Flag: "--match-title",
		Args: []string{regex},
	})
	return c
}

// UnsetMatchTitle unsets any flags that were previously set by
// MatchTitle().
func (c *Command) UnsetMatchTitle() *Command {
	c.removeFlagByID("matchtitle")
	return c
}

// RejectTitle maps to cli flags: --reject-title=REGEX (hidden).
func (c *Command) RejectTitle(regex string) *Command {
	c.addFlag(&Flag{
		ID:   "rejecttitle",
		Flag: "--reject-title",
		Args: []string{regex},
	})
	return c
}

// UnsetRejectTitle unsets any flags that were previously set by
// RejectTitle().
func (c *Command) UnsetRejectTitle() *Command {
	c.removeFlagByID("rejecttitle")
	return c
}

// Abort download if filesize is smaller than SIZE, e.g. 50k or 44.6M
//
// MinFilesize maps to cli flags: --min-filesize=SIZE.
func (c *Command) MinFilesize(size string) *Command {
	c.addFlag(&Flag{
		ID:   "min_filesize",
		Flag: "--min-filesize",
		Args: []string{size},
	})
	return c
}

// UnsetMinFilesize unsets any flags that were previously set by
// MinFilesize().
func (c *Command) UnsetMinFilesize() *Command {
	c.removeFlagByID("min_filesize")
	return c
}

// Abort download if filesize is larger than SIZE, e.g. 50k or 44.6M
//
// MaxFilesize maps to cli flags: --max-filesize=SIZE.
func (c *Command) MaxFilesize(size string) *Command {
	c.addFlag(&Flag{
		ID:   "max_filesize",
		Flag: "--max-filesize",
		Args: []string{size},
	})
	return c
}

// UnsetMaxFilesize unsets any flags that were previously set by
// MaxFilesize().
func (c *Command) UnsetMaxFilesize() *Command {
	c.removeFlagByID("max_filesize")
	return c
}

// Download only videos uploaded on this date. The date can be "YYYYMMDD" or in the
// format [now|today|yesterday][-N[day|week|month|year]]. E.g. "--date
// today-2weeks" downloads only videos uploaded on the same day two weeks ago
//
// Date maps to cli flags: --date=DATE.
func (c *Command) Date(date string) *Command {
	c.addFlag(&Flag{
		ID:   "date",
		Flag: "--date",
		Args: []string{date},
	})
	return c
}

// UnsetDate unsets any flags that were previously set by
// Date().
func (c *Command) UnsetDate() *Command {
	c.removeFlagByID("date")
	return c
}

// Download only videos uploaded on or before this date. The date formats accepted
// is the same as --date
//
// Datebefore maps to cli flags: --datebefore=DATE.
func (c *Command) Datebefore(date string) *Command {
	c.addFlag(&Flag{
		ID:   "datebefore",
		Flag: "--datebefore",
		Args: []string{date},
	})
	return c
}

// UnsetDatebefore unsets any flags that were previously set by
// Datebefore().
func (c *Command) UnsetDatebefore() *Command {
	c.removeFlagByID("datebefore")
	return c
}

// Download only videos uploaded on or after this date. The date formats accepted
// is the same as --date
//
// Dateafter maps to cli flags: --dateafter=DATE.
func (c *Command) Dateafter(date string) *Command {
	c.addFlag(&Flag{
		ID:   "dateafter",
		Flag: "--dateafter",
		Args: []string{date},
	})
	return c
}

// UnsetDateafter unsets any flags that were previously set by
// Dateafter().
func (c *Command) UnsetDateafter() *Command {
	c.removeFlagByID("dateafter")
	return c
}

// MinViews maps to cli flags: --min-views=COUNT (hidden).
func (c *Command) MinViews(count int) *Command {
	c.addFlag(&Flag{
		ID:   "min_views",
		Flag: "--min-views",
		Args: []string{
			strconv.Itoa(count),
		},
	})
	return c
}

// UnsetMinViews unsets any flags that were previously set by
// MinViews().
func (c *Command) UnsetMinViews() *Command {
	c.removeFlagByID("min_views")
	return c
}

// MaxViews maps to cli flags: --max-views=COUNT (hidden).
func (c *Command) MaxViews(count int) *Command {
	c.addFlag(&Flag{
		ID:   "max_views",
		Flag: "--max-views",
		Args: []string{
			strconv.Itoa(count),
		},
	})
	return c
}

// UnsetMaxViews unsets any flags that were previously set by
// MaxViews().
func (c *Command) UnsetMaxViews() *Command {
	c.removeFlagByID("max_views")
	return c
}

// Generic video filter. Any "OUTPUT TEMPLATE" field can be compared with a number
// or a string using the operators defined in "Filtering Formats". You can also
// simply specify a field to match if the field is present, use "!field" to check
// if the field is not present, and "&" to check multiple conditions. Use a "\" to
// escape "&" or quotes if needed. If used multiple times, the filter matches if
// atleast one of the conditions are met. E.g. --match-filter !is_live
// --match-filter "like_count>?100 & description~='(?i)\bcats \& dogs\b'" matches
// only videos that are not live OR those that have a like count more than 100 (or
// the like field is not available) and also has a description that contains the
// phrase "cats & dogs" (caseless). Use "--match-filter -" to interactively ask
// whether to download each video
//
// MatchFilters maps to cli flags: --match-filters=FILTER.
func (c *Command) MatchFilters(filter string) *Command {
	c.addFlag(&Flag{
		ID:   "match_filter",
		Flag: "--match-filters",
		Args: []string{filter},
	})
	return c
}

// UnsetMatchFilters unsets any flags that were previously set by
// MatchFilters().
func (c *Command) UnsetMatchFilters() *Command {
	c.removeFlagByID("match_filter")
	return c
}

// Same as "--match-filters" but stops the download process when a video is
// rejected
//
// BreakMatchFilters maps to cli flags: --break-match-filters=FILTER.
func (c *Command) BreakMatchFilters(filter string) *Command {
	c.addFlag(&Flag{
		ID:   "breaking_match_filter",
		Flag: "--break-match-filters",
		Args: []string{filter},
	})
	return c
}

// UnsetBreakMatchFilters unsets any flags that were previously set by
// BreakMatchFilters().
func (c *Command) UnsetBreakMatchFilters() *Command {
	c.removeFlagByID("breaking_match_filter")
	return c
}

// Download only the video, if the URL refers to a video and a playlist
//
// NoPlaylist maps to cli flags: --no-playlist.
func (c *Command) NoPlaylist() *Command {
	c.addFlag(&Flag{
		ID:   "noplaylist",
		Flag: "--no-playlist",
		Args: nil,
	})
	return c
}

// UnsetNoPlaylist unsets any flags that were previously set by
// NoPlaylist().
func (c *Command) UnsetNoPlaylist() *Command {
	c.removeFlagByID("noplaylist")
	return c
}

// Download the playlist, if the URL refers to a video and a playlist
//
// YesPlaylist maps to cli flags: --yes-playlist.
func (c *Command) YesPlaylist() *Command {
	c.addFlag(&Flag{
		ID:   "noplaylist",
		Flag: "--yes-playlist",
		Args: nil,
	})
	return c
}

// UnsetYesPlaylist unsets any flags that were previously set by
// YesPlaylist().
func (c *Command) UnsetYesPlaylist() *Command {
	c.removeFlagByID("noplaylist")
	return c
}

// Download only videos suitable for the given age
//
// AgeLimit maps to cli flags: --age-limit=YEARS.
func (c *Command) AgeLimit(years int) *Command {
	c.addFlag(&Flag{
		ID:   "age_limit",
		Flag: "--age-limit",
		Args: []string{
			strconv.Itoa(years),
		},
	})
	return c
}

// UnsetAgeLimit unsets any flags that were previously set by
// AgeLimit().
func (c *Command) UnsetAgeLimit() *Command {
	c.removeFlagByID("age_limit")
	return c
}

// Download only videos not listed in the archive file. Record the IDs of all
// downloaded videos in it
//
// DownloadArchive maps to cli flags: --download-archive=FILE.
func (c *Command) DownloadArchive(file string) *Command {
	c.addFlag(&Flag{
		ID:   "download_archive",
		Flag: "--download-archive",
		Args: []string{file},
	})
	return c
}

// UnsetDownloadArchive unsets any flags that were previously set by
// DownloadArchive().
func (c *Command) UnsetDownloadArchive() *Command {
	c.removeFlagByID("download_archive")
	return c
}

// Abort after downloading NUMBER files
//
// MaxDownloads maps to cli flags: --max-downloads=NUMBER.
func (c *Command) MaxDownloads(number int) *Command {
	c.addFlag(&Flag{
		ID:   "max_downloads",
		Flag: "--max-downloads",
		Args: []string{
			strconv.Itoa(number),
		},
	})
	return c
}

// UnsetMaxDownloads unsets any flags that were previously set by
// MaxDownloads().
func (c *Command) UnsetMaxDownloads() *Command {
	c.removeFlagByID("max_downloads")
	return c
}

// Stop the download process when encountering a file that is in the archive
//
// BreakOnExisting maps to cli flags: --break-on-existing.
func (c *Command) BreakOnExisting() *Command {
	c.addFlag(&Flag{
		ID:   "break_on_existing",
		Flag: "--break-on-existing",
		Args: nil,
	})
	return c
}

// UnsetBreakOnExisting unsets any flags that were previously set by
// BreakOnExisting().
func (c *Command) UnsetBreakOnExisting() *Command {
	c.removeFlagByID("break_on_existing")
	return c
}

// BreakOnReject sets the "break-on-reject" flag (no description specified).
//
// BreakOnReject maps to cli flags: --break-on-reject (hidden).
func (c *Command) BreakOnReject() *Command {
	c.addFlag(&Flag{
		ID:   "break_on_reject",
		Flag: "--break-on-reject",
		Args: nil,
	})
	return c
}

// UnsetBreakOnReject unsets any flags that were previously set by
// BreakOnReject().
func (c *Command) UnsetBreakOnReject() *Command {
	c.removeFlagByID("break_on_reject")
	return c
}

// Alters --max-downloads, --break-on-existing, --break-match-filter, and
// autonumber to reset per input URL
//
// BreakPerInput maps to cli flags: --break-per-input.
func (c *Command) BreakPerInput() *Command {
	c.addFlag(&Flag{
		ID:   "break_per_url",
		Flag: "--break-per-input",
		Args: nil,
	})
	return c
}

// UnsetBreakPerInput unsets any flags that were previously set by
// BreakPerInput().
func (c *Command) UnsetBreakPerInput() *Command {
	c.removeFlagByID("break_per_url")
	return c
}

// --break-on-existing and similar options terminates the entire download queue
//
// NoBreakPerInput maps to cli flags: --no-break-per-input.
func (c *Command) NoBreakPerInput() *Command {
	c.addFlag(&Flag{
		ID:   "break_per_url",
		Flag: "--no-break-per-input",
		Args: nil,
	})
	return c
}

// UnsetNoBreakPerInput unsets any flags that were previously set by
// NoBreakPerInput().
func (c *Command) UnsetNoBreakPerInput() *Command {
	c.removeFlagByID("break_per_url")
	return c
}

// Number of allowed failures until the rest of the playlist is skipped
//
// SkipPlaylistAfterErrors maps to cli flags: --skip-playlist-after-errors=N.
func (c *Command) SkipPlaylistAfterErrors(n int) *Command {
	c.addFlag(&Flag{
		ID:   "skip_playlist_after_errors",
		Flag: "--skip-playlist-after-errors",
		Args: []string{
			strconv.Itoa(n),
		},
	})
	return c
}

// UnsetSkipPlaylistAfterErrors unsets any flags that were previously set by
// SkipPlaylistAfterErrors().
func (c *Command) UnsetSkipPlaylistAfterErrors() *Command {
	c.removeFlagByID("skip_playlist_after_errors")
	return c
}

// IncludeAds sets the "include-ads" flag (no description specified).
//
// IncludeAds maps to cli flags: --include-ads (hidden).
func (c *Command) IncludeAds() *Command {
	c.addFlag(&Flag{
		ID:   "include_ads",
		Flag: "--include-ads",
		Args: nil,
	})
	return c
}

// UnsetIncludeAds unsets any flags that were previously set by
// IncludeAds().
func (c *Command) UnsetIncludeAds() *Command {
	c.removeFlagByID("include_ads")
	return c
}

// NoIncludeAds sets the "no-include-ads" flag (no description specified).
//
// NoIncludeAds maps to cli flags: --no-include-ads (hidden).
func (c *Command) NoIncludeAds() *Command {
	c.addFlag(&Flag{
		ID:   "include_ads",
		Flag: "--no-include-ads",
		Args: nil,
	})
	return c
}

// UnsetNoIncludeAds unsets any flags that were previously set by
// NoIncludeAds().
func (c *Command) UnsetNoIncludeAds() *Command {
	c.removeFlagByID("include_ads")
	return c
}
