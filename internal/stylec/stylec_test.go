package stylec

import (
	"testing"
)

func TestCompileTokens(t *testing.T) {
	tokens, err := Compile(`.hero {
  font-size: 44;
  bold: true;
  color: #2f6fed;
}
.dim { color: gray; }`)
	if err != nil {
		t.Fatalf("Compile 失败: %v", err)
	}
	hero := tokens["hero"]
	if hero == nil {
		t.Fatal("缺 .hero")
	}
	if hero["font-size"] != 44.0 {
		t.Errorf("font-size = %v", hero["font-size"])
	}
	if hero["bold"] != true {
		t.Errorf("bold = %v", hero["bold"])
	}
	if hero["color"] != "#2f6fed" {
		t.Errorf("color = %v", hero["color"])
	}
	if tokens["dim"]["color"] != "gray" {
		t.Errorf("dim.color = %v", tokens["dim"]["color"])
	}
}

func TestRejectAtRule(t *testing.T) {
	if _, err := Compile("@import url(x.css);"); err == nil {
		t.Fatal("@import 应拒绝")
	}
}

func TestRejectGlobalSelector(t *testing.T) {
	for _, src := range []string{"body { color: red; }", ".a, .b { color: red; }", ".a > .b { color: red; }", "* { color: red; }"} {
		if _, err := Compile(src); err == nil {
			t.Errorf("%q 应拒绝(只允许单个 .class)", src)
		}
	}
}

func TestRejectURL(t *testing.T) {
	if _, err := Compile(".a { color: url(http://x); }"); err == nil {
		t.Fatal("url() 应拒绝")
	}
}

func TestRejectUnknownProp(t *testing.T) {
	if _, err := Compile(".a { transform: scale(2); }"); err == nil {
		t.Fatal("白名单外属性应拒绝")
	}
}

func TestRejectDuplicateRule(t *testing.T) {
	if _, err := Compile(".a { color: red; } .a { color: blue; }"); err == nil {
		t.Fatal("重复规则应拒绝")
	}
}
