package messageHandler

func CreateMessage(receiver byte, id byte, msgType byte, content []byte) []byte {
	var message []byte
	if receiver != REC_NONE {
		message = []byte{MSG_START, receiver, id, msgType}		
		message = append(message, content...)
	} else {
		message = []byte{MSG_START, receiver}
	}
	return message
}

func MessageRequestCost(id byte, request []byte) []byte {
	return CreateMessage(REC_ALL, id, COST_REQUEST, request)
}

func MessageCostReply(receiver byte, id byte, floor byte, buttonType byte, cost int) []byte {
	return CreateMessage(receiver, id, COST_REPLY, []byte{floor, buttonType, byte(cost)})
}

func MessageOrderComplete(id byte, request []byte) []byte {
	return CreateMessage(REC_ALL, id, CLEAR_ORDER, request)
}

func MessageAssignOrder(id byte, assReq []byte) []byte {
	receiver := assReq[0]
	data := []byte{assReq[1], assReq[2]}
	return CreateMessage(receiver, id, ASSIGN_ORDER, data)
}

func MessageAcknowledgeOrder(receiver byte, id byte, floor byte, buttonType byte) []byte {
	return CreateMessage(receiver, id, ORDER_ACK, []byte{floor, buttonType})
}