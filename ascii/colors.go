package ascii

import (
	"strings"
)

func getColor(name string) (string, bool) {
	name = strings.ToLower(name)
	colorMap := map[string]string{
		"Aliceblue":            "240;248;255",
		"Antiquewhite":         "250;235;215",
		"Aqua":                 "0;255;255",
		"Aquamarine":           "127;255;212",
		"Azure":                "240;255;255",
		"Beige":                "245;245;220",
		"Bisque":               "255;228;196",
		"Black":                "0;0;0",
		"Blanchedalmond":       "255;235;205",
		"Blue":                 "0;0;255",
		"Blueviolet":           "138;43;226",
		"Brown":                "165;42;42",
		"Burlywood":            "222;184;135",
		"Cadetblue":            "95;158;160",
		"Chartreuse":           "127;255;0",
		"Chocolate":            "210;105;30",
		"Coral":                "255;127;80",
		"Cornflowerblue":       "100;149;237",
		"Cornsilk":             "255;248;220",
		"Crimson":              "220;20;60",
		"Cyan":                 "0;255;255",
		"Darkblue":             "0;0;139",
		"Darkcyan":             "0;139;139",
		"Darkgoldenrod":        "184;134;11",
		"Darkgray":             "169;169;169",
		"Darkgreen":            "0;100;0",
		"Darkgrey":             "169;169;169",
		"Darkkhaki":            "189;183;107",
		"Darkmagenta":          "139;0;139",
		"Darkolivegreen":       "85;107;47",
		"Darkorange":           "255;140;0",
		"Darkorchid":           "153;50;204",
		"Darkred":              "139;0;0",
		"Darksalmon":           "233;150;122",
		"Darkseagreen":         "143;188;143",
		"Darkslateblue":        "72;61;139",
		"Darkslategray":        "47;79;79",
		"Darkslategrey":        "47;79;79",
		"Darkturquoise":        "0;206;209",
		"Darkviolet":           "148;0;211",
		"Deeppink":             "255;20;147",
		"Deepskyblue":          "0;191;255",
		"Dimgray":              "105;105;105",
		"Dimgrey":              "105;105;105",
		"Dodgerblue":           "30;144;255",
		"Firebrick":            "178;34;34",
		"Floralwhite":          "255;250;240",
		"Forestgreen":          "34;139;34",
		"Fuchsia":              "255;0;255",
		"Gainsboro":            "220;220;220",
		"Ghostwhite":           "248;248;255",
		"Gold":                 "255;215;0",
		"Goldenrod":            "218;165;32",
		"Gray":                 "128;128;128",
		"Green":                "0;128;0",
		"Greenyellow":          "173;255;47",
		"Grey":                 "128;128;128",
		"Honeydew":             "240;255;240",
		"Hotpink":              "255;105;180",
		"Indianred":            "205;92;92",
		"Indigo":               "75;0;130",
		"Ivory":                "255;255;240",
		"Khaki":                "240;230;140",
		"Lavender":             "230;230;250",
		"Lavenderblush":        "255;240;245",
		"Lawngreen":            "124;252;0",
		"Lemonchiffon":         "255;250;205",
		"Lightblue":            "173;216;230",
		"Lightcoral":           "240;128;128",
		"Lightcyan":            "224;255;255",
		"Lightgoldenrodyellow": "250;250;210",
		"Lightgray":            "211;211;211",
		"Lightgreen":           "144;238;144",
		"Lightgrey":            "211;211;211",
		"Lightpink":            "255;182;193",
		"Lightsalmon":          "255;160;122",
		"Lightseagreen":        "32;178;170",
		"Lightskyblue":         "135;206;250",
		"Lightslategray":       "119;136;153",
		"Lightslategrey":       "119;136;153",
		"Lightsteelblue":       "176;196;222",
		"Lightyellow":          "255;255;224",
		"Lime":                 "0;255;0",
		"Limegreen":            "50;205;50",
		"Linen":                "250;240;230",
		"Magenta":              "255;0;255",
		"Maroon":               "128;0;0",
		"Mediumaquamarine":     "102;205;170",
		"Mediumblue":           "0;0;205",
		"Mediumorchid":         "186;85;211",
		"Mediumpurple":         "147;112;219",
		"Mediumseagreen":       "60;179;113",
		"Mediumslateblue":      "123;104;238",
		"Mediumspringgreen":    "0;250;154",
		"Mediumturquoise":      "72;209;204",
		"Mediumvioletred":      "199;21;133",
		"Midnightblue":         "25;25;112",
		"Mintcream":            "245;255;250",
		"Mistyrose":            "255;228;225",
		"Moccasin":             "255;228;181",
		"Navajowhite":          "255;222;173",
		"Navy":                 "0;0;128",
		"Oldlace":              "253;245;230",
		"Olive":                "128;128;0",
		"Olivedrab":            "107;142;35",
		"Orange":               "255;165;0",
		"Orangered":            "255;69;0",
		"Orchid":               "218;112;214",
		"Palegoldenrod":        "238;232;170",
		"Palegreen":            "152;251;152",
		"Paleturquoise":        "175;238;238",
		"Palevioletred":        "219;112;147",
		"Papayawhip":           "255;239;213",
		"Peachpuff":            "255;218;185",
		"Peru":                 "205;133;63",
		"Pink":                 "255;192;203",
		"Plum":                 "221;160;221",
		"Powderblue":           "176;224;230",
		"Purple":               "128;0;128",
		"Red":                  "255;0;0",
		"Rosybrown":            "188;143;143",
		"Royalblue":            "65;105;225",
		"Saddlebrown":          "139;69;19",
		"Salmon":               "250;128;114",
		"Sandybrown":           "244;164;96",
		"Seagreen":             "46;139;87",
		"Seashell":             "255;245;238",
		"Sienna":               "160;82;45",
		"Silver":               "192;192;192",
		"Skyblue":              "135;206;235",
		"Slateblue":            "106;90;205",
		"Slategray":            "112;128;144",
		"Slategrey":            "112;128;144",
		"Snow":                 "255;250;250",
		"Springgreen":          "0;255;127",
		"Steelblue":            "70;130;180",
		"Tan":                  "210;180;140",
		"Teal":                 "0;128;128",
		"Thistle":              "216;191;216",
		"Tomato":               "255;99;71",
		"Turquoise":            "64;224;208",
		"Violet":               "238;130;238",
		"Wheat":                "245;222;179",
		"White":                "255;255;255",
		"Whitesmoke":           "245;245;245",
		"Yellow":               "255;255;0",
		"Yellowgreen":          "154;205;50",
	}
	if len(colorMap[strings.Title(name)]) == 0 {
		return "", false
	}
	return colorMap[strings.Title(name)], true
}
