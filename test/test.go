package main

import (
    "fmt"
    "os"
    "time"
    "os/exec"
)

func clear() { 
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// clear["windows"] = func() {
// 	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }
// type intVal struct{
//     stat int
//     const name string
// }
// type floatVal struct{
//     stat float32
//     const name string
// }
// type Base struct{
// 	att intVal;
// 	balance intVal;
// 	speed intVal;
// 	addmg intVal;
// 	crit intVal;
// 	cdmg floatVal;
// }
// type Stats struct{	//dodelat'!!!!!
// 	// att_modifier int;
// 	balance_surplus intVal;
// 	balance_perc floatVal;
// 	addmg_perc floatVal;

// 	animation_speed floatVal;

// 	realcrit intVal;
// 	new_critDmg floatVal;
// 	crit_mod floatVal;

// 	base_stats Base;
// }


func cmov(col int,line int){
	fmt.Printf("\033[%d;%dH", line, col) // Set cursor position
}

type val_pos struct{
    val int
    y int //Y offset
    x int
}
///UNUSED
func returnMaps()(*map[string]val_pos,*map[string]val_pos){
    m:=map[string]val_pos{
        //"ATT": 10000,
        //                  {value,y,x}(y,x == print position)
        "Additional Damage": {2750,0,4},
        "Balance": {90,1,4},
        "Speed": {102,2,4},
        "Crit Chance": {115,3,4},
        "Crit Damage": {210,4,4},

        "dps":{0,5,4}

        "balance surplus": {0, 6,4},  
        "balance percent": {0,7,4},   // 0,1 - 1
        "adddmg percent": {0, 8,4},       
        "animation speed": {0, 9,4},  
        "real crit": {0,10,4},         // CRIT<=50
        "new crit damage": {0,11,4},   // crit damage with balance surplus
        "crit mod": {0, 12,4},         //dmg modifier from crits
    }
    n:=make(map[string]val_pos)
    for key,val:= range m{
        val.x=40
        n[key]=val
    }

    return &m,&n
}

func printStat(val int,name string,x int, y int){
    const name_offset =6

    cmov(x              ,y)
    fmt.Print(val)
    cmov(x+name_offset  ,y)
    fmt.Print(name)
}

func draw_stuff_offset(arr []int)(int, int){
    if(len(arr)==0){
        return 0,0
    }else if(len(arr)>1){    //len > 1 (2,3,4,5..)
        return arr[0],arr[1]
    }
    return arr[0],0             //if len=1
}

func printHeader(header string){
    lm:=23 //additional damage + 6 offset
    lh:=len(header)
    if(lh>=lm){
        fmt.Print(header)
        return
    }
    prinStart:=lm/2-lh/2
    i:=0
    for ;i<prinStart;i++{
        fmt.Print("_")
    }
    fmt.Print(header)
    for i+=lh+1;i<=lm;i++{
        fmt.Print("_")
    }
    //  0 1 2 3 4 5 6 7 8 9 
    //  A d d i t i o n a l
    //        M a i n
}

func draw_stuff(m map[string]*val_pos,header string, offset ...int){
    x,y:=draw_stuff_offset(offset)  //min x=1 , y=1
    //x+=3    //artificial offset
   
    cmov(x+m["Additional Damage"].x,   m["Additional Damage"].y-1)
    printHeader(header)
    //y++
    //add debug option in future (print all)
    
    for key, value := range m{
        if(key[0]<91){  //Print Only A-Z
            printStat(value.val,key,    x+value.x, y+value.y)
        }
    }
}

func debug(a int,b int){
    cmov(25,17)
    fmt.Print(a,b)
}

func calcDPS(m map[string]&val_pos){
        "Additional Damage": {2750,0,4},
        "Balance": {90,1,4},
        "Speed": {102,2,4},
        "Crit Chance": {115,3,4},
        "Crit Damage": {210,4,4},

        "dps":{0,5,4}
        
        "balance surplus": {0, 6,4},  
        "balance percent": {0,7,4},   // 0,1 - 1
        "adddmg percent": {0, 8,4},       
        "animation speed": {0, 9,4},  
        "real crit": {0,10,4},         // CRIT<=50
        "new crit damage": {0,11,4},   // crit damage with balance surplus
        "crit mod": {0, 12,4},         //dmg modifier from crits

    const addmg_const float32 = 0.0005740528129
    m["balance surplus"].val=   m["Balance"].val-m["Balance"].val%90
    m["balance percent"].val=   1000-(100-m["Balance"].val%90*1000)/2
    m["adddmg percent"].val =   int(float32(m["Additional Damage"].val)*addmg_const*1000)
    m["animation speed"].val=   int(1/((200+float32(m["Speed"].val))/200))
    m["real crit"].val      =   crit // CRIT<=50
    m["new crit damage"].val=
    m["crit mod"].val       =


    bal_perc:=
        dps:=finalStats.balance_perc+finalStats.addmg_perc
        //fmt.Println("DPS_perc: ",dps)
        dps=dps*finalStats.crit_mod
        //fmt.Println("mult: ",finalStats.crit_mod/finalStats.animation_speed)
        return dps/finalStats.animation_speed
    
}

func returnArray()(*[2][13]val_pos){
    var tmp =[13]val_pos{
    //{value,y,x}(y,x == print position)
        {2750,0,0},
        {90,1,0},
        {102,2,0},
        {115,3,0},
        {210,4,0},

        {0,5,0},

        {0, 6,0},  
       	{0,7,0},   // 0,1 - 1
        {0, 8,0},       
        {0, 9,0},  
        {0,10,0},         // CRIT<=50
        {0,11,0},   // crit damage with balance surplus
        {0, 12,0},         //dmg modifier from crits
    }
    var arr [2][13]val_pos
    arr[0]=tmp
    arr[1]=tmp

    for index := range arr[0]{
        arr[0][index].x=5
        arr[0][index].y=index+3
        arr[1][index].x=40
        arr[1][index].y=index+3
    }
    return &arr
}

func returnMap(arr *[13]val_pos)map[string]*val_pos{
	return map[string]*val_pos{
        "Additional Damage": &arr[0],
        "Balance": &arr[1],
        "Speed": &arr[2],
        "Crit Chance": &arr[3],
        "Crit Damage": &arr[4],

        "dps":&arr[5],

        "balance surplus": &arr[6],  
        "balance percent": &arr[7],   // 0,1 - 1
        "adddmg percent": &arr[8],       
        "animation speed": &arr[9],  
        "real crit": &arr[10],         // CRIT<=50
        "new crit damage": &arr[11],   // crit damage with balance surplus
        "crit mod": &arr[12],         //dmg modifier from crits
    }
}

var blink byte
func drawArrow(pos val_pos){
    offset:=3-int(blink%3)
    cmov(pos.x-offset,pos.y)
    arrow:=">"
    fmt.Print(arrow)
}

func handle_input(mov string,arr *[2][13]val_pos){
    
    switch mov[0] {
    case 'w':
        validate_arrow_y(-1)
    case 's':
        validate_arrow_y(1)
    case 'a':
        validate_arrow_x(-1)
    case 'd':
        validate_arrow_x(1)
    case 'e':
        set_number(&arr[user.x][user.y].val)
    case '0':
        accumulate_number(mov[0])
    case '1':
        accumulate_number(mov[0])
    case '2':
        accumulate_number(mov[0])
    case '3':
        accumulate_number(mov[0])
    case '4':
        accumulate_number(mov[0])
    case '5':
        accumulate_number(mov[0])
    case '6':
        accumulate_number(mov[0])
    case '7':
        accumulate_number(mov[0])
    case '8':
        accumulate_number(mov[0])
    case '9':
        accumulate_number(mov[0])
    default:
        return
    }
}

func set_number(num *int){
    fmt.Print(num)
    if(accum_number>=0){
        *num =accum_number
    }    
    accum_number=-1
}
func accumulate_number(by byte){
    deltha:=int(by-'0')
    if(accum_number<0){
        accum_number=deltha
        return
    }
    accum_number=accum_number*10+deltha
}
var accum_number int= -1

type arrow_pos struct{
    y int
    x int
}

func validate_arrow_x(movement int){
    accum_number=-1
    tmp:=user.x+movement
    if(tmp<0){
        return
    }
    if(tmp>1){
        return
    }
    user.x=tmp
}
func validate_arrow_y(movement int){
    accum_number=-1
    tmp:=user.y+movement
    if(tmp<0){
        return
    }
    if(tmp>4){
        return
    }
    user.y=tmp
}

var user =arrow_pos{x:0,y:0}
func main() {
    
    mainArr_ptr:=   returnArray()
    mainMap_ptr,compMap_ptr:=   returnMap(&mainArr_ptr[0]),returnMap(&mainArr_ptr[1])//returnMap(compArr_ptr)

    clear()


    draw_stuff(mainMap_ptr,"Main_Stats")
   
    draw_stuff(compMap_ptr,"Secondary_Stats")
    
    drawArrow(mainArr_ptr[user.x][user.y])
  debug(user.x,user.y)
fmt.Print("\n",mainArr_ptr[user.x][user.y])
    cmov(0,23)
   //return
    ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 1)
        for {
            os.Stdin.Read(b)
            ch <- string(b)
        }
        defer  exec.Command("stty", "-F", "/dev/tty", "echo").Run() ///??? doesnt work

    }(ch)

    for {
        select {
            case stdin, _ := <-ch:
                cmov(0,23)
                fmt.Print("Keys pressed:", stdin)
                handle_input(stdin,mainArr_ptr)
            default:
                
				
            //     fmt.Println("Working..")
        }
        draw_stuff(mainMap_ptr,"Main_Stats")
        draw_stuff(compMap_ptr,"Secondary_Stats")
        drawArrow(mainArr_ptr[user.x][user.y])
        cmov(0,23)
        fmt.Print(accum_number)
        time.Sleep(time.Millisecond * 175)
        clear()
		blink++
    }

}			