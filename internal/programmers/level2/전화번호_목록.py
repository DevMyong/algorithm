def solution(phone_book):
    pb = {v: v for _, v in enumerate(phone_book)}
    
    for phone in phone_book:
        prefix = ''
        
        for num in phone:
            prefix += num
            
            if prefix in pb and prefix != pb[phone]:
                return False
            
    return True