class Solution {
  public:
    /* Should return data of middle node. If linked list is empty, then -1 */
    int getMiddle(Node* head) {
        Node* temp = head;
        int count {0};
        int counter {0};
        
        
        while (temp != nullptr)
        {
            count++;
            temp = temp -> next;
        }
        int valIndex = count / 2;
        temp = head;
        while (temp != nullptr)
        {
            counter++;
            if (counter == valIndex + 1)
            {
                return temp -> data;
                
            }
            temp = temp -> next;
        }
        
    }
};
