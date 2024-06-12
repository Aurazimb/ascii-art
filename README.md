# AsciiArt
## About script
<br>This script will convert your text in the terminal to Ascii Art in the style you choose (if you don't specify a style, "standard" will be automatically selected).</br>
<br>Example: </br>
```sh
go run . "Hello" | cat -e
```

```sh
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```


## Installing
<br>For installing open your terminal and paste the next code</br>
```sh
git clone git@git.01.alem.school:aurazimb/ascii-art.git
```

## Usage
After downloaded, open the "ascii-art" repository using the next code
```sh
cd ascii-art
```
Use the following prompt to use the script
```sh
go run . "Your text" [style]
```

### Styles
You can use once of styles
- standard
- shadow
- thinkertoy

### Examples

#### STANDARD
```sh
go run . "123" standard
```

```sh
                    
 _   ____    _____  
/ | |___ \  |___ /  
| |   __) |   |_ \  
| |  / __/   ___) | 
|_| |_____| |____/  
                    
                    
```


#### THIKERTOY

```sh
go run . "abc" thinkertoy
```

```sh
               
     o         
     |         
 oo  O-o   o-o 
| |  |  | |    
o-o- o-o   o-o 
               
               
```

#### SHADOW


```sh
go run . "ABC" shadow
```

```sh
                           
  _|_|   _|_|_|     _|_|_| 
_|    _| _|    _| _|       
_|_|_|_| _|_|_|   _|       
_|    _| _|    _| _|       
_|    _| _|_|_|     _|_|_| 
                           
                           
```

## Extra

For newline you can use ``` \n ``` or you can send a few lines as text
```sh
go run . "ABC
abc
123
" standard
``` 

```sh
            ____     _____  
    /\     |  _ \   / ____| 
   /  \    | |_) | | |      
  / /\ \   |  _ <  | |      
 / ____ \  | |_) | | |____  
/_/    \_\ |____/   \_____| 
                            
                            
         _             
        | |            
  __ _  | |__     ___  
 / _` | | '_ \   / __| 
| (_| | | |_) | | (__  
 \__,_| |_.__/   \___| 
                       
                       
                    
 _   ____    _____  
/ | |___ \  |___ /  
| |   __) |   |_ \  
| |  / __/   ___) | 
|_| |_____| |____/  
                    
                    

``` 


# Authors:

- ynurmakh
- aurazimb