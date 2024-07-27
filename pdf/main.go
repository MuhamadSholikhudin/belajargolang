package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	htmlHeader := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
		<style>
			body{
				margin: 0%;
				background-origin: padding-box;
			}
			#experience{
				background-image: url("data:image/jpeg;base64,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");
				background-repeat: no-repeat;
				background-size: 300px 153px;
				background-position: center;
			}
			.font-sans-serif{
				/* font-family: gill sans, sans-serif; */
				font-family: sans-serif;
				line-height: 1.8;
			}
			.font-size-content{
				font-size: 14px;
	
			}
		</style>
	</head>
	<body><html><body>`
	htmlFooter := `</body>
	</html>`

	// var h1 string = ""

	// for i := 0; i < 1000; i++ {
	// 	h1 = fmt.Sprintf("%s <h1>Hello, PDF! %d</h1>", h1, i)
	// }

	// htmlContain := h1
	htmlContain := `    <div class="experience" id="experience" style="margin:0%; width:1000px; height:1000px;">
	<table style="margin:0%;">
		<tr>                
			<td style="width:690px; margin:0%; ">
				<div style="margin:0%; padding-left: 50px; font-size:20px; font-weight:800px;">
					<span>
						<img src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAoHCBESFRURERUZGBIYFBISEhkZHBkSGhoYGhgaHBkcGRgcIS4lHB4rIRoZJjgmKy8xNTU1GiU7QDszPy40NTEBDAwMEA8QHxISHzEnJCYxMTQ0MTQxMTQxNDQ0MT8/NDExNjUxNDc0NDQ0PzQxPzExNDQ/ND8xNzQ0PzQ0PzQ0P//AABEIAKkBKgMBIgACEQEDEQH/xAAcAAEAAgIDAQAAAAAAAAAAAAAAAQcFBgMECAL/xABJEAABAwICBAcLCQcEAwEAAAABAAIDBBEFEgYHITE0QVFxgZGxExQVIjJSYXJzobIWNVNigoOSs8EXQlR0otHSIzOTwuHi8UT/xAAaAQEAAgMBAAAAAAAAAAAAAAAABAUBAgMG/8QALhEBAAIBAgMGBgEFAAAAAAAAAAECEQMEBTFREiEyQXHBExQVJIGRoRYiM1Jh/9oADAMBAAIRAxEAPwC5kREBERAUKVCCVClEEIiICIiAoQrWtKtJmUTQAM0zhdrd1hyn0LFrRWMy6aWlfWvFKRmZbKiqB2nmIE3DmAcgbftKN08xAEEuYRyFtr9RUf5mn/Vt9B3WPL9rfUrWNFdJ2VrS0jLM22Zu8EcrTydi2YKRW0WjMKnV0r6N5peMTCVKIsuYiIgIiICIiAoUoghSiICIiAiIgIiICIiCF8SSNaLkgD0my62J1jYIpJ3eSxjnn02G5URjON1FW8vmebE+KwE5WjkA3dKkbfbW1s9+Ihpa8VX335F57fxBcjJWu8lwPMQV5uXewvFZ6VwfC9zSDci/in0EcalW4fOO6XONbrD0MpWNwLEm1UEc7RbO25HIRsI61klXTExOJd4nKVClFgQiKUEIpRB8qnNP3l1a8E7A1oHNZXGVTennDZeZnYo268H5XnAI+5n0lrqIirntWxaAvLa6ID94SNPNlJ7WhXIFTGgnDoPvPgcrnCsdr4Py8Vx+Puo9I930iIpKjEREBERAREQEREBERAREQEREBERAREQa9p1wCp9n/wBgqMV56dcAqfUHxNVGK24d4J9fZH1eYiIrFxXTq44DH6z+1bUtU1b8Bj9Z/atrXnNb/Jb1lMr4YSoRYrHcU73YC3a91w0fqubZlUVcS4tUPNzI7oNh1L48Iz/SP61rlt2VlIq18Iz/AEj+tPCM/wBI/rTJ2VkEKm9OuGS8zexZvwjP9I/rVbaW1chqXkvcT4vH6Fz1afEr2YT+Hbquz1Z1LRnMYdxFrPfL/PPWnfL/ADz1qP8AKT1Xv9Raf+k/tYug3DoPvPgKucBebNEquQVURD3A/wCpx/UcrLGIz/SP61I0qfDr2ZUXEd1Xd6vxKxjuwspFq2j2Oue4Qym5Pku478hW0rrEq+YwlERZYEREBERAREQEREBERAREQEREBERBr2nXAKn1B8TVRSvTTrgFT7MfE1UDXTFjbDeditNneKaNrT5T7OGpGbREOy57RvIUd1b5w61hCVC1niM+UM/Bjq9Aau6yJtDGHPaDmfvI5VtPhCH6Rn4gqa0R4M31n9qzVlXXv2rTbrKRWmIWX4Qh+kZ+ILT9NMQhzx/6jPJd+8OULCLTdOoznifbxcrm39OxaZyzjHe2nwjB9Iz8QTwjB9Iz8QVVKbJg7S2op2P8h7XcxBXIqnpKp8Tg9jiHA9HSrNwyqE0TJBszNBI9PGsTGGYnLtKudK+Ev5m9isZVzpXwl/M3sWasW5MOiItmrMaJ8Kj+8+AqxlXOifCo/t/A5WOtJ5tq8nZws/60XtY/iCstVthEZdPEANudh6AQT7grJCzDFkoiLZqIiICIiAihSgIouiCUUIglFF0uglERAREQa7p1wCp9QfE1efMU/d6f0XoPTvgFT7MfE1efsSYSA7kvfpU3SiZ21sdXK0/3wxyFEUJ2WJojwZvrP7Vm1hNEeDN9Z/as2tJdI5IXVxCgZOwxyDZvB4weULtKUZaTNoZJfxJGlvFmBB6bL4+Rs/ns/q/st5RMy1xDRRobP57P6v7KxNF9EphTMGdn73Ly8y6oF9g38SsTBacxwsY7yrXPOdqzHexPdya38k5vPZ7/AOyq3THBJG1cgLm3Aby8nMvQqprTvhkvMzsXLWtNK5hZcJ29NzrzTU5Yy0PwO/zmp4Hf5zVmkUT5nUej+h7TpP7TodgkjqyJoc25z8vmOVqfJObz2e/+y0rQTh0H3nwOVzKXoWm9cy85xbb022vFNPujET/LE4PgkdP418zyLZrWsPQOLnWYRF3VSUREBERAWka1a2WCi7pDI5j+6xDM05TYuAIut3Wg64+AffRfEEGN1PYnUTmp7vK+TLky53F1uZWgqN1V6R0lD3x31Jkz5Mniufe3qgqxP2j4R/Ef0Sf4oKqxvSCtbXyRtqZBGKlrQ0PNsuZuy3It21tVFVDFTzwSvYL5JMji0G4uCelVjidUyWvfLGc0b6ljmG1rjO3iKurWLQd3w2QAXcxrJW/Zsgx+qPGJKmmkbM9z5I5SMzjmJa4Bw29JHQuDW9jklPFDFA9zJHuL3FpynK3dtHpWs6lq7JVSwk7JIg8c7Hf+/uXQ1rVxqK8xM29zayJg+ud/vIQToHpTV9/07J53vje4xODnEi7mnKbesB1q7cZqxDBNMdmSOR45w0299l570pwp2G1TGsv4ghnjP1hZ3aFa+sXFm+Cs7DsnbE1npDgHdgQaDoFitdV18Eb6iQxhz5ZGlxtlYCQCOTMWhc2srHauHEJo4p5GRhsVmtcWgXjaTYLIakaDNLU1JGxjGQs53HM73NZ1rXta3znUerD+WxBd+jMjn0lM9xLnup4XOcdpJLG3JPKsssPonwKk/loPy2rMINd064BU+zHxNVFkcRV6adcAqfUHxNVGq24f4J9fZH1ebpvoGHdcc3/xR4PZyn3f2XdRSp22lPfiHPt26rP0B0chko2Pc54Jc/cW23+qtk+SlP50nW3/ABXU1b8Aj9Z/atqVFq1it5iOspdZnswwHyTp/Ok62/4p8lKfzpOtv+K2BFzw2zLX/kpT+dJ1t/xT5J0/nSdbf8VsCJgzLEUWAU8RzAFzhuLje3MBs9yy6IssIKpvTzhsvMzsVyFU7rAjLa15O5zWkdSjbnwflecAmPmp9Ja2iIq57Vn9BOHQfefA5XOFTmgEbnVsRA2NEjnc2Ut7XBXGFY7bwfl4rj8/dR6R7vpERSVGIiICIiAtB1x8A++i+ILfloOuPgH30XxBBW+guhrcT7reYxZMu5gkvfnIstv/AGNR/wAY7/ib/muPUfvqvsK20Hl6pohT1ncA7MGVEbM1st7PbtsNy9LVFOJIXRnc+Ms622XnXH/nKX+bZ8bV6Rh8lvqt7EHnDRms7wxBj3mzY5JI5PVs4dtl3NGYTX4qxztoM76h/qsNx78g6Vx6x6DuGITgCzXlsreSzx/4W0ak8OzSVFUR5LWwsPpPjO/6oOxrsw7g9UB50T/ibf3rT8ax7u2HUNLm8aMyiQehpyx/0q39ZWHd8UE4Au5gEredp2+4leeGMLiGt2ucQ1vOTYIL51SYf3LD2PIs6Z75jzXys/pa3rVaa1vnOo9WH8tivfBqMQQQwN3RxMYPstAVEa1vnOf1Ify2ILu0T4FSfy0H5bVl1h9E+BUn8tB+W1ZhBrunXAKn2Y+IKjVf+kVAammmgHlPjcG+tvb7wFQUsbmOcx4LXtNnA7wVa8OtHZtHnlH1o78vlF8r6YwuIa0XJNgBtJKsXFdGrfgEfrP7VtSweh+HGmpIon7H2LnDkLje3Ys6vO6sxN5mOspleUCKVC5thERARSoQQVr2k+jcda0XOWRt8rt/QfQthRYtWLRiXTS1b6V4vScTCpn6vqwEgFhHEbkdqM1fVhNiWAcZvf3BWzZLLh8tRafXd3jnH6a9ovo3HQtNjmkdbO7du3ADiG1bCild61isYhVaurfVvN7zmZSiIstBERAREQFjMbwWCtj7jUtLmZmvtct2g3G0LJogwuAaNUlDn72YW57ZrkuvbdvWaREGq1OgWGySunfETI5/dCczh41wb26Fs7W2AHILL7RBr2OaIUNa8S1EZc8NyAhzm7Ohd3AsDpqJhipmZGFxeRcuuTx3KyiIOKeFsjXMcLtc0tcOUEWK1an1d4XG5r2QnMx7Xtu5xGZpuNl+ULbkQRZa1i+hGH1crqieMulcGhxzObfK0NGwHkAWzIg69HSshYyJgsxjWsYN9mtFguwiIPlYTF9F6OqOaaMZ/Ob4rukjes4izFrVnMThiYiebUP2eYd5r/xLIYXonRUzs8cd3jc5xzEcyz6LedbUmMTaWIrWPIRSoXNslQpRB8OeBvIHObL6BVPa6p3smpsj3NvHLfK5zb7W77FWPog4mipiTcmFlydp3cZQZhzwN5A5zZQZG2vcW51VeuuZ7O9srnNuX3yuLeuxXHi8cj8Ap5WvfnYInlwc4G251ze5QWy1wO4396OeBvIHObKu9TNe6Smlje4udHKbFxLjlcAd59N1qet3FZHVoiY9zWxxtacrnN2u2m9igvBpB2g3HWoc8DeQOc2Ve6msQdJSyxPcXOjmNsxLjleA4bT6cy1DWpXSSYh3CN725WxRNDXOb4z+Y77uQXoCvgyN3ZhfnC6seSnhGY+JHF4xO3xWN2m/QvNWKYxPNNJP3R7S+R8gs9wDbm4sL2FtiD1EpWI0YxQVdLBUjfJG0uHI8bHjocCF09OceNBSPnYAZCWxxX2jO42BPoAuehBnZahjPLe1vOQO1THK121rgR6CD2Lz5hWjmJ4vmqMxe3MQXyPIBdvIaOLi3Ljd4RwSoY1ziw7HhocXMey9js3IPRd18d1b5w6wsJNiLanD31LNgfTPePQcpvt9BuqDwHD66teY6Z73PazObyPbs3b7oPS/dW+cOsKWvB3EHm2qhfkFjf1v+ZysHVjgdbRsnbWXu97Cy7zJsDbHfu2oN27uy9szb7rXF+pcyoKgkd4bAzG3fj9lzbceJX4g4u7svbM2+61xfqXMqCr5HeHHDMbd+MFrm3kt4lcOl2NihppKi13AZWDlcdyDMPka3ynAc5ARsjXbiDzEFedKWDFcXkc5jnvIN3EuLGNvuGzYFNRFiuDyNc5z2OO1vjF8b7cVjsKD0bdY+rxukhdklnYx3I5wBXQwPHe+6EVbBZxieXDfZ7AbjrC84VMrpHvfIS57nkuJ2naUHpf5T4f/ABMX4wjNJaBxDW1MRJIAAeLkncAqtj1QVDgHd8x7QD5DuPpXZotUlRHJHIaiMhkjJCMrtuVwdbf6EFwBSoClAREQEREEKUUIClFCCm9eH+9Tezk+Jqymj+sqggpoIZA/OyNrHWbcXHIVi9eH+9Tezk+JqyOAas6GopoZ3ukzvja91nWFzyCyDWdZOldNiPcO98/iZ82YZd+6y33BaMT4E2I/vUrvcL/otA1j6J02G9x7gXnPnzZzm3citPV80Ow2madxhAPSEFfalKvJUVELtmeJr7elhse0LE0tJ4TxWcHa1xqX9DWFrf6i1dKhrDhtfUE7MvfMR+2Dl99ltWpOjLpampdxNZGD6XEud/1QdfU1VGOrqKZx8uO9vrRvt2O9yxtIO/ccad7TVuk+zHdw6PFA6VFTP4MxiV+5jXTEerIx2X3lq7+p6k7rXTVB/cicftSPFj1Nd1oN61qYp3vQPY02fMRC31Ttf/SLfaVaYVoz3TCKmsy/6jZmPj5ckfiv+Nx+ysjrmxMyVMdK3aImXIHnv29dsqtLAcFZFQx0bxs7h3OT0l7fH97ig0nUniuaKejcdsbxNGPqP2OA5nNv9tbBrRwmSqoXiMFz43smDRvIbcOA5TZx6lV2hlQ7DcVZFIbDuslHLxXDnZWnmzBh5ld+MY5S0gaaqRsYfmDL38a1r2tzjrQUpoVp9JhzDA6MSQl5fa+VzXG2ax4xs41vTMdwXGXRtqG2lbmEbZPE8q1wHbjewXarNDMKxRgqoQWGS5bJH4l9pBJbuO0HeFWum2hMmGZJBIHxPfka62VzXWJAI5gdoQXXXUMdPQzQwtyxsglDW8gykqktXOkUOHzvlnDi10WQZRc3uDuVnaM4o+qwdz5Dd7YZo3E8eVpsT0WVY6udHYMQnfFPmyNizjKcpvcBBY/7WMO82X8K2rRzHoq+LviEODMzmeMLG7d+xat+yfDeWX8ZW06OYDDQRd7wZsmZz/GOY3dv2oKXoPnwfzj+wq/lQFL4uOC/8Y73g/3V/wB0FAV/z67+cZ8LVveud5FGxvnTNv0BaLV+NjrrcdY33NH9lu+urgkXtv0QdrU7E0UAeBtfNMXHlyvLR7gFxa5ogaJjjvbMyx5wbrs6nvm1ntaj8xy4dcnAB7eP9UHW1XOJwqYcjqoDpaT+qpZ3lfa/VXPqs+a6j1qn4VTDvK+1+qD1ZTeQz1W9i5VxU3kM9VvYvPlfpvijZJGtqXgNfIGjZsAJA4kHolFjNHZ3SUtPJIbvfBE95PG4sBJWTQEREBERAUKUQQpREFNa8P8AepvZyfE1WVobwGl9izsXPimBUlUWuqYI5S0ENL2h1gd9rru08DI2tjY0NY0ZWtGwAcQAQVXrw/8Ay/bW66u/m6l9m1ZTFMFparL3zCyTLfLnaHWvyXXZo6SOFjYomhkbRZrWiwA5AEFCa0qPuWISkbntZIOcjarI1Q0Pc6API2yySS84vlb7mhbNiWj9HVOD6mnjkeBlBe0OIHILrvUlLHCxsUTAyNgsxrRYAcgCCl9c9Dkq45hukiAJ+sw27LLP6madsVLU1TtgdIRf6kbbn3lysDE8GparKKmFkgbcsztDrX32uppsJp4ojTRxMbCQ8OY0ANObytnpuUFFYCx2J4u17xdrqh1Q/wBDIzmA5rhjelehLLFYZo9RUrjJT08cby0sLmNDTlJBIuOK4HUssgojW/hpgrm1DdjZ42vB3Wkjs13TbIekratLKR+L4VT1UIzTMa2YtG0uIaWStHpuCRy5Qt8xPB6aqDRUwskDSSzOA7KTa9r8w6ly4fh0NMwRU7GxxgkhrRYAnfYBBRuhesCTDmGnfH3WIOc5ovkcxx8oAnivxLr6Z6ZS4q6OJkWWNr8zGA53OeQQCSN5sSNnKrlxPQ7DalxfLTsLzvcLsJ5y0i65cK0UoKR2enp2MfxOtdw5nG5QYfA8GdR4S6GTZIYJZHjkc5pNujYqe0N0mdhsjpmsEhdH3OxJFtoN9i9ISxNe0scLtcCHA7QQd4KwfyKwv+Dh/AEFfftil/hmficts0D00fibpWviazI1rhYl17m3Gsr8i8L/AIOH8DV3sLwOkpS400LIy4AOyNDbgbr2QU1rMwiakrjVxgiOR7ZGPA2NeLXBPQF3263akR5TAwyWtnubX5cvKrhqqWOVpZKxr2He1wDgegrBHQbCr5u9Ir7/ACdnVuQVXq2wqasrhVvBLGPdK953F5vZoPStx11cFi9t+isClpY4mhkTGsYNga0BoHQF1sXwemq2BlTG2RgOYB24FBq2p/5tZ7Wo/McuHXJwAe3j/VblheGQ0sYhp2BkYLnBo3XJufevnFcKp6pnc6mNskdw7K7dcbig0TVZ811HrVHwKmHeV9r9V6iwvCKelYYqeNrIyS4tG65371q9XqwwuR7n5JG5iXFrXua0E8g4kGVg0ww0NaDVR3DW32nk5l51xF4dJK4G4L5CDxEZirw/ZRhfJL/yFS3VVhYIOWQgcRebH0FBsuivAqX+Wg+ALLrihiaxrWNFmtAa0DiAFgOpcqAiIgIiIIRSiAiKEBERARSoQEREBFKhBKIiAiIgIiICIoQSiIgIihAUoiAiIgIiICIiAiIg/9k=" alt="" style="width: 70px; height:40px;">
					</span>
					&nbsp;
					&nbsp;
					&nbsp;
					&nbsp;
					&nbsp;
					&nbsp;
					&nbsp;
					<span style="width: 1000px; font-size:40px;">PT. HWA SEUNG INDONESIA</span>                        
				</div>
				<div style="width:100%; text-align: center; font-weight:600px;"><b>Jalan Krasak Banyuputih RT.09 RW.03 Kecamatan Kalinyamatan Kabupaten Jepara</b></div>
				<div style="text-align: center; font-weight:600px;"><b>Provinsi Jawa Tengah, Indonesia 59467 Tel: (0291) 7512198 Fax: (0291) 7512191</b></div>
			</td>
		</tr>
	</table>

	<hr style="border: 2px solid black;">
	<div style="text-align: center; margin-top:10px; font-family: gill sans, sans-serif;"><b><u>SURAT KEPUTUSAN</u></b></div>
	<div style="text-align: center; margin-top:10px; font-size: 13px; font-family: 'Gill Sans', sans-serif;"><b>
		No. :no_letter/SK-HRD/HWI/month_letter/decided_date
		</b>
	</div>
	<div style="text-align: center; margin-top:9px; font-size: 13px; font-family: 'Gill Sans', sans-serif;"><b><u>PENGANGKATAN KARYAWAN TETAP</u></b></div>
	<div class="font-sans-serif font-size-content" style="margin-top:10px; padding-left: 50px; padding-right: 50px; ">
		<table>
			<tr>
				<td valign="top" style="padding-bottom: 5px; height: 40px;">
					Menimbang
				</td>
				<td valign="top">:</td>
				<td valign="top" style="text-align: justify;">Hasil evaluasi dan penilaian kerja karyawan yang dilakukan team manajemen PT. Hwa Seung Indonesia</td>
			</tr>
			<tr>
				<td valign="top" style="padding-bottom: 5px; height: 40px;">
					Mengingat
				</td>
				<td valign="top">:</td>
				<td valign="top"  style="text-align: justify;">
					Perjanjian Kerja Bersama PT. Hwa Seung Indonesia dan Perundang-undangan yang berlaku.
				</td>
			<tr>
				<td valign="top" style="padding-bottom: 5px; height: 20px;">
					Memperhatikan
				</td>
				<td valign="top">:</td>
				<td valign="top">
					Pemenuhan kebutuhan organisasi dan sumber daya manusia
				</td>
			</tr>
			<tr>
				<td valign="top" style="padding-bottom: 5px; height: 20px;">                        
				</td>
				<td valign="top"></td>
				<td valign="top" style="height: 30px;">
					
					<b>
						<u>
							MEMUTUSKAN ;
						</u>
					</b>
				</td>
			</tr>
			<tr>
				<td valign="top" style="height: 30;">                        
					Menetapkan
				</td>
				<td valign="top">:</td>
				<td valign="top">         
					<div>
						<table>
							<tr>
								<td valign="top" style="width:8px;">1.</td>
								<td >Pengangkatan Karyawan Tetap ;</td>
							</tr>
							<tr>
								<td></td>
								<td>
									<table>
										<tr>
											<td valign="top">Nama</td>
											<td valign="top">:</td>
											<td valign="top" style="height: 25px; width:200px;"> name</td>
											<td valign="top">Jabatan</td>
											<td valign="top">:</td>
											<td valign="top">job_level</td>
										</tr>
										<tr>
											<td valign="top" style="height: 25px;">NIK</td>
											<td valign="top">:</td>
											<td valign="top">number_of_employees</td>
											<td valign="top">Bagian</td>
											<td valign="top">:</td>
											<td valign="top">department</td>
										</tr>
									</table>
								</td>
							</tr>
							<tr>
								<td></td>
								<td style="height: 25px;">    
									Efektif Berlaku per tanggal, 
								</td>
							</tr>
							<tr>
								<td valign="top">2.</td>
								<td  style="text-align: justify; height:40px; ">
									Apabila terdapat kekeliuran, penyesuaian atau perubahan dalam Surat Keputusan ini, maka akan dilakukan perubahan seperlunya di kemudian hari.
								</td>
							</tr>
						</table>
					</div>               
					
				</td>
			</tr>
		</table>
		<br>
		<div style="height: 22px;">Diputuskan di : Jepara </div>
		<div style="height: 22px;">Pada tanggal&nbsp;&nbsp;: </div>
		<div style="height: 22px;">Human Resources Development,</div>
		<br>
		<br>
		<p>
			<b>
				<u>GUNTUR SUHENDRO</u> 
			</b>
			<br>
			<b>
				NIK. 1606001280
			</b>
		</p>
		<div>
			<span style="font-size: 8px;">Copy ; Arsip.</span>
		</div>
		<div style="margin-top:151px; font-size: 8px; align-content: flex-end; text-align: right;">             
       
			<p>building</p>
		</div>

	</div>
</div>`

	htmlFull := fmt.Sprintf("%s %s %s", htmlHeader, htmlContain, htmlFooter)

	htmlContent := fmt.Sprint(htmlFull)
	outputPath := "output.pdf"

	// Simpan konten HTML ke file sementara
	tmpFile, err := os.CreateTemp("", "html-*.html")
	if err != nil {
		fmt.Println("Error creating temporary HTML file:", err)
		return
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(htmlContent)
	if err != nil {
		fmt.Println("Error writing to temporary HTML file:", err)
		return
	}

	// Eksekusi perintah wkhtmltopdf melalui os/exec
	cmd := exec.Command("wkhtmltopdf", tmpFile.Name(), outputPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing wkhtmltopdf:", err)
		return
	}

	fmt.Println("PDF generated:", outputPath)
}
