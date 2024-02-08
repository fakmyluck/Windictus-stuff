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

func returnInputInt()int{
	// // string to int
	// var key_input string
	// for{
	// 	fmt.Scan(&key_input)
	// 	i, err := strconv.Atoi(key_input)
	// 	if err != nil {
	// 		fmt.Println("Something went WRONG, you can print this in Discord if it require a fix")
	// 		fmt.Print("Do it again: ") 
	// 		//panic(err)
	// 	}else{
	// 		return i
	// 	}
	// }
	return int(returnInputFloat())
}

func exit(i string){
	if(i[0]=='q'||i[0]=='Q'||i[0]=='E'||i[0]=='q'){
		panic("+bye")
	}
}

func returnInputFloat()float32{
   // string to float 
   var key_input string
   	for{
		fmt.Scan(&key_input)
		exit(key_input)
		i, err := strconv.ParseFloat(key_input, 32)
		if err != nil {
			fmt.Println("Something went WRONG, you can print this in Discord if it require a fix")
			fmt.Print("Do it again: ") 
			//panic(err)
		}else{
			return float32(i)
		}
	}
}

func (output *Base)inputNewBase(){
	fmt.Print("Enter your Balance:")
	output.balance=returnInputInt()
	fmt.Print("Enter your Speed:")
	output.speed=returnInputInt()
	fmt.Print("Enter your Additional Damage:")
	output.addmg=returnInputInt()
	fmt.Print("Enter your Crit Chance:")
	output.crit=returnInputInt()
	fmt.Print("Enter your Crit Damage:")
	output.cdmg=0
   	for(output.cdmg>=3 || output.cdmg<1.5){
		output.cdmg=returnInputFloat()
		if(output.cdmg>=3 || output.cdmg<1.5){
			fmt.Printf("Your Crit Damage (%v) is weird:",output.cdmg)
   		}
	}
}

func (output *Base) inputES(str string){
	fmt.Printf("Enter %s Balance:",str)
	output.balance=returnInputInt()
	fmt.Printf("Enter %s Speed:",str)
	output.speed=returnInputInt()
	fmt.Printf("Enter %s Additional Damage:",str)
	output.addmg=returnInputInt()
	fmt.Printf("Enter %s Crit Chance:",str)
	output.crit=returnInputInt()
	// fmt.Printf("Enter %s Crit Damage:")
	// output.cdmg=returnInputFloat()
}

func main(){
	var (
		totalStats Stats
		character_stats Base
		Empty_base Base
		ES1 Base
		ES2 Base
		// ES1 =Base{
		// 	speed: -34,
		// 	balance:0,
		// 	addmg: 600,
		// 	crit: 10,
		// 	cdmg: 0}
		// // ES2 =Base{
		// // 	speed: -48,
		// // 	balance: 9,
		// // 	addmg: 850,
		// // 	crit: 0,
		// // 	cdmg: 0,
		// // }
	)
	character_stats.setDefaultCharVal();
	
	//totalStats.SummScrollValStats(character_stats,ES1,ES2)
	totalStats.SummScrollValStats(character_stats)
	printStats(totalStats)
	fmt.Printf("\nYour Dps score with default stats is:\n\n\t> %.1f%s <\n\n\n",finalDPS(totalStats),"%")

	var key_input string
	for{
		ES1,ES2=Empty_base,Empty_base
		fmt.Print("do You wish to edit your character stats?\n(Y/N/Q) ")
		fmt.Scan(&key_input)
		if(key_input[0]=='y'||key_input[0]=='y'){
			//You wish to change default stats
			character_stats.inputNewBase()
			//fmt.Println("\n",character_stats)
		}
		exit(key_input)

		fmt.Print("do You wish to add Enchant scroll?\n(Y/N/Q) ")
		fmt.Scan(&key_input)
		if(key_input[0]=='y'||key_input[0]=='y'){
			ES1.inputES("1st ES")
		
			fmt.Println("do You wish to add another Enchant scroll?\n(Y/N/Q) ")
			fmt.Scan(&key_input)
			if(key_input[0]=='y'||key_input[0]=='y'){
				ES1.inputES("2st ES")
			}
			exit(key_input)
		}
		exit(key_input)

		totalStats.base_stats=character_stats
		totalStats.SummScrollValStats(ES1,ES2)
		printStats(totalStats)
		fmt.Printf("\nYour Dps score with those stats is:\n\n\t> %.1f%s <\n\n\n",finalDPS(totalStats),"%")
	}
}