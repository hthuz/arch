import random
import uuid
from collections import defaultdict

class Participant:
    def __init__(self, parent=None):
        self.id = uuid.uuid4()
        self.parent = parent  # 上线
        self.children = []
        self.referred_count = 0
        self.invested = 10000
        self.earned = 0
    
    def profit(self):
        return self.earned - self.invested

class PonziTrainingSystem:
    def __init__(self, referral_reward=2000, system_cut=8000):
        self.members = []  # 所有成员
        self.active_members = []  # 还能拉人的人
        self.institution_earning = 0
        self.referral_reward = referral_reward
        self.system_cut = system_cut

    def add_participant(self, parent=None):
        p = Participant(parent)
        if parent:
            parent.children.append(p)
            parent.referred_count += 1
            parent.earned += self.referral_reward
        self.members.append(p)
        self.active_members.append(p)
        self.institution_earning += self.system_cut
        return p

    def simulate(self, max_rounds=20):
        # 初始启动者
        root = self.add_participant(parent=None)

        for round in range(max_rounds):
            print(f"round {round} start: all members {len(self.members)}, active members {len(self.active_members)}")
            new_members = []
            for member in list(self.active_members):
                if member.referred_count >= 10:  # 限制最大拉人数量
                    self.active_members.remove(member)
                    continue
                # 推荐人数模拟（大部分人拉不到人，少部分人能拉很多）
                n = max(0, int(random.gauss(mu=1.5, sigma=1.5)))
                for _ in range(n):
                    new_p = self.add_participant(parent=member)
                    new_members.append(new_p)
            if not new_members:
                break  # 没人继续加入，系统崩盘
        return

    def summary(self):
        total_invested = sum(p.invested for p in self.members)
        total_earned_by_members = sum(p.earned for p in self.members)
        winners = sum(1 for p in self.members if p.profit() > 0)
        losers = len(self.members) - winners

        print(f"系统参与者总数：{len(self.members)}")
        print(f"系统机构总盈利：{self.institution_earning}")
        print(f"成员总投资：{total_invested}")
        print(f"成员返利总收入：{total_earned_by_members}")
        print(f"真正赚钱的人数：{winners} / {len(self.members)}")
        print(f"亏钱的人数：{losers} / {len(self.members)}")

        top_losers = sorted(self.members, key=lambda p: p.profit())[:5]
        print("\n亏得最惨的5个人（亏损金额）：")
        for p in top_losers:
            print(f"亏损：{p.profit()}，推荐人数：{p.referred_count}")

# 执行模拟
if __name__ == "__main__":
    system = PonziTrainingSystem()
    system.simulate(max_rounds=10)
    system.summary()

