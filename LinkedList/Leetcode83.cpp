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
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode* dummy = new ListNode(-1);
        ListNode* temp1 = head;
        ListNode* temp2 = dummy;

        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        while (temp1 != nullptr) {
            if (temp2 == dummy || temp1->val != temp2->val) {
                temp2->next = temp1;
                temp2 = temp2->next;
            }
            temp1 = temp1->next;
        }
        temp2->next = nullptr;
        return dummy->next;
    }
};
