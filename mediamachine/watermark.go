package mediamachine

// WatermarkPosition are references to named, pre-defined watermark locations
type WatermarkPosition = string

const (
	// PositionTopLeft places a watermark in the top left corner of the output
	PositionTopLeft WatermarkPosition = "topLeft"
	// PositionTopRight places a watermark in the top right corner of the output
	PositionTopRight WatermarkPosition = "topRight"
	// PositionBottomLeft places a watermark in the bottom left corner of the output
	PositionBottomLeft WatermarkPosition = "bottomLeft"
	// PositionBottomRight places a watermark in the bottom right corner of the output
	PositionBottomRight WatermarkPosition = "bottomRight"
)

// WatermarkText can be used for a simple text Watermark overlaid on the output
type WatermarkText struct {
	Text      string            // The text to display as the Watermark
	FontSize  uint              // Optional - defaults to 10
	FontColor string            // Optional - defaults to black
	Opacity   float32           // Opacity of Watermark between 0 and 1 inclusive
	Position  WatermarkPosition // Where the Watermark should be placed. See WatermarkPosition
}

// WatermarkImageURL can be used to supply an image url which will be used as a Watermark
type WatermarkImageURL struct {
	URL      string            // URL where the Watermark image should be fetched from (currently, bucket urls are not supported)
	Height   uint8             // Height of the Watermark
	Width    uint8             // Width of the Watermark
	Opacity  float32           // Opacity of Watermark between 0 and 1 inclusive
	Position WatermarkPosition // Where the Watermark should be placed. See WatermarkPosition
}

// WatermarkImageNamed can be used to provide a reference to a Watermark image uploaded to your mediamachine account
// You can easily upload your Watermark images via account settings. The uploaded image gets a unique name that can be used here.
type WatermarkImageNamed struct {
	ImageName string            // Name of a Watermark image uploaded on the mediamachine account
	Height    uint8             // Height of the Watermark
	Width     uint8             // Width of the Watermark
	Opacity   float32           // Opacity of Watermark between 0 and 1 inclusive
	Position  WatermarkPosition // Where the Watermark should be placed. See WatermarkPosition
}

// Watermark can be of multiple types - text watermark, image or a saved image reference watermark
type Watermark interface {
	isWatermark()
}

func (WatermarkText) isWatermark()       {}
func (WatermarkImageNamed) isWatermark() {}
func (WatermarkImageURL) isWatermark()   {}
