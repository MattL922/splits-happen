package main

import (
    "testing"
)


func TestScore(t *testing.T) {
    tests := []struct{
        line string
        score int
    }{
        {"XXXXXXXXXXXX", 300},
        {"9-9-9-9-9-9-9-9-9-9-", 90},
        {"5/5/5/5/5/5/5/5/5/5/5", 150},
        {"X7/9-X-88/-6XXX81", 167},
        {"--------------------", 0},
        {"------------------X--", 10},
        {"X--X--X--X--X--", 50},
        {"--X--X--X--X--X--", 50},
        {"-/-/-/-/-/-/-/-/-/-/-", 100},
        {"9/9/9/9/9/9/9/9/9/9/-", 181},
    }
    for _, test := range tests {
        l := NewLine(test.line)
        if l.Score() != test.score {
            t.Errorf("Line %q calculated score %d, expected %d", test.line, l.Score(), test.score)
        }
    }
}
