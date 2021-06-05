class Validate
    $winning_combo=[]
    $winning_combo.append([0,1,2])
    $winning_combo.append([3,4,5])
    $winning_combo.append([6,7,8])
    $winning_combo.append([0,3,6])
    $winning_combo.append([1,4,7])
    $winning_combo.append([2,5,8])
    $winning_combo.append([0,4,8])
    $winning_combo.append([2,4,6])
   
    def move(index)
        index.to_i
      
        if $board[index]=="X" || $board[index]=="Y"
            return 0
        end

        if index<0 || index>8
            return 0
        end
        
        return 1
    end

    def win()
       
        for c1,c2,c3 in $winning_combo
            if $board[c1]==$board[c2] && $board[c3]==$board[c2]
                if $board[c1]=="X"
                    return 1
                end
                if $board[c1]=="Y"
                    return 2
                end
            end
        end
        return 0
    end
    
    def draw()
        @@res=0
        $board.each do |i|
            if i=="X" || i=="Y"
                @@res+=1
            end
        end
        if @@res==9
            return 1
        end
        return 0
    
    end

    
end

