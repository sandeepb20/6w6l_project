require "./display.rb"
require "./validate.rb"
$board=[0,1,2,3,4,5,6,7,8]
$count=0
i=1

DisplayBoard.display_board($board)

My_class=Validate.new()

while true
    
    
    if ($count)%2==1
        puts "This is turn of player 2.Please Enter a Number"
        index=gets.chomp.to_i

        if My_class.move(index)==1
            $board[index]="Y"
            $count+=1
            DisplayBoard.display_board($board)
        else
            puts "INVALID______Please enter a valid number"
        end
    else
        puts "This is turn of player 1.Please Enter a Number"
        index=gets.chomp.to_i
       

        if My_class.move(index)==1
            $board[index]="X"
            $count+=1
            DisplayBoard.display_board($board)
        else
            puts "INVALID_______ Please enter a valid number"
        end
    end
    if My_class.win()==1
        
        puts "Player 1 Won the match :)"
        break
    end
    if My_class.win()==2
       
        puts "PLayer 2 Won the match"
        break
    end
    if My_class.draw()==1
        puts "This one is draw ....RETRY "
        break
    end
end

puts " DONE"