# 表结点类
class LNode:
  def __init__(self, elem, next_=None):
    self.elem = elem
    self.next = next_

# 异常类
class LinkedListUnderflow(ValueError):
  pass

class LList:
  def __init__(self):
    self._head = None
  
  def is_empty(self):
    return self._head is None
  
  def prepend(self, elem):
    self._head = LNode(elem, self._head)

  def pop(self):
    if self._head is None:
      raise LinkedListUnderflow("in pop")
    e = self._head.elem
    self._head = self._head.next
    return e

  def append(self, elem):
    if self._head is None:
      self._head = LNode(elem)
      return
    p = self._head
    while p.next is not None:
      p = p.next
    p.next = LNode(elem)

  def pop_last(self):
    if self._head is None:
      raise LinkedListUnderflow("in pop_last")
    p = self._head
    if p.next is None:
      e = p.elem
      self._head = None
      return e

    while p.next.next is not None:
      p = p.next
    e = p.next.elem
    p.next = None
    return e

  def filter(self, pred):
    p = self._head
    while p is not None:
      if pred(p.elem):
        yield p.elem
      p = p.next
  
  def printall(self):
    p = self._head
    while p is not None:
      print(p.elem, end='')
      if p.next is not None:
        print(', ', end='')
      p = p.next
    print('')

  def for_each(self, proc):
    p = self._head
    while p is not None:
      proc(p.elem)
      p = p.next
  
  def elements(self):
    p = self._head
    while p is not None:
      yield p.elem
      p = p.next

# 使用链表
mlist1 = LList()
for i in range(1, 10):
  mlist1.append(i)
  mlist1.append(i)

# mlist1.printall()
# mlist1.for_each(print)