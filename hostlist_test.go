package hostlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type testPair struct {
	input  string
	output []string
}

var testPairsNode = []testPair{
	{"n01p[001-005]", []string{"n01p001", "n01p002", "n01p003", "n01p004", "n01p005"}},
	{"ap[1-3]z", []string{"ap1z", "ap2z", "ap3z"}},
	{"dgx[01-03]", []string{"dgx01", "dgx02", "dgx03"}},
	{"adev[06,13,15]", []string{"adev06", "adev13", "adev15"}},
	{"lx[062-064,128]", []string{"lx062", "lx063", "lx064", "lx128"}},
	{"s02p[017,029-031]", []string{"s02p017", "s02p029", "s02p030", "s02p031"}},
	{"adev[009-101]", []string{"adev009", "adev010", "adev011", "adev012", "adev013", "adev014", "adev015", "adev016", "adev017", "adev018", "adev019", "adev020", "adev021", "adev022", "adev023", "adev024", "adev025", "adev026", "adev027", "adev028", "adev029", "adev030", "adev031", "adev032", "adev033", "adev034", "adev035", "adev036", "adev037", "adev038", "adev039", "adev040", "adev041", "adev042", "adev043", "adev044", "adev045", "adev046", "adev047", "adev048", "adev049", "adev050", "adev051", "adev052", "adev053", "adev054", "adev055", "adev056", "adev057", "adev058", "adev059", "adev060", "adev061", "adev062", "adev063", "adev064", "adev065", "adev066", "adev067", "adev068", "adev069", "adev070", "adev071", "adev072", "adev073", "adev074", "adev075", "adev076", "adev077", "adev078", "adev079", "adev080", "adev081", "adev082", "adev083", "adev084", "adev085", "adev086", "adev087", "adev088", "adev089", "adev090", "adev091", "adev092", "adev093", "adev094", "adev095", "adev096", "adev097", "adev098", "adev099", "adev100", "adev101"}},
	{"s02p[044,046,049]", []string{"s02p044", "s02p046", "s02p049"}},
	{"nodewithoutnumber", []string{"nodewithoutnumber"}},
	{"n04p[036,043-044,046-047]", []string{"n04p036", "n04p043", "n04p044", "n04p046", "n04p047"}},
	{"n05p[036,043-044],n06p[046-047]", []string{"n05p036", "n05p043", "n05p044", "n06p046", "n06p047"}},
	{"x[1-2]y[1-2][1-2]", []string{"x1y11", "x1y12", "x1y21", "x1y22", "x2y11", "x2y12", "x2y21", "x2y22"}},
}

func TestExpandNodeList(t *testing.T) {
	for _, pair := range testPairsNode {
		resultHostlist, err := ExpandHostlist(pair.input, false, false)
		if err != nil {
			t.Error(err)
		}
		// t.Log("input:", pair.input)
		// t.Log("output:", resultHostlist)
		if diff := cmp.Diff(resultHostlist, pair.output, cmpopts.SortSlices(func(x, y string) bool { return x < y })); diff != "" {
			t.Errorf("ExpandNodeList(%q) mismatch (-want +got):\n%s", pair.input, diff)
		}
	}
}
