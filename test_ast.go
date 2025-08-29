func parseExpression(tokens Instruction, pos *int) Node {
    left := parseTerm(tokens, pos)

    for *pos < len(tokens) {
        t := tokens[*pos]

		if t.Type != tok.OP {
			break
		}

		if t.BinOpType != tok.PLUS && t.BinOpType != tok.MIN {
			break
		}

        *pos++
        right := parseTerm(tokens, pos)
        left = BinOpNode{Left: left, Operator: t.Text, Right: right}
    }

    return left
}

func parseTerm(tokens Instruction, pos *int) Node {
    left := parseFactor(tokens, pos)

    for *pos < len(tokens) {
        t := tokens[*pos]

        if t.Type != tok.OP {
            break
        }

        if t.BinOpType != tok.MUL && t.BinOpType != tok.DIV {
            break
        }

        *pos++
        right := parseFactor(tokens, pos)
        left = BinOpNode{Left: left, Operator: t.Text, Right: right}
    }

    return left
}

func parseFactor(tokens Instruction, pos *int) Node {
    t := tokens[*pos]
    *pos++

    switch t.Type {
    case tok.NUMBER:
        var val int
        fmt.Sscanf(t.Text, "%d", &val)
        return NumberNode{Value: val}

    case tok.IDENTIFIER:
        return IdentifierNode{Name: t.Text}

    case tok.LPAREN:
        node := parseExpression(tokens, pos)
        *pos++
        return node
    }

    return nil
}
