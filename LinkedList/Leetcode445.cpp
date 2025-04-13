#include <iostream>

struct Node {
  int data;
  Node *next;
};

class LinkedList {
  Node *head;

public:
  LinkedList() : head(nullptr) {};

  Node *insertAtBegining(int val) {
    Node *newNode = new Node();
    newNode->data = val;
    newNode->next = nullptr;
    head = newNode;
    return head;
  }

  Node *insertAtPos(int val, int position) {
    Node *newNode = new Node();
    newNode->data = val;
    newNode->next = nullptr;
    if (position < 1) {
      return head;
    }
    if (position == 1) {
      Node *newHead = insertAtBegining(val);
      return newHead;
    }
    Node *temp = head;
    for (int i = 1; i < position - 1 && temp != nullptr; ++i) {
      temp = temp->next;
    }
    if (!temp) {
      std::cout << "Position out of range";
      delete newNode;
      return head;
    }
    newNode->next = temp->next;
    temp->next = newNode;
    return head;
  }

  Node *insertAtEnd(int val) {
    Node *newNode = new Node();
    newNode->data = val;
    if (!head) {
      head = newNode;
      return head;
    }
    Node *temp = head;
    while (temp->next) {
      temp = temp->next;
    }
    temp->next = newNode;
    newNode->next = nullptr;
    return head;
  }

  void displayLL() {
    if (!head) {
      std::cout << "List is empty" << std::endl;
      return;
    }
    Node *temp = head;
    while (temp) {
      std::cout << temp->data << "->";
      temp = temp->next;
    }
    std::cout << "nullptr" << std::endl;
  }
  Node *getHead() { return head; }
};

class Leetcode445 {
public:
  static Node *reverseLL(Node *head) {
    Node *backward{nullptr};
    Node *temp{head};
    while (temp) {
      Node *forward = temp->next;
      temp->next = backward;
      backward = temp;
      temp = forward;
    }
    return backward;
  }

  static Node *addTwoNumbers(Node *l1, Node *l2) {
    Node *temp1 = reverseLL(l1);
    Node *temp2 = reverseLL(l2);
    Node *head{nullptr};
    Node *temp3{head};
    int val1{0};
    int val2{0};
    int carry{0};

    while (temp1 || temp2 || carry > 0) {
      if (temp1) {
        val1 = temp1->data;
      } else {
        val1 = 0;
      }
      if (temp2) {
        val2 = temp2->data;
      } else {
        val2 = 0;
      }

      int sum = val1 + val2 + carry;
      carry = sum / 10;
      int digit = sum % 10;
      Node *newNode = new Node();
      newNode->data = digit;
      newNode->next = nullptr;
      if (!head) {
        head = newNode;
        temp3 = head;
      } else {
        temp3->next = newNode;
        temp3 = temp3->next;
      }
      if (temp1) {
        temp1 = temp1->next;
      }
      if (temp2) {
        temp2 = temp2->next;
      }
    }
    return reverseLL(head);
  }
};
int main() {
  LinkedList list1, list2;

  list1.insertAtEnd(7);
  list1.insertAtEnd(2);
  list1.insertAtEnd(4);
  list1.insertAtEnd(3);

  list2.insertAtEnd(5);
  list2.insertAtEnd(6);
  list2.insertAtEnd(4);

  std::cout << "List 1: ";
  list1.displayLL();

  std::cout << "List 2: ";
  list2.displayLL();

  Node *result = Leetcode445::addTwoNumbers(list1.getHead(), list2.getHead());

  LinkedList resultList;
  Node *temp = result;
  while (temp) {
    resultList.insertAtEnd(temp->data);
    temp = temp->next;
  }
  resultList.displayLL();
}
