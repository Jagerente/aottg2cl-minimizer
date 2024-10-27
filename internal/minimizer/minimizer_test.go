package minimizer

import (
	"strings"
	"testing"
)

func TestMinimizerAllInOne(t *testing.T) {
	m := New()

	input := `
		class Main
		{
			foo = faz;
			bar = a + b;

			/* Multiline comment
			inside code */

			# Single-line comment
			if (foo == faz || bar != baz) # Comment
			{
				foo += 1;
				bar -= 2; /* Comment */
				bar *= 3;
				bar /= 4;
			}
			elif (foo > faz && bar < baz)
			{
				foo = faz;
			}
			else
			{
				bar = baz;
			}

			while (foo <= faz)
			{
				foo += 1;
			}

			for (i in Range(0, 10, 1))
			{
				bar = bar * i;
			}

			foo = foo + (bar - baz) * (faz / 1);
		}
		
		extension Ext
		{
			foo = faz;
			bar = a + b;

			/* Multiline comment
			inside code */

			# Single-line comment
			if (foo == faz || bar != baz) # Comment
			{
				foo += 1;
				bar -= 2; /* Comment */
				bar *= 3;
				bar /= 4;
			}
			elif (foo > faz && bar < baz)
			{
				foo = faz;
			}
			else
			{
				bar = baz;
			}

			while (foo <= faz)
			{
				foo += 1;
			}

			for (i in Range(0, 10, 1))
			{
				bar = bar * i;
			}

			foo = foo + (bar - baz) * (faz / 1);
		}
	`

	expected := "class Main{foo=faz;bar=a+b;if(foo==faz||bar!=baz){foo+=1;bar-=2;bar*=3;bar/=4;}elif(foo>faz&&bar<baz){foo=faz;}else{bar=baz;}while(foo<=faz){foo+=1;}for(i in Range(0,10,1)){bar=bar*i;}foo=foo+(bar - baz)*(faz/1);}extension Ext{foo=faz;bar=a+b;if(foo==faz||bar!=baz){foo+=1;bar-=2;bar*=3;bar/=4;}elif(foo>faz&&bar<baz){foo=faz;}else{bar=baz;}while(foo<=faz){foo+=1;}for(i in Range(0,10,1)){bar=bar*i;}foo=foo+(bar - baz)*(faz/1);}"

	output, err := m.minimize([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output = strings.TrimSpace(output)

	if output != expected {
		t.Errorf("test failed: expected %q, got %q", expected, output)
	}
}
