package mediamachine

type WatermarkPosition = string

const (
	PositionTopLeft     WatermarkPosition = "topLeft"
	PositionTopRight    WatermarkPosition = "topRight"
	PositionBottomLeft  WatermarkPosition = "bottomLeft"
	PositionBottomRight WatermarkPosition = "bottomRight"
)

// WatermarkText can be used for a simple text watermark overlaid on the output
type WatermarkText struct {
	Text      string            // The text to display as the watermark
	FontSize  uint              // Optional - defaults to 10
	FontColor string            // Optional - defaults to black
	Opacity   float32           // Opacity of watermark between 0 and 1 inclusive
	Position  WatermarkPosition // Where the watermark should be placed. See WatermarkPosition
}

// WatermarkImageURL can be used to supply an image url which will be used as a watermark
type WatermarkImageURL struct {
	URL      string            // URL where the watermark image should be fetched from (currently, bucket urls are not supported)
	Height   uint8             // Height of the watermark
	Width    uint8             // Width of the watermark
	Opacity  float32           // Opacity of watermark between 0 and 1 inclusive
	Position WatermarkPosition // Where the watermark should be placed. See WatermarkPosition
}

// WatermarkImageNamed can be used to provide a reference to a watermark image uploaded to your mediamachine account
// You can easily upload your watermark images via account settings. The uploaded image gets a unique name that can be used here.
type WatermarkImageNamed struct {
	ImageName string            // Name of a watermark image uploaded on the mediamachine account
	Height    uint8             // Height of the watermark
	Width     uint8             // Width of the watermark
	Opacity   float32           // Opacity of watermark between 0 and 1 inclusive
	Position  WatermarkPosition // Where the watermark should be placed. See WatermarkPosition
}

type watermark interface {
	isWatermark()
}

func (WatermarkText) isWatermark()       {}
func (WatermarkImageNamed) isWatermark() {}
func (WatermarkImageURL) isWatermark()   {}
