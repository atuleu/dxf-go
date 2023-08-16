// Code generated at build time; DO NOT EDIT.

package dxf

import (
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
)

// Header contains the values common to an AutoCAD DXF drawing.
type Header struct {
	// The $ACADVER header variable.  The AutoCAD drawing database version number.
	Version AcadVersion
	// The $ACADMAINTVER header variable.  Maintenance version number (should be ignored).  Minimum AutoCAD version R14.
	MaintenanceVersion int16
	// The $DWGCODEPAGE header variable.  Drawing code page; set to the system code page when a new drawing is created, but not otherwise maintained by AutoCAD.  Minimum AutoCAD version R13.
	DrawingCodePage string
	// The $LASTSAVEDBY header variable.  Name of the last user to modify the file.  Minimum AutoCAD version R2004.
	LastSavedBy string
	// The $REQUIREDVERSIONS header variable.  Unknown.  Minimum AutoCAD version R2013.
	RequiredVersions int64
	// The $INSBASE header variable.  Insertion base set by BASE command (in WCS).
	InsertionBase Point
	// The $EXTMIN header variable.  X, Y, and Z drawing extents lower-left corner (in WCS).
	MinimumDrawingExtents Point
	// The $EXTMAX header variable.  X, Y, and Z drawing extents upper-right corner (in WCS).
	MaximumDrawingExtents Point
	// The $LIMMIN header variable.  XY drawing limits lower-left corner (in WCS).
	MinimumDrawingLimits Point
	// The $LIMMAX header variable.  XY drawing limits upper-right corner (in WCS).
	MaximumDrawingLimits Point
	// The $ORTHOMODE header variable.  Ortho mode on.
	DrawOrthoganalLines bool
	// The $REGENMODE header variable.  REGENAUTO mode on.
	UseRegenMode bool
	// The $FILLMODE header variable.  Fill mode on.
	FillModeOn bool
	// The $QTEXTMODE header variable.  Quick text mode on.
	UseQuickTextMode bool
	// The $MIRRTEXT header variable.  Mirror text.
	MirrorText bool
	// The $DRAGMODE header variable.  Controls the way dragged objects are displayed.  Maximum AutoCAD version R14.
	DragMode DragMode
	// The $LTSCALE header variable.  Global line type scale.
	LineTypeScale float64
	// The $OSMODE header variable.  Running object snap modes.  Maximum AutoCAD version R14.
	ObjectSnapFlags int
	// The $ATTMODE header variable.  Attribute visibility.
	AttributeVisibility AttributeVisibility
	// The $TEXTSIZE header variable.  Default text height.
	DefaultTextHeight float64
	// The $TRACEWID header variable.  Default trace width.
	TraceWidth float64
	// The $TEXTSTYLE header variable.  Current text style name.
	TextStyle string
	// The $CLAYER header variable.  Current layer name.
	CurrentLayer string
	// The $CELTYPE header variable.  Entity line type name, or BYBLOCK or BYLAYER.
	CurrentEntityLineType string
	// The $CECOLOR header variable.  Current entity color.
	CurrentEntityColor Color
	// The $CELTSCALE header variable.  Current entity line type scale.  Minimum AutoCAD version R13.
	CurrentEntityLineTypeScale float64
	// The $DELOBJ header variable.  Controls object deletion.  Minimum AutoCAD version R13.  Maximum AutoCAD version R14.
	RetainDeletedObjects bool
	// The $DISPSILH header variable.  Controls the display of silhouette curves of body objects in wireframe mode.  Minimum AutoCAD version R13.
	DisplaySilhouetteCurvesInWireframeMode bool
	// The $DRAGVS header variable.  Hard-pointer ID to visual style when creating 3D solid primitives.  Minimum AutoCAD version R2007.
	SolidVisualStylePointer Handle
	// The $DIMSCALE header variable.  Overall dimensioning scale factor.
	DimensioningScaleFactor float64
	// The $DIMASZ header variable.  Dimensioning arrow size.
	DimensioningArrowSize float64
	// The $DIMEXO header variable.  Extension line offset.
	DimensionExtensionLineOffset float64
	// The $DIMDLI header variable.  Dimension line increment.
	DimensionLineIncrement float64
	// The $DIMRND header variable.  Rounding value for dimension distances.
	DimensionDistanceRoundingValue float64
	// The $DIMDLE header variable.  Dimension line extension.
	DimensionLineExtension float64
	// The $DIMEXE header variable.  Extension line extension.
	DimensionExtensionLineExtension float64
	// The $DIMTP header variable.  Plus tolerance.
	DimensionPlusTolerance float64
	// The $DIMTM header variable.  Minus tolerance.
	DimensionMinusTolerance float64
	// The $DIMTXT header variable.  Dimensioning text height.
	DimensioningTextHeight float64
	// The $DIMCEN header variable.  Size of center mark/lines.
	CenterMarkSize float64
	// The $DIMTSZ header variable.  Dimensioning tick size.
	DimensioningTickSize float64
	// The $DIMTOL header variable.  Dimension tolerances generated.
	GenerateDimensionTolerances bool
	// The $DIMLIM header variable.  Dimension limits generated.
	GenerateDimensionLimits bool
	// The $DIMTIH header variable.  Text inside horizontal.
	DimensionTextInsideHorizontal bool
	// The $DIMTOH header variable.  Text outside horizontal.
	DimensionTextOutsideHorizontal bool
	// The $DIMSE1 header variable.  Suppression of first extension line.  Minimum AutoCAD version R12.
	SuppressFirstDimensionExtensionLine bool
	// The $DIMSE2 header variable.  Suppression of second extension line.  Minimum AutoCAD version R12.
	SuppressSecondDimensionExtensionLine bool
	// The $DIMTAD header variable.  Text above dimension line.
	TextAboveDimensionLine bool
	// The $DIMZIN header variable.  Controls suppression of zeros for primary unit values.
	DimensionUnitZeroSuppression UnitZeroSuppression
	// The $DIMBLK header variable.  Arrow block name.
	ArrowBlockName string
	// The $DIMASO header variable.  Controls associative dimensioning.
	CreateAssociativeDimensioning bool
	// The $DIMSHO header variable.  Recompute dimensions while dragging.
	RecomputeDimensionsWhileDragging bool
	// The $DIMPOST header variable.  General dimensioning suffix.
	DimensioningSuffix string
	// The $DIMAPOST header variable.  Alternate dimensioning suffix.
	AlternateDimensioningSuffix string
	// The $DIMALT header variable.  Alternate unit dimensioning performed.
	UseAlternateDimensioning bool
	// The $DIMALTD header variable.  Alternate unit decimal places.
	AlternateDimensioningDecimalPlaces int16
	// The $DIMALTF header variable.  Alternate unit scale factor.
	AlternateDimensioningScaleFactor float64
	// The $DIMLFAC header variable.  Linear measurements scale factor.
	DimensionLinearMeasurementsScaleFactor float64
	// The $DIMTOFL header variable.  If text is outside extensions, force line extensions between extensions.
	ForceDimensionLineExtensionsOutsideIfTextIs bool
	// The $DIMTVP header variable.  Text vertical position.
	DimensionVerticalTextPosition float64
	// The $DIMTIX header variable.  Force text inside extensions.
	ForceDimensionTextInsideExtensions bool
	// The $DIMSOXD header variable.  Suppress outside-extensions dimension lines.
	SuppressOutsideExtensionDimensionLines bool
	// The $DIMSAH header variable.  Use separate arrow blocks.
	UseSeparateArrowBlocksForDimensions bool
	// The $DIMBLK1 header variable.  First arrow block name.
	FirstArrowBlockName string
	// The $DIMBLK2 header variable.  Second arrow block name.
	SecondArrowBlockName string
	// The $DIMSTYLE header variable.  Dimension style name.
	DimensionStyleName string
	// The $DIMCLRD header variable.  Dimension line color.  Minimum AutoCAD version R11.
	DimensionLineColor Color
	// The $DIMCLRE header variable.  Dimension extension line color.  Minimum AutoCAD version R11.
	DimensionExtensionLineColor Color
	// The $DIMCLRT header variable.  Dimension text color.  Minimum AutoCAD version R11.
	DimensionTextColor Color
	// The $DIMTFAC header variable.  Dimension tolerance display factor.  Minimum AutoCAD version R12.
	DimensionToleranceDisplayScaleFactor float64
	// The $DIMGAP header variable.  Dimension line gap.  Minimum AutoCAD version R11.
	DimensionLineGap float64
	// The $DIMJUST header variable.  Horizontal dimension text position.  Minimum AutoCAD version R13.
	DimensionTextJustification DimensionTextJustification
	// The $DIMTOLJ header variable.  Vertical justification for tolerance values.  Minimum AutoCAD version R13.
	DimensionToleranceVerticalJustification Justification
	// The $DIMTZIN header variable.  Controls suppression of zeros for tolerance values.  Minimum AutoCAD version R13.
	DimensionToleranceZeroSuppression UnitZeroSuppression
	// The $DIMALTZ header variable.  Controls suppression of zeros for alternate unit dimension values.  Minimum AutoCAD version R13.
	AlternateDimensioningZeroSupression UnitZeroSuppression
	// The $DIMALTTZ header variable.  Controls suppression of zeros for alternate tolerance values.  Minimum AutoCAD version R13.
	AlternateDimensioningToleranceZeroSupression UnitZeroSuppression
	// The $DIMFIT header variable.  Placement of text and arrowheads.  Minimum AutoCAD version R13.  Maximum AutoCAD version R14.
	DimensionTextAndArrowPlacement DimensionFit
	// The $DIMUPT header variable.  Cursor functionality for user-positioned text.  Minimum AutoCAD version R13.
	DimensionCursorControlsTextPosition bool
	// The $DIMUNIT header variable.  Units format for all dimension style family members except angular.  Minimum AutoCAD version R13.  Maximum AutoCAD version R14.
	DimensionUnitFormat UnitFormat
	// The $DIMDEC header variable.  Number of decimal places for the tolerance values of a primary units dimension.  Minimum AutoCAD version R13.
	DimensionUnitToleranceDecimalPlaces int16
	// The $DIMTDEC header variable.  Number of decimal places to display the tolerance values.  Minimum AutoCAD version R13.
	DimensionToleranceDecimalPlaces int16
	// The $DIMALTU header variable.  Units format for alternate units of all dimension style family members except angular.  Minimum AutoCAD version R13.
	AlternateDimensioningUnits UnitFormat
	// The $DIMALTTD header variable.  Number of decimal places for tolerance values of an alternate units dimension.  Minimum AutoCAD version R13.
	AlternateDimensioningToleranceDecimalPlaces int16
	// The $DIMTXSTY header variable.  Dimension text style.  Minimum AutoCAD version R13.
	DimensionTextStyle string
	// The $DIMAUNIT header variable.  Angle format for angular dimensions.  Minimum AutoCAD version R13.
	DimensioningAngleFormat AngleFormat
	// The $DIMADEC header variable.  Number of precision places displayed in angular dimensions.  Minimum AutoCAD version R2000.
	AngularDimensionPrecision int16
	// The $DIMALTRND header variable.  Determines rounding of alternate units.  Minimum AutoCAD version R2000.
	AlternateDimensioningUnitRounding float64
	// The $DIMAZIN header variable.  Controls suppression of zeros for angular dimensions.  Minimum AutoCAD version R2000.
	DimensionAngleZeroSuppression UnitZeroSuppression
	// The $DIMDSEP header variable.  Single-character decimal separator used when creating dimensions whose unit format is decimal.  Minimum AutoCAD version R2000.
	DimensionDecimalSeparatorRune rune
	// The $DIMFRAC header variable.  Sets the fraction format when DIMLUNIT is set to Architectural or Fractional.  Minimum AutoCAD version R2000.
	DimensionTextHeightScaleFactor DimensionFractionFormat
	// The $DIMLDRBLK header variable.  Arrow block name for leaders.  Minimum AutoCAD version R2000.
	DimensionLeaderBlockName string
	// The $DIMLUNIT header variable.  Sets units for all dimension types except angular.  Minimum AutoCAD version R2000.
	DimensionNonAngularUnits NonAngularUnits
	// The $DIMLWD header variable.  Dimension line lineweight.  Minimum AutoCAD version R2000.
	DimensionLineWeight LineWeight
	// The $DIMLWE header variable.  Extension line lineweight.  Minimum AutoCAD version R2000.
	DimensionExtensionLineWeight LineWeight
	// The $DIMTMOVE header variable.  Dimension text movement rules.  Minimum AutoCAD version R2000.
	DimensionTextMovementRule DimensionTextMovementRule
	// The $DIMFXL header variable.  Sets the total length of the extension lines starting from the dimension line toward the dimension origin.  Minimum AutoCAD version R2007.
	DimensionLineFixedLength float64
	// The $DIMFXLON header variable.  Controls whether extension lines are set to a fixed length.  Minimum AutoCAD version R2007.
	DimensionLineFixedLengthOn bool
	// The $DIMJOGANG header variable.  Determines the angle of the transverse segment of the dimension line in a jogged radius dimension.  Minimum AutoCAD version R2007.
	DimensionTransverseSegmentAngleInJoggedRadius float64
	// The $DIMTFILL header variable.  Controls the background of dimension text.  Minimum AutoCAD version R2007.
	DimensionTextBackgroundColorMode DimensionTextBackgroundColorMode
	// The $DIMTFILLCLR header variable.  Sets the color for the text background in dimensions.  Minimum AutoCAD version R2007.
	DimensionTextBackgroundCustomColor Color
	// The $DIMARCSYM header variable.  Controls the display of the arc symbol in an arc length dimension.  Minimum AutoCAD version R2007.
	DimensionArcSymbolDisplayMode DimensionArcSymbolDisplayMode
	// The $DIMLTYPE header variable.  Sets the line type of the dimension line.  Minimum AutoCAD version R2007.
	DimensionLineType string
	// The $DIMLTEX1 header variable.  Sets the line type of the first extension line.  Minimum AutoCAD version R2007.
	DimensionFirstExtensionLineType string
	// The $DIMLTEX2 header variable.  Sets the line type of the second extension line.  Minimum AutoCAD version R2007.
	DimensionSecondExtensionLineType string
	// The $DIMTXTDIRECTION header variable.  Specifies the reading direction of the dimension text.  Minimum AutoCAD version R2010.
	DimensionTextDirection TextDirection
	// The $LUNITS header variable.  Units format for coordinates and distances.
	UnitFormat UnitFormat
	// The $LUPREC header variable.  Units precision for coordinates and distances.
	UnitPrecision int16
	// The $SKETCHINC header variable.  Sketch record increment.
	SketchRecordIncrement float64
	// The $FILLETRAD header variable.  Fillet radius.
	FilletRadius float64
	// The $AUNITS header variable.  Units format for angles.
	AngleUnitFormat AngleFormat
	// The $AUPREC header variable.  Units precision for angles.
	AngleUnitPrecision int16
	// The $MENU header variable.  Name of menu file.
	FileName string
	// The $ELEVATION header variable.  Current elevation set by ELEV command.
	Elevation float64
	// The $PELEVATION header variable.  Current paper space elevation.  Minimum AutoCAD version R11.
	PaperspaceElevation float64
	// The $THICKNESS header variable.  Current thickness set by ELEV command.
	Thickness float64
	// The $LIMCHECK header variable.  Limits checking.
	UseLimitsChecking bool
	// The $BLIPMODE header variable.  Display blips for click locations.  Maximum AutoCAD version R14.
	BlipMode bool
	// The $CHAMFERA header variable.  First chamfer distance.
	FirstChamferDistance float64
	// The $CHAMFERB header variable.  Second chamfer distance.
	SecondChamferDistance float64
	// The $CHAMFERC header variable.  Chamfer length.  Minimum AutoCAD version R14.
	ChamferLength float64
	// The $CHAMFERD header variable.  Chamfer angle.  Minimum AutoCAD version R14.
	ChamferAngle float64
	// The $SKPOLY header variable.  Controls polyline sketch mode.
	PolylineSketchMode PolySketchMode
	// The $TDCREATE header variable.  Local date/time of drawing creation.
	CreationDate time.Time
	// The $TDUCREATE header variable.  Universal date/time the drawing was created.  Minimum AutoCAD version R2000.
	CreationDateUniversal time.Time
	// The $TDUPDATE header variable.  Local date/time of last drawing update.
	UpdateDate time.Time
	// The $TDUUPDATE header variable.  Universal date/time of the last update/save.  Minimum AutoCAD version R2000.
	UpdateDateUniversal time.Time
	// The $TDINDWG header variable.  Cumulative editing time for this drawing.
	TimeInDrawing time.Duration
	// The $TDUSRTIMER header variable.  User-elapsed timer.
	UserElapsedTimer time.Duration
	// The $USRTIMER header variable.  User timer on.
	UserTimerOn bool
	// The $ANGBASE header variable.  Angle 0 direction.
	AngleZeroDirection float64
	// The $ANGDIR header variable.  Angle directions.
	AngleDirection AngleDirection
	// The $PDMODE header variable.  Point display mode.
	PointDisplayMode int
	// The $PDSIZE header variable.  Point display size.
	PointDisplaySize float64
	// The $PLINEWID header variable.  Default polyline width.
	DefaultPolylineWidth float64
	// The $COORDS header variable.  Controls the display of coordinates.  Maximum AutoCAD version R14.
	CoordinateDisplay CoordinateDisplay
	// The $SPLFRAME header variable.  Controls the display of helixes and smoothed mesh objects.  Maximum AutoCAD version R2013.
	DisplaySplinePolygonControl bool
	// The $SPLINETYPE header variable.  Spline curve type for PEDIT Spline.
	PEditSplineCurveType PolylineCurvedAndSmoothSurfaceType
	// The $SPLINESEGS header variable.  Number of line segments per spline hatch.
	LineSegmentsPerSplinePatch int16
	// The $ATTDIA header variable.  Controls whether the INSERT command uses a dialog box for attribute value entry.  Maximum AutoCAD version R14.
	ShowAttributeEntryDialogs bool
	// The $ATTREQ header variable.  Controls whether INSERT uses default attribute settings during insertion of blocks.  Maximum AutoCAD version R14.
	PromptForAttributeOnInsert bool
	// The $HANDLING header variable.  Handles available.  Maximum AutoCAD version R12.
	HandlesEnabled bool
	// The $HANDSEED header variable.  Next available handle.
	NextAvailableHandle Handle
	// The $SURFTAB1 header variable.  Number of mesh tabulations in first direction.
	MeshTabulationsInFirstDirection int16
	// The $SURFTAB2 header variable.  Number of mesh tabulations in second direction.
	MeshTabulationsInSecondDirection int16
	// The $SURFTYPE header variable.  Surface type for PEDIT Smooth.
	PEditSmoothSurfaceType PolylineCurvedAndSmoothSurfaceType
	// The $SURFU header variable.  Surface density (for PEDIT Smooth) in M direction.
	PEditSmoothMDensith int16
	// The $SURFV header variable.  Surface density (for PEDIT Smooth) in N direction.
	PEditSmoothNDensith int16
	// The $UCSBASE header variable.  Name of the UCS that defines the origin and orientation of orthographic UCS settings.  Minimum AutoCAD version R2000.
	UCSDefinitionName string
	// The $UCSNAME header variable.  Name of current UCS.
	UCSName string
	// The $UCSORG header variable.  Origin of current UCS (in WCS).
	UCSOrigin Point
	// The $UCSXDIR header variable.  Direction of the current UCS X axis (in WCS).
	UCSXAxis Vector
	// The $UCSYDIR header variable.  Direction of the current UCS Y axis (in WCS).
	UCSYAxis Vector
	// The $UCSORTHOREF header variable.  If model space UCS is orthographic (UCSORTHOVIEW not equal to 0), this is the name of the UCS that the orthographic UCS is relative to. If blank, UCS is relative to WORLD.  Minimum AutoCAD version R2000.
	OrthoUCSReference string
	// The $UCSORTHOVIEW header variable.  Orthographic view type of model space UCS.  Minimum AutoCAD version R2000.
	OrthgraphicViewType OrthographicViewType
	// The $UCSORGTOP header variable.  Point which becomes the new UCS origin after changing model space UCS to TOP when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginTop Point
	// The $UCSORGBOTTOM header variable.  Point which becomes the new UCS origin after changing model space UCS to BOTTOM when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginBottom Point
	// The $UCSORGLEFT header variable.  Point which becomes the new UCS origin after changing model space UCS to LEFT when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginLeft Point
	// The $UCSORGRIGHT header variable.  Point which becomes the new UCS origin after changing model space UCS to RIGHT when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginRight Point
	// The $UCSORGFRONT header variable.  Point which becomes the new UCS origin after changing model space UCS to FRONT when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginFront Point
	// The $UCSORGBACK header variable.  Point which becomes the new UCS origin after changing model space UCS to BACK when UCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	UCSOriginBack Point
	// The $PUCSBASE header variable.  Name of the UCS that defines the origin and orientation of orthographics UCS settings (paper space only).  Minimum AutoCAD version R2000.
	PaperspaceUCSDefinitionName string
	// The $PUCSNAME header variable.  Current paper space UCS name.  Minimum AutoCAD version R11.
	PaperspaceUCSName string
	// The $PUCSORG header variable.  Current paper space UCS origin.  Minimum AutoCAD version R11.
	PaperspaceUCSOrigin Point
	// The $PUCSXDIR header variable.  Current paper space UCS X axis.  Minimum AutoCAD version R11.
	PaperspaceXAxis Vector
	// The $PUCSYDIR header variable.  Current paper space UCS Y axis.  Minimum AutoCAD version R11.
	PaperspaceYAxis Vector
	// The $PUCSORTHOREF header variable.  If paper space UCS is orthographic (PUCSORTHOVIEW not equal to 0), this is the name of the UCS that the orthographic UCS is relative to. If blank, UCS is relative to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceOrthoUCSReference string
	// The $PUCSORTHOVIEW header variable.  Orthographic view type of paper space UCS.  Minimum AutoCAD version R2000.
	PaperspaceOrthographicViewType OrthographicViewType
	// The $PUCSORGTOP header variable.  Point which becomes the new UCS origin after changing paper space UCS to TOP when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginTop Point
	// The $PUCSORGBOTTOM header variable.  Point which becomes the new UCS origin after changing paper space UCS to BOTTOM when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginBottom Point
	// The $PUCSORGLEFT header variable.  Point which becomes the new UCS origin after changing paper space UCS to LEFT when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginLeft Point
	// The $PUCSORGRIGHT header variable.  Point which becomes the new UCS origin after changing paper space UCS to RIGHT when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginRight Point
	// The $PUCSORGFRONT header variable.  Point which becomes the new UCS origin after changing paper space UCS to FRONT when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginFront Point
	// The $PUCSORGBACK header variable.  Point which becomes the new UCS origin after changing paper space UCS to BACK when PUCSBASE is set to WORLD.  Minimum AutoCAD version R2000.
	PaperspaceUCSOriginBack Point
	// The $USERI1 header variable.  Integer variable intended for use by third-party developers.
	UserInt1 int16
	// The $USERI2 header variable.  Integer variable intended for use by third-party developers.
	UserInt2 int16
	// The $USERI3 header variable.  Integer variable intended for use by third-party developers.
	UserInt3 int16
	// The $USERI4 header variable.  Integer variable intended for use by third-party developers.
	UserInt4 int16
	// The $USERI5 header variable.  Integer variable intended for use by third-party developers.
	UserInt5 int16
	// The $USERR1 header variable.  Real variable indented for use by third-party developers.
	UserReal1 float64
	// The $USERR2 header variable.  Real variable indented for use by third-party developers.
	UserReal2 float64
	// The $USERR3 header variable.  Real variable indented for use by third-party developers.
	UserReal3 float64
	// The $USERR4 header variable.  Real variable indented for use by third-party developers.
	UserReal4 float64
	// The $USERR5 header variable.  Real variable indented for use by third-party developers.
	UserReal5 float64
	// The $WORLDVIEW header variable.  Set UCS to WCS during DVIEW/VPOINT.
	SetUCSToWCSInDViewOrVPoint bool
	// The $SHADEDGE header variable.  Controls shading of faces.  Minimum AutoCAD version R11.
	EdgeShading ShadeEdgeMode
	// The $SHADEDIF header variable.  Percent ambient/diffuse light; range 1-100.  Minimum AutoCAD version R11.
	PercentAmbientToDiffuse int16
	// The $TILEMODE header variable.  Use previous release compatibility mode.  Minimum AutoCAD version R11.
	PreviousReleaseTileCompatability bool
	// The $MAXACTVP header variable.  Sets the maximum number of viewports to be regenerated.  Minimum AutoCAD version R11.
	MaximumActiveViewports int16
	// The $PINSBASE header variable.  Paper space insertion base point.  Minimum AutoCAD version R14.
	PaperspaceInsertionBase Point
	// The $PLIMCHECK header variable.  Limits checking in paper space.  Minimum AutoCAD version R11.
	LimitCheckingInPaperspace bool
	// The $PEXTMIN header variable.  Minimum X, Y, and Z extents for paper space.  Minimum AutoCAD version R11.
	PaperspaceMinimumDrawingExtents Point
	// The $PEXTMAX header variable.  Maximum X, Y, and Z extents for paper space.  Minimum AutoCAD version R11.
	PaperspaceMaximumDrawingExtents Point
	// The $PLIMMIN header variable.  Minimum X and Y limits in paper space.  Minimum AutoCAD version R11.
	PaperspaceMinimumDrawingLimits Point
	// The $PLIMMAX header variable.  Maximum X and Y limits in paper space.  Minimum AutoCAD version R11.
	PaperspaceMaximumDrawingLimits Point
	// The $UNITMODE header variable.  Display fractions, feet-and-inches, and surveyor's angles in input format.  Minimum AutoCAD version R11.
	DisplayFractionsInInput bool
	// The $VISRETAIN header variable.  Retain xref-dependent visibility settings.  Minimum AutoCAD version R12.
	RetainXRefDependentVisibilitySettings bool
	// The $PLINEGEN header variable.  Governs the generation of line type patterns around the vertices of a 2D polyline.  Minimum AutoCAD version R11.
	IsPolylineContinuousAroundVerticies bool
	// The $PSLTSCALE header variable.  Controls paper space line type scaling.  Minimum AutoCAD version R11.
	ScaleLineTypesInPaperspace bool
	// The $TREEDEPTH header variable.  Specifies the maximum depth of the spatial index.  Minimum AutoCAD version R14.
	SpacialIndexMaxDepth int16
	// The $PICKSTYLE header variable.  Controls the group selection and associative hatch selection.  Minimum AutoCAD version R13.  Maximum AutoCAD version R14.
	PickStyle PickStyle
	// The $CMLSTYLE header variable.  Current multiline style name.  Minimum AutoCAD version R13.  Maximum AutoCAD version R13.
	CurrentMultilineStyle string
	// The $CMLJUST header variable.  Current multiline justification.  Minimum AutoCAD version R13.
	CurrentMultilineJustification Justification
	// The $CMLSCALE header variable.  Current multiline scale.  Minimum AutoCAD version R13.
	CurrentMultilineScale float64
	// The $PROXYGRAPHICS header variable.  Controls the saving of proxy object images.  Minimum AutoCAD version R14.
	SaveProxyGraphics bool
	// The $MEASUREMENT header variable.  Sets drawing units.  Minimum AutoCAD version R14.
	DrawingUnits DrawingUnits
	// The $CELWEIGHT header variable.  Lineweight of new objects.  Minimum AutoCAD version R2000.
	NewObjectLineWeight LineWeight
	// The $ENDCAPS header variable.  Lineweight endcaps setting for new objects.  Minimum AutoCAD version R2000.
	EndCapSetting EndCapSetting
	// The $JOINSTYLE header variable.  Lineweight join setting for new objects.  Minimum AutoCAD version R2000.
	LineweightJointSetting JoinStyle
	// The $LWDISPLAY header variable.  Controls the display of lineweights on the Model or Layout tab.  Minimum AutoCAD version R2000.
	DisplayLinewieghtInModelAndLayoutTab bool
	// The $INSUNITS header variable.  Default drawing units for AutoCAD DesignCenter blocks.  Minimum AutoCAD version R2000.
	DefaultDrawingUnits Units
	// The $HYPERLINKBASE header variable.  Path for all relative hyperlinks in the drawing.  If null, the drawing path is used.  Minimum AutoCAD version R2000.
	HyperlinkBase string
	// The $STYLESHEET header variable.  Path to the stylesheet for the drawing.  Minimum AutoCAD version R2000.
	Stylesheet string
	// The $XEDIT header variable.  Controls whether the current drawing can be edited in-place when being referenced by another drawing.  Minimum AutoCAD version R2000.
	CanUseInPlaceReferenceEditing bool
	// The $CEPSNID header variable.  PlotStyle handle of new objects.  Minimum AutoCAD version R2000.
	NewObjectPlotStyleHandle Handle
	// The $CEPSNTYPE header variable.  Plot style of new objects.  Minimum AutoCAD version R2000.
	NewObjectPlotStyle PlotStyle
	// The $PSTYLEMODE header variable.  Indicates whether the current drawing is in a Color-Dependent or Named Plot Style mode.  Minimum AutoCAD version R2000.
	UsesColorDependentPlotStyleTables bool
	// The $FINGERPRINTGUID header variable.  Set at creation time, uniquely identifies a particular drawing.  Minimum AutoCAD version R2000.
	FingerprintGuid uuid.UUID
	// The $VERSIONGUID header variable.  Uniquely identifies a particular version of a drawing.  Updated when the drawing is modified.  Minimum AutoCAD version R2000.
	VersionGuid uuid.UUID
	// The $EXTNAMES header variable.  Controls symbol table naming.  Minimum AutoCAD version R2000.
	UseACad2000SymbolTableNaming bool
	// The $PSVPSCALE header variable.  View scale factor for new viewports.  Minimum AutoCAD version R2000.
	ViewportViewScaleFactor float64
	// The $OLESTARTUP header variable.  Controls whether the source application of an embedded OLE object loads when plotting.  Minimum AutoCAD version R2000.
	OleStartup bool
	// The $SORTENTS header variable.  Controls the object sorting methods; accessible from the Options dialog box User Preferences tab.  Minimum AutoCAD version R2004.
	ObjectSortingMethodsFlags int
	// The $INDEXCTL header variable.  Controls whether layer and spatial indexes are created and saved in drawing files.  Minimum AutoCAD version R2004.
	LayerAndSpatialIndexSaveMode LayerAndSpatialIndexSaveMode
	// The $HIDETEXT header variable.  Ignore text objects.  Minimum AutoCAD version R2004.
	HideTextObjectsWhenProducintHiddenView bool
	// The $XCLIPFRAME header variable.  Controls the visibility of xref clipping boundaries.  Minimum AutoCAD version R2004.  Maximum AutoCAD version R2007.
	XRefClippingBoundaryVisible XrefClippingBoundaryVisibility
	// The $HALOGAP header variable.  Specifies a gap to be displayed where an object is hidden by another object; the value is specified as a percent of one unit and is independent of the zoom level.  A haloed line is shortened at the point where it is hidden when HIDE or the Hidden option of SHADEMODE is used.  Minimum AutoCAD version R2004.
	HaloGapPercent float64
	// The $OBSCOLOR header variable.  Specifies the color of obscured lines.  An obscured line is a hidden line made visible by changing its color and line type and is visible only when the HIDE or SHADEMODE command is used.  The OBSCUREDCOLOR setting is visible only if the OBSCUREDLTYPE is turned ON by setting it to a value other than 0.  Minimum AutoCAD version R2004.
	ObscuredLineColor Color
	// The $OBSLTYPE header variable.  Specifies the line type of obscured lines.  Obscured line types are independent of zoom level, unlike regular AutoCAD line types.  Value 0 turns off display of obscured lines and is the default.  Minimum AutoCAD version R2004.
	ObscuredLineTypeStyle LineTypeStyle
	// The $INTERSECTIONDISPLAY header variable.  Specifies the display of intersection polylines.  Minimum AutoCAD version R2004.
	DisplayIntersectionPolylines bool
	// The $INTERSECTIONCOLOR header variable.  Specifies the entity color of intersection polylines.  Minimum AutoCAD version R2004.
	IntersectionPolylineColor Color
	// The $DIMASSOC header variable.  Controls the associativity of dimension objects.  Minimum AutoCAD version R2004.
	DimensionObjectAssociativity DimensionAssociativity
	// The $PROJECTNAME header variable.  Assigns a project name to the current drawing.  Used when an external reference or image is not found on its original path.  The project name points to a section in the registry that can contain one or more search paths for each project name defined.  Project names and their search directories are created from the Files tab of the Options dialog box.  Minimum AutoCAD version R2004.
	ProjectName string
	// The $CAMERADISPLAY header variable.  Turns the display of camera objects on or off.  Minimum AutoCAD version R2007.
	UseCameraDisplay bool
	// The $LENSLENGTH header variable.  Stores the length of the lens in millimeters used in perspective viewing.  Minimum AutoCAD version R2007.
	LensLength float64
	// The $CAMERAHEIGHT header variable.  Specifies the default height for new camera objects.  Minimum AutoCAD version R2007.
	CameraHeight float64
	// The $STEPSPERSEC header variable.  Specifies the number of steps taken per second when you are in walk or fly mode.  Minimum AutoCAD version R2007.
	StepsPerSecondInWalkOrFlyMode float64
	// The $STEPSIZE header variable.  Specifies the size of each step when in walk or fly mode, in drawing units.  Minimum AutoCAD version R2007.
	StepSizeInWalkOrFlyMode float64
	// The $3DDWFPREC header variable.  Controls the precision of 3D DWF or 3D DWFx publishing.  Minimum AutoCAD version R2007.
	Dwf3DPrecision Dwf3DPrecision
	// The $PSOLWIDTH header variable.  Controls the default width for a swept solid object created with the POLYSOLID command.  Minimum AutoCAD version R2007.
	LastPolySolidWidth float64
	// The $PSOLHEIGHT header variable.  Controls the default height for a swept solid object created with the POLYSOLID command.  Minimum AutoCAD version R2007.
	LastPolySolidHeight float64
	// The $LOFTANG1 header variable.  Sets the draft angle through the first cross section in a loft operation.  Minimum AutoCAD version R2007.
	LoftOperationFirstDraftAngle float64
	// The $LOFTANG2 header variable.  Sets the draft angle through the second cross section in a loft operation.  Minimum AutoCAD version R2007.
	LoftOperationSecondDraftAngle float64
	// The $LOFTMAG1 header variable.  Sets the magnitude of the draft angle through the first cross section in a loft operation.  Minimum AutoCAD version R2007.
	LoftOperationFirstMagnitude float64
	// The $LOFTMAG2 header variable.  Sets the magnitude of the draft angle through the second cross section in a loft operation.  Minimum AutoCAD version R2007.
	LoftOperationSecondMagnitude float64
	// The $LOFTPARAM header variable.  Controls the shape of lofted solids and surfaces.  Minimum AutoCAD version R2007.
	LoftFlags int
	// The $LOFTNORMALS header variable.  Controls the normals of a lofted object where it passes through cross sections.  Minimum AutoCAD version R2007.
	LoftedObjectNormalMode LoftedObjectNormalMode
	// The $LATITUDE header variable.  The latitude of the geographic location assigned to the drawing.  Minimum AutoCAD version R2007.
	Latitude float64
	// The $LONGITUDE header variable.  The longitude of the geographic location assigned to the drawing.  Minimum AutoCAD version R2007.
	Longitude float64
	// The $NORTHDIRECTION header variable.  Specifies the angle between the Y axis of WCS and the grid north.  Minimum AutoCAD version R2007.
	AngleBetweenYAxisAndNorth float64
	// The $TIMEZONE header variable.  Sets the time zone for the sun in the drawing.  Minimum AutoCAD version R2007.
	TimeZone TimeZone
	// The $LIGHTGLYPHDISPLAY header variable.  Turns on and off the display of light glyphs.  Minimum AutoCAD version R2007.
	UseLightGlyphDisplay bool
	// The $TILEMODELIGHTSYNCH header variable.  Unknown.  Minimum AutoCAD version R2007.
	UseTileModeLightSync bool
	// The $CMATERIAL header variable.  Sets the material of new objects.  Minimum AutoCAD version R2007.
	CurrentMaterialHandle Handle
	// The $SOLIDHIST header variable.  Controls whether new composite solids retain a history of their original components.  Minimum AutoCAD version R2007.
	NewSolidsContainHistory bool
	// The $SHOWHIST header variable.  Controls the Show History property for solids in a drawing.  Minimum AutoCAD version R2007.
	SolidHistoryMode SolidHistoryMode
	// The $DWFFRAME header variable.  Determines whether DWF or DWFx underlay frames are visible or plotted in the current drawing.  Minimum AutoCAD version R2007.
	UnderlayFrameMode UnderlayFrameMode
	// The $REALWORLDSCALE header variable.  Drawing is scaled to the real world.  Minimum AutoCAD version R2007.
	UseRealWorldScale bool
	// The $INTERFERECOLOR header variable.  Represents the ACI color index of the "interference objects" created during the interfere command.  Minimum AutoCAD version R2007.
	InterferenceObjectColor Color
	// The $INTERFEREOBJVS header variable.  Hard-pointer ID to the visual stype for interference objects.  Minimum AutoCAD version R2007.
	InterferenceObjectVisualStylePointer Handle
	// The $INTERFEREVPVS header variable.  Hard-pointer ID to the visual styoe for the viewport during interference checking.  Minimum AutoCAD version R2007.
	InterferenceViewPortVisualStylePointer Handle
	// The $CSHADOW header variable.  Shadow mode for a 3D object.  Minimum AutoCAD version R2007.
	ShadowMode ShadowMode
	// The $SHADOWPLANELOCATION header variable.  Location of the ground shadow plane.  This is a Z axis ordinate.  Minimum AutoCAD version R2007.
	ShadowPlaneZOffset float64
	// The $AXISMODE header variable.  Axis on.  Maximum AutoCAD version R10.
	AxisOn bool
	// The $AXISUNIT header variable.  Axis X and Y tick spacing.  Maximum AutoCAD version R10.
	AxisTickSpacing Vector
	// The $FASTZOOM header variable.  Fast zoom enabled.  Maximum AutoCAD version R10.
	FastZoom bool
	// The $GRIDMODE header variable.  Grid mode on.  Maximum AutoCAD version R10.
	GridOn bool
	// The $GRIDUNIT header variable.  Grid X and Y spacing.  Maximum AutoCAD version R10.
	GridSpacing Vector
	// The $SNAPANG header variable.  Snap grid rotation angle.  Maximum AutoCAD version R10.
	SnapRotationAngle float64
	// The $SNAPBASE header variable.  Snap/grid/base point (in UCS).  Maximum AutoCAD version R10.
	SnapBasePoint Point
	// The $SNAPISOPAIR header variable.  Isometric plane.  Maximum AutoCAD version R10.
	SnapIsometricPlane SnapIsometricPlane
	// The $SNAPMODE header variable.  Snap mode on.  Maximum AutoCAD version R10.
	SnapOn bool
	// The $SNAPSTYLE header variable.  Snap style.  Maximum AutoCAD version R10.
	SnapStyle SnapStyle
	// The $SNAPUNIT header variable.  Snap grid X and Y spacing.  Maximum AutoCAD version R10.
	SnapSpacing Vector
	// The $VIEWCTR header variable.  XY center of current view on screen.  Maximum AutoCAD version R10.
	ViewCenter Point
	// The $VIEWDIR header variable.  Viewing direction (direction from target in WCS).  Maximum AutoCAD version R10.
	ViewDirection Vector
	// The $VIEWSIZE header variable.  Height of view.  Maximum AutoCAD version R10.
	ViewHeight float64
}

func NewHeader() *Header {
	return &Header{
		Version: R12,
		MaintenanceVersion: 0,
		DrawingCodePage: "ANSI_1252",
		LastSavedBy: "",
		RequiredVersions: 0,
		InsertionBase: *NewOrigin(),
		MinimumDrawingExtents: *NewOrigin(),
		MaximumDrawingExtents: *NewOrigin(),
		MinimumDrawingLimits: *NewOrigin(),
		MaximumDrawingLimits: Point{12.0, 9.0, 0.0},
		DrawOrthoganalLines: false,
		UseRegenMode: true,
		FillModeOn: true,
		UseQuickTextMode: false,
		MirrorText: false,
		DragMode: DragModeAuto,
		LineTypeScale: 1.0,
		ObjectSnapFlags: 37,
		AttributeVisibility: AttributeVisibilityNormal,
		DefaultTextHeight: 0.2,
		TraceWidth: 0.05,
		TextStyle: "STANDARD",
		CurrentLayer: "0",
		CurrentEntityLineType: "BYLAYER",
		CurrentEntityColor: ByLayer(),
		CurrentEntityLineTypeScale: 1.0,
		RetainDeletedObjects: true,
		DisplaySilhouetteCurvesInWireframeMode: false,
		SolidVisualStylePointer: Handle(0),
		DimensioningScaleFactor: 1.0,
		DimensioningArrowSize: 0.18,
		DimensionExtensionLineOffset: 0.0625,
		DimensionLineIncrement: 0.38,
		DimensionDistanceRoundingValue: 0.0,
		DimensionLineExtension: 0.0,
		DimensionExtensionLineExtension: 0.18,
		DimensionPlusTolerance: 0.0,
		DimensionMinusTolerance: 0.0,
		DimensioningTextHeight: 0.18,
		CenterMarkSize: 0.09,
		DimensioningTickSize: 0.0,
		GenerateDimensionTolerances: false,
		GenerateDimensionLimits: false,
		DimensionTextInsideHorizontal: true,
		DimensionTextOutsideHorizontal: true,
		SuppressFirstDimensionExtensionLine: false,
		SuppressSecondDimensionExtensionLine: false,
		TextAboveDimensionLine: false,
		DimensionUnitZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		ArrowBlockName: "",
		CreateAssociativeDimensioning: true,
		RecomputeDimensionsWhileDragging: true,
		DimensioningSuffix: "",
		AlternateDimensioningSuffix: "",
		UseAlternateDimensioning: false,
		AlternateDimensioningDecimalPlaces: 2,
		AlternateDimensioningScaleFactor: 25.4,
		DimensionLinearMeasurementsScaleFactor: 1.0,
		ForceDimensionLineExtensionsOutsideIfTextIs: false,
		DimensionVerticalTextPosition: 0.0,
		ForceDimensionTextInsideExtensions: false,
		SuppressOutsideExtensionDimensionLines: false,
		UseSeparateArrowBlocksForDimensions: false,
		FirstArrowBlockName: "",
		SecondArrowBlockName: "",
		DimensionStyleName: "STANDARD",
		DimensionLineColor: ByBlock(),
		DimensionExtensionLineColor: ByBlock(),
		DimensionTextColor: ByBlock(),
		DimensionToleranceDisplayScaleFactor: 1.0,
		DimensionLineGap: 0.09,
		DimensionTextJustification: DimensionTextJustificationAboveLineCenter,
		DimensionToleranceVerticalJustification: JustificationMiddle,
		DimensionToleranceZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		AlternateDimensioningZeroSupression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		AlternateDimensioningToleranceZeroSupression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		DimensionTextAndArrowPlacement: DimensionFitTextAndArrowsOutsideLines,
		DimensionCursorControlsTextPosition: false,
		DimensionUnitFormat: UnitFormatDecimal,
		DimensionUnitToleranceDecimalPlaces: 4,
		DimensionToleranceDecimalPlaces: 4,
		AlternateDimensioningUnits: UnitFormatDecimal,
		AlternateDimensioningToleranceDecimalPlaces: 2,
		DimensionTextStyle: "STANDARD",
		DimensioningAngleFormat: AngleFormatDecimalDegrees,
		AngularDimensionPrecision: 0,
		AlternateDimensioningUnitRounding: 0.0,
		DimensionAngleZeroSuppression: UnitZeroSuppressionSuppressZeroFeetAndZeroInches,
		DimensionDecimalSeparatorRune: '.',
		DimensionTextHeightScaleFactor: DimensionFractionFormatHorizontalStacking,
		DimensionLeaderBlockName: "",
		DimensionNonAngularUnits: NonAngularUnitsDecimal,
		DimensionLineWeight: NewLineWeightByLayer(),
		DimensionExtensionLineWeight: NewLineWeightByLayer(),
		DimensionTextMovementRule: DimensionTextMovementRuleMoveLineWithText,
		DimensionLineFixedLength: 1.0,
		DimensionLineFixedLengthOn: false,
		DimensionTransverseSegmentAngleInJoggedRadius: math.Pi / 4.0,
		DimensionTextBackgroundColorMode: DimensionTextBackgroundColorModeNone,
		DimensionTextBackgroundCustomColor: ByBlock(),
		DimensionArcSymbolDisplayMode: DimensionArcSymbolDisplayModeSymbolBeforeText,
		DimensionLineType: "",
		DimensionFirstExtensionLineType: "",
		DimensionSecondExtensionLineType: "",
		DimensionTextDirection: TextDirectionLeftToRight,
		UnitFormat: UnitFormatDecimal,
		UnitPrecision: 4,
		SketchRecordIncrement: 0.1,
		FilletRadius: 0.0,
		AngleUnitFormat: AngleFormatDecimalDegrees,
		AngleUnitPrecision: 0,
		FileName: ".",
		Elevation: 0.0,
		PaperspaceElevation: 0.0,
		Thickness: 0.0,
		UseLimitsChecking: false,
		BlipMode: false,
		FirstChamferDistance: 0.0,
		SecondChamferDistance: 0.0,
		ChamferLength: 0.0,
		ChamferAngle: 0.0,
		PolylineSketchMode: PolySketchModeSketchLines,
		CreationDate: time.Now(),
		CreationDateUniversal: time.Now().UTC(),
		UpdateDate: time.Now(),
		UpdateDateUniversal: time.Now().UTC(),
		TimeInDrawing: time.Duration(0),
		UserElapsedTimer: time.Duration(0),
		UserTimerOn: true,
		AngleZeroDirection: 0.0,
		AngleDirection: AngleDirectionCounterClockwise,
		PointDisplayMode: 0,
		PointDisplaySize: 0.0,
		DefaultPolylineWidth: 0.0,
		CoordinateDisplay: CoordinateDisplayContinuousUpdate,
		DisplaySplinePolygonControl: false,
		PEditSplineCurveType: PolylineCurvedAndSmoothSurfaceTypeCubicBSpline,
		LineSegmentsPerSplinePatch: 8,
		ShowAttributeEntryDialogs: true,
		PromptForAttributeOnInsert: true,
		HandlesEnabled: true,
		NextAvailableHandle: 0,
		MeshTabulationsInFirstDirection: 6,
		MeshTabulationsInSecondDirection: 6,
		PEditSmoothSurfaceType: PolylineCurvedAndSmoothSurfaceTypeCubicBSpline,
		PEditSmoothMDensith: 6,
		PEditSmoothNDensith: 6,
		UCSDefinitionName: "",
		UCSName: "",
		UCSOrigin: *NewOrigin(),
		UCSXAxis: *NewXAxis(),
		UCSYAxis: *NewYAxis(),
		OrthoUCSReference: "",
		OrthgraphicViewType: OrthographicViewTypeNone,
		UCSOriginTop: *NewOrigin(),
		UCSOriginBottom: *NewOrigin(),
		UCSOriginLeft: *NewOrigin(),
		UCSOriginRight: *NewOrigin(),
		UCSOriginFront: *NewOrigin(),
		UCSOriginBack: *NewOrigin(),
		PaperspaceUCSDefinitionName: "",
		PaperspaceUCSName: "",
		PaperspaceUCSOrigin: *NewOrigin(),
		PaperspaceXAxis: *NewXAxis(),
		PaperspaceYAxis: *NewYAxis(),
		PaperspaceOrthoUCSReference: "",
		PaperspaceOrthographicViewType: OrthographicViewTypeNone,
		PaperspaceUCSOriginTop: *NewOrigin(),
		PaperspaceUCSOriginBottom: *NewOrigin(),
		PaperspaceUCSOriginLeft: *NewOrigin(),
		PaperspaceUCSOriginRight: *NewOrigin(),
		PaperspaceUCSOriginFront: *NewOrigin(),
		PaperspaceUCSOriginBack: *NewOrigin(),
		UserInt1: 0,
		UserInt2: 0,
		UserInt3: 0,
		UserInt4: 0,
		UserInt5: 0,
		UserReal1: 0.0,
		UserReal2: 0.0,
		UserReal3: 0.0,
		UserReal4: 0.0,
		UserReal5: 0.0,
		SetUCSToWCSInDViewOrVPoint: true,
		EdgeShading: ShadeEdgeModeFacesInEntityColorEdgesInBlack,
		PercentAmbientToDiffuse: 70,
		PreviousReleaseTileCompatability: true,
		MaximumActiveViewports: 64,
		PaperspaceInsertionBase: *NewOrigin(),
		LimitCheckingInPaperspace: false,
		PaperspaceMinimumDrawingExtents: Point{1.0e20, 1.0e20, 1.0e20},
		PaperspaceMaximumDrawingExtents: Point{-1.0e20, -1.0e20, -1.0e20},
		PaperspaceMinimumDrawingLimits: *NewOrigin(),
		PaperspaceMaximumDrawingLimits: Point{12.0, 9.0, 0.0},
		DisplayFractionsInInput: false,
		RetainXRefDependentVisibilitySettings: true,
		IsPolylineContinuousAroundVerticies: false,
		ScaleLineTypesInPaperspace: true,
		SpacialIndexMaxDepth: 3020,
		PickStyle: PickStyleGroup,
		CurrentMultilineStyle: "STANDARD",
		CurrentMultilineJustification: JustificationTop,
		CurrentMultilineScale: 1.0,
		SaveProxyGraphics: true,
		DrawingUnits: DrawingUnitsEnglish,
		NewObjectLineWeight: LineWeightByBlock,
		EndCapSetting: EndCapSettingNone,
		LineweightJointSetting: JoinStyleNone,
		DisplayLinewieghtInModelAndLayoutTab: false,
		DefaultDrawingUnits: UnitsUnitless,
		HyperlinkBase: "",
		Stylesheet: "",
		CanUseInPlaceReferenceEditing: true,
		NewObjectPlotStyleHandle: 0,
		NewObjectPlotStyle: PlotStyleByLayer,
		UsesColorDependentPlotStyleTables: true,
		FingerprintGuid: uuid.New(),
		VersionGuid: uuid.New(),
		UseACad2000SymbolTableNaming: true,
		ViewportViewScaleFactor: 0.0,
		OleStartup: false,
		ObjectSortingMethodsFlags: 127,
		LayerAndSpatialIndexSaveMode: LayerAndSpatialIndexSaveModeNone,
		HideTextObjectsWhenProducintHiddenView: false,
		XRefClippingBoundaryVisible: XrefClippingBoundaryVisibilityDisplayedNotPlotted,
		HaloGapPercent: 0.0,
		ObscuredLineColor: ByEntity(),
		ObscuredLineTypeStyle: LineTypeStyleOff,
		DisplayIntersectionPolylines: false,
		IntersectionPolylineColor: ByEntity(),
		DimensionObjectAssociativity: DimensionAssociativityNonAssociativeObjects,
		ProjectName: "",
		UseCameraDisplay: false,
		LensLength: 50.0,
		CameraHeight: 0.0,
		StepsPerSecondInWalkOrFlyMode: 2.0,
		StepSizeInWalkOrFlyMode: 6.0,
		Dwf3DPrecision: Dwf3DPrecisionDeviation0_5,
		LastPolySolidWidth: 0.25,
		LastPolySolidHeight: 4.0,
		LoftOperationFirstDraftAngle: math.Pi / 2.0,
		LoftOperationSecondDraftAngle: math.Pi / 2.0,
		LoftOperationFirstMagnitude: 0.0,
		LoftOperationSecondMagnitude: 0.0,
		LoftFlags: 7,
		LoftedObjectNormalMode: LoftedObjectNormalModeSmoothFit,
		Latitude: 37.7950,
		Longitude: -122.3940,
		AngleBetweenYAxisAndNorth: 0.0,
		TimeZone: TimeZonePacificTime_US_Canada_SanFrancisco_Vancouver,
		UseLightGlyphDisplay: true,
		UseTileModeLightSync: true,
		CurrentMaterialHandle: 0,
		NewSolidsContainHistory: false,
		SolidHistoryMode: SolidHistoryModeDoesNotOverride,
		UnderlayFrameMode: UnderlayFrameModeDisplayNoPlot,
		UseRealWorldScale: true,
		InterferenceObjectColor: Color(1),
		InterferenceObjectVisualStylePointer: 0,
		InterferenceViewPortVisualStylePointer: 0,
		ShadowMode: ShadowModeCastsAndReceivesShadows,
		ShadowPlaneZOffset: 0.0,
		AxisOn: false,
		AxisTickSpacing: *NewZeroVector(),
		FastZoom: true,
		GridOn: false,
		GridSpacing: Vector{1.0, 1.0, 0.0},
		SnapRotationAngle: 0.0,
		SnapBasePoint: *NewOrigin(),
		SnapIsometricPlane: SnapIsometricPlaneLeft,
		SnapOn: false,
		SnapStyle: SnapStyleStandard,
		SnapSpacing: Vector{1.0, 1.0, 0.0},
		ViewCenter: *NewOrigin(),
		ViewDirection: *NewZAxis(),
		ViewHeight: 1.0,
	}
}

// EndPointSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) EndPointSnap() (val bool) {
	return h.ObjectSnapFlags & 1 != 0
}

// EndPointSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetEndPointSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 1
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^1
	}
}

// MidPointSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) MidPointSnap() (val bool) {
	return h.ObjectSnapFlags & 2 != 0
}

// MidPointSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetMidPointSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 2
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^2
	}
}

// CenterSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) CenterSnap() (val bool) {
	return h.ObjectSnapFlags & 4 != 0
}

// CenterSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetCenterSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 4
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^4
	}
}

// NodeSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) NodeSnap() (val bool) {
	return h.ObjectSnapFlags & 8 != 0
}

// NodeSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetNodeSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 8
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^8
	}
}

// QuadrantSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) QuadrantSnap() (val bool) {
	return h.ObjectSnapFlags & 16 != 0
}

// QuadrantSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetQuadrantSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 16
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^16
	}
}

// IntersectionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) IntersectionSnap() (val bool) {
	return h.ObjectSnapFlags & 32 != 0
}

// IntersectionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetIntersectionSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 32
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^32
	}
}

// InsertionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) InsertionSnap() (val bool) {
	return h.ObjectSnapFlags & 64 != 0
}

// InsertionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetInsertionSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 64
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^64
	}
}

// PerpendicularSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) PerpendicularSnap() (val bool) {
	return h.ObjectSnapFlags & 128 != 0
}

// PerpendicularSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetPerpendicularSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 128
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^128
	}
}

// TangentSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) TangentSnap() (val bool) {
	return h.ObjectSnapFlags & 256 != 0
}

// TangentSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetTangentSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 256
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^256
	}
}

// NearestSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) NearestSnap() (val bool) {
	return h.ObjectSnapFlags & 512 != 0
}

// NearestSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetNearestSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 512
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^512
	}
}

// ApparentIntersectionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) ApparentIntersectionSnap() (val bool) {
	return h.ObjectSnapFlags & 2048 != 0
}

// ApparentIntersectionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetApparentIntersectionSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 2048
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^2048
	}
}

// ExtensionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) ExtensionSnap() (val bool) {
	return h.ObjectSnapFlags & 4096 != 0
}

// ExtensionSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetExtensionSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 4096
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^4096
	}
}

// ParallelSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) ParallelSnap() (val bool) {
	return h.ObjectSnapFlags & 8192 != 0
}

// ParallelSnap status flag.  Maximum AutoCAD version R14.
func (h *Header) SetParallelSnap(val bool) {
	if val {
		h.ObjectSnapFlags = h.ObjectSnapFlags | 8192
	} else {
		h.ObjectSnapFlags = h.ObjectSnapFlags & ^8192
	}
}

// SortObjectsForObjectSelection status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForObjectSelection() (val bool) {
	return h.ObjectSortingMethodsFlags & 1 != 0
}

// SortObjectsForObjectSelection status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForObjectSelection(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 1
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^1
	}
}

// SortObjectsForObjectSnap status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForObjectSnap() (val bool) {
	return h.ObjectSortingMethodsFlags & 2 != 0
}

// SortObjectsForObjectSnap status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForObjectSnap(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 2
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^2
	}
}

// SortObjectsForRedraw status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForRedraw() (val bool) {
	return h.ObjectSortingMethodsFlags & 4 != 0
}

// SortObjectsForRedraw status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForRedraw(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 4
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^4
	}
}

// SortObjectsForMSlide status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForMSlide() (val bool) {
	return h.ObjectSortingMethodsFlags & 8 != 0
}

// SortObjectsForMSlide status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForMSlide(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 8
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^8
	}
}

// SortObjectsForRegen status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForRegen() (val bool) {
	return h.ObjectSortingMethodsFlags & 16 != 0
}

// SortObjectsForRegen status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForRegen(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 16
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^16
	}
}

// SortObjectsForPlotting status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForPlotting() (val bool) {
	return h.ObjectSortingMethodsFlags & 32 != 0
}

// SortObjectsForPlotting status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForPlotting(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 32
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^32
	}
}

// SortObjectsForPostScriptOutput status flag.  Minimum AutoCAD version R2004.
func (h *Header) SortObjectsForPostScriptOutput() (val bool) {
	return h.ObjectSortingMethodsFlags & 64 != 0
}

// SortObjectsForPostScriptOutput status flag.  Minimum AutoCAD version R2004.
func (h *Header) SetSortObjectsForPostScriptOutput(val bool) {
	if val {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags | 64
	} else {
		h.ObjectSortingMethodsFlags = h.ObjectSortingMethodsFlags & ^64
	}
}

// NoTwist status flag.  Minimum AutoCAD version R2007.
func (h *Header) NoTwist() (val bool) {
	return h.LoftFlags & 1 != 0
}

// NoTwist status flag.  Minimum AutoCAD version R2007.
func (h *Header) SetNoTwist(val bool) {
	if val {
		h.LoftFlags = h.LoftFlags | 1
	} else {
		h.LoftFlags = h.LoftFlags & ^1
	}
}

// AlignDirection status flag.  Minimum AutoCAD version R2007.
func (h *Header) AlignDirection() (val bool) {
	return h.LoftFlags & 2 != 0
}

// AlignDirection status flag.  Minimum AutoCAD version R2007.
func (h *Header) SetAlignDirection(val bool) {
	if val {
		h.LoftFlags = h.LoftFlags | 2
	} else {
		h.LoftFlags = h.LoftFlags & ^2
	}
}

// Simplify status flag.  Minimum AutoCAD version R2007.
func (h *Header) Simplify() (val bool) {
	return h.LoftFlags & 4 != 0
}

// Simplify status flag.  Minimum AutoCAD version R2007.
func (h *Header) SetSimplify(val bool) {
	if val {
		h.LoftFlags = h.LoftFlags | 4
	} else {
		h.LoftFlags = h.LoftFlags & ^4
	}
}

// Close status flag.  Minimum AutoCAD version R2007.
func (h *Header) Close() (val bool) {
	return h.LoftFlags & 8 != 0
}

// Close status flag.  Minimum AutoCAD version R2007.
func (h *Header) SetClose(val bool) {
	if val {
		h.LoftFlags = h.LoftFlags | 8
	} else {
		h.LoftFlags = h.LoftFlags & ^8
	}
}

func (h Header) writeHeaderSection(writer codePairWriter) error {
	pairs := make([]CodePair, 0)

	// $ACADVER
	pairs = append(pairs, NewStringCodePair(9, "$ACADVER"))
	pairs = append(pairs, NewStringCodePair(1, h.Version.String()))

	// $ACADMAINTVER
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$ACADMAINTVER"))
		pairs = append(pairs, NewShortCodePair(70, h.MaintenanceVersion))
	}

	// $DWGCODEPAGE
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DWGCODEPAGE"))
		pairs = append(pairs, NewStringCodePair(3, h.DrawingCodePage))
	}

	// $LASTSAVEDBY
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$LASTSAVEDBY"))
		pairs = append(pairs, NewStringCodePair(1, h.LastSavedBy))
	}

	// $REQUIREDVERSIONS
	if h.Version >= R2013 {
		pairs = append(pairs, NewStringCodePair(9, "$REQUIREDVERSIONS"))
		pairs = append(pairs, NewLongCodePair(160, h.RequiredVersions))
	}

	// $INSBASE
	pairs = append(pairs, NewStringCodePair(9, "$INSBASE"))
	pairs = append(pairs, NewDoubleCodePair(10, h.InsertionBase.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.InsertionBase.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.InsertionBase.Z))

	// $EXTMIN
	pairs = append(pairs, NewStringCodePair(9, "$EXTMIN"))
	pairs = append(pairs, NewDoubleCodePair(10, h.MinimumDrawingExtents.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.MinimumDrawingExtents.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.MinimumDrawingExtents.Z))

	// $EXTMAX
	pairs = append(pairs, NewStringCodePair(9, "$EXTMAX"))
	pairs = append(pairs, NewDoubleCodePair(10, h.MaximumDrawingExtents.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.MaximumDrawingExtents.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.MaximumDrawingExtents.Z))

	// $LIMMIN
	pairs = append(pairs, NewStringCodePair(9, "$LIMMIN"))
	pairs = append(pairs, NewDoubleCodePair(10, h.MinimumDrawingLimits.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.MinimumDrawingLimits.Y))

	// $LIMMAX
	pairs = append(pairs, NewStringCodePair(9, "$LIMMAX"))
	pairs = append(pairs, NewDoubleCodePair(10, h.MaximumDrawingLimits.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.MaximumDrawingLimits.Y))

	// $ORTHOMODE
	pairs = append(pairs, NewStringCodePair(9, "$ORTHOMODE"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DrawOrthoganalLines)))

	// $REGENMODE
	pairs = append(pairs, NewStringCodePair(9, "$REGENMODE"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UseRegenMode)))

	// $FILLMODE
	pairs = append(pairs, NewStringCodePair(9, "$FILLMODE"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.FillModeOn)))

	// $QTEXTMODE
	pairs = append(pairs, NewStringCodePair(9, "$QTEXTMODE"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UseQuickTextMode)))

	// $MIRRTEXT
	pairs = append(pairs, NewStringCodePair(9, "$MIRRTEXT"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.MirrorText)))

	// $DRAGMODE
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$DRAGMODE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DragMode)))
	}

	// $LTSCALE
	pairs = append(pairs, NewStringCodePair(9, "$LTSCALE"))
	pairs = append(pairs, NewDoubleCodePair(40, h.LineTypeScale))

	// $OSMODE
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$OSMODE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.ObjectSnapFlags)))
	}

	// $ATTMODE
	pairs = append(pairs, NewStringCodePair(9, "$ATTMODE"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.AttributeVisibility)))

	// $TEXTSIZE
	pairs = append(pairs, NewStringCodePair(9, "$TEXTSIZE"))
	pairs = append(pairs, NewDoubleCodePair(40, ensurePositiveOrDefault(h.DefaultTextHeight, 0.2)))

	// $TRACEWID
	pairs = append(pairs, NewStringCodePair(9, "$TRACEWID"))
	pairs = append(pairs, NewDoubleCodePair(40, ensurePositiveOrDefault(h.TraceWidth, 0.05)))

	// $TEXTSTYLE
	pairs = append(pairs, NewStringCodePair(9, "$TEXTSTYLE"))
	pairs = append(pairs, NewStringCodePair(7, defaultIfEmpty(h.TextStyle, "STANDARD")))

	// $CLAYER
	pairs = append(pairs, NewStringCodePair(9, "$CLAYER"))
	pairs = append(pairs, NewStringCodePair(8, defaultIfEmpty(h.CurrentLayer, "0")))

	// $CELTYPE
	pairs = append(pairs, NewStringCodePair(9, "$CELTYPE"))
	pairs = append(pairs, NewStringCodePair(6, defaultIfEmpty(h.CurrentEntityLineType, "BYLAYER")))

	// $CECOLOR
	pairs = append(pairs, NewStringCodePair(9, "$CECOLOR"))
	pairs = append(pairs, NewShortCodePair(62, int16(h.CurrentEntityColor)))

	// $CELTSCALE
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$CELTSCALE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.CurrentEntityLineTypeScale))
	}

	// $DELOBJ
	if h.Version >= R13 && h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$DELOBJ"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.RetainDeletedObjects)))
	}

	// $DISPSILH
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DISPSILH"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DisplaySilhouetteCurvesInWireframeMode)))
	}

	// $DRAGVS
	if h.Version >= R2007 && h.SolidVisualStylePointer != Handle(0) {
		pairs = append(pairs, NewStringCodePair(9, "$DRAGVS"))
		pairs = append(pairs, NewStringCodePair(349, stringFromHandle(h.SolidVisualStylePointer)))
	}

	// $DIMSCALE
	pairs = append(pairs, NewStringCodePair(9, "$DIMSCALE"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensioningScaleFactor))

	// $DIMASZ
	pairs = append(pairs, NewStringCodePair(9, "$DIMASZ"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensioningArrowSize))

	// $DIMEXO
	pairs = append(pairs, NewStringCodePair(9, "$DIMEXO"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionExtensionLineOffset))

	// $DIMDLI
	pairs = append(pairs, NewStringCodePair(9, "$DIMDLI"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionLineIncrement))

	// $DIMRND
	pairs = append(pairs, NewStringCodePair(9, "$DIMRND"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionDistanceRoundingValue))

	// $DIMDLE
	pairs = append(pairs, NewStringCodePair(9, "$DIMDLE"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionLineExtension))

	// $DIMEXE
	pairs = append(pairs, NewStringCodePair(9, "$DIMEXE"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionExtensionLineExtension))

	// $DIMTP
	pairs = append(pairs, NewStringCodePair(9, "$DIMTP"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionPlusTolerance))

	// $DIMTM
	pairs = append(pairs, NewStringCodePair(9, "$DIMTM"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionMinusTolerance))

	// $DIMTXT
	pairs = append(pairs, NewStringCodePair(9, "$DIMTXT"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensioningTextHeight))

	// $DIMCEN
	pairs = append(pairs, NewStringCodePair(9, "$DIMCEN"))
	pairs = append(pairs, NewDoubleCodePair(40, h.CenterMarkSize))

	// $DIMTSZ
	pairs = append(pairs, NewStringCodePair(9, "$DIMTSZ"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensioningTickSize))

	// $DIMTOL
	pairs = append(pairs, NewStringCodePair(9, "$DIMTOL"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.GenerateDimensionTolerances)))

	// $DIMLIM
	pairs = append(pairs, NewStringCodePair(9, "$DIMLIM"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.GenerateDimensionLimits)))

	// $DIMTIH
	pairs = append(pairs, NewStringCodePair(9, "$DIMTIH"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DimensionTextInsideHorizontal)))

	// $DIMTOH
	pairs = append(pairs, NewStringCodePair(9, "$DIMTOH"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DimensionTextOutsideHorizontal)))

	// $DIMSE1
	if h.Version >= R12 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMSE1"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SuppressFirstDimensionExtensionLine)))
	}

	// $DIMSE2
	if h.Version >= R12 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMSE2"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SuppressSecondDimensionExtensionLine)))
	}

	// $DIMTAD
	pairs = append(pairs, NewStringCodePair(9, "$DIMTAD"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.TextAboveDimensionLine)))

	// $DIMZIN
	pairs = append(pairs, NewStringCodePair(9, "$DIMZIN"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionUnitZeroSuppression)))

	// $DIMBLK
	pairs = append(pairs, NewStringCodePair(9, "$DIMBLK"))
	pairs = append(pairs, NewStringCodePair(1, h.ArrowBlockName))

	// $DIMASO
	pairs = append(pairs, NewStringCodePair(9, "$DIMASO"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.CreateAssociativeDimensioning)))

	// $DIMSHO
	pairs = append(pairs, NewStringCodePair(9, "$DIMSHO"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.RecomputeDimensionsWhileDragging)))

	// $DIMPOST
	pairs = append(pairs, NewStringCodePair(9, "$DIMPOST"))
	pairs = append(pairs, NewStringCodePair(1, h.DimensioningSuffix))

	// $DIMAPOST
	pairs = append(pairs, NewStringCodePair(9, "$DIMAPOST"))
	pairs = append(pairs, NewStringCodePair(1, h.AlternateDimensioningSuffix))

	// $DIMALT
	pairs = append(pairs, NewStringCodePair(9, "$DIMALT"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UseAlternateDimensioning)))

	// $DIMALTD
	pairs = append(pairs, NewStringCodePair(9, "$DIMALTD"))
	pairs = append(pairs, NewShortCodePair(70, h.AlternateDimensioningDecimalPlaces))

	// $DIMALTF
	pairs = append(pairs, NewStringCodePair(9, "$DIMALTF"))
	pairs = append(pairs, NewDoubleCodePair(40, h.AlternateDimensioningScaleFactor))

	// $DIMLFAC
	pairs = append(pairs, NewStringCodePair(9, "$DIMLFAC"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionLinearMeasurementsScaleFactor))

	// $DIMTOFL
	pairs = append(pairs, NewStringCodePair(9, "$DIMTOFL"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.ForceDimensionLineExtensionsOutsideIfTextIs)))

	// $DIMTVP
	pairs = append(pairs, NewStringCodePair(9, "$DIMTVP"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DimensionVerticalTextPosition))

	// $DIMTIX
	pairs = append(pairs, NewStringCodePair(9, "$DIMTIX"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.ForceDimensionTextInsideExtensions)))

	// $DIMSOXD
	pairs = append(pairs, NewStringCodePair(9, "$DIMSOXD"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SuppressOutsideExtensionDimensionLines)))

	// $DIMSAH
	pairs = append(pairs, NewStringCodePair(9, "$DIMSAH"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UseSeparateArrowBlocksForDimensions)))

	// $DIMBLK1
	pairs = append(pairs, NewStringCodePair(9, "$DIMBLK1"))
	pairs = append(pairs, NewStringCodePair(1, h.FirstArrowBlockName))

	// $DIMBLK2
	pairs = append(pairs, NewStringCodePair(9, "$DIMBLK2"))
	pairs = append(pairs, NewStringCodePair(1, h.SecondArrowBlockName))

	// $DIMSTYLE
	pairs = append(pairs, NewStringCodePair(9, "$DIMSTYLE"))
	pairs = append(pairs, NewStringCodePair(2, defaultIfEmpty(h.DimensionStyleName, "STANDARD")))

	// $DIMCLRD
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMCLRD"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionLineColor)))
	}

	// $DIMCLRE
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMCLRE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionExtensionLineColor)))
	}

	// $DIMCLRT
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMCLRT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextColor)))
	}

	// $DIMTFAC
	if h.Version >= R12 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTFAC"))
		pairs = append(pairs, NewDoubleCodePair(40, h.DimensionToleranceDisplayScaleFactor))
	}

	// $DIMGAP
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMGAP"))
		pairs = append(pairs, NewDoubleCodePair(40, h.DimensionLineGap))
	}

	// $DIMJUST
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMJUST"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextJustification)))
	}

	// $DIMSD1
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMSD1"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SuppressFirstDimensionExtensionLine)))
	}

	// $DIMSD2
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMSD2"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SuppressSecondDimensionExtensionLine)))
	}

	// $DIMTOLJ
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTOLJ"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionToleranceVerticalJustification)))
	}

	// $DIMTZIN
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTZIN"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionToleranceZeroSuppression)))
	}

	// $DIMALTZ
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMALTZ"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.AlternateDimensioningZeroSupression)))
	}

	// $DIMALTTZ
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMALTTZ"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.AlternateDimensioningToleranceZeroSupression)))
	}

	// $DIMFIT
	if h.Version >= R13 && h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMFIT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextAndArrowPlacement)))
	}

	// $DIMUPT
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMUPT"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DimensionCursorControlsTextPosition)))
	}

	// $DIMUNIT
	if h.Version >= R13 && h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMUNIT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionUnitFormat)))
	}

	// $DIMDEC
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMDEC"))
		pairs = append(pairs, NewShortCodePair(70, h.DimensionUnitToleranceDecimalPlaces))
	}

	// $DIMTDEC
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTDEC"))
		pairs = append(pairs, NewShortCodePair(70, h.DimensionToleranceDecimalPlaces))
	}

	// $DIMALTU
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMALTU"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.AlternateDimensioningUnits)))
	}

	// $DIMALTTD
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMALTTD"))
		pairs = append(pairs, NewShortCodePair(70, h.AlternateDimensioningToleranceDecimalPlaces))
	}

	// $DIMTXSTY
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTXSTY"))
		pairs = append(pairs, NewStringCodePair(7, h.DimensionTextStyle))
	}

	// $DIMAUNIT
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMAUNIT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensioningAngleFormat)))
	}

	// $DIMADEC
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMADEC"))
		pairs = append(pairs, NewShortCodePair(70, h.AngularDimensionPrecision))
	}

	// $DIMALTRND
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMALTRND"))
		pairs = append(pairs, NewDoubleCodePair(40, h.AlternateDimensioningUnitRounding))
	}

	// $DIMAZIN
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMAZIN"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionAngleZeroSuppression)))
	}

	// $DIMDSEP
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMDSEP"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionDecimalSeparatorRune)))
	}

	// $DIMATFIT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMATFIT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextAndArrowPlacement)))
	}

	// $DIMFRAC
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMFRAC"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextHeightScaleFactor)))
	}

	// $DIMLDRBLK
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLDRBLK"))
		pairs = append(pairs, NewStringCodePair(1, h.DimensionLeaderBlockName))
	}

	// $DIMLUNIT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLUNIT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionNonAngularUnits)))
	}

	// $DIMLWD
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLWD"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionLineWeight)))
	}

	// $DIMLWE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLWE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionExtensionLineWeight)))
	}

	// $DIMTMOVE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTMOVE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextMovementRule)))
	}

	// $DIMFXL
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMFXL"))
		pairs = append(pairs, NewDoubleCodePair(40, h.DimensionLineFixedLength))
	}

	// $DIMFXLON
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMFXLON"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DimensionLineFixedLengthOn)))
	}

	// $DIMJOGANG
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMJOGANG"))
		pairs = append(pairs, NewDoubleCodePair(40, h.DimensionTransverseSegmentAngleInJoggedRadius))
	}

	// $DIMTFILL
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTFILL"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextBackgroundColorMode)))
	}

	// $DIMTFILLCLR
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTFILLCLR"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextBackgroundCustomColor)))
	}

	// $DIMARCSYM
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMARCSYM"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionArcSymbolDisplayMode)))
	}

	// $DIMLTYPE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLTYPE"))
		pairs = append(pairs, NewStringCodePair(6, h.DimensionLineType))
	}

	// $DIMLTEX1
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLTEX1"))
		pairs = append(pairs, NewStringCodePair(6, h.DimensionFirstExtensionLineType))
	}

	// $DIMLTEX2
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMLTEX2"))
		pairs = append(pairs, NewStringCodePair(6, h.DimensionSecondExtensionLineType))
	}

	// $DIMTXTDIRECTION
	if h.Version >= R2010 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMTXTDIRECTION"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DimensionTextDirection)))
	}

	// $LUNITS
	pairs = append(pairs, NewStringCodePair(9, "$LUNITS"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.UnitFormat)))

	// $LUPREC
	pairs = append(pairs, NewStringCodePair(9, "$LUPREC"))
	pairs = append(pairs, NewShortCodePair(70, h.UnitPrecision))

	// $SKETCHINC
	pairs = append(pairs, NewStringCodePair(9, "$SKETCHINC"))
	pairs = append(pairs, NewDoubleCodePair(40, h.SketchRecordIncrement))

	// $FILLETRAD
	pairs = append(pairs, NewStringCodePair(9, "$FILLETRAD"))
	pairs = append(pairs, NewDoubleCodePair(40, h.FilletRadius))

	// $AUNITS
	pairs = append(pairs, NewStringCodePair(9, "$AUNITS"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.AngleUnitFormat)))

	// $AUPREC
	pairs = append(pairs, NewStringCodePair(9, "$AUPREC"))
	pairs = append(pairs, NewShortCodePair(70, h.AngleUnitPrecision))

	// $MENU
	pairs = append(pairs, NewStringCodePair(9, "$MENU"))
	pairs = append(pairs, NewStringCodePair(1, defaultIfEmpty(h.FileName, ".")))

	// $ELEVATION
	pairs = append(pairs, NewStringCodePair(9, "$ELEVATION"))
	pairs = append(pairs, NewDoubleCodePair(40, h.Elevation))

	// $PELEVATION
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PELEVATION"))
		pairs = append(pairs, NewDoubleCodePair(40, h.PaperspaceElevation))
	}

	// $THICKNESS
	pairs = append(pairs, NewStringCodePair(9, "$THICKNESS"))
	pairs = append(pairs, NewDoubleCodePair(40, h.Thickness))

	// $LIMCHECK
	pairs = append(pairs, NewStringCodePair(9, "$LIMCHECK"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UseLimitsChecking)))

	// $BLIPMODE
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$BLIPMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.BlipMode)))
	}

	// $CHAMFERA
	pairs = append(pairs, NewStringCodePair(9, "$CHAMFERA"))
	pairs = append(pairs, NewDoubleCodePair(40, h.FirstChamferDistance))

	// $CHAMFERB
	pairs = append(pairs, NewStringCodePair(9, "$CHAMFERB"))
	pairs = append(pairs, NewDoubleCodePair(40, h.SecondChamferDistance))

	// $CHAMFERC
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$CHAMFERC"))
		pairs = append(pairs, NewDoubleCodePair(40, h.ChamferLength))
	}

	// $CHAMFERD
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$CHAMFERD"))
		pairs = append(pairs, NewDoubleCodePair(40, h.ChamferAngle))
	}

	// $SKPOLY
	pairs = append(pairs, NewStringCodePair(9, "$SKPOLY"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.PolylineSketchMode)))

	// $TDCREATE
	pairs = append(pairs, NewStringCodePair(9, "$TDCREATE"))
	pairs = append(pairs, NewDoubleCodePair(40, julianDateFromTime(h.CreationDate)))

	// $TDUCREATE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$TDUCREATE"))
		pairs = append(pairs, NewDoubleCodePair(40, julianDateFromTime(h.CreationDateUniversal)))
	}

	// $TDUPDATE
	pairs = append(pairs, NewStringCodePair(9, "$TDUPDATE"))
	pairs = append(pairs, NewDoubleCodePair(40, julianDateFromTime(h.UpdateDate)))

	// $TDUUPDATE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$TDUUPDATE"))
		pairs = append(pairs, NewDoubleCodePair(40, julianDateFromTime(h.UpdateDateUniversal)))
	}

	// $TDINDWG
	pairs = append(pairs, NewStringCodePair(9, "$TDINDWG"))
	pairs = append(pairs, NewDoubleCodePair(40, daysFromDuration(h.TimeInDrawing)))

	// $TDUSRTIMER
	pairs = append(pairs, NewStringCodePair(9, "$TDUSRTIMER"))
	pairs = append(pairs, NewDoubleCodePair(40, daysFromDuration(h.UserElapsedTimer)))

	// $USRTIMER
	pairs = append(pairs, NewStringCodePair(9, "$USRTIMER"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.UserTimerOn)))

	// $ANGBASE
	pairs = append(pairs, NewStringCodePair(9, "$ANGBASE"))
	pairs = append(pairs, NewDoubleCodePair(50, h.AngleZeroDirection))

	// $ANGDIR
	pairs = append(pairs, NewStringCodePair(9, "$ANGDIR"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.AngleDirection)))

	// $PDMODE
	pairs = append(pairs, NewStringCodePair(9, "$PDMODE"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.PointDisplayMode)))

	// $PDSIZE
	pairs = append(pairs, NewStringCodePair(9, "$PDSIZE"))
	pairs = append(pairs, NewDoubleCodePair(40, h.PointDisplaySize))

	// $PLINEWID
	pairs = append(pairs, NewStringCodePair(9, "$PLINEWID"))
	pairs = append(pairs, NewDoubleCodePair(40, h.DefaultPolylineWidth))

	// $COORDS
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$COORDS"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.CoordinateDisplay)))
	}

	// $SPLFRAME
	if h.Version <= R2013 {
		pairs = append(pairs, NewStringCodePair(9, "$SPLFRAME"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DisplaySplinePolygonControl)))
	}

	// $SPLINETYPE
	pairs = append(pairs, NewStringCodePair(9, "$SPLINETYPE"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.PEditSplineCurveType)))

	// $SPLINESEGS
	pairs = append(pairs, NewStringCodePair(9, "$SPLINESEGS"))
	pairs = append(pairs, NewShortCodePair(70, h.LineSegmentsPerSplinePatch))

	// $ATTDIA
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$ATTDIA"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.ShowAttributeEntryDialogs)))
	}

	// $ATTREQ
	if h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$ATTREQ"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.PromptForAttributeOnInsert)))
	}

	// $HANDLING
	if h.Version <= R12 {
		pairs = append(pairs, NewStringCodePair(9, "$HANDLING"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.HandlesEnabled)))
	}

	// $HANDSEED
	pairs = append(pairs, NewStringCodePair(9, "$HANDSEED"))
	pairs = append(pairs, NewStringCodePair(5, stringFromHandle(h.NextAvailableHandle)))

	// $SURFTAB1
	pairs = append(pairs, NewStringCodePair(9, "$SURFTAB1"))
	pairs = append(pairs, NewShortCodePair(70, h.MeshTabulationsInFirstDirection))

	// $SURFTAB2
	pairs = append(pairs, NewStringCodePair(9, "$SURFTAB2"))
	pairs = append(pairs, NewShortCodePair(70, h.MeshTabulationsInSecondDirection))

	// $SURFTYPE
	pairs = append(pairs, NewStringCodePair(9, "$SURFTYPE"))
	pairs = append(pairs, NewShortCodePair(70, int16(h.PEditSmoothSurfaceType)))

	// $SURFU
	pairs = append(pairs, NewStringCodePair(9, "$SURFU"))
	pairs = append(pairs, NewShortCodePair(70, h.PEditSmoothMDensith))

	// $SURFV
	pairs = append(pairs, NewStringCodePair(9, "$SURFV"))
	pairs = append(pairs, NewShortCodePair(70, h.PEditSmoothNDensith))

	// $UCSBASE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSBASE"))
		pairs = append(pairs, NewStringCodePair(2, h.UCSDefinitionName))
	}

	// $UCSNAME
	pairs = append(pairs, NewStringCodePair(9, "$UCSNAME"))
	pairs = append(pairs, NewStringCodePair(2, h.UCSName))

	// $UCSORG
	pairs = append(pairs, NewStringCodePair(9, "$UCSORG"))
	pairs = append(pairs, NewDoubleCodePair(10, h.UCSOrigin.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.UCSOrigin.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.UCSOrigin.Z))

	// $UCSXDIR
	pairs = append(pairs, NewStringCodePair(9, "$UCSXDIR"))
	pairs = append(pairs, NewDoubleCodePair(10, h.UCSXAxis.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.UCSXAxis.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.UCSXAxis.Z))

	// $UCSYDIR
	pairs = append(pairs, NewStringCodePair(9, "$UCSYDIR"))
	pairs = append(pairs, NewDoubleCodePair(10, h.UCSYAxis.X))
	pairs = append(pairs, NewDoubleCodePair(20, h.UCSYAxis.Y))
	pairs = append(pairs, NewDoubleCodePair(30, h.UCSYAxis.Z))

	// $UCSORTHOREF
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORTHOREF"))
		pairs = append(pairs, NewStringCodePair(2, h.OrthoUCSReference))
	}

	// $UCSORTHOVIEW
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORTHOVIEW"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.OrthgraphicViewType)))
	}

	// $UCSORGTOP
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGTOP"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginTop.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginTop.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginTop.Z))
	}

	// $UCSORGBOTTOM
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGBOTTOM"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginBottom.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginBottom.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginBottom.Z))
	}

	// $UCSORGLEFT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGLEFT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginLeft.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginLeft.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginLeft.Z))
	}

	// $UCSORGRIGHT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGRIGHT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginRight.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginRight.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginRight.Z))
	}

	// $UCSORGFRONT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGFRONT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginFront.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginFront.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginFront.Z))
	}

	// $UCSORGBACK
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$UCSORGBACK"))
		pairs = append(pairs, NewDoubleCodePair(10, h.UCSOriginBack.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.UCSOriginBack.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.UCSOriginBack.Z))
	}

	// $PUCSBASE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSBASE"))
		pairs = append(pairs, NewStringCodePair(2, h.PaperspaceUCSDefinitionName))
	}

	// $PUCSNAME
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSNAME"))
		pairs = append(pairs, NewStringCodePair(2, h.PaperspaceUCSName))
	}

	// $PUCSORG
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORG"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOrigin.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOrigin.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOrigin.Z))
	}

	// $PUCSXDIR
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSXDIR"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceXAxis.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceXAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceXAxis.Z))
	}

	// $PUCSYDIR
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSYDIR"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceYAxis.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceYAxis.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceYAxis.Z))
	}

	// $PUCSORTHOREF
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORTHOREF"))
		pairs = append(pairs, NewStringCodePair(2, h.PaperspaceOrthoUCSReference))
	}

	// $PUCSORTHOVIEW
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORTHOVIEW"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.PaperspaceOrthographicViewType)))
	}

	// $PUCSORGTOP
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGTOP"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginTop.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginTop.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginTop.Z))
	}

	// $PUCSORGBOTTOM
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGBOTTOM"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginBottom.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginBottom.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginBottom.Z))
	}

	// $PUCSORGLEFT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGLEFT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginLeft.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginLeft.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginLeft.Z))
	}

	// $PUCSORGRIGHT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGRIGHT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginRight.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginRight.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginRight.Z))
	}

	// $PUCSORGFRONT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGFRONT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginFront.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginFront.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginFront.Z))
	}

	// $PUCSORGBACK
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PUCSORGBACK"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceUCSOriginBack.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceUCSOriginBack.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceUCSOriginBack.Z))
	}

	// $USERI1
	pairs = append(pairs, NewStringCodePair(9, "$USERI1"))
	pairs = append(pairs, NewShortCodePair(70, h.UserInt1))

	// $USERI2
	pairs = append(pairs, NewStringCodePair(9, "$USERI2"))
	pairs = append(pairs, NewShortCodePair(70, h.UserInt2))

	// $USERI3
	pairs = append(pairs, NewStringCodePair(9, "$USERI3"))
	pairs = append(pairs, NewShortCodePair(70, h.UserInt3))

	// $USERI4
	pairs = append(pairs, NewStringCodePair(9, "$USERI4"))
	pairs = append(pairs, NewShortCodePair(70, h.UserInt4))

	// $USERI5
	pairs = append(pairs, NewStringCodePair(9, "$USERI5"))
	pairs = append(pairs, NewShortCodePair(70, h.UserInt5))

	// $USERR1
	pairs = append(pairs, NewStringCodePair(9, "$USERR1"))
	pairs = append(pairs, NewDoubleCodePair(40, h.UserReal1))

	// $USERR2
	pairs = append(pairs, NewStringCodePair(9, "$USERR2"))
	pairs = append(pairs, NewDoubleCodePair(40, h.UserReal2))

	// $USERR3
	pairs = append(pairs, NewStringCodePair(9, "$USERR3"))
	pairs = append(pairs, NewDoubleCodePair(40, h.UserReal3))

	// $USERR4
	pairs = append(pairs, NewStringCodePair(9, "$USERR4"))
	pairs = append(pairs, NewDoubleCodePair(40, h.UserReal4))

	// $USERR5
	pairs = append(pairs, NewStringCodePair(9, "$USERR5"))
	pairs = append(pairs, NewDoubleCodePair(40, h.UserReal5))

	// $WORLDVIEW
	pairs = append(pairs, NewStringCodePair(9, "$WORLDVIEW"))
	pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SetUCSToWCSInDViewOrVPoint)))

	// $SHADEDGE
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$SHADEDGE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.EdgeShading)))
	}

	// $SHADEDIF
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$SHADEDIF"))
		pairs = append(pairs, NewShortCodePair(70, h.PercentAmbientToDiffuse))
	}

	// $TILEMODE
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$TILEMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.PreviousReleaseTileCompatability)))
	}

	// $MAXACTVP
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$MAXACTVP"))
		pairs = append(pairs, NewShortCodePair(70, h.MaximumActiveViewports))
	}

	// $PINSBASE
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$PINSBASE"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceInsertionBase.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceInsertionBase.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceInsertionBase.Z))
	}

	// $PLIMCHECK
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PLIMCHECK"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.LimitCheckingInPaperspace)))
	}

	// $PEXTMIN
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PEXTMIN"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceMinimumDrawingExtents.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceMinimumDrawingExtents.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceMinimumDrawingExtents.Z))
	}

	// $PEXTMAX
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PEXTMAX"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceMaximumDrawingExtents.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceMaximumDrawingExtents.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.PaperspaceMaximumDrawingExtents.Z))
	}

	// $PLIMMIN
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PLIMMIN"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceMinimumDrawingLimits.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceMinimumDrawingLimits.Y))
	}

	// $PLIMMAX
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PLIMMAX"))
		pairs = append(pairs, NewDoubleCodePair(10, h.PaperspaceMaximumDrawingLimits.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.PaperspaceMaximumDrawingLimits.Y))
	}

	// $UNITMODE
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$UNITMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.DisplayFractionsInInput)))
	}

	// $VISRETAIN
	if h.Version >= R12 {
		pairs = append(pairs, NewStringCodePair(9, "$VISRETAIN"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.RetainXRefDependentVisibilitySettings)))
	}

	// $PLINEGEN
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PLINEGEN"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.IsPolylineContinuousAroundVerticies)))
	}

	// $PSLTSCALE
	if h.Version >= R11 {
		pairs = append(pairs, NewStringCodePair(9, "$PSLTSCALE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.ScaleLineTypesInPaperspace)))
	}

	// $TREEDEPTH
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$TREEDEPTH"))
		pairs = append(pairs, NewShortCodePair(70, h.SpacialIndexMaxDepth))
	}

	// $PICKSTYLE
	if h.Version >= R13 && h.Version <= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$PICKSTYLE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.PickStyle)))
	}

	// $CMLSTYLE
	if h.Version >= R13 && h.Version <= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$CMLSTYLE"))
		pairs = append(pairs, NewStringCodePair(7, h.CurrentMultilineStyle))
	}

	// $CMLSTYLE
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$CMLSTYLE"))
		pairs = append(pairs, NewStringCodePair(2, h.CurrentMultilineStyle))
	}

	// $CMLJUST
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$CMLJUST"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.CurrentMultilineJustification)))
	}

	// $CMLSCALE
	if h.Version >= R13 {
		pairs = append(pairs, NewStringCodePair(9, "$CMLSCALE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.CurrentMultilineScale))
	}

	// $PROXYGRAPHICS
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$PROXYGRAPHICS"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SaveProxyGraphics)))
	}

	// $MEASUREMENT
	if h.Version >= R14 {
		pairs = append(pairs, NewStringCodePair(9, "$MEASUREMENT"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DrawingUnits)))
	}

	// $CELWEIGHT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$CELWEIGHT"))
		pairs = append(pairs, NewShortCodePair(370, int16(h.NewObjectLineWeight)))
	}

	// $ENDCAPS
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$ENDCAPS"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.EndCapSetting)))
	}

	// $JOINSTYLE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$JOINSTYLE"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.LineweightJointSetting)))
	}

	// $LWDISPLAY
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$LWDISPLAY"))
		pairs = append(pairs, NewBoolCodePair(290, h.DisplayLinewieghtInModelAndLayoutTab))
	}

	// $INSUNITS
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$INSUNITS"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.DefaultDrawingUnits)))
	}

	// $HYPERLINKBASE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$HYPERLINKBASE"))
		pairs = append(pairs, NewStringCodePair(1, h.HyperlinkBase))
	}

	// $STYLESHEET
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$STYLESHEET"))
		pairs = append(pairs, NewStringCodePair(1, h.Stylesheet))
	}

	// $XEDIT
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$XEDIT"))
		pairs = append(pairs, NewBoolCodePair(290, h.CanUseInPlaceReferenceEditing))
	}

	// $CEPSNID
	if h.Version >= R2000 && h.NewObjectPlotStyleHandle != 0 {
		pairs = append(pairs, NewStringCodePair(9, "$CEPSNID"))
		pairs = append(pairs, NewStringCodePair(390, stringFromHandle(h.NewObjectPlotStyleHandle)))
	}

	// $CEPSNTYPE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$CEPSNTYPE"))
		pairs = append(pairs, NewShortCodePair(380, int16(h.NewObjectPlotStyle)))
	}

	// $PSTYLEMODE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PSTYLEMODE"))
		pairs = append(pairs, NewBoolCodePair(290, h.UsesColorDependentPlotStyleTables))
	}

	// $FINGERPRINTGUID
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$FINGERPRINTGUID"))
		pairs = append(pairs, NewStringCodePair(2, h.FingerprintGuid.String()))
	}

	// $VERSIONGUID
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$VERSIONGUID"))
		pairs = append(pairs, NewStringCodePair(2, h.VersionGuid.String()))
	}

	// $EXTNAMES
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$EXTNAMES"))
		pairs = append(pairs, NewBoolCodePair(290, h.UseACad2000SymbolTableNaming))
	}

	// $PSVPSCALE
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$PSVPSCALE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.ViewportViewScaleFactor))
	}

	// $OLESTARTUP
	if h.Version >= R2000 {
		pairs = append(pairs, NewStringCodePair(9, "$OLESTARTUP"))
		pairs = append(pairs, NewBoolCodePair(290, h.OleStartup))
	}

	// $SORTENTS
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$SORTENTS"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.ObjectSortingMethodsFlags)))
	}

	// $INDEXCTL
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$INDEXCTL"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.LayerAndSpatialIndexSaveMode)))
	}

	// $HIDETEXT
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$HIDETEXT"))
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(h.HideTextObjectsWhenProducintHiddenView)))
	}

	// $XCLIPFRAME
	if h.Version >= R2004 && h.Version <= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$XCLIPFRAME"))
		pairs = append(pairs, NewBoolCodePair(290, boolFromShort(int16(h.XRefClippingBoundaryVisible))))
	}

	// $XCLIPFRAME
	if h.Version >= R2010 {
		pairs = append(pairs, NewStringCodePair(9, "$XCLIPFRAME"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.XRefClippingBoundaryVisible)))
	}

	// $HALOGAP
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$HALOGAP"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.HaloGapPercent)))
	}

	// $OBSCOLOR
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$OBSCOLOR"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.ObscuredLineColor)))
	}

	// $OBSLTYPE
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$OBSLTYPE"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.ObscuredLineTypeStyle)))
	}

	// $INTERSECTIONDISPLAY
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$INTERSECTIONDISPLAY"))
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(h.DisplayIntersectionPolylines)))
	}

	// $INTERSECTIONCOLOR
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$INTERSECTIONCOLOR"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.IntersectionPolylineColor)))
	}

	// $DIMASSOC
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$DIMASSOC"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.DimensionObjectAssociativity)))
	}

	// $PROJECTNAME
	if h.Version >= R2004 {
		pairs = append(pairs, NewStringCodePair(9, "$PROJECTNAME"))
		pairs = append(pairs, NewStringCodePair(1, h.ProjectName))
	}

	// $CAMERADISPLAY
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$CAMERADISPLAY"))
		pairs = append(pairs, NewBoolCodePair(290, h.UseCameraDisplay))
	}

	// $LENSLENGTH
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LENSLENGTH"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LensLength))
	}

	// $CAMERAHEIGHT
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$CAMERAHEIGHT"))
		pairs = append(pairs, NewDoubleCodePair(40, h.CameraHeight))
	}

	// $STEPSPERSEC
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$STEPSPERSEC"))
		pairs = append(pairs, NewDoubleCodePair(40, h.StepsPerSecondInWalkOrFlyMode))
	}

	// $STEPSIZE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$STEPSIZE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.StepSizeInWalkOrFlyMode))
	}

	// $3DDWFPREC
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$3DDWFPREC"))
		pairs = append(pairs, NewDoubleCodePair(40, float64(int16(h.Dwf3DPrecision))))
	}

	// $PSOLWIDTH
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$PSOLWIDTH"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LastPolySolidWidth))
	}

	// $PSOLHEIGHT
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$PSOLHEIGHT"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LastPolySolidHeight))
	}

	// $LOFTANG1
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTANG1"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LoftOperationFirstDraftAngle))
	}

	// $LOFTANG2
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTANG2"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LoftOperationSecondDraftAngle))
	}

	// $LOFTMAG1
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTMAG1"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LoftOperationFirstMagnitude))
	}

	// $LOFTMAG2
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTMAG2"))
		pairs = append(pairs, NewDoubleCodePair(40, h.LoftOperationSecondMagnitude))
	}

	// $LOFTPARAM
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTPARAM"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.LoftFlags)))
	}

	// $LOFTNORMALS
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LOFTNORMALS"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.LoftedObjectNormalMode)))
	}

	// $LATITUDE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LATITUDE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.Latitude))
	}

	// $LONGITUDE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LONGITUDE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.Longitude))
	}

	// $NORTHDIRECTION
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$NORTHDIRECTION"))
		pairs = append(pairs, NewDoubleCodePair(40, h.AngleBetweenYAxisAndNorth))
	}

	// $TIMEZONE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$TIMEZONE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.TimeZone)))
	}

	// $LIGHTGLYPHDISPLAY
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$LIGHTGLYPHDISPLAY"))
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(h.UseLightGlyphDisplay)))
	}

	// $TILEMODELIGHTSYNCH
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$TILEMODELIGHTSYNCH"))
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(h.UseTileModeLightSync)))
	}

	// $CMATERIAL
	if h.Version >= R2007 && h.CurrentMaterialHandle != 0 {
		pairs = append(pairs, NewStringCodePair(9, "$CMATERIAL"))
		pairs = append(pairs, NewStringCodePair(347, stringFromHandle(h.CurrentMaterialHandle)))
	}

	// $SOLIDHIST
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$SOLIDHIST"))
		pairs = append(pairs, NewShortCodePair(280, shortFromBool(h.NewSolidsContainHistory)))
	}

	// $SHOWHIST
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$SHOWHIST"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.SolidHistoryMode)))
	}

	// $DWFFRAME
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DWFFRAME"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.UnderlayFrameMode)))
	}

	// $DGNFRAME
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$DGNFRAME"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.UnderlayFrameMode)))
	}

	// $REALWORLDSCALE
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$REALWORLDSCALE"))
		pairs = append(pairs, NewBoolCodePair(290, h.UseRealWorldScale))
	}

	// $INTERFERECOLOR
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$INTERFERECOLOR"))
		pairs = append(pairs, NewShortCodePair(62, int16(h.InterferenceObjectColor)))
	}

	// $INTERFEREOBJVS
	if h.Version >= R2007 && h.InterferenceObjectVisualStylePointer != 0 {
		pairs = append(pairs, NewStringCodePair(9, "$INTERFEREOBJVS"))
		pairs = append(pairs, NewStringCodePair(345, stringFromHandle(h.InterferenceObjectVisualStylePointer)))
	}

	// $INTERFEREVPVS
	if h.Version >= R2007 && h.InterferenceViewPortVisualStylePointer != 0 {
		pairs = append(pairs, NewStringCodePair(9, "$INTERFEREVPVS"))
		pairs = append(pairs, NewStringCodePair(346, stringFromHandle(h.InterferenceViewPortVisualStylePointer)))
	}

	// $CSHADOW
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$CSHADOW"))
		pairs = append(pairs, NewShortCodePair(280, int16(h.ShadowMode)))
	}

	// $SHADOWPLANELOCATION
	if h.Version >= R2007 {
		pairs = append(pairs, NewStringCodePair(9, "$SHADOWPLANELOCATION"))
		pairs = append(pairs, NewDoubleCodePair(40, h.ShadowPlaneZOffset))
	}

	// $AXISMODE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$AXISMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.AxisOn)))
	}

	// $AXISUNIT
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$AXISUNIT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.AxisTickSpacing.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.AxisTickSpacing.Y))
	}

	// $FASTZOOM
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$FASTZOOM"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.FastZoom)))
	}

	// $GRIDMODE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$GRIDMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.GridOn)))
	}

	// $GRIDUNIT
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$GRIDUNIT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.GridSpacing.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.GridSpacing.Y))
	}

	// $SNAPANG
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPANG"))
		pairs = append(pairs, NewDoubleCodePair(50, h.SnapRotationAngle))
	}

	// $SNAPBASE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPBASE"))
		pairs = append(pairs, NewDoubleCodePair(10, h.SnapBasePoint.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.SnapBasePoint.Y))
	}

	// $SNAPISOPAIR
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPISOPAIR"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.SnapIsometricPlane)))
	}

	// $SNAPMODE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPMODE"))
		pairs = append(pairs, NewShortCodePair(70, shortFromBool(h.SnapOn)))
	}

	// $SNAPSTYLE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPSTYLE"))
		pairs = append(pairs, NewShortCodePair(70, int16(h.SnapStyle)))
	}

	// $SNAPUNIT
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$SNAPUNIT"))
		pairs = append(pairs, NewDoubleCodePair(10, h.SnapSpacing.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.SnapSpacing.Y))
	}

	// $VIEWCTR
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$VIEWCTR"))
		pairs = append(pairs, NewDoubleCodePair(10, h.ViewCenter.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.ViewCenter.Y))
	}

	// $VIEWDIR
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$VIEWDIR"))
		pairs = append(pairs, NewDoubleCodePair(10, h.ViewDirection.X))
		pairs = append(pairs, NewDoubleCodePair(20, h.ViewDirection.Y))
		pairs = append(pairs, NewDoubleCodePair(30, h.ViewDirection.Z))
	}

	// $VIEWSIZE
	if h.Version <= R10 {
		pairs = append(pairs, NewStringCodePair(9, "$VIEWSIZE"))
		pairs = append(pairs, NewDoubleCodePair(40, h.ViewHeight))
	}

	err := writeSectionStart(writer, "HEADER")
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		err = writer.writeCodePair(pair)
		if err != nil {
			return err
		}
	}
	err = writeSectionEnd(writer)
	if err != nil {
		return err
	}
	return nil
}

func readHeader(nextPair CodePair, reader codePairReader) (Header, CodePair, error) {
	header := *NewHeader()
	var err error
	var variableName string
	for nextPair.Code != 0 {
		if nextPair.Code == 9 {
			variableName = nextPair.Value.(StringCodePairValue).Value
		} else {
			switch variableName {
			case "$ACADVER":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.Version = parseAcadVersion(nextPair.Value.(StringCodePairValue).Value)
				if header.Version >= R2007 {
					reader.setUtf8Reader()
				}
			case "$ACADMAINTVER":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.MaintenanceVersion = nextPair.Value.(ShortCodePairValue).Value
			case "$DWGCODEPAGE":
				if nextPair.Code != 3 {
					return header, nextPair, errors.New("expected code 3")
				}
				header.DrawingCodePage = nextPair.Value.(StringCodePairValue).Value
			case "$LASTSAVEDBY":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.LastSavedBy = nextPair.Value.(StringCodePairValue).Value
			case "$REQUIREDVERSIONS":
				if nextPair.Code != 160 {
					return header, nextPair, errors.New("expected code 160")
				}
				header.RequiredVersions = nextPair.Value.(LongCodePairValue).Value
			case "$INSBASE":
				switch nextPair.Code {
				case 10:
					header.InsertionBase.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.InsertionBase.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.InsertionBase.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$EXTMIN":
				switch nextPair.Code {
				case 10:
					header.MinimumDrawingExtents.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.MinimumDrawingExtents.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.MinimumDrawingExtents.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$EXTMAX":
				switch nextPair.Code {
				case 10:
					header.MaximumDrawingExtents.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.MaximumDrawingExtents.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.MaximumDrawingExtents.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$LIMMIN":
				switch nextPair.Code {
				case 10:
					header.MinimumDrawingLimits.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.MinimumDrawingLimits.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$LIMMAX":
				switch nextPair.Code {
				case 10:
					header.MaximumDrawingLimits.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.MaximumDrawingLimits.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$ORTHOMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DrawOrthoganalLines = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$REGENMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UseRegenMode = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$FILLMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.FillModeOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$QTEXTMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UseQuickTextMode = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$MIRRTEXT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.MirrorText = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DRAGMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DragMode = DragMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$LTSCALE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LineTypeScale = nextPair.Value.(DoubleCodePairValue).Value
			case "$OSMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ObjectSnapFlags = int(nextPair.Value.(ShortCodePairValue).Value)
			case "$ATTMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AttributeVisibility = AttributeVisibility(nextPair.Value.(ShortCodePairValue).Value)
			case "$TEXTSIZE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DefaultTextHeight = nextPair.Value.(DoubleCodePairValue).Value
			case "$TRACEWID":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.TraceWidth = nextPair.Value.(DoubleCodePairValue).Value
			case "$TEXTSTYLE":
				if nextPair.Code != 7 {
					return header, nextPair, errors.New("expected code 7")
				}
				header.TextStyle = nextPair.Value.(StringCodePairValue).Value
			case "$CLAYER":
				if nextPair.Code != 8 {
					return header, nextPair, errors.New("expected code 8")
				}
				header.CurrentLayer = nextPair.Value.(StringCodePairValue).Value
			case "$CELTYPE":
				if nextPair.Code != 6 {
					return header, nextPair, errors.New("expected code 6")
				}
				header.CurrentEntityLineType = nextPair.Value.(StringCodePairValue).Value
			case "$CECOLOR":
				if nextPair.Code != 62 {
					return header, nextPair, errors.New("expected code 62")
				}
				header.CurrentEntityColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$CELTSCALE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CurrentEntityLineTypeScale = nextPair.Value.(DoubleCodePairValue).Value
			case "$DELOBJ":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.RetainDeletedObjects = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DISPSILH":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DisplaySilhouetteCurvesInWireframeMode = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DRAGVS":
				if nextPair.Code != 349 {
					return header, nextPair, errors.New("expected code 349")
				}
				header.SolidVisualStylePointer = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$DIMSCALE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensioningScaleFactor = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMASZ":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensioningArrowSize = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMEXO":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionExtensionLineOffset = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMDLI":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionLineIncrement = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMRND":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionDistanceRoundingValue = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMDLE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionLineExtension = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMEXE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionExtensionLineExtension = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTP":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionPlusTolerance = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTM":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionMinusTolerance = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTXT":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensioningTextHeight = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMCEN":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CenterMarkSize = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTSZ":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensioningTickSize = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTOL":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.GenerateDimensionTolerances = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMLIM":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.GenerateDimensionLimits = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTIH":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextInsideHorizontal = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTOH":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextOutsideHorizontal = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMSE1":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SuppressFirstDimensionExtensionLine = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMSE2":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SuppressSecondDimensionExtensionLine = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTAD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.TextAboveDimensionLine = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMZIN":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionUnitZeroSuppression = UnitZeroSuppression(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMBLK":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.ArrowBlockName = nextPair.Value.(StringCodePairValue).Value
			case "$DIMASO":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.CreateAssociativeDimensioning = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMSHO":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.RecomputeDimensionsWhileDragging = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMPOST":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.DimensioningSuffix = nextPair.Value.(StringCodePairValue).Value
			case "$DIMAPOST":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.AlternateDimensioningSuffix = nextPair.Value.(StringCodePairValue).Value
			case "$DIMALT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UseAlternateDimensioning = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMALTD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AlternateDimensioningDecimalPlaces = nextPair.Value.(ShortCodePairValue).Value
			case "$DIMALTF":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.AlternateDimensioningScaleFactor = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMLFAC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionLinearMeasurementsScaleFactor = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTOFL":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ForceDimensionLineExtensionsOutsideIfTextIs = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTVP":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionVerticalTextPosition = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTIX":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ForceDimensionTextInsideExtensions = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMSOXD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SuppressOutsideExtensionDimensionLines = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMSAH":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UseSeparateArrowBlocksForDimensions = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMBLK1":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.FirstArrowBlockName = nextPair.Value.(StringCodePairValue).Value
			case "$DIMBLK2":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.SecondArrowBlockName = nextPair.Value.(StringCodePairValue).Value
			case "$DIMSTYLE":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.DimensionStyleName = nextPair.Value.(StringCodePairValue).Value
			case "$DIMCLRD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionLineColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMCLRE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionExtensionLineColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMCLRT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTFAC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionToleranceDisplayScaleFactor = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMGAP":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionLineGap = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMJUST":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextJustification = DimensionTextJustification(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTOLJ":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionToleranceVerticalJustification = Justification(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTZIN":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionToleranceZeroSuppression = UnitZeroSuppression(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMALTZ":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AlternateDimensioningZeroSupression = UnitZeroSuppression(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMALTTZ":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AlternateDimensioningToleranceZeroSupression = UnitZeroSuppression(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMFIT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextAndArrowPlacement = DimensionFit(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMUPT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionCursorControlsTextPosition = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMUNIT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionUnitFormat = UnitFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMDEC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionUnitToleranceDecimalPlaces = nextPair.Value.(ShortCodePairValue).Value
			case "$DIMTDEC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionToleranceDecimalPlaces = nextPair.Value.(ShortCodePairValue).Value
			case "$DIMALTU":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AlternateDimensioningUnits = UnitFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMALTTD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AlternateDimensioningToleranceDecimalPlaces = nextPair.Value.(ShortCodePairValue).Value
			case "$DIMTXSTY":
				if nextPair.Code != 7 {
					return header, nextPair, errors.New("expected code 7")
				}
				header.DimensionTextStyle = nextPair.Value.(StringCodePairValue).Value
			case "$DIMAUNIT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensioningAngleFormat = AngleFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMADEC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AngularDimensionPrecision = nextPair.Value.(ShortCodePairValue).Value
			case "$DIMALTRND":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.AlternateDimensioningUnitRounding = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMAZIN":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionAngleZeroSuppression = UnitZeroSuppression(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMDSEP":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionDecimalSeparatorRune = rune(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMFRAC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextHeightScaleFactor = DimensionFractionFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMLDRBLK":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.DimensionLeaderBlockName = nextPair.Value.(StringCodePairValue).Value
			case "$DIMLUNIT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionNonAngularUnits = NonAngularUnits(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMLWD":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionLineWeight = LineWeight(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMLWE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionExtensionLineWeight = LineWeight(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTMOVE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextMovementRule = DimensionTextMovementRule(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMFXL":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionLineFixedLength = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMFXLON":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionLineFixedLengthOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMJOGANG":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DimensionTransverseSegmentAngleInJoggedRadius = nextPair.Value.(DoubleCodePairValue).Value
			case "$DIMTFILL":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextBackgroundColorMode = DimensionTextBackgroundColorMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMTFILLCLR":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextBackgroundCustomColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMARCSYM":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionArcSymbolDisplayMode = DimensionArcSymbolDisplayMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMLTYPE":
				if nextPair.Code != 6 {
					return header, nextPair, errors.New("expected code 6")
				}
				header.DimensionLineType = nextPair.Value.(StringCodePairValue).Value
			case "$DIMLTEX1":
				if nextPair.Code != 6 {
					return header, nextPair, errors.New("expected code 6")
				}
				header.DimensionFirstExtensionLineType = nextPair.Value.(StringCodePairValue).Value
			case "$DIMLTEX2":
				if nextPair.Code != 6 {
					return header, nextPair, errors.New("expected code 6")
				}
				header.DimensionSecondExtensionLineType = nextPair.Value.(StringCodePairValue).Value
			case "$DIMTXTDIRECTION":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DimensionTextDirection = TextDirection(nextPair.Value.(ShortCodePairValue).Value)
			case "$LUNITS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UnitFormat = UnitFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$LUPREC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UnitPrecision = nextPair.Value.(ShortCodePairValue).Value
			case "$SKETCHINC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.SketchRecordIncrement = nextPair.Value.(DoubleCodePairValue).Value
			case "$FILLETRAD":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.FilletRadius = nextPair.Value.(DoubleCodePairValue).Value
			case "$AUNITS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AngleUnitFormat = AngleFormat(nextPair.Value.(ShortCodePairValue).Value)
			case "$AUPREC":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AngleUnitPrecision = nextPair.Value.(ShortCodePairValue).Value
			case "$MENU":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.FileName = nextPair.Value.(StringCodePairValue).Value
			case "$ELEVATION":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.Elevation = nextPair.Value.(DoubleCodePairValue).Value
			case "$PELEVATION":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.PaperspaceElevation = nextPair.Value.(DoubleCodePairValue).Value
			case "$THICKNESS":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.Thickness = nextPair.Value.(DoubleCodePairValue).Value
			case "$LIMCHECK":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UseLimitsChecking = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$BLIPMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.BlipMode = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$CHAMFERA":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.FirstChamferDistance = nextPair.Value.(DoubleCodePairValue).Value
			case "$CHAMFERB":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.SecondChamferDistance = nextPair.Value.(DoubleCodePairValue).Value
			case "$CHAMFERC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.ChamferLength = nextPair.Value.(DoubleCodePairValue).Value
			case "$CHAMFERD":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.ChamferAngle = nextPair.Value.(DoubleCodePairValue).Value
			case "$SKPOLY":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PolylineSketchMode = PolySketchMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$TDCREATE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CreationDate = timeFromJulianDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$TDUCREATE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CreationDateUniversal = timeFromJulianDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$TDUPDATE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UpdateDate = timeFromJulianDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$TDUUPDATE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UpdateDateUniversal = timeFromJulianDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$TDINDWG":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.TimeInDrawing = durationFromDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$TDUSRTIMER":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserElapsedTimer = durationFromDays(nextPair.Value.(DoubleCodePairValue).Value)
			case "$USRTIMER":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserTimerOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$ANGBASE":
				if nextPair.Code != 50 {
					return header, nextPair, errors.New("expected code 50")
				}
				header.AngleZeroDirection = nextPair.Value.(DoubleCodePairValue).Value
			case "$ANGDIR":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AngleDirection = AngleDirection(nextPair.Value.(ShortCodePairValue).Value)
			case "$PDMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PointDisplayMode = int(nextPair.Value.(ShortCodePairValue).Value)
			case "$PDSIZE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.PointDisplaySize = nextPair.Value.(DoubleCodePairValue).Value
			case "$PLINEWID":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.DefaultPolylineWidth = nextPair.Value.(DoubleCodePairValue).Value
			case "$COORDS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.CoordinateDisplay = CoordinateDisplay(nextPair.Value.(ShortCodePairValue).Value)
			case "$SPLFRAME":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DisplaySplinePolygonControl = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$SPLINETYPE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PEditSplineCurveType = PolylineCurvedAndSmoothSurfaceType(nextPair.Value.(ShortCodePairValue).Value)
			case "$SPLINESEGS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.LineSegmentsPerSplinePatch = nextPair.Value.(ShortCodePairValue).Value
			case "$ATTDIA":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ShowAttributeEntryDialogs = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$ATTREQ":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PromptForAttributeOnInsert = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$HANDLING":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.HandlesEnabled = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$HANDSEED":
				if nextPair.Code != 5 {
					return header, nextPair, errors.New("expected code 5")
				}
				header.NextAvailableHandle = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$SURFTAB1":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.MeshTabulationsInFirstDirection = nextPair.Value.(ShortCodePairValue).Value
			case "$SURFTAB2":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.MeshTabulationsInSecondDirection = nextPair.Value.(ShortCodePairValue).Value
			case "$SURFTYPE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PEditSmoothSurfaceType = PolylineCurvedAndSmoothSurfaceType(nextPair.Value.(ShortCodePairValue).Value)
			case "$SURFU":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PEditSmoothMDensith = nextPair.Value.(ShortCodePairValue).Value
			case "$SURFV":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PEditSmoothNDensith = nextPair.Value.(ShortCodePairValue).Value
			case "$UCSBASE":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.UCSDefinitionName = nextPair.Value.(StringCodePairValue).Value
			case "$UCSNAME":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.UCSName = nextPair.Value.(StringCodePairValue).Value
			case "$UCSORG":
				switch nextPair.Code {
				case 10:
					header.UCSOrigin.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOrigin.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOrigin.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSXDIR":
				switch nextPair.Code {
				case 10:
					header.UCSXAxis.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSXAxis.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSXAxis.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSYDIR":
				switch nextPair.Code {
				case 10:
					header.UCSYAxis.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSYAxis.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSYAxis.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORTHOREF":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.OrthoUCSReference = nextPair.Value.(StringCodePairValue).Value
			case "$UCSORTHOVIEW":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.OrthgraphicViewType = OrthographicViewType(nextPair.Value.(ShortCodePairValue).Value)
			case "$UCSORGTOP":
				switch nextPair.Code {
				case 10:
					header.UCSOriginTop.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginTop.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginTop.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORGBOTTOM":
				switch nextPair.Code {
				case 10:
					header.UCSOriginBottom.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginBottom.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginBottom.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORGLEFT":
				switch nextPair.Code {
				case 10:
					header.UCSOriginLeft.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginLeft.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginLeft.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORGRIGHT":
				switch nextPair.Code {
				case 10:
					header.UCSOriginRight.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginRight.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginRight.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORGFRONT":
				switch nextPair.Code {
				case 10:
					header.UCSOriginFront.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginFront.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginFront.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UCSORGBACK":
				switch nextPair.Code {
				case 10:
					header.UCSOriginBack.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.UCSOriginBack.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.UCSOriginBack.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSBASE":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.PaperspaceUCSDefinitionName = nextPair.Value.(StringCodePairValue).Value
			case "$PUCSNAME":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.PaperspaceUCSName = nextPair.Value.(StringCodePairValue).Value
			case "$PUCSORG":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOrigin.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOrigin.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOrigin.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSXDIR":
				switch nextPair.Code {
				case 10:
					header.PaperspaceXAxis.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceXAxis.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceXAxis.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSYDIR":
				switch nextPair.Code {
				case 10:
					header.PaperspaceYAxis.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceYAxis.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceYAxis.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORTHOREF":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.PaperspaceOrthoUCSReference = nextPair.Value.(StringCodePairValue).Value
			case "$PUCSORTHOVIEW":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PaperspaceOrthographicViewType = OrthographicViewType(nextPair.Value.(ShortCodePairValue).Value)
			case "$PUCSORGTOP":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginTop.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginTop.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginTop.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORGBOTTOM":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginBottom.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginBottom.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginBottom.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORGLEFT":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginLeft.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginLeft.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginLeft.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORGRIGHT":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginRight.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginRight.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginRight.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORGFRONT":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginFront.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginFront.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginFront.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PUCSORGBACK":
				switch nextPair.Code {
				case 10:
					header.PaperspaceUCSOriginBack.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceUCSOriginBack.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceUCSOriginBack.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$USERI1":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserInt1 = nextPair.Value.(ShortCodePairValue).Value
			case "$USERI2":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserInt2 = nextPair.Value.(ShortCodePairValue).Value
			case "$USERI3":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserInt3 = nextPair.Value.(ShortCodePairValue).Value
			case "$USERI4":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserInt4 = nextPair.Value.(ShortCodePairValue).Value
			case "$USERI5":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.UserInt5 = nextPair.Value.(ShortCodePairValue).Value
			case "$USERR1":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserReal1 = nextPair.Value.(DoubleCodePairValue).Value
			case "$USERR2":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserReal2 = nextPair.Value.(DoubleCodePairValue).Value
			case "$USERR3":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserReal3 = nextPair.Value.(DoubleCodePairValue).Value
			case "$USERR4":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserReal4 = nextPair.Value.(DoubleCodePairValue).Value
			case "$USERR5":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.UserReal5 = nextPair.Value.(DoubleCodePairValue).Value
			case "$WORLDVIEW":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SetUCSToWCSInDViewOrVPoint = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$SHADEDGE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.EdgeShading = ShadeEdgeMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$SHADEDIF":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PercentAmbientToDiffuse = nextPair.Value.(ShortCodePairValue).Value
			case "$TILEMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PreviousReleaseTileCompatability = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$MAXACTVP":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.MaximumActiveViewports = nextPair.Value.(ShortCodePairValue).Value
			case "$PINSBASE":
				switch nextPair.Code {
				case 10:
					header.PaperspaceInsertionBase.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceInsertionBase.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceInsertionBase.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PLIMCHECK":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.LimitCheckingInPaperspace = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$PEXTMIN":
				switch nextPair.Code {
				case 10:
					header.PaperspaceMinimumDrawingExtents.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceMinimumDrawingExtents.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceMinimumDrawingExtents.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PEXTMAX":
				switch nextPair.Code {
				case 10:
					header.PaperspaceMaximumDrawingExtents.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceMaximumDrawingExtents.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.PaperspaceMaximumDrawingExtents.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PLIMMIN":
				switch nextPair.Code {
				case 10:
					header.PaperspaceMinimumDrawingLimits.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceMinimumDrawingLimits.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$PLIMMAX":
				switch nextPair.Code {
				case 10:
					header.PaperspaceMaximumDrawingLimits.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.PaperspaceMaximumDrawingLimits.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$UNITMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DisplayFractionsInInput = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$VISRETAIN":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.RetainXRefDependentVisibilitySettings = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$PLINEGEN":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.IsPolylineContinuousAroundVerticies = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$PSLTSCALE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ScaleLineTypesInPaperspace = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$TREEDEPTH":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SpacialIndexMaxDepth = nextPair.Value.(ShortCodePairValue).Value
			case "$PICKSTYLE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.PickStyle = PickStyle(nextPair.Value.(ShortCodePairValue).Value)
			case "$CMLSTYLE":
				switch nextPair.Code {
				case 7:
					header.CurrentMultilineStyle = nextPair.Value.(StringCodePairValue).Value
				case 2:
					header.CurrentMultilineStyle = nextPair.Value.(StringCodePairValue).Value
				default:
					return header, nextPair, errors.New("expected codes 7, 2")
				}
			case "$CMLJUST":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.CurrentMultilineJustification = Justification(nextPair.Value.(ShortCodePairValue).Value)
			case "$CMLSCALE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CurrentMultilineScale = nextPair.Value.(DoubleCodePairValue).Value
			case "$PROXYGRAPHICS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SaveProxyGraphics = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$MEASUREMENT":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DrawingUnits = DrawingUnits(nextPair.Value.(ShortCodePairValue).Value)
			case "$CELWEIGHT":
				if nextPair.Code != 370 {
					return header, nextPair, errors.New("expected code 370")
				}
				header.NewObjectLineWeight = LineWeight(nextPair.Value.(ShortCodePairValue).Value)
			case "$ENDCAPS":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.EndCapSetting = EndCapSetting(nextPair.Value.(ShortCodePairValue).Value)
			case "$JOINSTYLE":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.LineweightJointSetting = JoinStyle(nextPair.Value.(ShortCodePairValue).Value)
			case "$LWDISPLAY":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.DisplayLinewieghtInModelAndLayoutTab = nextPair.Value.(BoolCodePairValue).Value
			case "$INSUNITS":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.DefaultDrawingUnits = Units(nextPair.Value.(ShortCodePairValue).Value)
			case "$HYPERLINKBASE":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.HyperlinkBase = nextPair.Value.(StringCodePairValue).Value
			case "$STYLESHEET":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.Stylesheet = nextPair.Value.(StringCodePairValue).Value
			case "$XEDIT":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.CanUseInPlaceReferenceEditing = nextPair.Value.(BoolCodePairValue).Value
			case "$CEPSNID":
				if nextPair.Code != 390 {
					return header, nextPair, errors.New("expected code 390")
				}
				header.NewObjectPlotStyleHandle = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$CEPSNTYPE":
				if nextPair.Code != 380 {
					return header, nextPair, errors.New("expected code 380")
				}
				header.NewObjectPlotStyle = PlotStyle(nextPair.Value.(ShortCodePairValue).Value)
			case "$PSTYLEMODE":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.UsesColorDependentPlotStyleTables = nextPair.Value.(BoolCodePairValue).Value
			case "$FINGERPRINTGUID":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.FingerprintGuid = uuidFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$VERSIONGUID":
				if nextPair.Code != 2 {
					return header, nextPair, errors.New("expected code 2")
				}
				header.VersionGuid = uuidFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$EXTNAMES":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.UseACad2000SymbolTableNaming = nextPair.Value.(BoolCodePairValue).Value
			case "$PSVPSCALE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.ViewportViewScaleFactor = nextPair.Value.(DoubleCodePairValue).Value
			case "$OLESTARTUP":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.OleStartup = nextPair.Value.(BoolCodePairValue).Value
			case "$SORTENTS":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.ObjectSortingMethodsFlags = int(nextPair.Value.(ShortCodePairValue).Value)
			case "$INDEXCTL":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.LayerAndSpatialIndexSaveMode = LayerAndSpatialIndexSaveMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$HIDETEXT":
				switch nextPair.Code {
				case 280:
					header.HideTextObjectsWhenProducintHiddenView = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
				case 290:
					header.HideTextObjectsWhenProducintHiddenView = nextPair.Value.(BoolCodePairValue).Value
				default:
					return header, nextPair, errors.New("expected codes 280, 290")
				}
			case "$XCLIPFRAME":
				switch nextPair.Code {
				case 290:
					header.XRefClippingBoundaryVisible = XrefClippingBoundaryVisibility(shortFromBool(nextPair.Value.(BoolCodePairValue).Value))
				case 280:
					header.XRefClippingBoundaryVisible = XrefClippingBoundaryVisibility(nextPair.Value.(ShortCodePairValue).Value)
				default:
					return header, nextPair, errors.New("expected codes 290, 280")
				}
			case "$HALOGAP":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.HaloGapPercent = float64(nextPair.Value.(ShortCodePairValue).Value)
			case "$OBSCOLOR":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.ObscuredLineColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$OBSLTYPE":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.ObscuredLineTypeStyle = LineTypeStyle(nextPair.Value.(ShortCodePairValue).Value)
			case "$INTERSECTIONDISPLAY":
				switch nextPair.Code {
				case 280:
					header.DisplayIntersectionPolylines = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
				case 290:
					header.DisplayIntersectionPolylines = nextPair.Value.(BoolCodePairValue).Value
				default:
					return header, nextPair, errors.New("expected codes 280, 290")
				}
			case "$INTERSECTIONCOLOR":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.IntersectionPolylineColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$DIMASSOC":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.DimensionObjectAssociativity = DimensionAssociativity(nextPair.Value.(ShortCodePairValue).Value)
			case "$PROJECTNAME":
				if nextPair.Code != 1 {
					return header, nextPair, errors.New("expected code 1")
				}
				header.ProjectName = nextPair.Value.(StringCodePairValue).Value
			case "$CAMERADISPLAY":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.UseCameraDisplay = nextPair.Value.(BoolCodePairValue).Value
			case "$LENSLENGTH":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LensLength = nextPair.Value.(DoubleCodePairValue).Value
			case "$CAMERAHEIGHT":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.CameraHeight = nextPair.Value.(DoubleCodePairValue).Value
			case "$STEPSPERSEC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.StepsPerSecondInWalkOrFlyMode = nextPair.Value.(DoubleCodePairValue).Value
			case "$STEPSIZE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.StepSizeInWalkOrFlyMode = nextPair.Value.(DoubleCodePairValue).Value
			case "$3DDWFPREC":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.Dwf3DPrecision = Dwf3DPrecision(int16(nextPair.Value.(DoubleCodePairValue).Value))
			case "$PSOLWIDTH":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LastPolySolidWidth = nextPair.Value.(DoubleCodePairValue).Value
			case "$PSOLHEIGHT":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LastPolySolidHeight = nextPair.Value.(DoubleCodePairValue).Value
			case "$LOFTANG1":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LoftOperationFirstDraftAngle = nextPair.Value.(DoubleCodePairValue).Value
			case "$LOFTANG2":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LoftOperationSecondDraftAngle = nextPair.Value.(DoubleCodePairValue).Value
			case "$LOFTMAG1":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LoftOperationFirstMagnitude = nextPair.Value.(DoubleCodePairValue).Value
			case "$LOFTMAG2":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.LoftOperationSecondMagnitude = nextPair.Value.(DoubleCodePairValue).Value
			case "$LOFTPARAM":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.LoftFlags = int(nextPair.Value.(ShortCodePairValue).Value)
			case "$LOFTNORMALS":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.LoftedObjectNormalMode = LoftedObjectNormalMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$LATITUDE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.Latitude = nextPair.Value.(DoubleCodePairValue).Value
			case "$LONGITUDE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.Longitude = nextPair.Value.(DoubleCodePairValue).Value
			case "$NORTHDIRECTION":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.AngleBetweenYAxisAndNorth = nextPair.Value.(DoubleCodePairValue).Value
			case "$TIMEZONE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.TimeZone = TimeZone(nextPair.Value.(ShortCodePairValue).Value)
			case "$LIGHTGLYPHDISPLAY":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.UseLightGlyphDisplay = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$TILEMODELIGHTSYNCH":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.UseTileModeLightSync = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$CMATERIAL":
				if nextPair.Code != 347 {
					return header, nextPair, errors.New("expected code 347")
				}
				header.CurrentMaterialHandle = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$SOLIDHIST":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.NewSolidsContainHistory = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$SHOWHIST":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.SolidHistoryMode = SolidHistoryMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$DWFFRAME":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.UnderlayFrameMode = UnderlayFrameMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$REALWORLDSCALE":
				if nextPair.Code != 290 {
					return header, nextPair, errors.New("expected code 290")
				}
				header.UseRealWorldScale = nextPair.Value.(BoolCodePairValue).Value
			case "$INTERFERECOLOR":
				if nextPair.Code != 62 {
					return header, nextPair, errors.New("expected code 62")
				}
				header.InterferenceObjectColor = Color(nextPair.Value.(ShortCodePairValue).Value)
			case "$INTERFEREOBJVS":
				if nextPair.Code != 345 {
					return header, nextPair, errors.New("expected code 345")
				}
				header.InterferenceObjectVisualStylePointer = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$INTERFEREVPVS":
				if nextPair.Code != 346 {
					return header, nextPair, errors.New("expected code 346")
				}
				header.InterferenceViewPortVisualStylePointer = handleFromString(nextPair.Value.(StringCodePairValue).Value)
			case "$CSHADOW":
				if nextPair.Code != 280 {
					return header, nextPair, errors.New("expected code 280")
				}
				header.ShadowMode = ShadowMode(nextPair.Value.(ShortCodePairValue).Value)
			case "$SHADOWPLANELOCATION":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.ShadowPlaneZOffset = nextPair.Value.(DoubleCodePairValue).Value
			case "$AXISMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.AxisOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$AXISUNIT":
				switch nextPair.Code {
				case 10:
					header.AxisTickSpacing.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.AxisTickSpacing.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$FASTZOOM":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.FastZoom = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$GRIDMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.GridOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$GRIDUNIT":
				switch nextPair.Code {
				case 10:
					header.GridSpacing.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.GridSpacing.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$SNAPANG":
				if nextPair.Code != 50 {
					return header, nextPair, errors.New("expected code 50")
				}
				header.SnapRotationAngle = nextPair.Value.(DoubleCodePairValue).Value
			case "$SNAPBASE":
				switch nextPair.Code {
				case 10:
					header.SnapBasePoint.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.SnapBasePoint.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$SNAPISOPAIR":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SnapIsometricPlane = SnapIsometricPlane(nextPair.Value.(ShortCodePairValue).Value)
			case "$SNAPMODE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SnapOn = boolFromShort(nextPair.Value.(ShortCodePairValue).Value)
			case "$SNAPSTYLE":
				if nextPair.Code != 70 {
					return header, nextPair, errors.New("expected code 70")
				}
				header.SnapStyle = SnapStyle(nextPair.Value.(ShortCodePairValue).Value)
			case "$SNAPUNIT":
				switch nextPair.Code {
				case 10:
					header.SnapSpacing.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.SnapSpacing.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$VIEWCTR":
				switch nextPair.Code {
				case 10:
					header.ViewCenter.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.ViewCenter.Y = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$VIEWDIR":
				switch nextPair.Code {
				case 10:
					header.ViewDirection.X = nextPair.Value.(DoubleCodePairValue).Value
				case 20:
					header.ViewDirection.Y = nextPair.Value.(DoubleCodePairValue).Value
				case 30:
					header.ViewDirection.Z = nextPair.Value.(DoubleCodePairValue).Value
				}
			case "$VIEWSIZE":
				if nextPair.Code != 40 {
					return header, nextPair, errors.New("expected code 40")
				}
				header.ViewHeight = nextPair.Value.(DoubleCodePairValue).Value
			default:
				// ignore unsupported header variable
			}
		}

		nextPair, err = reader.readCodePair()
		if err != nil {
			return header, nextPair, err
		}
	}

	return header, nextPair, nil
}
