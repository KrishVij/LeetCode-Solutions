/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* temp = head;
        ListNode* prev = nullptr;
        int len {1};
        int counter {1};
        

        while (temp != nullptr){
            len++;
            temp = temp -> next;
        }
        temp = head;
        while (counter != len - n){
            prev = temp; 
            temp = temp -> next;
        }
        prev -> next = temp ->next;
        delete temp;
        return head;


        

    }
};
