include<unordered_map>

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        std::unordered_map<ListNode*,bool> valStore;
        ListNode* temp = head;

        while (temp != nullptr;){
            if (valStore[temp]){
                return true;
            }
            valStore[temp] = true;
            temp = temp -> next;
        }
        return false;
    }
};
