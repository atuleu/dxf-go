// Code generated at build time; DO NOT EDIT.

package dxf

import (
	"errors"
	"fmt"
	"math"
)

type Entity interface {
	codePairs(version AcadVersion) (pairs []CodePair)
	minVersion() AcadVersion
	maxVersion() AcadVersion
	tryApplyCodePair(pair CodePair)
	typeString() string
	pointers() (pointers []*pointer)
	Handle() Handle
	SetHandle(val Handle)
	IsInPaperSpace() bool
	SetIsInPaperSpace(val bool)
	Layer() string
	SetLayer(val string)
	LineTypeName() string
	SetLineTypeName(val string)
	Elevation() float64
	SetElevation(val float64)
	MaterialHandle() string
	SetMaterialHandle(val string)
	Color() Color
	SetColor(val Color)
	LineWeight() LineWeight
	SetLineWeight(val LineWeight)
	LineTypeScale() float64
	SetLineTypeScale(val float64)
	IsVisible() bool
	SetIsVisible(val bool)
	ImageByteCount() int
	SetImageByteCount(val int)
	PreviewImageData() []string
	SetPreviewImageData(val []string)
	Color24Bit() int
	SetColor24Bit(val int)
	ColorName() string
	SetColorName(val string)
	Transparency() int
	SetTransparency(val int)
	ShadowMode() ShadowMode
	SetShadowMode(val ShadowMode)
	Owner() *DrawingItem
	SetOwner(val *DrawingItem)
	getOwnerPointer() pointer
	setOwnerPointerHandle(h Handle)
	PlotStyle() *DrawingItem
	SetPlotStyle(val *DrawingItem)
	getPlotStylePointer() pointer
	setPlotStylePointerHandle(h Handle)
}

func tryApplyCodePairForEntity(this Entity, codePair CodePair) bool {
	switch codePair.Code {
	case 5:
		this.SetHandle(handleFromString(codePair.Value.(StringCodePairValue).Value))
	case 67:
		this.SetIsInPaperSpace(boolFromShort(codePair.Value.(ShortCodePairValue).Value))
	case 8:
		this.SetLayer(codePair.Value.(StringCodePairValue).Value)
	case 6:
		this.SetLineTypeName(codePair.Value.(StringCodePairValue).Value)
	case 38:
		this.SetElevation(codePair.Value.(DoubleCodePairValue).Value)
	case 347:
		this.SetMaterialHandle(codePair.Value.(StringCodePairValue).Value)
	case 62:
		this.SetColor(Color(codePair.Value.(ShortCodePairValue).Value))
	case 370:
		this.SetLineWeight(LineWeight(codePair.Value.(ShortCodePairValue).Value))
	case 48:
		this.SetLineTypeScale(codePair.Value.(DoubleCodePairValue).Value)
	case 60:
		this.SetIsVisible(!boolFromShort(codePair.Value.(ShortCodePairValue).Value))
	case 92:
		this.SetImageByteCount(codePair.Value.(IntCodePairValue).Value)
	case 310:
		this.SetPreviewImageData(append(this.PreviewImageData(), codePair.Value.(StringCodePairValue).Value))
	case 420:
		this.SetColor24Bit(codePair.Value.(IntCodePairValue).Value)
	case 430:
		this.SetColorName(codePair.Value.(StringCodePairValue).Value)
	case 440:
		this.SetTransparency(codePair.Value.(IntCodePairValue).Value)
	case 284:
		this.SetShadowMode(ShadowMode(codePair.Value.(ShortCodePairValue).Value))
	case 330:
		this.setOwnerPointerHandle(handleFromString(codePair.Value.(StringCodePairValue).Value))
	case 390:
		this.setPlotStylePointerHandle(handleFromString(codePair.Value.(StringCodePairValue).Value))
	default:
		return false
	}
	return true
}

func codePairsForEntity(this Entity, version AcadVersion) (pairs []CodePair) {
	if this.Handle() != 0 {
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(this.Handle())))
	}
	if this.getOwnerPointer().handle != 0 {
		pairs = append(pairs, NewStringCodePair(330, stringFromHandle(this.getOwnerPointer().handle)))
	}
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbEntity"))
	}
	if version >= R12 && this.IsInPaperSpace() != false {
		pairs = append(pairs, NewShortCodePair(67, shortFromBool(this.IsInPaperSpace())))
	}
	pairs = append(pairs, NewStringCodePair(8, defaultIfEmpty(this.Layer(), "0")))
	if this.LineTypeName() != "BYLAYER" {
		pairs = append(pairs, NewStringCodePair(6, this.LineTypeName()))
	}
	if version <= R12 && this.Elevation() != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(38, this.Elevation()))
	}
	if version >= R2007 && this.MaterialHandle() != "BYLAYER" {
		pairs = append(pairs, NewStringCodePair(347, this.MaterialHandle()))
	}
	if this.Color() != ByLayer() {
		pairs = append(pairs, NewShortCodePair(62, int16(this.Color())))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(370, int16(this.LineWeight())))
	}
	if version >= R13 && this.LineTypeScale() != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(48, this.LineTypeScale()))
	}
	if version >= R13 && this.IsVisible() != true {
		pairs = append(pairs, NewShortCodePair(60, shortFromBool(!this.IsVisible())))
	}
	if version >= R2000 && this.ImageByteCount() != 0 {
		pairs = append(pairs, NewIntCodePair(92, this.ImageByteCount()))
	}
	if version >= R2000 {
		for _, val := range this.PreviewImageData() {
			pairs = append(pairs, NewStringCodePair(310, val))
		}
	}
	if version >= R2004 && this.Color24Bit() != 0 {
		pairs = append(pairs, NewIntCodePair(420, this.Color24Bit()))
	}
	if version >= R2004 {
		pairs = append(pairs, NewStringCodePair(430, this.ColorName()))
	}
	if version >= R2004 {
		pairs = append(pairs, NewIntCodePair(440, this.Transparency()))
	}
	if this.getPlotStylePointer().handle != 0 {
		pairs = append(pairs, NewStringCodePair(390, stringFromHandle(this.getPlotStylePointer().handle)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(284, int16(this.ShadowMode())))
	}
	return
}

type Dimension interface {
	Version() Version
	SetVersion(val Version)
	BlockName() string
	SetBlockName(val string)
	DefinitionPoint1() Point
	SetDefinitionPoint1(val Point)
	TextMidPoint() Point
	SetTextMidPoint(val Point)
	DimensionType() DimensionType
	SetDimensionType(val DimensionType)
	AttachmentPoint() AttachmentPoint
	SetAttachmentPoint(val AttachmentPoint)
	TextLineSpacingStyle() TextLineSpacingStyle
	SetTextLineSpacingStyle(val TextLineSpacingStyle)
	TextLineSpacingFactor() float64
	SetTextLineSpacingFactor(val float64)
	ActualMeasurement() float64
	SetActualMeasurement(val float64)
	Text() string
	SetText(val string)
	TextRotationAngle() float64
	SetTextRotationAngle(val float64)
	HorizontalDirectionAngle() float64
	SetHorizontalDirectionAngle(val float64)
	Normal() Vector
	SetNormal(val Vector)
	DimensionStyleName() string
	SetDimensionStyleName(val string)
}

func tryApplyCodePairForDimension(this Dimension, codePair CodePair) bool {
	switch codePair.Code {
	case 280:
		this.SetVersion(Version(codePair.Value.(ShortCodePairValue).Value))
	case 2:
		this.SetBlockName(codePair.Value.(StringCodePairValue).Value)
	case 10:
		temp := this.DefinitionPoint1()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetDefinitionPoint1(temp)
	case 20:
		temp := this.DefinitionPoint1()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetDefinitionPoint1(temp)
	case 30:
		temp := this.DefinitionPoint1()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetDefinitionPoint1(temp)
	case 11:
		temp := this.TextMidPoint()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetTextMidPoint(temp)
	case 21:
		temp := this.TextMidPoint()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetTextMidPoint(temp)
	case 31:
		temp := this.TextMidPoint()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetTextMidPoint(temp)
	case 70:
		this.SetDimensionType(DimensionType(codePair.Value.(ShortCodePairValue).Value))
	case 71:
		this.SetAttachmentPoint(AttachmentPoint(codePair.Value.(ShortCodePairValue).Value))
	case 72:
		this.SetTextLineSpacingStyle(TextLineSpacingStyle(codePair.Value.(ShortCodePairValue).Value))
	case 41:
		this.SetTextLineSpacingFactor(codePair.Value.(DoubleCodePairValue).Value)
	case 42:
		this.SetActualMeasurement(codePair.Value.(DoubleCodePairValue).Value)
	case 1:
		this.SetText(codePair.Value.(StringCodePairValue).Value)
	case 53:
		this.SetTextRotationAngle(codePair.Value.(DoubleCodePairValue).Value)
	case 51:
		this.SetHorizontalDirectionAngle(codePair.Value.(DoubleCodePairValue).Value)
	case 210:
		temp := this.Normal()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 220:
		temp := this.Normal()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 230:
		temp := this.Normal()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 3:
		this.SetDimensionStyleName(codePair.Value.(StringCodePairValue).Value)
	default:
		return false
	}
	return true
}

func codePairsForDimension(this Dimension, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(100, "AcDbDimension"))
	if version >= R2010 {
		pairs = append(pairs, NewShortCodePair(280, int16(this.Version())))
	}
	pairs = append(pairs, NewStringCodePair(2, this.BlockName()))
	pairs = append(pairs, NewDoubleCodePair(10, this.DefinitionPoint1().X))
	pairs = append(pairs, NewDoubleCodePair(20, this.DefinitionPoint1().Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.DefinitionPoint1().Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.TextMidPoint().X))
	pairs = append(pairs, NewDoubleCodePair(21, this.TextMidPoint().Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.TextMidPoint().Z))
	pairs = append(pairs, NewShortCodePair(70, int16(this.DimensionType())))
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(71, int16(this.AttachmentPoint())))
	}
	if version >= R2000 && this.TextLineSpacingStyle() != TextLineSpacingStyleAtLeast {
		pairs = append(pairs, NewShortCodePair(72, int16(this.TextLineSpacingStyle())))
	}
	if version >= R2000 && this.TextLineSpacingFactor() != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.TextLineSpacingFactor()))
	}
	if version >= R2000 && this.ActualMeasurement() != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(42, this.ActualMeasurement()))
	}
	pairs = append(pairs, NewStringCodePair(1, this.Text()))
	if this.TextRotationAngle() != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(53, this.TextRotationAngle()))
	}
	if this.HorizontalDirectionAngle() != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(51, this.HorizontalDirectionAngle()))
	}
	if this.Normal() != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal().X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal().Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal().Z))
	}
	if version >= R12 {
		pairs = append(pairs, NewStringCodePair(3, this.DimensionStyleName()))
	}
	return
}

type RasterImage interface {
	subclassMarker() string
	SetsubclassMarker(val string)
	ClassVersion() int
	SetClassVersion(val int)
	Location() Point
	SetLocation(val Point)
	UVector() Vector
	SetUVector(val Vector)
	VVector() Vector
	SetVVector(val Vector)
	ImageSize() Vector
	SetImageSize(val Vector)
	imageDefinitionHandle() string
	SetimageDefinitionHandle(val string)
	DisplayOptionsFlags() int
	SetDisplayOptionsFlags(val int)
	UseClipping() bool
	SetUseClipping(val bool)
	Brightness() int16
	SetBrightness(val int16)
	Contrast() int16
	SetContrast(val int16)
	Fade() int16
	SetFade(val int16)
	imageDefinitionReactorHandle() string
	SetimageDefinitionReactorHandle(val string)
	ClippingType() ImageClippingBoundaryType
	SetClippingType(val ImageClippingBoundaryType)
	ClippingVertices() []Point
	SetClippingVertices(val []Point)
	clippingVertexCount() int
	SetclippingVertexCount(val int)
	clippingVerticesX() []float64
	SetclippingVerticesX(val []float64)
	clippingVerticesY() []float64
	SetclippingVerticesY(val []float64)
	IsInsideClipping() bool
	SetIsInsideClipping(val bool)
}

func tryApplyCodePairForRasterImage(this RasterImage, codePair CodePair) bool {
	switch codePair.Code {
	case 100:
		this.SetsubclassMarker(codePair.Value.(StringCodePairValue).Value)
	case 90:
		this.SetClassVersion(codePair.Value.(IntCodePairValue).Value)
	case 10:
		temp := this.Location()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetLocation(temp)
	case 20:
		temp := this.Location()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetLocation(temp)
	case 30:
		temp := this.Location()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetLocation(temp)
	case 11:
		temp := this.UVector()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetUVector(temp)
	case 21:
		temp := this.UVector()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetUVector(temp)
	case 31:
		temp := this.UVector()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetUVector(temp)
	case 12:
		temp := this.VVector()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetVVector(temp)
	case 22:
		temp := this.VVector()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetVVector(temp)
	case 32:
		temp := this.VVector()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetVVector(temp)
	case 13:
		temp := this.ImageSize()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetImageSize(temp)
	case 23:
		temp := this.ImageSize()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetImageSize(temp)
	case 340:
		this.SetimageDefinitionHandle(codePair.Value.(StringCodePairValue).Value)
	case 70:
		this.SetDisplayOptionsFlags(int(codePair.Value.(ShortCodePairValue).Value))
	case 280:
		this.SetUseClipping(boolFromShort(codePair.Value.(ShortCodePairValue).Value))
	case 281:
		this.SetBrightness(codePair.Value.(ShortCodePairValue).Value)
	case 282:
		this.SetContrast(codePair.Value.(ShortCodePairValue).Value)
	case 283:
		this.SetFade(codePair.Value.(ShortCodePairValue).Value)
	case 360:
		this.SetimageDefinitionReactorHandle(codePair.Value.(StringCodePairValue).Value)
	case 71:
		this.SetClippingType(ImageClippingBoundaryType(codePair.Value.(ShortCodePairValue).Value))
	case 91:
		this.SetclippingVertexCount(codePair.Value.(IntCodePairValue).Value)
	case 14:
		this.SetclippingVerticesX(append(this.clippingVerticesX(), codePair.Value.(DoubleCodePairValue).Value))
	case 24:
		this.SetclippingVerticesY(append(this.clippingVerticesY(), codePair.Value.(DoubleCodePairValue).Value))
	case 290:
		this.SetIsInsideClipping(codePair.Value.(BoolCodePairValue).Value)
	default:
		return false
	}
	return true
}

func codePairsForRasterImage(this RasterImage, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(100, this.subclassMarker()))
	pairs = append(pairs, NewIntCodePair(90, this.ClassVersion()))
	pairs = append(pairs, NewDoubleCodePair(10, this.Location().X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location().Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location().Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.UVector().X))
	pairs = append(pairs, NewDoubleCodePair(21, this.UVector().Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.UVector().Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.VVector().X))
	pairs = append(pairs, NewDoubleCodePair(22, this.VVector().Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.VVector().Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.ImageSize().X))
	pairs = append(pairs, NewDoubleCodePair(23, this.ImageSize().Y))
	if this.imageDefinitionHandle() != "" {
		pairs = append(pairs, NewStringCodePair(340, this.imageDefinitionHandle()))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(this.DisplayOptionsFlags())))
	pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.UseClipping())))
	pairs = append(pairs, NewShortCodePair(281, this.Brightness()))
	pairs = append(pairs, NewShortCodePair(282, this.Contrast()))
	pairs = append(pairs, NewShortCodePair(283, this.Fade()))
	if this.imageDefinitionReactorHandle() != "" {
		pairs = append(pairs, NewStringCodePair(360, this.imageDefinitionReactorHandle()))
	}
	pairs = append(pairs, NewShortCodePair(71, int16(this.ClippingType())))
	pairs = append(pairs, NewIntCodePair(91, len(this.ClippingVertices())))
	for _, item := range this.ClippingVertices() {
		pairs = append(pairs, NewDoubleCodePair(14, item.X))
		pairs = append(pairs, NewDoubleCodePair(24, item.Y))
	}
	if version >= R2010 {
		pairs = append(pairs, NewBoolCodePair(290, this.IsInsideClipping()))
	}
	return
}

type Underlay interface {
	ObjectHandle() string
	SetObjectHandle(val string)
	InsertionPoint() Point
	SetInsertionPoint(val Point)
	XScale() float64
	SetXScale(val float64)
	YScale() float64
	SetYScale(val float64)
	ZScale() float64
	SetZScale(val float64)
	RotationAngle() float64
	SetRotationAngle(val float64)
	Normal() Vector
	SetNormal(val Vector)
	Flags() int
	SetFlags(val int)
	Contrast() int16
	SetContrast(val int16)
	Fade() int16
	SetFade(val int16)
	BoundaryPoints() []Point
	SetBoundaryPoints(val []Point)
	pointX() []float64
	SetpointX(val []float64)
	pointY() []float64
	SetpointY(val []float64)
}

func tryApplyCodePairForUnderlay(this Underlay, codePair CodePair) bool {
	switch codePair.Code {
	case 340:
		this.SetObjectHandle(codePair.Value.(StringCodePairValue).Value)
	case 10:
		temp := this.InsertionPoint()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetInsertionPoint(temp)
	case 20:
		temp := this.InsertionPoint()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetInsertionPoint(temp)
	case 30:
		temp := this.InsertionPoint()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetInsertionPoint(temp)
	case 41:
		this.SetXScale(codePair.Value.(DoubleCodePairValue).Value)
	case 42:
		this.SetYScale(codePair.Value.(DoubleCodePairValue).Value)
	case 43:
		this.SetZScale(codePair.Value.(DoubleCodePairValue).Value)
	case 50:
		this.SetRotationAngle(codePair.Value.(DoubleCodePairValue).Value)
	case 210:
		temp := this.Normal()
		temp.X = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 220:
		temp := this.Normal()
		temp.Y = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 230:
		temp := this.Normal()
		temp.Z = codePair.Value.(DoubleCodePairValue).Value
		this.SetNormal(temp)
	case 280:
		this.SetFlags(int(codePair.Value.(ShortCodePairValue).Value))
	case 281:
		this.SetContrast(codePair.Value.(ShortCodePairValue).Value)
	case 282:
		this.SetFade(codePair.Value.(ShortCodePairValue).Value)
	case 11:
		this.SetpointX(append(this.pointX(), codePair.Value.(DoubleCodePairValue).Value))
	case 21:
		this.SetpointY(append(this.pointY(), codePair.Value.(DoubleCodePairValue).Value))
	default:
		return false
	}
	return true
}

func codePairsForUnderlay(this Underlay, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(100, "AcDbUnderlayReference"))
	pairs = append(pairs, NewStringCodePair(340, this.ObjectHandle()))
	pairs = append(pairs, NewDoubleCodePair(10, this.InsertionPoint().X))
	pairs = append(pairs, NewDoubleCodePair(20, this.InsertionPoint().Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.InsertionPoint().Z))
	pairs = append(pairs, NewDoubleCodePair(41, this.XScale()))
	pairs = append(pairs, NewDoubleCodePair(42, this.YScale()))
	pairs = append(pairs, NewDoubleCodePair(43, this.ZScale()))
	pairs = append(pairs, NewDoubleCodePair(50, this.RotationAngle()))
	pairs = append(pairs, NewDoubleCodePair(210, this.Normal().X))
	pairs = append(pairs, NewDoubleCodePair(220, this.Normal().Y))
	pairs = append(pairs, NewDoubleCodePair(230, this.Normal().Z))
	pairs = append(pairs, NewShortCodePair(280, int16(this.Flags())))
	pairs = append(pairs, NewShortCodePair(281, this.Contrast()))
	pairs = append(pairs, NewShortCodePair(282, this.Fade()))
	for _, item := range this.BoundaryPoints() {
		pairs = append(pairs, NewDoubleCodePair(11, item.X))
		pairs = append(pairs, NewDoubleCodePair(21, item.Y))
	}
	return
}

type Face struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FirstCorner Point
	SecondCorner Point
	ThirdCorner Point
	FourthCorner Point
	EdgeFlags int
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewFace() *Face {
	return &Face{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FirstCorner: *NewOrigin(),
		SecondCorner: *NewOrigin(),
		ThirdCorner: *NewOrigin(),
		FourthCorner: *NewOrigin(),
		EdgeFlags: 0,
	}
}

func (e *Face) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Face) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Face) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Face) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Face) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Face) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Face) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Face) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Face) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Face) Handle() Handle {
	return this.handle
}

func (this *Face) SetHandle(val Handle) {
	this.handle = val
}

func (this *Face) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Face) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Face) Layer() string {
	return this.layer
}

func (this *Face) SetLayer(val string) {
	this.layer = val
}

func (this *Face) LineTypeName() string {
	return this.lineTypeName
}

func (this *Face) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Face) Elevation() float64 {
	return this.elevation
}

func (this *Face) SetElevation(val float64) {
	this.elevation = val
}

func (this *Face) MaterialHandle() string {
	return this.materialHandle
}

func (this *Face) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Face) Color() Color {
	return this.color
}

func (this *Face) SetColor(val Color) {
	this.color = val
}

func (this *Face) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Face) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Face) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Face) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Face) IsVisible() bool {
	return this.isVisible
}

func (this *Face) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Face) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Face) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Face) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Face) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Face) Color24Bit() int {
	return this.color24Bit
}

func (this *Face) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Face) ColorName() string {
	return this.colorName
}

func (this *Face) SetColorName(val string) {
	this.colorName = val
}

func (this *Face) Transparency() int {
	return this.transparency
}

func (this *Face) SetTransparency(val int) {
	this.transparency = val
}

func (this *Face) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Face) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// FirstEdgeInvisible status flag.
func (this *Face) FirstEdgeInvisible() bool {
	return this.EdgeFlags & 1 != 0
}

// FirstEdgeInvisible status flag.
func (this *Face) SetFirstEdgeInvisible(val bool) {
	if val {
		this.EdgeFlags = this.EdgeFlags | 1
	} else {
		this.EdgeFlags = this.EdgeFlags & ^1
	}
}

// SecondEdgeInvisible status flag.
func (this *Face) SecondEdgeInvisible() bool {
	return this.EdgeFlags & 2 != 0
}

// SecondEdgeInvisible status flag.
func (this *Face) SetSecondEdgeInvisible(val bool) {
	if val {
		this.EdgeFlags = this.EdgeFlags | 2
	} else {
		this.EdgeFlags = this.EdgeFlags & ^2
	}
}

// ThirdEdgeInvisible status flag.
func (this *Face) ThirdEdgeInvisible() bool {
	return this.EdgeFlags & 4 != 0
}

// ThirdEdgeInvisible status flag.
func (this *Face) SetThirdEdgeInvisible(val bool) {
	if val {
		this.EdgeFlags = this.EdgeFlags | 4
	} else {
		this.EdgeFlags = this.EdgeFlags & ^4
	}
}

// FourthEdgeInvisible status flag.
func (this *Face) FourthEdgeInvisible() bool {
	return this.EdgeFlags & 8 != 0
}

// FourthEdgeInvisible status flag.
func (this *Face) SetFourthEdgeInvisible(val bool) {
	if val {
		this.EdgeFlags = this.EdgeFlags | 8
	} else {
		this.EdgeFlags = this.EdgeFlags & ^8
	}
}

func (this *Face) typeString() string {
	return "3DFACE"
}

func (this *Face) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Face) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Face) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.FirstCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.FirstCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.FirstCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.SecondCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.SecondCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.SecondCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.ThirdCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.ThirdCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.ThirdCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 13:
		this.FourthCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.FourthCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.FourthCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 70:
		this.EdgeFlags = int(codePair.Value.(ShortCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Face) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "3DFACE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbFace"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.FirstCorner.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.FirstCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.FirstCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.SecondCorner.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.SecondCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.SecondCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.ThirdCorner.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.ThirdCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.ThirdCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.FourthCorner.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.FourthCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.FourthCorner.Z))
	if this.EdgeFlags != 0 {
		pairs = append(pairs, NewShortCodePair(70, int16(this.EdgeFlags)))
	}
	return
}

type Solid3D struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FormatVersionNumber int16
	CustomData []string
	CustomData2 []string
	pointerOwner pointer
	pointerPlotStyle pointer
	pointerHistoryObject pointer
}

func NewSolid3D() *Solid3D {
	return &Solid3D{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FormatVersionNumber: 1,
		CustomData: []string{},
		CustomData2: []string{},
	}
}

func (e *Solid3D) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	pointers = append(pointers, &e.pointerHistoryObject)
	return
}

func (e *Solid3D) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Solid3D) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Solid3D) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Solid3D) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Solid3D) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Solid3D) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Solid3D) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Solid3D) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (e *Solid3D) HistoryObject() *DrawingItem {
	return e.pointerHistoryObject.value
}

func (e *Solid3D) SetHistoryObject(val *DrawingItem) {
	e.pointerHistoryObject.value = val
}

func (this *Solid3D) Handle() Handle {
	return this.handle
}

func (this *Solid3D) SetHandle(val Handle) {
	this.handle = val
}

func (this *Solid3D) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Solid3D) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Solid3D) Layer() string {
	return this.layer
}

func (this *Solid3D) SetLayer(val string) {
	this.layer = val
}

func (this *Solid3D) LineTypeName() string {
	return this.lineTypeName
}

func (this *Solid3D) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Solid3D) Elevation() float64 {
	return this.elevation
}

func (this *Solid3D) SetElevation(val float64) {
	this.elevation = val
}

func (this *Solid3D) MaterialHandle() string {
	return this.materialHandle
}

func (this *Solid3D) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Solid3D) Color() Color {
	return this.color
}

func (this *Solid3D) SetColor(val Color) {
	this.color = val
}

func (this *Solid3D) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Solid3D) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Solid3D) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Solid3D) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Solid3D) IsVisible() bool {
	return this.isVisible
}

func (this *Solid3D) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Solid3D) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Solid3D) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Solid3D) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Solid3D) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Solid3D) Color24Bit() int {
	return this.color24Bit
}

func (this *Solid3D) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Solid3D) ColorName() string {
	return this.colorName
}

func (this *Solid3D) SetColorName(val string) {
	this.colorName = val
}

func (this *Solid3D) Transparency() int {
	return this.transparency
}

func (this *Solid3D) SetTransparency(val int) {
	this.transparency = val
}

func (this *Solid3D) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Solid3D) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Solid3D) AddCustomData(val string) {
	this.CustomData = append(this.CustomData, val)
}

func (this *Solid3D) ClearCustomData() {
	this.CustomData = []string{}
}

func (this *Solid3D) AddCustomData2(val string) {
	this.CustomData2 = append(this.CustomData2, val)
}

func (this *Solid3D) ClearCustomData2() {
	this.CustomData2 = []string{}
}

func (this *Solid3D) typeString() string {
	return "3DSOLID"
}

func (this *Solid3D) minVersion() (version AcadVersion) {
	return R13
}

func (this *Solid3D) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Solid3D) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 70:
		this.FormatVersionNumber = codePair.Value.(ShortCodePairValue).Value
	case 1:
		this.CustomData = append(this.CustomData, codePair.Value.(StringCodePairValue).Value)
	case 3:
		this.CustomData2 = append(this.CustomData2, codePair.Value.(StringCodePairValue).Value)
	case 350:
		this.pointerHistoryObject.handle = handleFromString(codePair.Value.(StringCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Solid3D) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "3DSOLID"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbModelerGeometry"))
	pairs = append(pairs, NewShortCodePair(70, this.FormatVersionNumber))
	for _, val := range this.CustomData {
		pairs = append(pairs, NewStringCodePair(1, val))
	}
	for _, val := range this.CustomData2 {
		pairs = append(pairs, NewStringCodePair(3, val))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(100, "AcDb3dSolid"))
	}
	if this.pointerHistoryObject.handle != 0 {
		pairs = append(pairs, NewStringCodePair(350, stringFromHandle(this.pointerHistoryObject.handle)))
	}
	return
}

type ProxyEntity struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	readingGraphicsData bool
	ProxyEntityClassId int
	ApplicationEntityClassId int
	GraphicsData []byte
	graphicsDataSize int
	graphicsDataString []string
	EntityData []byte
	entityDataSize int
	entityDataString []string
	ObjectID1 []string
	ObjectID2 []string
	ObjectID3 []string
	ObjectID4 []string
	Terminator int
	objectDrawingFormat uint
	OriginalDataFormatIsDxf bool
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewProxyEntity() *ProxyEntity {
	return &ProxyEntity{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		readingGraphicsData: true,
		ProxyEntityClassId: 498,
		ApplicationEntityClassId: 500,
		GraphicsData: []byte{},
		graphicsDataSize: 0,
		graphicsDataString: []string{},
		EntityData: []byte{},
		entityDataSize: 0,
		entityDataString: []string{},
		ObjectID1: []string{},
		ObjectID2: []string{},
		ObjectID3: []string{},
		ObjectID4: []string{},
		Terminator: 0,
		objectDrawingFormat: 0,
		OriginalDataFormatIsDxf: true,
	}
}

func (e *ProxyEntity) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *ProxyEntity) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *ProxyEntity) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *ProxyEntity) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *ProxyEntity) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *ProxyEntity) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *ProxyEntity) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *ProxyEntity) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *ProxyEntity) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *ProxyEntity) Handle() Handle {
	return this.handle
}

func (this *ProxyEntity) SetHandle(val Handle) {
	this.handle = val
}

func (this *ProxyEntity) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *ProxyEntity) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *ProxyEntity) Layer() string {
	return this.layer
}

func (this *ProxyEntity) SetLayer(val string) {
	this.layer = val
}

func (this *ProxyEntity) LineTypeName() string {
	return this.lineTypeName
}

func (this *ProxyEntity) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *ProxyEntity) Elevation() float64 {
	return this.elevation
}

func (this *ProxyEntity) SetElevation(val float64) {
	this.elevation = val
}

func (this *ProxyEntity) MaterialHandle() string {
	return this.materialHandle
}

func (this *ProxyEntity) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *ProxyEntity) Color() Color {
	return this.color
}

func (this *ProxyEntity) SetColor(val Color) {
	this.color = val
}

func (this *ProxyEntity) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *ProxyEntity) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *ProxyEntity) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *ProxyEntity) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *ProxyEntity) IsVisible() bool {
	return this.isVisible
}

func (this *ProxyEntity) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *ProxyEntity) ImageByteCount() int {
	return this.imageByteCount
}

func (this *ProxyEntity) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *ProxyEntity) PreviewImageData() []string {
	return this.previewImageData
}

func (this *ProxyEntity) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *ProxyEntity) Color24Bit() int {
	return this.color24Bit
}

func (this *ProxyEntity) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *ProxyEntity) ColorName() string {
	return this.colorName
}

func (this *ProxyEntity) SetColorName(val string) {
	this.colorName = val
}

func (this *ProxyEntity) Transparency() int {
	return this.transparency
}

func (this *ProxyEntity) SetTransparency(val int) {
	this.transparency = val
}

func (this *ProxyEntity) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *ProxyEntity) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *ProxyEntity) AddgraphicsDataString(val string) {
	this.graphicsDataString = append(this.graphicsDataString, val)
}

func (this *ProxyEntity) CleargraphicsDataString() {
	this.graphicsDataString = []string{}
}

func (this *ProxyEntity) AddentityDataString(val string) {
	this.entityDataString = append(this.entityDataString, val)
}

func (this *ProxyEntity) ClearentityDataString() {
	this.entityDataString = []string{}
}

func (this *ProxyEntity) AddObjectID1(val string) {
	this.ObjectID1 = append(this.ObjectID1, val)
}

func (this *ProxyEntity) ClearObjectID1() {
	this.ObjectID1 = []string{}
}

func (this *ProxyEntity) AddObjectID2(val string) {
	this.ObjectID2 = append(this.ObjectID2, val)
}

func (this *ProxyEntity) ClearObjectID2() {
	this.ObjectID2 = []string{}
}

func (this *ProxyEntity) AddObjectID3(val string) {
	this.ObjectID3 = append(this.ObjectID3, val)
}

func (this *ProxyEntity) ClearObjectID3() {
	this.ObjectID3 = []string{}
}

func (this *ProxyEntity) AddObjectID4(val string) {
	this.ObjectID4 = append(this.ObjectID4, val)
}

func (this *ProxyEntity) ClearObjectID4() {
	this.ObjectID4 = []string{}
}

func (this *ProxyEntity) typeString() string {
	return "ACAD_PROXY_ENTITY"
}

func (this *ProxyEntity) minVersion() (version AcadVersion) {
	return R14
}

func (this *ProxyEntity) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *ProxyEntity) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ACAD_PROXY_ENTITY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbProxyEntity"))
	pairs = append(pairs, NewIntCodePair(90, this.ProxyEntityClassId))
	pairs = append(pairs, NewIntCodePair(91, this.ApplicationEntityClassId))
	pairs = append(pairs, NewIntCodePair(92, len(this.GraphicsData)))
	for _, val := range this.graphicsDataString {
		pairs = append(pairs, NewStringCodePair(310, val))
	}
	pairs = append(pairs, NewIntCodePair(93, len(this.EntityData)))
	for _, val := range this.entityDataString {
		pairs = append(pairs, NewStringCodePair(310, val))
	}
	for _, val := range this.ObjectID1 {
		pairs = append(pairs, NewStringCodePair(330, val))
	}
	for _, val := range this.ObjectID2 {
		pairs = append(pairs, NewStringCodePair(340, val))
	}
	for _, val := range this.ObjectID3 {
		pairs = append(pairs, NewStringCodePair(350, val))
	}
	for _, val := range this.ObjectID4 {
		pairs = append(pairs, NewStringCodePair(360, val))
	}
	pairs = append(pairs, NewIntCodePair(94, this.Terminator))
	if version >= R2000 {
		pairs = append(pairs, NewIntCodePair(95, int(this.objectDrawingFormat)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(this.OriginalDataFormatIsDxf)))
	}
	return
}

type Arc struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Center Point
	Radius float64
	Normal Vector
	StartAngle float64 // Arc start angle in degrees
	EndAngle float64 // Arc end angle in degrees
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewArc() *Arc {
	return &Arc{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Center: *NewOrigin(),
		Radius: 0.0,
		Normal: *NewZAxis(),
		StartAngle: 0.0,
		EndAngle: 360.0,
	}
}

func (e *Arc) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Arc) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Arc) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Arc) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Arc) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Arc) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Arc) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Arc) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Arc) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Arc) Handle() Handle {
	return this.handle
}

func (this *Arc) SetHandle(val Handle) {
	this.handle = val
}

func (this *Arc) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Arc) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Arc) Layer() string {
	return this.layer
}

func (this *Arc) SetLayer(val string) {
	this.layer = val
}

func (this *Arc) LineTypeName() string {
	return this.lineTypeName
}

func (this *Arc) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Arc) Elevation() float64 {
	return this.elevation
}

func (this *Arc) SetElevation(val float64) {
	this.elevation = val
}

func (this *Arc) MaterialHandle() string {
	return this.materialHandle
}

func (this *Arc) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Arc) Color() Color {
	return this.color
}

func (this *Arc) SetColor(val Color) {
	this.color = val
}

func (this *Arc) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Arc) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Arc) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Arc) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Arc) IsVisible() bool {
	return this.isVisible
}

func (this *Arc) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Arc) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Arc) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Arc) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Arc) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Arc) Color24Bit() int {
	return this.color24Bit
}

func (this *Arc) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Arc) ColorName() string {
	return this.colorName
}

func (this *Arc) SetColorName(val string) {
	this.colorName = val
}

func (this *Arc) Transparency() int {
	return this.transparency
}

func (this *Arc) SetTransparency(val int) {
	this.transparency = val
}

func (this *Arc) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Arc) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Arc) typeString() string {
	return "ARC"
}

func (this *Arc) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Arc) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Arc) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.Center.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Center.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Center.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.Radius = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.StartAngle = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.EndAngle = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Arc) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ARC"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbCircle"))
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Center.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Center.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Center.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.Radius))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	pairs = append(pairs, NewStringCodePair(100, "AcDbArc"))
	pairs = append(pairs, NewDoubleCodePair(50, this.StartAngle))
	pairs = append(pairs, NewDoubleCodePair(51, this.EndAngle))
	return
}

type ArcAlignedText struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Text string
	FontName string
	BigfontName string
	TextStyleName string
	CenterPoint Point
	ArcRadius float64
	WidthFactor float64
	TextHeight float64
	CharacterSpacing float64
	OffsetFromArc float64
	RightOffset float64
	LeftOffset float64
	StartAngle float64
	EndAngle float64
	IsCharacterOrderReversed bool
	DirectionFlag int16
	AlignmentFlag int16
	SideFlag int16
	IsBold bool
	IsItalic bool
	IsUnderline bool
	CharacterSetValue int16
	PitchAndFamilyValue int16
	FontType FontType
	ColorIndex int
	ExtrusionDirection Vector
	WizardFlag int16
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewArcAlignedText() *ArcAlignedText {
	return &ArcAlignedText{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Text: "",
		FontName: "",
		BigfontName: "",
		TextStyleName: "",
		CenterPoint: *NewOrigin(),
		ArcRadius: 0.0,
		WidthFactor: 1.0,
		TextHeight: 0.0,
		CharacterSpacing: 0.0,
		OffsetFromArc: 0.0,
		RightOffset: 0.0,
		LeftOffset: 0.0,
		StartAngle: 0.0,
		EndAngle: 0.0,
		IsCharacterOrderReversed: false,
		DirectionFlag: 0,
		AlignmentFlag: 0,
		SideFlag: 0,
		IsBold: false,
		IsItalic: false,
		IsUnderline: false,
		CharacterSetValue: 0,
		PitchAndFamilyValue: 0,
		FontType: FontTypeTTF,
		ColorIndex: 0,
		ExtrusionDirection: *NewZAxis(),
		WizardFlag: 0,
	}
}

func (e *ArcAlignedText) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *ArcAlignedText) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *ArcAlignedText) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *ArcAlignedText) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *ArcAlignedText) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *ArcAlignedText) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *ArcAlignedText) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *ArcAlignedText) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *ArcAlignedText) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *ArcAlignedText) Handle() Handle {
	return this.handle
}

func (this *ArcAlignedText) SetHandle(val Handle) {
	this.handle = val
}

func (this *ArcAlignedText) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *ArcAlignedText) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *ArcAlignedText) Layer() string {
	return this.layer
}

func (this *ArcAlignedText) SetLayer(val string) {
	this.layer = val
}

func (this *ArcAlignedText) LineTypeName() string {
	return this.lineTypeName
}

func (this *ArcAlignedText) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *ArcAlignedText) Elevation() float64 {
	return this.elevation
}

func (this *ArcAlignedText) SetElevation(val float64) {
	this.elevation = val
}

func (this *ArcAlignedText) MaterialHandle() string {
	return this.materialHandle
}

func (this *ArcAlignedText) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *ArcAlignedText) Color() Color {
	return this.color
}

func (this *ArcAlignedText) SetColor(val Color) {
	this.color = val
}

func (this *ArcAlignedText) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *ArcAlignedText) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *ArcAlignedText) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *ArcAlignedText) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *ArcAlignedText) IsVisible() bool {
	return this.isVisible
}

func (this *ArcAlignedText) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *ArcAlignedText) ImageByteCount() int {
	return this.imageByteCount
}

func (this *ArcAlignedText) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *ArcAlignedText) PreviewImageData() []string {
	return this.previewImageData
}

func (this *ArcAlignedText) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *ArcAlignedText) Color24Bit() int {
	return this.color24Bit
}

func (this *ArcAlignedText) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *ArcAlignedText) ColorName() string {
	return this.colorName
}

func (this *ArcAlignedText) SetColorName(val string) {
	this.colorName = val
}

func (this *ArcAlignedText) Transparency() int {
	return this.transparency
}

func (this *ArcAlignedText) SetTransparency(val int) {
	this.transparency = val
}

func (this *ArcAlignedText) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *ArcAlignedText) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *ArcAlignedText) typeString() string {
	return "ARCALIGNEDTEXT"
}

func (this *ArcAlignedText) minVersion() (version AcadVersion) {
	return R2000
}

func (this *ArcAlignedText) maxVersion() (version AcadVersion) {
	return R2000
}

func (this *ArcAlignedText) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 1:
		this.Text = codePair.Value.(StringCodePairValue).Value
	case 2:
		this.FontName = codePair.Value.(StringCodePairValue).Value
	case 3:
		this.BigfontName = codePair.Value.(StringCodePairValue).Value
	case 7:
		this.TextStyleName = codePair.Value.(StringCodePairValue).Value
	case 10:
		this.CenterPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.CenterPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.CenterPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.ArcRadius = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.WidthFactor = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.TextHeight = codePair.Value.(DoubleCodePairValue).Value
	case 43:
		this.CharacterSpacing = codePair.Value.(DoubleCodePairValue).Value
	case 44:
		this.OffsetFromArc = codePair.Value.(DoubleCodePairValue).Value
	case 45:
		this.RightOffset = codePair.Value.(DoubleCodePairValue).Value
	case 46:
		this.LeftOffset = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.StartAngle = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.EndAngle = codePair.Value.(DoubleCodePairValue).Value
	case 70:
		this.IsCharacterOrderReversed = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 71:
		this.DirectionFlag = codePair.Value.(ShortCodePairValue).Value
	case 72:
		this.AlignmentFlag = codePair.Value.(ShortCodePairValue).Value
	case 73:
		this.SideFlag = codePair.Value.(ShortCodePairValue).Value
	case 74:
		this.IsBold = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 75:
		this.IsItalic = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 76:
		this.IsUnderline = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 77:
		this.CharacterSetValue = codePair.Value.(ShortCodePairValue).Value
	case 78:
		this.PitchAndFamilyValue = codePair.Value.(ShortCodePairValue).Value
	case 79:
		this.FontType = FontType(codePair.Value.(ShortCodePairValue).Value)
	case 90:
		this.ColorIndex = codePair.Value.(IntCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 280:
		this.WizardFlag = codePair.Value.(ShortCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *ArcAlignedText) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ARCALIGNEDTEXT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbArcAlignedText"))
	}
	pairs = append(pairs, NewStringCodePair(1, this.Text))
	pairs = append(pairs, NewStringCodePair(2, this.FontName))
	pairs = append(pairs, NewStringCodePair(3, this.BigfontName))
	pairs = append(pairs, NewStringCodePair(7, this.TextStyleName))
	pairs = append(pairs, NewDoubleCodePair(10, this.CenterPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.CenterPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.CenterPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.ArcRadius))
	pairs = append(pairs, NewDoubleCodePair(41, this.WidthFactor))
	pairs = append(pairs, NewDoubleCodePair(42, this.TextHeight))
	pairs = append(pairs, NewDoubleCodePair(43, this.CharacterSpacing))
	pairs = append(pairs, NewDoubleCodePair(44, this.OffsetFromArc))
	pairs = append(pairs, NewDoubleCodePair(45, this.RightOffset))
	pairs = append(pairs, NewDoubleCodePair(46, this.LeftOffset))
	pairs = append(pairs, NewDoubleCodePair(50, this.StartAngle))
	pairs = append(pairs, NewDoubleCodePair(51, this.EndAngle))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(this.IsCharacterOrderReversed)))
	pairs = append(pairs, NewShortCodePair(71, this.DirectionFlag))
	pairs = append(pairs, NewShortCodePair(72, this.AlignmentFlag))
	pairs = append(pairs, NewShortCodePair(73, this.SideFlag))
	pairs = append(pairs, NewShortCodePair(74, shortFromBool(this.IsBold)))
	pairs = append(pairs, NewShortCodePair(75, shortFromBool(this.IsItalic)))
	pairs = append(pairs, NewShortCodePair(76, shortFromBool(this.IsUnderline)))
	pairs = append(pairs, NewShortCodePair(77, this.CharacterSetValue))
	pairs = append(pairs, NewShortCodePair(78, this.PitchAndFamilyValue))
	pairs = append(pairs, NewShortCodePair(79, int16(this.FontType)))
	pairs = append(pairs, NewIntCodePair(90, this.ColorIndex))
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	pairs = append(pairs, NewShortCodePair(280, this.WizardFlag))
	return
}

type AttributeDefinition struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Location Point
	TextHeight float64
	Value string
	Rotation float64
	RelativeXScaleFactor float64
	ObliqueAngle float64
	TextStyleName string
	TextGenerationFlags int
	HorizontalTextJustification HorizontalTextJustification
	SecondAlignmentPoint Point
	Normal Vector
	Version Version
	Prompt string
	TextTag string
	Flags int
	FieldLength int16
	VerticalTextJustification VerticalTextJustification
	IsLockedInBlock bool
	KeepDuplicateRecords bool
	MTextFlag MTextFlag
	IsReallyLocked bool
	secondaryAttributeHandlesCount int
	SecondaryAttributeHandles []string
	AlignmentPoint Point
	AnnotationScale float64
	XRecordTag string
	MText MText
	lastSubclassMarker string
	isVersionSet bool
	xrecCode70Count int
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewAttributeDefinition() *AttributeDefinition {
	return &AttributeDefinition{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Location: *NewOrigin(),
		TextHeight: 1.0,
		Value: "",
		Rotation: 0.0,
		RelativeXScaleFactor: 1.0,
		ObliqueAngle: 0.0,
		TextStyleName: "STANDARD",
		TextGenerationFlags: 0,
		HorizontalTextJustification: HorizontalTextJustificationLeft,
		SecondAlignmentPoint: *NewOrigin(),
		Normal: *NewZAxis(),
		Version: VersionR2010,
		Prompt: "",
		TextTag: "",
		Flags: 0,
		FieldLength: 0,
		VerticalTextJustification: VerticalTextJustificationBaseline,
		IsLockedInBlock: false,
		KeepDuplicateRecords: false,
		MTextFlag: MTextFlagMultilineAttribute,
		IsReallyLocked: false,
		secondaryAttributeHandlesCount: 0,
		SecondaryAttributeHandles: []string{},
		AlignmentPoint: *NewOrigin(),
		AnnotationScale: 1.0,
		XRecordTag: "",
		MText: *NewMText(),
		lastSubclassMarker: "",
		isVersionSet: false,
		xrecCode70Count: 0,
	}
}

func (e *AttributeDefinition) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *AttributeDefinition) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *AttributeDefinition) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *AttributeDefinition) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *AttributeDefinition) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *AttributeDefinition) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *AttributeDefinition) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *AttributeDefinition) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *AttributeDefinition) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *AttributeDefinition) Handle() Handle {
	return this.handle
}

func (this *AttributeDefinition) SetHandle(val Handle) {
	this.handle = val
}

func (this *AttributeDefinition) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *AttributeDefinition) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *AttributeDefinition) Layer() string {
	return this.layer
}

func (this *AttributeDefinition) SetLayer(val string) {
	this.layer = val
}

func (this *AttributeDefinition) LineTypeName() string {
	return this.lineTypeName
}

func (this *AttributeDefinition) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *AttributeDefinition) Elevation() float64 {
	return this.elevation
}

func (this *AttributeDefinition) SetElevation(val float64) {
	this.elevation = val
}

func (this *AttributeDefinition) MaterialHandle() string {
	return this.materialHandle
}

func (this *AttributeDefinition) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *AttributeDefinition) Color() Color {
	return this.color
}

func (this *AttributeDefinition) SetColor(val Color) {
	this.color = val
}

func (this *AttributeDefinition) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *AttributeDefinition) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *AttributeDefinition) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *AttributeDefinition) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *AttributeDefinition) IsVisible() bool {
	return this.isVisible
}

func (this *AttributeDefinition) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *AttributeDefinition) ImageByteCount() int {
	return this.imageByteCount
}

func (this *AttributeDefinition) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *AttributeDefinition) PreviewImageData() []string {
	return this.previewImageData
}

func (this *AttributeDefinition) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *AttributeDefinition) Color24Bit() int {
	return this.color24Bit
}

func (this *AttributeDefinition) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *AttributeDefinition) ColorName() string {
	return this.colorName
}

func (this *AttributeDefinition) SetColorName(val string) {
	this.colorName = val
}

func (this *AttributeDefinition) Transparency() int {
	return this.transparency
}

func (this *AttributeDefinition) SetTransparency(val int) {
	this.transparency = val
}

func (this *AttributeDefinition) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *AttributeDefinition) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsTextBackward status flag.
func (this *AttributeDefinition) IsTextBackward() bool {
	return this.TextGenerationFlags & 2 != 0
}

// IsTextBackward status flag.
func (this *AttributeDefinition) SetIsTextBackward(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 2
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^2
	}
}

// IsTextUpsideDown status flag.
func (this *AttributeDefinition) IsTextUpsideDown() bool {
	return this.TextGenerationFlags & 4 != 0
}

// IsTextUpsideDown status flag.
func (this *AttributeDefinition) SetIsTextUpsideDown(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 4
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^4
	}
}

// IsInvisible status flag.
func (this *AttributeDefinition) IsInvisible() bool {
	return this.Flags & 1 != 0
}

// IsInvisible status flag.
func (this *AttributeDefinition) SetIsInvisible(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// IsConstant status flag.
func (this *AttributeDefinition) IsConstant() bool {
	return this.Flags & 2 != 0
}

// IsConstant status flag.
func (this *AttributeDefinition) SetIsConstant(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// IsInputVerificationRequired status flag.
func (this *AttributeDefinition) IsInputVerificationRequired() bool {
	return this.Flags & 4 != 0
}

// IsInputVerificationRequired status flag.
func (this *AttributeDefinition) SetIsInputVerificationRequired(val bool) {
	if val {
		this.Flags = this.Flags | 4
	} else {
		this.Flags = this.Flags & ^4
	}
}

// IsAttributePresent status flag.
func (this *AttributeDefinition) IsAttributePresent() bool {
	return this.Flags & 8 != 0
}

// IsAttributePresent status flag.
func (this *AttributeDefinition) SetIsAttributePresent(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

func (this *AttributeDefinition) AddSecondaryAttributeHandles(val string) {
	this.SecondaryAttributeHandles = append(this.SecondaryAttributeHandles, val)
}

func (this *AttributeDefinition) ClearSecondaryAttributeHandles() {
	this.SecondaryAttributeHandles = []string{}
}

func (this *AttributeDefinition) typeString() string {
	return "ATTDEF"
}

func (this *AttributeDefinition) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *AttributeDefinition) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *AttributeDefinition) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ATTDEF"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbText"))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.TextHeight))
	pairs = append(pairs, NewStringCodePair(1, this.Value))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbAttributeDefinition"))
	}
	if this.Rotation != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.Rotation))
	}
	if this.RelativeXScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.RelativeXScaleFactor))
	}
	if this.ObliqueAngle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(51, this.ObliqueAngle))
	}
	if this.TextStyleName != "STANDARD" {
		pairs = append(pairs, NewStringCodePair(7, this.TextStyleName))
	}
	if this.TextGenerationFlags != 0 {
		pairs = append(pairs, NewShortCodePair(71, int16(this.TextGenerationFlags)))
	}
	if this.HorizontalTextJustification != HorizontalTextJustificationLeft {
		pairs = append(pairs, NewShortCodePair(72, int16(this.HorizontalTextJustification)))
	}
	pairs = append(pairs, NewDoubleCodePair(11, this.SecondAlignmentPoint.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.SecondAlignmentPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.SecondAlignmentPoint.Z))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbAttributeDefinition"))
	}
	if version >= R2010 {
		pairs = append(pairs, NewShortCodePair(280, int16(this.Version)))
	}
	pairs = append(pairs, NewStringCodePair(3, this.Prompt))
	pairs = append(pairs, NewStringCodePair(2, this.TextTag))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	if this.FieldLength != 0 {
		pairs = append(pairs, NewShortCodePair(73, this.FieldLength))
	}
	if version >= R12 && this.VerticalTextJustification != VerticalTextJustificationBaseline {
		pairs = append(pairs, NewShortCodePair(74, int16(this.VerticalTextJustification)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.IsLockedInBlock)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbXrecord"))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.KeepDuplicateRecords)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, int16(this.MTextFlag)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(this.IsReallyLocked)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, int16(len(this.SecondaryAttributeHandles))))
	}
	if version >= R2007 {
		for _, val := range this.SecondaryAttributeHandles {
			pairs = append(pairs, NewStringCodePair(340, val))
		}
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(10, this.AlignmentPoint.X))
		pairs = append(pairs, NewDoubleCodePair(20, this.AlignmentPoint.Y))
		pairs = append(pairs, NewDoubleCodePair(30, this.AlignmentPoint.Z))
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(40, this.AnnotationScale))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(2, this.XRecordTag))
	}
	return
}

type Attribute struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Location Point
	TextHeight float64
	Value string
	Version Version
	AttributeTag string
	Flags int
	FieldLength int16
	Rotation float64
	RelativeXScaleFactor float64
	ObliqueAngle float64
	TextStyleName string
	TextGenerationFlags int
	HorizontalTextJustification HorizontalTextJustification
	VerticalTextJustification VerticalTextJustification
	SecondAlignmentPoint Point
	Normal Vector
	IsLockedInBlock bool
	KeepDuplicateRecords bool
	MTextFlag MTextFlag
	IsReallyLocked bool
	secondaryAttributeCount int
	secondaryAttributeHandles []string
	AlignmentPoint Point
	AnnotationScale float64
	XRecordTag string
	MText MText
	lastSubclassMarker string
	isVersionSet bool
	xrecCode70Count int
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewAttribute() *Attribute {
	return &Attribute{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Location: *NewOrigin(),
		TextHeight: 1.0,
		Value: "",
		Version: VersionR2010,
		AttributeTag: "",
		Flags: 0,
		FieldLength: 0,
		Rotation: 0.0,
		RelativeXScaleFactor: 1.0,
		ObliqueAngle: 0.0,
		TextStyleName: "STANDARD",
		TextGenerationFlags: 0,
		HorizontalTextJustification: HorizontalTextJustificationLeft,
		VerticalTextJustification: VerticalTextJustificationBaseline,
		SecondAlignmentPoint: *NewOrigin(),
		Normal: *NewZAxis(),
		IsLockedInBlock: false,
		KeepDuplicateRecords: false,
		MTextFlag: MTextFlagMultilineAttribute,
		IsReallyLocked: false,
		secondaryAttributeCount: 0,
		secondaryAttributeHandles: []string{},
		AlignmentPoint: *NewOrigin(),
		AnnotationScale: 1.0,
		XRecordTag: "",
		MText: *NewMText(),
		lastSubclassMarker: "",
		isVersionSet: false,
		xrecCode70Count: 0,
	}
}

func (e *Attribute) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Attribute) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Attribute) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Attribute) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Attribute) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Attribute) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Attribute) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Attribute) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Attribute) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Attribute) Handle() Handle {
	return this.handle
}

func (this *Attribute) SetHandle(val Handle) {
	this.handle = val
}

func (this *Attribute) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Attribute) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Attribute) Layer() string {
	return this.layer
}

func (this *Attribute) SetLayer(val string) {
	this.layer = val
}

func (this *Attribute) LineTypeName() string {
	return this.lineTypeName
}

func (this *Attribute) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Attribute) Elevation() float64 {
	return this.elevation
}

func (this *Attribute) SetElevation(val float64) {
	this.elevation = val
}

func (this *Attribute) MaterialHandle() string {
	return this.materialHandle
}

func (this *Attribute) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Attribute) Color() Color {
	return this.color
}

func (this *Attribute) SetColor(val Color) {
	this.color = val
}

func (this *Attribute) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Attribute) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Attribute) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Attribute) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Attribute) IsVisible() bool {
	return this.isVisible
}

func (this *Attribute) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Attribute) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Attribute) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Attribute) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Attribute) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Attribute) Color24Bit() int {
	return this.color24Bit
}

func (this *Attribute) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Attribute) ColorName() string {
	return this.colorName
}

func (this *Attribute) SetColorName(val string) {
	this.colorName = val
}

func (this *Attribute) Transparency() int {
	return this.transparency
}

func (this *Attribute) SetTransparency(val int) {
	this.transparency = val
}

func (this *Attribute) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Attribute) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsInvisible status flag.
func (this *Attribute) IsInvisible() bool {
	return this.Flags & 1 != 0
}

// IsInvisible status flag.
func (this *Attribute) SetIsInvisible(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// IsConstant status flag.
func (this *Attribute) IsConstant() bool {
	return this.Flags & 2 != 0
}

// IsConstant status flag.
func (this *Attribute) SetIsConstant(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// IsInputVerificationRequired status flag.
func (this *Attribute) IsInputVerificationRequired() bool {
	return this.Flags & 4 != 0
}

// IsInputVerificationRequired status flag.
func (this *Attribute) SetIsInputVerificationRequired(val bool) {
	if val {
		this.Flags = this.Flags | 4
	} else {
		this.Flags = this.Flags & ^4
	}
}

// IsAttributePresent status flag.
func (this *Attribute) IsAttributePresent() bool {
	return this.Flags & 8 != 0
}

// IsAttributePresent status flag.
func (this *Attribute) SetIsAttributePresent(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

// IsTextBackward status flag.
func (this *Attribute) IsTextBackward() bool {
	return this.TextGenerationFlags & 2 != 0
}

// IsTextBackward status flag.
func (this *Attribute) SetIsTextBackward(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 2
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^2
	}
}

// IsTextUpsideDown status flag.
func (this *Attribute) IsTextUpsideDown() bool {
	return this.TextGenerationFlags & 4 != 0
}

// IsTextUpsideDown status flag.
func (this *Attribute) SetIsTextUpsideDown(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 4
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^4
	}
}

func (this *Attribute) AddsecondaryAttributeHandles(val string) {
	this.secondaryAttributeHandles = append(this.secondaryAttributeHandles, val)
}

func (this *Attribute) ClearsecondaryAttributeHandles() {
	this.secondaryAttributeHandles = []string{}
}

func (this *Attribute) typeString() string {
	return "ATTRIB"
}

func (this *Attribute) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Attribute) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Attribute) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ATTRIB"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbText"))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.TextHeight))
	pairs = append(pairs, NewStringCodePair(1, this.Value))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbAttribute"))
	}
	if version >= R2010 {
		pairs = append(pairs, NewShortCodePair(280, int16(this.Version)))
	}
	pairs = append(pairs, NewStringCodePair(2, this.AttributeTag))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	if this.FieldLength != 0 {
		pairs = append(pairs, NewShortCodePair(73, this.FieldLength))
	}
	if this.Rotation != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.Rotation))
	}
	if this.RelativeXScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.RelativeXScaleFactor))
	}
	if this.ObliqueAngle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(51, this.ObliqueAngle))
	}
	if this.TextStyleName != "STANDARD" {
		pairs = append(pairs, NewStringCodePair(7, this.TextStyleName))
	}
	if this.TextGenerationFlags != 0 {
		pairs = append(pairs, NewShortCodePair(71, int16(this.TextGenerationFlags)))
	}
	if this.HorizontalTextJustification != HorizontalTextJustificationLeft {
		pairs = append(pairs, NewShortCodePair(72, int16(this.HorizontalTextJustification)))
	}
	if version >= R12 && this.VerticalTextJustification != VerticalTextJustificationBaseline {
		pairs = append(pairs, NewShortCodePair(74, int16(this.VerticalTextJustification)))
	}
	pairs = append(pairs, NewDoubleCodePair(11, this.SecondAlignmentPoint.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.SecondAlignmentPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.SecondAlignmentPoint.Z))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.IsLockedInBlock)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbXrecord"))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.KeepDuplicateRecords)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, int16(this.MTextFlag)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(this.IsReallyLocked)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, int16(len(this.secondaryAttributeHandles))))
	}
	if version >= R2007 {
		for _, val := range this.secondaryAttributeHandles {
			pairs = append(pairs, NewStringCodePair(340, val))
		}
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(10, this.AlignmentPoint.X))
		pairs = append(pairs, NewDoubleCodePair(20, this.AlignmentPoint.Y))
		pairs = append(pairs, NewDoubleCodePair(30, this.AlignmentPoint.Z))
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(40, this.AnnotationScale))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(2, this.XRecordTag))
	}
	return
}

type Body struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FormatVersionNumber int16
	CustomData []string
	CustomData2 []string
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewBody() *Body {
	return &Body{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FormatVersionNumber: 1,
		CustomData: []string{},
		CustomData2: []string{},
	}
}

func (e *Body) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Body) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Body) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Body) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Body) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Body) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Body) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Body) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Body) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Body) Handle() Handle {
	return this.handle
}

func (this *Body) SetHandle(val Handle) {
	this.handle = val
}

func (this *Body) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Body) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Body) Layer() string {
	return this.layer
}

func (this *Body) SetLayer(val string) {
	this.layer = val
}

func (this *Body) LineTypeName() string {
	return this.lineTypeName
}

func (this *Body) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Body) Elevation() float64 {
	return this.elevation
}

func (this *Body) SetElevation(val float64) {
	this.elevation = val
}

func (this *Body) MaterialHandle() string {
	return this.materialHandle
}

func (this *Body) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Body) Color() Color {
	return this.color
}

func (this *Body) SetColor(val Color) {
	this.color = val
}

func (this *Body) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Body) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Body) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Body) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Body) IsVisible() bool {
	return this.isVisible
}

func (this *Body) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Body) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Body) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Body) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Body) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Body) Color24Bit() int {
	return this.color24Bit
}

func (this *Body) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Body) ColorName() string {
	return this.colorName
}

func (this *Body) SetColorName(val string) {
	this.colorName = val
}

func (this *Body) Transparency() int {
	return this.transparency
}

func (this *Body) SetTransparency(val int) {
	this.transparency = val
}

func (this *Body) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Body) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Body) AddCustomData(val string) {
	this.CustomData = append(this.CustomData, val)
}

func (this *Body) ClearCustomData() {
	this.CustomData = []string{}
}

func (this *Body) AddCustomData2(val string) {
	this.CustomData2 = append(this.CustomData2, val)
}

func (this *Body) ClearCustomData2() {
	this.CustomData2 = []string{}
}

func (this *Body) typeString() string {
	return "BODY"
}

func (this *Body) minVersion() (version AcadVersion) {
	return R13
}

func (this *Body) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Body) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 70:
		this.FormatVersionNumber = codePair.Value.(ShortCodePairValue).Value
	case 1:
		this.CustomData = append(this.CustomData, codePair.Value.(StringCodePairValue).Value)
	case 3:
		this.CustomData2 = append(this.CustomData2, codePair.Value.(StringCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Body) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "BODY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbModelerGeometry"))
	}
	pairs = append(pairs, NewShortCodePair(70, this.FormatVersionNumber))
	for _, val := range this.CustomData {
		pairs = append(pairs, NewStringCodePair(1, val))
	}
	for _, val := range this.CustomData2 {
		pairs = append(pairs, NewStringCodePair(3, val))
	}
	return
}

type Circle struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Center Point
	Radius float64
	Normal Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewCircle() *Circle {
	return &Circle{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Center: *NewOrigin(),
		Radius: 0.0,
		Normal: *NewZAxis(),
	}
}

func (e *Circle) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Circle) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Circle) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Circle) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Circle) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Circle) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Circle) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Circle) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Circle) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Circle) Handle() Handle {
	return this.handle
}

func (this *Circle) SetHandle(val Handle) {
	this.handle = val
}

func (this *Circle) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Circle) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Circle) Layer() string {
	return this.layer
}

func (this *Circle) SetLayer(val string) {
	this.layer = val
}

func (this *Circle) LineTypeName() string {
	return this.lineTypeName
}

func (this *Circle) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Circle) Elevation() float64 {
	return this.elevation
}

func (this *Circle) SetElevation(val float64) {
	this.elevation = val
}

func (this *Circle) MaterialHandle() string {
	return this.materialHandle
}

func (this *Circle) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Circle) Color() Color {
	return this.color
}

func (this *Circle) SetColor(val Color) {
	this.color = val
}

func (this *Circle) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Circle) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Circle) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Circle) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Circle) IsVisible() bool {
	return this.isVisible
}

func (this *Circle) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Circle) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Circle) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Circle) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Circle) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Circle) Color24Bit() int {
	return this.color24Bit
}

func (this *Circle) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Circle) ColorName() string {
	return this.colorName
}

func (this *Circle) SetColorName(val string) {
	this.colorName = val
}

func (this *Circle) Transparency() int {
	return this.transparency
}

func (this *Circle) SetTransparency(val int) {
	this.transparency = val
}

func (this *Circle) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Circle) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Circle) typeString() string {
	return "CIRCLE"
}

func (this *Circle) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Circle) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Circle) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.Center.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Center.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Center.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.Radius = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Circle) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "CIRCLE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbCircle"))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Center.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Center.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Center.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.Radius))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	return
}

type dimensionHelper struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	collectedPairs []CodePair
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewdimensionHelper() *dimensionHelper {
	return &dimensionHelper{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		collectedPairs: []CodePair{},
	}
}

func (e *dimensionHelper) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *dimensionHelper) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *dimensionHelper) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *dimensionHelper) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *dimensionHelper) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *dimensionHelper) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *dimensionHelper) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *dimensionHelper) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *dimensionHelper) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *dimensionHelper) Handle() Handle {
	return this.handle
}

func (this *dimensionHelper) SetHandle(val Handle) {
	this.handle = val
}

func (this *dimensionHelper) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *dimensionHelper) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *dimensionHelper) Layer() string {
	return this.layer
}

func (this *dimensionHelper) SetLayer(val string) {
	this.layer = val
}

func (this *dimensionHelper) LineTypeName() string {
	return this.lineTypeName
}

func (this *dimensionHelper) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *dimensionHelper) Elevation() float64 {
	return this.elevation
}

func (this *dimensionHelper) SetElevation(val float64) {
	this.elevation = val
}

func (this *dimensionHelper) MaterialHandle() string {
	return this.materialHandle
}

func (this *dimensionHelper) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *dimensionHelper) Color() Color {
	return this.color
}

func (this *dimensionHelper) SetColor(val Color) {
	this.color = val
}

func (this *dimensionHelper) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *dimensionHelper) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *dimensionHelper) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *dimensionHelper) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *dimensionHelper) IsVisible() bool {
	return this.isVisible
}

func (this *dimensionHelper) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *dimensionHelper) ImageByteCount() int {
	return this.imageByteCount
}

func (this *dimensionHelper) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *dimensionHelper) PreviewImageData() []string {
	return this.previewImageData
}

func (this *dimensionHelper) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *dimensionHelper) Color24Bit() int {
	return this.color24Bit
}

func (this *dimensionHelper) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *dimensionHelper) ColorName() string {
	return this.colorName
}

func (this *dimensionHelper) SetColorName(val string) {
	this.colorName = val
}

func (this *dimensionHelper) Transparency() int {
	return this.transparency
}

func (this *dimensionHelper) SetTransparency(val int) {
	this.transparency = val
}

func (this *dimensionHelper) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *dimensionHelper) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *dimensionHelper) Version() Version {
	return this.version
}

func (this *dimensionHelper) SetVersion(val Version) {
	this.version = val
}

func (this *dimensionHelper) BlockName() string {
	return this.blockName
}

func (this *dimensionHelper) SetBlockName(val string) {
	this.blockName = val
}

func (this *dimensionHelper) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *dimensionHelper) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *dimensionHelper) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *dimensionHelper) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *dimensionHelper) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *dimensionHelper) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *dimensionHelper) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *dimensionHelper) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *dimensionHelper) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *dimensionHelper) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *dimensionHelper) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *dimensionHelper) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *dimensionHelper) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *dimensionHelper) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *dimensionHelper) Text() string {
	return this.text
}

func (this *dimensionHelper) SetText(val string) {
	this.text = val
}

func (this *dimensionHelper) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *dimensionHelper) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *dimensionHelper) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *dimensionHelper) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *dimensionHelper) Normal() Vector {
	return this.normal
}

func (this *dimensionHelper) SetNormal(val Vector) {
	this.normal = val
}

func (this *dimensionHelper) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *dimensionHelper) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *dimensionHelper) AddcollectedPairs(val CodePair) {
	this.collectedPairs = append(this.collectedPairs, val)
}

func (this *dimensionHelper) ClearcollectedPairs() {
	this.collectedPairs = []CodePair{}
}

func (this *dimensionHelper) typeString() string {
	return "DIMENSION"
}

func (this *dimensionHelper) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *dimensionHelper) maxVersion() (version AcadVersion) {
	return R2018
}

type AlignedDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	DefinitionPoint2 Point
	DefinitionPoint3 Point
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewAlignedDimension() *AlignedDimension {
	return &AlignedDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		DefinitionPoint2: *NewOrigin(),
		DefinitionPoint3: *NewOrigin(),
	}
}

func (e *AlignedDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *AlignedDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *AlignedDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *AlignedDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *AlignedDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *AlignedDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *AlignedDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *AlignedDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *AlignedDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *AlignedDimension) Handle() Handle {
	return this.handle
}

func (this *AlignedDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *AlignedDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *AlignedDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *AlignedDimension) Layer() string {
	return this.layer
}

func (this *AlignedDimension) SetLayer(val string) {
	this.layer = val
}

func (this *AlignedDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *AlignedDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *AlignedDimension) Elevation() float64 {
	return this.elevation
}

func (this *AlignedDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *AlignedDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *AlignedDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *AlignedDimension) Color() Color {
	return this.color
}

func (this *AlignedDimension) SetColor(val Color) {
	this.color = val
}

func (this *AlignedDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *AlignedDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *AlignedDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *AlignedDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *AlignedDimension) IsVisible() bool {
	return this.isVisible
}

func (this *AlignedDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *AlignedDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *AlignedDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *AlignedDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *AlignedDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *AlignedDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *AlignedDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *AlignedDimension) ColorName() string {
	return this.colorName
}

func (this *AlignedDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *AlignedDimension) Transparency() int {
	return this.transparency
}

func (this *AlignedDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *AlignedDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *AlignedDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *AlignedDimension) Version() Version {
	return this.version
}

func (this *AlignedDimension) SetVersion(val Version) {
	this.version = val
}

func (this *AlignedDimension) BlockName() string {
	return this.blockName
}

func (this *AlignedDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *AlignedDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *AlignedDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *AlignedDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *AlignedDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *AlignedDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *AlignedDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *AlignedDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *AlignedDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *AlignedDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *AlignedDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *AlignedDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *AlignedDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *AlignedDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *AlignedDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *AlignedDimension) Text() string {
	return this.text
}

func (this *AlignedDimension) SetText(val string) {
	this.text = val
}

func (this *AlignedDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *AlignedDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *AlignedDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *AlignedDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *AlignedDimension) Normal() Vector {
	return this.normal
}

func (this *AlignedDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *AlignedDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *AlignedDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *AlignedDimension) typeString() string {
	return "DIMENSION"
}

func (this *AlignedDimension) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *AlignedDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *AlignedDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 13:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 14:
		this.DefinitionPoint3.X = codePair.Value.(DoubleCodePairValue).Value
	case 24:
		this.DefinitionPoint3.Y = codePair.Value.(DoubleCodePairValue).Value
	case 34:
		this.DefinitionPoint3.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *AlignedDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbAlignedDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(13, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(14, this.DefinitionPoint3.X))
	pairs = append(pairs, NewDoubleCodePair(24, this.DefinitionPoint3.Y))
	pairs = append(pairs, NewDoubleCodePair(34, this.DefinitionPoint3.Z))
	return
}

type RotatedDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	InsertionPoint Point
	DefinitionPoint2 Point
	DefinitionPoint3 Point
	RotationAngle float64
	ExtensionLineAngle float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewRotatedDimension() *RotatedDimension {
	return &RotatedDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		InsertionPoint: *NewOrigin(),
		DefinitionPoint2: *NewOrigin(),
		DefinitionPoint3: *NewOrigin(),
		RotationAngle: 0.0,
		ExtensionLineAngle: 0.0,
	}
}

func (e *RotatedDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *RotatedDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *RotatedDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *RotatedDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *RotatedDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *RotatedDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *RotatedDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *RotatedDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *RotatedDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *RotatedDimension) Handle() Handle {
	return this.handle
}

func (this *RotatedDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *RotatedDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *RotatedDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *RotatedDimension) Layer() string {
	return this.layer
}

func (this *RotatedDimension) SetLayer(val string) {
	this.layer = val
}

func (this *RotatedDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *RotatedDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *RotatedDimension) Elevation() float64 {
	return this.elevation
}

func (this *RotatedDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *RotatedDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *RotatedDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *RotatedDimension) Color() Color {
	return this.color
}

func (this *RotatedDimension) SetColor(val Color) {
	this.color = val
}

func (this *RotatedDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *RotatedDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *RotatedDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *RotatedDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *RotatedDimension) IsVisible() bool {
	return this.isVisible
}

func (this *RotatedDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *RotatedDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *RotatedDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *RotatedDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *RotatedDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *RotatedDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *RotatedDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *RotatedDimension) ColorName() string {
	return this.colorName
}

func (this *RotatedDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *RotatedDimension) Transparency() int {
	return this.transparency
}

func (this *RotatedDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *RotatedDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *RotatedDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *RotatedDimension) Version() Version {
	return this.version
}

func (this *RotatedDimension) SetVersion(val Version) {
	this.version = val
}

func (this *RotatedDimension) BlockName() string {
	return this.blockName
}

func (this *RotatedDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *RotatedDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *RotatedDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *RotatedDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *RotatedDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *RotatedDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *RotatedDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *RotatedDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *RotatedDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *RotatedDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *RotatedDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *RotatedDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *RotatedDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *RotatedDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *RotatedDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *RotatedDimension) Text() string {
	return this.text
}

func (this *RotatedDimension) SetText(val string) {
	this.text = val
}

func (this *RotatedDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *RotatedDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *RotatedDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *RotatedDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *RotatedDimension) Normal() Vector {
	return this.normal
}

func (this *RotatedDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *RotatedDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *RotatedDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *RotatedDimension) typeString() string {
	return "DIMENSION"
}

func (this *RotatedDimension) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *RotatedDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *RotatedDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 12:
		this.InsertionPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.InsertionPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.InsertionPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 13:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 14:
		this.DefinitionPoint3.X = codePair.Value.(DoubleCodePairValue).Value
	case 24:
		this.DefinitionPoint3.Y = codePair.Value.(DoubleCodePairValue).Value
	case 34:
		this.DefinitionPoint3.Z = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.RotationAngle = codePair.Value.(DoubleCodePairValue).Value
	case 52:
		this.ExtensionLineAngle = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *RotatedDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbAlignedDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(12, this.InsertionPoint.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.InsertionPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.InsertionPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(14, this.DefinitionPoint3.X))
	pairs = append(pairs, NewDoubleCodePair(24, this.DefinitionPoint3.Y))
	pairs = append(pairs, NewDoubleCodePair(34, this.DefinitionPoint3.Z))
	pairs = append(pairs, NewDoubleCodePair(50, this.RotationAngle))
	pairs = append(pairs, NewDoubleCodePair(52, this.ExtensionLineAngle))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbRotatedDimension"))
	}
	return
}

type RadialDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	DefinitionPoint2 Point
	LeaderLength float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewRadialDimension() *RadialDimension {
	return &RadialDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		DefinitionPoint2: *NewOrigin(),
		LeaderLength: 0.0,
	}
}

func (e *RadialDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *RadialDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *RadialDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *RadialDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *RadialDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *RadialDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *RadialDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *RadialDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *RadialDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *RadialDimension) Handle() Handle {
	return this.handle
}

func (this *RadialDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *RadialDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *RadialDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *RadialDimension) Layer() string {
	return this.layer
}

func (this *RadialDimension) SetLayer(val string) {
	this.layer = val
}

func (this *RadialDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *RadialDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *RadialDimension) Elevation() float64 {
	return this.elevation
}

func (this *RadialDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *RadialDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *RadialDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *RadialDimension) Color() Color {
	return this.color
}

func (this *RadialDimension) SetColor(val Color) {
	this.color = val
}

func (this *RadialDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *RadialDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *RadialDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *RadialDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *RadialDimension) IsVisible() bool {
	return this.isVisible
}

func (this *RadialDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *RadialDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *RadialDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *RadialDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *RadialDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *RadialDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *RadialDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *RadialDimension) ColorName() string {
	return this.colorName
}

func (this *RadialDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *RadialDimension) Transparency() int {
	return this.transparency
}

func (this *RadialDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *RadialDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *RadialDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *RadialDimension) Version() Version {
	return this.version
}

func (this *RadialDimension) SetVersion(val Version) {
	this.version = val
}

func (this *RadialDimension) BlockName() string {
	return this.blockName
}

func (this *RadialDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *RadialDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *RadialDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *RadialDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *RadialDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *RadialDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *RadialDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *RadialDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *RadialDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *RadialDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *RadialDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *RadialDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *RadialDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *RadialDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *RadialDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *RadialDimension) Text() string {
	return this.text
}

func (this *RadialDimension) SetText(val string) {
	this.text = val
}

func (this *RadialDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *RadialDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *RadialDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *RadialDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *RadialDimension) Normal() Vector {
	return this.normal
}

func (this *RadialDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *RadialDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *RadialDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *RadialDimension) typeString() string {
	return "DIMENSION"
}

func (this *RadialDimension) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *RadialDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *RadialDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 15:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 25:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 35:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.LeaderLength = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *RadialDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbRadialDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(15, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(25, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(35, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.LeaderLength))
	return
}

type DiameterDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	DefinitionPoint2 Point
	LeaderLength float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewDiameterDimension() *DiameterDimension {
	return &DiameterDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		DefinitionPoint2: *NewOrigin(),
		LeaderLength: 0.0,
	}
}

func (e *DiameterDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *DiameterDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *DiameterDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *DiameterDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *DiameterDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *DiameterDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *DiameterDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *DiameterDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *DiameterDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *DiameterDimension) Handle() Handle {
	return this.handle
}

func (this *DiameterDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *DiameterDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *DiameterDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *DiameterDimension) Layer() string {
	return this.layer
}

func (this *DiameterDimension) SetLayer(val string) {
	this.layer = val
}

func (this *DiameterDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *DiameterDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *DiameterDimension) Elevation() float64 {
	return this.elevation
}

func (this *DiameterDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *DiameterDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *DiameterDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *DiameterDimension) Color() Color {
	return this.color
}

func (this *DiameterDimension) SetColor(val Color) {
	this.color = val
}

func (this *DiameterDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *DiameterDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *DiameterDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *DiameterDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *DiameterDimension) IsVisible() bool {
	return this.isVisible
}

func (this *DiameterDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *DiameterDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *DiameterDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *DiameterDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *DiameterDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *DiameterDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *DiameterDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *DiameterDimension) ColorName() string {
	return this.colorName
}

func (this *DiameterDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *DiameterDimension) Transparency() int {
	return this.transparency
}

func (this *DiameterDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *DiameterDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *DiameterDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *DiameterDimension) Version() Version {
	return this.version
}

func (this *DiameterDimension) SetVersion(val Version) {
	this.version = val
}

func (this *DiameterDimension) BlockName() string {
	return this.blockName
}

func (this *DiameterDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *DiameterDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *DiameterDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *DiameterDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *DiameterDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *DiameterDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *DiameterDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *DiameterDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *DiameterDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *DiameterDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *DiameterDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *DiameterDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *DiameterDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *DiameterDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *DiameterDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *DiameterDimension) Text() string {
	return this.text
}

func (this *DiameterDimension) SetText(val string) {
	this.text = val
}

func (this *DiameterDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *DiameterDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *DiameterDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *DiameterDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *DiameterDimension) Normal() Vector {
	return this.normal
}

func (this *DiameterDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *DiameterDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *DiameterDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *DiameterDimension) typeString() string {
	return "DIMENSION"
}

func (this *DiameterDimension) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *DiameterDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *DiameterDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 15:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 25:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 35:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.LeaderLength = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *DiameterDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbDiametricDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(15, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(25, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(35, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.LeaderLength))
	return
}

type AngularThreePointDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	DefinitionPoint2 Point
	DefinitionPoint3 Point
	DefinitionPoint4 Point
	DefinitionPoint5 Point
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewAngularThreePointDimension() *AngularThreePointDimension {
	return &AngularThreePointDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		DefinitionPoint2: *NewOrigin(),
		DefinitionPoint3: *NewOrigin(),
		DefinitionPoint4: *NewOrigin(),
		DefinitionPoint5: *NewOrigin(),
	}
}

func (e *AngularThreePointDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *AngularThreePointDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *AngularThreePointDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *AngularThreePointDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *AngularThreePointDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *AngularThreePointDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *AngularThreePointDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *AngularThreePointDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *AngularThreePointDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *AngularThreePointDimension) Handle() Handle {
	return this.handle
}

func (this *AngularThreePointDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *AngularThreePointDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *AngularThreePointDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *AngularThreePointDimension) Layer() string {
	return this.layer
}

func (this *AngularThreePointDimension) SetLayer(val string) {
	this.layer = val
}

func (this *AngularThreePointDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *AngularThreePointDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *AngularThreePointDimension) Elevation() float64 {
	return this.elevation
}

func (this *AngularThreePointDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *AngularThreePointDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *AngularThreePointDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *AngularThreePointDimension) Color() Color {
	return this.color
}

func (this *AngularThreePointDimension) SetColor(val Color) {
	this.color = val
}

func (this *AngularThreePointDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *AngularThreePointDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *AngularThreePointDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *AngularThreePointDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *AngularThreePointDimension) IsVisible() bool {
	return this.isVisible
}

func (this *AngularThreePointDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *AngularThreePointDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *AngularThreePointDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *AngularThreePointDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *AngularThreePointDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *AngularThreePointDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *AngularThreePointDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *AngularThreePointDimension) ColorName() string {
	return this.colorName
}

func (this *AngularThreePointDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *AngularThreePointDimension) Transparency() int {
	return this.transparency
}

func (this *AngularThreePointDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *AngularThreePointDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *AngularThreePointDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *AngularThreePointDimension) Version() Version {
	return this.version
}

func (this *AngularThreePointDimension) SetVersion(val Version) {
	this.version = val
}

func (this *AngularThreePointDimension) BlockName() string {
	return this.blockName
}

func (this *AngularThreePointDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *AngularThreePointDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *AngularThreePointDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *AngularThreePointDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *AngularThreePointDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *AngularThreePointDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *AngularThreePointDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *AngularThreePointDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *AngularThreePointDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *AngularThreePointDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *AngularThreePointDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *AngularThreePointDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *AngularThreePointDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *AngularThreePointDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *AngularThreePointDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *AngularThreePointDimension) Text() string {
	return this.text
}

func (this *AngularThreePointDimension) SetText(val string) {
	this.text = val
}

func (this *AngularThreePointDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *AngularThreePointDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *AngularThreePointDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *AngularThreePointDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *AngularThreePointDimension) Normal() Vector {
	return this.normal
}

func (this *AngularThreePointDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *AngularThreePointDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *AngularThreePointDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *AngularThreePointDimension) typeString() string {
	return "DIMENSION"
}

func (this *AngularThreePointDimension) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *AngularThreePointDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *AngularThreePointDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 13:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 14:
		this.DefinitionPoint3.X = codePair.Value.(DoubleCodePairValue).Value
	case 24:
		this.DefinitionPoint3.Y = codePair.Value.(DoubleCodePairValue).Value
	case 34:
		this.DefinitionPoint3.Z = codePair.Value.(DoubleCodePairValue).Value
	case 15:
		this.DefinitionPoint4.X = codePair.Value.(DoubleCodePairValue).Value
	case 25:
		this.DefinitionPoint4.Y = codePair.Value.(DoubleCodePairValue).Value
	case 35:
		this.DefinitionPoint4.Z = codePair.Value.(DoubleCodePairValue).Value
	case 16:
		this.DefinitionPoint5.X = codePair.Value.(DoubleCodePairValue).Value
	case 26:
		this.DefinitionPoint5.Y = codePair.Value.(DoubleCodePairValue).Value
	case 36:
		this.DefinitionPoint5.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *AngularThreePointDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDb3PointAngularDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(13, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(14, this.DefinitionPoint3.X))
	pairs = append(pairs, NewDoubleCodePair(24, this.DefinitionPoint3.Y))
	pairs = append(pairs, NewDoubleCodePair(34, this.DefinitionPoint3.Z))
	pairs = append(pairs, NewDoubleCodePair(15, this.DefinitionPoint4.X))
	pairs = append(pairs, NewDoubleCodePair(25, this.DefinitionPoint4.Y))
	pairs = append(pairs, NewDoubleCodePair(35, this.DefinitionPoint4.Z))
	pairs = append(pairs, NewDoubleCodePair(16, this.DefinitionPoint5.X))
	pairs = append(pairs, NewDoubleCodePair(26, this.DefinitionPoint5.Y))
	pairs = append(pairs, NewDoubleCodePair(36, this.DefinitionPoint5.Z))
	return
}

type OrdinateDimension struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Dimension interface
	version Version
	blockName string
	definitionPoint1 Point
	textMidPoint Point
	dimensionType DimensionType
	attachmentPoint AttachmentPoint
	textLineSpacingStyle TextLineSpacingStyle
	textLineSpacingFactor float64
	actualMeasurement float64
	text string
	textRotationAngle float64
	horizontalDirectionAngle float64
	normal Vector
	dimensionStyleName string
	DefinitionPoint2 Point
	DefinitionPoint3 Point
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewOrdinateDimension() *OrdinateDimension {
	return &OrdinateDimension{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		version: VersionR2010,
		blockName: "*MODEL_SPACE",
		definitionPoint1: *NewOrigin(),
		textMidPoint: *NewOrigin(),
		dimensionType: DimensionTypeAligned,
		attachmentPoint: AttachmentPointTopLeft,
		textLineSpacingStyle: TextLineSpacingStyleAtLeast,
		textLineSpacingFactor: 1.0,
		actualMeasurement: 0.0,
		text: "<>",
		textRotationAngle: 0.0,
		horizontalDirectionAngle: 0.0,
		normal: *NewZAxis(),
		dimensionStyleName: "STANDARD",
		DefinitionPoint2: *NewOrigin(),
		DefinitionPoint3: *NewOrigin(),
	}
}

func (e *OrdinateDimension) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *OrdinateDimension) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *OrdinateDimension) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *OrdinateDimension) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *OrdinateDimension) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *OrdinateDimension) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *OrdinateDimension) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *OrdinateDimension) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *OrdinateDimension) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *OrdinateDimension) Handle() Handle {
	return this.handle
}

func (this *OrdinateDimension) SetHandle(val Handle) {
	this.handle = val
}

func (this *OrdinateDimension) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *OrdinateDimension) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *OrdinateDimension) Layer() string {
	return this.layer
}

func (this *OrdinateDimension) SetLayer(val string) {
	this.layer = val
}

func (this *OrdinateDimension) LineTypeName() string {
	return this.lineTypeName
}

func (this *OrdinateDimension) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *OrdinateDimension) Elevation() float64 {
	return this.elevation
}

func (this *OrdinateDimension) SetElevation(val float64) {
	this.elevation = val
}

func (this *OrdinateDimension) MaterialHandle() string {
	return this.materialHandle
}

func (this *OrdinateDimension) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *OrdinateDimension) Color() Color {
	return this.color
}

func (this *OrdinateDimension) SetColor(val Color) {
	this.color = val
}

func (this *OrdinateDimension) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *OrdinateDimension) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *OrdinateDimension) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *OrdinateDimension) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *OrdinateDimension) IsVisible() bool {
	return this.isVisible
}

func (this *OrdinateDimension) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *OrdinateDimension) ImageByteCount() int {
	return this.imageByteCount
}

func (this *OrdinateDimension) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *OrdinateDimension) PreviewImageData() []string {
	return this.previewImageData
}

func (this *OrdinateDimension) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *OrdinateDimension) Color24Bit() int {
	return this.color24Bit
}

func (this *OrdinateDimension) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *OrdinateDimension) ColorName() string {
	return this.colorName
}

func (this *OrdinateDimension) SetColorName(val string) {
	this.colorName = val
}

func (this *OrdinateDimension) Transparency() int {
	return this.transparency
}

func (this *OrdinateDimension) SetTransparency(val int) {
	this.transparency = val
}

func (this *OrdinateDimension) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *OrdinateDimension) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *OrdinateDimension) Version() Version {
	return this.version
}

func (this *OrdinateDimension) SetVersion(val Version) {
	this.version = val
}

func (this *OrdinateDimension) BlockName() string {
	return this.blockName
}

func (this *OrdinateDimension) SetBlockName(val string) {
	this.blockName = val
}

func (this *OrdinateDimension) DefinitionPoint1() Point {
	return this.definitionPoint1
}

func (this *OrdinateDimension) SetDefinitionPoint1(val Point) {
	this.definitionPoint1 = val
}

func (this *OrdinateDimension) TextMidPoint() Point {
	return this.textMidPoint
}

func (this *OrdinateDimension) SetTextMidPoint(val Point) {
	this.textMidPoint = val
}

func (this *OrdinateDimension) DimensionType() DimensionType {
	return this.dimensionType
}

func (this *OrdinateDimension) SetDimensionType(val DimensionType) {
	this.dimensionType = val
}

func (this *OrdinateDimension) AttachmentPoint() AttachmentPoint {
	return this.attachmentPoint
}

func (this *OrdinateDimension) SetAttachmentPoint(val AttachmentPoint) {
	this.attachmentPoint = val
}

func (this *OrdinateDimension) TextLineSpacingStyle() TextLineSpacingStyle {
	return this.textLineSpacingStyle
}

func (this *OrdinateDimension) SetTextLineSpacingStyle(val TextLineSpacingStyle) {
	this.textLineSpacingStyle = val
}

func (this *OrdinateDimension) TextLineSpacingFactor() float64 {
	return this.textLineSpacingFactor
}

func (this *OrdinateDimension) SetTextLineSpacingFactor(val float64) {
	this.textLineSpacingFactor = val
}

func (this *OrdinateDimension) ActualMeasurement() float64 {
	return this.actualMeasurement
}

func (this *OrdinateDimension) SetActualMeasurement(val float64) {
	this.actualMeasurement = val
}

func (this *OrdinateDimension) Text() string {
	return this.text
}

func (this *OrdinateDimension) SetText(val string) {
	this.text = val
}

func (this *OrdinateDimension) TextRotationAngle() float64 {
	return this.textRotationAngle
}

func (this *OrdinateDimension) SetTextRotationAngle(val float64) {
	this.textRotationAngle = val
}

func (this *OrdinateDimension) HorizontalDirectionAngle() float64 {
	return this.horizontalDirectionAngle
}

func (this *OrdinateDimension) SetHorizontalDirectionAngle(val float64) {
	this.horizontalDirectionAngle = val
}

func (this *OrdinateDimension) Normal() Vector {
	return this.normal
}

func (this *OrdinateDimension) SetNormal(val Vector) {
	this.normal = val
}

func (this *OrdinateDimension) DimensionStyleName() string {
	return this.dimensionStyleName
}

func (this *OrdinateDimension) SetDimensionStyleName(val string) {
	this.dimensionStyleName = val
}

func (this *OrdinateDimension) typeString() string {
	return "DIMENSION"
}

func (this *OrdinateDimension) minVersion() (version AcadVersion) {
	return R11
}

func (this *OrdinateDimension) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *OrdinateDimension) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 13:
		this.DefinitionPoint2.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.DefinitionPoint2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.DefinitionPoint2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 14:
		this.DefinitionPoint3.X = codePair.Value.(DoubleCodePairValue).Value
	case 24:
		this.DefinitionPoint3.Y = codePair.Value.(DoubleCodePairValue).Value
	case 34:
		this.DefinitionPoint3.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForDimension(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *OrdinateDimension) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DIMENSION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForDimension(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbOrdinateDimension"))
	}
	pairs = append(pairs, NewDoubleCodePair(13, this.DefinitionPoint2.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.DefinitionPoint2.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.DefinitionPoint2.Z))
	pairs = append(pairs, NewDoubleCodePair(14, this.DefinitionPoint3.X))
	pairs = append(pairs, NewDoubleCodePair(24, this.DefinitionPoint3.Y))
	pairs = append(pairs, NewDoubleCodePair(34, this.DefinitionPoint3.Z))
	return
}

type Ellipse struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Center Point
	MajorAxis Vector
	Normal Vector
	MinorAxisRatio float64
	StartAngle float64 // Ellipse start angle in radians.
	EndAngle float64 // Ellipse end angle in radians.
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewEllipse() *Ellipse {
	return &Ellipse{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Center: *NewOrigin(),
		MajorAxis: *NewXAxis(),
		Normal: *NewZAxis(),
		MinorAxisRatio: 1.0,
		StartAngle: 0.0,
		EndAngle: math.Pi * 2.0,
	}
}

func (e *Ellipse) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Ellipse) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Ellipse) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Ellipse) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Ellipse) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Ellipse) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Ellipse) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Ellipse) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Ellipse) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Ellipse) Handle() Handle {
	return this.handle
}

func (this *Ellipse) SetHandle(val Handle) {
	this.handle = val
}

func (this *Ellipse) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Ellipse) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Ellipse) Layer() string {
	return this.layer
}

func (this *Ellipse) SetLayer(val string) {
	this.layer = val
}

func (this *Ellipse) LineTypeName() string {
	return this.lineTypeName
}

func (this *Ellipse) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Ellipse) Elevation() float64 {
	return this.elevation
}

func (this *Ellipse) SetElevation(val float64) {
	this.elevation = val
}

func (this *Ellipse) MaterialHandle() string {
	return this.materialHandle
}

func (this *Ellipse) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Ellipse) Color() Color {
	return this.color
}

func (this *Ellipse) SetColor(val Color) {
	this.color = val
}

func (this *Ellipse) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Ellipse) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Ellipse) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Ellipse) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Ellipse) IsVisible() bool {
	return this.isVisible
}

func (this *Ellipse) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Ellipse) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Ellipse) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Ellipse) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Ellipse) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Ellipse) Color24Bit() int {
	return this.color24Bit
}

func (this *Ellipse) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Ellipse) ColorName() string {
	return this.colorName
}

func (this *Ellipse) SetColorName(val string) {
	this.colorName = val
}

func (this *Ellipse) Transparency() int {
	return this.transparency
}

func (this *Ellipse) SetTransparency(val int) {
	this.transparency = val
}

func (this *Ellipse) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Ellipse) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Ellipse) typeString() string {
	return "ELLIPSE"
}

func (this *Ellipse) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Ellipse) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Ellipse) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.Center.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Center.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Center.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.MajorAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.MajorAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.MajorAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.MinorAxisRatio = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.StartAngle = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.EndAngle = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Ellipse) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "ELLIPSE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbEllipse"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Center.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Center.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Center.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.MajorAxis.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.MajorAxis.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.MajorAxis.Z))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	pairs = append(pairs, NewDoubleCodePair(40, this.MinorAxisRatio))
	pairs = append(pairs, NewDoubleCodePair(41, this.StartAngle))
	pairs = append(pairs, NewDoubleCodePair(42, this.EndAngle))
	return
}

type Helix struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	MajorReleaseNumber int
	MaintainenceReleaseNumber int
	AxisBasePoint Point
	StartPoint Point
	AxisVector Vector
	Radius float64
	NumberOfTurns float64
	TurnHeight float64
	IsRightHanded bool
	Constraint HelixConstraint
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewHelix() *Helix {
	return &Helix{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		MajorReleaseNumber: 0,
		MaintainenceReleaseNumber: 0,
		AxisBasePoint: *NewOrigin(),
		StartPoint: *NewOrigin(),
		AxisVector: *NewZAxis(),
		Radius: 0.0,
		NumberOfTurns: 0.0,
		TurnHeight: 0.0,
		IsRightHanded: false,
		Constraint: HelixConstraintConstrainTurnHeight,
	}
}

func (e *Helix) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Helix) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Helix) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Helix) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Helix) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Helix) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Helix) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Helix) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Helix) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Helix) Handle() Handle {
	return this.handle
}

func (this *Helix) SetHandle(val Handle) {
	this.handle = val
}

func (this *Helix) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Helix) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Helix) Layer() string {
	return this.layer
}

func (this *Helix) SetLayer(val string) {
	this.layer = val
}

func (this *Helix) LineTypeName() string {
	return this.lineTypeName
}

func (this *Helix) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Helix) Elevation() float64 {
	return this.elevation
}

func (this *Helix) SetElevation(val float64) {
	this.elevation = val
}

func (this *Helix) MaterialHandle() string {
	return this.materialHandle
}

func (this *Helix) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Helix) Color() Color {
	return this.color
}

func (this *Helix) SetColor(val Color) {
	this.color = val
}

func (this *Helix) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Helix) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Helix) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Helix) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Helix) IsVisible() bool {
	return this.isVisible
}

func (this *Helix) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Helix) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Helix) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Helix) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Helix) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Helix) Color24Bit() int {
	return this.color24Bit
}

func (this *Helix) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Helix) ColorName() string {
	return this.colorName
}

func (this *Helix) SetColorName(val string) {
	this.colorName = val
}

func (this *Helix) Transparency() int {
	return this.transparency
}

func (this *Helix) SetTransparency(val int) {
	this.transparency = val
}

func (this *Helix) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Helix) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Helix) typeString() string {
	return "HELIX"
}

func (this *Helix) minVersion() (version AcadVersion) {
	return R2007
}

func (this *Helix) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Helix) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 90:
		this.MajorReleaseNumber = codePair.Value.(IntCodePairValue).Value
	case 91:
		this.MaintainenceReleaseNumber = codePair.Value.(IntCodePairValue).Value
	case 10:
		this.AxisBasePoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.AxisBasePoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.AxisBasePoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.StartPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.StartPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.StartPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.AxisVector.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.AxisVector.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.AxisVector.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.Radius = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.NumberOfTurns = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.TurnHeight = codePair.Value.(DoubleCodePairValue).Value
	case 290:
		this.IsRightHanded = codePair.Value.(BoolCodePairValue).Value
	case 280:
		this.Constraint = HelixConstraint(codePair.Value.(ShortCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Helix) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "HELIX"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbHelix"))
	}
	pairs = append(pairs, NewIntCodePair(90, this.MajorReleaseNumber))
	pairs = append(pairs, NewIntCodePair(91, this.MaintainenceReleaseNumber))
	pairs = append(pairs, NewDoubleCodePair(10, this.AxisBasePoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.AxisBasePoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.AxisBasePoint.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.StartPoint.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.StartPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.StartPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.AxisVector.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.AxisVector.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.AxisVector.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.Radius))
	pairs = append(pairs, NewDoubleCodePair(41, this.NumberOfTurns))
	pairs = append(pairs, NewDoubleCodePair(42, this.TurnHeight))
	pairs = append(pairs, NewBoolCodePair(290, this.IsRightHanded))
	pairs = append(pairs, NewShortCodePair(280, int16(this.Constraint)))
	return
}

type Image struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for RasterImage interface
	_subclassMarker string
	classVersion int
	location Point // The location of the bottom-left corner of the image
	uVector Vector
	vVector Vector
	imageSize Vector // Image size in pixels
	_imageDefinitionHandle string
	displayOptionsFlags int
	useClipping bool
	brightness int16
	contrast int16
	fade int16
	_imageDefinitionReactorHandle string
	clippingType ImageClippingBoundaryType
	clippingVertices []Point
	_clippingVertexCount int
	_clippingVerticesX []float64
	_clippingVerticesY []float64
	isInsideClipping bool
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewImage() *Image {
	return &Image{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		_subclassMarker: "",
		classVersion: 0,
		location: *NewOrigin(),
		uVector: *NewXAxis(),
		vVector: *NewYAxis(),
		imageSize: *NewZeroVector(),
		_imageDefinitionHandle: "",
		displayOptionsFlags: 1,
		useClipping: true,
		brightness: 50,
		contrast: 50,
		fade: 0,
		_imageDefinitionReactorHandle: "",
		clippingType: ImageClippingBoundaryTypeRectangular,
		clippingVertices: []Point{},
		_clippingVertexCount: 0,
		_clippingVerticesX: []float64{},
		_clippingVerticesY: []float64{},
		isInsideClipping: false,
	}
}

func (e *Image) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Image) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Image) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Image) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Image) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Image) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Image) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Image) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Image) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Image) Handle() Handle {
	return this.handle
}

func (this *Image) SetHandle(val Handle) {
	this.handle = val
}

func (this *Image) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Image) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Image) Layer() string {
	return this.layer
}

func (this *Image) SetLayer(val string) {
	this.layer = val
}

func (this *Image) LineTypeName() string {
	return this.lineTypeName
}

func (this *Image) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Image) Elevation() float64 {
	return this.elevation
}

func (this *Image) SetElevation(val float64) {
	this.elevation = val
}

func (this *Image) MaterialHandle() string {
	return this.materialHandle
}

func (this *Image) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Image) Color() Color {
	return this.color
}

func (this *Image) SetColor(val Color) {
	this.color = val
}

func (this *Image) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Image) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Image) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Image) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Image) IsVisible() bool {
	return this.isVisible
}

func (this *Image) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Image) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Image) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Image) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Image) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Image) Color24Bit() int {
	return this.color24Bit
}

func (this *Image) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Image) ColorName() string {
	return this.colorName
}

func (this *Image) SetColorName(val string) {
	this.colorName = val
}

func (this *Image) Transparency() int {
	return this.transparency
}

func (this *Image) SetTransparency(val int) {
	this.transparency = val
}

func (this *Image) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Image) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Image) subclassMarker() string {
	return this._subclassMarker
}

func (this *Image) SetsubclassMarker(val string) {
	this._subclassMarker = val
}

func (this *Image) ClassVersion() int {
	return this.classVersion
}

func (this *Image) SetClassVersion(val int) {
	this.classVersion = val
}

func (this *Image) Location() Point {
	return this.location
}

func (this *Image) SetLocation(val Point) {
	this.location = val
}

func (this *Image) UVector() Vector {
	return this.uVector
}

func (this *Image) SetUVector(val Vector) {
	this.uVector = val
}

func (this *Image) VVector() Vector {
	return this.vVector
}

func (this *Image) SetVVector(val Vector) {
	this.vVector = val
}

func (this *Image) ImageSize() Vector {
	return this.imageSize
}

func (this *Image) SetImageSize(val Vector) {
	this.imageSize = val
}

func (this *Image) imageDefinitionHandle() string {
	return this._imageDefinitionHandle
}

func (this *Image) SetimageDefinitionHandle(val string) {
	this._imageDefinitionHandle = val
}

func (this *Image) DisplayOptionsFlags() int {
	return this.displayOptionsFlags
}

func (this *Image) SetDisplayOptionsFlags(val int) {
	this.displayOptionsFlags = val
}

func (this *Image) UseClipping() bool {
	return this.useClipping
}

func (this *Image) SetUseClipping(val bool) {
	this.useClipping = val
}

func (this *Image) Brightness() int16 {
	return this.brightness
}

func (this *Image) SetBrightness(val int16) {
	this.brightness = val
}

func (this *Image) Contrast() int16 {
	return this.contrast
}

func (this *Image) SetContrast(val int16) {
	this.contrast = val
}

func (this *Image) Fade() int16 {
	return this.fade
}

func (this *Image) SetFade(val int16) {
	this.fade = val
}

func (this *Image) imageDefinitionReactorHandle() string {
	return this._imageDefinitionReactorHandle
}

func (this *Image) SetimageDefinitionReactorHandle(val string) {
	this._imageDefinitionReactorHandle = val
}

func (this *Image) ClippingType() ImageClippingBoundaryType {
	return this.clippingType
}

func (this *Image) SetClippingType(val ImageClippingBoundaryType) {
	this.clippingType = val
}

func (this *Image) ClippingVertices() []Point {
	return this.clippingVertices
}

func (this *Image) SetClippingVertices(val []Point) {
	this.clippingVertices = val
}

func (this *Image) clippingVertexCount() int {
	return this._clippingVertexCount
}

func (this *Image) SetclippingVertexCount(val int) {
	this._clippingVertexCount = val
}

func (this *Image) clippingVerticesX() []float64 {
	return this._clippingVerticesX
}

func (this *Image) SetclippingVerticesX(val []float64) {
	this._clippingVerticesX = val
}

func (this *Image) clippingVerticesY() []float64 {
	return this._clippingVerticesY
}

func (this *Image) SetclippingVerticesY(val []float64) {
	this._clippingVerticesY = val
}

func (this *Image) IsInsideClipping() bool {
	return this.isInsideClipping
}

func (this *Image) SetIsInsideClipping(val bool) {
	this.isInsideClipping = val
}

func (this *Image) typeString() string {
	return "IMAGE"
}

func (this *Image) minVersion() (version AcadVersion) {
	return R14
}

func (this *Image) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Image) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForRasterImage(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Image) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "IMAGE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForRasterImage(this, version)...)
	return
}

type Insert struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	HasAttributes bool
	Name string
	Location Point
	XScaleFactor float64
	YScaleFactor float64
	ZScaleFactor float64
	Rotation float64
	ColumnCount int16
	RowCount int16
	ColumnSpacing float64
	RowSpacing float64
	ExtrusionDirection Vector
	Attributes []Attribute
	seqend Seqend
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewInsert() *Insert {
	return &Insert{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		HasAttributes: false,
		Name: "",
		Location: *NewOrigin(),
		XScaleFactor: 1.0,
		YScaleFactor: 1.0,
		ZScaleFactor: 1.0,
		Rotation: 0.0,
		ColumnCount: 1,
		RowCount: 1,
		ColumnSpacing: 0.0,
		RowSpacing: 0.0,
		ExtrusionDirection: *NewZAxis(),
		Attributes: []Attribute{},
		seqend: *NewSeqend(),
	}
}

func (e *Insert) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Insert) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Insert) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Insert) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Insert) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Insert) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Insert) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Insert) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Insert) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Insert) Handle() Handle {
	return this.handle
}

func (this *Insert) SetHandle(val Handle) {
	this.handle = val
}

func (this *Insert) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Insert) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Insert) Layer() string {
	return this.layer
}

func (this *Insert) SetLayer(val string) {
	this.layer = val
}

func (this *Insert) LineTypeName() string {
	return this.lineTypeName
}

func (this *Insert) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Insert) Elevation() float64 {
	return this.elevation
}

func (this *Insert) SetElevation(val float64) {
	this.elevation = val
}

func (this *Insert) MaterialHandle() string {
	return this.materialHandle
}

func (this *Insert) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Insert) Color() Color {
	return this.color
}

func (this *Insert) SetColor(val Color) {
	this.color = val
}

func (this *Insert) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Insert) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Insert) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Insert) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Insert) IsVisible() bool {
	return this.isVisible
}

func (this *Insert) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Insert) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Insert) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Insert) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Insert) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Insert) Color24Bit() int {
	return this.color24Bit
}

func (this *Insert) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Insert) ColorName() string {
	return this.colorName
}

func (this *Insert) SetColorName(val string) {
	this.colorName = val
}

func (this *Insert) Transparency() int {
	return this.transparency
}

func (this *Insert) SetTransparency(val int) {
	this.transparency = val
}

func (this *Insert) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Insert) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Insert) AddAttributes(val Attribute) {
	this.Attributes = append(this.Attributes, val)
}

func (this *Insert) ClearAttributes() {
	this.Attributes = []Attribute{}
}

func (this *Insert) typeString() string {
	return "INSERT"
}

func (this *Insert) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Insert) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Insert) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 66:
		this.HasAttributes = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 10:
		this.Location.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Location.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Location.Z = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.XScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.YScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 43:
		this.ZScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.Rotation = codePair.Value.(DoubleCodePairValue).Value
	case 70:
		this.ColumnCount = codePair.Value.(ShortCodePairValue).Value
	case 71:
		this.RowCount = codePair.Value.(ShortCodePairValue).Value
	case 44:
		this.ColumnSpacing = codePair.Value.(DoubleCodePairValue).Value
	case 45:
		this.RowSpacing = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Insert) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "INSERT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbBlockReference"))
	}
	if this.HasAttributes != false {
		pairs = append(pairs, NewShortCodePair(66, shortFromBool(this.HasAttributes)))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	if this.XScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.XScaleFactor))
	}
	if this.YScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(42, this.YScaleFactor))
	}
	if this.ZScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(43, this.ZScaleFactor))
	}
	if this.Rotation != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.Rotation))
	}
	if this.ColumnCount != 1 {
		pairs = append(pairs, NewShortCodePair(70, this.ColumnCount))
	}
	if this.RowCount != 1 {
		pairs = append(pairs, NewShortCodePair(71, this.RowCount))
	}
	if this.ColumnSpacing != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(44, this.ColumnSpacing))
	}
	if this.RowSpacing != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(45, this.RowSpacing))
	}
	pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
	pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
	pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	return
}

type Leader struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	DimensionStyleName string
	UseArrowheads bool
	PathType LeaderPathType
	AnnotationType LeaderCreationAnnotationType
	HooklineDirection LeaderHooklineDirection
	UseHookline bool
	TextAnnotationHeight float64
	TextAnnotationWidth float64
	Vertices []Point
	vertexCount int
	verticesX []float64
	verticesY []float64
	verticesZ []float64
	OverrideColor Color
	AssociatedAnnotationReference string
	Normal Vector
	Right Vector
	BlockOffset Vector
	AnnotationOffset Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewLeader() *Leader {
	return &Leader{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		DimensionStyleName: "",
		UseArrowheads: true,
		PathType: LeaderPathTypeStraightLineSegments,
		AnnotationType: LeaderCreationAnnotationTypeNoAnnotation,
		HooklineDirection: LeaderHooklineDirectionOppositeFromHorizontalVector,
		UseHookline: true,
		TextAnnotationHeight: 1.0,
		TextAnnotationWidth: 1.0,
		Vertices: []Point{},
		vertexCount: 0,
		verticesX: []float64{},
		verticesY: []float64{},
		verticesZ: []float64{},
		OverrideColor: ByBlock(),
		AssociatedAnnotationReference: "",
		Normal: *NewZAxis(),
		Right: *NewXAxis(),
		BlockOffset: *NewZeroVector(),
		AnnotationOffset: *NewXAxis(),
	}
}

func (e *Leader) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Leader) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Leader) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Leader) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Leader) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Leader) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Leader) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Leader) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Leader) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Leader) Handle() Handle {
	return this.handle
}

func (this *Leader) SetHandle(val Handle) {
	this.handle = val
}

func (this *Leader) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Leader) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Leader) Layer() string {
	return this.layer
}

func (this *Leader) SetLayer(val string) {
	this.layer = val
}

func (this *Leader) LineTypeName() string {
	return this.lineTypeName
}

func (this *Leader) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Leader) Elevation() float64 {
	return this.elevation
}

func (this *Leader) SetElevation(val float64) {
	this.elevation = val
}

func (this *Leader) MaterialHandle() string {
	return this.materialHandle
}

func (this *Leader) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Leader) Color() Color {
	return this.color
}

func (this *Leader) SetColor(val Color) {
	this.color = val
}

func (this *Leader) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Leader) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Leader) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Leader) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Leader) IsVisible() bool {
	return this.isVisible
}

func (this *Leader) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Leader) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Leader) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Leader) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Leader) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Leader) Color24Bit() int {
	return this.color24Bit
}

func (this *Leader) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Leader) ColorName() string {
	return this.colorName
}

func (this *Leader) SetColorName(val string) {
	this.colorName = val
}

func (this *Leader) Transparency() int {
	return this.transparency
}

func (this *Leader) SetTransparency(val int) {
	this.transparency = val
}

func (this *Leader) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Leader) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Leader) AddVertices(val Point) {
	this.Vertices = append(this.Vertices, val)
}

func (this *Leader) ClearVertices() {
	this.Vertices = []Point{}
}

func (this *Leader) AddverticesX(val float64) {
	this.verticesX = append(this.verticesX, val)
}

func (this *Leader) ClearverticesX() {
	this.verticesX = []float64{}
}

func (this *Leader) AddverticesY(val float64) {
	this.verticesY = append(this.verticesY, val)
}

func (this *Leader) ClearverticesY() {
	this.verticesY = []float64{}
}

func (this *Leader) AddverticesZ(val float64) {
	this.verticesZ = append(this.verticesZ, val)
}

func (this *Leader) ClearverticesZ() {
	this.verticesZ = []float64{}
}

func (this *Leader) typeString() string {
	return "LEADER"
}

func (this *Leader) minVersion() (version AcadVersion) {
	return R13
}

func (this *Leader) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Leader) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 3:
		this.DimensionStyleName = codePair.Value.(StringCodePairValue).Value
	case 71:
		this.UseArrowheads = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.PathType = LeaderPathType(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.AnnotationType = LeaderCreationAnnotationType(codePair.Value.(ShortCodePairValue).Value)
	case 74:
		this.HooklineDirection = LeaderHooklineDirection(codePair.Value.(ShortCodePairValue).Value)
	case 75:
		this.UseHookline = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 40:
		this.TextAnnotationHeight = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.TextAnnotationWidth = codePair.Value.(DoubleCodePairValue).Value
	case 76:
		this.vertexCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 10:
		this.verticesX = append(this.verticesX, codePair.Value.(DoubleCodePairValue).Value)
	case 20:
		this.verticesY = append(this.verticesY, codePair.Value.(DoubleCodePairValue).Value)
	case 30:
		this.verticesZ = append(this.verticesZ, codePair.Value.(DoubleCodePairValue).Value)
	case 77:
		this.OverrideColor = Color(codePair.Value.(ShortCodePairValue).Value)
	case 340:
		this.AssociatedAnnotationReference = codePair.Value.(StringCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	case 211:
		this.Right.X = codePair.Value.(DoubleCodePairValue).Value
	case 221:
		this.Right.Y = codePair.Value.(DoubleCodePairValue).Value
	case 231:
		this.Right.Z = codePair.Value.(DoubleCodePairValue).Value
	case 212:
		this.BlockOffset.X = codePair.Value.(DoubleCodePairValue).Value
	case 222:
		this.BlockOffset.Y = codePair.Value.(DoubleCodePairValue).Value
	case 232:
		this.BlockOffset.Z = codePair.Value.(DoubleCodePairValue).Value
	case 213:
		this.AnnotationOffset.X = codePair.Value.(DoubleCodePairValue).Value
	case 223:
		this.AnnotationOffset.Y = codePair.Value.(DoubleCodePairValue).Value
	case 233:
		this.AnnotationOffset.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Leader) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "LEADER"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbLeader"))
	pairs = append(pairs, NewStringCodePair(3, this.DimensionStyleName))
	pairs = append(pairs, NewShortCodePair(71, shortFromBool(this.UseArrowheads)))
	pairs = append(pairs, NewShortCodePair(72, int16(this.PathType)))
	pairs = append(pairs, NewShortCodePair(73, int16(this.AnnotationType)))
	pairs = append(pairs, NewShortCodePair(74, int16(this.HooklineDirection)))
	pairs = append(pairs, NewShortCodePair(75, shortFromBool(this.UseHookline)))
	pairs = append(pairs, NewDoubleCodePair(40, this.TextAnnotationHeight))
	pairs = append(pairs, NewDoubleCodePair(41, this.TextAnnotationWidth))
	pairs = append(pairs, NewShortCodePair(76, int16(len(this.Vertices))))
	for _, item := range this.Vertices {
		pairs = append(pairs, NewDoubleCodePair(10, item.X))
		pairs = append(pairs, NewDoubleCodePair(20, item.Y))
		pairs = append(pairs, NewDoubleCodePair(30, item.Z))
	}
	if this.OverrideColor != ByBlock() {
		pairs = append(pairs, NewShortCodePair(77, int16(this.OverrideColor)))
	}
	pairs = append(pairs, NewStringCodePair(340, this.AssociatedAnnotationReference))
	pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
	pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
	pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	pairs = append(pairs, NewDoubleCodePair(211, this.Right.X))
	pairs = append(pairs, NewDoubleCodePair(221, this.Right.Y))
	pairs = append(pairs, NewDoubleCodePair(231, this.Right.Z))
	pairs = append(pairs, NewDoubleCodePair(212, this.BlockOffset.X))
	pairs = append(pairs, NewDoubleCodePair(222, this.BlockOffset.Y))
	pairs = append(pairs, NewDoubleCodePair(232, this.BlockOffset.Z))
	if version >= R14 {
		pairs = append(pairs, NewDoubleCodePair(213, this.AnnotationOffset.X))
		pairs = append(pairs, NewDoubleCodePair(223, this.AnnotationOffset.Y))
		pairs = append(pairs, NewDoubleCodePair(233, this.AnnotationOffset.Z))
	}
	return
}

type Light struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	VersionNumber int
	Name string
	LightType LightType
	IsActive bool
	PlotGlyph bool
	Intensity float64
	Position Point
	TargetLocation Point
	AttentuationType LightAttenuationType
	UseAttenuationLimits bool
	AttenuationStartLimit float64
	AttenuationEndLimit float64
	HotspotAngle float64
	FalloffAngle float64
	CastShadows bool
	ShadowType ShadowType
	ShadowMapSize int
	ShadowMapSoftness int16
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewLight() *Light {
	return &Light{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		VersionNumber: 0,
		Name: "",
		LightType: LightTypeDistant,
		IsActive: true,
		PlotGlyph: true,
		Intensity: 1.0,
		Position: *NewOrigin(),
		TargetLocation: *NewOrigin(),
		AttentuationType: LightAttenuationTypeNone,
		UseAttenuationLimits: true,
		AttenuationStartLimit: 0.0,
		AttenuationEndLimit: 1.0,
		HotspotAngle: 0.0,
		FalloffAngle: 0.0,
		CastShadows: true,
		ShadowType: ShadowTypeRayTraced,
		ShadowMapSize: 0,
		ShadowMapSoftness: 0,
	}
}

func (e *Light) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Light) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Light) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Light) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Light) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Light) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Light) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Light) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Light) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Light) Handle() Handle {
	return this.handle
}

func (this *Light) SetHandle(val Handle) {
	this.handle = val
}

func (this *Light) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Light) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Light) Layer() string {
	return this.layer
}

func (this *Light) SetLayer(val string) {
	this.layer = val
}

func (this *Light) LineTypeName() string {
	return this.lineTypeName
}

func (this *Light) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Light) Elevation() float64 {
	return this.elevation
}

func (this *Light) SetElevation(val float64) {
	this.elevation = val
}

func (this *Light) MaterialHandle() string {
	return this.materialHandle
}

func (this *Light) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Light) Color() Color {
	return this.color
}

func (this *Light) SetColor(val Color) {
	this.color = val
}

func (this *Light) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Light) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Light) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Light) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Light) IsVisible() bool {
	return this.isVisible
}

func (this *Light) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Light) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Light) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Light) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Light) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Light) Color24Bit() int {
	return this.color24Bit
}

func (this *Light) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Light) ColorName() string {
	return this.colorName
}

func (this *Light) SetColorName(val string) {
	this.colorName = val
}

func (this *Light) Transparency() int {
	return this.transparency
}

func (this *Light) SetTransparency(val int) {
	this.transparency = val
}

func (this *Light) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Light) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Light) typeString() string {
	return "LIGHT"
}

func (this *Light) minVersion() (version AcadVersion) {
	return R2007
}

func (this *Light) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Light) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 90:
		this.VersionNumber = codePair.Value.(IntCodePairValue).Value
	case 1:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.LightType = LightType(codePair.Value.(ShortCodePairValue).Value)
	case 290:
		this.IsActive = codePair.Value.(BoolCodePairValue).Value
	case 291:
		this.PlotGlyph = codePair.Value.(BoolCodePairValue).Value
	case 40:
		this.Intensity = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.Position.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Position.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Position.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.TargetLocation.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.TargetLocation.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.TargetLocation.Z = codePair.Value.(DoubleCodePairValue).Value
	case 72:
		this.AttentuationType = LightAttenuationType(codePair.Value.(ShortCodePairValue).Value)
	case 292:
		this.UseAttenuationLimits = codePair.Value.(BoolCodePairValue).Value
	case 41:
		this.AttenuationStartLimit = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.AttenuationEndLimit = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.HotspotAngle = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.FalloffAngle = codePair.Value.(DoubleCodePairValue).Value
	case 293:
		this.CastShadows = codePair.Value.(BoolCodePairValue).Value
	case 73:
		this.ShadowType = ShadowType(codePair.Value.(ShortCodePairValue).Value)
	case 91:
		this.ShadowMapSize = codePair.Value.(IntCodePairValue).Value
	case 280:
		this.ShadowMapSoftness = codePair.Value.(ShortCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Light) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "LIGHT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbLight"))
	}
	pairs = append(pairs, NewIntCodePair(90, this.VersionNumber))
	pairs = append(pairs, NewStringCodePair(1, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.LightType)))
	pairs = append(pairs, NewBoolCodePair(290, this.IsActive))
	pairs = append(pairs, NewBoolCodePair(291, this.PlotGlyph))
	pairs = append(pairs, NewDoubleCodePair(40, this.Intensity))
	pairs = append(pairs, NewDoubleCodePair(10, this.Position.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Position.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Position.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.TargetLocation.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.TargetLocation.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.TargetLocation.Z))
	pairs = append(pairs, NewShortCodePair(72, int16(this.AttentuationType)))
	pairs = append(pairs, NewBoolCodePair(292, this.UseAttenuationLimits))
	pairs = append(pairs, NewDoubleCodePair(41, this.AttenuationStartLimit))
	pairs = append(pairs, NewDoubleCodePair(42, this.AttenuationEndLimit))
	pairs = append(pairs, NewDoubleCodePair(50, this.HotspotAngle))
	pairs = append(pairs, NewDoubleCodePair(51, this.FalloffAngle))
	pairs = append(pairs, NewBoolCodePair(293, this.CastShadows))
	pairs = append(pairs, NewShortCodePair(73, int16(this.ShadowType)))
	pairs = append(pairs, NewIntCodePair(91, this.ShadowMapSize))
	pairs = append(pairs, NewShortCodePair(280, this.ShadowMapSoftness))
	return
}

type Line struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	P1 Point
	P2 Point
	ExtrusionDirection Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewLine() *Line {
	return &Line{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		P1: *NewOrigin(),
		P2: *NewOrigin(),
		ExtrusionDirection: *NewZAxis(),
	}
}

func (e *Line) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Line) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Line) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Line) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Line) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Line) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Line) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Line) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Line) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Line) Handle() Handle {
	return this.handle
}

func (this *Line) SetHandle(val Handle) {
	this.handle = val
}

func (this *Line) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Line) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Line) Layer() string {
	return this.layer
}

func (this *Line) SetLayer(val string) {
	this.layer = val
}

func (this *Line) LineTypeName() string {
	return this.lineTypeName
}

func (this *Line) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Line) Elevation() float64 {
	return this.elevation
}

func (this *Line) SetElevation(val float64) {
	this.elevation = val
}

func (this *Line) MaterialHandle() string {
	return this.materialHandle
}

func (this *Line) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Line) Color() Color {
	return this.color
}

func (this *Line) SetColor(val Color) {
	this.color = val
}

func (this *Line) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Line) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Line) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Line) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Line) IsVisible() bool {
	return this.isVisible
}

func (this *Line) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Line) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Line) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Line) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Line) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Line) Color24Bit() int {
	return this.color24Bit
}

func (this *Line) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Line) ColorName() string {
	return this.colorName
}

func (this *Line) SetColorName(val string) {
	this.colorName = val
}

func (this *Line) Transparency() int {
	return this.transparency
}

func (this *Line) SetTransparency(val int) {
	this.transparency = val
}

func (this *Line) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Line) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Line) typeString() string {
	return "LINE"
}

func (this *Line) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Line) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Line) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.P1.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.P1.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.P1.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.P2.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.P2.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.P2.Z = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Line) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "LINE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbLine"))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.P1.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.P1.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.P1.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.P2.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.P2.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.P2.Z))
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	return
}

type LWPolyline struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	flags int
	ConstantWidth float64
	Thickness float64
	ExtrusionDirection Vector
	vertexCount int
	Vertices []LwVertex
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewLWPolyline() *LWPolyline {
	return &LWPolyline{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		flags: 0,
		ConstantWidth: 0.0,
		Thickness: 0.0,
		ExtrusionDirection: *NewZAxis(),
		vertexCount: 0,
		Vertices: []LwVertex{},
	}
}

func (e *LWPolyline) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *LWPolyline) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *LWPolyline) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *LWPolyline) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *LWPolyline) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *LWPolyline) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *LWPolyline) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *LWPolyline) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *LWPolyline) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *LWPolyline) Handle() Handle {
	return this.handle
}

func (this *LWPolyline) SetHandle(val Handle) {
	this.handle = val
}

func (this *LWPolyline) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *LWPolyline) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *LWPolyline) Layer() string {
	return this.layer
}

func (this *LWPolyline) SetLayer(val string) {
	this.layer = val
}

func (this *LWPolyline) LineTypeName() string {
	return this.lineTypeName
}

func (this *LWPolyline) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *LWPolyline) Elevation() float64 {
	return this.elevation
}

func (this *LWPolyline) SetElevation(val float64) {
	this.elevation = val
}

func (this *LWPolyline) MaterialHandle() string {
	return this.materialHandle
}

func (this *LWPolyline) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *LWPolyline) Color() Color {
	return this.color
}

func (this *LWPolyline) SetColor(val Color) {
	this.color = val
}

func (this *LWPolyline) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *LWPolyline) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *LWPolyline) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *LWPolyline) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *LWPolyline) IsVisible() bool {
	return this.isVisible
}

func (this *LWPolyline) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *LWPolyline) ImageByteCount() int {
	return this.imageByteCount
}

func (this *LWPolyline) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *LWPolyline) PreviewImageData() []string {
	return this.previewImageData
}

func (this *LWPolyline) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *LWPolyline) Color24Bit() int {
	return this.color24Bit
}

func (this *LWPolyline) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *LWPolyline) ColorName() string {
	return this.colorName
}

func (this *LWPolyline) SetColorName(val string) {
	this.colorName = val
}

func (this *LWPolyline) Transparency() int {
	return this.transparency
}

func (this *LWPolyline) SetTransparency(val int) {
	this.transparency = val
}

func (this *LWPolyline) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *LWPolyline) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsClosed status flag.
func (this *LWPolyline) IsClosed() bool {
	return this.flags & 1 != 0
}

// IsClosed status flag.
func (this *LWPolyline) SetIsClosed(val bool) {
	if val {
		this.flags = this.flags | 1
	} else {
		this.flags = this.flags & ^1
	}
}

// IsPLineGen status flag.
func (this *LWPolyline) IsPLineGen() bool {
	return this.flags & 128 != 0
}

// IsPLineGen status flag.
func (this *LWPolyline) SetIsPLineGen(val bool) {
	if val {
		this.flags = this.flags | 128
	} else {
		this.flags = this.flags & ^128
	}
}

func (this *LWPolyline) AddVertices(val LwVertex) {
	this.Vertices = append(this.Vertices, val)
}

func (this *LWPolyline) ClearVertices() {
	this.Vertices = []LwVertex{}
}

func (this *LWPolyline) typeString() string {
	return "LWPOLYLINE"
}

func (this *LWPolyline) minVersion() (version AcadVersion) {
	return R14
}

func (this *LWPolyline) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *LWPolyline) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "LWPOLYLINE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbPolyline"))
	pairs = append(pairs, NewIntCodePair(90, len(this.Vertices)))
	pairs = append(pairs, NewShortCodePair(70, int16(this.flags)))
	if this.ConstantWidth != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(43, this.ConstantWidth))
	}
	if this.Elevation() != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(38, this.Elevation()))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	for _, item := range this.Vertices {
		pairs = append(pairs, NewDoubleCodePair(10, item.X))
		pairs = append(pairs, NewDoubleCodePair(20, item.Y))
	if version >= R2013 && item.ID != 0 {
			pairs = append(pairs, NewIntCodePair(91, item.ID))
	}
	if item.StartingWidth != 0.0 {
			pairs = append(pairs, NewDoubleCodePair(40, item.StartingWidth))
	}
	if item.EndingWidth != 0.0 {
			pairs = append(pairs, NewDoubleCodePair(41, item.EndingWidth))
	}
	if item.Bulge != 0.0 {
			pairs = append(pairs, NewDoubleCodePair(42, item.Bulge))
	}
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	return
}

type MLine struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	StyleName string
	styleHandle string
	ScaleFactor float64
	Justification Justification
	Flags int
	vertexCount int
	StyleElementCount int
	StartPoint Point
	Normal Vector
	Vertices []Point
	vertexX []float64
	vertexY []float64
	vertexZ []float64
	SegmentDirections []Point
	segmentDirectionX []float64
	segmentDirectionY []float64
	segmentDirectionZ []float64
	MiterDirections []Point
	miterDirectionX []float64
	miterDirectionY []float64
	miterDirectionZ []float64
	parameterCount int
	Parameters []float64
	areaFillParameterCount int
	AreaFillParameters []float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewMLine() *MLine {
	return &MLine{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		StyleName: "",
		styleHandle: "",
		ScaleFactor: 1.0,
		Justification: JustificationTop,
		Flags: 0,
		vertexCount: 0,
		StyleElementCount: 0,
		StartPoint: *NewOrigin(),
		Normal: *NewZAxis(),
		Vertices: []Point{},
		vertexX: []float64{},
		vertexY: []float64{},
		vertexZ: []float64{},
		SegmentDirections: []Point{},
		segmentDirectionX: []float64{},
		segmentDirectionY: []float64{},
		segmentDirectionZ: []float64{},
		MiterDirections: []Point{},
		miterDirectionX: []float64{},
		miterDirectionY: []float64{},
		miterDirectionZ: []float64{},
		parameterCount: 0,
		Parameters: []float64{},
		areaFillParameterCount: 0,
		AreaFillParameters: []float64{},
	}
}

func (e *MLine) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *MLine) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *MLine) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *MLine) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *MLine) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *MLine) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *MLine) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *MLine) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *MLine) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *MLine) Handle() Handle {
	return this.handle
}

func (this *MLine) SetHandle(val Handle) {
	this.handle = val
}

func (this *MLine) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *MLine) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *MLine) Layer() string {
	return this.layer
}

func (this *MLine) SetLayer(val string) {
	this.layer = val
}

func (this *MLine) LineTypeName() string {
	return this.lineTypeName
}

func (this *MLine) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *MLine) Elevation() float64 {
	return this.elevation
}

func (this *MLine) SetElevation(val float64) {
	this.elevation = val
}

func (this *MLine) MaterialHandle() string {
	return this.materialHandle
}

func (this *MLine) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *MLine) Color() Color {
	return this.color
}

func (this *MLine) SetColor(val Color) {
	this.color = val
}

func (this *MLine) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *MLine) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *MLine) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *MLine) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *MLine) IsVisible() bool {
	return this.isVisible
}

func (this *MLine) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *MLine) ImageByteCount() int {
	return this.imageByteCount
}

func (this *MLine) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *MLine) PreviewImageData() []string {
	return this.previewImageData
}

func (this *MLine) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *MLine) Color24Bit() int {
	return this.color24Bit
}

func (this *MLine) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *MLine) ColorName() string {
	return this.colorName
}

func (this *MLine) SetColorName(val string) {
	this.colorName = val
}

func (this *MLine) Transparency() int {
	return this.transparency
}

func (this *MLine) SetTransparency(val int) {
	this.transparency = val
}

func (this *MLine) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *MLine) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// HasAtLeastOneVertex status flag.
func (this *MLine) HasAtLeastOneVertex() bool {
	return this.Flags & 1 != 0
}

// HasAtLeastOneVertex status flag.
func (this *MLine) SetHasAtLeastOneVertex(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// IsClosed status flag.
func (this *MLine) IsClosed() bool {
	return this.Flags & 2 != 0
}

// IsClosed status flag.
func (this *MLine) SetIsClosed(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// SuppressStartCaps status flag.
func (this *MLine) SuppressStartCaps() bool {
	return this.Flags & 4 != 0
}

// SuppressStartCaps status flag.
func (this *MLine) SetSuppressStartCaps(val bool) {
	if val {
		this.Flags = this.Flags | 4
	} else {
		this.Flags = this.Flags & ^4
	}
}

// SuppressEndCaps status flag.
func (this *MLine) SuppressEndCaps() bool {
	return this.Flags & 8 != 0
}

// SuppressEndCaps status flag.
func (this *MLine) SetSuppressEndCaps(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

func (this *MLine) AddVertices(val Point) {
	this.Vertices = append(this.Vertices, val)
}

func (this *MLine) ClearVertices() {
	this.Vertices = []Point{}
}

func (this *MLine) AddvertexX(val float64) {
	this.vertexX = append(this.vertexX, val)
}

func (this *MLine) ClearvertexX() {
	this.vertexX = []float64{}
}

func (this *MLine) AddvertexY(val float64) {
	this.vertexY = append(this.vertexY, val)
}

func (this *MLine) ClearvertexY() {
	this.vertexY = []float64{}
}

func (this *MLine) AddvertexZ(val float64) {
	this.vertexZ = append(this.vertexZ, val)
}

func (this *MLine) ClearvertexZ() {
	this.vertexZ = []float64{}
}

func (this *MLine) AddSegmentDirections(val Point) {
	this.SegmentDirections = append(this.SegmentDirections, val)
}

func (this *MLine) ClearSegmentDirections() {
	this.SegmentDirections = []Point{}
}

func (this *MLine) AddsegmentDirectionX(val float64) {
	this.segmentDirectionX = append(this.segmentDirectionX, val)
}

func (this *MLine) ClearsegmentDirectionX() {
	this.segmentDirectionX = []float64{}
}

func (this *MLine) AddsegmentDirectionY(val float64) {
	this.segmentDirectionY = append(this.segmentDirectionY, val)
}

func (this *MLine) ClearsegmentDirectionY() {
	this.segmentDirectionY = []float64{}
}

func (this *MLine) AddsegmentDirectionZ(val float64) {
	this.segmentDirectionZ = append(this.segmentDirectionZ, val)
}

func (this *MLine) ClearsegmentDirectionZ() {
	this.segmentDirectionZ = []float64{}
}

func (this *MLine) AddMiterDirections(val Point) {
	this.MiterDirections = append(this.MiterDirections, val)
}

func (this *MLine) ClearMiterDirections() {
	this.MiterDirections = []Point{}
}

func (this *MLine) AddmiterDirectionX(val float64) {
	this.miterDirectionX = append(this.miterDirectionX, val)
}

func (this *MLine) ClearmiterDirectionX() {
	this.miterDirectionX = []float64{}
}

func (this *MLine) AddmiterDirectionY(val float64) {
	this.miterDirectionY = append(this.miterDirectionY, val)
}

func (this *MLine) ClearmiterDirectionY() {
	this.miterDirectionY = []float64{}
}

func (this *MLine) AddmiterDirectionZ(val float64) {
	this.miterDirectionZ = append(this.miterDirectionZ, val)
}

func (this *MLine) ClearmiterDirectionZ() {
	this.miterDirectionZ = []float64{}
}

func (this *MLine) AddParameters(val float64) {
	this.Parameters = append(this.Parameters, val)
}

func (this *MLine) ClearParameters() {
	this.Parameters = []float64{}
}

func (this *MLine) AddAreaFillParameters(val float64) {
	this.AreaFillParameters = append(this.AreaFillParameters, val)
}

func (this *MLine) ClearAreaFillParameters() {
	this.AreaFillParameters = []float64{}
}

func (this *MLine) typeString() string {
	return "MLINE"
}

func (this *MLine) minVersion() (version AcadVersion) {
	return R13
}

func (this *MLine) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *MLine) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.StyleName = codePair.Value.(StringCodePairValue).Value
	case 340:
		this.styleHandle = codePair.Value.(StringCodePairValue).Value
	case 40:
		this.ScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 70:
		this.Justification = Justification(codePair.Value.(ShortCodePairValue).Value)
	case 71:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.vertexCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.StyleElementCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 10:
		this.StartPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.StartPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.StartPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.vertexX = append(this.vertexX, codePair.Value.(DoubleCodePairValue).Value)
	case 21:
		this.vertexY = append(this.vertexY, codePair.Value.(DoubleCodePairValue).Value)
	case 31:
		this.vertexZ = append(this.vertexZ, codePair.Value.(DoubleCodePairValue).Value)
	case 12:
		this.segmentDirectionX = append(this.segmentDirectionX, codePair.Value.(DoubleCodePairValue).Value)
	case 22:
		this.segmentDirectionY = append(this.segmentDirectionY, codePair.Value.(DoubleCodePairValue).Value)
	case 32:
		this.segmentDirectionZ = append(this.segmentDirectionZ, codePair.Value.(DoubleCodePairValue).Value)
	case 13:
		this.miterDirectionX = append(this.miterDirectionX, codePair.Value.(DoubleCodePairValue).Value)
	case 23:
		this.miterDirectionY = append(this.miterDirectionY, codePair.Value.(DoubleCodePairValue).Value)
	case 33:
		this.miterDirectionZ = append(this.miterDirectionZ, codePair.Value.(DoubleCodePairValue).Value)
	case 74:
		this.parameterCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 41:
		this.Parameters = append(this.Parameters, codePair.Value.(DoubleCodePairValue).Value)
	case 75:
		this.areaFillParameterCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 42:
		this.AreaFillParameters = append(this.AreaFillParameters, codePair.Value.(DoubleCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *MLine) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "MLINE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbMline"))
	pairs = append(pairs, NewStringCodePair(2, this.StyleName))
	if this.styleHandle != "" {
		pairs = append(pairs, NewStringCodePair(340, this.styleHandle))
	}
	pairs = append(pairs, NewDoubleCodePair(40, this.ScaleFactor))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Justification)))
	pairs = append(pairs, NewShortCodePair(71, int16(this.Flags)))
	pairs = append(pairs, NewShortCodePair(72, int16(len(this.Vertices))))
	pairs = append(pairs, NewShortCodePair(73, int16(this.StyleElementCount)))
	pairs = append(pairs, NewDoubleCodePair(10, this.StartPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.StartPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.StartPoint.Z))
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	for _, item := range this.Vertices {
		pairs = append(pairs, NewDoubleCodePair(10, item.X))
		pairs = append(pairs, NewDoubleCodePair(20, item.Y))
		pairs = append(pairs, NewDoubleCodePair(30, item.Z))
	}
	for _, item := range this.SegmentDirections {
		pairs = append(pairs, NewDoubleCodePair(11, item.X))
		pairs = append(pairs, NewDoubleCodePair(21, item.Y))
		pairs = append(pairs, NewDoubleCodePair(31, item.Z))
	}
	for _, item := range this.MiterDirections {
		pairs = append(pairs, NewDoubleCodePair(12, item.X))
		pairs = append(pairs, NewDoubleCodePair(22, item.Y))
		pairs = append(pairs, NewDoubleCodePair(32, item.Z))
	}
	pairs = append(pairs, NewShortCodePair(74, int16(len(this.Parameters))))
	for _, val := range this.Parameters {
		pairs = append(pairs, NewDoubleCodePair(41, val))
	}
	pairs = append(pairs, NewShortCodePair(75, int16(len(this.AreaFillParameters))))
	for _, val := range this.AreaFillParameters {
		pairs = append(pairs, NewDoubleCodePair(42, val))
	}
	return
}

type MText struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	InsertionPoint Point
	InitialTextHeight float64
	ReferenceRectangleWidth float64
	AttachmentPoint AttachmentPoint
	DrawingDirection DrawingDirection
	ExtendedText []string
	Text string
	TextStyleName string
	ExtrusionDirection Vector
	XAxisDirection Vector
	HorizontalWidth float64
	VerticalHeight float64
	RotationAngle float64
	LineSpacingStyle MTextLineSpacingStyle
	LineSpacingFactor float64
	BackgroundFillSetting BackgroundFillSetting
	BackgroundColorRGB int
	BackgroundColorName string
	FillBoxScale float64
	BackgroundFillColor Color
	BackgroundFillColorTransparency int
	ColumnType int16
	columnCount int16
	IsColumnFlowReversed bool
	IsColumnAutoHeight bool
	ColumnWidth float64
	ColumnGutter float64
	ColumnHeights []float64
	readingColumnData bool
	readColumnCount bool
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewMText() *MText {
	return &MText{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		InsertionPoint: *NewOrigin(),
		InitialTextHeight: 1.0,
		ReferenceRectangleWidth: 1.0,
		AttachmentPoint: AttachmentPointTopLeft,
		DrawingDirection: DrawingDirectionLeftToRight,
		ExtendedText: []string{},
		Text: "",
		TextStyleName: "STANDARD",
		ExtrusionDirection: *NewZAxis(),
		XAxisDirection: *NewXAxis(),
		HorizontalWidth: 1.0,
		VerticalHeight: 1.0,
		RotationAngle: 0.0,
		LineSpacingStyle: MTextLineSpacingStyleAtLeast,
		LineSpacingFactor: 1.0,
		BackgroundFillSetting: BackgroundFillSettingOff,
		BackgroundColorRGB: 0,
		BackgroundColorName: "",
		FillBoxScale: 1.0,
		BackgroundFillColor: ByLayer(),
		BackgroundFillColorTransparency: 0,
		ColumnType: 0,
		columnCount: 0,
		IsColumnFlowReversed: false,
		IsColumnAutoHeight: true,
		ColumnWidth: 0.0,
		ColumnGutter: 0.0,
		ColumnHeights: []float64{},
		readingColumnData: false,
		readColumnCount: false,
	}
}

func (e *MText) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *MText) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *MText) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *MText) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *MText) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *MText) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *MText) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *MText) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *MText) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *MText) Handle() Handle {
	return this.handle
}

func (this *MText) SetHandle(val Handle) {
	this.handle = val
}

func (this *MText) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *MText) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *MText) Layer() string {
	return this.layer
}

func (this *MText) SetLayer(val string) {
	this.layer = val
}

func (this *MText) LineTypeName() string {
	return this.lineTypeName
}

func (this *MText) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *MText) Elevation() float64 {
	return this.elevation
}

func (this *MText) SetElevation(val float64) {
	this.elevation = val
}

func (this *MText) MaterialHandle() string {
	return this.materialHandle
}

func (this *MText) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *MText) Color() Color {
	return this.color
}

func (this *MText) SetColor(val Color) {
	this.color = val
}

func (this *MText) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *MText) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *MText) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *MText) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *MText) IsVisible() bool {
	return this.isVisible
}

func (this *MText) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *MText) ImageByteCount() int {
	return this.imageByteCount
}

func (this *MText) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *MText) PreviewImageData() []string {
	return this.previewImageData
}

func (this *MText) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *MText) Color24Bit() int {
	return this.color24Bit
}

func (this *MText) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *MText) ColorName() string {
	return this.colorName
}

func (this *MText) SetColorName(val string) {
	this.colorName = val
}

func (this *MText) Transparency() int {
	return this.transparency
}

func (this *MText) SetTransparency(val int) {
	this.transparency = val
}

func (this *MText) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *MText) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *MText) AddExtendedText(val string) {
	this.ExtendedText = append(this.ExtendedText, val)
}

func (this *MText) ClearExtendedText() {
	this.ExtendedText = []string{}
}

func (this *MText) AddColumnHeights(val float64) {
	this.ColumnHeights = append(this.ColumnHeights, val)
}

func (this *MText) ClearColumnHeights() {
	this.ColumnHeights = []float64{}
}

func (this *MText) typeString() string {
	return "MTEXT"
}

func (this *MText) minVersion() (version AcadVersion) {
	return R13
}

func (this *MText) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *MText) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "MTEXT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbMText"))
	pairs = append(pairs, NewDoubleCodePair(10, this.InsertionPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.InsertionPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.InsertionPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.InitialTextHeight))
	pairs = append(pairs, NewDoubleCodePair(41, this.ReferenceRectangleWidth))
	pairs = append(pairs, NewShortCodePair(71, int16(this.AttachmentPoint)))
	pairs = append(pairs, NewShortCodePair(72, int16(this.DrawingDirection)))
	for _, val := range this.ExtendedText {
		pairs = append(pairs, NewStringCodePair(3, val))
	}
	pairs = append(pairs, NewStringCodePair(1, this.Text))
	if this.TextStyleName != "STANDARD" {
		pairs = append(pairs, NewStringCodePair(7, this.TextStyleName))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	pairs = append(pairs, NewDoubleCodePair(11, this.XAxisDirection.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.XAxisDirection.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.XAxisDirection.Z))
	pairs = append(pairs, NewDoubleCodePair(42, this.HorizontalWidth))
	pairs = append(pairs, NewDoubleCodePair(43, this.VerticalHeight))
	pairs = append(pairs, NewDoubleCodePair(50, this.RotationAngle))
	pairs = append(pairs, NewShortCodePair(73, int16(this.LineSpacingStyle)))
	pairs = append(pairs, NewDoubleCodePair(44, this.LineSpacingFactor))
	pairs = append(pairs, NewIntCodePair(90, int(this.BackgroundFillSetting)))
	pairs = append(pairs, NewIntCodePair(420, this.BackgroundColorRGB))
	pairs = append(pairs, NewStringCodePair(430, this.BackgroundColorName))
	if this.FillBoxScale != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(45, this.FillBoxScale))
	}
	pairs = append(pairs, NewShortCodePair(63, int16(this.BackgroundFillColor)))
	pairs = append(pairs, NewIntCodePair(441, this.BackgroundFillColorTransparency))
	pairs = append(pairs, NewShortCodePair(75, this.ColumnType))
	pairs = append(pairs, NewShortCodePair(76, int16(this.columnCount)))
	pairs = append(pairs, NewShortCodePair(78, shortFromBool(this.IsColumnFlowReversed)))
	pairs = append(pairs, NewShortCodePair(79, shortFromBool(this.IsColumnAutoHeight)))
	pairs = append(pairs, NewDoubleCodePair(48, this.ColumnWidth))
	pairs = append(pairs, NewDoubleCodePair(49, this.ColumnGutter))
	pairs = append(pairs, NewDoubleCodePair(50, float64(len(this.ColumnHeights))))
	for _, val := range this.ColumnHeights {
		pairs = append(pairs, NewDoubleCodePair(50, val))
	}
	return
}

type OleFrame struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	VersionNumber int
	binaryDataLength int
	binaryDataStrings []string
	BinaryData []byte
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewOleFrame() *OleFrame {
	return &OleFrame{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		VersionNumber: 0,
		binaryDataLength: 0,
		binaryDataStrings: []string{},
		BinaryData: []byte{},
	}
}

func (e *OleFrame) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *OleFrame) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *OleFrame) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *OleFrame) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *OleFrame) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *OleFrame) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *OleFrame) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *OleFrame) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *OleFrame) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *OleFrame) Handle() Handle {
	return this.handle
}

func (this *OleFrame) SetHandle(val Handle) {
	this.handle = val
}

func (this *OleFrame) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *OleFrame) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *OleFrame) Layer() string {
	return this.layer
}

func (this *OleFrame) SetLayer(val string) {
	this.layer = val
}

func (this *OleFrame) LineTypeName() string {
	return this.lineTypeName
}

func (this *OleFrame) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *OleFrame) Elevation() float64 {
	return this.elevation
}

func (this *OleFrame) SetElevation(val float64) {
	this.elevation = val
}

func (this *OleFrame) MaterialHandle() string {
	return this.materialHandle
}

func (this *OleFrame) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *OleFrame) Color() Color {
	return this.color
}

func (this *OleFrame) SetColor(val Color) {
	this.color = val
}

func (this *OleFrame) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *OleFrame) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *OleFrame) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *OleFrame) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *OleFrame) IsVisible() bool {
	return this.isVisible
}

func (this *OleFrame) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *OleFrame) ImageByteCount() int {
	return this.imageByteCount
}

func (this *OleFrame) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *OleFrame) PreviewImageData() []string {
	return this.previewImageData
}

func (this *OleFrame) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *OleFrame) Color24Bit() int {
	return this.color24Bit
}

func (this *OleFrame) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *OleFrame) ColorName() string {
	return this.colorName
}

func (this *OleFrame) SetColorName(val string) {
	this.colorName = val
}

func (this *OleFrame) Transparency() int {
	return this.transparency
}

func (this *OleFrame) SetTransparency(val int) {
	this.transparency = val
}

func (this *OleFrame) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *OleFrame) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *OleFrame) AddbinaryDataStrings(val string) {
	this.binaryDataStrings = append(this.binaryDataStrings, val)
}

func (this *OleFrame) ClearbinaryDataStrings() {
	this.binaryDataStrings = []string{}
}

func (this *OleFrame) AddBinaryData(val byte) {
	this.BinaryData = append(this.BinaryData, val)
}

func (this *OleFrame) ClearBinaryData() {
	this.BinaryData = []byte{}
}

func (this *OleFrame) typeString() string {
	return "OLEFRAME"
}

func (this *OleFrame) minVersion() (version AcadVersion) {
	return R13
}

func (this *OleFrame) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *OleFrame) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 70:
		this.VersionNumber = int(codePair.Value.(ShortCodePairValue).Value)
	case 90:
		this.binaryDataLength = codePair.Value.(IntCodePairValue).Value
	case 310:
		this.binaryDataStrings = append(this.binaryDataStrings, codePair.Value.(StringCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *OleFrame) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "OLEFRAME"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbOleFrame"))
	pairs = append(pairs, NewShortCodePair(70, int16(this.VersionNumber)))
	pairs = append(pairs, NewIntCodePair(90, this.binaryDataLength))
	for _, item := range this.binaryDataStrings {
		pairs = append(pairs, NewStringCodePair(310, item))
	}
	pairs = append(pairs, NewStringCodePair(1, "OLE"))
	return
}

type Ole2Frame struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	VersionNumber int
	Description string
	UpperLeftCorner Point
	LowerRightCorner Point
	ObjectType OleObjectType
	TileMode TileModeDescriptor
	binaryDataLength int
	binaryDataStrings []string
	BinaryData []byte
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewOle2Frame() *Ole2Frame {
	return &Ole2Frame{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		VersionNumber: 0,
		Description: "",
		UpperLeftCorner: *NewOrigin(),
		LowerRightCorner: *NewOrigin(),
		ObjectType: OleObjectTypeStatic,
		TileMode: TileModeDescriptorInTiledViewport,
		binaryDataLength: 0,
		binaryDataStrings: []string{},
		BinaryData: []byte{},
	}
}

func (e *Ole2Frame) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Ole2Frame) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Ole2Frame) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Ole2Frame) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Ole2Frame) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Ole2Frame) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Ole2Frame) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Ole2Frame) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Ole2Frame) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Ole2Frame) Handle() Handle {
	return this.handle
}

func (this *Ole2Frame) SetHandle(val Handle) {
	this.handle = val
}

func (this *Ole2Frame) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Ole2Frame) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Ole2Frame) Layer() string {
	return this.layer
}

func (this *Ole2Frame) SetLayer(val string) {
	this.layer = val
}

func (this *Ole2Frame) LineTypeName() string {
	return this.lineTypeName
}

func (this *Ole2Frame) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Ole2Frame) Elevation() float64 {
	return this.elevation
}

func (this *Ole2Frame) SetElevation(val float64) {
	this.elevation = val
}

func (this *Ole2Frame) MaterialHandle() string {
	return this.materialHandle
}

func (this *Ole2Frame) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Ole2Frame) Color() Color {
	return this.color
}

func (this *Ole2Frame) SetColor(val Color) {
	this.color = val
}

func (this *Ole2Frame) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Ole2Frame) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Ole2Frame) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Ole2Frame) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Ole2Frame) IsVisible() bool {
	return this.isVisible
}

func (this *Ole2Frame) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Ole2Frame) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Ole2Frame) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Ole2Frame) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Ole2Frame) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Ole2Frame) Color24Bit() int {
	return this.color24Bit
}

func (this *Ole2Frame) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Ole2Frame) ColorName() string {
	return this.colorName
}

func (this *Ole2Frame) SetColorName(val string) {
	this.colorName = val
}

func (this *Ole2Frame) Transparency() int {
	return this.transparency
}

func (this *Ole2Frame) SetTransparency(val int) {
	this.transparency = val
}

func (this *Ole2Frame) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Ole2Frame) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Ole2Frame) AddbinaryDataStrings(val string) {
	this.binaryDataStrings = append(this.binaryDataStrings, val)
}

func (this *Ole2Frame) ClearbinaryDataStrings() {
	this.binaryDataStrings = []string{}
}

func (this *Ole2Frame) AddBinaryData(val byte) {
	this.BinaryData = append(this.BinaryData, val)
}

func (this *Ole2Frame) ClearBinaryData() {
	this.BinaryData = []byte{}
}

func (this *Ole2Frame) typeString() string {
	return "OLE2FRAME"
}

func (this *Ole2Frame) minVersion() (version AcadVersion) {
	return R14
}

func (this *Ole2Frame) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Ole2Frame) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 70:
		this.VersionNumber = int(codePair.Value.(ShortCodePairValue).Value)
	case 3:
		this.Description = codePair.Value.(StringCodePairValue).Value
	case 10:
		this.UpperLeftCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.UpperLeftCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.UpperLeftCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.LowerRightCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.LowerRightCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.LowerRightCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.ObjectType = OleObjectType(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.TileMode = TileModeDescriptor(codePair.Value.(ShortCodePairValue).Value)
	case 90:
		this.binaryDataLength = codePair.Value.(IntCodePairValue).Value
	case 310:
		this.binaryDataStrings = append(this.binaryDataStrings, codePair.Value.(StringCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Ole2Frame) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "OLE2FRAME"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbOle2Frame"))
	pairs = append(pairs, NewShortCodePair(70, int16(this.VersionNumber)))
	pairs = append(pairs, NewStringCodePair(3, this.Description))
	pairs = append(pairs, NewDoubleCodePair(10, this.UpperLeftCorner.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.UpperLeftCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.UpperLeftCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.LowerRightCorner.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.LowerRightCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.LowerRightCorner.Z))
	pairs = append(pairs, NewShortCodePair(71, int16(this.ObjectType)))
	pairs = append(pairs, NewShortCodePair(72, int16(this.TileMode)))
	pairs = append(pairs, NewIntCodePair(90, this.binaryDataLength))
	for _, item := range this.binaryDataStrings {
		pairs = append(pairs, NewStringCodePair(310, item))
	}
	pairs = append(pairs, NewStringCodePair(1, "OLE"))
	return
}

type ModelPoint struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Location Point
	Thickness float64
	ExtrusionDirection Vector
	Angle float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewModelPoint() *ModelPoint {
	return &ModelPoint{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Location: *NewOrigin(),
		Thickness: 0.0,
		ExtrusionDirection: *NewZAxis(),
		Angle: 0.0,
	}
}

func (e *ModelPoint) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *ModelPoint) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *ModelPoint) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *ModelPoint) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *ModelPoint) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *ModelPoint) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *ModelPoint) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *ModelPoint) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *ModelPoint) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *ModelPoint) Handle() Handle {
	return this.handle
}

func (this *ModelPoint) SetHandle(val Handle) {
	this.handle = val
}

func (this *ModelPoint) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *ModelPoint) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *ModelPoint) Layer() string {
	return this.layer
}

func (this *ModelPoint) SetLayer(val string) {
	this.layer = val
}

func (this *ModelPoint) LineTypeName() string {
	return this.lineTypeName
}

func (this *ModelPoint) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *ModelPoint) Elevation() float64 {
	return this.elevation
}

func (this *ModelPoint) SetElevation(val float64) {
	this.elevation = val
}

func (this *ModelPoint) MaterialHandle() string {
	return this.materialHandle
}

func (this *ModelPoint) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *ModelPoint) Color() Color {
	return this.color
}

func (this *ModelPoint) SetColor(val Color) {
	this.color = val
}

func (this *ModelPoint) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *ModelPoint) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *ModelPoint) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *ModelPoint) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *ModelPoint) IsVisible() bool {
	return this.isVisible
}

func (this *ModelPoint) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *ModelPoint) ImageByteCount() int {
	return this.imageByteCount
}

func (this *ModelPoint) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *ModelPoint) PreviewImageData() []string {
	return this.previewImageData
}

func (this *ModelPoint) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *ModelPoint) Color24Bit() int {
	return this.color24Bit
}

func (this *ModelPoint) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *ModelPoint) ColorName() string {
	return this.colorName
}

func (this *ModelPoint) SetColorName(val string) {
	this.colorName = val
}

func (this *ModelPoint) Transparency() int {
	return this.transparency
}

func (this *ModelPoint) SetTransparency(val int) {
	this.transparency = val
}

func (this *ModelPoint) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *ModelPoint) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *ModelPoint) typeString() string {
	return "POINT"
}

func (this *ModelPoint) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *ModelPoint) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *ModelPoint) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.Location.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Location.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Location.Z = codePair.Value.(DoubleCodePairValue).Value
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.Angle = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *ModelPoint) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "POINT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbPoint"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	if this.Angle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.Angle))
	}
	return
}

type Polyline struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Location Point
	Thickness float64
	Flags int
	DefaultStartingWidth float64
	DefaultEndingWidth float64
	PolygonMeshMVertexCount int
	PolygonMeshNVertexCount int
	SmoothSurfaceMDensity int
	SmoothSurfaceNDensity int
	SurfaceType PolylineCurvedAndSmoothSurfaceType
	CLO_PolylineType PolylineType
	Normal Vector
	Vertices []Vertex
	seqend Seqend
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewPolyline() *Polyline {
	return &Polyline{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Location: *NewOrigin(),
		Thickness: 0.0,
		Flags: 0,
		DefaultStartingWidth: 0.0,
		DefaultEndingWidth: 0.0,
		PolygonMeshMVertexCount: 0,
		PolygonMeshNVertexCount: 0,
		SmoothSurfaceMDensity: 0,
		SmoothSurfaceNDensity: 0,
		SurfaceType: PolylineCurvedAndSmoothSurfaceTypeNone,
		CLO_PolylineType: PolylineTypeBaseline,
		Normal: *NewZAxis(),
		Vertices: []Vertex{},
		seqend: *NewSeqend(),
	}
}

func (e *Polyline) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Polyline) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Polyline) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Polyline) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Polyline) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Polyline) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Polyline) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Polyline) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Polyline) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Polyline) Handle() Handle {
	return this.handle
}

func (this *Polyline) SetHandle(val Handle) {
	this.handle = val
}

func (this *Polyline) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Polyline) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Polyline) Layer() string {
	return this.layer
}

func (this *Polyline) SetLayer(val string) {
	this.layer = val
}

func (this *Polyline) LineTypeName() string {
	return this.lineTypeName
}

func (this *Polyline) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Polyline) Elevation() float64 {
	return this.elevation
}

func (this *Polyline) SetElevation(val float64) {
	this.elevation = val
}

func (this *Polyline) MaterialHandle() string {
	return this.materialHandle
}

func (this *Polyline) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Polyline) Color() Color {
	return this.color
}

func (this *Polyline) SetColor(val Color) {
	this.color = val
}

func (this *Polyline) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Polyline) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Polyline) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Polyline) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Polyline) IsVisible() bool {
	return this.isVisible
}

func (this *Polyline) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Polyline) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Polyline) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Polyline) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Polyline) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Polyline) Color24Bit() int {
	return this.color24Bit
}

func (this *Polyline) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Polyline) ColorName() string {
	return this.colorName
}

func (this *Polyline) SetColorName(val string) {
	this.colorName = val
}

func (this *Polyline) Transparency() int {
	return this.transparency
}

func (this *Polyline) SetTransparency(val int) {
	this.transparency = val
}

func (this *Polyline) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Polyline) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsClosed status flag.
func (this *Polyline) IsClosed() bool {
	return this.Flags & 1 != 0
}

// IsClosed status flag.
func (this *Polyline) SetIsClosed(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// CurveFitVerticiesAdded status flag.
func (this *Polyline) CurveFitVerticiesAdded() bool {
	return this.Flags & 2 != 0
}

// CurveFitVerticiesAdded status flag.
func (this *Polyline) SetCurveFitVerticiesAdded(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// SplineFitVerticiesAdded status flag.
func (this *Polyline) SplineFitVerticiesAdded() bool {
	return this.Flags & 4 != 0
}

// SplineFitVerticiesAdded status flag.
func (this *Polyline) SetSplineFitVerticiesAdded(val bool) {
	if val {
		this.Flags = this.Flags | 4
	} else {
		this.Flags = this.Flags & ^4
	}
}

// Is3DPolyline status flag.
func (this *Polyline) Is3DPolyline() bool {
	return this.Flags & 8 != 0
}

// Is3DPolyline status flag.
func (this *Polyline) SetIs3DPolyline(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

// Is3DPolygonMesh status flag.
func (this *Polyline) Is3DPolygonMesh() bool {
	return this.Flags & 16 != 0
}

// Is3DPolygonMesh status flag.
func (this *Polyline) SetIs3DPolygonMesh(val bool) {
	if val {
		this.Flags = this.Flags | 16
	} else {
		this.Flags = this.Flags & ^16
	}
}

// IsPolygonMeshClosedInNDirection status flag.
func (this *Polyline) IsPolygonMeshClosedInNDirection() bool {
	return this.Flags & 32 != 0
}

// IsPolygonMeshClosedInNDirection status flag.
func (this *Polyline) SetIsPolygonMeshClosedInNDirection(val bool) {
	if val {
		this.Flags = this.Flags | 32
	} else {
		this.Flags = this.Flags & ^32
	}
}

// IsPolyfaceMesh status flag.
func (this *Polyline) IsPolyfaceMesh() bool {
	return this.Flags & 64 != 0
}

// IsPolyfaceMesh status flag.
func (this *Polyline) SetIsPolyfaceMesh(val bool) {
	if val {
		this.Flags = this.Flags | 64
	} else {
		this.Flags = this.Flags & ^64
	}
}

// IsLineTypePatternGeneratedContinuously status flag.
func (this *Polyline) IsLineTypePatternGeneratedContinuously() bool {
	return this.Flags & 128 != 0
}

// IsLineTypePatternGeneratedContinuously status flag.
func (this *Polyline) SetIsLineTypePatternGeneratedContinuously(val bool) {
	if val {
		this.Flags = this.Flags | 128
	} else {
		this.Flags = this.Flags & ^128
	}
}

func (this *Polyline) AddVertices(val Vertex) {
	this.Vertices = append(this.Vertices, val)
}

func (this *Polyline) ClearVertices() {
	this.Vertices = []Vertex{}
}

func (this *Polyline) typeString() string {
	return "POLYLINE"
}

func (this *Polyline) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Polyline) maxVersion() (version AcadVersion) {
	return R2018
}

type Ray struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	StartPoint Point
	UnitDirectionVector Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewRay() *Ray {
	return &Ray{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		StartPoint: *NewOrigin(),
		UnitDirectionVector: *NewXAxis(),
	}
}

func (e *Ray) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Ray) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Ray) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Ray) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Ray) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Ray) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Ray) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Ray) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Ray) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Ray) Handle() Handle {
	return this.handle
}

func (this *Ray) SetHandle(val Handle) {
	this.handle = val
}

func (this *Ray) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Ray) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Ray) Layer() string {
	return this.layer
}

func (this *Ray) SetLayer(val string) {
	this.layer = val
}

func (this *Ray) LineTypeName() string {
	return this.lineTypeName
}

func (this *Ray) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Ray) Elevation() float64 {
	return this.elevation
}

func (this *Ray) SetElevation(val float64) {
	this.elevation = val
}

func (this *Ray) MaterialHandle() string {
	return this.materialHandle
}

func (this *Ray) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Ray) Color() Color {
	return this.color
}

func (this *Ray) SetColor(val Color) {
	this.color = val
}

func (this *Ray) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Ray) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Ray) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Ray) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Ray) IsVisible() bool {
	return this.isVisible
}

func (this *Ray) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Ray) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Ray) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Ray) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Ray) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Ray) Color24Bit() int {
	return this.color24Bit
}

func (this *Ray) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Ray) ColorName() string {
	return this.colorName
}

func (this *Ray) SetColorName(val string) {
	this.colorName = val
}

func (this *Ray) Transparency() int {
	return this.transparency
}

func (this *Ray) SetTransparency(val int) {
	this.transparency = val
}

func (this *Ray) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Ray) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Ray) typeString() string {
	return "RAY"
}

func (this *Ray) minVersion() (version AcadVersion) {
	return R13
}

func (this *Ray) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Ray) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.StartPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.StartPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.StartPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.UnitDirectionVector.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.UnitDirectionVector.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.UnitDirectionVector.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Ray) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "RAY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbRay"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.StartPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.StartPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.StartPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.UnitDirectionVector.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.UnitDirectionVector.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.UnitDirectionVector.Z))
	return
}

type Region struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FormatVersionNumber int16
	CustomData []string
	CustomData2 []string
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewRegion() *Region {
	return &Region{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FormatVersionNumber: 1,
		CustomData: []string{},
		CustomData2: []string{},
	}
}

func (e *Region) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Region) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Region) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Region) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Region) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Region) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Region) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Region) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Region) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Region) Handle() Handle {
	return this.handle
}

func (this *Region) SetHandle(val Handle) {
	this.handle = val
}

func (this *Region) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Region) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Region) Layer() string {
	return this.layer
}

func (this *Region) SetLayer(val string) {
	this.layer = val
}

func (this *Region) LineTypeName() string {
	return this.lineTypeName
}

func (this *Region) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Region) Elevation() float64 {
	return this.elevation
}

func (this *Region) SetElevation(val float64) {
	this.elevation = val
}

func (this *Region) MaterialHandle() string {
	return this.materialHandle
}

func (this *Region) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Region) Color() Color {
	return this.color
}

func (this *Region) SetColor(val Color) {
	this.color = val
}

func (this *Region) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Region) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Region) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Region) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Region) IsVisible() bool {
	return this.isVisible
}

func (this *Region) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Region) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Region) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Region) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Region) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Region) Color24Bit() int {
	return this.color24Bit
}

func (this *Region) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Region) ColorName() string {
	return this.colorName
}

func (this *Region) SetColorName(val string) {
	this.colorName = val
}

func (this *Region) Transparency() int {
	return this.transparency
}

func (this *Region) SetTransparency(val int) {
	this.transparency = val
}

func (this *Region) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Region) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Region) AddCustomData(val string) {
	this.CustomData = append(this.CustomData, val)
}

func (this *Region) ClearCustomData() {
	this.CustomData = []string{}
}

func (this *Region) AddCustomData2(val string) {
	this.CustomData2 = append(this.CustomData2, val)
}

func (this *Region) ClearCustomData2() {
	this.CustomData2 = []string{}
}

func (this *Region) typeString() string {
	return "REGION"
}

func (this *Region) minVersion() (version AcadVersion) {
	return R13
}

func (this *Region) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Region) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 70:
		this.FormatVersionNumber = codePair.Value.(ShortCodePairValue).Value
	case 1:
		this.CustomData = append(this.CustomData, codePair.Value.(StringCodePairValue).Value)
	case 3:
		this.CustomData2 = append(this.CustomData2, codePair.Value.(StringCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Region) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "REGION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbModelerGeometry"))
	}
	pairs = append(pairs, NewShortCodePair(70, this.FormatVersionNumber))
	for _, val := range this.CustomData {
		pairs = append(pairs, NewStringCodePair(1, val))
	}
	for _, val := range this.CustomData2 {
		pairs = append(pairs, NewStringCodePair(3, val))
	}
	return
}

type RText struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	InsertionPoint Point
	ExtrusionDirection Vector
	RotationAngle float64
	TextHeight float64
	TextStyle string
	TypeFlags int
	Contents string
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewRText() *RText {
	return &RText{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		InsertionPoint: *NewOrigin(),
		ExtrusionDirection: *NewZAxis(),
		RotationAngle: 0.0,
		TextHeight: 0.0,
		TextStyle: "STANDARD",
		TypeFlags: 0,
		Contents: "",
	}
}

func (e *RText) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *RText) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *RText) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *RText) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *RText) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *RText) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *RText) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *RText) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *RText) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *RText) Handle() Handle {
	return this.handle
}

func (this *RText) SetHandle(val Handle) {
	this.handle = val
}

func (this *RText) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *RText) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *RText) Layer() string {
	return this.layer
}

func (this *RText) SetLayer(val string) {
	this.layer = val
}

func (this *RText) LineTypeName() string {
	return this.lineTypeName
}

func (this *RText) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *RText) Elevation() float64 {
	return this.elevation
}

func (this *RText) SetElevation(val float64) {
	this.elevation = val
}

func (this *RText) MaterialHandle() string {
	return this.materialHandle
}

func (this *RText) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *RText) Color() Color {
	return this.color
}

func (this *RText) SetColor(val Color) {
	this.color = val
}

func (this *RText) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *RText) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *RText) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *RText) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *RText) IsVisible() bool {
	return this.isVisible
}

func (this *RText) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *RText) ImageByteCount() int {
	return this.imageByteCount
}

func (this *RText) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *RText) PreviewImageData() []string {
	return this.previewImageData
}

func (this *RText) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *RText) Color24Bit() int {
	return this.color24Bit
}

func (this *RText) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *RText) ColorName() string {
	return this.colorName
}

func (this *RText) SetColorName(val string) {
	this.colorName = val
}

func (this *RText) Transparency() int {
	return this.transparency
}

func (this *RText) SetTransparency(val int) {
	this.transparency = val
}

func (this *RText) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *RText) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsExpression status flag.
func (this *RText) IsExpression() bool {
	return this.TypeFlags & 1 != 0
}

// IsExpression status flag.
func (this *RText) SetIsExpression(val bool) {
	if val {
		this.TypeFlags = this.TypeFlags | 1
	} else {
		this.TypeFlags = this.TypeFlags & ^1
	}
}

// IsInlineMTextSequencesEnabled status flag.
func (this *RText) IsInlineMTextSequencesEnabled() bool {
	return this.TypeFlags & 2 != 0
}

// IsInlineMTextSequencesEnabled status flag.
func (this *RText) SetIsInlineMTextSequencesEnabled(val bool) {
	if val {
		this.TypeFlags = this.TypeFlags | 2
	} else {
		this.TypeFlags = this.TypeFlags & ^2
	}
}

func (this *RText) typeString() string {
	return "RTEXT"
}

func (this *RText) minVersion() (version AcadVersion) {
	return R2000
}

func (this *RText) maxVersion() (version AcadVersion) {
	return R2000
}

func (this *RText) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.InsertionPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.InsertionPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.InsertionPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.RotationAngle = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.TextHeight = codePair.Value.(DoubleCodePairValue).Value
	case 7:
		this.TextStyle = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.TypeFlags = int(codePair.Value.(ShortCodePairValue).Value)
	case 1:
		this.Contents = codePair.Value.(StringCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *RText) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "RTEXT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "RText"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.InsertionPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.InsertionPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.InsertionPoint.Z))
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	pairs = append(pairs, NewDoubleCodePair(50, this.RotationAngle))
	pairs = append(pairs, NewDoubleCodePair(40, this.TextHeight))
	pairs = append(pairs, NewStringCodePair(7, this.TextStyle))
	pairs = append(pairs, NewShortCodePair(70, int16(this.TypeFlags)))
	pairs = append(pairs, NewStringCodePair(1, this.Contents))
	return
}

type Section struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	State int
	Flags int
	Name string
	VerticalDirection Vector
	TopHeight float64
	BottomHeight float64
	IndicatorTransparency int16
	IndicatorColor Color
	IndicatorColorName string
	vertexCount int
	Vertices []Point
	backLineVertexCount int
	BackLineVertices []Point
	GeometrySettingsHandle string
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewSection() *Section {
	return &Section{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		State: 0,
		Flags: 0,
		Name: "",
		VerticalDirection: *NewZAxis(),
		TopHeight: 0.0,
		BottomHeight: 0.0,
		IndicatorTransparency: 0,
		IndicatorColor: ByLayer(),
		IndicatorColorName: "",
		vertexCount: 0,
		Vertices: []Point{},
		backLineVertexCount: 0,
		BackLineVertices: []Point{},
		GeometrySettingsHandle: "",
	}
}

func (e *Section) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Section) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Section) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Section) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Section) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Section) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Section) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Section) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Section) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Section) Handle() Handle {
	return this.handle
}

func (this *Section) SetHandle(val Handle) {
	this.handle = val
}

func (this *Section) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Section) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Section) Layer() string {
	return this.layer
}

func (this *Section) SetLayer(val string) {
	this.layer = val
}

func (this *Section) LineTypeName() string {
	return this.lineTypeName
}

func (this *Section) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Section) Elevation() float64 {
	return this.elevation
}

func (this *Section) SetElevation(val float64) {
	this.elevation = val
}

func (this *Section) MaterialHandle() string {
	return this.materialHandle
}

func (this *Section) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Section) Color() Color {
	return this.color
}

func (this *Section) SetColor(val Color) {
	this.color = val
}

func (this *Section) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Section) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Section) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Section) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Section) IsVisible() bool {
	return this.isVisible
}

func (this *Section) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Section) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Section) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Section) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Section) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Section) Color24Bit() int {
	return this.color24Bit
}

func (this *Section) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Section) ColorName() string {
	return this.colorName
}

func (this *Section) SetColorName(val string) {
	this.colorName = val
}

func (this *Section) Transparency() int {
	return this.transparency
}

func (this *Section) SetTransparency(val int) {
	this.transparency = val
}

func (this *Section) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Section) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Section) AddVertices(val Point) {
	this.Vertices = append(this.Vertices, val)
}

func (this *Section) ClearVertices() {
	this.Vertices = []Point{}
}

func (this *Section) AddBackLineVertices(val Point) {
	this.BackLineVertices = append(this.BackLineVertices, val)
}

func (this *Section) ClearBackLineVertices() {
	this.BackLineVertices = []Point{}
}

func (this *Section) typeString() string {
	return "SECTION"
}

func (this *Section) minVersion() (version AcadVersion) {
	return R2007
}

func (this *Section) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Section) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "SECTION"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbSection"))
	pairs = append(pairs, NewIntCodePair(90, this.State))
	pairs = append(pairs, NewIntCodePair(91, this.Flags))
	pairs = append(pairs, NewStringCodePair(1, this.Name))
	pairs = append(pairs, NewDoubleCodePair(10, this.VerticalDirection.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.VerticalDirection.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.VerticalDirection.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.TopHeight))
	pairs = append(pairs, NewDoubleCodePair(41, this.BottomHeight))
	pairs = append(pairs, NewShortCodePair(70, this.IndicatorTransparency))
	pairs = append(pairs, NewShortCodePair(63, int16(this.IndicatorColor)))
	pairs = append(pairs, NewStringCodePair(411, this.IndicatorColorName))
	pairs = append(pairs, NewIntCodePair(92, len(this.Vertices)))
	for _, item := range this.Vertices {
		pairs = append(pairs, NewDoubleCodePair(11, item.X))
		pairs = append(pairs, NewDoubleCodePair(21, item.Y))
		pairs = append(pairs, NewDoubleCodePair(31, item.Z))
	}
	pairs = append(pairs, NewIntCodePair(93, len(this.BackLineVertices)))
	for _, item := range this.BackLineVertices {
		pairs = append(pairs, NewDoubleCodePair(12, item.X))
		pairs = append(pairs, NewDoubleCodePair(22, item.Y))
		pairs = append(pairs, NewDoubleCodePair(32, item.Z))
	}
	if this.GeometrySettingsHandle != "" {
		pairs = append(pairs, NewStringCodePair(360, this.GeometrySettingsHandle))
	}
	return
}

type Seqend struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewSeqend() *Seqend {
	return &Seqend{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
	}
}

func (e *Seqend) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Seqend) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Seqend) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Seqend) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Seqend) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Seqend) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Seqend) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Seqend) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Seqend) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Seqend) Handle() Handle {
	return this.handle
}

func (this *Seqend) SetHandle(val Handle) {
	this.handle = val
}

func (this *Seqend) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Seqend) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Seqend) Layer() string {
	return this.layer
}

func (this *Seqend) SetLayer(val string) {
	this.layer = val
}

func (this *Seqend) LineTypeName() string {
	return this.lineTypeName
}

func (this *Seqend) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Seqend) Elevation() float64 {
	return this.elevation
}

func (this *Seqend) SetElevation(val float64) {
	this.elevation = val
}

func (this *Seqend) MaterialHandle() string {
	return this.materialHandle
}

func (this *Seqend) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Seqend) Color() Color {
	return this.color
}

func (this *Seqend) SetColor(val Color) {
	this.color = val
}

func (this *Seqend) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Seqend) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Seqend) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Seqend) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Seqend) IsVisible() bool {
	return this.isVisible
}

func (this *Seqend) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Seqend) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Seqend) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Seqend) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Seqend) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Seqend) Color24Bit() int {
	return this.color24Bit
}

func (this *Seqend) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Seqend) ColorName() string {
	return this.colorName
}

func (this *Seqend) SetColorName(val string) {
	this.colorName = val
}

func (this *Seqend) Transparency() int {
	return this.transparency
}

func (this *Seqend) SetTransparency(val int) {
	this.transparency = val
}

func (this *Seqend) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Seqend) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Seqend) typeString() string {
	return "SEQEND"
}

func (this *Seqend) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Seqend) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Seqend) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "SEQEND"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	return
}

type Shape struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Location Point
	Size float64
	Name string
	RotationAngle float64
	RelativeXScaleFactor float64
	ObliqueAngle float64
	ExtrusionDirection Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewShape() *Shape {
	return &Shape{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Location: *NewOrigin(),
		Size: 0.0,
		Name: "",
		RotationAngle: 0.0,
		RelativeXScaleFactor: 1.0,
		ObliqueAngle: 0.0,
		ExtrusionDirection: *NewZAxis(),
	}
}

func (e *Shape) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Shape) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Shape) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Shape) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Shape) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Shape) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Shape) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Shape) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Shape) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Shape) Handle() Handle {
	return this.handle
}

func (this *Shape) SetHandle(val Handle) {
	this.handle = val
}

func (this *Shape) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Shape) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Shape) Layer() string {
	return this.layer
}

func (this *Shape) SetLayer(val string) {
	this.layer = val
}

func (this *Shape) LineTypeName() string {
	return this.lineTypeName
}

func (this *Shape) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Shape) Elevation() float64 {
	return this.elevation
}

func (this *Shape) SetElevation(val float64) {
	this.elevation = val
}

func (this *Shape) MaterialHandle() string {
	return this.materialHandle
}

func (this *Shape) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Shape) Color() Color {
	return this.color
}

func (this *Shape) SetColor(val Color) {
	this.color = val
}

func (this *Shape) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Shape) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Shape) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Shape) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Shape) IsVisible() bool {
	return this.isVisible
}

func (this *Shape) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Shape) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Shape) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Shape) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Shape) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Shape) Color24Bit() int {
	return this.color24Bit
}

func (this *Shape) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Shape) ColorName() string {
	return this.colorName
}

func (this *Shape) SetColorName(val string) {
	this.colorName = val
}

func (this *Shape) Transparency() int {
	return this.transparency
}

func (this *Shape) SetTransparency(val int) {
	this.transparency = val
}

func (this *Shape) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Shape) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Shape) typeString() string {
	return "SHAPE"
}

func (this *Shape) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Shape) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Shape) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.Location.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Location.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Location.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.Size = codePair.Value.(DoubleCodePairValue).Value
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 50:
		this.RotationAngle = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.RelativeXScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.ObliqueAngle = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Shape) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "SHAPE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbShape"))
	}
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.Size))
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	if this.RotationAngle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.RotationAngle))
	}
	if this.RelativeXScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.RelativeXScaleFactor))
	}
	if this.ObliqueAngle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(51, this.ObliqueAngle))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	return
}

type Solid struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FirstCorner Point
	SecondCorner Point
	ThirdCorner Point
	FourthCorner Point
	Thickness float64
	ExtrusionDirection Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewSolid() *Solid {
	return &Solid{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FirstCorner: *NewOrigin(),
		SecondCorner: *NewOrigin(),
		ThirdCorner: *NewOrigin(),
		FourthCorner: *NewOrigin(),
		Thickness: 0.0,
		ExtrusionDirection: *NewZAxis(),
	}
}

func (e *Solid) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Solid) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Solid) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Solid) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Solid) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Solid) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Solid) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Solid) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Solid) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Solid) Handle() Handle {
	return this.handle
}

func (this *Solid) SetHandle(val Handle) {
	this.handle = val
}

func (this *Solid) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Solid) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Solid) Layer() string {
	return this.layer
}

func (this *Solid) SetLayer(val string) {
	this.layer = val
}

func (this *Solid) LineTypeName() string {
	return this.lineTypeName
}

func (this *Solid) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Solid) Elevation() float64 {
	return this.elevation
}

func (this *Solid) SetElevation(val float64) {
	this.elevation = val
}

func (this *Solid) MaterialHandle() string {
	return this.materialHandle
}

func (this *Solid) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Solid) Color() Color {
	return this.color
}

func (this *Solid) SetColor(val Color) {
	this.color = val
}

func (this *Solid) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Solid) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Solid) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Solid) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Solid) IsVisible() bool {
	return this.isVisible
}

func (this *Solid) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Solid) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Solid) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Solid) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Solid) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Solid) Color24Bit() int {
	return this.color24Bit
}

func (this *Solid) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Solid) ColorName() string {
	return this.colorName
}

func (this *Solid) SetColorName(val string) {
	this.colorName = val
}

func (this *Solid) Transparency() int {
	return this.transparency
}

func (this *Solid) SetTransparency(val int) {
	this.transparency = val
}

func (this *Solid) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Solid) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Solid) typeString() string {
	return "SOLID"
}

func (this *Solid) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Solid) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Solid) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.FirstCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.FirstCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.FirstCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.SecondCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.SecondCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.SecondCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.ThirdCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.ThirdCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.ThirdCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 13:
		this.FourthCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.FourthCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.FourthCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Solid) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "SOLID"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbTrace"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.FirstCorner.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.FirstCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.FirstCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.SecondCorner.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.SecondCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.SecondCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.ThirdCorner.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.ThirdCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.ThirdCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.FourthCorner.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.FourthCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.FourthCorner.Z))
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	return
}

type Spline struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Normal Vector
	Flags int
	DegreeOfCurve int
	numberOfKnots int
	numberOfControlPoints int
	numberOfFitPoints int
	KnotTolerance float64
	ControlPointTolerance float64
	FitTolerance float64
	StartTangent Vector
	EndTangent Vector
	KnotValues []float64
	weights []float64
	ControlPoints []ControlPoint
	FitPoints []Point
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewSpline() *Spline {
	return &Spline{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Normal: *NewZAxis(),
		Flags: 0,
		DegreeOfCurve: 1,
		numberOfKnots: 0,
		numberOfControlPoints: 0,
		numberOfFitPoints: 0,
		KnotTolerance: 0.0000001,
		ControlPointTolerance: 0.0000001,
		FitTolerance: 0.0000001,
		StartTangent: *NewXAxis(),
		EndTangent: *NewXAxis(),
		KnotValues: []float64{},
		weights: []float64{},
		ControlPoints: []ControlPoint{},
		FitPoints: []Point{},
	}
}

func (e *Spline) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Spline) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Spline) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Spline) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Spline) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Spline) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Spline) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Spline) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Spline) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Spline) Handle() Handle {
	return this.handle
}

func (this *Spline) SetHandle(val Handle) {
	this.handle = val
}

func (this *Spline) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Spline) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Spline) Layer() string {
	return this.layer
}

func (this *Spline) SetLayer(val string) {
	this.layer = val
}

func (this *Spline) LineTypeName() string {
	return this.lineTypeName
}

func (this *Spline) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Spline) Elevation() float64 {
	return this.elevation
}

func (this *Spline) SetElevation(val float64) {
	this.elevation = val
}

func (this *Spline) MaterialHandle() string {
	return this.materialHandle
}

func (this *Spline) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Spline) Color() Color {
	return this.color
}

func (this *Spline) SetColor(val Color) {
	this.color = val
}

func (this *Spline) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Spline) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Spline) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Spline) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Spline) IsVisible() bool {
	return this.isVisible
}

func (this *Spline) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Spline) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Spline) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Spline) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Spline) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Spline) Color24Bit() int {
	return this.color24Bit
}

func (this *Spline) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Spline) ColorName() string {
	return this.colorName
}

func (this *Spline) SetColorName(val string) {
	this.colorName = val
}

func (this *Spline) Transparency() int {
	return this.transparency
}

func (this *Spline) SetTransparency(val int) {
	this.transparency = val
}

func (this *Spline) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Spline) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsClosed status flag.
func (this *Spline) IsClosed() bool {
	return this.Flags & 1 != 0
}

// IsClosed status flag.
func (this *Spline) SetIsClosed(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// IsPeriodic status flag.
func (this *Spline) IsPeriodic() bool {
	return this.Flags & 2 != 0
}

// IsPeriodic status flag.
func (this *Spline) SetIsPeriodic(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// IsRational status flag.
func (this *Spline) IsRational() bool {
	return this.Flags & 4 != 0
}

// IsRational status flag.
func (this *Spline) SetIsRational(val bool) {
	if val {
		this.Flags = this.Flags | 4
	} else {
		this.Flags = this.Flags & ^4
	}
}

// IsPlanar status flag.
func (this *Spline) IsPlanar() bool {
	return this.Flags & 8 != 0
}

// IsPlanar status flag.
func (this *Spline) SetIsPlanar(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

// IsLinear status flag.
func (this *Spline) IsLinear() bool {
	return this.Flags & 16 != 0
}

// IsLinear status flag.
func (this *Spline) SetIsLinear(val bool) {
	if val {
		this.Flags = this.Flags | 16
	} else {
		this.Flags = this.Flags & ^16
	}
}

func (this *Spline) AddKnotValues(val float64) {
	this.KnotValues = append(this.KnotValues, val)
}

func (this *Spline) ClearKnotValues() {
	this.KnotValues = []float64{}
}

func (this *Spline) Addweights(val float64) {
	this.weights = append(this.weights, val)
}

func (this *Spline) Clearweights() {
	this.weights = []float64{}
}

func (this *Spline) AddControlPoints(val ControlPoint) {
	this.ControlPoints = append(this.ControlPoints, val)
}

func (this *Spline) ClearControlPoints() {
	this.ControlPoints = []ControlPoint{}
}

func (this *Spline) AddFitPoints(val Point) {
	this.FitPoints = append(this.FitPoints, val)
}

func (this *Spline) ClearFitPoints() {
	this.FitPoints = []Point{}
}

func (this *Spline) typeString() string {
	return "SPLINE"
}

func (this *Spline) minVersion() (version AcadVersion) {
	return R13
}

func (this *Spline) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Spline) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "SPLINE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbSpline"))
	pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
	pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
	pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewShortCodePair(71, int16(this.DegreeOfCurve)))
	pairs = append(pairs, NewShortCodePair(72, int16(len(this.KnotValues))))
	pairs = append(pairs, NewShortCodePair(73, int16(len(this.ControlPoints))))
	pairs = append(pairs, NewShortCodePair(74, int16(len(this.FitPoints))))
	if this.KnotTolerance != 0.0000001 {
		pairs = append(pairs, NewDoubleCodePair(42, this.KnotTolerance))
	}
	if this.ControlPointTolerance != 0.0000001 {
		pairs = append(pairs, NewDoubleCodePair(43, this.ControlPointTolerance))
	}
	if this.FitTolerance != 0.0000001 {
		pairs = append(pairs, NewDoubleCodePair(44, this.FitTolerance))
	}
	pairs = append(pairs, NewDoubleCodePair(12, this.StartTangent.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.StartTangent.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.StartTangent.Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.EndTangent.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.EndTangent.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.EndTangent.Z))
	for _, val := range this.KnotValues {
		pairs = append(pairs, NewDoubleCodePair(40, val))
	}
	if this.shouldWriteWeights() {
		for _, item := range this.ControlPoints {
			pairs = append(pairs, NewDoubleCodePair(41, item.Weight))
		}
	}
	for _, item := range this.ControlPoints {
		pairs = append(pairs, NewDoubleCodePair(10, item.Point.X))
		pairs = append(pairs, NewDoubleCodePair(20, item.Point.Y))
		pairs = append(pairs, NewDoubleCodePair(30, item.Point.Z))
	}
	for _, item := range this.FitPoints {
		pairs = append(pairs, NewDoubleCodePair(11, item.X))
		pairs = append(pairs, NewDoubleCodePair(21, item.Y))
		pairs = append(pairs, NewDoubleCodePair(31, item.Z))
	}
	return
}

type Text struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Thickness float64
	Location Point
	Height float64
	Value string
	Rotation float64
	RelativeXScaleFactor float64
	ObliqueAngle float64
	TextStyleName string
	TextGenerationFlags int
	HorizontalTextJustification HorizontalTextJustification
	SecondAlignmentPoint Point
	Normal Vector
	VerticalTextJustification VerticalTextJustification
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewText() *Text {
	return &Text{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Thickness: 0.0,
		Location: *NewOrigin(),
		Height: 1.0,
		Value: "",
		Rotation: 0.0,
		RelativeXScaleFactor: 1.0,
		ObliqueAngle: 0.0,
		TextStyleName: "STANDARD",
		TextGenerationFlags: 0,
		HorizontalTextJustification: HorizontalTextJustificationLeft,
		SecondAlignmentPoint: *NewOrigin(),
		Normal: *NewZAxis(),
		VerticalTextJustification: VerticalTextJustificationBaseline,
	}
}

func (e *Text) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Text) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Text) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Text) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Text) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Text) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Text) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Text) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Text) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Text) Handle() Handle {
	return this.handle
}

func (this *Text) SetHandle(val Handle) {
	this.handle = val
}

func (this *Text) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Text) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Text) Layer() string {
	return this.layer
}

func (this *Text) SetLayer(val string) {
	this.layer = val
}

func (this *Text) LineTypeName() string {
	return this.lineTypeName
}

func (this *Text) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Text) Elevation() float64 {
	return this.elevation
}

func (this *Text) SetElevation(val float64) {
	this.elevation = val
}

func (this *Text) MaterialHandle() string {
	return this.materialHandle
}

func (this *Text) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Text) Color() Color {
	return this.color
}

func (this *Text) SetColor(val Color) {
	this.color = val
}

func (this *Text) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Text) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Text) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Text) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Text) IsVisible() bool {
	return this.isVisible
}

func (this *Text) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Text) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Text) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Text) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Text) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Text) Color24Bit() int {
	return this.color24Bit
}

func (this *Text) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Text) ColorName() string {
	return this.colorName
}

func (this *Text) SetColorName(val string) {
	this.colorName = val
}

func (this *Text) Transparency() int {
	return this.transparency
}

func (this *Text) SetTransparency(val int) {
	this.transparency = val
}

func (this *Text) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Text) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsTextBackwards status flag.
func (this *Text) IsTextBackwards() bool {
	return this.TextGenerationFlags & 2 != 0
}

// IsTextBackwards status flag.
func (this *Text) SetIsTextBackwards(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 2
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^2
	}
}

// IsTextUpsideDown status flag.
func (this *Text) IsTextUpsideDown() bool {
	return this.TextGenerationFlags & 4 != 0
}

// IsTextUpsideDown status flag.
func (this *Text) SetIsTextUpsideDown(val bool) {
	if val {
		this.TextGenerationFlags = this.TextGenerationFlags | 4
	} else {
		this.TextGenerationFlags = this.TextGenerationFlags & ^4
	}
}

func (this *Text) typeString() string {
	return "TEXT"
}

func (this *Text) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Text) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Text) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.Location.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Location.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Location.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.Height = codePair.Value.(DoubleCodePairValue).Value
	case 1:
		this.Value = codePair.Value.(StringCodePairValue).Value
	case 50:
		this.Rotation = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.RelativeXScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.ObliqueAngle = codePair.Value.(DoubleCodePairValue).Value
	case 7:
		this.TextStyleName = codePair.Value.(StringCodePairValue).Value
	case 71:
		this.TextGenerationFlags = int(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.HorizontalTextJustification = HorizontalTextJustification(codePair.Value.(ShortCodePairValue).Value)
	case 11:
		this.SecondAlignmentPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.SecondAlignmentPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.SecondAlignmentPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.Normal.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.Normal.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.Normal.Z = codePair.Value.(DoubleCodePairValue).Value
	case 73:
		this.VerticalTextJustification = VerticalTextJustification(codePair.Value.(ShortCodePairValue).Value)
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Text) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TEXT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, NewStringCodePair(100, "AcDbText"))
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.Location.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Location.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Location.Z))
	pairs = append(pairs, NewDoubleCodePair(40, this.Height))
	pairs = append(pairs, NewStringCodePair(1, this.Value))
	if this.Rotation != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(50, this.Rotation))
	}
	if this.RelativeXScaleFactor != 1.0 {
		pairs = append(pairs, NewDoubleCodePair(41, this.RelativeXScaleFactor))
	}
	if this.ObliqueAngle != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(51, this.ObliqueAngle))
	}
	if this.TextStyleName != "STANDARD" {
		pairs = append(pairs, NewStringCodePair(7, this.TextStyleName))
	}
	if this.TextGenerationFlags != 0 {
		pairs = append(pairs, NewShortCodePair(71, int16(this.TextGenerationFlags)))
	}
	if this.HorizontalTextJustification != HorizontalTextJustificationLeft {
		pairs = append(pairs, NewShortCodePair(72, int16(this.HorizontalTextJustification)))
	}
	if this.SecondAlignmentPoint != *NewOrigin() {
		pairs = append(pairs, NewDoubleCodePair(11, this.SecondAlignmentPoint.X))
		pairs = append(pairs, NewDoubleCodePair(21, this.SecondAlignmentPoint.Y))
		pairs = append(pairs, NewDoubleCodePair(31, this.SecondAlignmentPoint.Z))
	}
	if this.Normal != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.Normal.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.Normal.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.Normal.Z))
	}
	if this.VerticalTextJustification != VerticalTextJustificationBaseline {
		pairs = append(pairs, NewShortCodePair(73, int16(this.VerticalTextJustification)))
	}
	pairs = append(pairs, NewStringCodePair(100, "AcDbText"))
	return
}

type Tolerance struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	DimensionStyleName string
	InsertionPoint Point
	DisplayText string
	ExtrusionDirection Vector
	DirectionVector Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewTolerance() *Tolerance {
	return &Tolerance{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		DimensionStyleName: "",
		InsertionPoint: *NewOrigin(),
		DisplayText: "",
		ExtrusionDirection: *NewZAxis(),
		DirectionVector: *NewXAxis(),
	}
}

func (e *Tolerance) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Tolerance) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Tolerance) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Tolerance) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Tolerance) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Tolerance) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Tolerance) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Tolerance) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Tolerance) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Tolerance) Handle() Handle {
	return this.handle
}

func (this *Tolerance) SetHandle(val Handle) {
	this.handle = val
}

func (this *Tolerance) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Tolerance) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Tolerance) Layer() string {
	return this.layer
}

func (this *Tolerance) SetLayer(val string) {
	this.layer = val
}

func (this *Tolerance) LineTypeName() string {
	return this.lineTypeName
}

func (this *Tolerance) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Tolerance) Elevation() float64 {
	return this.elevation
}

func (this *Tolerance) SetElevation(val float64) {
	this.elevation = val
}

func (this *Tolerance) MaterialHandle() string {
	return this.materialHandle
}

func (this *Tolerance) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Tolerance) Color() Color {
	return this.color
}

func (this *Tolerance) SetColor(val Color) {
	this.color = val
}

func (this *Tolerance) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Tolerance) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Tolerance) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Tolerance) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Tolerance) IsVisible() bool {
	return this.isVisible
}

func (this *Tolerance) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Tolerance) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Tolerance) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Tolerance) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Tolerance) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Tolerance) Color24Bit() int {
	return this.color24Bit
}

func (this *Tolerance) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Tolerance) ColorName() string {
	return this.colorName
}

func (this *Tolerance) SetColorName(val string) {
	this.colorName = val
}

func (this *Tolerance) Transparency() int {
	return this.transparency
}

func (this *Tolerance) SetTransparency(val int) {
	this.transparency = val
}

func (this *Tolerance) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Tolerance) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Tolerance) typeString() string {
	return "TOLERANCE"
}

func (this *Tolerance) minVersion() (version AcadVersion) {
	return R13
}

func (this *Tolerance) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Tolerance) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 3:
		this.DimensionStyleName = codePair.Value.(StringCodePairValue).Value
	case 10:
		this.InsertionPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.InsertionPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.InsertionPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 1:
		this.DisplayText = codePair.Value.(StringCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.DirectionVector.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.DirectionVector.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.DirectionVector.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Tolerance) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TOLERANCE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbFcf"))
	}
	pairs = append(pairs, NewStringCodePair(3, this.DimensionStyleName))
	pairs = append(pairs, NewDoubleCodePair(10, this.InsertionPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.InsertionPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.InsertionPoint.Z))
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(1, this.DisplayText))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	pairs = append(pairs, NewDoubleCodePair(11, this.DirectionVector.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.DirectionVector.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.DirectionVector.Z))
	return
}

type Trace struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FirstCorner Point
	SecondCorner Point
	ThirdCorner Point
	FourthCorner Point
	Thickness float64
	ExtrusionDirection Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewTrace() *Trace {
	return &Trace{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FirstCorner: *NewOrigin(),
		SecondCorner: *NewOrigin(),
		ThirdCorner: *NewOrigin(),
		FourthCorner: *NewOrigin(),
		Thickness: 0.0,
		ExtrusionDirection: *NewZAxis(),
	}
}

func (e *Trace) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Trace) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Trace) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Trace) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Trace) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Trace) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Trace) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Trace) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Trace) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Trace) Handle() Handle {
	return this.handle
}

func (this *Trace) SetHandle(val Handle) {
	this.handle = val
}

func (this *Trace) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Trace) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Trace) Layer() string {
	return this.layer
}

func (this *Trace) SetLayer(val string) {
	this.layer = val
}

func (this *Trace) LineTypeName() string {
	return this.lineTypeName
}

func (this *Trace) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Trace) Elevation() float64 {
	return this.elevation
}

func (this *Trace) SetElevation(val float64) {
	this.elevation = val
}

func (this *Trace) MaterialHandle() string {
	return this.materialHandle
}

func (this *Trace) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Trace) Color() Color {
	return this.color
}

func (this *Trace) SetColor(val Color) {
	this.color = val
}

func (this *Trace) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Trace) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Trace) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Trace) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Trace) IsVisible() bool {
	return this.isVisible
}

func (this *Trace) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Trace) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Trace) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Trace) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Trace) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Trace) Color24Bit() int {
	return this.color24Bit
}

func (this *Trace) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Trace) ColorName() string {
	return this.colorName
}

func (this *Trace) SetColorName(val string) {
	this.colorName = val
}

func (this *Trace) Transparency() int {
	return this.transparency
}

func (this *Trace) SetTransparency(val int) {
	this.transparency = val
}

func (this *Trace) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Trace) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Trace) typeString() string {
	return "TRACE"
}

func (this *Trace) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Trace) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Trace) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.FirstCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.FirstCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.FirstCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.SecondCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.SecondCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.SecondCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.ThirdCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.ThirdCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.ThirdCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 13:
		this.FourthCorner.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.FourthCorner.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.FourthCorner.Z = codePair.Value.(DoubleCodePairValue).Value
	case 39:
		this.Thickness = codePair.Value.(DoubleCodePairValue).Value
	case 210:
		this.ExtrusionDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 220:
		this.ExtrusionDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 230:
		this.ExtrusionDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Trace) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TRACE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbTrace"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.FirstCorner.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.FirstCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.FirstCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.SecondCorner.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.SecondCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.SecondCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.ThirdCorner.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.ThirdCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.ThirdCorner.Z))
	pairs = append(pairs, NewDoubleCodePair(13, this.FourthCorner.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.FourthCorner.Y))
	pairs = append(pairs, NewDoubleCodePair(33, this.FourthCorner.Z))
	if this.Thickness != 0.0 {
		pairs = append(pairs, NewDoubleCodePair(39, this.Thickness))
	}
	if this.ExtrusionDirection != *NewZAxis() {
		pairs = append(pairs, NewDoubleCodePair(210, this.ExtrusionDirection.X))
		pairs = append(pairs, NewDoubleCodePair(220, this.ExtrusionDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(230, this.ExtrusionDirection.Z))
	}
	return
}

type DgnUnderlay struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Underlay interface
	objectHandle string
	insertionPoint Point
	xScale float64
	yScale float64
	zScale float64
	rotationAngle float64
	normal Vector
	flags int
	contrast int16
	fade int16
	boundaryPoints []Point
	_pointX []float64
	_pointY []float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewDgnUnderlay() *DgnUnderlay {
	return &DgnUnderlay{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		objectHandle: "",
		insertionPoint: *NewOrigin(),
		xScale: 1.0,
		yScale: 1.0,
		zScale: 1.0,
		rotationAngle: 0.0,
		normal: *NewZAxis(),
		flags: 0,
		contrast: 100,
		fade: 0,
		boundaryPoints: []Point{},
		_pointX: []float64{},
		_pointY: []float64{},
	}
}

func (e *DgnUnderlay) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *DgnUnderlay) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *DgnUnderlay) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *DgnUnderlay) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *DgnUnderlay) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *DgnUnderlay) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *DgnUnderlay) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *DgnUnderlay) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *DgnUnderlay) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *DgnUnderlay) Handle() Handle {
	return this.handle
}

func (this *DgnUnderlay) SetHandle(val Handle) {
	this.handle = val
}

func (this *DgnUnderlay) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *DgnUnderlay) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *DgnUnderlay) Layer() string {
	return this.layer
}

func (this *DgnUnderlay) SetLayer(val string) {
	this.layer = val
}

func (this *DgnUnderlay) LineTypeName() string {
	return this.lineTypeName
}

func (this *DgnUnderlay) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *DgnUnderlay) Elevation() float64 {
	return this.elevation
}

func (this *DgnUnderlay) SetElevation(val float64) {
	this.elevation = val
}

func (this *DgnUnderlay) MaterialHandle() string {
	return this.materialHandle
}

func (this *DgnUnderlay) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *DgnUnderlay) Color() Color {
	return this.color
}

func (this *DgnUnderlay) SetColor(val Color) {
	this.color = val
}

func (this *DgnUnderlay) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *DgnUnderlay) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *DgnUnderlay) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *DgnUnderlay) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *DgnUnderlay) IsVisible() bool {
	return this.isVisible
}

func (this *DgnUnderlay) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *DgnUnderlay) ImageByteCount() int {
	return this.imageByteCount
}

func (this *DgnUnderlay) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *DgnUnderlay) PreviewImageData() []string {
	return this.previewImageData
}

func (this *DgnUnderlay) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *DgnUnderlay) Color24Bit() int {
	return this.color24Bit
}

func (this *DgnUnderlay) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *DgnUnderlay) ColorName() string {
	return this.colorName
}

func (this *DgnUnderlay) SetColorName(val string) {
	this.colorName = val
}

func (this *DgnUnderlay) Transparency() int {
	return this.transparency
}

func (this *DgnUnderlay) SetTransparency(val int) {
	this.transparency = val
}

func (this *DgnUnderlay) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *DgnUnderlay) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *DgnUnderlay) ObjectHandle() string {
	return this.objectHandle
}

func (this *DgnUnderlay) SetObjectHandle(val string) {
	this.objectHandle = val
}

func (this *DgnUnderlay) InsertionPoint() Point {
	return this.insertionPoint
}

func (this *DgnUnderlay) SetInsertionPoint(val Point) {
	this.insertionPoint = val
}

func (this *DgnUnderlay) XScale() float64 {
	return this.xScale
}

func (this *DgnUnderlay) SetXScale(val float64) {
	this.xScale = val
}

func (this *DgnUnderlay) YScale() float64 {
	return this.yScale
}

func (this *DgnUnderlay) SetYScale(val float64) {
	this.yScale = val
}

func (this *DgnUnderlay) ZScale() float64 {
	return this.zScale
}

func (this *DgnUnderlay) SetZScale(val float64) {
	this.zScale = val
}

func (this *DgnUnderlay) RotationAngle() float64 {
	return this.rotationAngle
}

func (this *DgnUnderlay) SetRotationAngle(val float64) {
	this.rotationAngle = val
}

func (this *DgnUnderlay) Normal() Vector {
	return this.normal
}

func (this *DgnUnderlay) SetNormal(val Vector) {
	this.normal = val
}

func (this *DgnUnderlay) Flags() int {
	return this.flags
}

func (this *DgnUnderlay) SetFlags(val int) {
	this.flags = val
}

func (this *DgnUnderlay) Contrast() int16 {
	return this.contrast
}

func (this *DgnUnderlay) SetContrast(val int16) {
	this.contrast = val
}

func (this *DgnUnderlay) Fade() int16 {
	return this.fade
}

func (this *DgnUnderlay) SetFade(val int16) {
	this.fade = val
}

func (this *DgnUnderlay) BoundaryPoints() []Point {
	return this.boundaryPoints
}

func (this *DgnUnderlay) SetBoundaryPoints(val []Point) {
	this.boundaryPoints = val
}

func (this *DgnUnderlay) pointX() []float64 {
	return this._pointX
}

func (this *DgnUnderlay) SetpointX(val []float64) {
	this._pointX = val
}

func (this *DgnUnderlay) pointY() []float64 {
	return this._pointY
}

func (this *DgnUnderlay) SetpointY(val []float64) {
	this._pointY = val
}

func (this *DgnUnderlay) typeString() string {
	return "DGNUNDERLAY"
}

func (this *DgnUnderlay) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *DgnUnderlay) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *DgnUnderlay) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForUnderlay(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *DgnUnderlay) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DGNUNDERLAY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForUnderlay(this, version)...)
	return
}

type DwfUnderlay struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Underlay interface
	objectHandle string
	insertionPoint Point
	xScale float64
	yScale float64
	zScale float64
	rotationAngle float64
	normal Vector
	flags int
	contrast int16
	fade int16
	boundaryPoints []Point
	_pointX []float64
	_pointY []float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewDwfUnderlay() *DwfUnderlay {
	return &DwfUnderlay{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		objectHandle: "",
		insertionPoint: *NewOrigin(),
		xScale: 1.0,
		yScale: 1.0,
		zScale: 1.0,
		rotationAngle: 0.0,
		normal: *NewZAxis(),
		flags: 0,
		contrast: 100,
		fade: 0,
		boundaryPoints: []Point{},
		_pointX: []float64{},
		_pointY: []float64{},
	}
}

func (e *DwfUnderlay) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *DwfUnderlay) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *DwfUnderlay) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *DwfUnderlay) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *DwfUnderlay) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *DwfUnderlay) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *DwfUnderlay) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *DwfUnderlay) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *DwfUnderlay) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *DwfUnderlay) Handle() Handle {
	return this.handle
}

func (this *DwfUnderlay) SetHandle(val Handle) {
	this.handle = val
}

func (this *DwfUnderlay) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *DwfUnderlay) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *DwfUnderlay) Layer() string {
	return this.layer
}

func (this *DwfUnderlay) SetLayer(val string) {
	this.layer = val
}

func (this *DwfUnderlay) LineTypeName() string {
	return this.lineTypeName
}

func (this *DwfUnderlay) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *DwfUnderlay) Elevation() float64 {
	return this.elevation
}

func (this *DwfUnderlay) SetElevation(val float64) {
	this.elevation = val
}

func (this *DwfUnderlay) MaterialHandle() string {
	return this.materialHandle
}

func (this *DwfUnderlay) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *DwfUnderlay) Color() Color {
	return this.color
}

func (this *DwfUnderlay) SetColor(val Color) {
	this.color = val
}

func (this *DwfUnderlay) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *DwfUnderlay) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *DwfUnderlay) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *DwfUnderlay) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *DwfUnderlay) IsVisible() bool {
	return this.isVisible
}

func (this *DwfUnderlay) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *DwfUnderlay) ImageByteCount() int {
	return this.imageByteCount
}

func (this *DwfUnderlay) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *DwfUnderlay) PreviewImageData() []string {
	return this.previewImageData
}

func (this *DwfUnderlay) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *DwfUnderlay) Color24Bit() int {
	return this.color24Bit
}

func (this *DwfUnderlay) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *DwfUnderlay) ColorName() string {
	return this.colorName
}

func (this *DwfUnderlay) SetColorName(val string) {
	this.colorName = val
}

func (this *DwfUnderlay) Transparency() int {
	return this.transparency
}

func (this *DwfUnderlay) SetTransparency(val int) {
	this.transparency = val
}

func (this *DwfUnderlay) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *DwfUnderlay) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *DwfUnderlay) ObjectHandle() string {
	return this.objectHandle
}

func (this *DwfUnderlay) SetObjectHandle(val string) {
	this.objectHandle = val
}

func (this *DwfUnderlay) InsertionPoint() Point {
	return this.insertionPoint
}

func (this *DwfUnderlay) SetInsertionPoint(val Point) {
	this.insertionPoint = val
}

func (this *DwfUnderlay) XScale() float64 {
	return this.xScale
}

func (this *DwfUnderlay) SetXScale(val float64) {
	this.xScale = val
}

func (this *DwfUnderlay) YScale() float64 {
	return this.yScale
}

func (this *DwfUnderlay) SetYScale(val float64) {
	this.yScale = val
}

func (this *DwfUnderlay) ZScale() float64 {
	return this.zScale
}

func (this *DwfUnderlay) SetZScale(val float64) {
	this.zScale = val
}

func (this *DwfUnderlay) RotationAngle() float64 {
	return this.rotationAngle
}

func (this *DwfUnderlay) SetRotationAngle(val float64) {
	this.rotationAngle = val
}

func (this *DwfUnderlay) Normal() Vector {
	return this.normal
}

func (this *DwfUnderlay) SetNormal(val Vector) {
	this.normal = val
}

func (this *DwfUnderlay) Flags() int {
	return this.flags
}

func (this *DwfUnderlay) SetFlags(val int) {
	this.flags = val
}

func (this *DwfUnderlay) Contrast() int16 {
	return this.contrast
}

func (this *DwfUnderlay) SetContrast(val int16) {
	this.contrast = val
}

func (this *DwfUnderlay) Fade() int16 {
	return this.fade
}

func (this *DwfUnderlay) SetFade(val int16) {
	this.fade = val
}

func (this *DwfUnderlay) BoundaryPoints() []Point {
	return this.boundaryPoints
}

func (this *DwfUnderlay) SetBoundaryPoints(val []Point) {
	this.boundaryPoints = val
}

func (this *DwfUnderlay) pointX() []float64 {
	return this._pointX
}

func (this *DwfUnderlay) SetpointX(val []float64) {
	this._pointX = val
}

func (this *DwfUnderlay) pointY() []float64 {
	return this._pointY
}

func (this *DwfUnderlay) SetpointY(val []float64) {
	this._pointY = val
}

func (this *DwfUnderlay) typeString() string {
	return "DWFUNDERLAY"
}

func (this *DwfUnderlay) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *DwfUnderlay) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *DwfUnderlay) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForUnderlay(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *DwfUnderlay) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "DWFUNDERLAY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForUnderlay(this, version)...)
	return
}

type PdfUnderlay struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for Underlay interface
	objectHandle string
	insertionPoint Point
	xScale float64
	yScale float64
	zScale float64
	rotationAngle float64
	normal Vector
	flags int
	contrast int16
	fade int16
	boundaryPoints []Point
	_pointX []float64
	_pointY []float64
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewPdfUnderlay() *PdfUnderlay {
	return &PdfUnderlay{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		objectHandle: "",
		insertionPoint: *NewOrigin(),
		xScale: 1.0,
		yScale: 1.0,
		zScale: 1.0,
		rotationAngle: 0.0,
		normal: *NewZAxis(),
		flags: 0,
		contrast: 100,
		fade: 0,
		boundaryPoints: []Point{},
		_pointX: []float64{},
		_pointY: []float64{},
	}
}

func (e *PdfUnderlay) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *PdfUnderlay) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *PdfUnderlay) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *PdfUnderlay) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *PdfUnderlay) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *PdfUnderlay) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *PdfUnderlay) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *PdfUnderlay) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *PdfUnderlay) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *PdfUnderlay) Handle() Handle {
	return this.handle
}

func (this *PdfUnderlay) SetHandle(val Handle) {
	this.handle = val
}

func (this *PdfUnderlay) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *PdfUnderlay) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *PdfUnderlay) Layer() string {
	return this.layer
}

func (this *PdfUnderlay) SetLayer(val string) {
	this.layer = val
}

func (this *PdfUnderlay) LineTypeName() string {
	return this.lineTypeName
}

func (this *PdfUnderlay) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *PdfUnderlay) Elevation() float64 {
	return this.elevation
}

func (this *PdfUnderlay) SetElevation(val float64) {
	this.elevation = val
}

func (this *PdfUnderlay) MaterialHandle() string {
	return this.materialHandle
}

func (this *PdfUnderlay) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *PdfUnderlay) Color() Color {
	return this.color
}

func (this *PdfUnderlay) SetColor(val Color) {
	this.color = val
}

func (this *PdfUnderlay) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *PdfUnderlay) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *PdfUnderlay) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *PdfUnderlay) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *PdfUnderlay) IsVisible() bool {
	return this.isVisible
}

func (this *PdfUnderlay) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *PdfUnderlay) ImageByteCount() int {
	return this.imageByteCount
}

func (this *PdfUnderlay) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *PdfUnderlay) PreviewImageData() []string {
	return this.previewImageData
}

func (this *PdfUnderlay) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *PdfUnderlay) Color24Bit() int {
	return this.color24Bit
}

func (this *PdfUnderlay) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *PdfUnderlay) ColorName() string {
	return this.colorName
}

func (this *PdfUnderlay) SetColorName(val string) {
	this.colorName = val
}

func (this *PdfUnderlay) Transparency() int {
	return this.transparency
}

func (this *PdfUnderlay) SetTransparency(val int) {
	this.transparency = val
}

func (this *PdfUnderlay) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *PdfUnderlay) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *PdfUnderlay) ObjectHandle() string {
	return this.objectHandle
}

func (this *PdfUnderlay) SetObjectHandle(val string) {
	this.objectHandle = val
}

func (this *PdfUnderlay) InsertionPoint() Point {
	return this.insertionPoint
}

func (this *PdfUnderlay) SetInsertionPoint(val Point) {
	this.insertionPoint = val
}

func (this *PdfUnderlay) XScale() float64 {
	return this.xScale
}

func (this *PdfUnderlay) SetXScale(val float64) {
	this.xScale = val
}

func (this *PdfUnderlay) YScale() float64 {
	return this.yScale
}

func (this *PdfUnderlay) SetYScale(val float64) {
	this.yScale = val
}

func (this *PdfUnderlay) ZScale() float64 {
	return this.zScale
}

func (this *PdfUnderlay) SetZScale(val float64) {
	this.zScale = val
}

func (this *PdfUnderlay) RotationAngle() float64 {
	return this.rotationAngle
}

func (this *PdfUnderlay) SetRotationAngle(val float64) {
	this.rotationAngle = val
}

func (this *PdfUnderlay) Normal() Vector {
	return this.normal
}

func (this *PdfUnderlay) SetNormal(val Vector) {
	this.normal = val
}

func (this *PdfUnderlay) Flags() int {
	return this.flags
}

func (this *PdfUnderlay) SetFlags(val int) {
	this.flags = val
}

func (this *PdfUnderlay) Contrast() int16 {
	return this.contrast
}

func (this *PdfUnderlay) SetContrast(val int16) {
	this.contrast = val
}

func (this *PdfUnderlay) Fade() int16 {
	return this.fade
}

func (this *PdfUnderlay) SetFade(val int16) {
	this.fade = val
}

func (this *PdfUnderlay) BoundaryPoints() []Point {
	return this.boundaryPoints
}

func (this *PdfUnderlay) SetBoundaryPoints(val []Point) {
	this.boundaryPoints = val
}

func (this *PdfUnderlay) pointX() []float64 {
	return this._pointX
}

func (this *PdfUnderlay) SetpointX(val []float64) {
	this._pointX = val
}

func (this *PdfUnderlay) pointY() []float64 {
	return this._pointY
}

func (this *PdfUnderlay) SetpointY(val []float64) {
	this._pointY = val
}

func (this *PdfUnderlay) typeString() string {
	return "PDFUNDERLAY"
}

func (this *PdfUnderlay) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *PdfUnderlay) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *PdfUnderlay) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForUnderlay(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *PdfUnderlay) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "PDFUNDERLAY"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForUnderlay(this, version)...)
	return
}

type Vertex struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	Location Point
	StartingWidth float64
	EndingWidth float64
	Bulge float64
	Flags int
	CurveFitTangentDirection float64
	PolyfaceMeshVertexIndex1 int
	PolyfaceMeshVertexIndex2 int
	PolyfaceMeshVertexIndex3 int
	PolyfaceMeshVertexIndex4 int
	Identifier int
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewVertex() *Vertex {
	return &Vertex{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		Location: *NewOrigin(),
		StartingWidth: 0.0,
		EndingWidth: 0.0,
		Bulge: 0.0,
		Flags: 0,
		CurveFitTangentDirection: 0.0,
		PolyfaceMeshVertexIndex1: 0,
		PolyfaceMeshVertexIndex2: 0,
		PolyfaceMeshVertexIndex3: 0,
		PolyfaceMeshVertexIndex4: 0,
		Identifier: 0,
	}
}

func (e *Vertex) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Vertex) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Vertex) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Vertex) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Vertex) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Vertex) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Vertex) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Vertex) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Vertex) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Vertex) Handle() Handle {
	return this.handle
}

func (this *Vertex) SetHandle(val Handle) {
	this.handle = val
}

func (this *Vertex) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Vertex) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Vertex) Layer() string {
	return this.layer
}

func (this *Vertex) SetLayer(val string) {
	this.layer = val
}

func (this *Vertex) LineTypeName() string {
	return this.lineTypeName
}

func (this *Vertex) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Vertex) Elevation() float64 {
	return this.elevation
}

func (this *Vertex) SetElevation(val float64) {
	this.elevation = val
}

func (this *Vertex) MaterialHandle() string {
	return this.materialHandle
}

func (this *Vertex) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Vertex) Color() Color {
	return this.color
}

func (this *Vertex) SetColor(val Color) {
	this.color = val
}

func (this *Vertex) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Vertex) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Vertex) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Vertex) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Vertex) IsVisible() bool {
	return this.isVisible
}

func (this *Vertex) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Vertex) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Vertex) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Vertex) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Vertex) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Vertex) Color24Bit() int {
	return this.color24Bit
}

func (this *Vertex) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Vertex) ColorName() string {
	return this.colorName
}

func (this *Vertex) SetColorName(val string) {
	this.colorName = val
}

func (this *Vertex) Transparency() int {
	return this.transparency
}

func (this *Vertex) SetTransparency(val int) {
	this.transparency = val
}

func (this *Vertex) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Vertex) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

// IsExtraCreatedByCurveFit status flag.
func (this *Vertex) IsExtraCreatedByCurveFit() bool {
	return this.Flags & 1 != 0
}

// IsExtraCreatedByCurveFit status flag.
func (this *Vertex) SetIsExtraCreatedByCurveFit(val bool) {
	if val {
		this.Flags = this.Flags | 1
	} else {
		this.Flags = this.Flags & ^1
	}
}

// IsCurveFitTangentDefined status flag.
func (this *Vertex) IsCurveFitTangentDefined() bool {
	return this.Flags & 2 != 0
}

// IsCurveFitTangentDefined status flag.
func (this *Vertex) SetIsCurveFitTangentDefined(val bool) {
	if val {
		this.Flags = this.Flags | 2
	} else {
		this.Flags = this.Flags & ^2
	}
}

// IsSplineVertexCreatedBySplineFitting status flag.
func (this *Vertex) IsSplineVertexCreatedBySplineFitting() bool {
	return this.Flags & 8 != 0
}

// IsSplineVertexCreatedBySplineFitting status flag.
func (this *Vertex) SetIsSplineVertexCreatedBySplineFitting(val bool) {
	if val {
		this.Flags = this.Flags | 8
	} else {
		this.Flags = this.Flags & ^8
	}
}

// IsSplineFrameControlPoint status flag.
func (this *Vertex) IsSplineFrameControlPoint() bool {
	return this.Flags & 16 != 0
}

// IsSplineFrameControlPoint status flag.
func (this *Vertex) SetIsSplineFrameControlPoint(val bool) {
	if val {
		this.Flags = this.Flags | 16
	} else {
		this.Flags = this.Flags & ^16
	}
}

// Is3DPolylineVertex status flag.
func (this *Vertex) Is3DPolylineVertex() bool {
	return this.Flags & 32 != 0
}

// Is3DPolylineVertex status flag.
func (this *Vertex) SetIs3DPolylineVertex(val bool) {
	if val {
		this.Flags = this.Flags | 32
	} else {
		this.Flags = this.Flags & ^32
	}
}

// Is3DPolygonMesh status flag.
func (this *Vertex) Is3DPolygonMesh() bool {
	return this.Flags & 64 != 0
}

// Is3DPolygonMesh status flag.
func (this *Vertex) SetIs3DPolygonMesh(val bool) {
	if val {
		this.Flags = this.Flags | 64
	} else {
		this.Flags = this.Flags & ^64
	}
}

// IsPolyfaceMeshVertex status flag.
func (this *Vertex) IsPolyfaceMeshVertex() bool {
	return this.Flags & 128 != 0
}

// IsPolyfaceMeshVertex status flag.
func (this *Vertex) SetIsPolyfaceMeshVertex(val bool) {
	if val {
		this.Flags = this.Flags | 128
	} else {
		this.Flags = this.Flags & ^128
	}
}

func (this *Vertex) typeString() string {
	return "VERTEX"
}

func (this *Vertex) minVersion() (version AcadVersion) {
	return Version1_0
}

func (this *Vertex) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Vertex) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.Location.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Location.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Location.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.StartingWidth = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.EndingWidth = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.Bulge = codePair.Value.(DoubleCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 50:
		this.CurveFitTangentDirection = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.PolyfaceMeshVertexIndex1 = int(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.PolyfaceMeshVertexIndex2 = int(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.PolyfaceMeshVertexIndex3 = int(codePair.Value.(ShortCodePairValue).Value)
	case 74:
		this.PolyfaceMeshVertexIndex4 = int(codePair.Value.(ShortCodePairValue).Value)
	case 91:
		this.Identifier = codePair.Value.(IntCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

type Wipeout struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	// fields for RasterImage interface
	_subclassMarker string
	classVersion int
	location Point // The location of the bottom-left corner of the image
	uVector Vector
	vVector Vector
	imageSize Vector // Image size in pixels
	_imageDefinitionHandle string
	displayOptionsFlags int
	useClipping bool
	brightness int16
	contrast int16
	fade int16
	_imageDefinitionReactorHandle string
	clippingType ImageClippingBoundaryType
	clippingVertices []Point
	_clippingVertexCount int
	_clippingVerticesX []float64
	_clippingVerticesY []float64
	isInsideClipping bool
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewWipeout() *Wipeout {
	return &Wipeout{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		_subclassMarker: "",
		classVersion: 0,
		location: *NewOrigin(),
		uVector: *NewXAxis(),
		vVector: *NewYAxis(),
		imageSize: *NewZeroVector(),
		_imageDefinitionHandle: "",
		displayOptionsFlags: 1,
		useClipping: true,
		brightness: 50,
		contrast: 50,
		fade: 0,
		_imageDefinitionReactorHandle: "",
		clippingType: ImageClippingBoundaryTypeRectangular,
		clippingVertices: []Point{},
		_clippingVertexCount: 0,
		_clippingVerticesX: []float64{},
		_clippingVerticesY: []float64{},
		isInsideClipping: false,
	}
}

func (e *Wipeout) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *Wipeout) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *Wipeout) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *Wipeout) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *Wipeout) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *Wipeout) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *Wipeout) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *Wipeout) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *Wipeout) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *Wipeout) Handle() Handle {
	return this.handle
}

func (this *Wipeout) SetHandle(val Handle) {
	this.handle = val
}

func (this *Wipeout) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *Wipeout) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *Wipeout) Layer() string {
	return this.layer
}

func (this *Wipeout) SetLayer(val string) {
	this.layer = val
}

func (this *Wipeout) LineTypeName() string {
	return this.lineTypeName
}

func (this *Wipeout) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *Wipeout) Elevation() float64 {
	return this.elevation
}

func (this *Wipeout) SetElevation(val float64) {
	this.elevation = val
}

func (this *Wipeout) MaterialHandle() string {
	return this.materialHandle
}

func (this *Wipeout) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *Wipeout) Color() Color {
	return this.color
}

func (this *Wipeout) SetColor(val Color) {
	this.color = val
}

func (this *Wipeout) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *Wipeout) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *Wipeout) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *Wipeout) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *Wipeout) IsVisible() bool {
	return this.isVisible
}

func (this *Wipeout) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *Wipeout) ImageByteCount() int {
	return this.imageByteCount
}

func (this *Wipeout) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *Wipeout) PreviewImageData() []string {
	return this.previewImageData
}

func (this *Wipeout) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *Wipeout) Color24Bit() int {
	return this.color24Bit
}

func (this *Wipeout) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *Wipeout) ColorName() string {
	return this.colorName
}

func (this *Wipeout) SetColorName(val string) {
	this.colorName = val
}

func (this *Wipeout) Transparency() int {
	return this.transparency
}

func (this *Wipeout) SetTransparency(val int) {
	this.transparency = val
}

func (this *Wipeout) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *Wipeout) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *Wipeout) subclassMarker() string {
	return this._subclassMarker
}

func (this *Wipeout) SetsubclassMarker(val string) {
	this._subclassMarker = val
}

func (this *Wipeout) ClassVersion() int {
	return this.classVersion
}

func (this *Wipeout) SetClassVersion(val int) {
	this.classVersion = val
}

func (this *Wipeout) Location() Point {
	return this.location
}

func (this *Wipeout) SetLocation(val Point) {
	this.location = val
}

func (this *Wipeout) UVector() Vector {
	return this.uVector
}

func (this *Wipeout) SetUVector(val Vector) {
	this.uVector = val
}

func (this *Wipeout) VVector() Vector {
	return this.vVector
}

func (this *Wipeout) SetVVector(val Vector) {
	this.vVector = val
}

func (this *Wipeout) ImageSize() Vector {
	return this.imageSize
}

func (this *Wipeout) SetImageSize(val Vector) {
	this.imageSize = val
}

func (this *Wipeout) imageDefinitionHandle() string {
	return this._imageDefinitionHandle
}

func (this *Wipeout) SetimageDefinitionHandle(val string) {
	this._imageDefinitionHandle = val
}

func (this *Wipeout) DisplayOptionsFlags() int {
	return this.displayOptionsFlags
}

func (this *Wipeout) SetDisplayOptionsFlags(val int) {
	this.displayOptionsFlags = val
}

func (this *Wipeout) UseClipping() bool {
	return this.useClipping
}

func (this *Wipeout) SetUseClipping(val bool) {
	this.useClipping = val
}

func (this *Wipeout) Brightness() int16 {
	return this.brightness
}

func (this *Wipeout) SetBrightness(val int16) {
	this.brightness = val
}

func (this *Wipeout) Contrast() int16 {
	return this.contrast
}

func (this *Wipeout) SetContrast(val int16) {
	this.contrast = val
}

func (this *Wipeout) Fade() int16 {
	return this.fade
}

func (this *Wipeout) SetFade(val int16) {
	this.fade = val
}

func (this *Wipeout) imageDefinitionReactorHandle() string {
	return this._imageDefinitionReactorHandle
}

func (this *Wipeout) SetimageDefinitionReactorHandle(val string) {
	this._imageDefinitionReactorHandle = val
}

func (this *Wipeout) ClippingType() ImageClippingBoundaryType {
	return this.clippingType
}

func (this *Wipeout) SetClippingType(val ImageClippingBoundaryType) {
	this.clippingType = val
}

func (this *Wipeout) ClippingVertices() []Point {
	return this.clippingVertices
}

func (this *Wipeout) SetClippingVertices(val []Point) {
	this.clippingVertices = val
}

func (this *Wipeout) clippingVertexCount() int {
	return this._clippingVertexCount
}

func (this *Wipeout) SetclippingVertexCount(val int) {
	this._clippingVertexCount = val
}

func (this *Wipeout) clippingVerticesX() []float64 {
	return this._clippingVerticesX
}

func (this *Wipeout) SetclippingVerticesX(val []float64) {
	this._clippingVerticesX = val
}

func (this *Wipeout) clippingVerticesY() []float64 {
	return this._clippingVerticesY
}

func (this *Wipeout) SetclippingVerticesY(val []float64) {
	this._clippingVerticesY = val
}

func (this *Wipeout) IsInsideClipping() bool {
	return this.isInsideClipping
}

func (this *Wipeout) SetIsInsideClipping(val bool) {
	this.isInsideClipping = val
}

func (this *Wipeout) typeString() string {
	return "WIPEOUT"
}

func (this *Wipeout) minVersion() (version AcadVersion) {
	return R2000
}

func (this *Wipeout) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *Wipeout) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForRasterImage(this, codePair)
		}
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *Wipeout) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "WIPEOUT"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	pairs = append(pairs, codePairsForRasterImage(this, version)...)
	return
}

type XLine struct {
	// fields for Entity interface
	handle Handle
	isInPaperSpace bool
	layer string
	lineTypeName string
	elevation float64
	materialHandle string
	color Color
	lineWeight LineWeight
	lineTypeScale float64
	isVisible bool
	imageByteCount int
	previewImageData []string
	color24Bit int
	colorName string
	transparency int
	shadowMode ShadowMode
	FirstPoint Point
	UnitDirectionVector Vector
	pointerOwner pointer
	pointerPlotStyle pointer
}

func NewXLine() *XLine {
	return &XLine{
		handle: 0,
		isInPaperSpace: false,
		layer: "0",
		lineTypeName: "BYLAYER",
		elevation: 0.0,
		materialHandle: "BYLAYER",
		color: ByLayer(),
		lineWeight: NewLineWeightStandard(),
		lineTypeScale: 1.0,
		isVisible: true,
		imageByteCount: 0,
		previewImageData: []string{},
		color24Bit: 0,
		colorName: "",
		transparency: 0,
		shadowMode: ShadowModeCastsAndReceivesShadows,
		FirstPoint: *NewOrigin(),
		UnitDirectionVector: *NewXAxis(),
	}
}

func (e *XLine) pointers() (pointers []*pointer) {
	pointers = append(pointers, &e.pointerOwner)
	pointers = append(pointers, &e.pointerPlotStyle)
	return
}

func (e *XLine) getOwnerPointer() pointer {
	return e.pointerOwner
}

func (e *XLine) setOwnerPointerHandle(h Handle) {
	e.pointerOwner.handle = h
}

func (e *XLine) Owner() *DrawingItem {
	return e.pointerOwner.value
}

func (e *XLine) SetOwner(val *DrawingItem) {
	e.pointerOwner.value = val
}

func (e *XLine) getPlotStylePointer() pointer {
	return e.pointerPlotStyle
}

func (e *XLine) setPlotStylePointerHandle(h Handle) {
	e.pointerPlotStyle.handle = h
}

func (e *XLine) PlotStyle() *DrawingItem {
	return e.pointerPlotStyle.value
}

func (e *XLine) SetPlotStyle(val *DrawingItem) {
	e.pointerPlotStyle.value = val
}

func (this *XLine) Handle() Handle {
	return this.handle
}

func (this *XLine) SetHandle(val Handle) {
	this.handle = val
}

func (this *XLine) IsInPaperSpace() bool {
	return this.isInPaperSpace
}

func (this *XLine) SetIsInPaperSpace(val bool) {
	this.isInPaperSpace = val
}

func (this *XLine) Layer() string {
	return this.layer
}

func (this *XLine) SetLayer(val string) {
	this.layer = val
}

func (this *XLine) LineTypeName() string {
	return this.lineTypeName
}

func (this *XLine) SetLineTypeName(val string) {
	this.lineTypeName = val
}

func (this *XLine) Elevation() float64 {
	return this.elevation
}

func (this *XLine) SetElevation(val float64) {
	this.elevation = val
}

func (this *XLine) MaterialHandle() string {
	return this.materialHandle
}

func (this *XLine) SetMaterialHandle(val string) {
	this.materialHandle = val
}

func (this *XLine) Color() Color {
	return this.color
}

func (this *XLine) SetColor(val Color) {
	this.color = val
}

func (this *XLine) LineWeight() LineWeight {
	return this.lineWeight
}

func (this *XLine) SetLineWeight(val LineWeight) {
	this.lineWeight = val
}

func (this *XLine) LineTypeScale() float64 {
	return this.lineTypeScale
}

func (this *XLine) SetLineTypeScale(val float64) {
	this.lineTypeScale = val
}

func (this *XLine) IsVisible() bool {
	return this.isVisible
}

func (this *XLine) SetIsVisible(val bool) {
	this.isVisible = val
}

func (this *XLine) ImageByteCount() int {
	return this.imageByteCount
}

func (this *XLine) SetImageByteCount(val int) {
	this.imageByteCount = val
}

func (this *XLine) PreviewImageData() []string {
	return this.previewImageData
}

func (this *XLine) SetPreviewImageData(val []string) {
	this.previewImageData = val
}

func (this *XLine) Color24Bit() int {
	return this.color24Bit
}

func (this *XLine) SetColor24Bit(val int) {
	this.color24Bit = val
}

func (this *XLine) ColorName() string {
	return this.colorName
}

func (this *XLine) SetColorName(val string) {
	this.colorName = val
}

func (this *XLine) Transparency() int {
	return this.transparency
}

func (this *XLine) SetTransparency(val int) {
	this.transparency = val
}

func (this *XLine) ShadowMode() ShadowMode {
	return this.shadowMode
}

func (this *XLine) SetShadowMode(val ShadowMode) {
	this.shadowMode = val
}

func (this *XLine) typeString() string {
	return "XLINE"
}

func (this *XLine) minVersion() (version AcadVersion) {
	return R13
}

func (this *XLine) maxVersion() (version AcadVersion) {
	return R2018
}

func (this *XLine) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 10:
		this.FirstPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.FirstPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.FirstPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.UnitDirectionVector.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.UnitDirectionVector.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.UnitDirectionVector.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
		appliedCodePair := false
		if !appliedCodePair {
			appliedCodePair = tryApplyCodePairForEntity(this, codePair)
		}
	}
}

func (this *XLine) codePairs(version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "XLINE"))
	pairs = append(pairs, codePairsForEntity(this, version)...)
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbXline"))
	}
	pairs = append(pairs, NewDoubleCodePair(10, this.FirstPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.FirstPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.FirstPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.UnitDirectionVector.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.UnitDirectionVector.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.UnitDirectionVector.Z))
	return
}

func createAndPopulateDimension(temp *dimensionHelper) (dimension Entity, error error) {
	switch temp.DimensionType() {
	case DimensionTypeAligned:
		dimension = NewAlignedDimension()
	case DimensionTypeRotatedHorizontalOrVertical:
		dimension = NewRotatedDimension()
	case DimensionTypeRadius:
		dimension = NewRadialDimension()
	case DimensionTypeDiameter:
		dimension = NewDiameterDimension()
	case DimensionTypeAngularThreePoint:
		dimension = NewAngularThreePointDimension()
	case DimensionTypeOrdinate:
		dimension = NewOrdinateDimension()
	default:
		error = errors.New(fmt.Sprintf("Unsupported dimension type %s", temp.DimensionType()))
		return
	}

	for _, pair := range temp.collectedPairs {
		dimension.tryApplyCodePair(pair)
	}

	return
}

func createEntity(entityType string) (entity Entity, ok bool) {
	ok = true
	switch entityType {
	case "3DFACE":
		entity = NewFace()
	case "3DSOLID":
		entity = NewSolid3D()
	case "ACAD_PROXY_ENTITY":
		entity = NewProxyEntity()
	case "ARC":
		entity = NewArc()
	case "ARCALIGNEDTEXT":
		entity = NewArcAlignedText()
	case "ATTDEF":
		entity = NewAttributeDefinition()
	case "ATTRIB":
		entity = NewAttribute()
	case "BODY":
		entity = NewBody()
	case "CIRCLE":
		entity = NewCircle()
	case "DIMENSION":
		entity = NewdimensionHelper()
	case "ELLIPSE":
		entity = NewEllipse()
	case "HELIX":
		entity = NewHelix()
	case "IMAGE":
		entity = NewImage()
	case "INSERT":
		entity = NewInsert()
	case "LEADER":
		entity = NewLeader()
	case "LIGHT":
		entity = NewLight()
	case "LINE", "3DLINE":
		entity = NewLine()
	case "LWPOLYLINE":
		entity = NewLWPolyline()
	case "MLINE":
		entity = NewMLine()
	case "MTEXT":
		entity = NewMText()
	case "OLEFRAME":
		entity = NewOleFrame()
	case "OLE2FRAME":
		entity = NewOle2Frame()
	case "POINT":
		entity = NewModelPoint()
	case "POLYLINE":
		entity = NewPolyline()
	case "RAY":
		entity = NewRay()
	case "REGION":
		entity = NewRegion()
	case "RTEXT":
		entity = NewRText()
	case "SECTION":
		entity = NewSection()
	case "SEQEND":
		entity = NewSeqend()
	case "SHAPE":
		entity = NewShape()
	case "SOLID":
		entity = NewSolid()
	case "SPLINE":
		entity = NewSpline()
	case "TEXT":
		entity = NewText()
	case "TOLERANCE":
		entity = NewTolerance()
	case "TRACE":
		entity = NewTrace()
	case "DGNUNDERLAY":
		entity = NewDgnUnderlay()
	case "DWFUNDERLAY":
		entity = NewDwfUnderlay()
	case "PDFUNDERLAY":
		entity = NewPdfUnderlay()
	case "VERTEX":
		entity = NewVertex()
	case "WIPEOUT":
		entity = NewWipeout()
	case "XLINE":
		entity = NewXLine()
	default:
		ok = false
	}
	return
}
