package main
import "fmt"

func main() {
	var matricula string //primeiro valor da linha
	var p1, p2, p3, p4, p5, p6, p7, p8, l1, l2, l3, l4, l5, f, pr float64 //próximos valores da linha, p = prova, l = lista, f = trabalho final, pr = presença
	var mp, ml, nota_final float64 //médias
	var situação string
	
	fmt.Scanln(&matricula, &p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &l1, &l2, &l3, &l4, &l5, &f, &pr)
	
	if(matricula != "-1") {
    	mp = (p1+p2+p3+p4+p5+p6+p7+p8)/8
    	ml = (l1+l2+l3+l4+l5)/5
    	nota_final = mp*0.7+ml*0.15+f*0.15
    	
    	if(nota_final >= 6 && pr >= 96 ) {
    	    situação = "APROVADO"
    	}else if(nota_final >= 6 && pr < 96) {
    	    situação = "REPROVADO POR FREQUENCIA"
    	}else if(nota_final < 6 && pr >= 96) {
    	    situação = "REPROVADO POR NOTA"
    	}else if(nota_final < 6 && pr < 96) {
    	    situação = "REPROVADO POR NOTA E POR FREQUENCIA"
    	}
    	
    	fmt.Printf("Matrícula: %v, Nota Final: %0.2f, Situação final: %v \n", matricula, nota_final, situação)
    	main()
    }
    
}