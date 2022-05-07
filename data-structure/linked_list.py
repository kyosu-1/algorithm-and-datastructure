from typing import List, Any, Optional


class Node:
    def __init__(self, data: Optional[Any] = None, next_node: Optional['Node'] = None) -> None:
        self.data = data
        self.next_node = next_node


class LinkedList:
    def __init__(self) -> None:
        self.head: Optional[Node] = None
        self.last_node: Optional[Node] = None
    
    def to_list(self) -> List[Any]:
        """Convert a linear list to a standard list and return it.

        Returns:
            List[Any]: list
        """
        l = []
        if self.head is None:
            return l
        
        node = self.head
        while node:
            l.append(node.data)
            node = node.next_node
        return l
    
    def print_ll(self):
        """Output a linear list as a string to the standard output
        """
        ll_string = ""
        node = self.head
        if node is None:
            print(None)
        while node:
            ll_string += f"{str(node.data)} ->"
            node = node.next_node

        ll_string += " None"
        print(ll_string)
    
    def insert_beginning(self, data: Any) -> None:
        """Insert at the top

        Args:
            data (Any): data
        """
        if self.head is None:
            self.head = Node(data, None)
            self.last_node = self.head
        
        new_node = Node(data, self.head)
        self.head = new_node
    
    def insert_at_end(self, data: Any) -> None:
        """Insert at the end

        Args:
            data (Any): data
        """
        if self.head is None:
            self.insert_beginning(data)
            return
        
        self.last_node.next_node = Node(data, None)
        self.last_node = self.last_node.next_node
