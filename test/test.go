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

        "dps":{0,5,4},

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
    const name_offset =10

    cmov(x              ,y)
    fmt.Print(val)
    cmov(x+name_offset  ,y)
    fmt.Print(name)
}

func printDPS(val int,x int,y int){
    const name_offset =10

    cmov(x              ,y)
    tmp:=val%1000/10
    if(tmp%10>=5){
        tmp+=10
    }
    tmp=tmp/10
    if(tmp%10==0){
        tmp=tmp/10
    }
    
    fmt.Printf("%v,%v%s",val/1000,tmp,"%")
    cmov(x+name_offset  ,y)
    fmt.Print("Damage Potential")
}

func drawDif(main val_pos, sec val_pos){
    const offset =40
    

    intHax:=main.val*100000/sec.val

    str:=intToString(intHax-100000)
    cmov(offset-len(str)/2     ,main.y+2)
    fmt.Print(str)
}

func intToString(num int)string{
    num=num/10+num%10/5
    num=num/10+num%10/5
 
    if(num==0){
        return "No Difference"
    }
    if(num>999){
        return "Secondary suck"
    }
    var sign string
    if(num<0){
        sign="-"
        num=num*-1
    }else{
        sign="+"
    }
    
    byt:=byte(num%10)
    dec:=[]byte{byt/10+'0',byt%10+'0'}        // <0
    byt=byte(num/10)
    digit:=[]byte{byt/10+'0',byt%10+'0'}      // >0
    if(digit[0]=='0'){
        digit=digit[1:2]    // ummm.... yes?
    }
    return sign+string(digit)+"."+string(dec)+"%"
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
    lm:=27 //additional damage + 10 offset
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
        // if(key[0]<91){
        if(key[0]<91){  //Print Only A-Z
            printStat(value.val,key,    x+value.x, y+value.y)
        }
        // }else{
        //     printStat(value.val,key,    x+value.x, y+value.y+1)
        // }
        if(key=="dps"){
            printDPS(value.val,    x+value.x, y+value.y+1)
        }
    }
}

func debug(a int,b int){
    cmov(25,17)
    fmt.Print(a,b)
}

func calcDPS(m map[string]*val_pos){
    const ten=100000
    //const addmg_const  = 0.0005740528129
    const addmg_const = 5740528129

    tmp:=0
    if(m["Balance"].val>90){
        tmp=m["Balance"].val-90
    }
    m["balance surplus"].val=   tmp
    m["real balance"].val   =   m["Balance"].val-m["balance surplus"].val
    m["balance percent"].val=   (100-(100-m["real balance"].val)/2)  *ten /100         //           <100
    m["adddmg percent"].val =   m["Additional Damage"].val*addmg_const/1000/ten    // /100000
    m["animation speed"].val=   ten*ten/(ten*(200+m["Speed"].val)/200)             // /100000
    m["real crit"].val      =   m["Crit Chance"].val
    if(m["Crit Chance"].val>50){
        m["real crit"].val  =   50
    }
    m["new crit damage"].val=   m["Crit Damage"].val+m["balance surplus"].val/3
    m["crit mod"].val       =   (m["real crit"].val*(m["new crit damage"].val-100)+10000)*10 // /100000
    dps:=(m["balance percent"].val+m["adddmg percent"].val)*m["crit mod"].val

    m["dps"].val=dps/m["animation speed"].val
    
    
}

func returnArray()(*[2][14]val_pos){
    var tmp =[14]val_pos{
    //{value,y,x}(y,x == print position)
        {2750,0,0},
        {103,1,0},
        {102,2,0},
        {115,3,0},
        {210,4,0},

        {0, 5,0},  
       	{0,6,0},   // 0,1 - 1
        {0, 7,0},       
        {0, 8,0},  
        {0,9,0},         // CRIT<=50
        {0,10,0},   // crit damage with balance surplus
        {0, 11,0},         //dmg modifier from crits
        {0,12,0},
        {0,13,0},
    }
    var arr [2][14]val_pos
    arr[0]=tmp
    arr[1]=tmp

    for index := range arr[0]{
        arr[0][index].x=5
        arr[0][index].y=index+3
        arr[1][index].x=50
        arr[1][index].y=index+3
    }
    return &arr
}

func returnMap(arr *[14]val_pos)map[string]*val_pos{
	return map[string]*val_pos{
        "Additional Damage": &arr[0],
        "Balance": &arr[1],
        "Speed": &arr[2],
        "Crit Chance": &arr[3],
        "Crit Damage": &arr[4],

        "dps":  &arr[5],
        "real balance":    &arr[6],
        "balance surplus": &arr[7],  
        "balance percent": &arr[8],   // 0,1 - 1
        "adddmg percent": &arr[9],       
        "animation speed": &arr[10],  
        "real crit": &arr[11],         // CRIT<=50
        "new crit damage": &arr[12],   // crit damage with balance surplus
        "crit mod": &arr[13],         //dmg modifier from crits
    }
}

var blink byte
func drawArrow(pos val_pos){
    var arrow =[]byte{' ',' ',' '}
    arrow[blink%3]='>'
    cmov(pos.x-4,pos.y)

    fmt.Print(string(arrow)+">")
}

func check_acc_num(pos val_pos){
    if(accum_number<0){
        return
    }
    cmov(pos.x,pos.y)
    fmt.Print("          ")
    cmov(pos.x,pos.y)
    if(blink%3!=0){
        fmt.Print(accum_number)
    }

    cmov(pos.x,13)
    fmt.Print("Pres E to confirm")
}

func checkarrowinput(mov string){
    if(len(mov)<3){
        return
    }                
    if(mov[1]==91){
        switch mov[2] {
        case 65:   // 27 91 65 up
            validate_arrow_y(-1)
        case 66:   // 27 91 66 down 
            validate_arrow_y(1)
        case 67:   // 27 91 67 >
            validate_arrow_x(1)
        case 68:   // 27 91 68 <
            validate_arrow_x(-1)
        }
    }
}

var exit byte=0
func handle_input(mov string,arr *[2][14]val_pos){
    switch mov[0] {
    case 'w':
        validate_arrow_y(-1)
    case 's':
        validate_arrow_y(1)
    case 'd':
        validate_arrow_x(1)
    case 'a':
        validate_arrow_x(-1)
    case 27:
        checkarrowinput(mov)
    case 'q':
        accum_number=-1
        exit++
        if(exit>5){
            panic("Exit successfull!")
        }
    case 'Q':
        accum_number=-1
        exit++
        if(exit>5){
            panic("Exit successfull!")
        }
    case 'e':
        set_number(&arr[user.x][user.y].val)
    case 'E':
        set_number(&arr[user.x][user.y].val)
    case 10:    //enter
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
    case 127:
        accum_number=accum_number/10
    default:
        cmov(11,17)
        fmt.Print("Use latin layout")
        cmov(11,18)
        fmt.Print("W A S D + numbers")
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

func drwar_tutorial(x int,y int){
    cmov(x+2,y)
    fmt.Print("  W        enter numbers")
    cmov(x,y+2)
    fmt.Print("A   e   D    E: to confirm")
    cmov(x+2,y+4)
    fmt.Print("  S        Q(hold): to Exit")
}
var user =arrow_pos{x:0,y:0}
func main() {
    //var last_exit byte
    mainArr_ptr:=   returnArray()
    mainMap_ptr,compMap_ptr:=   returnMap(&mainArr_ptr[0]),returnMap(&mainArr_ptr[1])//returnMap(compArr_ptr)

    clear()
    drwar_tutorial(11,15)
    
    tmp:=*compMap_ptr["Balance"]
    tmp.val=90
    *compMap_ptr["Balance"]=tmp
    calcDPS(mainMap_ptr)
    calcDPS(compMap_ptr)
    draw_stuff(mainMap_ptr,"Main_Stats")
   
    draw_stuff(compMap_ptr,"Secondary_Stats")

   

    //drawArrow(mainArr_ptr[user.x][user.y])

   
    drawDif(mainArr_ptr[0][5],mainArr_ptr[1][5])
// fmt.Print("\n",mainArr_ptr[user.x][user.y])
//     cmov(0,23)
   //return
    ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 3)
        for {
            os.Stdin.Read(b)
            ch <- string(b)
        }
        defer  exec.Command("stty", "-F", "/dev/tty", "echo").Run() ///??? doesnt work

    }(ch)

    for {
        select {    //
            case stdin, _ := <-ch:
                clear()
                cmov(0,23)
                fmt.Printf("Keys pressed: %s[%v]", stdin,stdin[0])
                handle_input(stdin,mainArr_ptr)
                
                calcDPS(mainMap_ptr)
                calcDPS(compMap_ptr)
                draw_stuff(mainMap_ptr,"Main_Stats")
                draw_stuff(compMap_ptr,"Secondary_Stats")
                
                drawDif(mainArr_ptr[0][5],mainArr_ptr[1][5])
                
            default:
                exit=0
            //     fmt.Println("Working..")
        }
       
        check_acc_num(mainArr_ptr[user.x][user.y])
        drawArrow(mainArr_ptr[user.x][user.y])
        



        cmov(0,23)
        time.Sleep(time.Millisecond * 175)
        
		blink++
    }

}			