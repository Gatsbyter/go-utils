package link_list

type Node struct {
	val  interface{}
	next *Node
	prev *Node
}

func New(slice []interface{}) *Node {
	if len(slice) == 0 {
		return nil
	}

	head := &Node{
		val:  slice[0],
		next: nil,
		prev: nil,
	}

	node := head

	for i := 1; i < len(slice); i++ {
		node = node.InsertNext(slice[i])
	}

	return head
}

func (object *Node) InsertNext(val interface{}) *Node {
	node := &Node{
		val:  val,
		next: object.next,
		prev: object,
	}

	if object.next != nil {
		object.next.prev = node
	}
	object.next = node

	return node
}

func (object *Node) InsertPrev(val interface{}) {
	node := &Node{
		val:  val,
		next: object,
		prev: object.prev,
	}

	if object.prev != nil {
		object.prev.next = node
	}
	object.prev = node
}

func Delete(node *Node) {
	if node == nil {
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
}

func (object *Node) ToSlice() (res []interface{}) {
	node := object

	for node != nil {
		res = append(res, node.val)
		node = node.next
	}

	return res
}
