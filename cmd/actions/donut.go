package actions

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

func MakeDonut() {
	angleA := 0.0
	angleB := 0.0
	var angleI, angleJ float64
	depthBuffer := make([]float64, 1760)
	charBuffer := make([]byte, 1760)

	// clears the screen on start
	fmt.Print("\033[2J")
	for {
		// resets the character and depth buffers
		for k := range charBuffer {
			charBuffer[k] = ' '
		}
		for k := range depthBuffer {
			depthBuffer[k] = 0
		}
		// loops through angles angleJ and angleI to generate points on the torus
		for angleJ = 0; angleJ < 6.28; angleJ += 0.07 { // outer loop for the vertical circle
			for angleI = 0; angleI < 6.28; angleI += 0.02 { // inner loop for the horizontal circle
				sinI := math.Sin(angleI)                           // sine of angle I
				cosJ := math.Cos(angleJ)                           // cosine of angle J
				sinA := math.Sin(angleA)                           // sine of rotation angle A
				sinJ := math.Sin(angleJ)                           // sine of angle J
				cosA := math.Cos(angleA)                           // cosine of rotation angle A
				offset := cosJ + 2                                 // offset to ensure positive radius
				invDepth := 1 / (sinI*offset*sinA + sinJ*cosA + 5) // depth factor
				cosI := math.Cos(angleI)                           // cosine of angle I
				cosB := math.Cos(angleB)                           // cosine of rotation angle B
				sinB := math.Sin(angleB)                           // sine of rotation angle B
				transform := sinI*offset*cosA - sinJ*sinA
				screenX := int(40 + 30*invDepth*(cosI*offset*cosB-transform*sinB))                                         // X-coordinate on the screen
				screenY := int(12 + 15*invDepth*(cosI*offset*sinB+transform*cosB))                                         // Y-coordinate on the screen
				index := screenX + 80*screenY                                                                              // 1D index in the screen buffer
				luminanceIndex := int(8 * ((sinJ*sinA-sinI*cosJ*cosA)*cosB - sinI*cosJ*sinA - sinJ*cosA - cosI*cosJ*sinB)) // intensity of the character

				// updates the buffer if the point is within bounds and closer than the previous point
				if 22 > screenY && screenY > 0 && screenX > 0 && 80 > screenX && invDepth > depthBuffer[index] {
					depthBuffer[index] = invDepth
					if luminanceIndex > 0 {
						charBuffer[index] = ".,-~:;=!*#$@"[luminanceIndex]
					} else {
						charBuffer[index] = '.'
					}
				}
			}
		}
		// move the cursor to the top-left corner
		fmt.Print("\033[H")
		// print the buffer to the screen
		for k := 0; k < 1760; k++ {
			if k%80 == 0 {
				fmt.Print("\n")
			} else {
				fmt.Printf("%c", charBuffer[k])
			}
		}
		angleA += 0.04                    // increment rotation angle A
		angleB += 0.02                    // increment rotation angle B
		time.Sleep(50 * time.Millisecond) // control frame rate
	}
}

var DonutCmd = &cobra.Command{
	Use:   "please",
	Short: "Gives to the user a awesome donut :)",
	Long:  `Show a Donut in the spectacular conditions and render the 3d edition`,
	Run: func(cmd *cobra.Command, args []string) {
		MakeDonut()
	},
}
