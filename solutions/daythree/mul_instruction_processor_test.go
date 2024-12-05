package daythree

import "testing"

func TestMulInstructionProcessor(t *testing.T) {
	tests := []struct {
		line          string
		expectedCount int
	}{
		/*
		 * mul instruction
		 */
		{"mumfordandsons2,3)", 0},

		{"mul(0,0)", 1},
		{"mul(1,2)", 1},
		{"mul(12,3)", 1},
		{"mul(123,4)", 1},

		{"mul(1,23)", 1},
		{"mul(1,234)", 1},

		{"mul(12,34)", 1},
		{"mul(123,45)", 1},

		{"mul(12,345)", 1},
		{"mul(123,456)", 1},
		{"mul(1234,5678)", 1},

		{`mul(485,849)%'&/} #why()mul(337,670)}&when(292,584)mul(291,636)$<,}mul(873,787)mul(583,597)from()%-{-(+!do()]?mul(508,735)where()+<[*,[^#mul(356,730) why()^what()@?where(558,563)!mul(254,728)mul(573,97)'when(){[!mul(847,358)>+-%[mul(445,707)~<+< *&who(192,588)/mul(339,368)why()?>how())*&when() mul(807,43)don't()$@%who();&(how()mul(719,162)#{(from(),;mul:~-where()%who()how(84,44),mul(970,95)where()(/]how():mul(223[;'?why()!/how()@@who()mul(507,673))^mul(977,524)from()select()^}{#)when()*mul(454,303)%select()*@,mul(563,289)*}who()<}<;why()&mul(673,375)^where()~mul(637,689)'*^who()?don't()who()where()<mul(729,329),?+mul(452,183)where();mul(412,329)select()&{^select()mul(304,748/-<~{$'$mul(250,311)^[!*@don't()-]%when()who()mul(275,612)mul(808,359)%:?how()'/$what()mul(125,6)>mul(347,707)[#mul(714,896)]mul(249,387)mul(991,762)from()when()where()how(553,916)-mul(116,473)& {who()mul(934,426)#)}from()mul(305,521)why(583,561)-;<}/how()mulwhat()<@&mul(667,932)%select()%mul(196,687))select()*do()from()/what()select()}mul(141,484)+]from()!'^how()who()[$mul(493,253):)&<mul(106,339)[#}[mul(527,610)/@&,#)/;who())mul(945,633)#who()!&,when(),]mul(699,592)*%mul(533,983)#:where()who()>$;[<]mul(663,997) {#@~:#-!mul(869,651)?*@/<#--do()mul(584,519)%select()what()who()-(!@how()don't()<,[!!;where()<^)mul(670,564);$what(),}mul(584,531)-(*mul(727,546)$:how(410,917)select()mul(447,208)'who()( ')&mul(476,15)how()where()^what()mul(332,458)`, 53},
	}

	for _, tt := range tests {
		mulInstructionProcessor := NewMulInstructionProcessor()

		t.Run("should find expected mul instructions", func(t *testing.T) {
			mulInstructionProcessor.Collect(tt.line)

			instructionsLen := len(*mulInstructionProcessor.Instructions)

			if instructionsLen != tt.expectedCount {
				t.Errorf("Expected %d instructions, got %d", tt.expectedCount, instructionsLen)
			}
		})
	}
}
