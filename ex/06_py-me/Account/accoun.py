# ... ing

class Account:

    def __init__(self, accountNo, ownString, balance):
        self.accountNo = accountNo
        self.ownString = ownString
        self.balance = balance

    def deposit(self, amount):
        self.balance += amount

    def withdraw(self, amount):
        if self.balance < amount:
            print("not enough money")
            return 0
            self.balance -= amount
        return amount

if __name__ == '__main__':
    acc = Account("223344", "Obama", 30000)
    acc.deposit(700)
    money = acc.withdraw(2200)
    print(money)
    print(acc.balance)

