
import torch
from torch import nn
import torch.nn.functional
from torch.utils.data import Dataset
from torch.utils.data import DataLoader
from torchvision import datasets
from torchvision.transforms import ToTensor
import matplotlib.pyplot as plt


device = "cuda" if torch.cuda.is_available() else "cpu"

training_data = datasets.FashionMNIST(
    root= "data",
    train=True,
    transform=ToTensor(),
    download=True,
)

test_data = datasets.FashionMNIST(
    root= "data",
    train=False,
    transform=ToTensor(),
    download=True,
)

class NeuralNetwork(nn.Module):
    def __init__(self):
        super().__init__()
        self.flatten = nn.Flatten()
        self.linear_relu_stack = nn.Sequential(
            nn.Linear(28 * 28, 512),
            nn.ReLU(),
            nn.Linear(512, 512),
            nn.ReLU(),
            nn.Linear(512, 10),
        )
    def forward(self, x):
        x = self.flatten(x)
        return self.linear_relu_stack(x)

def train_loop(dataloader, model, loss_fn, optimizer):
    size = len(dataloader.dataset)

    model.train()
    for batch, (X,y) in enumerate(dataloader):
        X, y = X.to(device), y.to(device)
        pred = model(X)
        loss = loss_fn(pred,y)

        loss.backward()
        optimizer.step()
        optimizer.zero_grad()

        if batch % 100 == 0:
            loss, current = loss.item(), batch * batch_size + len(X)
            print(f"loss: {loss:>7f}  [{current:>5d}/{size:>5d}]")

def test_loop(dataloader, model, loss_fn):
    model.eval()
    size = len(dataloader.dataset)
    num_batchs = len(dataloader)
    test_loss, correct = 0, 0

    with torch.no_grad():
        for X, y in dataloader:
            X, y = X.to(device), y.to(device)
            pred = model(X)
            test_loss += loss_fn(pred, y).item()
            correct += (pred.argmax(1) == y).type(torch.float).sum().item()
    test_loss /= num_batchs
    correct /= size
    print(f"Test Error: \n Accuracy: {(100 * correct):>0.1f}%, Avg loss: {test_loss:>8f} \n")


model = NeuralNetwork().to(device)
# Train model
learning_rate = 1e-3
batch_size = 64
epochs = 10
loss_fn = nn.CrossEntropyLoss()
optimizer = torch.optim.SGD(params=model.parameters(), lr=learning_rate)
training_dataloader = DataLoader(dataset=training_data, batch_size=batch_size, shuffle=True)
test_dataloader = DataLoader(dataset=test_data, batch_size=batch_size, shuffle=True)

for t in range(epochs):
    print(f"Epoch {t+1}\n--------------")
    train_loop(training_dataloader, model, loss_fn, optimizer)
    test_loop(test_dataloader, model, loss_fn)

torch.save(model.state_dict(), 'model_weights.pth')



# Use model to do prediction
# model.load_state_dict(torch.load('model_weights.pth'))
# model.eval()

# idx = 500
# plt.imshow(training_data[idx][0].squeeze(), cmap="Greens")
# plt.plot()
# plt.savefig("./data.png")

# X = training_data[idx][0].to(device)
# pred = model(X)
# print(pred)
# print(pred.argmax(1), training_data[idx][1])
