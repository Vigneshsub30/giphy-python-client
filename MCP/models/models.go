package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Count int `json:"count,omitempty"` // Total number of items returned.
	Offset int `json:"offset,omitempty"` // Position in pagination.
	Total_count int `json:"total_count,omitempty"` // Total number of items available.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Twitter string `json:"twitter,omitempty"` // The Twitter username associated with this user, if applicable.
	Username string `json:"username,omitempty"` // The username associated with this user.
	Avatar_url string `json:"avatar_url,omitempty"` // The URL for this user's avatar image.
	Banner_url string `json:"banner_url,omitempty"` // The URL for the banner image that appears atop this user's profile page.
	Display_name string `json:"display_name,omitempty"` // The display name associated with this user (contains formatting the base username might not).
	Profile_url string `json:"profile_url,omitempty"` // The URL for this user's profile.
}

// Gif represents the Gif schema from the OpenAPI specification
type Gif struct {
	Update_datetime string `json:"update_datetime,omitempty"` // The date on which this GIF was last updated.
	Bitly_url string `json:"bitly_url,omitempty"` // The unique bit.ly URL for this GIF
	Trending_datetime string `json:"trending_datetime,omitempty"` // The date on which this gif was marked trending, if applicable.
	User User `json:"user,omitempty"` // The User Object contains information about the user associated with a GIF and URLs to assets such as that user's avatar image, profile, and more.
	Source string `json:"source,omitempty"` // The page on which this GIF was found
	Tags []string `json:"tags,omitempty"` // An array of tags for this GIF (Note: Not available when using the Public Beta Key)
	Create_datetime string `json:"create_datetime,omitempty"` // The date this GIF was added to the GIPHY database.
	Id string `json:"id,omitempty"` // This GIF's unique ID
	Import_datetime string `json:"import_datetime,omitempty"` // The creation or upload date from this GIF's source.
	Slug string `json:"slug,omitempty"` // The unique slug used in this GIF's URL
	Url string `json:"url,omitempty"` // The unique URL for this GIF
	TypeField string `json:"type,omitempty"` // Type of the gif. By default, this is almost always gif
	Images map[string]interface{} `json:"images,omitempty"` // An object containing data for various available formats and sizes of this GIF.
	Username string `json:"username,omitempty"` // The username this GIF is attached to, if applicable
	Content_url string `json:"content_url,omitempty"` // Currently unused
	Rating string `json:"rating,omitempty"` // The MPAA-style rating for this content. Examples include Y, G, PG, PG-13 and R
	Embded_url string `json:"embded_url,omitempty"` // A URL used for embedding this GIF
	Source_post_url string `json:"source_post_url,omitempty"` // The URL of the webpage on which this GIF was found.
	Source_tld string `json:"source_tld,omitempty"` // The top level domain of the source URL.
	Featured_tags []string `json:"featured_tags,omitempty"` // An array of featured tags for this GIF (Note: Not available when using the Public Beta Key)
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Width string `json:"width,omitempty"` // The width of this GIF in pixels.
	Mp4_size string `json:"mp4_size,omitempty"` // The size in bytes of the .MP4 file corresponding to this GIF.
	Webp string `json:"webp,omitempty"` // The URL for this GIF in .webp format.
	Webp_size string `json:"webp_size,omitempty"` // The size in bytes of the .webp file corresponding to this GIF.
	Height string `json:"height,omitempty"` // The height of this GIF in pixels.
	Mp4 string `json:"mp4,omitempty"` // The URL for this GIF in .MP4 format.
	Size string `json:"size,omitempty"` // The size of this GIF in bytes.
	Url string `json:"url,omitempty"` // The publicly-accessible direct URL for this GIF.
	Frames string `json:"frames,omitempty"` // The number of frames in this GIF.
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Status int `json:"status,omitempty"` // HTTP Response Code
	Msg string `json:"msg,omitempty"` // HTTP Response Message
	Response_id string `json:"response_id,omitempty"` // A unique ID paired with this response from the API.
}
