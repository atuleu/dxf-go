// Code generated at build time; DO NOT EDIT.

package dxf

import (
	"fmt"
)

type AngleFormat int16

const (
	AngleFormatDecimalDegrees AngleFormat = iota
	AngleFormatDegreesMinutesSeconds
	AngleFormatGradians
	AngleFormatRadians
	AngleFormatSurveyorUnits
)

func (this AngleFormat) String() string {
	switch this {
	case AngleFormatDecimalDegrees:
		return "AngleFormatDecimalDegrees"
	case AngleFormatDegreesMinutesSeconds:
		return "AngleFormatDegreesMinutesSeconds"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type AngleDirection int16

const (
	AngleDirectionCounterClockwise AngleDirection = iota
	AngleDirectionClockwise
)

func (this AngleDirection) String() string {
	switch this {
	case AngleDirectionCounterClockwise:
		return "AngleDirectionCounterClockwise"
	case AngleDirectionClockwise:
		return "AngleDirectionClockwise"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type AttachmentPoint int16

const (
	AttachmentPointTopLeft AttachmentPoint = 1
	AttachmentPointTopCenter AttachmentPoint = 2
	AttachmentPointTopRight AttachmentPoint = 3
	AttachmentPointMiddleLeft AttachmentPoint = 4
	AttachmentPointMiddleCenter AttachmentPoint = 5
	AttachmentPointMiddleRight AttachmentPoint = 6
	AttachmentPointBottomLeft AttachmentPoint = 7
	AttachmentPointBottomCenter AttachmentPoint = 8
	AttachmentPointBottomRight AttachmentPoint = 9
)

func (this AttachmentPoint) String() string {
	switch this {
	case AttachmentPointTopLeft:
		return "AttachmentPointTopLeft"
	case AttachmentPointTopCenter:
		return "AttachmentPointTopCenter"
	case AttachmentPointTopRight:
		return "AttachmentPointTopRight"
	case AttachmentPointMiddleLeft:
		return "AttachmentPointMiddleLeft"
	case AttachmentPointMiddleCenter:
		return "AttachmentPointMiddleCenter"
	case AttachmentPointMiddleRight:
		return "AttachmentPointMiddleRight"
	case AttachmentPointBottomLeft:
		return "AttachmentPointBottomLeft"
	case AttachmentPointBottomCenter:
		return "AttachmentPointBottomCenter"
	case AttachmentPointBottomRight:
		return "AttachmentPointBottomRight"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type AttributeVisibility int16

const (
	AttributeVisibilityNone AttributeVisibility = iota
	AttributeVisibilityNormal
	AttributeVisibilityAll
)

func (this AttributeVisibility) String() string {
	switch this {
	case AttributeVisibilityNone:
		return "AttributeVisibilityNone"
	case AttributeVisibilityNormal:
		return "AttributeVisibilityNormal"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type BackgroundFillSetting int16

const (
	BackgroundFillSettingOff BackgroundFillSetting = iota
	BackgroundFillSettingUseBackgroundFillColor
	BackgroundFillSettingUseDrawingWindowColor
)

func (this BackgroundFillSetting) String() string {
	switch this {
	case BackgroundFillSettingOff:
		return "BackgroundFillSettingOff"
	case BackgroundFillSettingUseBackgroundFillColor:
		return "BackgroundFillSettingUseBackgroundFillColor"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type CoordinateDisplay int16

const (
	CoordinateDisplayStatic CoordinateDisplay = iota
	CoordinateDisplayContinuousUpdate
	CoordinateDisplayDistanceAngleFormat
)

func (this CoordinateDisplay) String() string {
	switch this {
	case CoordinateDisplayStatic:
		return "CoordinateDisplayStatic"
	case CoordinateDisplayContinuousUpdate:
		return "CoordinateDisplayContinuousUpdate"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DefaultLightingType int16

const (
	DefaultLightingTypeOneDistantLight DefaultLightingType = iota
	DefaultLightingTypeTwoDistanceLights
)

func (this DefaultLightingType) String() string {
	switch this {
	case DefaultLightingTypeOneDistantLight:
		return "DefaultLightingTypeOneDistantLight"
	case DefaultLightingTypeTwoDistanceLights:
		return "DefaultLightingTypeTwoDistanceLights"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionArcSymbolDisplayMode int16

const (
	DimensionArcSymbolDisplayModeSymbolBeforeText DimensionArcSymbolDisplayMode = iota
	DimensionArcSymbolDisplayModeSymbolAboveText
	DimensionArcSymbolDisplayModeSuppress
)

func (this DimensionArcSymbolDisplayMode) String() string {
	switch this {
	case DimensionArcSymbolDisplayModeSymbolBeforeText:
		return "DimensionArcSymbolDisplayModeSymbolBeforeText"
	case DimensionArcSymbolDisplayModeSymbolAboveText:
		return "DimensionArcSymbolDisplayModeSymbolAboveText"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionAssociativity int16

const (
	DimensionAssociativityNoAssociationExploded DimensionAssociativity = iota
	DimensionAssociativityNonAssociativeObjects
	DimensionAssociativityAssociativeObjects
)

func (this DimensionAssociativity) String() string {
	switch this {
	case DimensionAssociativityNoAssociationExploded:
		return "DimensionAssociativityNoAssociationExploded"
	case DimensionAssociativityNonAssociativeObjects:
		return "DimensionAssociativityNonAssociativeObjects"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionFit int16

const (
	DimensionFitTextAndArrowsOutsideLines DimensionFit = iota
	DimensionFitMoveArrowsFirst
	DimensionFitMoveTextFirst
	DimensionFitMoveEitherForBestFit
)

func (this DimensionFit) String() string {
	switch this {
	case DimensionFitTextAndArrowsOutsideLines:
		return "DimensionFitTextAndArrowsOutsideLines"
	case DimensionFitMoveArrowsFirst:
		return "DimensionFitMoveArrowsFirst"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionFractionFormat int16

const (
	DimensionFractionFormatHorizontalStacking DimensionFractionFormat = iota
	DimensionFractionFormatDiagonalStacking
	DimensionFractionFormatNotStacked
)

func (this DimensionFractionFormat) String() string {
	switch this {
	case DimensionFractionFormatHorizontalStacking:
		return "DimensionFractionFormatHorizontalStacking"
	case DimensionFractionFormatDiagonalStacking:
		return "DimensionFractionFormatDiagonalStacking"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionTextBackgroundColorMode int16

const (
	DimensionTextBackgroundColorModeNone DimensionTextBackgroundColorMode = iota
	DimensionTextBackgroundColorModeUseDrawingBackground
	DimensionTextBackgroundColorModeCustom
)

func (this DimensionTextBackgroundColorMode) String() string {
	switch this {
	case DimensionTextBackgroundColorModeNone:
		return "DimensionTextBackgroundColorModeNone"
	case DimensionTextBackgroundColorModeUseDrawingBackground:
		return "DimensionTextBackgroundColorModeUseDrawingBackground"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionTextJustification int16

const (
	DimensionTextJustificationAboveLineCenter DimensionTextJustification = iota
	DimensionTextJustificationAboveLineNextToFirstExtension
	DimensionTextJustificationAboveLineNextToSecondExtension
	DimensionTextJustificationAboveLineCenteredOnFirstExtension
	DimensionTextJustificationAboveLineCenteredOnSecondExtension
)

func (this DimensionTextJustification) String() string {
	switch this {
	case DimensionTextJustificationAboveLineCenter:
		return "DimensionTextJustificationAboveLineCenter"
	case DimensionTextJustificationAboveLineNextToFirstExtension:
		return "DimensionTextJustificationAboveLineNextToFirstExtension"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionTextMovementRule int16

const (
	DimensionTextMovementRuleMoveLineWithText DimensionTextMovementRule = iota
	DimensionTextMovementRuleAddLeaderWhenTextIsMoved
	DimensionTextMovementRuleMoveTextFreely
)

func (this DimensionTextMovementRule) String() string {
	switch this {
	case DimensionTextMovementRuleMoveLineWithText:
		return "DimensionTextMovementRuleMoveLineWithText"
	case DimensionTextMovementRuleAddLeaderWhenTextIsMoved:
		return "DimensionTextMovementRuleAddLeaderWhenTextIsMoved"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DimensionType int16

const (
	DimensionTypeRotatedHorizontalOrVertical DimensionType = iota
	DimensionTypeAligned
	DimensionTypeAngular
	DimensionTypeDiameter
	DimensionTypeRadius
	DimensionTypeAngularThreePoint
	DimensionTypeOrdinate
)

func (this DimensionType) String() string {
	switch this {
	case DimensionTypeRotatedHorizontalOrVertical:
		return "DimensionTypeRotatedHorizontalOrVertical"
	case DimensionTypeAligned:
		return "DimensionTypeAligned"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DragMode int16

const (
	DragModeOff DragMode = iota
	DragModeOn
	DragModeAuto
)

func (this DragMode) String() string {
	switch this {
	case DragModeOff:
		return "DragModeOff"
	case DragModeOn:
		return "DragModeOn"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DrawingDirection int16

const (
	DrawingDirectionLeftToRight DrawingDirection = 1
	DrawingDirectionTopToBottom DrawingDirection = 3
	DrawingDirectionByStyle DrawingDirection = 5
)

func (this DrawingDirection) String() string {
	switch this {
	case DrawingDirectionLeftToRight:
		return "DrawingDirectionLeftToRight"
	case DrawingDirectionTopToBottom:
		return "DrawingDirectionTopToBottom"
	case DrawingDirectionByStyle:
		return "DrawingDirectionByStyle"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type DrawingUnits int16

const (
	DrawingUnitsEnglish DrawingUnits = iota
	DrawingUnitsMetric
)

func (this DrawingUnits) String() string {
	switch this {
	case DrawingUnitsEnglish:
		return "DrawingUnitsEnglish"
	case DrawingUnitsMetric:
		return "DrawingUnitsMetric"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type Dwf3DPrecision int16

const (
	Dwf3DPrecision_ Dwf3DPrecision = iota
	Dwf3DPrecisionDeviation1
	Dwf3DPrecisionDeviation0_5
	Dwf3DPrecisionDeviation0_2
	Dwf3DPrecisionDeviation0_1
	Dwf3DPrecisionDeviation0_01
	Dwf3DPrecisionDeviation0_001
)

func (this Dwf3DPrecision) String() string {
	switch this {
	case Dwf3DPrecision_:
		return "Dwf3DPrecision_"
	case Dwf3DPrecisionDeviation1:
		return "Dwf3DPrecisionDeviation1"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type EndCapSetting int16

const (
	EndCapSettingNone EndCapSetting = iota
	EndCapSettingRound
	EndCapSettingAngle
	EndCapSettingSquare
)

func (this EndCapSetting) String() string {
	switch this {
	case EndCapSettingNone:
		return "EndCapSettingNone"
	case EndCapSettingRound:
		return "EndCapSettingRound"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type FontType int16

const (
	FontTypeTTF FontType = iota
	FontTypeSHX
)

func (this FontType) String() string {
	switch this {
	case FontTypeTTF:
		return "FontTypeTTF"
	case FontTypeSHX:
		return "FontTypeSHX"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type HelixConstraint int16

const (
	HelixConstraintConstrainTurnHeight HelixConstraint = iota
	HelixConstraintConstrainTurns
	HelixConstraintConstrainHeight
)

func (this HelixConstraint) String() string {
	switch this {
	case HelixConstraintConstrainTurnHeight:
		return "HelixConstraintConstrainTurnHeight"
	case HelixConstraintConstrainTurns:
		return "HelixConstraintConstrainTurns"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type HorizontalTextJustification int16

const (
	HorizontalTextJustificationLeft HorizontalTextJustification = iota
	HorizontalTextJustificationCenter
	HorizontalTextJustificationRight
	HorizontalTextJustificationAligned
	HorizontalTextJustificationMiddle
	HorizontalTextJustificationFit
)

func (this HorizontalTextJustification) String() string {
	switch this {
	case HorizontalTextJustificationLeft:
		return "HorizontalTextJustificationLeft"
	case HorizontalTextJustificationCenter:
		return "HorizontalTextJustificationCenter"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type ImageClippingBoundaryType int16

const (
	ImageClippingBoundaryTypeRectangular ImageClippingBoundaryType = 1
	ImageClippingBoundaryTypePolygonal ImageClippingBoundaryType = 2
)

func (this ImageClippingBoundaryType) String() string {
	switch this {
	case ImageClippingBoundaryTypeRectangular:
		return "ImageClippingBoundaryTypeRectangular"
	case ImageClippingBoundaryTypePolygonal:
		return "ImageClippingBoundaryTypePolygonal"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type JoinStyle int16

const (
	JoinStyleNone JoinStyle = iota
	JoinStyleRound
	JoinStyleAngle
	JoinStyleFlat
)

func (this JoinStyle) String() string {
	switch this {
	case JoinStyleNone:
		return "JoinStyleNone"
	case JoinStyleRound:
		return "JoinStyleRound"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type Justification int16

const (
	JustificationTop Justification = iota
	JustificationMiddle
	JustificationBottom
)

func (this Justification) String() string {
	switch this {
	case JustificationTop:
		return "JustificationTop"
	case JustificationMiddle:
		return "JustificationMiddle"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LayerAndSpatialIndexSaveMode int16

const (
	LayerAndSpatialIndexSaveModeNone LayerAndSpatialIndexSaveMode = iota
	LayerAndSpatialIndexSaveModeLayerIndex
	LayerAndSpatialIndexSaveModeSpatialIndex
	LayerAndSpatialIndexSaveModeLayerAndSpatialIndex
)

func (this LayerAndSpatialIndexSaveMode) String() string {
	switch this {
	case LayerAndSpatialIndexSaveModeNone:
		return "LayerAndSpatialIndexSaveModeNone"
	case LayerAndSpatialIndexSaveModeLayerIndex:
		return "LayerAndSpatialIndexSaveModeLayerIndex"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LeaderCreationAnnotationType int16

const (
	LeaderCreationAnnotationTypeWithTextAnnotation LeaderCreationAnnotationType = iota
	LeaderCreationAnnotationTypeWithToleranceAnnotation
	LeaderCreationAnnotationTypeWithBlockReferenceAnnotation
	LeaderCreationAnnotationTypeNoAnnotation
)

func (this LeaderCreationAnnotationType) String() string {
	switch this {
	case LeaderCreationAnnotationTypeWithTextAnnotation:
		return "LeaderCreationAnnotationTypeWithTextAnnotation"
	case LeaderCreationAnnotationTypeWithToleranceAnnotation:
		return "LeaderCreationAnnotationTypeWithToleranceAnnotation"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LeaderHooklineDirection int16

const (
	LeaderHooklineDirectionOppositeFromHorizontalVector LeaderHooklineDirection = iota
	LeaderHooklineDirectionSameAsHorizontalVector
)

func (this LeaderHooklineDirection) String() string {
	switch this {
	case LeaderHooklineDirectionOppositeFromHorizontalVector:
		return "LeaderHooklineDirectionOppositeFromHorizontalVector"
	case LeaderHooklineDirectionSameAsHorizontalVector:
		return "LeaderHooklineDirectionSameAsHorizontalVector"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LeaderPathType int16

const (
	LeaderPathTypeStraightLineSegments LeaderPathType = iota
	LeaderPathTypeSpline
)

func (this LeaderPathType) String() string {
	switch this {
	case LeaderPathTypeStraightLineSegments:
		return "LeaderPathTypeStraightLineSegments"
	case LeaderPathTypeSpline:
		return "LeaderPathTypeSpline"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LightType int16

const (
	LightTypeDistant LightType = 1
	LightTypePoint LightType = 2
	LightTypeSpot LightType = 3
)

func (this LightType) String() string {
	switch this {
	case LightTypeDistant:
		return "LightTypeDistant"
	case LightTypePoint:
		return "LightTypePoint"
	case LightTypeSpot:
		return "LightTypeSpot"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LightAttenuationType int16

const (
	LightAttenuationTypeNone LightAttenuationType = iota
	LightAttenuationTypeInverseLinear
	LightAttenuationTypeInverseSquare
)

func (this LightAttenuationType) String() string {
	switch this {
	case LightAttenuationTypeNone:
		return "LightAttenuationTypeNone"
	case LightAttenuationTypeInverseLinear:
		return "LightAttenuationTypeInverseLinear"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LineTypeStyle int16

const (
	LineTypeStyleOff LineTypeStyle = iota
	LineTypeStyleSolid
	LineTypeStyleDashed
	LineTypeStyleDotted
	LineTypeStyleShortDash
	LineTypeStyleMediumDash
	LineTypeStyleLongDash
	LineTypeStyleDoubleShortDash
	LineTypeStyleDoubleMediumDash
	LineTypeStyleDoubleLongDash
	LineTypeStyleMediumLongDash
	LineTypeStyleSparseDot
)

func (this LineTypeStyle) String() string {
	switch this {
	case LineTypeStyleOff:
		return "LineTypeStyleOff"
	case LineTypeStyleSolid:
		return "LineTypeStyleSolid"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type LoftedObjectNormalMode int16

const (
	LoftedObjectNormalModeRuled LoftedObjectNormalMode = iota
	LoftedObjectNormalModeSmoothFit
	LoftedObjectNormalModeStartCrossSection
	LoftedObjectNormalModeEndCrossSection
	LoftedObjectNormalModeStartAndEndCrossSections
	LoftedObjectNormalModeAllCrossSections
	LoftedObjectNormalModeUseDraftAngleAndMagnitude
)

func (this LoftedObjectNormalMode) String() string {
	switch this {
	case LoftedObjectNormalModeRuled:
		return "LoftedObjectNormalModeRuled"
	case LoftedObjectNormalModeSmoothFit:
		return "LoftedObjectNormalModeSmoothFit"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type MTextFlag int16

const (
	MTextFlagMultilineAttribute MTextFlag = 2
	MTextFlagConstantMultilineAttributeDefinition MTextFlag = 4
)

func (this MTextFlag) String() string {
	switch this {
	case MTextFlagMultilineAttribute:
		return "MTextFlagMultilineAttribute"
	case MTextFlagConstantMultilineAttributeDefinition:
		return "MTextFlagConstantMultilineAttributeDefinition"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type MTextLineSpacingStyle int16

const (
	MTextLineSpacingStyleAtLeast MTextLineSpacingStyle = 1
	MTextLineSpacingStyleExact MTextLineSpacingStyle = 2
)

func (this MTextLineSpacingStyle) String() string {
	switch this {
	case MTextLineSpacingStyleAtLeast:
		return "MTextLineSpacingStyleAtLeast"
	case MTextLineSpacingStyleExact:
		return "MTextLineSpacingStyleExact"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type NonAngularUnits int16

const (
	NonAngularUnitsScientific NonAngularUnits = iota
	NonAngularUnitsDecimal
	NonAngularUnitsEngineering
	NonAngularUnitsArchitectural
	NonAngularUnitsFractional
	NonAngularUnitsWindowsDesktop
)

func (this NonAngularUnits) String() string {
	switch this {
	case NonAngularUnitsScientific:
		return "NonAngularUnitsScientific"
	case NonAngularUnitsDecimal:
		return "NonAngularUnitsDecimal"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type OleObjectType int16

const (
	OleObjectTypeLink OleObjectType = 1
	OleObjectTypeEmbedded OleObjectType = 2
	OleObjectTypeStatic OleObjectType = 3
)

func (this OleObjectType) String() string {
	switch this {
	case OleObjectTypeLink:
		return "OleObjectTypeLink"
	case OleObjectTypeEmbedded:
		return "OleObjectTypeEmbedded"
	case OleObjectTypeStatic:
		return "OleObjectTypeStatic"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type OrthographicViewType int16

const (
	OrthographicViewTypeNone OrthographicViewType = iota
	OrthographicViewTypeTop
	OrthographicViewTypeBottom
	OrthographicViewTypeFront
	OrthographicViewTypeBack
	OrthographicViewTypeLeft
	OrthographicViewTypeRight
)

func (this OrthographicViewType) String() string {
	switch this {
	case OrthographicViewTypeNone:
		return "OrthographicViewTypeNone"
	case OrthographicViewTypeTop:
		return "OrthographicViewTypeTop"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type PickStyle int16

const (
	PickStyleNone PickStyle = iota
	PickStyleGroup
	PickStyleAssociativeHatch
	PickStyleGroupAndAssociativeHatch
)

func (this PickStyle) String() string {
	switch this {
	case PickStyleNone:
		return "PickStyleNone"
	case PickStyleGroup:
		return "PickStyleGroup"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type PolylineCurvedAndSmoothSurfaceType int16

const (
	PolylineCurvedAndSmoothSurfaceTypeNone PolylineCurvedAndSmoothSurfaceType = iota
	PolylineCurvedAndSmoothSurfaceTypeQuadraticBSpline
	PolylineCurvedAndSmoothSurfaceTypeCubicBSpline
	PolylineCurvedAndSmoothSurfaceTypeBezier
)

func (this PolylineCurvedAndSmoothSurfaceType) String() string {
	switch this {
	case PolylineCurvedAndSmoothSurfaceTypeNone:
		return "PolylineCurvedAndSmoothSurfaceTypeNone"
	case PolylineCurvedAndSmoothSurfaceTypeQuadraticBSpline:
		return "PolylineCurvedAndSmoothSurfaceTypeQuadraticBSpline"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type PolylineType int16

const (
	PolylineTypeBaseline PolylineType = iota
	PolylineTypeInternal
	PolylineTypeOutline
)

func (this PolylineType) String() string {
	switch this {
	case PolylineTypeBaseline:
		return "PolylineTypeBaseline"
	case PolylineTypeInternal:
		return "PolylineTypeInternal"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type PolySketchMode int16

const (
	PolySketchModeSketchLines PolySketchMode = iota
	PolySketchModeSketchPolyLines
)

func (this PolySketchMode) String() string {
	switch this {
	case PolySketchModeSketchLines:
		return "PolySketchModeSketchLines"
	case PolySketchModeSketchPolyLines:
		return "PolySketchModeSketchPolyLines"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type PlotStyle int16

const (
	PlotStyleByLayer PlotStyle = iota
	PlotStyleByBlock
	PlotStyleByDictionaryDefault
	PlotStyleByObjectId
)

func (this PlotStyle) String() string {
	switch this {
	case PlotStyleByLayer:
		return "PlotStyleByLayer"
	case PlotStyleByBlock:
		return "PlotStyleByBlock"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type ShadeEdgeMode int16

const (
	ShadeEdgeModeFacesShadedEdgeNotHighlighted ShadeEdgeMode = iota
	ShadeEdgeModeFacesShadedEdgesHighlightedInBlack
	ShadeEdgeModeFacesNotFilledEdgesInEntityColor
	ShadeEdgeModeFacesInEntityColorEdgesInBlack
)

func (this ShadeEdgeMode) String() string {
	switch this {
	case ShadeEdgeModeFacesShadedEdgeNotHighlighted:
		return "ShadeEdgeModeFacesShadedEdgeNotHighlighted"
	case ShadeEdgeModeFacesShadedEdgesHighlightedInBlack:
		return "ShadeEdgeModeFacesShadedEdgesHighlightedInBlack"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type ShadowMode int16

const (
	ShadowModeCastsAndReceivesShadows ShadowMode = iota
	ShadowModeCastsShadows
	ShadowModeReceivesShadows
	ShadowModeIgnoresShadows
)

func (this ShadowMode) String() string {
	switch this {
	case ShadowModeCastsAndReceivesShadows:
		return "ShadowModeCastsAndReceivesShadows"
	case ShadowModeCastsShadows:
		return "ShadowModeCastsShadows"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type ShadowType int16

const (
	ShadowTypeRayTraced ShadowType = iota
	ShadowTypeShadowMaps
)

func (this ShadowType) String() string {
	switch this {
	case ShadowTypeRayTraced:
		return "ShadowTypeRayTraced"
	case ShadowTypeShadowMaps:
		return "ShadowTypeShadowMaps"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type SnapIsometricPlane int16

const (
	SnapIsometricPlaneLeft SnapIsometricPlane = iota
	SnapIsometricPlaneTop
	SnapIsometricPlaneRight
)

func (this SnapIsometricPlane) String() string {
	switch this {
	case SnapIsometricPlaneLeft:
		return "SnapIsometricPlaneLeft"
	case SnapIsometricPlaneTop:
		return "SnapIsometricPlaneTop"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type SnapStyle int16

const (
	SnapStyleStandard SnapStyle = iota
	SnapStyleIsometric
)

func (this SnapStyle) String() string {
	switch this {
	case SnapStyleStandard:
		return "SnapStyleStandard"
	case SnapStyleIsometric:
		return "SnapStyleIsometric"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type SolidHistoryMode int16

const (
	SolidHistoryModeNone SolidHistoryMode = iota
	SolidHistoryModeDoesNotOverride
	SolidHistoryModeOverride
)

func (this SolidHistoryMode) String() string {
	switch this {
	case SolidHistoryModeNone:
		return "SolidHistoryModeNone"
	case SolidHistoryModeDoesNotOverride:
		return "SolidHistoryModeDoesNotOverride"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type TextDirection int16

const (
	TextDirectionLeftToRight TextDirection = iota
	TextDirectionRightToLeft
)

func (this TextDirection) String() string {
	switch this {
	case TextDirectionLeftToRight:
		return "TextDirectionLeftToRight"
	case TextDirectionRightToLeft:
		return "TextDirectionRightToLeft"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type TextLineSpacingStyle int16

const (
	TextLineSpacingStyleAtLeast TextLineSpacingStyle = 1
	TextLineSpacingStyleExact TextLineSpacingStyle = 2
)

func (this TextLineSpacingStyle) String() string {
	switch this {
	case TextLineSpacingStyleAtLeast:
		return "TextLineSpacingStyleAtLeast"
	case TextLineSpacingStyleExact:
		return "TextLineSpacingStyleExact"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type TileModeDescriptor int16

const (
	TileModeDescriptorInTiledViewport TileModeDescriptor = iota
	TileModeDescriptorInNonTiledViewport
)

func (this TileModeDescriptor) String() string {
	switch this {
	case TileModeDescriptorInTiledViewport:
		return "TileModeDescriptorInTiledViewport"
	case TileModeDescriptorInNonTiledViewport:
		return "TileModeDescriptorInNonTiledViewport"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type TimeZone int16

const (
	TimeZoneInternationalDateLineWest TimeZone = -12000
	TimeZoneMidwayIsland_Samoa TimeZone = -11000
	TimeZoneHawaii TimeZone = -10000
	TimeZoneAlaska TimeZone = -9000
	TimeZonePacificTime_US_Canada_SanFrancisco_Vancouver TimeZone = -8000
	TimeZoneArizona TimeZone = -7000
	TimeZoneChihuahua_LaPaz_Mazatlan TimeZone = -7000
	TimeZoneMountainTime_US_Canada TimeZone = -7000
	TimeZoneMazatlan TimeZone = -7002
	TimeZoneCentralAmerica TimeZone = -6000
	TimeZoneCentralTime_US_Canada TimeZone = -6001
	TimeZoneGuadalajara_MexicoCity_Monterrey TimeZone = -6002
	TimeZoneSaskatchewan TimeZone = -6003
	TimeZoneEasternTime_US_Canada_ TimeZone = -5000
	TimeZoneIndiana_East_ TimeZone = -5001
	TimeZoneBogota_Lima_Quito TimeZone = -5002
	TimeZoneAtlanticTime_Canada_ TimeZone = -4000
	TimeZoneCaracas_LaPaz TimeZone = -4001
	TimeZoneSantiago TimeZone = -4002
	TimeZoneNewfoundland TimeZone = -3300
	TimeZoneBrasilia TimeZone = -3000
	TimeZoneBuenosAires_Georgetown TimeZone = -3001
	TimeZoneGreenland TimeZone = -3002
	TimeZoneMidAtlantic TimeZone = -2000
	TimeZoneAzores TimeZone = -1000
	TimeZoneCapeVerdeIs TimeZone = -1001
	TimeZoneUniversalCoordinatedTime TimeZone = 0
	TimeZoneGreenwichMeanTime TimeZone = 1
	TimeZoneCasablanca_Monrovia TimeZone = 2
	TimeZoneAmsterdam_Berlin_Bern_Rome_Stockholm TimeZone = 1000
	TimeZoneBrussels_Madrid_Copenhagen_Paris TimeZone = 1001
	TimeZoneBelgrade_Bratislava_Budapest_Ljubljana_Prague TimeZone = 1002
	TimeZoneSarajevo_Skopje_Warsaw_Zagreb TimeZone = 1003
	TimeZoneWestCentralAfrica TimeZone = 1004
	TimeZoneAthens_Beirut_Istanbul_Minsk TimeZone = 2000
	TimeZoneBucharest TimeZone = 2001
	TimeZoneCairo TimeZone = 2002
	TimeZoneHarare_Pretoria TimeZone = 2003
	TimeZoneHelsinki_Kyiv_Sofia_Talinn_Vilnius TimeZone = 2004
	TimeZoneJerusalem TimeZone = 2005
	TimeZoneMoscow_StPetersburg_Volograd TimeZone = 3000
	TimeZoneKuwait_Riyadh TimeZone = 3001
	TimeZoneBaghdad TimeZone = 3002
	TimeZoneNairobi TimeZone = 3003
	TimeZoneTehran TimeZone = 3300
	TimeZoneAbuDhabi_Muscat TimeZone = 4000
	TimeZoneBaku_Tbilisi_Yerevan TimeZone = 4001
	TimeZoneKabul TimeZone = 4300
	TimeZoneEkaterinburg TimeZone = 5000
	TimeZoneIslamabad_Karachi_Tashkent TimeZone = 5001
	TimeZoneChennai_Kolkata_Mumbai_NewDelhi TimeZone = 5300
	TimeZoneKathmandu TimeZone = 5450
	TimeZoneAlmaty_Novosibirsk TimeZone = 6000
	TimeZoneAstana_Dhaka TimeZone = 6001
	TimeZoneSriJayawardenepura TimeZone = 6002
	TimeZoneRangoon TimeZone = 6300
	TimeZoneBangkok_Hanoi_Jakarta TimeZone = 7000
	TimeZoneKrasnoyarsk TimeZone = 7001
	TimeZoneBeijing_Chongqing_HongKong_Urumqi TimeZone = 8000
	TimeZoneKualaLumpur_Singapore TimeZone = 8001
	TimeZoneTaipei TimeZone = 8002
	TimeZoneIrkutsk_UlaanBataar TimeZone = 8003
	TimeZonePerth TimeZone = 8004
	TimeZoneOsaka_Sapporo_Tokyo TimeZone = 9000
	TimeZoneSeoul TimeZone = 9001
	TimeZoneYakutsk TimeZone = 9002
	TimeZoneAdelaide TimeZone = 9300
	TimeZoneDarwin TimeZone = 9301
	TimeZoneCanberra_Melbourne_Sydney TimeZone = 10000
	TimeZoneGuam_PortMoresby TimeZone = 10001
	TimeZoneBrisbane TimeZone = 10002
	TimeZoneHobart TimeZone = 10003
	TimeZoneVladivostok TimeZone = 10004
	TimeZoneMagadan_SolomonIs_NewCaledonia TimeZone = 11000
	TimeZoneAuckland_Wellington TimeZone = 12000
	TimeZoneFiji_Kamchatka_MarshallIs TimeZone = 12001
	TimeZoneNukualofa_Tonga TimeZone = 13000
)

func (this TimeZone) String() string {
	switch this {
	case TimeZoneInternationalDateLineWest:
		return "TimeZoneInternationalDateLineWest"
	case TimeZoneMidwayIsland_Samoa:
		return "TimeZoneMidwayIsland_Samoa"
	case TimeZoneHawaii:
		return "TimeZoneHawaii"
	case TimeZoneAlaska:
		return "TimeZoneAlaska"
	case TimeZonePacificTime_US_Canada_SanFrancisco_Vancouver:
		return "TimeZonePacificTime_US_Canada_SanFrancisco_Vancouver"
	case TimeZoneArizona:
		return "TimeZoneArizona"
	case TimeZoneMazatlan:
		return "TimeZoneMazatlan"
	case TimeZoneCentralAmerica:
		return "TimeZoneCentralAmerica"
	case TimeZoneCentralTime_US_Canada:
		return "TimeZoneCentralTime_US_Canada"
	case TimeZoneGuadalajara_MexicoCity_Monterrey:
		return "TimeZoneGuadalajara_MexicoCity_Monterrey"
	case TimeZoneSaskatchewan:
		return "TimeZoneSaskatchewan"
	case TimeZoneEasternTime_US_Canada_:
		return "TimeZoneEasternTime_US_Canada_"
	case TimeZoneIndiana_East_:
		return "TimeZoneIndiana_East_"
	case TimeZoneBogota_Lima_Quito:
		return "TimeZoneBogota_Lima_Quito"
	case TimeZoneAtlanticTime_Canada_:
		return "TimeZoneAtlanticTime_Canada_"
	case TimeZoneCaracas_LaPaz:
		return "TimeZoneCaracas_LaPaz"
	case TimeZoneSantiago:
		return "TimeZoneSantiago"
	case TimeZoneNewfoundland:
		return "TimeZoneNewfoundland"
	case TimeZoneBrasilia:
		return "TimeZoneBrasilia"
	case TimeZoneBuenosAires_Georgetown:
		return "TimeZoneBuenosAires_Georgetown"
	case TimeZoneGreenland:
		return "TimeZoneGreenland"
	case TimeZoneMidAtlantic:
		return "TimeZoneMidAtlantic"
	case TimeZoneAzores:
		return "TimeZoneAzores"
	case TimeZoneCapeVerdeIs:
		return "TimeZoneCapeVerdeIs"
	case TimeZoneUniversalCoordinatedTime:
		return "TimeZoneUniversalCoordinatedTime"
	case TimeZoneGreenwichMeanTime:
		return "TimeZoneGreenwichMeanTime"
	case TimeZoneCasablanca_Monrovia:
		return "TimeZoneCasablanca_Monrovia"
	case TimeZoneAmsterdam_Berlin_Bern_Rome_Stockholm:
		return "TimeZoneAmsterdam_Berlin_Bern_Rome_Stockholm"
	case TimeZoneBrussels_Madrid_Copenhagen_Paris:
		return "TimeZoneBrussels_Madrid_Copenhagen_Paris"
	case TimeZoneBelgrade_Bratislava_Budapest_Ljubljana_Prague:
		return "TimeZoneBelgrade_Bratislava_Budapest_Ljubljana_Prague"
	case TimeZoneSarajevo_Skopje_Warsaw_Zagreb:
		return "TimeZoneSarajevo_Skopje_Warsaw_Zagreb"
	case TimeZoneWestCentralAfrica:
		return "TimeZoneWestCentralAfrica"
	case TimeZoneAthens_Beirut_Istanbul_Minsk:
		return "TimeZoneAthens_Beirut_Istanbul_Minsk"
	case TimeZoneBucharest:
		return "TimeZoneBucharest"
	case TimeZoneCairo:
		return "TimeZoneCairo"
	case TimeZoneHarare_Pretoria:
		return "TimeZoneHarare_Pretoria"
	case TimeZoneHelsinki_Kyiv_Sofia_Talinn_Vilnius:
		return "TimeZoneHelsinki_Kyiv_Sofia_Talinn_Vilnius"
	case TimeZoneJerusalem:
		return "TimeZoneJerusalem"
	case TimeZoneMoscow_StPetersburg_Volograd:
		return "TimeZoneMoscow_StPetersburg_Volograd"
	case TimeZoneKuwait_Riyadh:
		return "TimeZoneKuwait_Riyadh"
	case TimeZoneBaghdad:
		return "TimeZoneBaghdad"
	case TimeZoneNairobi:
		return "TimeZoneNairobi"
	case TimeZoneTehran:
		return "TimeZoneTehran"
	case TimeZoneAbuDhabi_Muscat:
		return "TimeZoneAbuDhabi_Muscat"
	case TimeZoneBaku_Tbilisi_Yerevan:
		return "TimeZoneBaku_Tbilisi_Yerevan"
	case TimeZoneKabul:
		return "TimeZoneKabul"
	case TimeZoneEkaterinburg:
		return "TimeZoneEkaterinburg"
	case TimeZoneIslamabad_Karachi_Tashkent:
		return "TimeZoneIslamabad_Karachi_Tashkent"
	case TimeZoneChennai_Kolkata_Mumbai_NewDelhi:
		return "TimeZoneChennai_Kolkata_Mumbai_NewDelhi"
	case TimeZoneKathmandu:
		return "TimeZoneKathmandu"
	case TimeZoneAlmaty_Novosibirsk:
		return "TimeZoneAlmaty_Novosibirsk"
	case TimeZoneAstana_Dhaka:
		return "TimeZoneAstana_Dhaka"
	case TimeZoneSriJayawardenepura:
		return "TimeZoneSriJayawardenepura"
	case TimeZoneRangoon:
		return "TimeZoneRangoon"
	case TimeZoneBangkok_Hanoi_Jakarta:
		return "TimeZoneBangkok_Hanoi_Jakarta"
	case TimeZoneKrasnoyarsk:
		return "TimeZoneKrasnoyarsk"
	case TimeZoneBeijing_Chongqing_HongKong_Urumqi:
		return "TimeZoneBeijing_Chongqing_HongKong_Urumqi"
	case TimeZoneKualaLumpur_Singapore:
		return "TimeZoneKualaLumpur_Singapore"
	case TimeZoneTaipei:
		return "TimeZoneTaipei"
	case TimeZoneIrkutsk_UlaanBataar:
		return "TimeZoneIrkutsk_UlaanBataar"
	case TimeZonePerth:
		return "TimeZonePerth"
	case TimeZoneOsaka_Sapporo_Tokyo:
		return "TimeZoneOsaka_Sapporo_Tokyo"
	case TimeZoneSeoul:
		return "TimeZoneSeoul"
	case TimeZoneYakutsk:
		return "TimeZoneYakutsk"
	case TimeZoneAdelaide:
		return "TimeZoneAdelaide"
	case TimeZoneDarwin:
		return "TimeZoneDarwin"
	case TimeZoneCanberra_Melbourne_Sydney:
		return "TimeZoneCanberra_Melbourne_Sydney"
	case TimeZoneGuam_PortMoresby:
		return "TimeZoneGuam_PortMoresby"
	case TimeZoneBrisbane:
		return "TimeZoneBrisbane"
	case TimeZoneHobart:
		return "TimeZoneHobart"
	case TimeZoneVladivostok:
		return "TimeZoneVladivostok"
	case TimeZoneMagadan_SolomonIs_NewCaledonia:
		return "TimeZoneMagadan_SolomonIs_NewCaledonia"
	case TimeZoneAuckland_Wellington:
		return "TimeZoneAuckland_Wellington"
	case TimeZoneFiji_Kamchatka_MarshallIs:
		return "TimeZoneFiji_Kamchatka_MarshallIs"
	case TimeZoneNukualofa_Tonga:
		return "TimeZoneNukualofa_Tonga"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type UnderlayFrameMode int16

const (
	UnderlayFrameModeNone UnderlayFrameMode = iota
	UnderlayFrameModeDisplayAndPlot
	UnderlayFrameModeDisplayNoPlot
)

func (this UnderlayFrameMode) String() string {
	switch this {
	case UnderlayFrameModeNone:
		return "UnderlayFrameModeNone"
	case UnderlayFrameModeDisplayAndPlot:
		return "UnderlayFrameModeDisplayAndPlot"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type UnitFormat int16

const (
	UnitFormat_ UnitFormat = iota
	UnitFormatScientific UnitFormat = iota
	UnitFormatDecimal
	UnitFormatEngineering
	UnitFormatArchitecturalStacked
	UnitFormatFractionalStacked
	UnitFormatArchitectural
	UnitFormatFractional
)

func (this UnitFormat) String() string {
	switch this {
	case UnitFormat_:
		return "UnitFormat_"
	case UnitFormatDecimal:
		return "UnitFormatDecimal"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type Units int16

const (
	UnitsUnitless Units = iota
	UnitsInches
	UnitsFeet
	UnitsMiles
	UnitsMillimeters
	UnitsCentimeters
	UnitsMeters
	UnitsKilometers
	UnitsMicroinches
	UnitsMils
	UnitsYards
	UnitsAngstroms
	UnitsNanometers
	UnitsMicrons
	UnitsDecimeters
	UnitsDecameters
	UnitsHectometers
	UnitsGigameters
	UnitsAstronomicalUnits
	UnitsLightYears
	UnitsParsecs
	UnitsUSSurveyFeet
	UnitsUSSurveyInch
	UnitsUSSurveyYard
	UnitsUSSurveyMile
)

func (this Units) String() string {
	switch this {
	case UnitsUnitless:
		return "UnitsUnitless"
	case UnitsInches:
		return "UnitsInches"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type UnitZeroSuppression int16

const (
	UnitZeroSuppressionSuppressZeroFeetAndZeroInches UnitZeroSuppression = iota
	UnitZeroSuppressionIncludeZeroFeetAndZeroInches
	UnitZeroSuppressionIncludeZeroFeetAndSuppressZeroInches
	UnitZeroSuppressionIncludeZeroInchesAndSuppressZeroFeet
)

func (this UnitZeroSuppression) String() string {
	switch this {
	case UnitZeroSuppressionSuppressZeroFeetAndZeroInches:
		return "UnitZeroSuppressionSuppressZeroFeetAndZeroInches"
	case UnitZeroSuppressionIncludeZeroFeetAndZeroInches:
		return "UnitZeroSuppressionIncludeZeroFeetAndZeroInches"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type Version int16

const (
	VersionR2010 Version = iota
)

func (this Version) String() string {
	switch this {
	case VersionR2010:
		return "VersionR2010"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type VerticalTextJustification int16

const (
	VerticalTextJustificationBaseline VerticalTextJustification = iota
	VerticalTextJustificationBottom
	VerticalTextJustificationMiddle
	VerticalTextJustificationTop
)

func (this VerticalTextJustification) String() string {
	switch this {
	case VerticalTextJustificationBaseline:
		return "VerticalTextJustificationBaseline"
	case VerticalTextJustificationBottom:
		return "VerticalTextJustificationBottom"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type ViewRenderMode int16

const (
	ViewRenderModeClassic2D ViewRenderMode = iota
	ViewRenderModeWireframe
	ViewRenderModeHiddenLine
	ViewRenderModeFlatShaded
	ViewRenderModeGouraudShaded
	ViewRenderModeFlatShadedWithWireframe
	ViewRenderModeGouraudShadedWithWireframe
)

func (this ViewRenderMode) String() string {
	switch this {
	case ViewRenderModeClassic2D:
		return "ViewRenderModeClassic2D"
	case ViewRenderModeWireframe:
		return "ViewRenderModeWireframe"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

type XrefClippingBoundaryVisibility int16

const (
	XrefClippingBoundaryVisibilityNotDisplayedNotPlotted XrefClippingBoundaryVisibility = iota
	XrefClippingBoundaryVisibilityDisplayedAndPlotted
	XrefClippingBoundaryVisibilityDisplayedNotPlotted
)

func (this XrefClippingBoundaryVisibility) String() string {
	switch this {
	case XrefClippingBoundaryVisibilityNotDisplayedNotPlotted:
		return "XrefClippingBoundaryVisibilityNotDisplayedNotPlotted"
	case XrefClippingBoundaryVisibilityDisplayedAndPlotted:
		return "XrefClippingBoundaryVisibilityDisplayedAndPlotted"
	default:
		return fmt.Sprintf("%v", (this))
	}
}

