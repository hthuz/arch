
import torch
import torch.nn as nn
import torch.nn.functional as F
import torch.optim as optim
from torch.utils.data import DataLoader, TensorDataset
from sklearn.datasets import make_classification
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
import matplotlib.pyplot as plt

# 设置随机种子确保可重复性
torch.manual_seed(42)

# 1. 生成示例数据
X, y = make_classification(
    n_samples=1000,  # 样本数量
    n_features=20,   # 特征数量
    n_informative=15, # 有效特征
    n_redundant=5,   # 冗余特征
    n_classes=2,     # 二分类
    random_state=42
)


# 2. 数据预处理
# 划分训练集和测试集
X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.2, random_state=42
)

# 标准化特征
scaler = StandardScaler()
X_train = scaler.fit_transform(X_train)
X_test = scaler.transform(X_test)

# 转换为PyTorch张量
X_train_tensor = torch.FloatTensor(X_train)
y_train_tensor = torch.LongTensor(y_train)
X_test_tensor = torch.FloatTensor(X_test)
y_test_tensor = torch.LongTensor(y_test)

# 创建数据加载器
train_dataset = TensorDataset(X_train_tensor, y_train_tensor)
test_dataset = TensorDataset(X_test_tensor, y_test_tensor)
train_loader = DataLoader(train_dataset, batch_size=32, shuffle=True)
test_loader = DataLoader(test_dataset, batch_size=32)

class MLP(nn.Module):
    def __init__(self, input_size, hidden_size, output_size):
        super(MLP, self).__init__()

        self.fc1 = nn.Linear(input_size, hidden_size)  # 输入层到隐藏层
        self.fc2 = nn.Linear(hidden_size, hidden_size) # 隐藏层到隐藏层
        self.fc3 = nn.Linear(hidden_size, output_size) # 隐藏层到输出层
        self.dropout = nn.Dropout(0.2)  # 防止过拟合
        
    def forward(self, x):
        # 前向传播
        x = F.relu(self.fc1(x))  # 第一层 + ReLU激活函数
        x = self.dropout(x)
        x = F.relu(self.fc2(x))  # 第二层 + ReLU激活函数
        x = self.dropout(x)
        x = self.fc3(x)           # 输出层（无激活函数，用交叉熵损失时内部会做softmax）
        return x

# 4. 初始化模型
input_size = X.shape[1]  # 输入特征数
hidden_size = 64         # 隐藏层神经元数量
output_size = 2          # 输出类别数（二分类）

model = MLP(input_size, hidden_size, output_size)

# 5. 定义损失函数和优化器
loss_fn = nn.CrossEntropyLoss()
optimizer = optim.Adam(model.parameters(), lr=0.001)

def train_model(model: nn.Module, train_loader: DataLoader, loss_fn: nn.CrossEntropyLoss, optimizer: optim.Adam, epochs=50):
    model.train()
    train_losses = []
    train_accuracies = []
    
    for epoch in range(epochs):
        running_loss = 0.0
        correct = 0
        total = 0
        
        for batch_idx, (X, y) in enumerate(train_loader):
            X: torch.Tensor
            y: torch.Tensor
            # 梯度清零
            optimizer.zero_grad()
            
            outputs: torch.Tensor = model(X)
            print(outputs, outputs.shape)
            exit()
            loss: torch.Tensor = loss_fn(outputs, y)
            loss.backward()
            optimizer.step()
            
            # 统计
            running_loss += loss.item()
            _, predicted = torch.max(outputs.data, 1)
            total += y.size(0)
            correct += (predicted == y).sum().item()
        
        # 计算每个epoch的平均损失和准确率
        epoch_loss = running_loss / len(train_loader)
        epoch_acc = 100 * correct / total
        train_losses.append(epoch_loss)
        train_accuracies.append(epoch_acc)
        
        if (epoch + 1) % 10 == 0:
            print(f'Epoch [{epoch+1}/{epochs}], Loss: {epoch_loss:.4f}, Accuracy: {epoch_acc:.2f}%')
    
    return train_losses, train_accuracies

# 7. 测试函数
def test_model(model: MLP, test_loader):
    model.eval()  # 设置为评估模式
    correct = 0
    total = 0
    
    with torch.no_grad():  # 不需要计算梯度
        for data, target in test_loader:
            outputs = model(data)
            _, predicted = torch.max(outputs.data, 1)
            total += target.size(0)
            correct += (predicted == target).sum().item()
    
    accuracy = 100 * correct / total
    print(f'测试集准确率: {accuracy:.2f}%')
    return accuracy

# 8. 训练模型
train_losses, train_accuracies = train_model(
    model, train_loader, loss_fn, optimizer, epochs=50
)

# 9. 测试模型
test_accuracy = test_model(model, test_loader)

# 10. 可视化训练过程
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(12, 4))

# 绘制损失曲线
ax1.plot(train_losses)
ax1.set_xlabel('Epoch')
ax1.set_ylabel('Loss')
ax1.set_title('Training Loss')
ax1.grid(True)

# 绘制准确率曲线
ax2.plot(train_accuracies)
ax2.set_xlabel('Epoch')
ax2.set_ylabel('Accuracy (%)')
ax2.set_title('Training Accuracy')
ax2.grid(True)

plt.tight_layout()
plt.show()

# 11. 保存模型（可选）
# torch.save(model.state_dict(), 'mlp_model.pth')

# 12. 加载模型（可选）
# model = MLP(input_size, hidden_size, output_size)
# model.load_state_dict(torch.load('mlp_model.pth'))
