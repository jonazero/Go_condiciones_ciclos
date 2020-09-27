package main
import "fmt"
func main(){
	var dia int;
	var mes int;
	fmt.Scan(&dia);
	fmt.Scan(&mes);
	if mes == 12{
		if dia < 22{
			fmt.Println("Sagitario")
		}else{
			fmt.Println("Capriconio")
		}		
	} else if mes == 1{
		if dia < 20{
			fmt.Println("Capricornio")
		}else{
			fmt.Println("Acuario")
		}
	}else if mes == 2{
		if dia < 19{
			fmt.Println("Acuario")
		}else{
			fmt.Println("Piscis")
		}
	}else if mes == 3{
		if dia < 21{
			fmt.Println("Piscis")
		}else{
			fmt.Println("Aries")
		}	
	}else if mes == 4{
		if dia < 20{
			fmt.Println("Aries")
		}else{
			fmt.Println("Tauro")
		}
	}else if mes == 5{
		if dia < 21{
			fmt.Println("Tauro")
		}else{
			fmt.Println("Geminis")
		}
	}else if mes == 6{
		if dia < 21{
			fmt.Println("Geminis")
		}else{
			fmt.Println("Cancer")
		}
	}else if mes == 7{
		if dia < 23{
			fmt.Println("Cancer")
		}else{
			fmt.Println("Leo")
		}
	}else if mes == 8{
		if dia < 23{
			fmt.Println("Leo")
		}else{
			fmt.Println("Virgo")
		}
	}else if mes == 9{
		if dia < 23{
			fmt.Println("Virgo")
		}else{
			fmt.Println("Libra")
		}
	}else if mes == 10{
		if dia < 23{
			fmt.Println("Libra")
		}else{
			fmt.Println("Scorpio")
		}
	}else if mes == 11{
		if dia < 22{
			fmt.Println("Scorpio")
		}else{
			fmt.Println("Sagitario")
		}
	}

}

