package main

import ("fmt"
"strconv")

const dmg_const = (274.0/623.0/750.0);

type Base struct{
	att int;
	balance int;
	speed int;
	addmg int;
	crit int;
	cdmg float32;
}

type Stats struct{	//dodelat'!!!!!
	// att_modifier int;
	balance_surplus int;
	balance_perc float32;
	addmg_perc float32;

	animation_speed float32;

	realcrit int;
	new_critDmg float32;
	crit_mod float32;
	

	base_stats Base;
}

func (scroll *Base)setDefaultCharVal(){
	scroll.att =10000;
	scroll.balance =90;
	scroll.speed =102;
	scroll.addmg =2750;
	scroll.crit =115;
	scroll.cdmg =2.1; 
}

func (summ *Base)addBaseVal(scroll Base){
	summ.att +=scroll.att;
	summ.balance +=scroll.balance;
	summ.speed +=scroll.speed;
	summ.addmg +=scroll.addmg;
	summ.crit +=scroll.crit;
	summ.cdmg  +=scroll.cdmg;
} 

func getBalanceSurp(bal int)int{
	if bal = bal-90; bal > 0 {
		return bal
	}
	return 0
}

func getRealBalance(bal int)int{
	if bal > 90 {
		return 90
	}
	return bal
}

var critMax=50;
func getRealCrit(crit int)int{
	if crit > critMax {
		return critMax
	}
	return crit
}

func (totalStats *Stats)SummScrollValStats(scrolls ...Base){
	for j := range scrolls { 
        totalStats.base_stats.addBaseVal(scrolls[j]) 
    } 
	totalStats.addmg_perc=float32(totalStats.base_stats.addmg)*dmg_const*100.0
	totalStats.balance_surplus= getBalanceSurp(totalStats.base_stats.balance)
	totalStats.balance_perc=float32(getRealBalance(totalStats.base_stats.balance)+100)/2
	totalStats.animation_speed= 1/((200+float32(totalStats.base_stats.speed))/200)
	totalStats.realcrit=getRealCrit(totalStats.base_stats.crit)
	totalStats.new_critDmg= totalStats.base_stats.cdmg+(float32(totalStats.balance_surplus)*0.01/3)
	totalStats.crit_mod=(float32(totalStats.realcrit)*(totalStats.new_critDmg-1))/100+1
}

func printBase(base Base){
	fmt.Println("Balance:\t ",base.balance)
	fmt.Println("Speed:\t\t ",base.speed)
	fmt.Println("Aditional damage:",base.addmg)
	fmt.Println("Critical chance: ",base.crit)
	fmt.Println("Crit Damage: \t ",base.cdmg)
}

func printStats(base Stats){
	fmt.Println("new_critdmg:\t",base.new_critDmg)
	fmt.Println("real_critrate:\t",base.realcrit)
	fmt.Println("animation_speed:",base.animation_speed)
	fmt.Println("Crit_mod:\t",base.crit_mod)
	fmt.Println("balance_surplus:",base.balance_surplus)
	
	fmt.Println("addmg_perc: \t",base.addmg_perc)
	fmt.Println("balance_perc: \t",base.balance_perc)
	fmt.Println("new_Damage: \t",base.new_critDmg)
	fmt.Println("\n\tBase stats: \t")
	printBase(base.base_stats)
}

func finalDPS(finalStats Stats)float32{
	dps:=finalStats.balance_perc+finalStats.addmg_perc
	//fmt.Println("DPS_perc: ",dps)
	dps=dps*finalStats.crit_mod
	//fmt.Println("mult: ",finalStats.crit_mod/finalStats.animation_speed)
	return dps/finalStats.animation_speed
}

func exit(i string){
	if(i[0]=='q'||i[0]=='Q'||i[0]=='E'||i[0]=='q'){
		panic("+bye")
	}
}

func NewInputInt(num *int){
	float:=float32(*num)
	NewInputFloat(&float)
	*num=int(float)
}

func NewInputFloat(float *float32){
   // string to float 
   var key_input string
   	for{
		fmt.Scanln(&key_input)
		if(len(key_input)==0){
			return
		}
		exit(key_input)
		i, err := strconv.ParseFloat(key_input, 32)
		if err != nil {
			fmt.Println("Something went WRONG, you can print this in Discord if it require a fix")
			fmt.Print("Do it again: ") 
			//panic(err)
		}else{
			*float=float32(i)
			return
		}
	}
}

func (output *Base)inputNewBase(){
	fmt.Printf("Enter your Balance[%v]:",output.balance)
	NewInputInt(&output.balance)
	fmt.Printf("Enter your Speed[%v]:",output.speed)
	NewInputInt(&output.speed)
	fmt.Printf("Enter your Additional Damage[%v]:",output.addmg)
	NewInputInt(&output.addmg)
	fmt.Printf("Enter your Crit Chance[%v]:",output.crit)
	NewInputInt(&output.crit)
	fmt.Printf("Enter your Crit Damage[%v]:",output.cdmg)
	NewInputFloat(&output.cdmg)
   	for(output.cdmg>=3 || output.cdmg<1.5){
		if(output.cdmg>=3 || output.cdmg<1.5){
			fmt.Printf("Your Crit Damage (%v) is weird, try again:",output.cdmg)
			NewInputFloat(&output.cdmg)
   		}
	}
}

func (output *Base) inputES(str string){
	fmt.Printf("Enter %s Balance[%v]:",str,output.balance)
	NewInputInt(&output.balance)
	fmt.Printf("Enter %s Speed[%v]:",str,output.speed)
	NewInputInt(&output.speed)
	fmt.Printf("Enter %s Additional Damage[%v]:",str,output.addmg)
	NewInputInt(&output.addmg)
	fmt.Printf("Enter %s Crit Chance[%v]:",str,output.crit)
	NewInputInt(&output.crit)
	// fmt.Printf("Enter %s Crit Damage:")
	// output.cdmg=NewInputFloat(&)
}

func ask(question string)bool{
	var key_input string
	//>> ADD fmt print here with question
	fmt.Printf("%s\n(Y/N/Q) ",question)
	fmt.Scanln(&key_input)
	if(len(key_input)==0){
		return false
	}
	if(key_input[0]=='y'||key_input[0]=='Y'){
		return true
	}
	exit(key_input)
	return false
}

func main(){
	var (
		totalStats Stats
		character_stats Base
		Empty_base Base
		ES1 Base
		ES2 Base
	)
	character_stats.setDefaultCharVal();
	
	//totalStats.SummScrollValStats(character_stats,ES1,ES2)
	totalStats.SummScrollValStats(character_stats)
	printStats(totalStats)
	fmt.Printf("\nYour Dps score with default stats is:\n\n\t> %.1f%s <\n\n\n",finalDPS(totalStats),"%")
	var naked_DPS,scrolled_DPS float32
	var scrolls_changed bool
	var IncDec string ="ERROR"
	for{
		ES1,ES2,scrolls_changed=Empty_base,Empty_base,false
		
		if(ask("do You wish to edit your character stats?")){
			//You wish to change default stats
			character_stats.inputNewBase()
		}

		if(ask("do You wish to add Enchant scroll?")){
			ES1.inputES("1st ES")
	
			if(ask("do You wish to add another Enchant scroll?")){
				ES2.inputES("2st ES")
			}
			scrolls_changed=true
		}

		totalStats.base_stats=character_stats
		naked_DPS=finalDPS(totalStats)
		totalStats.SummScrollValStats(ES1,ES2)
		scrolled_DPS=finalDPS(totalStats)
		fmt.Printf("\nYour Dps score with those stats is:\n\n\t> %.1f%s <\n\n\n",scrolled_DPS,"%")
		if(scrolls_changed){
			if(scrolled_DPS-naked_DPS<=0){
				if(scrolled_DPS==naked_DPS){
					IncDec="(no diff)"
				}else{
					IncDec="decrease"
				}
			}else{
				IncDec="increase"
			}
			fmt.Printf("DPS diff: %.2f%s\n\n",((scrolled_DPS/naked_DPS-1)*100),"%")
		}else{
			fmt.Println()
		}
	}
}