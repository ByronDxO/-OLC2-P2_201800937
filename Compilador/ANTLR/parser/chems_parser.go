// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/expression"
import "OLC2/Compilador/instruction"
import "OLC2/Compilador/instruction/loops"
import "OLC2/Compilador/instruction/control"
import "OLC2/Compilador/instruction/variable"
import "OLC2/Compilador/instruction/ternario"
import "OLC2/Compilador/instruction/function"
import "OLC2/Compilador/instruction/casteo"
import "OLC2/Compilador/instruction/transferencia"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 769,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 105, 10, 3, 13, 3, 14, 3, 106,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 114, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 163, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 181, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 241, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 276, 10, 10, 3, 11, 7,
	11, 279, 10, 11, 12, 11, 14, 11, 282, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 313, 10, 12, 3, 13, 7, 13, 316, 10,
	13, 12, 13, 14, 13, 319, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 338, 10, 14, 3, 15, 6, 15, 341, 10, 15, 13, 15, 14, 15, 342,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 360, 10, 16, 3, 17, 6, 17, 363,
	10, 17, 13, 17, 14, 17, 364, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 5, 18, 376, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 394, 10, 20, 3, 21, 6, 21, 397, 10, 21, 13, 21, 14, 21, 398,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 418, 10, 22, 3, 23,
	6, 23, 421, 10, 23, 13, 23, 14, 23, 422, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	5, 24, 440, 10, 24, 3, 25, 6, 25, 443, 10, 25, 13, 25, 14, 25, 444, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 456,
	10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 471, 10, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 500, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 5, 32, 520, 10, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 5, 35, 560, 10, 35, 3, 36, 7, 36, 563, 10, 36, 12, 36, 14, 36,
	566, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 581, 10, 37, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 593, 10, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 5, 39, 607, 10, 39, 3, 40, 6, 40, 610, 10, 40, 13, 40, 14, 40, 611,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 623,
	10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 5, 42, 637, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 5, 43, 647, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 7, 43, 654, 10, 43, 12, 43, 14, 43, 657, 11, 43, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 7, 44, 668, 10, 44, 12, 44,
	14, 44, 671, 11, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 686, 10, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 698, 10,
	45, 12, 45, 14, 45, 701, 11, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 5, 47, 730, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 756, 10, 48, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 5, 50, 767, 10,
	50, 3, 50, 2, 5, 84, 86, 88, 51, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	96, 98, 2, 6, 3, 2, 57, 58, 4, 2, 38, 39, 42, 45, 3, 2, 46, 48, 3, 2, 49,
	50, 2, 798, 2, 100, 3, 2, 2, 2, 4, 104, 3, 2, 2, 2, 6, 113, 3, 2, 2, 2,
	8, 162, 3, 2, 2, 2, 10, 180, 3, 2, 2, 2, 12, 182, 3, 2, 2, 2, 14, 240,
	3, 2, 2, 2, 16, 242, 3, 2, 2, 2, 18, 275, 3, 2, 2, 2, 20, 280, 3, 2, 2,
	2, 22, 312, 3, 2, 2, 2, 24, 317, 3, 2, 2, 2, 26, 337, 3, 2, 2, 2, 28, 340,
	3, 2, 2, 2, 30, 359, 3, 2, 2, 2, 32, 362, 3, 2, 2, 2, 34, 375, 3, 2, 2,
	2, 36, 377, 3, 2, 2, 2, 38, 393, 3, 2, 2, 2, 40, 396, 3, 2, 2, 2, 42, 417,
	3, 2, 2, 2, 44, 420, 3, 2, 2, 2, 46, 439, 3, 2, 2, 2, 48, 442, 3, 2, 2,
	2, 50, 455, 3, 2, 2, 2, 52, 470, 3, 2, 2, 2, 54, 472, 3, 2, 2, 2, 56, 499,
	3, 2, 2, 2, 58, 501, 3, 2, 2, 2, 60, 507, 3, 2, 2, 2, 62, 519, 3, 2, 2,
	2, 64, 521, 3, 2, 2, 2, 66, 524, 3, 2, 2, 2, 68, 559, 3, 2, 2, 2, 70, 564,
	3, 2, 2, 2, 72, 580, 3, 2, 2, 2, 74, 592, 3, 2, 2, 2, 76, 606, 3, 2, 2,
	2, 78, 609, 3, 2, 2, 2, 80, 622, 3, 2, 2, 2, 82, 636, 3, 2, 2, 2, 84, 646,
	3, 2, 2, 2, 86, 658, 3, 2, 2, 2, 88, 685, 3, 2, 2, 2, 90, 702, 3, 2, 2,
	2, 92, 729, 3, 2, 2, 2, 94, 755, 3, 2, 2, 2, 96, 757, 3, 2, 2, 2, 98, 766,
	3, 2, 2, 2, 100, 101, 5, 4, 3, 2, 101, 102, 8, 2, 1, 2, 102, 3, 3, 2, 2,
	2, 103, 105, 5, 8, 5, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109,
	8, 3, 1, 2, 109, 5, 3, 2, 2, 2, 110, 111, 7, 34, 2, 2, 111, 114, 8, 4,
	1, 2, 112, 114, 8, 4, 1, 2, 113, 110, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2,
	114, 7, 3, 2, 2, 2, 115, 116, 5, 10, 6, 2, 116, 117, 5, 6, 4, 2, 117, 118,
	8, 5, 1, 2, 118, 163, 3, 2, 2, 2, 119, 120, 5, 12, 7, 2, 120, 121, 8, 5,
	1, 2, 121, 163, 3, 2, 2, 2, 122, 123, 5, 14, 8, 2, 123, 124, 8, 5, 1, 2,
	124, 163, 3, 2, 2, 2, 125, 126, 5, 16, 9, 2, 126, 127, 8, 5, 1, 2, 127,
	163, 3, 2, 2, 2, 128, 129, 5, 18, 10, 2, 129, 130, 8, 5, 1, 2, 130, 163,
	3, 2, 2, 2, 131, 132, 5, 26, 14, 2, 132, 133, 8, 5, 1, 2, 133, 163, 3,
	2, 2, 2, 134, 135, 5, 54, 28, 2, 135, 136, 8, 5, 1, 2, 136, 163, 3, 2,
	2, 2, 137, 138, 5, 56, 29, 2, 138, 139, 8, 5, 1, 2, 139, 163, 3, 2, 2,
	2, 140, 141, 5, 58, 30, 2, 141, 142, 8, 5, 1, 2, 142, 163, 3, 2, 2, 2,
	143, 144, 5, 62, 32, 2, 144, 145, 5, 6, 4, 2, 145, 146, 8, 5, 1, 2, 146,
	163, 3, 2, 2, 2, 147, 148, 5, 64, 33, 2, 148, 149, 5, 6, 4, 2, 149, 150,
	8, 5, 1, 2, 150, 163, 3, 2, 2, 2, 151, 152, 5, 68, 35, 2, 152, 153, 8,
	5, 1, 2, 153, 163, 3, 2, 2, 2, 154, 155, 5, 74, 38, 2, 155, 156, 5, 6,
	4, 2, 156, 157, 8, 5, 1, 2, 157, 163, 3, 2, 2, 2, 158, 159, 5, 66, 34,
	2, 159, 160, 5, 6, 4, 2, 160, 161, 8, 5, 1, 2, 161, 163, 3, 2, 2, 2, 162,
	115, 3, 2, 2, 2, 162, 119, 3, 2, 2, 2, 162, 122, 3, 2, 2, 2, 162, 125,
	3, 2, 2, 2, 162, 128, 3, 2, 2, 2, 162, 131, 3, 2, 2, 2, 162, 134, 3, 2,
	2, 2, 162, 137, 3, 2, 2, 2, 162, 140, 3, 2, 2, 2, 162, 143, 3, 2, 2, 2,
	162, 147, 3, 2, 2, 2, 162, 151, 3, 2, 2, 2, 162, 154, 3, 2, 2, 2, 162,
	158, 3, 2, 2, 2, 163, 9, 3, 2, 2, 2, 164, 165, 7, 4, 2, 2, 165, 166, 7,
	51, 2, 2, 166, 167, 5, 92, 47, 2, 167, 168, 7, 52, 2, 2, 168, 169, 7, 34,
	2, 2, 169, 170, 8, 6, 1, 2, 170, 181, 3, 2, 2, 2, 171, 172, 7, 4, 2, 2,
	172, 173, 7, 51, 2, 2, 173, 174, 7, 29, 2, 2, 174, 175, 7, 35, 2, 2, 175,
	176, 5, 78, 40, 2, 176, 177, 7, 52, 2, 2, 177, 178, 7, 34, 2, 2, 178, 179,
	8, 6, 1, 2, 179, 181, 3, 2, 2, 2, 180, 164, 3, 2, 2, 2, 180, 171, 3, 2,
	2, 2, 181, 11, 3, 2, 2, 2, 182, 183, 7, 18, 2, 2, 183, 184, 7, 17, 2, 2,
	184, 185, 7, 51, 2, 2, 185, 186, 7, 52, 2, 2, 186, 187, 7, 53, 2, 2, 187,
	188, 5, 4, 3, 2, 188, 189, 7, 54, 2, 2, 189, 190, 8, 7, 1, 2, 190, 13,
	3, 2, 2, 2, 191, 192, 7, 5, 2, 2, 192, 193, 7, 6, 2, 2, 193, 194, 7, 31,
	2, 2, 194, 195, 7, 37, 2, 2, 195, 196, 5, 82, 42, 2, 196, 197, 7, 34, 2,
	2, 197, 198, 8, 8, 1, 2, 198, 241, 3, 2, 2, 2, 199, 200, 7, 5, 2, 2, 200,
	201, 7, 6, 2, 2, 201, 202, 7, 31, 2, 2, 202, 203, 7, 36, 2, 2, 203, 204,
	5, 76, 39, 2, 204, 205, 7, 34, 2, 2, 205, 206, 8, 8, 1, 2, 206, 241, 3,
	2, 2, 2, 207, 208, 7, 5, 2, 2, 208, 209, 7, 6, 2, 2, 209, 210, 7, 31, 2,
	2, 210, 211, 7, 36, 2, 2, 211, 212, 5, 76, 39, 2, 212, 213, 7, 37, 2, 2,
	213, 214, 5, 82, 42, 2, 214, 215, 7, 34, 2, 2, 215, 216, 8, 8, 1, 2, 216,
	241, 3, 2, 2, 2, 217, 218, 7, 5, 2, 2, 218, 219, 7, 31, 2, 2, 219, 220,
	7, 37, 2, 2, 220, 221, 5, 82, 42, 2, 221, 222, 7, 34, 2, 2, 222, 223, 8,
	8, 1, 2, 223, 241, 3, 2, 2, 2, 224, 225, 7, 5, 2, 2, 225, 226, 7, 31, 2,
	2, 226, 227, 7, 36, 2, 2, 227, 228, 5, 76, 39, 2, 228, 229, 7, 34, 2, 2,
	229, 230, 8, 8, 1, 2, 230, 241, 3, 2, 2, 2, 231, 232, 7, 5, 2, 2, 232,
	233, 7, 31, 2, 2, 233, 234, 7, 36, 2, 2, 234, 235, 5, 76, 39, 2, 235, 236,
	7, 37, 2, 2, 236, 237, 5, 82, 42, 2, 237, 238, 7, 34, 2, 2, 238, 239, 8,
	8, 1, 2, 239, 241, 3, 2, 2, 2, 240, 191, 3, 2, 2, 2, 240, 199, 3, 2, 2,
	2, 240, 207, 3, 2, 2, 2, 240, 217, 3, 2, 2, 2, 240, 224, 3, 2, 2, 2, 240,
	231, 3, 2, 2, 2, 241, 15, 3, 2, 2, 2, 242, 243, 7, 31, 2, 2, 243, 244,
	7, 37, 2, 2, 244, 245, 5, 82, 42, 2, 245, 246, 7, 34, 2, 2, 246, 247, 8,
	9, 1, 2, 247, 17, 3, 2, 2, 2, 248, 249, 7, 7, 2, 2, 249, 250, 5, 82, 42,
	2, 250, 251, 7, 53, 2, 2, 251, 252, 5, 4, 3, 2, 252, 253, 7, 54, 2, 2,
	253, 254, 8, 10, 1, 2, 254, 276, 3, 2, 2, 2, 255, 256, 7, 7, 2, 2, 256,
	257, 5, 82, 42, 2, 257, 258, 7, 53, 2, 2, 258, 259, 5, 4, 3, 2, 259, 260,
	7, 54, 2, 2, 260, 261, 7, 8, 2, 2, 261, 262, 7, 53, 2, 2, 262, 263, 5,
	4, 3, 2, 263, 264, 7, 54, 2, 2, 264, 265, 8, 10, 1, 2, 265, 276, 3, 2,
	2, 2, 266, 267, 7, 7, 2, 2, 267, 268, 5, 82, 42, 2, 268, 269, 7, 53, 2,
	2, 269, 270, 5, 4, 3, 2, 270, 271, 7, 54, 2, 2, 271, 272, 7, 8, 2, 2, 272,
	273, 5, 20, 11, 2, 273, 274, 8, 10, 1, 2, 274, 276, 3, 2, 2, 2, 275, 248,
	3, 2, 2, 2, 275, 255, 3, 2, 2, 2, 275, 266, 3, 2, 2, 2, 276, 19, 3, 2,
	2, 2, 277, 279, 5, 18, 10, 2, 278, 277, 3, 2, 2, 2, 279, 282, 3, 2, 2,
	2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 283, 3, 2, 2, 2, 282,
	280, 3, 2, 2, 2, 283, 284, 8, 11, 1, 2, 284, 21, 3, 2, 2, 2, 285, 286,
	7, 7, 2, 2, 286, 287, 5, 82, 42, 2, 287, 288, 7, 53, 2, 2, 288, 289, 5,
	82, 42, 2, 289, 290, 7, 54, 2, 2, 290, 291, 8, 12, 1, 2, 291, 313, 3, 2,
	2, 2, 292, 293, 7, 7, 2, 2, 293, 294, 5, 82, 42, 2, 294, 295, 7, 53, 2,
	2, 295, 296, 5, 82, 42, 2, 296, 297, 7, 54, 2, 2, 297, 298, 7, 8, 2, 2,
	298, 299, 7, 53, 2, 2, 299, 300, 5, 82, 42, 2, 300, 301, 7, 54, 2, 2, 301,
	302, 8, 12, 1, 2, 302, 313, 3, 2, 2, 2, 303, 304, 7, 7, 2, 2, 304, 305,
	5, 82, 42, 2, 305, 306, 7, 53, 2, 2, 306, 307, 5, 82, 42, 2, 307, 308,
	7, 54, 2, 2, 308, 309, 7, 8, 2, 2, 309, 310, 5, 24, 13, 2, 310, 311, 8,
	12, 1, 2, 311, 313, 3, 2, 2, 2, 312, 285, 3, 2, 2, 2, 312, 292, 3, 2, 2,
	2, 312, 303, 3, 2, 2, 2, 313, 23, 3, 2, 2, 2, 314, 316, 5, 22, 12, 2, 315,
	314, 3, 2, 2, 2, 316, 319, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 318,
	3, 2, 2, 2, 318, 320, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 320, 321, 8, 13,
	1, 2, 321, 25, 3, 2, 2, 2, 322, 323, 7, 9, 2, 2, 323, 324, 5, 82, 42, 2,
	324, 325, 7, 53, 2, 2, 325, 326, 5, 28, 15, 2, 326, 327, 5, 40, 21, 2,
	327, 328, 7, 54, 2, 2, 328, 329, 8, 14, 1, 2, 329, 338, 3, 2, 2, 2, 330,
	331, 7, 9, 2, 2, 331, 332, 5, 82, 42, 2, 332, 333, 7, 53, 2, 2, 333, 334,
	5, 40, 21, 2, 334, 335, 7, 54, 2, 2, 335, 336, 8, 14, 1, 2, 336, 338, 3,
	2, 2, 2, 337, 322, 3, 2, 2, 2, 337, 330, 3, 2, 2, 2, 338, 27, 3, 2, 2,
	2, 339, 341, 5, 30, 16, 2, 340, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2,
	342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	345, 8, 15, 1, 2, 345, 29, 3, 2, 2, 2, 346, 347, 5, 32, 17, 2, 347, 348,
	7, 40, 2, 2, 348, 349, 7, 53, 2, 2, 349, 350, 5, 4, 3, 2, 350, 351, 7,
	54, 2, 2, 351, 352, 8, 16, 1, 2, 352, 360, 3, 2, 2, 2, 353, 354, 5, 32,
	17, 2, 354, 355, 7, 40, 2, 2, 355, 356, 5, 36, 19, 2, 356, 357, 7, 35,
	2, 2, 357, 358, 8, 16, 1, 2, 358, 360, 3, 2, 2, 2, 359, 346, 3, 2, 2, 2,
	359, 353, 3, 2, 2, 2, 360, 31, 3, 2, 2, 2, 361, 363, 5, 34, 18, 2, 362,
	361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365,
	3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 8, 17, 1, 2, 367, 33, 3, 2,
	2, 2, 368, 369, 5, 82, 42, 2, 369, 370, 7, 59, 2, 2, 370, 371, 8, 18, 1,
	2, 371, 376, 3, 2, 2, 2, 372, 373, 5, 82, 42, 2, 373, 374, 8, 18, 1, 2,
	374, 376, 3, 2, 2, 2, 375, 368, 3, 2, 2, 2, 375, 372, 3, 2, 2, 2, 376,
	35, 3, 2, 2, 2, 377, 378, 5, 8, 5, 2, 378, 379, 8, 19, 1, 2, 379, 37, 3,
	2, 2, 2, 380, 381, 7, 31, 2, 2, 381, 382, 7, 40, 2, 2, 382, 383, 7, 53,
	2, 2, 383, 384, 5, 4, 3, 2, 384, 385, 7, 54, 2, 2, 385, 386, 8, 20, 1,
	2, 386, 394, 3, 2, 2, 2, 387, 388, 7, 31, 2, 2, 388, 389, 7, 40, 2, 2,
	389, 390, 5, 36, 19, 2, 390, 391, 7, 35, 2, 2, 391, 392, 8, 20, 1, 2, 392,
	394, 3, 2, 2, 2, 393, 380, 3, 2, 2, 2, 393, 387, 3, 2, 2, 2, 394, 39, 3,
	2, 2, 2, 395, 397, 5, 38, 20, 2, 396, 395, 3, 2, 2, 2, 397, 398, 3, 2,
	2, 2, 398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2,
	400, 401, 8, 21, 1, 2, 401, 41, 3, 2, 2, 2, 402, 403, 7, 9, 2, 2, 403,
	404, 5, 82, 42, 2, 404, 405, 7, 53, 2, 2, 405, 406, 5, 44, 23, 2, 406,
	407, 5, 52, 27, 2, 407, 408, 7, 54, 2, 2, 408, 409, 8, 22, 1, 2, 409, 418,
	3, 2, 2, 2, 410, 411, 7, 9, 2, 2, 411, 412, 5, 82, 42, 2, 412, 413, 7,
	53, 2, 2, 413, 414, 5, 52, 27, 2, 414, 415, 7, 54, 2, 2, 415, 416, 8, 22,
	1, 2, 416, 418, 3, 2, 2, 2, 417, 402, 3, 2, 2, 2, 417, 410, 3, 2, 2, 2,
	418, 43, 3, 2, 2, 2, 419, 421, 5, 46, 24, 2, 420, 419, 3, 2, 2, 2, 421,
	422, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423, 424,
	3, 2, 2, 2, 424, 425, 8, 23, 1, 2, 425, 45, 3, 2, 2, 2, 426, 427, 5, 48,
	25, 2, 427, 428, 7, 40, 2, 2, 428, 429, 5, 82, 42, 2, 429, 430, 7, 35,
	2, 2, 430, 431, 8, 24, 1, 2, 431, 440, 3, 2, 2, 2, 432, 433, 5, 48, 25,
	2, 433, 434, 7, 40, 2, 2, 434, 435, 7, 53, 2, 2, 435, 436, 5, 82, 42, 2,
	436, 437, 7, 54, 2, 2, 437, 438, 8, 24, 1, 2, 438, 440, 3, 2, 2, 2, 439,
	426, 3, 2, 2, 2, 439, 432, 3, 2, 2, 2, 440, 47, 3, 2, 2, 2, 441, 443, 5,
	50, 26, 2, 442, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 442, 3, 2,
	2, 2, 444, 445, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 447, 8, 25, 1, 2,
	447, 49, 3, 2, 2, 2, 448, 449, 5, 82, 42, 2, 449, 450, 7, 59, 2, 2, 450,
	451, 8, 26, 1, 2, 451, 456, 3, 2, 2, 2, 452, 453, 5, 82, 42, 2, 453, 454,
	8, 26, 1, 2, 454, 456, 3, 2, 2, 2, 455, 448, 3, 2, 2, 2, 455, 452, 3, 2,
	2, 2, 456, 51, 3, 2, 2, 2, 457, 458, 7, 31, 2, 2, 458, 459, 7, 40, 2, 2,
	459, 460, 5, 82, 42, 2, 460, 461, 7, 35, 2, 2, 461, 462, 8, 27, 1, 2, 462,
	471, 3, 2, 2, 2, 463, 464, 7, 31, 2, 2, 464, 465, 7, 40, 2, 2, 465, 466,
	7, 53, 2, 2, 466, 467, 5, 82, 42, 2, 467, 468, 7, 54, 2, 2, 468, 469, 8,
	27, 1, 2, 469, 471, 3, 2, 2, 2, 470, 457, 3, 2, 2, 2, 470, 463, 3, 2, 2,
	2, 471, 53, 3, 2, 2, 2, 472, 473, 7, 10, 2, 2, 473, 474, 5, 82, 42, 2,
	474, 475, 7, 53, 2, 2, 475, 476, 5, 4, 3, 2, 476, 477, 7, 54, 2, 2, 477,
	478, 8, 28, 1, 2, 478, 55, 3, 2, 2, 2, 479, 480, 7, 12, 2, 2, 480, 481,
	7, 31, 2, 2, 481, 482, 7, 13, 2, 2, 482, 483, 5, 82, 42, 2, 483, 484, 7,
	32, 2, 2, 484, 485, 5, 82, 42, 2, 485, 486, 7, 53, 2, 2, 486, 487, 5, 4,
	3, 2, 487, 488, 7, 54, 2, 2, 488, 489, 8, 29, 1, 2, 489, 500, 3, 2, 2,
	2, 490, 491, 7, 12, 2, 2, 491, 492, 7, 31, 2, 2, 492, 493, 7, 13, 2, 2,
	493, 494, 5, 82, 42, 2, 494, 495, 7, 53, 2, 2, 495, 496, 5, 4, 3, 2, 496,
	497, 7, 54, 2, 2, 497, 498, 8, 29, 1, 2, 498, 500, 3, 2, 2, 2, 499, 479,
	3, 2, 2, 2, 499, 490, 3, 2, 2, 2, 500, 57, 3, 2, 2, 2, 501, 502, 7, 11,
	2, 2, 502, 503, 7, 53, 2, 2, 503, 504, 5, 4, 3, 2, 504, 505, 7, 54, 2,
	2, 505, 506, 8, 30, 1, 2, 506, 59, 3, 2, 2, 2, 507, 508, 7, 11, 2, 2, 508,
	509, 7, 53, 2, 2, 509, 510, 5, 4, 3, 2, 510, 511, 7, 54, 2, 2, 511, 512,
	8, 31, 1, 2, 512, 61, 3, 2, 2, 2, 513, 514, 7, 14, 2, 2, 514, 520, 8, 32,
	1, 2, 515, 516, 7, 14, 2, 2, 516, 517, 5, 82, 42, 2, 517, 518, 8, 32, 1,
	2, 518, 520, 3, 2, 2, 2, 519, 513, 3, 2, 2, 2, 519, 515, 3, 2, 2, 2, 520,
	63, 3, 2, 2, 2, 521, 522, 7, 15, 2, 2, 522, 523, 8, 33, 1, 2, 523, 65,
	3, 2, 2, 2, 524, 525, 7, 16, 2, 2, 525, 526, 5, 82, 42, 2, 526, 527, 8,
	34, 1, 2, 527, 67, 3, 2, 2, 2, 528, 529, 7, 18, 2, 2, 529, 530, 7, 31,
	2, 2, 530, 531, 7, 51, 2, 2, 531, 532, 7, 52, 2, 2, 532, 533, 7, 53, 2,
	2, 533, 534, 5, 4, 3, 2, 534, 535, 7, 54, 2, 2, 535, 536, 8, 35, 1, 2,
	536, 560, 3, 2, 2, 2, 537, 538, 7, 18, 2, 2, 538, 539, 7, 31, 2, 2, 539,
	540, 7, 51, 2, 2, 540, 541, 5, 70, 36, 2, 541, 542, 7, 52, 2, 2, 542, 543,
	7, 53, 2, 2, 543, 544, 5, 4, 3, 2, 544, 545, 7, 54, 2, 2, 545, 546, 8,
	35, 1, 2, 546, 560, 3, 2, 2, 2, 547, 548, 7, 18, 2, 2, 548, 549, 7, 31,
	2, 2, 549, 550, 7, 51, 2, 2, 550, 551, 5, 70, 36, 2, 551, 552, 7, 52, 2,
	2, 552, 553, 7, 41, 2, 2, 553, 554, 5, 76, 39, 2, 554, 555, 7, 53, 2, 2,
	555, 556, 5, 4, 3, 2, 556, 557, 7, 54, 2, 2, 557, 558, 8, 35, 1, 2, 558,
	560, 3, 2, 2, 2, 559, 528, 3, 2, 2, 2, 559, 537, 3, 2, 2, 2, 559, 547,
	3, 2, 2, 2, 560, 69, 3, 2, 2, 2, 561, 563, 5, 72, 37, 2, 562, 561, 3, 2,
	2, 2, 563, 566, 3, 2, 2, 2, 564, 562, 3, 2, 2, 2, 564, 565, 3, 2, 2, 2,
	565, 567, 3, 2, 2, 2, 566, 564, 3, 2, 2, 2, 567, 568, 8, 36, 1, 2, 568,
	71, 3, 2, 2, 2, 569, 570, 7, 31, 2, 2, 570, 571, 7, 36, 2, 2, 571, 572,
	5, 76, 39, 2, 572, 573, 7, 35, 2, 2, 573, 574, 8, 37, 1, 2, 574, 581, 3,
	2, 2, 2, 575, 576, 7, 31, 2, 2, 576, 577, 7, 36, 2, 2, 577, 578, 5, 76,
	39, 2, 578, 579, 8, 37, 1, 2, 579, 581, 3, 2, 2, 2, 580, 569, 3, 2, 2,
	2, 580, 575, 3, 2, 2, 2, 581, 73, 3, 2, 2, 2, 582, 583, 7, 31, 2, 2, 583,
	584, 7, 51, 2, 2, 584, 585, 7, 52, 2, 2, 585, 593, 8, 38, 1, 2, 586, 587,
	7, 31, 2, 2, 587, 588, 7, 51, 2, 2, 588, 589, 5, 78, 40, 2, 589, 590, 7,
	52, 2, 2, 590, 591, 8, 38, 1, 2, 591, 593, 3, 2, 2, 2, 592, 582, 3, 2,
	2, 2, 592, 586, 3, 2, 2, 2, 593, 75, 3, 2, 2, 2, 594, 595, 7, 20, 2, 2,
	595, 607, 8, 39, 1, 2, 596, 597, 7, 21, 2, 2, 597, 607, 8, 39, 1, 2, 598,
	599, 7, 22, 2, 2, 599, 607, 8, 39, 1, 2, 600, 601, 7, 25, 2, 2, 601, 607,
	8, 39, 1, 2, 602, 603, 7, 23, 2, 2, 603, 607, 8, 39, 1, 2, 604, 605, 7,
	24, 2, 2, 605, 607, 8, 39, 1, 2, 606, 594, 3, 2, 2, 2, 606, 596, 3, 2,
	2, 2, 606, 598, 3, 2, 2, 2, 606, 600, 3, 2, 2, 2, 606, 602, 3, 2, 2, 2,
	606, 604, 3, 2, 2, 2, 607, 77, 3, 2, 2, 2, 608, 610, 5, 80, 41, 2, 609,
	608, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 609, 3, 2, 2, 2, 611, 612,
	3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 614, 8, 40, 1, 2, 614, 79, 3, 2,
	2, 2, 615, 616, 5, 82, 42, 2, 616, 617, 7, 35, 2, 2, 617, 618, 8, 41, 1,
	2, 618, 623, 3, 2, 2, 2, 619, 620, 5, 82, 42, 2, 620, 621, 8, 41, 1, 2,
	621, 623, 3, 2, 2, 2, 622, 615, 3, 2, 2, 2, 622, 619, 3, 2, 2, 2, 623,
	81, 3, 2, 2, 2, 624, 625, 5, 84, 43, 2, 625, 626, 8, 42, 1, 2, 626, 637,
	3, 2, 2, 2, 627, 628, 5, 86, 44, 2, 628, 629, 8, 42, 1, 2, 629, 637, 3,
	2, 2, 2, 630, 631, 5, 88, 45, 2, 631, 632, 8, 42, 1, 2, 632, 637, 3, 2,
	2, 2, 633, 634, 5, 96, 49, 2, 634, 635, 8, 42, 1, 2, 635, 637, 3, 2, 2,
	2, 636, 624, 3, 2, 2, 2, 636, 627, 3, 2, 2, 2, 636, 630, 3, 2, 2, 2, 636,
	633, 3, 2, 2, 2, 637, 83, 3, 2, 2, 2, 638, 639, 8, 43, 1, 2, 639, 640,
	7, 60, 2, 2, 640, 641, 5, 84, 43, 5, 641, 642, 8, 43, 1, 2, 642, 647, 3,
	2, 2, 2, 643, 644, 5, 86, 44, 2, 644, 645, 8, 43, 1, 2, 645, 647, 3, 2,
	2, 2, 646, 638, 3, 2, 2, 2, 646, 643, 3, 2, 2, 2, 647, 655, 3, 2, 2, 2,
	648, 649, 12, 4, 2, 2, 649, 650, 9, 2, 2, 2, 650, 651, 5, 84, 43, 5, 651,
	652, 8, 43, 1, 2, 652, 654, 3, 2, 2, 2, 653, 648, 3, 2, 2, 2, 654, 657,
	3, 2, 2, 2, 655, 653, 3, 2, 2, 2, 655, 656, 3, 2, 2, 2, 656, 85, 3, 2,
	2, 2, 657, 655, 3, 2, 2, 2, 658, 659, 8, 44, 1, 2, 659, 660, 5, 88, 45,
	2, 660, 661, 8, 44, 1, 2, 661, 669, 3, 2, 2, 2, 662, 663, 12, 4, 2, 2,
	663, 664, 9, 3, 2, 2, 664, 665, 5, 86, 44, 5, 665, 666, 8, 44, 1, 2, 666,
	668, 3, 2, 2, 2, 667, 662, 3, 2, 2, 2, 668, 671, 3, 2, 2, 2, 669, 667,
	3, 2, 2, 2, 669, 670, 3, 2, 2, 2, 670, 87, 3, 2, 2, 2, 671, 669, 3, 2,
	2, 2, 672, 673, 8, 45, 1, 2, 673, 674, 7, 50, 2, 2, 674, 675, 5, 88, 45,
	7, 675, 676, 8, 45, 1, 2, 676, 686, 3, 2, 2, 2, 677, 678, 5, 90, 46, 2,
	678, 679, 8, 45, 1, 2, 679, 686, 3, 2, 2, 2, 680, 681, 7, 51, 2, 2, 681,
	682, 5, 82, 42, 2, 682, 683, 7, 52, 2, 2, 683, 684, 8, 45, 1, 2, 684, 686,
	3, 2, 2, 2, 685, 672, 3, 2, 2, 2, 685, 677, 3, 2, 2, 2, 685, 680, 3, 2,
	2, 2, 686, 699, 3, 2, 2, 2, 687, 688, 12, 6, 2, 2, 688, 689, 9, 4, 2, 2,
	689, 690, 5, 88, 45, 7, 690, 691, 8, 45, 1, 2, 691, 698, 3, 2, 2, 2, 692,
	693, 12, 5, 2, 2, 693, 694, 9, 5, 2, 2, 694, 695, 5, 88, 45, 6, 695, 696,
	8, 45, 1, 2, 696, 698, 3, 2, 2, 2, 697, 687, 3, 2, 2, 2, 697, 692, 3, 2,
	2, 2, 698, 701, 3, 2, 2, 2, 699, 697, 3, 2, 2, 2, 699, 700, 3, 2, 2, 2,
	700, 89, 3, 2, 2, 2, 701, 699, 3, 2, 2, 2, 702, 703, 5, 92, 47, 2, 703,
	704, 8, 46, 1, 2, 704, 91, 3, 2, 2, 2, 705, 706, 7, 26, 2, 2, 706, 730,
	8, 47, 1, 2, 707, 708, 7, 27, 2, 2, 708, 730, 8, 47, 1, 2, 709, 710, 7,
	29, 2, 2, 710, 730, 8, 47, 1, 2, 711, 712, 7, 30, 2, 2, 712, 730, 8, 47,
	1, 2, 713, 714, 7, 28, 2, 2, 714, 730, 8, 47, 1, 2, 715, 716, 7, 31, 2,
	2, 716, 730, 8, 47, 1, 2, 717, 718, 5, 94, 48, 2, 718, 719, 8, 47, 1, 2,
	719, 730, 3, 2, 2, 2, 720, 721, 5, 22, 12, 2, 721, 722, 8, 47, 1, 2, 722,
	730, 3, 2, 2, 2, 723, 724, 5, 42, 22, 2, 724, 725, 8, 47, 1, 2, 725, 730,
	3, 2, 2, 2, 726, 727, 5, 60, 31, 2, 727, 728, 8, 47, 1, 2, 728, 730, 3,
	2, 2, 2, 729, 705, 3, 2, 2, 2, 729, 707, 3, 2, 2, 2, 729, 709, 3, 2, 2,
	2, 729, 711, 3, 2, 2, 2, 729, 713, 3, 2, 2, 2, 729, 715, 3, 2, 2, 2, 729,
	717, 3, 2, 2, 2, 729, 720, 3, 2, 2, 2, 729, 723, 3, 2, 2, 2, 729, 726,
	3, 2, 2, 2, 730, 93, 3, 2, 2, 2, 731, 732, 7, 26, 2, 2, 732, 733, 7, 19,
	2, 2, 733, 734, 7, 20, 2, 2, 734, 756, 8, 48, 1, 2, 735, 736, 7, 26, 2,
	2, 736, 737, 7, 19, 2, 2, 737, 738, 7, 21, 2, 2, 738, 756, 8, 48, 1, 2,
	739, 740, 7, 27, 2, 2, 740, 741, 7, 19, 2, 2, 741, 742, 7, 20, 2, 2, 742,
	756, 8, 48, 1, 2, 743, 744, 7, 27, 2, 2, 744, 745, 7, 19, 2, 2, 745, 746,
	7, 21, 2, 2, 746, 756, 8, 48, 1, 2, 747, 748, 7, 30, 2, 2, 748, 749, 7,
	19, 2, 2, 749, 750, 7, 20, 2, 2, 750, 756, 8, 48, 1, 2, 751, 752, 7, 30,
	2, 2, 752, 753, 7, 19, 2, 2, 753, 754, 7, 21, 2, 2, 754, 756, 8, 48, 1,
	2, 755, 731, 3, 2, 2, 2, 755, 735, 3, 2, 2, 2, 755, 739, 3, 2, 2, 2, 755,
	743, 3, 2, 2, 2, 755, 747, 3, 2, 2, 2, 755, 751, 3, 2, 2, 2, 756, 95, 3,
	2, 2, 2, 757, 758, 5, 88, 45, 2, 758, 759, 7, 19, 2, 2, 759, 760, 5, 98,
	50, 2, 760, 761, 8, 49, 1, 2, 761, 97, 3, 2, 2, 2, 762, 763, 7, 20, 2,
	2, 763, 767, 8, 50, 1, 2, 764, 765, 7, 21, 2, 2, 765, 767, 8, 50, 1, 2,
	766, 762, 3, 2, 2, 2, 766, 764, 3, 2, 2, 2, 767, 99, 3, 2, 2, 2, 43, 106,
	113, 162, 180, 240, 275, 280, 312, 317, 337, 342, 359, 364, 375, 393, 398,
	417, 422, 439, 444, 455, 470, 499, 519, 559, 564, 580, 592, 606, 611, 622,
	636, 646, 655, 669, 685, 697, 699, 729, 755, 766,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'print!'", "'println!'", "'let'", "'mut'", "'if'", "'else'", "'match'",
	"'while'", "'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'",
	"'main'", "'fn'", "'as'", "'i64'", "'f64'", "'String'", "'bool'", "'char'",
	"'&str'", "", "", "", "", "", "", "'..'", "'.'", "';'", "','", "':'", "'='",
	"'=='", "'>='", "'=>'", "'->'", "'<='", "'!='", "'>'", "'<'", "'*'", "'/'",
	"'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'",
	"'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "R_LET", "R_MUT", "R_IF", "R_ELSE", "R_MATCH",
	"R_WHILE", "R_LOOP", "R_FOR", "R_IN", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_MAIN", "R_FUNCTION", "R_AS", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"R_CHAR", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID",
	"TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS",
	"TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR",
	"TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV",
	"TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC",
	"TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE",
	"TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_println",
	"instr_main", "instr_declaracion", "instr_asignacion", "instr_if", "instr_else_if",
	"instr_ternario", "instr_else_if_ternario", "instr_match", "list_case",
	"instr_case", "list_expre_case", "block_case", "block_instr_match", "instr_default",
	"block_default", "instr_match_ter", "list_case_ternario", "instr_case_ter",
	"list_expre_case_ter", "block_case_ter", "instr_default_ter", "instr_while",
	"instr_for_in", "instr_loop", "instr_loop_ternario", "instr_break", "instr_continue",
	"instr_return", "instr_func", "list_function_parameters", "block_parameters_fn",
	"instr_llamada", "instr_tipo", "list_expression", "block_list_expression",
	"expressions", "expre_log", "expre_rel", "expre_arit", "expre_valor", "primitivo",
	"primitivo_casteo", "expre_casteo", "type_casteo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsR_PRINT       = 1
	ChemsR_PRINTLN     = 2
	ChemsR_LET         = 3
	ChemsR_MUT         = 4
	ChemsR_IF          = 5
	ChemsR_ELSE        = 6
	ChemsR_MATCH       = 7
	ChemsR_WHILE       = 8
	ChemsR_LOOP        = 9
	ChemsR_FOR         = 10
	ChemsR_IN          = 11
	ChemsR_BREAK       = 12
	ChemsR_CONTINUE    = 13
	ChemsR_RETURN      = 14
	ChemsR_MAIN        = 15
	ChemsR_FUNCTION    = 16
	ChemsR_AS          = 17
	ChemsR_INT         = 18
	ChemsR_FLOAT       = 19
	ChemsR_STRING      = 20
	ChemsR_BOOL        = 21
	ChemsR_CHAR        = 22
	ChemsR_STR         = 23
	ChemsNUMBER        = 24
	ChemsDOUBLE        = 25
	ChemsCHAR          = 26
	ChemsSTRING        = 27
	ChemsBOOLEAN       = 28
	ChemsID            = 29
	ChemsTK_DOBLEPUNTO = 30
	ChemsTK_PUNTO      = 31
	ChemsTK_PUNTOCOMA  = 32
	ChemsTK_COMA       = 33
	ChemsTK_DOSPUNTOS  = 34
	ChemsTK_IGUAL      = 35
	ChemsTK_IGUALIGUAL = 36
	ChemsTK_MAYORIGUAL = 37
	ChemsTK_IGUALMAYOR = 38
	ChemsTK_MENOSMAYOR = 39
	ChemsTK_MENORIGUAL = 40
	ChemsTK_DIFIGUAL   = 41
	ChemsTK_MAYOR      = 42
	ChemsTK_MENOR      = 43
	ChemsTK_MULT       = 44
	ChemsTK_DIV        = 45
	ChemsTK_MODULO     = 46
	ChemsTK_MAS        = 47
	ChemsTK_MENOS      = 48
	ChemsTK_PARA       = 49
	ChemsTK_PARC       = 50
	ChemsTK_LLAVEA     = 51
	ChemsTK_LLAVEC     = 52
	ChemsTK_CORA       = 53
	ChemsTK_CORC       = 54
	ChemsTK_AND        = 55
	ChemsTK_OR         = 56
	ChemsTK_BARRA      = 57
	ChemsTK_NOT        = 58
	ChemsWHITESPACE    = 59
	ChemsTK_MULTI      = 60
	ChemsTK_LINE       = 61
)

// Chems rules.
const (
	ChemsRULE_start                    = 0
	ChemsRULE_instrucciones            = 1
	ChemsRULE_end_instr                = 2
	ChemsRULE_instruccion              = 3
	ChemsRULE_instr_println            = 4
	ChemsRULE_instr_main               = 5
	ChemsRULE_instr_declaracion        = 6
	ChemsRULE_instr_asignacion         = 7
	ChemsRULE_instr_if                 = 8
	ChemsRULE_instr_else_if            = 9
	ChemsRULE_instr_ternario           = 10
	ChemsRULE_instr_else_if_ternario   = 11
	ChemsRULE_instr_match              = 12
	ChemsRULE_list_case                = 13
	ChemsRULE_instr_case               = 14
	ChemsRULE_list_expre_case          = 15
	ChemsRULE_block_case               = 16
	ChemsRULE_block_instr_match        = 17
	ChemsRULE_instr_default            = 18
	ChemsRULE_block_default            = 19
	ChemsRULE_instr_match_ter          = 20
	ChemsRULE_list_case_ternario       = 21
	ChemsRULE_instr_case_ter           = 22
	ChemsRULE_list_expre_case_ter      = 23
	ChemsRULE_block_case_ter           = 24
	ChemsRULE_instr_default_ter        = 25
	ChemsRULE_instr_while              = 26
	ChemsRULE_instr_for_in             = 27
	ChemsRULE_instr_loop               = 28
	ChemsRULE_instr_loop_ternario      = 29
	ChemsRULE_instr_break              = 30
	ChemsRULE_instr_continue           = 31
	ChemsRULE_instr_return             = 32
	ChemsRULE_instr_func               = 33
	ChemsRULE_list_function_parameters = 34
	ChemsRULE_block_parameters_fn      = 35
	ChemsRULE_instr_llamada            = 36
	ChemsRULE_instr_tipo               = 37
	ChemsRULE_list_expression          = 38
	ChemsRULE_block_list_expression    = 39
	ChemsRULE_expressions              = 40
	ChemsRULE_expre_log                = 41
	ChemsRULE_expre_rel                = 42
	ChemsRULE_expre_arit               = 43
	ChemsRULE_expre_valor              = 44
	ChemsRULE_primitivo                = 45
	ChemsRULE_primitivo_casteo         = 46
	ChemsRULE_expre_casteo             = 47
	ChemsRULE_type_casteo              = 48
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetList() []IInstruccionContext { return s.list }

func (s *InstruccionesContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_WHILE)|(1<<ChemsR_LOOP)|(1<<ChemsR_FOR)|(1<<ChemsR_BREAK)|(1<<ChemsR_CONTINUE)|(1<<ChemsR_RETURN)|(1<<ChemsR_FUNCTION)|(1<<ChemsID))) != 0) {
		{
			p.SetState(101)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IEnd_instrContext is an interface to support dynamic dispatch.
type IEnd_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEnd_instrContext differentiates from other interfaces.
	IsEnd_instrContext()
}

type End_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
}

func NewEmptyEnd_instrContext() *End_instrContext {
	var p = new(End_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_end_instr
	return p
}

func (*End_instrContext) IsEnd_instrContext() {}

func NewEnd_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *End_instrContext {
	var p = new(End_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_end_instr

	return p
}

func (s *End_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *End_instrContext) GetV() int { return s.v }

func (s *End_instrContext) SetV(v int) { s.v = v }

func (s *End_instrContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *End_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *End_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *End_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterEnd_instr(s)
	}
}

func (s *End_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitEnd_instr(s)
	}
}

func (p *Chems) End_instr() (localctx IEnd_instrContext) {
	localctx = NewEnd_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_end_instr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_LET, ChemsR_IF, ChemsR_MATCH, ChemsR_WHILE, ChemsR_LOOP, ChemsR_FOR, ChemsR_BREAK, ChemsR_CONTINUE, ChemsR_RETURN, ChemsR_FUNCTION, ChemsID, ChemsTK_COMA, ChemsTK_LLAVEC:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*End_instrContext).v = 0

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_println returns the _instr_println rule contexts.
	Get_instr_println() IInstr_printlnContext

	// Get_instr_main returns the _instr_main rule contexts.
	Get_instr_main() IInstr_mainContext

	// Get_instr_declaracion returns the _instr_declaracion rule contexts.
	Get_instr_declaracion() IInstr_declaracionContext

	// Get_instr_asignacion returns the _instr_asignacion rule contexts.
	Get_instr_asignacion() IInstr_asignacionContext

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Get_instr_match returns the _instr_match rule contexts.
	Get_instr_match() IInstr_matchContext

	// Get_instr_while returns the _instr_while rule contexts.
	Get_instr_while() IInstr_whileContext

	// Get_instr_for_in returns the _instr_for_in rule contexts.
	Get_instr_for_in() IInstr_for_inContext

	// Get_instr_loop returns the _instr_loop rule contexts.
	Get_instr_loop() IInstr_loopContext

	// Get_instr_break returns the _instr_break rule contexts.
	Get_instr_break() IInstr_breakContext

	// Get_instr_continue returns the _instr_continue rule contexts.
	Get_instr_continue() IInstr_continueContext

	// Get_instr_func returns the _instr_func rule contexts.
	Get_instr_func() IInstr_funcContext

	// Get_instr_llamada returns the _instr_llamada rule contexts.
	Get_instr_llamada() IInstr_llamadaContext

	// Get_instr_return returns the _instr_return rule contexts.
	Get_instr_return() IInstr_returnContext

	// Set_instr_println sets the _instr_println rule contexts.
	Set_instr_println(IInstr_printlnContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// Set_instr_match sets the _instr_match rule contexts.
	Set_instr_match(IInstr_matchContext)

	// Set_instr_while sets the _instr_while rule contexts.
	Set_instr_while(IInstr_whileContext)

	// Set_instr_for_in sets the _instr_for_in rule contexts.
	Set_instr_for_in(IInstr_for_inContext)

	// Set_instr_loop sets the _instr_loop rule contexts.
	Set_instr_loop(IInstr_loopContext)

	// Set_instr_break sets the _instr_break rule contexts.
	Set_instr_break(IInstr_breakContext)

	// Set_instr_continue sets the _instr_continue rule contexts.
	Set_instr_continue(IInstr_continueContext)

	// Set_instr_func sets the _instr_func rule contexts.
	Set_instr_func(IInstr_funcContext)

	// Set_instr_llamada sets the _instr_llamada rule contexts.
	Set_instr_llamada(IInstr_llamadaContext)

	// Set_instr_return sets the _instr_return rule contexts.
	Set_instr_return(IInstr_returnContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_instr_println     IInstr_printlnContext
	_instr_main        IInstr_mainContext
	_instr_declaracion IInstr_declaracionContext
	_instr_asignacion  IInstr_asignacionContext
	_instr_if          IInstr_ifContext
	_instr_match       IInstr_matchContext
	_instr_while       IInstr_whileContext
	_instr_for_in      IInstr_for_inContext
	_instr_loop        IInstr_loopContext
	_instr_break       IInstr_breakContext
	_instr_continue    IInstr_continueContext
	_instr_func        IInstr_funcContext
	_instr_llamada     IInstr_llamadaContext
	_instr_return      IInstr_returnContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_instr_println() IInstr_printlnContext { return s._instr_println }

func (s *InstruccionContext) Get_instr_main() IInstr_mainContext { return s._instr_main }

func (s *InstruccionContext) Get_instr_declaracion() IInstr_declaracionContext {
	return s._instr_declaracion
}

func (s *InstruccionContext) Get_instr_asignacion() IInstr_asignacionContext {
	return s._instr_asignacion
}

func (s *InstruccionContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *InstruccionContext) Get_instr_match() IInstr_matchContext { return s._instr_match }

func (s *InstruccionContext) Get_instr_while() IInstr_whileContext { return s._instr_while }

func (s *InstruccionContext) Get_instr_for_in() IInstr_for_inContext { return s._instr_for_in }

func (s *InstruccionContext) Get_instr_loop() IInstr_loopContext { return s._instr_loop }

func (s *InstruccionContext) Get_instr_break() IInstr_breakContext { return s._instr_break }

func (s *InstruccionContext) Get_instr_continue() IInstr_continueContext { return s._instr_continue }

func (s *InstruccionContext) Get_instr_func() IInstr_funcContext { return s._instr_func }

func (s *InstruccionContext) Get_instr_llamada() IInstr_llamadaContext { return s._instr_llamada }

func (s *InstruccionContext) Get_instr_return() IInstr_returnContext { return s._instr_return }

func (s *InstruccionContext) Set_instr_println(v IInstr_printlnContext) { s._instr_println = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

func (s *InstruccionContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *InstruccionContext) Set_instr_match(v IInstr_matchContext) { s._instr_match = v }

func (s *InstruccionContext) Set_instr_while(v IInstr_whileContext) { s._instr_while = v }

func (s *InstruccionContext) Set_instr_for_in(v IInstr_for_inContext) { s._instr_for_in = v }

func (s *InstruccionContext) Set_instr_loop(v IInstr_loopContext) { s._instr_loop = v }

func (s *InstruccionContext) Set_instr_break(v IInstr_breakContext) { s._instr_break = v }

func (s *InstruccionContext) Set_instr_continue(v IInstr_continueContext) { s._instr_continue = v }

func (s *InstruccionContext) Set_instr_func(v IInstr_funcContext) { s._instr_func = v }

func (s *InstruccionContext) Set_instr_llamada(v IInstr_llamadaContext) { s._instr_llamada = v }

func (s *InstruccionContext) Set_instr_return(v IInstr_returnContext) { s._instr_return = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) Instr_println() IInstr_printlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_printlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_printlnContext)
}

func (s *InstruccionContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *InstruccionContext) Instr_main() IInstr_mainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_mainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_mainContext)
}

func (s *InstruccionContext) Instr_declaracion() IInstr_declaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_declaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_declaracionContext)
}

func (s *InstruccionContext) Instr_asignacion() IInstr_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_asignacionContext)
}

func (s *InstruccionContext) Instr_if() IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *InstruccionContext) Instr_match() IInstr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_matchContext)
}

func (s *InstruccionContext) Instr_while() IInstr_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_whileContext)
}

func (s *InstruccionContext) Instr_for_in() IInstr_for_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_for_inContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_for_inContext)
}

func (s *InstruccionContext) Instr_loop() IInstr_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loopContext)
}

func (s *InstruccionContext) Instr_break() IInstr_breakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_breakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_breakContext)
}

func (s *InstruccionContext) Instr_continue() IInstr_continueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_continueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_continueContext)
}

func (s *InstruccionContext) Instr_func() IInstr_funcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_funcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_funcContext)
}

func (s *InstruccionContext) Instr_llamada() IInstr_llamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_llamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_llamadaContext)
}

func (s *InstruccionContext) Instr_return() IInstr_returnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_returnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_returnContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(114)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)

			var _x = p.Instr_while()

			localctx.(*InstruccionContext)._instr_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_while().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)

			var _x = p.Instr_for_in()

			localctx.(*InstruccionContext)._instr_for_in = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_for_in().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(138)

			var _x = p.Instr_loop()

			localctx.(*InstruccionContext)._instr_loop = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_loop().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(141)

			var _x = p.Instr_break()

			localctx.(*InstruccionContext)._instr_break = _x
		}
		{
			p.SetState(142)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_break().GetInstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(145)

			var _x = p.Instr_continue()

			localctx.(*InstruccionContext)._instr_continue = _x
		}
		{
			p.SetState(146)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_continue().GetInstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(149)

			var _x = p.Instr_func()

			localctx.(*InstruccionContext)._instr_func = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_func().GetInstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(152)

			var _x = p.Instr_llamada()

			localctx.(*InstruccionContext)._instr_llamada = _x
		}
		{
			p.SetState(153)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_llamada().GetInstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(156)

			var _x = p.Instr_return()

			localctx.(*InstruccionContext)._instr_return = _x
		}
		{
			p.SetState(157)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_return().GetInstr()

	}

	return localctx
}

// IInstr_printlnContext is an interface to support dynamic dispatch.
type IInstr_printlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_PRINTLN returns the _R_PRINTLN token.
	Get_R_PRINTLN() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_R_PRINTLN sets the _R_PRINTLN token.
	Set_R_PRINTLN(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_printlnContext differentiates from other interfaces.
	IsInstr_printlnContext()
}

type Instr_printlnContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_R_PRINTLN       antlr.Token
	_primitivo       IPrimitivoContext
	_STRING          antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_printlnContext() *Instr_printlnContext {
	var p = new(Instr_printlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_println
	return p
}

func (*Instr_printlnContext) IsInstr_printlnContext() {}

func NewInstr_printlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_printlnContext {
	var p = new(Instr_printlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_println

	return p
}

func (s *Instr_printlnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_printlnContext) Get_R_PRINTLN() antlr.Token { return s._R_PRINTLN }

func (s *Instr_printlnContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Instr_printlnContext) Set_R_PRINTLN(v antlr.Token) { s._R_PRINTLN = v }

func (s *Instr_printlnContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Instr_printlnContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Instr_printlnContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_printlnContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Instr_printlnContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_printlnContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *Instr_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_printlnContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Instr_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_printlnContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *Instr_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_printlnContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_printlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_printlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_printlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_println(s)
	}
}

func (s *Instr_printlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_println(s)
	}
}

func (p *Chems) Instr_println() (localctx IInstr_printlnContext) {
	localctx = NewInstr_printlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_instr_println)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(163)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(164)

			var _x = p.Primitivo()

			localctx.(*Instr_printlnContext)._primitivo = _x
		}
		{
			p.SetState(165)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(166)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(nil, localctx.(*Instr_printlnContext).Get_primitivo().GetP(), "-1", (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(170)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(171)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(172)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(173)

			var _x = p.List_expression()

			localctx.(*Instr_printlnContext)._list_expression = _x
		}
		{
			p.SetState(174)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(175)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(localctx.(*Instr_printlnContext).Get_list_expression().GetL(), nil, (func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}())[1:len((func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}()))-1], (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	}

	return localctx
}

// IInstr_mainContext is an interface to support dynamic dispatch.
type IInstr_mainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MAIN returns the _R_MAIN token.
	Get_R_MAIN() antlr.Token

	// Set_R_MAIN sets the _R_MAIN token.
	Set_R_MAIN(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_mainContext differentiates from other interfaces.
	IsInstr_mainContext()
}

type Instr_mainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MAIN        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_mainContext() *Instr_mainContext {
	var p = new(Instr_mainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_main
	return p
}

func (*Instr_mainContext) IsInstr_mainContext() {}

func NewInstr_mainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_mainContext {
	var p = new(Instr_mainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_main

	return p
}

func (s *Instr_mainContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_mainContext) Get_R_MAIN() antlr.Token { return s._R_MAIN }

func (s *Instr_mainContext) Set_R_MAIN(v antlr.Token) { s._R_MAIN = v }

func (s *Instr_mainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_mainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_mainContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_mainContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_mainContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_mainContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(ChemsR_MAIN, 0)
}

func (s *Instr_mainContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_mainContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_mainContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_mainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_mainContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_mainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_mainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_mainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_main(s)
	}
}

func (s *Instr_mainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_main(s)
	}
}

func (p *Chems) Instr_main() (localctx IInstr_mainContext) {
	localctx = NewInstr_mainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_instr_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(181)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(182)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(183)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(184)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(185)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(186)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_mainContext).instr = instruction.NewMain(localctx.(*Instr_mainContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_mainContext).Get_R_MAIN() == nil {
			return 0
		} else {
			return localctx.(*Instr_mainContext).Get_R_MAIN().GetLine()
		}
	}()), localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn())

	return localctx
}

// IInstr_declaracionContext is an interface to support dynamic dispatch.
type IInstr_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LET returns the _R_LET token.
	Get_R_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_LET sets the _R_LET token.
	Set_R_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_declaracionContext differentiates from other interfaces.
	IsInstr_declaracionContext()
}

type Instr_declaracionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_LET       antlr.Token
	_ID          antlr.Token
	_expressions IExpressionsContext
	_instr_tipo  IInstr_tipoContext
}

func NewEmptyInstr_declaracionContext() *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_declaracion
	return p
}

func (*Instr_declaracionContext) IsInstr_declaracionContext() {}

func NewInstr_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_declaracion

	return p
}

func (s *Instr_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_declaracionContext) Get_R_LET() antlr.Token { return s._R_LET }

func (s *Instr_declaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_declaracionContext) Set_R_LET(v antlr.Token) { s._R_LET = v }

func (s *Instr_declaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_declaracionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_declaracionContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_declaracionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_declaracionContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_declaracionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_declaracionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_declaracionContext) R_LET() antlr.TerminalNode {
	return s.GetToken(ChemsR_LET, 0)
}

func (s *Instr_declaracionContext) R_MUT() antlr.TerminalNode {
	return s.GetToken(ChemsR_MUT, 0)
}

func (s *Instr_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_declaracionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_declaracionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_declaracionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_declaracionContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Instr_declaracionContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_declaracion(s)
	}
}

func (s *Instr_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_declaracion(s)
	}
}

func (p *Chems) Instr_declaracion() (localctx IInstr_declaracionContext) {
	localctx = NewInstr_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_instr_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(190)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(191)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(192)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(193)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(194)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(198)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(199)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(200)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(201)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(202)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), nil, true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(206)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(207)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(208)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(209)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(210)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(211)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(212)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(215)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(216)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(217)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(218)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(219)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(223)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(224)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(225)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(226)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), nil, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(229)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(230)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(231)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(232)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(233)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(234)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(235)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	}

	return localctx
}

// IInstr_asignacionContext is an interface to support dynamic dispatch.
type IInstr_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_asignacionContext differentiates from other interfaces.
	IsInstr_asignacionContext()
}

type Instr_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_ID          antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_asignacionContext() *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_asignacion
	return p
}

func (*Instr_asignacionContext) IsInstr_asignacionContext() {}

func NewInstr_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_asignacion

	return p
}

func (s *Instr_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_asignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_asignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_asignacionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_asignacionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_asignacionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_asignacionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_asignacionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_asignacionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_asignacionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_asignacion(s)
	}
}

func (s *Instr_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_asignacion(s)
	}
}

func (p *Chems) Instr_asignacion() (localctx IInstr_asignacionContext) {
	localctx = NewInstr_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ChemsRULE_instr_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(241)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(242)

		var _x = p.Expressions()

		localctx.(*Instr_asignacionContext)._expressions = _x
	}
	{
		p.SetState(243)
		p.Match(ChemsTK_PUNTOCOMA)
	}
	localctx.(*Instr_asignacionContext).instr = variable.NewAssignment((func() string {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_asignacionContext).Get_expressions().GetP(), (func() int {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instr_asignacionContext).Get_ID().GetColumn())

	return localctx
}

// IInstr_ifContext is an interface to support dynamic dispatch.
type IInstr_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IInstruccionesContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IInstruccionesContext

	// Get_instr_else_if returns the _instr_else_if rule contexts.
	Get_instr_else_if() IInstr_else_ifContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IInstruccionesContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IInstruccionesContext)

	// Set_instr_else_if sets the _instr_else_if rule contexts.
	Set_instr_else_if(IInstr_else_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_ifContext differentiates from other interfaces.
	IsInstr_ifContext()
}

type Instr_ifContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_IF          antlr.Token
	_expressions   IExpressionsContext
	left_instr     IInstruccionesContext
	right_intr     IInstruccionesContext
	_instr_else_if IInstr_else_ifContext
}

func NewEmptyInstr_ifContext() *Instr_ifContext {
	var p = new(Instr_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_if
	return p
}

func (*Instr_ifContext) IsInstr_ifContext() {}

func NewInstr_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ifContext {
	var p = new(Instr_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_if

	return p
}

func (s *Instr_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ifContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ifContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ifContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_ifContext) GetLeft_instr() IInstruccionesContext { return s.left_instr }

func (s *Instr_ifContext) GetRight_intr() IInstruccionesContext { return s.right_intr }

func (s *Instr_ifContext) Get_instr_else_if() IInstr_else_ifContext { return s._instr_else_if }

func (s *Instr_ifContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_ifContext) SetLeft_instr(v IInstruccionesContext) { s.left_instr = v }

func (s *Instr_ifContext) SetRight_intr(v IInstruccionesContext) { s.right_intr = v }

func (s *Instr_ifContext) Set_instr_else_if(v IInstr_else_ifContext) { s._instr_else_if = v }

func (s *Instr_ifContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_ifContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_ifContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ifContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_ifContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ifContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ifContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ifContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ifContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *Instr_ifContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_ifContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ifContext) Instr_else_if() IInstr_else_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_ifContext)
}

func (s *Instr_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_if(s)
	}
}

func (s *Instr_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_if(s)
	}
}

func (p *Chems) Instr_if() (localctx IInstr_ifContext) {
	localctx = NewInstr_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_instr_if)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(247)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(248)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(249)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(250)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(254)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(255)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(256)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(257)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(258)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(259)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(260)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(261)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), localctx.(*Instr_ifContext).GetRight_intr().GetL(), nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(265)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(266)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(267)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(268)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(269)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(270)

			var _x = p.Instr_else_if()

			localctx.(*Instr_ifContext)._instr_else_if = _x
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, localctx.(*Instr_ifContext).Get_instr_else_if().GetL(), (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	}

	return localctx
}

// IInstr_else_ifContext is an interface to support dynamic dispatch.
type IInstr_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ifContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ifContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_ifContext differentiates from other interfaces.
	IsInstr_else_ifContext()
}

type Instr_else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	_instr_if IInstr_ifContext
	e         []IInstr_ifContext
}

func NewEmptyInstr_else_ifContext() *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if
	return p
}

func (*Instr_else_ifContext) IsInstr_else_ifContext() {}

func NewInstr_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if

	return p
}

func (s *Instr_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_ifContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *Instr_else_ifContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *Instr_else_ifContext) GetE() []IInstr_ifContext { return s.e }

func (s *Instr_else_ifContext) SetE(v []IInstr_ifContext) { s.e = v }

func (s *Instr_else_ifContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_ifContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_ifContext) AllInstr_if() []IInstr_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem())
	var tst = make([]IInstr_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ifContext)
		}
	}

	return tst
}

func (s *Instr_else_ifContext) Instr_if(i int) IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *Instr_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if(s)
	}
}

func (s *Instr_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if(s)
	}
}

func (p *Chems) Instr_else_if() (localctx IInstr_else_ifContext) {
	localctx = NewInstr_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChemsRULE_instr_else_if)

	localctx.(*Instr_else_ifContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(275)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_ifContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_ifContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_ternarioContext is an interface to support dynamic dispatch.
type IInstr_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IExpressionsContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IExpressionsContext

	// Get_instr_else_if_ternario returns the _instr_else_if_ternario rule contexts.
	Get_instr_else_if_ternario() IInstr_else_if_ternarioContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IExpressionsContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IExpressionsContext)

	// Set_instr_else_if_ternario sets the _instr_else_if_ternario rule contexts.
	Set_instr_else_if_ternario(IInstr_else_if_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_ternarioContext differentiates from other interfaces.
	IsInstr_ternarioContext()
}

type Instr_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser                  antlr.Parser
	p                       interfaces.Expression
	_R_IF                   antlr.Token
	cond                    IExpressionsContext
	left_instr              IExpressionsContext
	right_intr              IExpressionsContext
	_instr_else_if_ternario IInstr_else_if_ternarioContext
}

func NewEmptyInstr_ternarioContext() *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_ternario
	return p
}

func (*Instr_ternarioContext) IsInstr_ternarioContext() {}

func NewInstr_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_ternario

	return p
}

func (s *Instr_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ternarioContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ternarioContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ternarioContext) GetCond() IExpressionsContext { return s.cond }

func (s *Instr_ternarioContext) GetLeft_instr() IExpressionsContext { return s.left_instr }

func (s *Instr_ternarioContext) GetRight_intr() IExpressionsContext { return s.right_intr }

func (s *Instr_ternarioContext) Get_instr_else_if_ternario() IInstr_else_if_ternarioContext {
	return s._instr_else_if_ternario
}

func (s *Instr_ternarioContext) SetCond(v IExpressionsContext) { s.cond = v }

func (s *Instr_ternarioContext) SetLeft_instr(v IExpressionsContext) { s.left_instr = v }

func (s *Instr_ternarioContext) SetRight_intr(v IExpressionsContext) { s.right_intr = v }

func (s *Instr_ternarioContext) Set_instr_else_if_ternario(v IInstr_else_if_ternarioContext) {
	s._instr_else_if_ternario = v
}

func (s *Instr_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_ternarioContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ternarioContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ternarioContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ternarioContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ternarioContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ternarioContext) AllExpressions() []IExpressionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionsContext)(nil)).Elem())
	var tst = make([]IExpressionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionsContext)
		}
	}

	return tst
}

func (s *Instr_ternarioContext) Expressions(i int) IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_ternarioContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ternarioContext) Instr_else_if_ternario() IInstr_else_if_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_if_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_if_ternarioContext)
}

func (s *Instr_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_ternario(s)
	}
}

func (s *Instr_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_ternario(s)
	}
}

func (p *Chems) Instr_ternario() (localctx IInstr_ternarioContext) {
	localctx = NewInstr_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChemsRULE_instr_ternario)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(283)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(284)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(285)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(286)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(287)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(291)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(292)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(293)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(294)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(295)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(296)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(297)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(298)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), localctx.(*Instr_ternarioContext).GetRight_intr().GetP(), nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(302)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(303)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(304)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(305)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(306)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(307)

			var _x = p.Instr_else_if_ternario()

			localctx.(*Instr_ternarioContext)._instr_else_if_ternario = _x
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, localctx.(*Instr_ternarioContext).Get_instr_else_if_ternario().GetL(), (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	}

	return localctx
}

// IInstr_else_if_ternarioContext is an interface to support dynamic dispatch.
type IInstr_else_if_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ternarioContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ternarioContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_if_ternarioContext differentiates from other interfaces.
	IsInstr_else_if_ternarioContext()
}

type Instr_else_if_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_ternario IInstr_ternarioContext
	e               []IInstr_ternarioContext
}

func NewEmptyInstr_else_if_ternarioContext() *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if_ternario
	return p
}

func (*Instr_else_if_ternarioContext) IsInstr_else_if_ternarioContext() {}

func NewInstr_else_if_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if_ternario

	return p
}

func (s *Instr_else_if_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_if_ternarioContext) Get_instr_ternario() IInstr_ternarioContext {
	return s._instr_ternario
}

func (s *Instr_else_if_ternarioContext) Set_instr_ternario(v IInstr_ternarioContext) {
	s._instr_ternario = v
}

func (s *Instr_else_if_ternarioContext) GetE() []IInstr_ternarioContext { return s.e }

func (s *Instr_else_if_ternarioContext) SetE(v []IInstr_ternarioContext) { s.e = v }

func (s *Instr_else_if_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_if_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_if_ternarioContext) AllInstr_ternario() []IInstr_ternarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem())
	var tst = make([]IInstr_ternarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ternarioContext)
		}
	}

	return tst
}

func (s *Instr_else_if_ternarioContext) Instr_ternario(i int) IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *Instr_else_if_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_if_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_if_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if_ternario(s)
	}
}

func (s *Instr_else_if_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if_ternario(s)
	}
}

func (p *Chems) Instr_else_if_ternario() (localctx IInstr_else_if_ternarioContext) {
	localctx = NewInstr_else_if_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChemsRULE_instr_else_if_ternario)

	localctx.(*Instr_else_if_ternarioContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(312)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_if_ternarioContext).l.Add(e.GetP())
	}

	return localctx
}

// IInstr_matchContext is an interface to support dynamic dispatch.
type IInstr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_list_case returns the _list_case rule contexts.
	Get_list_case() IList_caseContext

	// Get_block_default returns the _block_default rule contexts.
	Get_block_default() IBlock_defaultContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_list_case sets the _list_case rule contexts.
	Set_list_case(IList_caseContext)

	// Set_block_default sets the _block_default rule contexts.
	Set_block_default(IBlock_defaultContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_matchContext differentiates from other interfaces.
	IsInstr_matchContext()
}

type Instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MATCH       antlr.Token
	_expressions   IExpressionsContext
	_list_case     IList_caseContext
	_block_default IBlock_defaultContext
}

func NewEmptyInstr_matchContext() *Instr_matchContext {
	var p = new(Instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match
	return p
}

func (*Instr_matchContext) IsInstr_matchContext() {}

func NewInstr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_matchContext {
	var p = new(Instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match

	return p
}

func (s *Instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_matchContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_matchContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_matchContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_matchContext) Get_list_case() IList_caseContext { return s._list_case }

func (s *Instr_matchContext) Get_block_default() IBlock_defaultContext { return s._block_default }

func (s *Instr_matchContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_matchContext) Set_list_case(v IList_caseContext) { s._list_case = v }

func (s *Instr_matchContext) Set_block_default(v IBlock_defaultContext) { s._block_default = v }

func (s *Instr_matchContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_matchContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_matchContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_matchContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_matchContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_matchContext) List_case() IList_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_caseContext)
}

func (s *Instr_matchContext) Block_default() IBlock_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_defaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_defaultContext)
}

func (s *Instr_matchContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match(s)
	}
}

func (s *Instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match(s)
	}
}

func (p *Chems) Instr_match() (localctx IInstr_matchContext) {
	localctx = NewInstr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChemsRULE_instr_match)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(321)

			var _x = p.Expressions()

			localctx.(*Instr_matchContext)._expressions = _x
		}
		{
			p.SetState(322)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(323)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(324)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(325)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expressions().GetP(), localctx.(*Instr_matchContext).Get_list_case().GetL(), localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(329)

			var _x = p.Expressions()

			localctx.(*Instr_matchContext)._expressions = _x
		}
		{
			p.SetState(330)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(331)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(332)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expressions().GetP(), nil, localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn())

	}

	return localctx
}

// IList_caseContext is an interface to support dynamic dispatch.
type IList_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case returns the _instr_case rule contexts.
	Get_instr_case() IInstr_caseContext

	// Set_instr_case sets the _instr_case rule contexts.
	Set_instr_case(IInstr_caseContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_caseContext

	// SetE sets the e rule context list.
	SetE([]IInstr_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_caseContext differentiates from other interfaces.
	IsList_caseContext()
}

type List_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_instr_case IInstr_caseContext
	e           []IInstr_caseContext
}

func NewEmptyList_caseContext() *List_caseContext {
	var p = new(List_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case
	return p
}

func (*List_caseContext) IsList_caseContext() {}

func NewList_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_caseContext {
	var p = new(List_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case

	return p
}

func (s *List_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_caseContext) Get_instr_case() IInstr_caseContext { return s._instr_case }

func (s *List_caseContext) Set_instr_case(v IInstr_caseContext) { s._instr_case = v }

func (s *List_caseContext) GetE() []IInstr_caseContext { return s.e }

func (s *List_caseContext) SetE(v []IInstr_caseContext) { s.e = v }

func (s *List_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_caseContext) AllInstr_case() []IInstr_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem())
	var tst = make([]IInstr_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_caseContext)
		}
	}

	return tst
}

func (s *List_caseContext) Instr_case(i int) IInstr_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_caseContext)
}

func (s *List_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case(s)
	}
}

func (s *List_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case(s)
	}
}

func (p *Chems) List_case() (localctx IList_caseContext) {
	localctx = NewList_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChemsRULE_list_case)

	localctx.(*List_caseContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(337)

				var _x = p.Instr_case()

				localctx.(*List_caseContext)._instr_case = _x
			}
			localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_caseContext is an interface to support dynamic dispatch.
type IInstr_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case returns the _list_expre_case rule contexts.
	Get_list_expre_case() IList_expre_caseContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_list_expre_case sets the _list_expre_case rule contexts.
	Set_list_expre_case(IList_expre_caseContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expression

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expression)

	// IsInstr_caseContext differentiates from other interfaces.
	IsInstr_caseContext()
}

type Instr_caseContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Expression
	_list_expre_case   IList_expre_caseContext
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_caseContext() *Instr_caseContext {
	var p = new(Instr_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case
	return p
}

func (*Instr_caseContext) IsInstr_caseContext() {}

func NewInstr_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_caseContext {
	var p = new(Instr_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case

	return p
}

func (s *Instr_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_caseContext) Get_list_expre_case() IList_expre_caseContext { return s._list_expre_case }

func (s *Instr_caseContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_caseContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_caseContext) Set_list_expre_case(v IList_expre_caseContext) { s._list_expre_case = v }

func (s *Instr_caseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_caseContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_caseContext) GetInstr() interfaces.Expression { return s.instr }

func (s *Instr_caseContext) SetInstr(v interfaces.Expression) { s.instr = v }

func (s *Instr_caseContext) List_expre_case() IList_expre_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_caseContext)
}

func (s *Instr_caseContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_caseContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_caseContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_caseContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_caseContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_caseContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case(s)
	}
}

func (s *Instr_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case(s)
	}
}

func (p *Chems) Instr_case() (localctx IInstr_caseContext) {
	localctx = NewInstr_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_instr_case)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(345)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(346)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(347)

			var _x = p.Instrucciones()

			localctx.(*Instr_caseContext)._instrucciones = _x
		}
		{
			p.SetState(348)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(352)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(353)

			var _x = p.Block_instr_match()

			localctx.(*Instr_caseContext)._block_instr_match = _x
		}
		{
			p.SetState(354)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IList_expre_caseContext is an interface to support dynamic dispatch.
type IList_expre_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case returns the _block_case rule contexts.
	Get_block_case() IBlock_caseContext

	// Set_block_case sets the _block_case rule contexts.
	Set_block_case(IBlock_caseContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_caseContext

	// SetE sets the e rule context list.
	SetE([]IBlock_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_caseContext differentiates from other interfaces.
	IsList_expre_caseContext()
}

type List_expre_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_block_case IBlock_caseContext
	e           []IBlock_caseContext
}

func NewEmptyList_expre_caseContext() *List_expre_caseContext {
	var p = new(List_expre_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case
	return p
}

func (*List_expre_caseContext) IsList_expre_caseContext() {}

func NewList_expre_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_caseContext {
	var p = new(List_expre_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case

	return p
}

func (s *List_expre_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_caseContext) Get_block_case() IBlock_caseContext { return s._block_case }

func (s *List_expre_caseContext) Set_block_case(v IBlock_caseContext) { s._block_case = v }

func (s *List_expre_caseContext) GetE() []IBlock_caseContext { return s.e }

func (s *List_expre_caseContext) SetE(v []IBlock_caseContext) { s.e = v }

func (s *List_expre_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_caseContext) AllBlock_case() []IBlock_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem())
	var tst = make([]IBlock_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_caseContext)
		}
	}

	return tst
}

func (s *List_expre_caseContext) Block_case(i int) IBlock_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_caseContext)
}

func (s *List_expre_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case(s)
	}
}

func (s *List_expre_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case(s)
	}
}

func (p *Chems) List_expre_case() (localctx IList_expre_caseContext) {
	localctx = NewList_expre_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ChemsRULE_list_expre_case)

	localctx.(*List_expre_caseContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(359)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_caseContext is an interface to support dynamic dispatch.
type IBlock_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expression

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expression)

	// IsBlock_caseContext differentiates from other interfaces.
	IsBlock_caseContext()
}

type Block_caseContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_caseContext() *Block_caseContext {
	var p = new(Block_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case
	return p
}

func (*Block_caseContext) IsBlock_caseContext() {}

func NewBlock_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_caseContext {
	var p = new(Block_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case

	return p
}

func (s *Block_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_caseContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_caseContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_caseContext) GetInstr() interfaces.Expression { return s.instr }

func (s *Block_caseContext) SetInstr(v interfaces.Expression) { s.instr = v }

func (s *Block_caseContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_caseContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case(s)
	}
}

func (s *Block_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case(s)
	}
}

func (p *Chems) Block_case() (localctx IBlock_caseContext) {
	localctx = NewBlock_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_block_case)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)

			var _x = p.Expressions()

			localctx.(*Block_caseContext)._expressions = _x
		}
		{
			p.SetState(367)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expressions().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)

			var _x = p.Expressions()

			localctx.(*Block_caseContext)._expressions = _x
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expressions().GetP(), nil, nil)

	}

	return localctx
}

// IBlock_instr_matchContext is an interface to support dynamic dispatch.
type IBlock_instr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_instr_matchContext differentiates from other interfaces.
	IsBlock_instr_matchContext()
}

type Block_instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyBlock_instr_matchContext() *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_instr_match
	return p
}

func (*Block_instr_matchContext) IsBlock_instr_matchContext() {}

func NewBlock_instr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_instr_match

	return p
}

func (s *Block_instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_instr_matchContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *Block_instr_matchContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *Block_instr_matchContext) GetList() []IInstruccionContext { return s.list }

func (s *Block_instr_matchContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *Block_instr_matchContext) GetL() *arrayList.List { return s.l }

func (s *Block_instr_matchContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_instr_matchContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *Block_instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_instr_match(s)
	}
}

func (s *Block_instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_instr_match(s)
	}
}

func (p *Chems) Block_instr_match() (localctx IBlock_instr_matchContext) {
	localctx = NewBlock_instr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ChemsRULE_block_instr_match)

	localctx.(*Block_instr_matchContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)

		var _x = p.Instruccion()

		localctx.(*Block_instr_matchContext)._instruccion = _x
	}
	localctx.(*Block_instr_matchContext).list = append(localctx.(*Block_instr_matchContext).list, localctx.(*Block_instr_matchContext)._instruccion)

	listInt := localctx.(*Block_instr_matchContext).GetList()
	for _, e := range listInt {
		localctx.(*Block_instr_matchContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_defaultContext is an interface to support dynamic dispatch.
type IInstr_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_defaultContext differentiates from other interfaces.
	IsInstr_defaultContext()
}

type Instr_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_defaultContext() *Instr_defaultContext {
	var p = new(Instr_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default
	return p
}

func (*Instr_defaultContext) IsInstr_defaultContext() {}

func NewInstr_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_defaultContext {
	var p = new(Instr_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default

	return p
}

func (s *Instr_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_defaultContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_defaultContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_defaultContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_defaultContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_defaultContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_defaultContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_defaultContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_defaultContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_defaultContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_defaultContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_defaultContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_defaultContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_defaultContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default(s)
	}
}

func (s *Instr_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default(s)
	}
}

func (p *Chems) Instr_default() (localctx IInstr_defaultContext) {
	localctx = NewInstr_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ChemsRULE_instr_default)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Match(ChemsID)
		}
		{
			p.SetState(379)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(380)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(381)

			var _x = p.Instrucciones()

			localctx.(*Instr_defaultContext)._instrucciones = _x
		}
		{
			p.SetState(382)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Match(ChemsID)
		}
		{
			p.SetState(386)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(387)

			var _x = p.Block_instr_match()

			localctx.(*Instr_defaultContext)._block_instr_match = _x
		}
		{
			p.SetState(388)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IBlock_defaultContext is an interface to support dynamic dispatch.
type IBlock_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_default returns the _instr_default rule contexts.
	Get_instr_default() IInstr_defaultContext

	// Set_instr_default sets the _instr_default rule contexts.
	Set_instr_default(IInstr_defaultContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_defaultContext

	// SetE sets the e rule context list.
	SetE([]IInstr_defaultContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_defaultContext differentiates from other interfaces.
	IsBlock_defaultContext()
}

type Block_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	l              *arrayList.List
	_instr_default IInstr_defaultContext
	e              []IInstr_defaultContext
}

func NewEmptyBlock_defaultContext() *Block_defaultContext {
	var p = new(Block_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_default
	return p
}

func (*Block_defaultContext) IsBlock_defaultContext() {}

func NewBlock_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_defaultContext {
	var p = new(Block_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_default

	return p
}

func (s *Block_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_defaultContext) Get_instr_default() IInstr_defaultContext { return s._instr_default }

func (s *Block_defaultContext) Set_instr_default(v IInstr_defaultContext) { s._instr_default = v }

func (s *Block_defaultContext) GetE() []IInstr_defaultContext { return s.e }

func (s *Block_defaultContext) SetE(v []IInstr_defaultContext) { s.e = v }

func (s *Block_defaultContext) GetL() *arrayList.List { return s.l }

func (s *Block_defaultContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_defaultContext) AllInstr_default() []IInstr_defaultContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem())
	var tst = make([]IInstr_defaultContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_defaultContext)
		}
	}

	return tst
}

func (s *Block_defaultContext) Instr_default(i int) IInstr_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_defaultContext)
}

func (s *Block_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_default(s)
	}
}

func (s *Block_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_default(s)
	}
}

func (p *Chems) Block_default() (localctx IBlock_defaultContext) {
	localctx = NewBlock_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ChemsRULE_block_default)

	localctx.(*Block_defaultContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsID {
		{
			p.SetState(393)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*Block_defaultContext).GetE()
	for _, e := range listInt {
		localctx.(*Block_defaultContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_match_terContext is an interface to support dynamic dispatch.
type IInstr_match_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_list_case_ternario returns the _list_case_ternario rule contexts.
	Get_list_case_ternario() IList_case_ternarioContext

	// Get_instr_default_ter returns the _instr_default_ter rule contexts.
	Get_instr_default_ter() IInstr_default_terContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_list_case_ternario sets the _list_case_ternario rule contexts.
	Set_list_case_ternario(IList_case_ternarioContext)

	// Set_instr_default_ter sets the _instr_default_ter rule contexts.
	Set_instr_default_ter(IInstr_default_terContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_match_terContext differentiates from other interfaces.
	IsInstr_match_terContext()
}

type Instr_match_terContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	p                   interfaces.Expression
	_R_MATCH            antlr.Token
	_expressions        IExpressionsContext
	_list_case_ternario IList_case_ternarioContext
	_instr_default_ter  IInstr_default_terContext
}

func NewEmptyInstr_match_terContext() *Instr_match_terContext {
	var p = new(Instr_match_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match_ter
	return p
}

func (*Instr_match_terContext) IsInstr_match_terContext() {}

func NewInstr_match_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_match_terContext {
	var p = new(Instr_match_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match_ter

	return p
}

func (s *Instr_match_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_match_terContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_match_terContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_match_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_match_terContext) Get_list_case_ternario() IList_case_ternarioContext {
	return s._list_case_ternario
}

func (s *Instr_match_terContext) Get_instr_default_ter() IInstr_default_terContext {
	return s._instr_default_ter
}

func (s *Instr_match_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_match_terContext) Set_list_case_ternario(v IList_case_ternarioContext) {
	s._list_case_ternario = v
}

func (s *Instr_match_terContext) Set_instr_default_ter(v IInstr_default_terContext) {
	s._instr_default_ter = v
}

func (s *Instr_match_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_match_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_match_terContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_match_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_match_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_match_terContext) List_case_ternario() IList_case_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_case_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_case_ternarioContext)
}

func (s *Instr_match_terContext) Instr_default_ter() IInstr_default_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_default_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_default_terContext)
}

func (s *Instr_match_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_match_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_match_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_match_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match_ter(s)
	}
}

func (s *Instr_match_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match_ter(s)
	}
}

func (p *Chems) Instr_match_ter() (localctx IInstr_match_terContext) {
	localctx = NewInstr_match_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ChemsRULE_instr_match_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(401)

			var _x = p.Expressions()

			localctx.(*Instr_match_terContext)._expressions = _x
		}
		{
			p.SetState(402)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(403)

			var _x = p.List_case_ternario()

			localctx.(*Instr_match_terContext)._list_case_ternario = _x
		}
		{
			p.SetState(404)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(405)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).p = ternario.NewMatch(localctx.(*Instr_match_terContext).Get_expressions().GetP(), localctx.(*Instr_match_terContext).Get_list_case_ternario().GetL(), localctx.(*Instr_match_terContext).Get_instr_default_ter().GetP(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(409)

			var _x = p.Expressions()

			localctx.(*Instr_match_terContext)._expressions = _x
		}
		{
			p.SetState(410)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(411)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(412)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).p = ternario.NewMatch(localctx.(*Instr_match_terContext).Get_expressions().GetP(), nil, localctx.(*Instr_match_terContext).Get_instr_default_ter().GetP(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn())

	}

	return localctx
}

// IList_case_ternarioContext is an interface to support dynamic dispatch.
type IList_case_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case_ter returns the _instr_case_ter rule contexts.
	Get_instr_case_ter() IInstr_case_terContext

	// Set_instr_case_ter sets the _instr_case_ter rule contexts.
	Set_instr_case_ter(IInstr_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_case_terContext

	// SetE sets the e rule context list.
	SetE([]IInstr_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_case_ternarioContext differentiates from other interfaces.
	IsList_case_ternarioContext()
}

type List_case_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_case_ter IInstr_case_terContext
	e               []IInstr_case_terContext
}

func NewEmptyList_case_ternarioContext() *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case_ternario
	return p
}

func (*List_case_ternarioContext) IsList_case_ternarioContext() {}

func NewList_case_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case_ternario

	return p
}

func (s *List_case_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *List_case_ternarioContext) Get_instr_case_ter() IInstr_case_terContext {
	return s._instr_case_ter
}

func (s *List_case_ternarioContext) Set_instr_case_ter(v IInstr_case_terContext) {
	s._instr_case_ter = v
}

func (s *List_case_ternarioContext) GetE() []IInstr_case_terContext { return s.e }

func (s *List_case_ternarioContext) SetE(v []IInstr_case_terContext) { s.e = v }

func (s *List_case_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *List_case_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_case_ternarioContext) AllInstr_case_ter() []IInstr_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem())
	var tst = make([]IInstr_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_case_terContext)
		}
	}

	return tst
}

func (s *List_case_ternarioContext) Instr_case_ter(i int) IInstr_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_case_terContext)
}

func (s *List_case_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_case_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_case_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case_ternario(s)
	}
}

func (s *List_case_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case_ternario(s)
	}
}

func (p *Chems) List_case_ternario() (localctx IList_case_ternarioContext) {
	localctx = NewList_case_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ChemsRULE_list_case_ternario)

	localctx.(*List_case_ternarioContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(417)

				var _x = p.Instr_case_ter()

				localctx.(*List_case_ternarioContext)._instr_case_ter = _x
			}
			localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_case_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*List_case_ternarioContext).l.Add(e.GetP())
	}

	return localctx
}

// IInstr_case_terContext is an interface to support dynamic dispatch.
type IInstr_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case_ter returns the _list_expre_case_ter rule contexts.
	Get_list_expre_case_ter() IList_expre_case_terContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_list_expre_case_ter sets the _list_expre_case_ter rule contexts.
	Set_list_expre_case_ter(IList_expre_case_terContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_case_terContext differentiates from other interfaces.
	IsInstr_case_terContext()
}

type Instr_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	p                    interfaces.Expression
	_list_expre_case_ter IList_expre_case_terContext
	_expressions         IExpressionsContext
}

func NewEmptyInstr_case_terContext() *Instr_case_terContext {
	var p = new(Instr_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case_ter
	return p
}

func (*Instr_case_terContext) IsInstr_case_terContext() {}

func NewInstr_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_case_terContext {
	var p = new(Instr_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case_ter

	return p
}

func (s *Instr_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_case_terContext) Get_list_expre_case_ter() IList_expre_case_terContext {
	return s._list_expre_case_ter
}

func (s *Instr_case_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_case_terContext) Set_list_expre_case_ter(v IList_expre_case_terContext) {
	s._list_expre_case_ter = v
}

func (s *Instr_case_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_case_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_case_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_case_terContext) List_expre_case_ter() IList_expre_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_case_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_case_terContext)
}

func (s *Instr_case_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_case_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_case_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_case_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_case_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case_ter(s)
	}
}

func (s *Instr_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case_ter(s)
	}
}

func (p *Chems) Instr_case_ter() (localctx IInstr_case_terContext) {
	localctx = NewInstr_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ChemsRULE_instr_case_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)

			var _x = p.List_expre_case_ter()

			localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
		}
		{
			p.SetState(425)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(426)

			var _x = p.Expressions()

			localctx.(*Instr_case_terContext)._expressions = _x
		}
		{
			p.SetState(427)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_case_terContext).p = ternario.NewCase(nil, localctx.(*Instr_case_terContext).Get_list_expre_case_ter().GetL(), localctx.(*Instr_case_terContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)

			var _x = p.List_expre_case_ter()

			localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
		}
		{
			p.SetState(431)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(432)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(433)

			var _x = p.Expressions()

			localctx.(*Instr_case_terContext)._expressions = _x
		}
		{
			p.SetState(434)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_case_terContext).p = ternario.NewCase(nil, localctx.(*Instr_case_terContext).Get_list_expre_case_ter().GetL(), localctx.(*Instr_case_terContext).Get_expressions().GetP())

	}

	return localctx
}

// IList_expre_case_terContext is an interface to support dynamic dispatch.
type IList_expre_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case_ter returns the _block_case_ter rule contexts.
	Get_block_case_ter() IBlock_case_terContext

	// Set_block_case_ter sets the _block_case_ter rule contexts.
	Set_block_case_ter(IBlock_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_case_terContext

	// SetE sets the e rule context list.
	SetE([]IBlock_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_case_terContext differentiates from other interfaces.
	IsList_expre_case_terContext()
}

type List_expre_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_block_case_ter IBlock_case_terContext
	e               []IBlock_case_terContext
}

func NewEmptyList_expre_case_terContext() *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case_ter
	return p
}

func (*List_expre_case_terContext) IsList_expre_case_terContext() {}

func NewList_expre_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case_ter

	return p
}

func (s *List_expre_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_case_terContext) Get_block_case_ter() IBlock_case_terContext {
	return s._block_case_ter
}

func (s *List_expre_case_terContext) Set_block_case_ter(v IBlock_case_terContext) {
	s._block_case_ter = v
}

func (s *List_expre_case_terContext) GetE() []IBlock_case_terContext { return s.e }

func (s *List_expre_case_terContext) SetE(v []IBlock_case_terContext) { s.e = v }

func (s *List_expre_case_terContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_case_terContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_case_terContext) AllBlock_case_ter() []IBlock_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem())
	var tst = make([]IBlock_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_case_terContext)
		}
	}

	return tst
}

func (s *List_expre_case_terContext) Block_case_ter(i int) IBlock_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_case_terContext)
}

func (s *List_expre_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case_ter(s)
	}
}

func (s *List_expre_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case_ter(s)
	}
}

func (p *Chems) List_expre_case_ter() (localctx IList_expre_case_terContext) {
	localctx = NewList_expre_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ChemsRULE_list_expre_case_ter)

	localctx.(*List_expre_case_terContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(439)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_case_terContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_case_terContext).l.Add(e.GetP())
	}

	return localctx
}

// IBlock_case_terContext is an interface to support dynamic dispatch.
type IBlock_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsBlock_case_terContext differentiates from other interfaces.
	IsBlock_case_terContext()
}

type Block_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_case_terContext() *Block_case_terContext {
	var p = new(Block_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case_ter
	return p
}

func (*Block_case_terContext) IsBlock_case_terContext() {}

func NewBlock_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_case_terContext {
	var p = new(Block_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case_ter

	return p
}

func (s *Block_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_case_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_case_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_case_terContext) GetP() interfaces.Expression { return s.p }

func (s *Block_case_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Block_case_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_case_terContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case_ter(s)
	}
}

func (s *Block_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case_ter(s)
	}
}

func (p *Chems) Block_case_ter() (localctx IBlock_case_terContext) {
	localctx = NewBlock_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ChemsRULE_block_case_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)

			var _x = p.Expressions()

			localctx.(*Block_case_terContext)._expressions = _x
		}
		{
			p.SetState(447)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_case_terContext).p = ternario.NewCase(localctx.(*Block_case_terContext).Get_expressions().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)

			var _x = p.Expressions()

			localctx.(*Block_case_terContext)._expressions = _x
		}
		localctx.(*Block_case_terContext).p = ternario.NewCase(localctx.(*Block_case_terContext).Get_expressions().GetP(), nil, nil)

	}

	return localctx
}

// IInstr_default_terContext is an interface to support dynamic dispatch.
type IInstr_default_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_default_terContext differentiates from other interfaces.
	IsInstr_default_terContext()
}

type Instr_default_terContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyInstr_default_terContext() *Instr_default_terContext {
	var p = new(Instr_default_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default_ter
	return p
}

func (*Instr_default_terContext) IsInstr_default_terContext() {}

func NewInstr_default_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_default_terContext {
	var p = new(Instr_default_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default_ter

	return p
}

func (s *Instr_default_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_default_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_default_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_default_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_default_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_default_terContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_default_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_default_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_default_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_default_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_default_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_default_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_default_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_default_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default_ter(s)
	}
}

func (s *Instr_default_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default_ter(s)
	}
}

func (p *Chems) Instr_default_ter() (localctx IInstr_default_terContext) {
	localctx = NewInstr_default_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ChemsRULE_instr_default_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(455)
			p.Match(ChemsID)
		}
		{
			p.SetState(456)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(457)

			var _x = p.Expressions()

			localctx.(*Instr_default_terContext)._expressions = _x
		}
		{
			p.SetState(458)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_default_terContext).p = ternario.NewDefault(localctx.(*Instr_default_terContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.Match(ChemsID)
		}
		{
			p.SetState(462)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(463)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(464)

			var _x = p.Expressions()

			localctx.(*Instr_default_terContext)._expressions = _x
		}
		{
			p.SetState(465)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_default_terContext).p = ternario.NewDefault(localctx.(*Instr_default_terContext).Get_expressions().GetP())

	}

	return localctx
}

// IInstr_whileContext is an interface to support dynamic dispatch.
type IInstr_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_WHILE returns the _R_WHILE token.
	Get_R_WHILE() antlr.Token

	// Set_R_WHILE sets the _R_WHILE token.
	Set_R_WHILE(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_whileContext differentiates from other interfaces.
	IsInstr_whileContext()
}

type Instr_whileContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_WHILE       antlr.Token
	_expressions   IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_whileContext() *Instr_whileContext {
	var p = new(Instr_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_while
	return p
}

func (*Instr_whileContext) IsInstr_whileContext() {}

func NewInstr_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_whileContext {
	var p = new(Instr_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_while

	return p
}

func (s *Instr_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_whileContext) Get_R_WHILE() antlr.Token { return s._R_WHILE }

func (s *Instr_whileContext) Set_R_WHILE(v antlr.Token) { s._R_WHILE = v }

func (s *Instr_whileContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_whileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_whileContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_whileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_whileContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_whileContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_whileContext) R_WHILE() antlr.TerminalNode {
	return s.GetToken(ChemsR_WHILE, 0)
}

func (s *Instr_whileContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_whileContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_whileContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_whileContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_while(s)
	}
}

func (s *Instr_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_while(s)
	}
}

func (p *Chems) Instr_while() (localctx IInstr_whileContext) {
	localctx = NewInstr_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ChemsRULE_instr_while)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)

		var _m = p.Match(ChemsR_WHILE)

		localctx.(*Instr_whileContext)._R_WHILE = _m
	}
	{
		p.SetState(471)

		var _x = p.Expressions()

		localctx.(*Instr_whileContext)._expressions = _x
	}
	{
		p.SetState(472)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(473)

		var _x = p.Instrucciones()

		localctx.(*Instr_whileContext)._instrucciones = _x
	}
	{
		p.SetState(474)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_whileContext).instr = loops.NewWhile(localctx.(*Instr_whileContext).Get_expressions().GetP(), localctx.(*Instr_whileContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_whileContext).Get_R_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Instr_whileContext).Get_R_WHILE().GetLine()
		}
	}()), localctx.(*Instr_whileContext).Get_R_WHILE().GetColumn())

	return localctx
}

// IInstr_for_inContext is an interface to support dynamic dispatch.
type IInstr_for_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FOR returns the _R_FOR token.
	Get_R_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FOR sets the _R_FOR token.
	Set_R_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_for_inContext differentiates from other interfaces.
	IsInstr_for_inContext()
}

type Instr_for_inContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_FOR         antlr.Token
	_ID            antlr.Token
	left           IExpressionsContext
	right          IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_for_inContext() *Instr_for_inContext {
	var p = new(Instr_for_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_for_in
	return p
}

func (*Instr_for_inContext) IsInstr_for_inContext() {}

func NewInstr_for_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_for_inContext {
	var p = new(Instr_for_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_for_in

	return p
}

func (s *Instr_for_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_for_inContext) Get_R_FOR() antlr.Token { return s._R_FOR }

func (s *Instr_for_inContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_for_inContext) Set_R_FOR(v antlr.Token) { s._R_FOR = v }

func (s *Instr_for_inContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_for_inContext) GetLeft() IExpressionsContext { return s.left }

func (s *Instr_for_inContext) GetRight() IExpressionsContext { return s.right }

func (s *Instr_for_inContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_for_inContext) SetLeft(v IExpressionsContext) { s.left = v }

func (s *Instr_for_inContext) SetRight(v IExpressionsContext) { s.right = v }

func (s *Instr_for_inContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_for_inContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_for_inContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_for_inContext) R_FOR() antlr.TerminalNode {
	return s.GetToken(ChemsR_FOR, 0)
}

func (s *Instr_for_inContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_for_inContext) R_IN() antlr.TerminalNode {
	return s.GetToken(ChemsR_IN, 0)
}

func (s *Instr_for_inContext) TK_DOBLEPUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOBLEPUNTO, 0)
}

func (s *Instr_for_inContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_for_inContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_for_inContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_for_inContext) AllExpressions() []IExpressionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionsContext)(nil)).Elem())
	var tst = make([]IExpressionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionsContext)
		}
	}

	return tst
}

func (s *Instr_for_inContext) Expressions(i int) IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_for_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_for_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_for_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_for_in(s)
	}
}

func (s *Instr_for_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_for_in(s)
	}
}

func (p *Chems) Instr_for_in() (localctx IInstr_for_inContext) {
	localctx = NewInstr_for_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ChemsRULE_instr_for_in)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(478)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(479)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(480)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(481)
			p.Match(ChemsTK_DOBLEPUNTO)
		}
		{
			p.SetState(482)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).right = _x
		}
		{
			p.SetState(483)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(484)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(485)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.INTEGER, localctx.(*Instr_for_inContext).GetLeft().GetP(), localctx.(*Instr_for_inContext).GetRight().GetP(), localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(488)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(489)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(490)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(491)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(492)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(493)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(494)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.STRING, localctx.(*Instr_for_inContext).GetLeft().GetP(), nil, localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	}

	return localctx
}

// IInstr_loopContext is an interface to support dynamic dispatch.
type IInstr_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LOOP returns the _R_LOOP token.
	Get_R_LOOP() antlr.Token

	// Set_R_LOOP sets the _R_LOOP token.
	Set_R_LOOP(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_loopContext differentiates from other interfaces.
	IsInstr_loopContext()
}

type Instr_loopContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_LOOP        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loopContext() *Instr_loopContext {
	var p = new(Instr_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop
	return p
}

func (*Instr_loopContext) IsInstr_loopContext() {}

func NewInstr_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loopContext {
	var p = new(Instr_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop

	return p
}

func (s *Instr_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loopContext) Get_R_LOOP() antlr.Token { return s._R_LOOP }

func (s *Instr_loopContext) Set_R_LOOP(v antlr.Token) { s._R_LOOP = v }

func (s *Instr_loopContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_loopContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loopContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_loopContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_loopContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loopContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loopContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loopContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop(s)
	}
}

func (s *Instr_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop(s)
	}
}

func (p *Chems) Instr_loop() (localctx IInstr_loopContext) {
	localctx = NewInstr_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ChemsRULE_instr_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)

		var _m = p.Match(ChemsR_LOOP)

		localctx.(*Instr_loopContext)._R_LOOP = _m
	}
	{
		p.SetState(500)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(501)

		var _x = p.Instrucciones()

		localctx.(*Instr_loopContext)._instrucciones = _x
	}
	{
		p.SetState(502)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loopContext).instr = loops.NewLoop(localctx.(*Instr_loopContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_loopContext).Get_R_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Instr_loopContext).Get_R_LOOP().GetLine()
		}
	}()), localctx.(*Instr_loopContext).Get_R_LOOP().GetColumn())

	return localctx
}

// IInstr_loop_ternarioContext is an interface to support dynamic dispatch.
type IInstr_loop_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LOOP returns the _R_LOOP token.
	Get_R_LOOP() antlr.Token

	// Set_R_LOOP sets the _R_LOOP token.
	Set_R_LOOP(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_loop_ternarioContext differentiates from other interfaces.
	IsInstr_loop_ternarioContext()
}

type Instr_loop_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              interfaces.Expression
	_R_LOOP        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loop_ternarioContext() *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop_ternario
	return p
}

func (*Instr_loop_ternarioContext) IsInstr_loop_ternarioContext() {}

func NewInstr_loop_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop_ternario

	return p
}

func (s *Instr_loop_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loop_ternarioContext) Get_R_LOOP() antlr.Token { return s._R_LOOP }

func (s *Instr_loop_ternarioContext) Set_R_LOOP(v antlr.Token) { s._R_LOOP = v }

func (s *Instr_loop_ternarioContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instr_loop_ternarioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loop_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_loop_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_loop_ternarioContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loop_ternarioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loop_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loop_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loop_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop_ternario(s)
	}
}

func (s *Instr_loop_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop_ternario(s)
	}
}

func (p *Chems) Instr_loop_ternario() (localctx IInstr_loop_ternarioContext) {
	localctx = NewInstr_loop_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ChemsRULE_instr_loop_ternario)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)

		var _m = p.Match(ChemsR_LOOP)

		localctx.(*Instr_loop_ternarioContext)._R_LOOP = _m
	}
	{
		p.SetState(506)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(507)

		var _x = p.Instrucciones()

		localctx.(*Instr_loop_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(508)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loop_ternarioContext).p = ternario.NewLoop(localctx.(*Instr_loop_ternarioContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_loop_ternarioContext).Get_R_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Instr_loop_ternarioContext).Get_R_LOOP().GetLine()
		}
	}()), localctx.(*Instr_loop_ternarioContext).Get_R_LOOP().GetColumn())

	return localctx
}

// IInstr_breakContext is an interface to support dynamic dispatch.
type IInstr_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_BREAK returns the _R_BREAK token.
	Get_R_BREAK() antlr.Token

	// Set_R_BREAK sets the _R_BREAK token.
	Set_R_BREAK(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_breakContext differentiates from other interfaces.
	IsInstr_breakContext()
}

type Instr_breakContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_BREAK     antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_breakContext() *Instr_breakContext {
	var p = new(Instr_breakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_break
	return p
}

func (*Instr_breakContext) IsInstr_breakContext() {}

func NewInstr_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_breakContext {
	var p = new(Instr_breakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_break

	return p
}

func (s *Instr_breakContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_breakContext) Get_R_BREAK() antlr.Token { return s._R_BREAK }

func (s *Instr_breakContext) Set_R_BREAK(v antlr.Token) { s._R_BREAK = v }

func (s *Instr_breakContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_breakContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_breakContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_breakContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_breakContext) R_BREAK() antlr.TerminalNode {
	return s.GetToken(ChemsR_BREAK, 0)
}

func (s *Instr_breakContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_break(s)
	}
}

func (s *Instr_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_break(s)
	}
}

func (p *Chems) Instr_break() (localctx IInstr_breakContext) {
	localctx = NewInstr_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ChemsRULE_instr_break)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(511)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(nil, (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(513)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(514)

			var _x = p.Expressions()

			localctx.(*Instr_breakContext)._expressions = _x
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(localctx.(*Instr_breakContext).Get_expressions().GetP(), (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	}

	return localctx
}

// IInstr_continueContext is an interface to support dynamic dispatch.
type IInstr_continueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_CONTINUE returns the _R_CONTINUE token.
	Get_R_CONTINUE() antlr.Token

	// Set_R_CONTINUE sets the _R_CONTINUE token.
	Set_R_CONTINUE(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_continueContext differentiates from other interfaces.
	IsInstr_continueContext()
}

type Instr_continueContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_R_CONTINUE antlr.Token
}

func NewEmptyInstr_continueContext() *Instr_continueContext {
	var p = new(Instr_continueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_continue
	return p
}

func (*Instr_continueContext) IsInstr_continueContext() {}

func NewInstr_continueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_continueContext {
	var p = new(Instr_continueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_continue

	return p
}

func (s *Instr_continueContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_continueContext) Get_R_CONTINUE() antlr.Token { return s._R_CONTINUE }

func (s *Instr_continueContext) Set_R_CONTINUE(v antlr.Token) { s._R_CONTINUE = v }

func (s *Instr_continueContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_continueContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_continueContext) R_CONTINUE() antlr.TerminalNode {
	return s.GetToken(ChemsR_CONTINUE, 0)
}

func (s *Instr_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_continueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_continue(s)
	}
}

func (s *Instr_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_continue(s)
	}
}

func (p *Chems) Instr_continue() (localctx IInstr_continueContext) {
	localctx = NewInstr_continueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ChemsRULE_instr_continue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)

		var _m = p.Match(ChemsR_CONTINUE)

		localctx.(*Instr_continueContext)._R_CONTINUE = _m
	}
	localctx.(*Instr_continueContext).instr = transferencia.NewContinue((func() int {
		if localctx.(*Instr_continueContext).Get_R_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*Instr_continueContext).Get_R_CONTINUE().GetLine()
		}
	}()), localctx.(*Instr_continueContext).Get_R_CONTINUE().GetColumn())

	return localctx
}

// IInstr_returnContext is an interface to support dynamic dispatch.
type IInstr_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_RETURN returns the _R_RETURN token.
	Get_R_RETURN() antlr.Token

	// Set_R_RETURN sets the _R_RETURN token.
	Set_R_RETURN(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_returnContext differentiates from other interfaces.
	IsInstr_returnContext()
}

type Instr_returnContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_RETURN    antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_returnContext() *Instr_returnContext {
	var p = new(Instr_returnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_return
	return p
}

func (*Instr_returnContext) IsInstr_returnContext() {}

func NewInstr_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_returnContext {
	var p = new(Instr_returnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_return

	return p
}

func (s *Instr_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_returnContext) Get_R_RETURN() antlr.Token { return s._R_RETURN }

func (s *Instr_returnContext) Set_R_RETURN(v antlr.Token) { s._R_RETURN = v }

func (s *Instr_returnContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_returnContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_returnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_returnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_returnContext) R_RETURN() antlr.TerminalNode {
	return s.GetToken(ChemsR_RETURN, 0)
}

func (s *Instr_returnContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_return(s)
	}
}

func (s *Instr_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_return(s)
	}
}

func (p *Chems) Instr_return() (localctx IInstr_returnContext) {
	localctx = NewInstr_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ChemsRULE_instr_return)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)

		var _m = p.Match(ChemsR_RETURN)

		localctx.(*Instr_returnContext)._R_RETURN = _m
	}
	{
		p.SetState(523)

		var _x = p.Expressions()

		localctx.(*Instr_returnContext)._expressions = _x
	}
	localctx.(*Instr_returnContext).instr = transferencia.NewReturn(localctx.(*Instr_returnContext).Get_expressions().GetP(), (func() int {
		if localctx.(*Instr_returnContext).Get_R_RETURN() == nil {
			return 0
		} else {
			return localctx.(*Instr_returnContext).Get_R_RETURN().GetLine()
		}
	}()), localctx.(*Instr_returnContext).Get_R_RETURN().GetColumn())

	return localctx
}

// IInstr_funcContext is an interface to support dynamic dispatch.
type IInstr_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FUNCTION returns the _R_FUNCTION token.
	Get_R_FUNCTION() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FUNCTION sets the _R_FUNCTION token.
	Set_R_FUNCTION(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_list_function_parameters returns the _list_function_parameters rule contexts.
	Get_list_function_parameters() IList_function_parametersContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_list_function_parameters sets the _list_function_parameters rule contexts.
	Set_list_function_parameters(IList_function_parametersContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_funcContext differentiates from other interfaces.
	IsInstr_funcContext()
}

type Instr_funcContext struct {
	*antlr.BaseParserRuleContext
	parser                    antlr.Parser
	instr                     interfaces.Instruction
	_R_FUNCTION               antlr.Token
	_ID                       antlr.Token
	_instrucciones            IInstruccionesContext
	_list_function_parameters IList_function_parametersContext
	_instr_tipo               IInstr_tipoContext
}

func NewEmptyInstr_funcContext() *Instr_funcContext {
	var p = new(Instr_funcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_func
	return p
}

func (*Instr_funcContext) IsInstr_funcContext() {}

func NewInstr_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_funcContext {
	var p = new(Instr_funcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_func

	return p
}

func (s *Instr_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_funcContext) Get_R_FUNCTION() antlr.Token { return s._R_FUNCTION }

func (s *Instr_funcContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_funcContext) Set_R_FUNCTION(v antlr.Token) { s._R_FUNCTION = v }

func (s *Instr_funcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_funcContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_funcContext) Get_list_function_parameters() IList_function_parametersContext {
	return s._list_function_parameters
}

func (s *Instr_funcContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_funcContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_funcContext) Set_list_function_parameters(v IList_function_parametersContext) {
	s._list_function_parameters = v
}

func (s *Instr_funcContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_funcContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_funcContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_funcContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_funcContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_funcContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_funcContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_funcContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_funcContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_funcContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_funcContext) List_function_parameters() IList_function_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_function_parametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_function_parametersContext)
}

func (s *Instr_funcContext) TK_MENOSMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOSMAYOR, 0)
}

func (s *Instr_funcContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_func(s)
	}
}

func (s *Instr_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_func(s)
	}
}

func (p *Chems) Instr_func() (localctx IInstr_funcContext) {
	localctx = NewInstr_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ChemsRULE_instr_func)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(526)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(527)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(528)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(529)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(530)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(531)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(532)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*Instr_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(536)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(537)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(538)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(539)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(540)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(541)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(542)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(545)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(546)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(547)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(548)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(549)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(550)
			p.Match(ChemsTK_MENOSMAYOR)
		}
		{
			p.SetState(551)

			var _x = p.Instr_tipo()

			localctx.(*Instr_funcContext)._instr_tipo = _x
		}
		{
			p.SetState(552)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(553)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(554)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), localctx.(*Instr_funcContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	}

	return localctx
}

// IList_function_parametersContext is an interface to support dynamic dispatch.
type IList_function_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_parameters_fn returns the _block_parameters_fn rule contexts.
	Get_block_parameters_fn() IBlock_parameters_fnContext

	// Set_block_parameters_fn sets the _block_parameters_fn rule contexts.
	Set_block_parameters_fn(IBlock_parameters_fnContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_parameters_fnContext

	// SetE sets the e rule context list.
	SetE([]IBlock_parameters_fnContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_function_parametersContext differentiates from other interfaces.
	IsList_function_parametersContext()
}

type List_function_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	l                    *arrayList.List
	_block_parameters_fn IBlock_parameters_fnContext
	e                    []IBlock_parameters_fnContext
}

func NewEmptyList_function_parametersContext() *List_function_parametersContext {
	var p = new(List_function_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_function_parameters
	return p
}

func (*List_function_parametersContext) IsList_function_parametersContext() {}

func NewList_function_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_function_parametersContext {
	var p = new(List_function_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_function_parameters

	return p
}

func (s *List_function_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *List_function_parametersContext) Get_block_parameters_fn() IBlock_parameters_fnContext {
	return s._block_parameters_fn
}

func (s *List_function_parametersContext) Set_block_parameters_fn(v IBlock_parameters_fnContext) {
	s._block_parameters_fn = v
}

func (s *List_function_parametersContext) GetE() []IBlock_parameters_fnContext { return s.e }

func (s *List_function_parametersContext) SetE(v []IBlock_parameters_fnContext) { s.e = v }

func (s *List_function_parametersContext) GetL() *arrayList.List { return s.l }

func (s *List_function_parametersContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_function_parametersContext) AllBlock_parameters_fn() []IBlock_parameters_fnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem())
	var tst = make([]IBlock_parameters_fnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_parameters_fnContext)
		}
	}

	return tst
}

func (s *List_function_parametersContext) Block_parameters_fn(i int) IBlock_parameters_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_parameters_fnContext)
}

func (s *List_function_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_function_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_function_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_function_parameters(s)
	}
}

func (s *List_function_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_function_parameters(s)
	}
}

func (p *Chems) List_function_parameters() (localctx IList_function_parametersContext) {
	localctx = NewList_function_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ChemsRULE_list_function_parameters)

	localctx.(*List_function_parametersContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChemsID {
		{
			p.SetState(559)

			var _x = p.Block_parameters_fn()

			localctx.(*List_function_parametersContext)._block_parameters_fn = _x
		}
		localctx.(*List_function_parametersContext).e = append(localctx.(*List_function_parametersContext).e, localctx.(*List_function_parametersContext)._block_parameters_fn)

		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_function_parametersContext).GetE()
	for _, e := range listInt {
		localctx.(*List_function_parametersContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_parameters_fnContext is an interface to support dynamic dispatch.
type IBlock_parameters_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsBlock_parameters_fnContext differentiates from other interfaces.
	IsBlock_parameters_fnContext()
}

type Block_parameters_fnContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_ID         antlr.Token
	_instr_tipo IInstr_tipoContext
}

func NewEmptyBlock_parameters_fnContext() *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_parameters_fn
	return p
}

func (*Block_parameters_fnContext) IsBlock_parameters_fnContext() {}

func NewBlock_parameters_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_parameters_fn

	return p
}

func (s *Block_parameters_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_parameters_fnContext) Get_ID() antlr.Token { return s._ID }

func (s *Block_parameters_fnContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Block_parameters_fnContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Block_parameters_fnContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Block_parameters_fnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Block_parameters_fnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Block_parameters_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Block_parameters_fnContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Block_parameters_fnContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Block_parameters_fnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_parameters_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_parameters_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_parameters_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_parameters_fn(s)
	}
}

func (s *Block_parameters_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_parameters_fn(s)
	}
}

func (p *Chems) Block_parameters_fn() (localctx IBlock_parameters_fnContext) {
	localctx = NewBlock_parameters_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ChemsRULE_block_parameters_fn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(567)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(568)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(569)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		{
			p.SetState(570)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListExpre((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(574)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(575)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListExpre((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	}

	return localctx
}

// IInstr_llamadaContext is an interface to support dynamic dispatch.
type IInstr_llamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_llamadaContext differentiates from other interfaces.
	IsInstr_llamadaContext()
}

type Instr_llamadaContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_ID              antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_llamadaContext() *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_llamada
	return p
}

func (*Instr_llamadaContext) IsInstr_llamadaContext() {}

func NewInstr_llamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_llamada

	return p
}

func (s *Instr_llamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_llamadaContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_llamadaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_llamadaContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_llamadaContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_llamadaContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_llamadaContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_llamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_llamadaContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_llamadaContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_llamadaContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_llamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_llamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_llamada(s)
	}
}

func (s *Instr_llamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_llamada(s)
	}
}

func (p *Chems) Instr_llamada() (localctx IInstr_llamadaContext) {
	localctx = NewInstr_llamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ChemsRULE_instr_llamada)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(580)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_llamadaContext)._ID = _m
		}
		{
			p.SetState(581)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(582)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Instr_llamadaContext).instr = function.NewLlamada((func() string {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetText()
			}
		}()), nil, (func() int {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetLine()
			}
		}()), localctx.(*Instr_llamadaContext).Get_ID().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(584)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_llamadaContext)._ID = _m
		}
		{
			p.SetState(585)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(586)

			var _x = p.List_expression()

			localctx.(*Instr_llamadaContext)._list_expression = _x
		}
		{
			p.SetState(587)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Instr_llamadaContext).instr = function.NewLlamada((func() string {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_llamadaContext).Get_list_expression().GetL(), (func() int {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetLine()
			}
		}()), localctx.(*Instr_llamadaContext).Get_ID().GetColumn())

	}

	return localctx
}

// IInstr_tipoContext is an interface to support dynamic dispatch.
type IInstr_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TypeExpression

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TypeExpression)

	// IsInstr_tipoContext differentiates from other interfaces.
	IsInstr_tipoContext()
}

type Instr_tipoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TypeExpression
}

func NewEmptyInstr_tipoContext() *Instr_tipoContext {
	var p = new(Instr_tipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_tipo
	return p
}

func (*Instr_tipoContext) IsInstr_tipoContext() {}

func NewInstr_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_tipoContext {
	var p = new(Instr_tipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_tipo

	return p
}

func (s *Instr_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_tipoContext) GetTipo_exp() interfaces.TypeExpression { return s.tipo_exp }

func (s *Instr_tipoContext) SetTipo_exp(v interfaces.TypeExpression) { s.tipo_exp = v }

func (s *Instr_tipoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Instr_tipoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Instr_tipoContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsR_STRING, 0)
}

func (s *Instr_tipoContext) R_STR() antlr.TerminalNode {
	return s.GetToken(ChemsR_STR, 0)
}

func (s *Instr_tipoContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(ChemsR_BOOL, 0)
}

func (s *Instr_tipoContext) R_CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsR_CHAR, 0)
}

func (s *Instr_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_tipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_tipo(s)
	}
}

func (s *Instr_tipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_tipo(s)
	}
}

func (p *Chems) Instr_tipo() (localctx IInstr_tipoContext) {
	localctx = NewInstr_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ChemsRULE_instr_tipo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(604)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(592)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(594)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(596)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(598)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STR

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(600)
			p.Match(ChemsR_BOOL)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.BOOLEAN

	case ChemsR_CHAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(602)
			p.Match(ChemsR_CHAR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.CHAR

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_list_expression returns the _block_list_expression rule contexts.
	Get_block_list_expression() IBlock_list_expressionContext

	// Set_block_list_expression sets the _block_list_expression rule contexts.
	Set_block_list_expression(IBlock_list_expressionContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_list_expressionContext

	// SetE sets the e rule context list.
	SetE([]IBlock_list_expressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser                 antlr.Parser
	l                      *arrayList.List
	_block_list_expression IBlock_list_expressionContext
	e                      []IBlock_list_expressionContext
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expression
	return p
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) Get_block_list_expression() IBlock_list_expressionContext {
	return s._block_list_expression
}

func (s *List_expressionContext) Set_block_list_expression(v IBlock_list_expressionContext) {
	s._block_list_expression = v
}

func (s *List_expressionContext) GetE() []IBlock_list_expressionContext { return s.e }

func (s *List_expressionContext) SetE(v []IBlock_list_expressionContext) { s.e = v }

func (s *List_expressionContext) GetL() *arrayList.List { return s.l }

func (s *List_expressionContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expressionContext) AllBlock_list_expression() []IBlock_list_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem())
	var tst = make([]IBlock_list_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_list_expressionContext)
		}
	}

	return tst
}

func (s *List_expressionContext) Block_list_expression(i int) IBlock_list_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_list_expressionContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *Chems) List_expression() (localctx IList_expressionContext) {
	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ChemsRULE_list_expression)

	localctx.(*List_expressionContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(606)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(609)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expressionContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expressionContext).l.Add(e.GetP())
	}

	return localctx
}

// IBlock_list_expressionContext is an interface to support dynamic dispatch.
type IBlock_list_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsBlock_list_expressionContext differentiates from other interfaces.
	IsBlock_list_expressionContext()
}

type Block_list_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_list_expressionContext() *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_list_expression
	return p
}

func (*Block_list_expressionContext) IsBlock_list_expressionContext() {}

func NewBlock_list_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_list_expression

	return p
}

func (s *Block_list_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_list_expressionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_list_expressionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_list_expressionContext) GetP() interfaces.Expression { return s.p }

func (s *Block_list_expressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Block_list_expressionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_list_expressionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_list_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_list_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_list_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_list_expression(s)
	}
}

func (s *Block_list_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_list_expression(s)
	}
}

func (p *Chems) Block_list_expression() (localctx IBlock_list_expressionContext) {
	localctx = NewBlock_list_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ChemsRULE_block_list_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(613)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		{
			p.SetState(614)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(617)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	}

	return localctx
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expre_log returns the _expre_log rule contexts.
	Get_expre_log() IExpre_logContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// Get_expre_casteo returns the _expre_casteo rule contexts.
	Get_expre_casteo() IExpre_casteoContext

	// Set_expre_log sets the _expre_log rule contexts.
	Set_expre_log(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// Set_expre_casteo sets the _expre_casteo rule contexts.
	Set_expre_casteo(IExpre_casteoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	p             interfaces.Expression
	_expre_log    IExpre_logContext
	_expre_rel    IExpre_relContext
	_expre_arit   IExpre_aritContext
	_expre_casteo IExpre_casteoContext
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Get_expre_log() IExpre_logContext { return s._expre_log }

func (s *ExpressionsContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *ExpressionsContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *ExpressionsContext) Get_expre_casteo() IExpre_casteoContext { return s._expre_casteo }

func (s *ExpressionsContext) Set_expre_log(v IExpre_logContext) { s._expre_log = v }

func (s *ExpressionsContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *ExpressionsContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *ExpressionsContext) Set_expre_casteo(v IExpre_casteoContext) { s._expre_casteo = v }

func (s *ExpressionsContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionsContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionsContext) Expre_log() IExpre_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_logContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *ExpressionsContext) Expre_rel() IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *ExpressionsContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *ExpressionsContext) Expre_casteo() IExpre_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_casteoContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *Chems) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ChemsRULE_expressions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(634)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(622)

			var _x = p.expre_log(0)

			localctx.(*ExpressionsContext)._expre_log = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(625)

			var _x = p.expre_rel(0)

			localctx.(*ExpressionsContext)._expre_rel = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_rel().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(628)

			var _x = p.expre_arit(0)

			localctx.(*ExpressionsContext)._expre_arit = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_arit().GetP()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(631)

			var _x = p.Expre_casteo()

			localctx.(*ExpressionsContext)._expre_casteo = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_casteo().GetP()

	}

	return localctx
}

// IExpre_logContext is an interface to support dynamic dispatch.
type IExpre_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_logContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_logContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_logContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_logContext differentiates from other interfaces.
	IsExpre_logContext()
}

type Expre_logContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	left       IExpre_logContext
	op         antlr.Token
	_expre_rel IExpre_relContext
	right      IExpre_logContext
}

func NewEmptyExpre_logContext() *Expre_logContext {
	var p = new(Expre_logContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_log
	return p
}

func (*Expre_logContext) IsExpre_logContext() {}

func NewExpre_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_logContext {
	var p = new(Expre_logContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_log

	return p
}

func (s *Expre_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_logContext) GetOp() antlr.Token { return s.op }

func (s *Expre_logContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_logContext) GetLeft() IExpre_logContext { return s.left }

func (s *Expre_logContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *Expre_logContext) GetRight() IExpre_logContext { return s.right }

func (s *Expre_logContext) SetLeft(v IExpre_logContext) { s.left = v }

func (s *Expre_logContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *Expre_logContext) SetRight(v IExpre_logContext) { s.right = v }

func (s *Expre_logContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_logContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_logContext) TK_NOT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_NOT, 0)
}

func (s *Expre_logContext) AllExpre_log() []IExpre_logContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_logContext)(nil)).Elem())
	var tst = make([]IExpre_logContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_logContext)
		}
	}

	return tst
}

func (s *Expre_logContext) Expre_log(i int) IExpre_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_logContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *Expre_logContext) Expre_rel() IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_logContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(ChemsTK_AND, 0)
}

func (s *Expre_logContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_OR, 0)
}

func (s *Expre_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_log(s)
	}
}

func (s *Expre_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_log(s)
	}
}

func (p *Chems) Expre_log() (localctx IExpre_logContext) {
	return p.expre_log(0)
}

func (p *Chems) expre_log(_p int) (localctx IExpre_logContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_logContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_logContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, ChemsRULE_expre_log, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(644)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_NOT:
		{
			p.SetState(637)

			var _m = p.Match(ChemsTK_NOT)

			localctx.(*Expre_logContext).op = _m
		}
		{
			p.SetState(638)

			var _x = p.expre_log(3)

			localctx.(*Expre_logContext).left = _x
		}
		localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_logContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_logContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_logContext).GetOp().GetColumn())

	case ChemsR_IF, ChemsR_MATCH, ChemsR_LOOP, ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID, ChemsTK_MENOS, ChemsTK_PARA:
		{
			p.SetState(641)

			var _x = p.expre_rel(0)

			localctx.(*Expre_logContext)._expre_rel = _x
		}
		localctx.(*Expre_logContext).p = localctx.(*Expre_logContext).Get_expre_rel().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(653)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_logContext(p, _parentctx, _parentState)
			localctx.(*Expre_logContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_log)
			p.SetState(646)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(647)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_logContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ChemsTK_AND || _la == ChemsTK_OR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_logContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(648)

				var _x = p.expre_log(3)

				localctx.(*Expre_logContext).right = _x
			}
			localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_logContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_logContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_logContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_logContext).GetOp().GetColumn())

		}
		p.SetState(655)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_relContext is an interface to support dynamic dispatch.
type IExpre_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_relContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_relContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_relContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_relContext differentiates from other interfaces.
	IsExpre_relContext()
}

type Expre_relContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	left        IExpre_relContext
	_expre_arit IExpre_aritContext
	op          antlr.Token
	right       IExpre_relContext
}

func NewEmptyExpre_relContext() *Expre_relContext {
	var p = new(Expre_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_rel
	return p
}

func (*Expre_relContext) IsExpre_relContext() {}

func NewExpre_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_relContext {
	var p = new(Expre_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_rel

	return p
}

func (s *Expre_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_relContext) GetOp() antlr.Token { return s.op }

func (s *Expre_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_relContext) GetLeft() IExpre_relContext { return s.left }

func (s *Expre_relContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *Expre_relContext) GetRight() IExpre_relContext { return s.right }

func (s *Expre_relContext) SetLeft(v IExpre_relContext) { s.left = v }

func (s *Expre_relContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *Expre_relContext) SetRight(v IExpre_relContext) { s.right = v }

func (s *Expre_relContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_relContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_relContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_relContext) AllExpre_rel() []IExpre_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_relContext)(nil)).Elem())
	var tst = make([]IExpre_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_relContext)
		}
	}

	return tst
}

func (s *Expre_relContext) Expre_rel(i int) IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_relContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOR, 0)
}

func (s *Expre_relContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYOR, 0)
}

func (s *Expre_relContext) TK_DIFIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIFIGUAL, 0)
}

func (s *Expre_relContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALIGUAL, 0)
}

func (s *Expre_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_rel(s)
	}
}

func (s *Expre_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_rel(s)
	}
}

func (p *Chems) Expre_rel() (localctx IExpre_relContext) {
	return p.expre_rel(0)
}

func (p *Chems) expre_rel(_p int) (localctx IExpre_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 84
	p.EnterRecursionRule(localctx, 84, ChemsRULE_expre_rel, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(657)

		var _x = p.expre_arit(0)

		localctx.(*Expre_relContext)._expre_arit = _x
	}
	localctx.(*Expre_relContext).p = localctx.(*Expre_relContext).Get_expre_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(667)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_relContext(p, _parentctx, _parentState)
			localctx.(*Expre_relContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_rel)
			p.SetState(660)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(661)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(662)

				var _x = p.expre_rel(3)

				localctx.(*Expre_relContext).right = _x
			}
			localctx.(*Expre_relContext).p = expression.NewOperacion(localctx.(*Expre_relContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_relContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_relContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_relContext).GetOp().GetColumn())

		}
		p.SetState(669)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_aritContext is an interface to support dynamic dispatch.
type IExpre_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_aritContext

	// Get_expre_valor returns the _expre_valor rule contexts.
	Get_expre_valor() IExpre_valorContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_aritContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_aritContext)

	// Set_expre_valor sets the _expre_valor rule contexts.
	Set_expre_valor(IExpre_valorContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_aritContext differentiates from other interfaces.
	IsExpre_aritContext()
}

type Expre_aritContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	left         IExpre_aritContext
	op           antlr.Token
	_expre_valor IExpre_valorContext
	_expressions IExpressionsContext
	right        IExpre_aritContext
}

func NewEmptyExpre_aritContext() *Expre_aritContext {
	var p = new(Expre_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_arit
	return p
}

func (*Expre_aritContext) IsExpre_aritContext() {}

func NewExpre_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_aritContext {
	var p = new(Expre_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_arit

	return p
}

func (s *Expre_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expre_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_aritContext) GetLeft() IExpre_aritContext { return s.left }

func (s *Expre_aritContext) Get_expre_valor() IExpre_valorContext { return s._expre_valor }

func (s *Expre_aritContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Expre_aritContext) GetRight() IExpre_aritContext { return s.right }

func (s *Expre_aritContext) SetLeft(v IExpre_aritContext) { s.left = v }

func (s *Expre_aritContext) Set_expre_valor(v IExpre_valorContext) { s._expre_valor = v }

func (s *Expre_aritContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Expre_aritContext) SetRight(v IExpre_aritContext) { s.right = v }

func (s *Expre_aritContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_aritContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_aritContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOS, 0)
}

func (s *Expre_aritContext) AllExpre_arit() []IExpre_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem())
	var tst = make([]IExpre_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_aritContext)
		}
	}

	return tst
}

func (s *Expre_aritContext) Expre_arit(i int) IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_aritContext) Expre_valor() IExpre_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_valorContext)
}

func (s *Expre_aritContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Expre_aritContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Expre_aritContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Expre_aritContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MULT, 0)
}

func (s *Expre_aritContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIV, 0)
}

func (s *Expre_aritContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MODULO, 0)
}

func (s *Expre_aritContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAS, 0)
}

func (s *Expre_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_arit(s)
	}
}

func (s *Expre_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_arit(s)
	}
}

func (p *Chems) Expre_arit() (localctx IExpre_aritContext) {
	return p.expre_arit(0)
}

func (p *Chems) expre_arit(_p int) (localctx IExpre_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 86
	p.EnterRecursionRule(localctx, 86, ChemsRULE_expre_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(683)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_MENOS:
		{
			p.SetState(671)

			var _m = p.Match(ChemsTK_MENOS)

			localctx.(*Expre_aritContext).op = _m
		}
		{
			p.SetState(672)

			var _x = p.expre_arit(5)

			localctx.(*Expre_aritContext).left = _x
		}
		localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

	case ChemsR_IF, ChemsR_MATCH, ChemsR_LOOP, ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID:
		{
			p.SetState(675)

			var _x = p.Expre_valor()

			localctx.(*Expre_aritContext)._expre_valor = _x
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expre_valor().GetP()

	case ChemsTK_PARA:
		{
			p.SetState(678)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(679)

			var _x = p.Expressions()

			localctx.(*Expre_aritContext)._expressions = _x
		}
		{
			p.SetState(680)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expressions().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(697)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(695)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_arit)
				p.SetState(685)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(686)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(ChemsTK_MULT-44))|(1<<(ChemsTK_DIV-44))|(1<<(ChemsTK_MODULO-44)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(687)

					var _x = p.expre_arit(5)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_arit)
				p.SetState(690)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(691)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_MAS || _la == ChemsTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(692)

					var _x = p.expre_arit(4)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			}

		}
		p.SetState(699)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_valorContext is an interface to support dynamic dispatch.
type IExpre_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_valorContext differentiates from other interfaces.
	IsExpre_valorContext()
}

type Expre_valorContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	_primitivo IPrimitivoContext
}

func NewEmptyExpre_valorContext() *Expre_valorContext {
	var p = new(Expre_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_valor
	return p
}

func (*Expre_valorContext) IsExpre_valorContext() {}

func NewExpre_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_valorContext {
	var p = new(Expre_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_valor

	return p
}

func (s *Expre_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_valorContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expre_valorContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expre_valorContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_valorContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_valorContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expre_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_valor(s)
	}
}

func (s *Expre_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_valor(s)
	}
}

func (p *Chems) Expre_valor() (localctx IExpre_valorContext) {
	localctx = NewExpre_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ChemsRULE_expre_valor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(700)

		var _x = p.Primitivo()

		localctx.(*Expre_valorContext)._primitivo = _x
	}
	localctx.(*Expre_valorContext).p = localctx.(*Expre_valorContext).Get_primitivo().GetP()

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_primitivo_casteo returns the _primitivo_casteo rule contexts.
	Get_primitivo_casteo() IPrimitivo_casteoContext

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Get_instr_match_ter returns the _instr_match_ter rule contexts.
	Get_instr_match_ter() IInstr_match_terContext

	// Get_instr_loop_ternario returns the _instr_loop_ternario rule contexts.
	Get_instr_loop_ternario() IInstr_loop_ternarioContext

	// Set_primitivo_casteo sets the _primitivo_casteo rule contexts.
	Set_primitivo_casteo(IPrimitivo_casteoContext)

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// Set_instr_match_ter sets the _instr_match_ter rule contexts.
	Set_instr_match_ter(IInstr_match_terContext)

	// Set_instr_loop_ternario sets the _instr_loop_ternario rule contexts.
	Set_instr_loop_ternario(IInstr_loop_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	p                    interfaces.Expression
	_NUMBER              antlr.Token
	_DOUBLE              antlr.Token
	_STRING              antlr.Token
	_BOOLEAN             antlr.Token
	_CHAR                antlr.Token
	_ID                  antlr.Token
	_primitivo_casteo    IPrimitivo_casteoContext
	_instr_ternario      IInstr_ternarioContext
	_instr_match_ter     IInstr_match_terContext
	_instr_loop_ternario IInstr_loop_ternarioContext
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitivoContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitivoContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) Get_primitivo_casteo() IPrimitivo_casteoContext { return s._primitivo_casteo }

func (s *PrimitivoContext) Get_instr_ternario() IInstr_ternarioContext { return s._instr_ternario }

func (s *PrimitivoContext) Get_instr_match_ter() IInstr_match_terContext { return s._instr_match_ter }

func (s *PrimitivoContext) Get_instr_loop_ternario() IInstr_loop_ternarioContext {
	return s._instr_loop_ternario
}

func (s *PrimitivoContext) Set_primitivo_casteo(v IPrimitivo_casteoContext) { s._primitivo_casteo = v }

func (s *PrimitivoContext) Set_instr_ternario(v IInstr_ternarioContext) { s._instr_ternario = v }

func (s *PrimitivoContext) Set_instr_match_ter(v IInstr_match_terContext) { s._instr_match_ter = v }

func (s *PrimitivoContext) Set_instr_loop_ternario(v IInstr_loop_ternarioContext) {
	s._instr_loop_ternario = v
}

func (s *PrimitivoContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *PrimitivoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsCHAR, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *PrimitivoContext) Primitivo_casteo() IPrimitivo_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivo_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivo_casteoContext)
}

func (s *PrimitivoContext) Instr_ternario() IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *PrimitivoContext) Instr_match_ter() IInstr_match_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_match_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_match_terContext)
}

func (s *PrimitivoContext) Instr_loop_ternario() IInstr_loop_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loop_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loop_ternarioContext)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(727)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(703)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(705)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(707)

			var _m = p.Match(ChemsSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_STRING().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(709)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
		}

		// str:= (func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(711)

			var _m = p.Match(ChemsCHAR)

			localctx.(*PrimitivoContext)._CHAR = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetText()
			}
		}())[1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(713)

			var _m = p.Match(ChemsID)

			localctx.(*PrimitivoContext)._ID = _m
		}
		localctx.(*PrimitivoContext).p = variable.NewIdentifier((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_ID().GetColumn())

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(715)

			var _x = p.Primitivo_casteo()

			localctx.(*PrimitivoContext)._primitivo_casteo = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_primitivo_casteo().GetP()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(718)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(721)

			var _x = p.Instr_match_ter()

			localctx.(*PrimitivoContext)._instr_match_ter = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_match_ter().GetP()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(724)

			var _x = p.Instr_loop_ternario()

			localctx.(*PrimitivoContext)._instr_loop_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_loop_ternario().GetP()

	}

	return localctx
}

// IPrimitivo_casteoContext is an interface to support dynamic dispatch.
type IPrimitivo_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivo_casteoContext differentiates from other interfaces.
	IsPrimitivo_casteoContext()
}

type Primitivo_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        interfaces.Expression
	_NUMBER  antlr.Token
	_DOUBLE  antlr.Token
	_BOOLEAN antlr.Token
}

func NewEmptyPrimitivo_casteoContext() *Primitivo_casteoContext {
	var p = new(Primitivo_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo_casteo
	return p
}

func (*Primitivo_casteoContext) IsPrimitivo_casteoContext() {}

func NewPrimitivo_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitivo_casteoContext {
	var p = new(Primitivo_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo_casteo

	return p
}

func (s *Primitivo_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitivo_casteoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *Primitivo_casteoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *Primitivo_casteoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *Primitivo_casteoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *Primitivo_casteoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *Primitivo_casteoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *Primitivo_casteoContext) GetP() interfaces.Expression { return s.p }

func (s *Primitivo_casteoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Primitivo_casteoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *Primitivo_casteoContext) R_AS() antlr.TerminalNode {
	return s.GetToken(ChemsR_AS, 0)
}

func (s *Primitivo_casteoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Primitivo_casteoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Primitivo_casteoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *Primitivo_casteoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *Primitivo_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitivo_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo_casteo(s)
	}
}

func (s *Primitivo_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo_casteo(s)
	}
}

func (p *Chems) Primitivo_casteo() (localctx IPrimitivo_casteoContext) {
	localctx = NewPrimitivo_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ChemsRULE_primitivo_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(753)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(729)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(730)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(731)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(733)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(734)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(735)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(737)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(738)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(739)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(741)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(742)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(743)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(745)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(746)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(747)
			p.Match(ChemsR_INT)
		}

		// str:= (func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(749)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(750)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(751)
			p.Match(ChemsR_FLOAT)
		}

		// str:= (func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())

	}

	return localctx
}

// IExpre_casteoContext is an interface to support dynamic dispatch.
type IExpre_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_AS returns the _R_AS token.
	Get_R_AS() antlr.Token

	// Set_R_AS sets the _R_AS token.
	Set_R_AS(antlr.Token)

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// Get_type_casteo returns the _type_casteo rule contexts.
	Get_type_casteo() IType_casteoContext

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// Set_type_casteo sets the _type_casteo rule contexts.
	Set_type_casteo(IType_casteoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_casteoContext differentiates from other interfaces.
	IsExpre_casteoContext()
}

type Expre_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expre_arit  IExpre_aritContext
	_R_AS        antlr.Token
	_type_casteo IType_casteoContext
}

func NewEmptyExpre_casteoContext() *Expre_casteoContext {
	var p = new(Expre_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_casteo
	return p
}

func (*Expre_casteoContext) IsExpre_casteoContext() {}

func NewExpre_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_casteoContext {
	var p = new(Expre_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_casteo

	return p
}

func (s *Expre_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_casteoContext) Get_R_AS() antlr.Token { return s._R_AS }

func (s *Expre_casteoContext) Set_R_AS(v antlr.Token) { s._R_AS = v }

func (s *Expre_casteoContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *Expre_casteoContext) Get_type_casteo() IType_casteoContext { return s._type_casteo }

func (s *Expre_casteoContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *Expre_casteoContext) Set_type_casteo(v IType_casteoContext) { s._type_casteo = v }

func (s *Expre_casteoContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_casteoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_casteoContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_casteoContext) R_AS() antlr.TerminalNode {
	return s.GetToken(ChemsR_AS, 0)
}

func (s *Expre_casteoContext) Type_casteo() IType_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_casteoContext)
}

func (s *Expre_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_casteo(s)
	}
}

func (s *Expre_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_casteo(s)
	}
}

func (p *Chems) Expre_casteo() (localctx IExpre_casteoContext) {
	localctx = NewExpre_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ChemsRULE_expre_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(755)

		var _x = p.expre_arit(0)

		localctx.(*Expre_casteoContext)._expre_arit = _x
	}
	{
		p.SetState(756)

		var _m = p.Match(ChemsR_AS)

		localctx.(*Expre_casteoContext)._R_AS = _m
	}
	{
		p.SetState(757)

		var _x = p.Type_casteo()

		localctx.(*Expre_casteoContext)._type_casteo = _x
	}
	localctx.(*Expre_casteoContext).p = casteo.NewCasteo(localctx.(*Expre_casteoContext).Get_expre_arit().GetP(), localctx.(*Expre_casteoContext).Get_type_casteo().GetTipo_exp(), (func() int {
		if localctx.(*Expre_casteoContext).Get_R_AS() == nil {
			return 0
		} else {
			return localctx.(*Expre_casteoContext).Get_R_AS().GetLine()
		}
	}()), localctx.(*Expre_casteoContext).Get_R_AS().GetColumn())

	return localctx
}

// IType_casteoContext is an interface to support dynamic dispatch.
type IType_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TypeExpression

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TypeExpression)

	// IsType_casteoContext differentiates from other interfaces.
	IsType_casteoContext()
}

type Type_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TypeExpression
}

func NewEmptyType_casteoContext() *Type_casteoContext {
	var p = new(Type_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_type_casteo
	return p
}

func (*Type_casteoContext) IsType_casteoContext() {}

func NewType_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_casteoContext {
	var p = new(Type_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_type_casteo

	return p
}

func (s *Type_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_casteoContext) GetTipo_exp() interfaces.TypeExpression { return s.tipo_exp }

func (s *Type_casteoContext) SetTipo_exp(v interfaces.TypeExpression) { s.tipo_exp = v }

func (s *Type_casteoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Type_casteoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Type_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterType_casteo(s)
	}
}

func (s *Type_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitType_casteo(s)
	}
}

func (p *Chems) Type_casteo() (localctx IType_casteoContext) {
	localctx = NewType_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ChemsRULE_type_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(764)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(760)
			p.Match(ChemsR_INT)
		}
		localctx.(*Type_casteoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(762)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Type_casteoContext).tipo_exp = interfaces.FLOAT

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 41:
		var t *Expre_logContext = nil
		if localctx != nil {
			t = localctx.(*Expre_logContext)
		}
		return p.Expre_log_Sempred(t, predIndex)

	case 42:
		var t *Expre_relContext = nil
		if localctx != nil {
			t = localctx.(*Expre_relContext)
		}
		return p.Expre_rel_Sempred(t, predIndex)

	case 43:
		var t *Expre_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expre_aritContext)
		}
		return p.Expre_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Expre_log_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Expre_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Expre_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
