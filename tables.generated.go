// Code generated at build time; DO NOT EDIT.

package dxf

import (
	"fmt"
)

type ViewPort struct {
	handle Handle
	Name string
	Flags int
	LowerLeft Point
	UpperRight Point
	ViewCenter Point
	SnapBasePoint Point
	SnapSpacing Vector
	GridSpacing Vector
	ViewDirection Vector
	TargetViewPoint Point
	ViewHeight float64
	ViewPortAspectRatio float64
	LensLength float64
	FrontClippingPlane float64
	BackClippingPlane float64
	SnapRotationAngle float64
	ViewTwistAngle float64
	ViewMode ViewMode
	CircleSides int
	FastZoom bool
	UCSIcon int16
	SnapOn bool
	GridOn bool
	SnapStyle SnapStyle
	SnapIsometricPlane SnapIsometricPlane
	PlotStyleSheet string
	RenderMode ViewRenderMode
	HasOwnUCS bool
	UCSOrigin Point
	UCSXAxis Vector
	UCSYAxis Vector
	OrthographicViewType OrthographicViewType
	UCSElevation float64
	UCSHandle string
	BaseUCSHandle string
	ShadePlotSetting ShadeEdgeMode
	MajorGridLines bool
	BackgroundObjectHandle string
	ShadePlotObjectHandle string
	VisualStyleObjectHandle string
	IsDefaultLightingOn bool
	DefaultLightingType DefaultLightingType
	Brightness float64
	Contrast float64
	AmbientColor Color
	AmbientColorInt int
	AmbientColorName string
}

func NewViewPort() *ViewPort {
	return &ViewPort{
		Name: "",
		Flags: 0,
		LowerLeft: *NewOrigin(),
		UpperRight: Point{1.0, 1.0, 0.0},
		ViewCenter: *NewOrigin(),
		SnapBasePoint: *NewOrigin(),
		SnapSpacing: Vector{1.0, 1.0, 0.0},
		GridSpacing: Vector{1.0, 1.0, 0.0},
		ViewDirection: *NewZAxis(),
		TargetViewPoint: *NewOrigin(),
		ViewHeight: 1.0,
		ViewPortAspectRatio: 1.0,
		LensLength: 50.0,
		FrontClippingPlane: 0.0,
		BackClippingPlane: 0.0,
		SnapRotationAngle: 0.0,
		ViewTwistAngle: 0.0,
		ViewMode: ViewMode(0),
		CircleSides: 1000,
		FastZoom: true,
		UCSIcon: 3,
		SnapOn: false,
		GridOn: false,
		SnapStyle: SnapStyleStandard,
		SnapIsometricPlane: SnapIsometricPlaneLeft,
		PlotStyleSheet: "",
		RenderMode: ViewRenderModeClassic2D,
		HasOwnUCS: false,
		UCSOrigin: *NewOrigin(),
		UCSXAxis: *NewXAxis(),
		UCSYAxis: *NewYAxis(),
		OrthographicViewType: OrthographicViewTypeNone,
		UCSElevation: 0.0,
		UCSHandle: "",
		BaseUCSHandle: "",
		ShadePlotSetting: ShadeEdgeModeFacesShadedEdgeNotHighlighted,
		MajorGridLines: false,
		BackgroundObjectHandle: "",
		ShadePlotObjectHandle: "",
		VisualStyleObjectHandle: "",
		IsDefaultLightingOn: true,
		DefaultLightingType: DefaultLightingTypeOneDistantLight,
		Brightness: 0.0,
		Contrast: 0.0,
		AmbientColor: Color(7),
		AmbientColorInt: 0,
		AmbientColorName: "BLACK",
	}
}

func (this *ViewPort) Handle() Handle {
	return this.handle
}

func (this *ViewPort) SetHandle(val Handle) {
	this.handle = val
}

func readViewPorts(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "VPORT" {
			error = fmt.Errorf("expected [0/VPORT] but found [%s]", nextPair.String())
			return
		}
		item := *NewViewPort()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.ViewPorts = append(drawing.ViewPorts, item)
	}
	return
}

func (this *ViewPort) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 10:
		this.LowerLeft.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.LowerLeft.Y = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.UpperRight.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.UpperRight.Y = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.ViewCenter.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.ViewCenter.Y = codePair.Value.(DoubleCodePairValue).Value
	case 13:
		this.SnapBasePoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.SnapBasePoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 14:
		this.SnapSpacing.X = codePair.Value.(DoubleCodePairValue).Value
	case 24:
		this.SnapSpacing.Y = codePair.Value.(DoubleCodePairValue).Value
	case 15:
		this.GridSpacing.X = codePair.Value.(DoubleCodePairValue).Value
	case 25:
		this.GridSpacing.Y = codePair.Value.(DoubleCodePairValue).Value
	case 16:
		this.ViewDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 26:
		this.ViewDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 36:
		this.ViewDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 17:
		this.TargetViewPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 27:
		this.TargetViewPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 37:
		this.TargetViewPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 40:
		this.ViewHeight = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.ViewPortAspectRatio = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.LensLength = codePair.Value.(DoubleCodePairValue).Value
	case 43:
		this.FrontClippingPlane = codePair.Value.(DoubleCodePairValue).Value
	case 44:
		this.BackClippingPlane = codePair.Value.(DoubleCodePairValue).Value
	case 45:
		this.ViewHeight = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.SnapRotationAngle = codePair.Value.(DoubleCodePairValue).Value
	case 51:
		this.ViewTwistAngle = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.ViewMode = ViewMode(int(codePair.Value.(ShortCodePairValue).Value))
	case 72:
		this.CircleSides = int(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.FastZoom = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 74:
		this.UCSIcon = codePair.Value.(ShortCodePairValue).Value
	case 75:
		this.SnapOn = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 76:
		this.GridOn = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 77:
		this.SnapStyle = SnapStyle(codePair.Value.(ShortCodePairValue).Value)
	case 78:
		this.SnapIsometricPlane = SnapIsometricPlane(codePair.Value.(ShortCodePairValue).Value)
	case 1:
		this.PlotStyleSheet = codePair.Value.(StringCodePairValue).Value
	case 281:
		this.RenderMode = ViewRenderMode(codePair.Value.(ShortCodePairValue).Value)
	case 65:
		this.HasOwnUCS = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 110:
		this.UCSOrigin.X = codePair.Value.(DoubleCodePairValue).Value
	case 120:
		this.UCSOrigin.Y = codePair.Value.(DoubleCodePairValue).Value
	case 130:
		this.UCSOrigin.Z = codePair.Value.(DoubleCodePairValue).Value
	case 111:
		this.UCSXAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 121:
		this.UCSXAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 131:
		this.UCSXAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 112:
		this.UCSYAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 122:
		this.UCSYAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 132:
		this.UCSYAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 79:
		this.OrthographicViewType = OrthographicViewType(codePair.Value.(ShortCodePairValue).Value)
	case 146:
		this.UCSElevation = codePair.Value.(DoubleCodePairValue).Value
	case 345:
		this.UCSHandle = codePair.Value.(StringCodePairValue).Value
	case 346:
		this.BaseUCSHandle = codePair.Value.(StringCodePairValue).Value
	case 170:
		this.ShadePlotSetting = ShadeEdgeMode(codePair.Value.(ShortCodePairValue).Value)
	case 61:
		this.MajorGridLines = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 332:
		this.BackgroundObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 333:
		this.ShadePlotObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 348:
		this.VisualStyleObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 292:
		this.IsDefaultLightingOn = codePair.Value.(BoolCodePairValue).Value
	case 282:
		this.DefaultLightingType = DefaultLightingType(codePair.Value.(ShortCodePairValue).Value)
	case 141:
		this.Brightness = codePair.Value.(DoubleCodePairValue).Value
	case 142:
		this.Contrast = codePair.Value.(DoubleCodePairValue).Value
	case 62:
		this.AmbientColor = Color(codePair.Value.(ShortCodePairValue).Value)
	case 421:
		this.AmbientColorInt = codePair.Value.(IntCodePairValue).Value
	case 431:
		this.AmbientColorName = codePair.Value.(StringCodePairValue).Value
	default:
	}
}

func tablePairsViewPorts(tableHandle Handle, items []ViewPort, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "VPORT"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "VPORT"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *ViewPort) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbViewportTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewDoubleCodePair(10, this.LowerLeft.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.LowerLeft.Y))
	pairs = append(pairs, NewDoubleCodePair(11, this.UpperRight.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.UpperRight.Y))
	pairs = append(pairs, NewDoubleCodePair(12, this.ViewCenter.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.ViewCenter.Y))
	pairs = append(pairs, NewDoubleCodePair(13, this.SnapBasePoint.X))
	pairs = append(pairs, NewDoubleCodePair(23, this.SnapBasePoint.Y))
	pairs = append(pairs, NewDoubleCodePair(14, this.SnapSpacing.X))
	pairs = append(pairs, NewDoubleCodePair(24, this.SnapSpacing.Y))
	pairs = append(pairs, NewDoubleCodePair(15, this.GridSpacing.X))
	pairs = append(pairs, NewDoubleCodePair(25, this.GridSpacing.Y))
	pairs = append(pairs, NewDoubleCodePair(16, this.ViewDirection.X))
	pairs = append(pairs, NewDoubleCodePair(26, this.ViewDirection.Y))
	pairs = append(pairs, NewDoubleCodePair(36, this.ViewDirection.Z))
	pairs = append(pairs, NewDoubleCodePair(17, this.TargetViewPoint.X))
	pairs = append(pairs, NewDoubleCodePair(27, this.TargetViewPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(37, this.TargetViewPoint.Z))
	if version <= R2004 {
		pairs = append(pairs, NewDoubleCodePair(40, ensurePositiveOrDefault(this.ViewHeight, 1.0)))
	}
	if version <= R2004 {
		pairs = append(pairs, NewDoubleCodePair(41, ensurePositiveOrDefault(this.ViewPortAspectRatio, 1.0)))
	}
	pairs = append(pairs, NewDoubleCodePair(42, ensurePositiveOrDefault(this.LensLength, 50.0)))
	pairs = append(pairs, NewDoubleCodePair(43, this.FrontClippingPlane))
	pairs = append(pairs, NewDoubleCodePair(44, this.BackClippingPlane))
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(45, ensurePositiveOrDefault(this.ViewHeight, 1.0)))
	}
	pairs = append(pairs, NewDoubleCodePair(50, this.SnapRotationAngle))
	pairs = append(pairs, NewDoubleCodePair(51, this.ViewTwistAngle))
	pairs = append(pairs, NewShortCodePair(71, int16(this.ViewMode)))
	pairs = append(pairs, NewShortCodePair(72, int16(ensurePositiveOrDefault(float64(this.CircleSides), 1000.0))))
	if version <= R2004 {
		pairs = append(pairs, NewShortCodePair(73, shortFromBool(this.FastZoom)))
	}
	pairs = append(pairs, NewShortCodePair(74, int16(ensurePositiveOrDefault(float64(this.UCSIcon), 3.0))))
	if version <= R2004 {
		pairs = append(pairs, NewShortCodePair(75, shortFromBool(this.SnapOn)))
	}
	if version <= R2004 {
		pairs = append(pairs, NewShortCodePair(76, shortFromBool(this.GridOn)))
	}
	if version <= R2004 {
		pairs = append(pairs, NewShortCodePair(77, int16(this.SnapStyle)))
	}
	if version <= R2004 {
		pairs = append(pairs, NewShortCodePair(78, int16(this.SnapIsometricPlane)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(1, this.PlotStyleSheet))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(281, int16(this.RenderMode)))
	}
	if version >= R2000 && version <= R2004 {
		pairs = append(pairs, NewShortCodePair(65, shortFromBool(this.HasOwnUCS)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(110, this.UCSOrigin.X))
		pairs = append(pairs, NewDoubleCodePair(120, this.UCSOrigin.Y))
		pairs = append(pairs, NewDoubleCodePair(130, this.UCSOrigin.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(111, this.UCSXAxis.X))
		pairs = append(pairs, NewDoubleCodePair(121, this.UCSXAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(131, this.UCSXAxis.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(112, this.UCSYAxis.X))
		pairs = append(pairs, NewDoubleCodePair(122, this.UCSYAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(132, this.UCSYAxis.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(79, int16(this.OrthographicViewType)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(146, this.UCSElevation))
	}
	if version >= R2000 && this.UCSHandle != "" {
		pairs = append(pairs, NewStringCodePair(345, this.UCSHandle))
	}
	if version >= R2000 && this.BaseUCSHandle != "" {
		pairs = append(pairs, NewStringCodePair(346, this.BaseUCSHandle))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(170, int16(this.ShadePlotSetting)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(61, shortFromBool(this.MajorGridLines)))
	}
	if version >= R2007 && this.BackgroundObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(332, this.BackgroundObjectHandle))
	}
	if version >= R2007 && this.ShadePlotObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(333, this.ShadePlotObjectHandle))
	}
	if version >= R2007 && this.VisualStyleObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(348, this.VisualStyleObjectHandle))
	}
	if version >= R2007 {
		pairs = append(pairs, NewBoolCodePair(292, this.IsDefaultLightingOn))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(282, int16(this.DefaultLightingType)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(141, this.Brightness))
	}
	if version >= R2007 {
		pairs = append(pairs, NewDoubleCodePair(142, this.Contrast))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(62, int16(this.AmbientColor)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewIntCodePair(421, this.AmbientColorInt))
	}
	if version >= R2007 {
		pairs = append(pairs, NewStringCodePair(431, this.AmbientColorName))
	}
	return
}

type LineType struct {
	handle Handle
	Name string
	Flags int
	Description string
	AlignmentCode int
	ElementCount int
	TotalPatternLength float64
	DashDotSpaceLengths []float64
	ComplexLineTypeElementTypes []int16
	ShapeNumbers []int16
	StyleHandles []string
	ScaleValues []float64
	RotationAngles []float64
	XOffsets []float64
	YOffsets []float64
	TextStrings []string
}

func NewLineType() *LineType {
	return &LineType{
		Name: "",
		Flags: 0,
		Description: "",
		AlignmentCode: int('A'),
		ElementCount: 0,
		TotalPatternLength: 0.0,
		DashDotSpaceLengths: []float64{},
		ComplexLineTypeElementTypes: []int16{},
		ShapeNumbers: []int16{},
		StyleHandles: []string{},
		ScaleValues: []float64{},
		RotationAngles: []float64{},
		XOffsets: []float64{},
		YOffsets: []float64{},
		TextStrings: []string{},
	}
}

func (this *LineType) Handle() Handle {
	return this.handle
}

func (this *LineType) SetHandle(val Handle) {
	this.handle = val
}

func readLineTypes(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "LTYPE" {
			error = fmt.Errorf("expected [0/LTYPE] but found [%s]", nextPair.String())
			return
		}
		item := *NewLineType()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.LineTypes = append(drawing.LineTypes, item)
	}
	return
}

func (this *LineType) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 3:
		this.Description = codePair.Value.(StringCodePairValue).Value
	case 72:
		this.AlignmentCode = int(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.ElementCount = int(codePair.Value.(ShortCodePairValue).Value)
	case 40:
		this.TotalPatternLength = codePair.Value.(DoubleCodePairValue).Value
	case 49:
		this.DashDotSpaceLengths = append(this.DashDotSpaceLengths, codePair.Value.(DoubleCodePairValue).Value)
	case 74:
		this.ComplexLineTypeElementTypes = append(this.ComplexLineTypeElementTypes, codePair.Value.(ShortCodePairValue).Value)
	case 75:
		this.ShapeNumbers = append(this.ShapeNumbers, codePair.Value.(ShortCodePairValue).Value)
	case 340:
		this.StyleHandles = append(this.StyleHandles, codePair.Value.(StringCodePairValue).Value)
	case 46:
		this.ScaleValues = append(this.ScaleValues, codePair.Value.(DoubleCodePairValue).Value)
	case 50:
		this.RotationAngles = append(this.RotationAngles, codePair.Value.(DoubleCodePairValue).Value)
	case 44:
		this.XOffsets = append(this.XOffsets, codePair.Value.(DoubleCodePairValue).Value)
	case 45:
		this.YOffsets = append(this.YOffsets, codePair.Value.(DoubleCodePairValue).Value)
	case 9:
		this.TextStrings = append(this.TextStrings, codePair.Value.(StringCodePairValue).Value)
	default:
	}
}

func tablePairsLineTypes(tableHandle Handle, items []LineType, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "LTYPE"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "LTYPE"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *LineType) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbLinetypeTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewStringCodePair(3, this.Description))
	pairs = append(pairs, NewShortCodePair(72, int16(this.AlignmentCode)))
	pairs = append(pairs, NewShortCodePair(73, int16(this.ElementCount)))
	pairs = append(pairs, NewDoubleCodePair(40, this.TotalPatternLength))
	for _, val := range this.DashDotSpaceLengths {
		pairs = append(pairs, NewDoubleCodePair(49, val))
	}
	if version >= R13 {
		for _, val := range this.ComplexLineTypeElementTypes {
			pairs = append(pairs, NewShortCodePair(74, val))
		}
	}
	if version >= R13 {
		for _, val := range this.ShapeNumbers {
			pairs = append(pairs, NewShortCodePair(75, val))
		}
	}
	if version >= R13 {
		for _, val := range this.StyleHandles {
			pairs = append(pairs, NewStringCodePair(340, val))
		}
	}
	if version >= R13 {
		for _, val := range this.ScaleValues {
			pairs = append(pairs, NewDoubleCodePair(46, val))
		}
	}
	if version >= R13 {
		for _, val := range this.RotationAngles {
			pairs = append(pairs, NewDoubleCodePair(50, val))
		}
	}
	if version >= R13 {
		for _, val := range this.XOffsets {
			pairs = append(pairs, NewDoubleCodePair(44, val))
		}
	}
	if version >= R13 {
		for _, val := range this.YOffsets {
			pairs = append(pairs, NewDoubleCodePair(45, val))
		}
	}
	if version >= R13 {
		for _, val := range this.TextStrings {
			pairs = append(pairs, NewStringCodePair(9, val))
		}
	}
	return
}

type Layer struct {
	handle Handle
	Name string
	Flags int
	Color Color
	LineTypeName string
	IsLayerPlotted bool
	LineWeight LineWeight
	PlotStyleHandle string
	MaterialHandle string
}

func NewLayer() *Layer {
	return &Layer{
		Name: "",
		Flags: 0,
		Color: Color(7),
		LineTypeName: "CONTINUOUS",
		IsLayerPlotted: true,
		LineWeight: NewLineWeightStandard(),
		PlotStyleHandle: "",
		MaterialHandle: "",
	}
}

func (this *Layer) Handle() Handle {
	return this.handle
}

func (this *Layer) SetHandle(val Handle) {
	this.handle = val
}

func readLayers(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "LAYER" {
			error = fmt.Errorf("expected [0/LAYER] but found [%s]", nextPair.String())
			return
		}
		item := *NewLayer()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.Layers = append(drawing.Layers, item)
	}
	return
}

func (this *Layer) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 62:
		this.Color = Color(codePair.Value.(ShortCodePairValue).Value)
	case 6:
		this.LineTypeName = codePair.Value.(StringCodePairValue).Value
	case 290:
		this.IsLayerPlotted = codePair.Value.(BoolCodePairValue).Value
	case 370:
		this.LineWeight = LineWeight(codePair.Value.(ShortCodePairValue).Value)
	case 390:
		this.PlotStyleHandle = codePair.Value.(StringCodePairValue).Value
	case 347:
		this.MaterialHandle = codePair.Value.(StringCodePairValue).Value
	default:
	}
}

func tablePairsLayers(tableHandle Handle, items []Layer, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "LAYER"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "LAYER"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *Layer) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbLayerTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewShortCodePair(62, int16(this.Color)))
	pairs = append(pairs, NewStringCodePair(6, this.LineTypeName))
	if version >= R2000 {
		pairs = append(pairs, NewBoolCodePair(290, this.IsLayerPlotted))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(370, int16(this.LineWeight)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(390, this.PlotStyleHandle))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(347, this.MaterialHandle))
	}
	return
}

type Style struct {
	handle Handle
	Name string
	Flags int
	TextHeight float64
	WidthFactor float64
	ObliqueAngle float64
	TextGenerationFlags int
	LastHeightUsed float64
	PrimaryFontFileName string
	BigFontFileName string
	FontFlags int
}

func NewStyle() *Style {
	return &Style{
		Name: "",
		Flags: 0,
		TextHeight: 0.0,
		WidthFactor: 1.0,
		ObliqueAngle: 0.0,
		TextGenerationFlags: 0,
		LastHeightUsed: 0.2,
		PrimaryFontFileName: "txt",
		BigFontFileName: "",
		FontFlags: 0,
	}
}

func (this *Style) Handle() Handle {
	return this.handle
}

func (this *Style) SetHandle(val Handle) {
	this.handle = val
}

func readStyles(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "STYLE" {
			error = fmt.Errorf("expected [0/STYLE] but found [%s]", nextPair.String())
			return
		}
		item := *NewStyle()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.Styles = append(drawing.Styles, item)
	}
	return
}

func (this *Style) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 40:
		this.TextHeight = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.WidthFactor = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.ObliqueAngle = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.TextGenerationFlags = int(codePair.Value.(ShortCodePairValue).Value)
	case 42:
		this.LastHeightUsed = codePair.Value.(DoubleCodePairValue).Value
	case 3:
		this.PrimaryFontFileName = codePair.Value.(StringCodePairValue).Value
	case 4:
		this.BigFontFileName = codePair.Value.(StringCodePairValue).Value
	case 1071:
		this.FontFlags = codePair.Value.(IntCodePairValue).Value
	default:
	}
}

func tablePairsStyles(tableHandle Handle, items []Style, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "STYLE"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "STYLE"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *Style) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbTextStyleTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewDoubleCodePair(40, this.TextHeight))
	pairs = append(pairs, NewDoubleCodePair(41, this.WidthFactor))
	pairs = append(pairs, NewDoubleCodePair(50, this.ObliqueAngle))
	pairs = append(pairs, NewShortCodePair(71, int16(this.TextGenerationFlags)))
	pairs = append(pairs, NewDoubleCodePair(42, this.LastHeightUsed))
	pairs = append(pairs, NewStringCodePair(3, this.PrimaryFontFileName))
	pairs = append(pairs, NewStringCodePair(4, this.BigFontFileName))
	if version >= R2010 {
		pairs = append(pairs, NewIntCodePair(1071, this.FontFlags))
	}
	return
}

type View struct {
	handle Handle
	Name string
	Flags int
	ViewHeight float64
	ViewCenterPoint Point
	ViewWidth float64
	ViewDirection Vector
	TargetPoint Point
	LensLength float64
	FrontClippingPlane float64
	BackClippingPlane float64
	TwistAngle float64
	ViewMode int16
	RenderMode ViewRenderMode
	IsAssociatedUCSPresent bool
	IsCameraPlottable bool
	BackgroundObjectHandle string
	SelectionObjectHandle string
	VisualStyleObjectHandle string
	SunOwnershipHandle string
	UCSOrigin Point
	UCSXAxis Vector
	UCSYAxis Vector
	OrthographicViewType OrthographicViewType
	UCSElevation float64
	UCSHandle string
	BaseUCSHandle string
}

func NewView() *View {
	return &View{
		Name: "",
		Flags: 0,
		ViewHeight: 1.0,
		ViewCenterPoint: *NewOrigin(),
		ViewWidth: 1.0,
		ViewDirection: *NewZAxis(),
		TargetPoint: *NewOrigin(),
		LensLength: 1.0,
		FrontClippingPlane: 0.0,
		BackClippingPlane: 1.0,
		TwistAngle: 0.0,
		ViewMode: 0,
		RenderMode: ViewRenderModeClassic2D,
		IsAssociatedUCSPresent: false,
		IsCameraPlottable: false,
		BackgroundObjectHandle: "",
		SelectionObjectHandle: "",
		VisualStyleObjectHandle: "",
		SunOwnershipHandle: "",
		UCSOrigin: *NewOrigin(),
		UCSXAxis: *NewXAxis(),
		UCSYAxis: *NewYAxis(),
		OrthographicViewType: OrthographicViewTypeNone,
		UCSElevation: 0.0,
		UCSHandle: "",
		BaseUCSHandle: "",
	}
}

func (this *View) Handle() Handle {
	return this.handle
}

func (this *View) SetHandle(val Handle) {
	this.handle = val
}

func readViews(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "VIEW" {
			error = fmt.Errorf("expected [0/VIEW] but found [%s]", nextPair.String())
			return
		}
		item := *NewView()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.Views = append(drawing.Views, item)
	}
	return
}

func (this *View) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 40:
		this.ViewHeight = codePair.Value.(DoubleCodePairValue).Value
	case 10:
		this.ViewCenterPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.ViewCenterPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.ViewWidth = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.ViewDirection.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.ViewDirection.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.ViewDirection.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.TargetPoint.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.TargetPoint.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.TargetPoint.Z = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.LensLength = codePair.Value.(DoubleCodePairValue).Value
	case 43:
		this.FrontClippingPlane = codePair.Value.(DoubleCodePairValue).Value
	case 44:
		this.BackClippingPlane = codePair.Value.(DoubleCodePairValue).Value
	case 50:
		this.TwistAngle = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.ViewMode = codePair.Value.(ShortCodePairValue).Value
	case 281:
		this.RenderMode = ViewRenderMode(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.IsAssociatedUCSPresent = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.IsCameraPlottable = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 332:
		this.BackgroundObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 334:
		this.SelectionObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 348:
		this.VisualStyleObjectHandle = codePair.Value.(StringCodePairValue).Value
	case 361:
		this.SunOwnershipHandle = codePair.Value.(StringCodePairValue).Value
	case 110:
		this.UCSOrigin.X = codePair.Value.(DoubleCodePairValue).Value
	case 120:
		this.UCSOrigin.Y = codePair.Value.(DoubleCodePairValue).Value
	case 130:
		this.UCSOrigin.Z = codePair.Value.(DoubleCodePairValue).Value
	case 111:
		this.UCSXAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 121:
		this.UCSXAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 131:
		this.UCSXAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 112:
		this.UCSYAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 122:
		this.UCSYAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 132:
		this.UCSYAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 79:
		this.OrthographicViewType = OrthographicViewType(codePair.Value.(ShortCodePairValue).Value)
	case 146:
		this.UCSElevation = codePair.Value.(DoubleCodePairValue).Value
	case 345:
		this.UCSHandle = codePair.Value.(StringCodePairValue).Value
	case 346:
		this.BaseUCSHandle = codePair.Value.(StringCodePairValue).Value
	default:
	}
}

func tablePairsViews(tableHandle Handle, items []View, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "VIEW"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "VIEW"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *View) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbViewTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewDoubleCodePair(40, ensurePositiveOrDefault(this.ViewHeight, 1.0)))
	pairs = append(pairs, NewDoubleCodePair(10, this.ViewCenterPoint.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.ViewCenterPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(41, ensurePositiveOrDefault(this.ViewWidth, 1.0)))
	pairs = append(pairs, NewDoubleCodePair(11, this.ViewDirection.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.ViewDirection.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.ViewDirection.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.TargetPoint.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.TargetPoint.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.TargetPoint.Z))
	pairs = append(pairs, NewDoubleCodePair(42, ensurePositiveOrDefault(this.LensLength, 1.0)))
	pairs = append(pairs, NewDoubleCodePair(43, this.FrontClippingPlane))
	pairs = append(pairs, NewDoubleCodePair(44, this.BackClippingPlane))
	pairs = append(pairs, NewDoubleCodePair(50, this.TwistAngle))
	pairs = append(pairs, NewShortCodePair(71, this.ViewMode))
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(281, int16(this.RenderMode)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(72, shortFromBool(this.IsAssociatedUCSPresent)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(73, shortFromBool(this.IsCameraPlottable)))
	}
	if version >= R2007 && this.BackgroundObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(332, this.BackgroundObjectHandle))
	}
	if version >= R2007 && this.SelectionObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(334, this.SelectionObjectHandle))
	}
	if version >= R2007 && this.VisualStyleObjectHandle != "" {
		pairs = append(pairs, NewStringCodePair(348, this.VisualStyleObjectHandle))
	}
	if version >= R2010 && this.SunOwnershipHandle != "" {
		pairs = append(pairs, NewStringCodePair(361, this.SunOwnershipHandle))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(110, this.UCSOrigin.X))
		pairs = append(pairs, NewDoubleCodePair(120, this.UCSOrigin.Y))
		pairs = append(pairs, NewDoubleCodePair(130, this.UCSOrigin.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(111, this.UCSXAxis.X))
		pairs = append(pairs, NewDoubleCodePair(121, this.UCSXAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(131, this.UCSXAxis.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(112, this.UCSYAxis.X))
		pairs = append(pairs, NewDoubleCodePair(122, this.UCSYAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(132, this.UCSYAxis.Z))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(79, int16(this.OrthographicViewType)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(146, this.UCSElevation))
	}
	if version >= R2000 && this.UCSHandle != "" {
		pairs = append(pairs, NewStringCodePair(345, this.UCSHandle))
	}
	if version >= R2000 && this.BaseUCSHandle != "" {
		pairs = append(pairs, NewStringCodePair(346, this.BaseUCSHandle))
	}
	return
}

type Ucs struct {
	handle Handle
	Name string
	Flags int
	Origin Point
	XAxis Vector
	YAxis Vector
	OrthographicViewType OrthographicViewType
	Elevation float64
	BaseUcsHandle string
	OrthographicType OrthographicViewType
	OrthographicOrigin Point
}

func NewUcs() *Ucs {
	return &Ucs{
		Name: "",
		Flags: 0,
		Origin: *NewOrigin(),
		XAxis: *NewXAxis(),
		YAxis: *NewXAxis(),
		OrthographicViewType: OrthographicViewTypeNone,
		Elevation: 0.0,
		BaseUcsHandle: "",
		OrthographicType: OrthographicViewTypeTop,
		OrthographicOrigin: *NewOrigin(),
	}
}

func (this *Ucs) Handle() Handle {
	return this.handle
}

func (this *Ucs) SetHandle(val Handle) {
	this.handle = val
}

func readUcss(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "UCS" {
			error = fmt.Errorf("expected [0/UCS] but found [%s]", nextPair.String())
			return
		}
		item := *NewUcs()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.Ucss = append(drawing.Ucss, item)
	}
	return
}

func (this *Ucs) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 10:
		this.Origin.X = codePair.Value.(DoubleCodePairValue).Value
	case 20:
		this.Origin.Y = codePair.Value.(DoubleCodePairValue).Value
	case 30:
		this.Origin.Z = codePair.Value.(DoubleCodePairValue).Value
	case 11:
		this.XAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 21:
		this.XAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 31:
		this.XAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 12:
		this.YAxis.X = codePair.Value.(DoubleCodePairValue).Value
	case 22:
		this.YAxis.Y = codePair.Value.(DoubleCodePairValue).Value
	case 32:
		this.YAxis.Z = codePair.Value.(DoubleCodePairValue).Value
	case 79:
		this.OrthographicViewType = OrthographicViewType(codePair.Value.(ShortCodePairValue).Value)
	case 146:
		this.Elevation = codePair.Value.(DoubleCodePairValue).Value
	case 346:
		this.BaseUcsHandle = codePair.Value.(StringCodePairValue).Value
	case 71:
		this.OrthographicType = OrthographicViewType(codePair.Value.(ShortCodePairValue).Value)
	case 13:
		this.OrthographicOrigin.X = codePair.Value.(DoubleCodePairValue).Value
	case 23:
		this.OrthographicOrigin.Y = codePair.Value.(DoubleCodePairValue).Value
	case 33:
		this.OrthographicOrigin.Z = codePair.Value.(DoubleCodePairValue).Value
	default:
	}
}

func tablePairsUcss(tableHandle Handle, items []Ucs, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "UCS"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "UCS"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *Ucs) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbUCSTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewDoubleCodePair(10, this.Origin.X))
	pairs = append(pairs, NewDoubleCodePair(20, this.Origin.Y))
	pairs = append(pairs, NewDoubleCodePair(30, this.Origin.Z))
	pairs = append(pairs, NewDoubleCodePair(11, this.XAxis.X))
	pairs = append(pairs, NewDoubleCodePair(21, this.XAxis.Y))
	pairs = append(pairs, NewDoubleCodePair(31, this.XAxis.Z))
	pairs = append(pairs, NewDoubleCodePair(12, this.YAxis.X))
	pairs = append(pairs, NewDoubleCodePair(22, this.YAxis.Y))
	pairs = append(pairs, NewDoubleCodePair(32, this.YAxis.Z))
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(79, int16(this.OrthographicViewType)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(146, this.Elevation))
	}
	if version >= R2000 && this.BaseUcsHandle != "" {
		pairs = append(pairs, NewStringCodePair(346, this.BaseUcsHandle))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(71, int16(this.OrthographicType)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(13, this.OrthographicOrigin.X))
		pairs = append(pairs, NewDoubleCodePair(23, this.OrthographicOrigin.Y))
		pairs = append(pairs, NewDoubleCodePair(33, this.OrthographicOrigin.Z))
	}
	return
}

type AppId struct {
	handle Handle
	Name string
	Flags int
}

func NewAppId() *AppId {
	return &AppId{
		Name: "",
		Flags: 0,
	}
}

func (this *AppId) Handle() Handle {
	return this.handle
}

func (this *AppId) SetHandle(val Handle) {
	this.handle = val
}

func readAppIds(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "APPID" {
			error = fmt.Errorf("expected [0/APPID] but found [%s]", nextPair.String())
			return
		}
		item := *NewAppId()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.AppIds = append(drawing.AppIds, item)
	}
	return
}

func (this *AppId) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	default:
	}
}

func tablePairsAppIds(tableHandle Handle, items []AppId, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "APPID"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "APPID"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *AppId) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbRegAppTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	return
}

type DimStyle struct {
	handle Handle
	Name string
	Flags int
	DimensioningSuffix string
	AlternateDimensioningSuffix string
	ArrowBlockName string
	FirstArrowBlockName string
	SecondArrowBlockName string
	DimensioningScaleFactor float64
	DimensioningArrowSize float64
	DimensionExtensionLineOffset float64
	DimensionLineIncrement float64
	DimensionExtensionLineIncrement float64
	DimensionDistanceRoundingValue float64
	DimensionLineExtension float64
	DimensionPlusTolerance float64
	DimensionMinusTolerance float64
	GenerateDimensionTolerances bool
	GenerateDimensionLimits bool
	DimensionTextInsideHorizontal bool
	DimensionTextOutsideHorizontal bool
	SuppressFirstDimensionExtensionLine bool
	SuppressSecondDimensionExtensionLine bool
	TextAboveDimensionLine bool
	DimensionUnitZeroSuppression UnitZeroSuppression
	DimensionAngleZeroSuppression UnitZeroSuppression
	DimensioningTextHeight float64
	CenterMarkSize float64
	DimensioningTickSize float64
	AlternateDimensioningScaleFactor float64
	DimensionLinearMeasurementScaleFactor float64
	DimensionVerticalTextPosition float64
	DimensionToleranceDisplaceScaleFactor float64
	DimensionLineGap float64
	AlternateDimensioningUnitRounding float64
	UseAlternateDimensioning bool
	AlternateDimensioningDecimalPlaces int16
	ForceDimensionLineExtensionsOutsideIfTextExists bool
	UseSeparateArrowBlocksForDimensions bool
	ForceDimensionTextInsideExtensions bool
	SuppressOutsideExtensionDimensionLines bool
	DimensionLineColor Color
	DimensionExtensionLineColor Color
	DimensionTextColor Color
	AngularDimensionPrecision int16
	DimensionUnitFormat UnitFormat
	DimensionUnitToleranceDecimalPlaces int16
	DimensionToleraceDecimalPlaces int16
	AlternateDimensioningUnits UnitFormat
	AlternateDimensioningToleranceDecimalPlaces int16
	DimensioningAngleFormat AngleFormat
	DimensionPrecision int16
	DimensionNonAngularUnits NonAngularUnits
	DimensionDecilamSeparatorChar rune
	DimensionTextMovementRule DimensionTextMovementRule
	DimensionTextJustification DimensionTextJustification
	DimensionToleranceVerticalJustification Justification
	DimensionToleranceZeroSuppression UnitZeroSuppression
	AlternateDimensioningZeroSuppression UnitZeroSuppression
	AlternateDimensioningToleranceZeroSuppression UnitZeroSuppression
	DimensionTextAndArrowPlacement DimensionFit
	DimensionCursorControlsTextPosition bool
	DimensionTextStyle string
	DimensionLeaderBlockName string
	DimensionLineWeight LineWeight
	DimensionExtensionLineWeight LineWeight
}

func NewDimStyle() *DimStyle {
	return &DimStyle{
		Name: "",
		Flags: 0,
		DimensioningSuffix: "",
		AlternateDimensioningSuffix: "",
		ArrowBlockName: "",
		FirstArrowBlockName: "",
		SecondArrowBlockName: "",
		DimensioningScaleFactor: 1.0,
		DimensioningArrowSize: 0.18,
		DimensionExtensionLineOffset: 0.0625,
		DimensionLineIncrement: 0.38,
		DimensionExtensionLineIncrement: 0.18,
		DimensionDistanceRoundingValue: 0.0,
		DimensionLineExtension: 0.0,
		DimensionPlusTolerance: 0.0,
		DimensionMinusTolerance: 0.0,
		GenerateDimensionTolerances: false,
		GenerateDimensionLimits: false,
		DimensionTextInsideHorizontal: true,
		DimensionTextOutsideHorizontal: true,
		SuppressFirstDimensionExtensionLine: false,
		SuppressSecondDimensionExtensionLine: false,
		TextAboveDimensionLine: false,
		DimensionUnitZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		DimensionAngleZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		DimensioningTextHeight: 0.18,
		CenterMarkSize: 0.09,
		DimensioningTickSize: 0.0,
		AlternateDimensioningScaleFactor: 25.4,
		DimensionLinearMeasurementScaleFactor: 1.0,
		DimensionVerticalTextPosition: 0.0,
		DimensionToleranceDisplaceScaleFactor: 1.0,
		DimensionLineGap: 0.09,
		AlternateDimensioningUnitRounding: 0.0,
		UseAlternateDimensioning: false,
		AlternateDimensioningDecimalPlaces: 2,
		ForceDimensionLineExtensionsOutsideIfTextExists: false,
		UseSeparateArrowBlocksForDimensions: false,
		ForceDimensionTextInsideExtensions: false,
		SuppressOutsideExtensionDimensionLines: false,
		DimensionLineColor: ByBlock(),
		DimensionExtensionLineColor: ByBlock(),
		DimensionTextColor: ByBlock(),
		AngularDimensionPrecision: 12,
		DimensionUnitFormat: UnitFormatScientific,
		DimensionUnitToleranceDecimalPlaces: 0,
		DimensionToleraceDecimalPlaces: 0,
		AlternateDimensioningUnits: UnitFormatScientific,
		AlternateDimensioningToleranceDecimalPlaces: 0,
		DimensioningAngleFormat: AngleFormatDecimalDegrees,
		DimensionPrecision: 12,
		DimensionNonAngularUnits: NonAngularUnitsScientific,
		DimensionDecilamSeparatorChar: '.',
		DimensionTextMovementRule: DimensionTextMovementRuleMoveLineWithText,
		DimensionTextJustification: DimensionTextJustificationAboveLineCenter,
		DimensionToleranceVerticalJustification: JustificationTop,
		DimensionToleranceZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		AlternateDimensioningZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		AlternateDimensioningToleranceZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		DimensionTextAndArrowPlacement: DimensionFitTextAndArrowsOutsideLines,
		DimensionCursorControlsTextPosition: true,
		DimensionTextStyle: "",
		DimensionLeaderBlockName: "",
		DimensionLineWeight: NewLineWeightStandard(),
		DimensionExtensionLineWeight: NewLineWeightStandard(),
	}
}

func (this *DimStyle) Handle() Handle {
	return this.handle
}

func (this *DimStyle) SetHandle(val Handle) {
	this.handle = val
}

func readDimStyles(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "DIMSTYLE" {
			error = fmt.Errorf("expected [0/DIMSTYLE] but found [%s]", nextPair.String())
			return
		}
		item := *NewDimStyle()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.DimStyles = append(drawing.DimStyles, item)
	}
	return
}

func (this *DimStyle) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.Flags = int(codePair.Value.(ShortCodePairValue).Value)
	case 3:
		this.DimensioningSuffix = codePair.Value.(StringCodePairValue).Value
	case 4:
		this.AlternateDimensioningSuffix = codePair.Value.(StringCodePairValue).Value
	case 5:
		this.ArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 6:
		this.FirstArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 7:
		this.SecondArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 40:
		this.DimensioningScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 41:
		this.DimensioningArrowSize = codePair.Value.(DoubleCodePairValue).Value
	case 42:
		this.DimensionExtensionLineOffset = codePair.Value.(DoubleCodePairValue).Value
	case 43:
		this.DimensionLineIncrement = codePair.Value.(DoubleCodePairValue).Value
	case 44:
		this.DimensionExtensionLineIncrement = codePair.Value.(DoubleCodePairValue).Value
	case 45:
		this.DimensionDistanceRoundingValue = codePair.Value.(DoubleCodePairValue).Value
	case 46:
		this.DimensionLineExtension = codePair.Value.(DoubleCodePairValue).Value
	case 47:
		this.DimensionPlusTolerance = codePair.Value.(DoubleCodePairValue).Value
	case 48:
		this.DimensionMinusTolerance = codePair.Value.(DoubleCodePairValue).Value
	case 71:
		this.GenerateDimensionTolerances = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 72:
		this.GenerateDimensionLimits = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 73:
		this.DimensionTextInsideHorizontal = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 74:
		this.DimensionTextOutsideHorizontal = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 75:
		this.SuppressFirstDimensionExtensionLine = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 76:
		this.SuppressSecondDimensionExtensionLine = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 77:
		this.TextAboveDimensionLine = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 78:
		this.DimensionUnitZeroSuppression = UnitZeroSuppression(codePair.Value.(ShortCodePairValue).Value)
	case 79:
		this.DimensionAngleZeroSuppression = UnitZeroSuppression(codePair.Value.(ShortCodePairValue).Value)
	case 140:
		this.DimensioningTextHeight = codePair.Value.(DoubleCodePairValue).Value
	case 141:
		this.CenterMarkSize = codePair.Value.(DoubleCodePairValue).Value
	case 142:
		this.DimensioningTickSize = codePair.Value.(DoubleCodePairValue).Value
	case 143:
		this.AlternateDimensioningScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 144:
		this.DimensionLinearMeasurementScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 145:
		this.DimensionVerticalTextPosition = codePair.Value.(DoubleCodePairValue).Value
	case 146:
		this.DimensionToleranceDisplaceScaleFactor = codePair.Value.(DoubleCodePairValue).Value
	case 147:
		this.DimensionLineGap = codePair.Value.(DoubleCodePairValue).Value
	case 148:
		this.AlternateDimensioningUnitRounding = codePair.Value.(DoubleCodePairValue).Value
	case 170:
		this.UseAlternateDimensioning = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 171:
		this.AlternateDimensioningDecimalPlaces = codePair.Value.(ShortCodePairValue).Value
	case 172:
		this.ForceDimensionLineExtensionsOutsideIfTextExists = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 173:
		this.UseSeparateArrowBlocksForDimensions = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 174:
		this.ForceDimensionTextInsideExtensions = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 175:
		this.SuppressOutsideExtensionDimensionLines = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 176:
		this.DimensionLineColor = Color(codePair.Value.(ShortCodePairValue).Value)
	case 177:
		this.DimensionExtensionLineColor = Color(codePair.Value.(ShortCodePairValue).Value)
	case 178:
		this.DimensionTextColor = Color(codePair.Value.(ShortCodePairValue).Value)
	case 179:
		this.AngularDimensionPrecision = codePair.Value.(ShortCodePairValue).Value
	case 270:
		this.DimensionUnitFormat = UnitFormat(codePair.Value.(ShortCodePairValue).Value)
	case 271:
		this.DimensionUnitToleranceDecimalPlaces = codePair.Value.(ShortCodePairValue).Value
	case 272:
		this.DimensionToleraceDecimalPlaces = codePair.Value.(ShortCodePairValue).Value
	case 273:
		this.AlternateDimensioningUnits = UnitFormat(codePair.Value.(ShortCodePairValue).Value)
	case 274:
		this.AlternateDimensioningToleranceDecimalPlaces = codePair.Value.(ShortCodePairValue).Value
	case 275:
		this.DimensioningAngleFormat = AngleFormat(codePair.Value.(ShortCodePairValue).Value)
	case 276:
		this.DimensionPrecision = codePair.Value.(ShortCodePairValue).Value
	case 277:
		this.DimensionNonAngularUnits = NonAngularUnits(codePair.Value.(ShortCodePairValue).Value)
	case 278:
		this.DimensionDecilamSeparatorChar = rune(codePair.Value.(ShortCodePairValue).Value)
	case 279:
		this.DimensionTextMovementRule = DimensionTextMovementRule(codePair.Value.(ShortCodePairValue).Value)
	case 280:
		this.DimensionTextJustification = DimensionTextJustification(codePair.Value.(ShortCodePairValue).Value)
	case 281:
		this.SuppressFirstDimensionExtensionLine = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 282:
		this.SuppressSecondDimensionExtensionLine = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 283:
		this.DimensionToleranceVerticalJustification = Justification(codePair.Value.(ShortCodePairValue).Value)
	case 284:
		this.DimensionToleranceZeroSuppression = UnitZeroSuppression(codePair.Value.(ShortCodePairValue).Value)
	case 285:
		this.AlternateDimensioningZeroSuppression = UnitZeroSuppression(codePair.Value.(ShortCodePairValue).Value)
	case 286:
		this.AlternateDimensioningToleranceZeroSuppression = UnitZeroSuppression(codePair.Value.(ShortCodePairValue).Value)
	case 287:
		this.DimensionTextAndArrowPlacement = DimensionFit(codePair.Value.(ShortCodePairValue).Value)
	case 288:
		this.DimensionCursorControlsTextPosition = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 289:
		this.DimensionTextAndArrowPlacement = DimensionFit(codePair.Value.(ShortCodePairValue).Value)
	case 340:
		this.DimensionTextStyle = codePair.Value.(StringCodePairValue).Value
	case 341:
		this.DimensionLeaderBlockName = codePair.Value.(StringCodePairValue).Value
	case 342:
		this.ArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 343:
		this.FirstArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 344:
		this.SecondArrowBlockName = codePair.Value.(StringCodePairValue).Value
	case 371:
		this.DimensionLineWeight = LineWeight(codePair.Value.(ShortCodePairValue).Value)
	case 372:
		this.DimensionExtensionLineWeight = LineWeight(codePair.Value.(ShortCodePairValue).Value)
	default:
	}
}

func tablePairsDimStyles(tableHandle Handle, items []DimStyle, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "DIMSTYLE"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "DIMSTYLE"))
		pairs = append(pairs, NewStringCodePair(105, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *DimStyle) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbDimStyleTable"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	pairs = append(pairs, NewShortCodePair(70, int16(this.Flags)))
	pairs = append(pairs, NewStringCodePair(3, this.DimensioningSuffix))
	pairs = append(pairs, NewStringCodePair(4, this.AlternateDimensioningSuffix))
	if version <= R14 {
		pairs = append(pairs, NewStringCodePair(5, this.ArrowBlockName))
	}
	if version <= R14 {
		pairs = append(pairs, NewStringCodePair(6, this.FirstArrowBlockName))
	}
	if version <= R14 {
		pairs = append(pairs, NewStringCodePair(7, this.SecondArrowBlockName))
	}
	pairs = append(pairs, NewDoubleCodePair(40, this.DimensioningScaleFactor))
	pairs = append(pairs, NewDoubleCodePair(41, this.DimensioningArrowSize))
	pairs = append(pairs, NewDoubleCodePair(42, this.DimensionExtensionLineOffset))
	pairs = append(pairs, NewDoubleCodePair(43, this.DimensionLineIncrement))
	pairs = append(pairs, NewDoubleCodePair(44, this.DimensionExtensionLineIncrement))
	pairs = append(pairs, NewDoubleCodePair(45, this.DimensionDistanceRoundingValue))
	pairs = append(pairs, NewDoubleCodePair(46, this.DimensionLineExtension))
	pairs = append(pairs, NewDoubleCodePair(47, this.DimensionPlusTolerance))
	pairs = append(pairs, NewDoubleCodePair(48, this.DimensionMinusTolerance))
	pairs = append(pairs, NewShortCodePair(71, shortFromBool(this.GenerateDimensionTolerances)))
	pairs = append(pairs, NewShortCodePair(72, shortFromBool(this.GenerateDimensionLimits)))
	pairs = append(pairs, NewShortCodePair(73, shortFromBool(this.DimensionTextInsideHorizontal)))
	pairs = append(pairs, NewShortCodePair(74, shortFromBool(this.DimensionTextOutsideHorizontal)))
	pairs = append(pairs, NewShortCodePair(75, shortFromBool(this.SuppressFirstDimensionExtensionLine)))
	pairs = append(pairs, NewShortCodePair(76, shortFromBool(this.SuppressSecondDimensionExtensionLine)))
	pairs = append(pairs, NewShortCodePair(77, shortFromBool(this.TextAboveDimensionLine)))
	pairs = append(pairs, NewShortCodePair(78, int16(this.DimensionUnitZeroSuppression)))
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(79, int16(this.DimensionAngleZeroSuppression)))
	}
	pairs = append(pairs, NewDoubleCodePair(140, this.DimensioningTextHeight))
	pairs = append(pairs, NewDoubleCodePair(141, this.CenterMarkSize))
	pairs = append(pairs, NewDoubleCodePair(142, this.DimensioningTickSize))
	pairs = append(pairs, NewDoubleCodePair(143, this.AlternateDimensioningScaleFactor))
	pairs = append(pairs, NewDoubleCodePair(144, this.DimensionLinearMeasurementScaleFactor))
	pairs = append(pairs, NewDoubleCodePair(145, this.DimensionVerticalTextPosition))
	pairs = append(pairs, NewDoubleCodePair(146, this.DimensionToleranceDisplaceScaleFactor))
	pairs = append(pairs, NewDoubleCodePair(147, this.DimensionLineGap))
	if version >= R2000 {
		pairs = append(pairs, NewDoubleCodePair(148, this.AlternateDimensioningUnitRounding))
	}
	pairs = append(pairs, NewShortCodePair(170, shortFromBool(this.UseAlternateDimensioning)))
	pairs = append(pairs, NewShortCodePair(171, this.AlternateDimensioningDecimalPlaces))
	pairs = append(pairs, NewShortCodePair(172, shortFromBool(this.ForceDimensionLineExtensionsOutsideIfTextExists)))
	pairs = append(pairs, NewShortCodePair(173, shortFromBool(this.UseSeparateArrowBlocksForDimensions)))
	pairs = append(pairs, NewShortCodePair(174, shortFromBool(this.ForceDimensionTextInsideExtensions)))
	pairs = append(pairs, NewShortCodePair(175, shortFromBool(this.SuppressOutsideExtensionDimensionLines)))
	pairs = append(pairs, NewShortCodePair(176, int16(this.DimensionLineColor)))
	pairs = append(pairs, NewShortCodePair(177, int16(this.DimensionExtensionLineColor)))
	pairs = append(pairs, NewShortCodePair(178, int16(this.DimensionTextColor)))
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(179, this.AngularDimensionPrecision))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(270, int16(this.DimensionUnitFormat)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(271, this.DimensionUnitToleranceDecimalPlaces))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(272, this.DimensionToleraceDecimalPlaces))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(273, int16(this.AlternateDimensioningUnits)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(274, this.AlternateDimensioningToleranceDecimalPlaces))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(275, int16(this.DimensioningAngleFormat)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(276, this.DimensionPrecision))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(277, int16(this.DimensionNonAngularUnits)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(278, int16(this.DimensionDecilamSeparatorChar)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(279, int16(this.DimensionTextMovementRule)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(280, int16(this.DimensionTextJustification)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(281, shortFromBool(this.SuppressFirstDimensionExtensionLine)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(282, shortFromBool(this.SuppressSecondDimensionExtensionLine)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(283, int16(this.DimensionToleranceVerticalJustification)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(284, int16(this.DimensionToleranceZeroSuppression)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(285, int16(this.AlternateDimensioningZeroSuppression)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(286, int16(this.AlternateDimensioningToleranceZeroSuppression)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(287, int16(this.DimensionTextAndArrowPlacement)))
	}
	if version >= R13 {
		pairs = append(pairs, NewShortCodePair(288, shortFromBool(this.DimensionCursorControlsTextPosition)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(289, int16(this.DimensionTextAndArrowPlacement)))
	}
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(340, this.DimensionTextStyle))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(341, this.DimensionLeaderBlockName))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(342, this.ArrowBlockName))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(343, this.FirstArrowBlockName))
	}
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(344, this.SecondArrowBlockName))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(371, int16(this.DimensionLineWeight)))
	}
	if version >= R2000 {
		pairs = append(pairs, NewShortCodePair(372, int16(this.DimensionExtensionLineWeight)))
	}
	return
}

type BlockRecord struct {
	handle Handle
	Name string
	LayoutHandle string
	InsertionUnits Units
	Explodability bool
	Scalability bool
	bitmapPreviewData []string
}

func NewBlockRecord() *BlockRecord {
	return &BlockRecord{
		Name: "",
		LayoutHandle: "",
		InsertionUnits: UnitsUnitless,
		Explodability: true,
		Scalability: true,
		bitmapPreviewData: []string{},
	}
}

func (this *BlockRecord) Handle() Handle {
	return this.handle
}

func (this *BlockRecord) SetHandle(val Handle) {
	this.handle = val
}

func readBlockRecords(drawing *Drawing, np CodePair, reader codePairReader) (nextPair CodePair, error error) {
	nextPair = np
	for error == nil && !nextPair.isEndTable() {
		if nextPair.Code != 0 || nextPair.Value.(StringCodePairValue).Value != "BLOCK_RECORD" {
			error = fmt.Errorf("expected [0/BLOCK_RECORD] but found [%s]", nextPair.String())
			return
		}
		item := *NewBlockRecord()
		nextPair, error = reader.readCodePair()
		for error == nil && nextPair.Code != 0 {
			item.tryApplyCodePair(nextPair)
			nextPair, error = reader.readCodePair()
		}
		drawing.BlockRecords = append(drawing.BlockRecords, item)
	}
	return
}

func (this *BlockRecord) tryApplyCodePair(codePair CodePair) {
	switch codePair.Code {
	case 2:
		this.Name = codePair.Value.(StringCodePairValue).Value
	case 340:
		this.LayoutHandle = codePair.Value.(StringCodePairValue).Value
	case 70:
		this.InsertionUnits = Units(codePair.Value.(ShortCodePairValue).Value)
	case 280:
		this.Explodability = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 281:
		this.Scalability = boolFromShort(codePair.Value.(ShortCodePairValue).Value)
	case 310:
		this.bitmapPreviewData = append(this.bitmapPreviewData, codePair.Value.(StringCodePairValue).Value)
	default:
	}
}

func tablePairsBlockRecords(tableHandle Handle, items []BlockRecord, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, NewStringCodePair(0, "TABLE"))
	pairs = append(pairs, NewStringCodePair(2, "BLOCK_RECORD"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(tableHandle)))
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTable"))
	}
	pairs = append(pairs, NewShortCodePair(70, int16(len(items))))
	for _, item := range items {
		pairs = append(pairs, NewStringCodePair(0, "BLOCK_RECORD"))
		pairs = append(pairs, NewStringCodePair(5, stringFromHandle(item.Handle())))
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, item.codePairs(version)...)
	}
	pairs = append(pairs, NewStringCodePair(0, "ENDTAB"))
	return
}

func (this *BlockRecord) codePairs(version AcadVersion) (pairs []CodePair) {
	if version >= R13 {
		pairs = append(pairs, NewStringCodePair(100, "AcDbSymbolTableRecord"))
		pairs = append(pairs, NewStringCodePair(100, "AcDbBlockTableRecord"))
	}
	pairs = append(pairs, NewStringCodePair(2, this.Name))
	if version >= R2000 {
		pairs = append(pairs, NewStringCodePair(340, this.LayoutHandle))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(70, int16(this.InsertionUnits)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(this.Explodability)))
	}
	if version >= R2007 {
		pairs = append(pairs, NewShortCodePair(281, shortFromBool(this.Scalability)))
	}
	if version >= R2000 {
		for _, val := range this.bitmapPreviewData {
			pairs = append(pairs, NewStringCodePair(310, val))
		}
	}
	return
}

func readSpecificTable(drawing *Drawing, np CodePair, reader codePairReader, tableType string) (nextPair CodePair, error error) {
	nextPair = np
	if error != nil {
		return
	}

	// swallow until 0/<item>
	for error == nil && nextPair.Code != 0 {
		nextPair, error = reader.readCodePair()
	}

	switch tableType {
	case "VPORT":
		nextPair, error = readViewPorts(drawing, nextPair, reader)
	case "LTYPE":
		nextPair, error = readLineTypes(drawing, nextPair, reader)
	case "LAYER":
		nextPair, error = readLayers(drawing, nextPair, reader)
	case "STYLE":
		nextPair, error = readStyles(drawing, nextPair, reader)
	case "VIEW":
		nextPair, error = readViews(drawing, nextPair, reader)
	case "UCS":
		nextPair, error = readUcss(drawing, nextPair, reader)
	case "APPID":
		nextPair, error = readAppIds(drawing, nextPair, reader)
	case "DIMSTYLE":
		nextPair, error = readDimStyles(drawing, nextPair, reader)
	case "BLOCK_RECORD":
		nextPair, error = readBlockRecords(drawing, nextPair, reader)
	}
	return
}

func getTablePairs(drawing *Drawing, version AcadVersion) (pairs []CodePair) {
	pairs = append(pairs, tablePairsViewPorts(drawing.viewPortTableHandle, drawing.ViewPorts, version)...)
	pairs = append(pairs, tablePairsLineTypes(drawing.lineTypeTableHandle, drawing.LineTypes, version)...)
	pairs = append(pairs, tablePairsLayers(drawing.layerTableHandle, drawing.Layers, version)...)
	pairs = append(pairs, tablePairsStyles(drawing.styleTableHandle, drawing.Styles, version)...)
	pairs = append(pairs, tablePairsViews(drawing.viewTableHandle, drawing.Views, version)...)
	pairs = append(pairs, tablePairsUcss(drawing.ucsTableHandle, drawing.Ucss, version)...)
	if version >= R12 {
		pairs = append(pairs, tablePairsAppIds(drawing.appIdTableHandle, drawing.AppIds, version)...)
	}
	if version >= R12 {
		pairs = append(pairs, tablePairsDimStyles(drawing.dimStyleTableHandle, drawing.DimStyles, version)...)
	}
	if version >= R13 {
		pairs = append(pairs, tablePairsBlockRecords(drawing.blockRecordTableHandle, drawing.BlockRecords, version)...)
	}
	return
}
func assignTableHandles(drawing *Drawing, nextHandle Handle) Handle {
	drawing.viewPortTableHandle = nextHandle
	nextHandle++
	for i := range drawing.ViewPorts {
		item := &drawing.ViewPorts[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.lineTypeTableHandle = nextHandle
	nextHandle++
	for i := range drawing.LineTypes {
		item := &drawing.LineTypes[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.layerTableHandle = nextHandle
	nextHandle++
	for i := range drawing.Layers {
		item := &drawing.Layers[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.styleTableHandle = nextHandle
	nextHandle++
	for i := range drawing.Styles {
		item := &drawing.Styles[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.viewTableHandle = nextHandle
	nextHandle++
	for i := range drawing.Views {
		item := &drawing.Views[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.ucsTableHandle = nextHandle
	nextHandle++
	for i := range drawing.Ucss {
		item := &drawing.Ucss[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.appIdTableHandle = nextHandle
	nextHandle++
	for i := range drawing.AppIds {
		item := &drawing.AppIds[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.dimStyleTableHandle = nextHandle
	nextHandle++
	for i := range drawing.DimStyles {
		item := &drawing.DimStyles[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	drawing.blockRecordTableHandle = nextHandle
	nextHandle++
	for i := range drawing.BlockRecords {
		item := &drawing.BlockRecords[i]
		if (*item).Handle() == 0 {
			(*item).SetHandle(nextHandle)
			nextHandle++
		}
	}
	return nextHandle
}

