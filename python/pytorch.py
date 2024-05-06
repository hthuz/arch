
import torch


x = torch.zeros(2,2,requires_grad=True)
y = x + 2
z = (y * y).mean()
print(y.grad_fn)
print(z.backward())
