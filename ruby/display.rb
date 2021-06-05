module DisplayBoard
    def DisplayBoard.display_board(board)
        line='_______'
        puts "#{board[0]} " "|" "#{board[1]} " "|" "#{board[2]} " 
        puts line
        puts "#{board[3]} " "|" "#{board[4]} " "|" "#{board[5]} "
        puts line 
        puts "#{board[6]} " "|" "#{board[7]} " "|" "#{board[8]} "
    end
end
        
