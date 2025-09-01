import leon, kaisa, klaudia, lena, maliah, joana

team_members = [leon, kaisa, klaudia, lena, maliah, joana]

def group_intro():
    print('This is Team robocops. We are:')
    for member in team_members:
        print(member.name())
group_intro()

def group_bonus():
   for member in team_members:
      print(member.act1())
      print(member.act2())
      print(member.act3())
    
group_bonus()