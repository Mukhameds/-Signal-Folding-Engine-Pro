package output

import (
	"fmt"
	"os"
	"signal_folding_pro/core"
)

func ExportPDB(p *core.Protein, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("[ERROR] creating PDB:", err)
		return
	}
	defer f.Close()

	for i, r := range p.Chain {
		recordCA := fmt.Sprintf("ATOM  %5d  CA  %-3s A%4d    %8.3f%8.3f%8.3f  1.00  0.00           C\n",
			i+1, r.Type, i+1, r.X_CA, r.Y_CA, r.Z_CA)
		recordN := fmt.Sprintf("ATOM  %5d  N   %-3s A%4d    %8.3f%8.3f%8.3f  1.00  0.00           N\n",
			i+1, r.Type, i+1, r.X_N, r.Y_N, r.Z_N)
		recordC := fmt.Sprintf("ATOM  %5d  C   %-3s A%4d    %8.3f%8.3f%8.3f  1.00  0.00           C\n",
			i+1, r.Type, i+1, r.X_C, r.Y_C, r.Z_C)

		f.WriteString(recordN)
		f.WriteString(recordCA)
		f.WriteString(recordC)
	}
	f.WriteString("END\n")
}
