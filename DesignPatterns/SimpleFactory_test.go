package DesignPatterns

import "testing"

func TestGirlFactory_CreateGirl(t *testing.T) {
	factor := &GirlFactory{}

	Fat := factor.CreateGirl("fat")
	Fat.weight()
	Thin := factor.CreateGirl("thin")
	Thin.weight()
}
