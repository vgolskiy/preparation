# Complexity O(n) n - number of competitions,
# Space O(k) k - number of teams
def tournamentWinner(competitions, results):
    winners = {}
    winner = ""
    for i, teams in enumerate(competitions):
        team = teams[0] if results[i] else teams[1]
        if team in winners:
            winners[team] += 3
        else:
            winners[team] = 3
        if winner != "" and winners[team] > winners[winner]:
            winner = team
        if winner == "":
            winner = team
    return winner


competitions = [['HTML', 'C#'], ['C#', 'Python'], ['Python', 'HTML']]
results = [0, 0, 1]

if __name__ == '__main__':
    print(tournamentWinner(competitions, results))
