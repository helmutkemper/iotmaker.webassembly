// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lighting-color
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lighting-color
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="diffuseLighting1" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting in="SourceGraphic" lighting-color="white">
//        <fePointLight x="60" y="60" z="20" />
//      </feDiffuseLighting>
//    </filter>
//    <filter id="diffuseLighting2" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting in="SourceGraphic" lighting-color="blue">
//        <fePointLight x="60" y="60" z="20" />
//      </feDiffuseLighting>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#diffuseLighting1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#diffuseLighting2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("diffuseLighting1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).LightingColor(factoryColor.NewWhite()).Append(
				factoryBrowser.NewTagSvgFePointLight().X(60).Y(60).Z(20),
			),
		),

		factoryBrowser.NewTagSvgFilter().Id("diffuseLighting2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).LightingColor(factoryColor.NewBlue()).Append(
				factoryBrowser.NewTagSvgFePointLight().X(60).Y(60).Z(20),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#diffuseLighting1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#diffuseLighting2); transform: translateX(220px);"),
	)

	stage.Append(s1)

	<-done
}

//
//
//
//
//
//
//
//
//
//
//
//
//
