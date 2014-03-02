package m65go2

func ExampleNesTest() {
	Setup()

	cpu.EnableDecode()
	cpu.DisableDecimalMode()

	cpu.Registers.P = 0x24
	cpu.Registers.SP = 0xfd
	cpu.Registers.PC = 0xc000

	cpu.Memory.(*BasicMemory).load("test-roms/nestest/nestest.nes")

	cpu.Memory.Store(0x4004, 0xff)
	cpu.Memory.Store(0x4005, 0xff)
	cpu.Memory.Store(0x4006, 0xff)
	cpu.Memory.Store(0x4007, 0xff)
	cpu.Memory.Store(0x4015, 0xff)

	cpu.Run()

	Teardown()

	// Output:
	// C000  4C F5 C5  JMP $C5F5                       A:00 X:00 Y:00 P:24 SP:FD
	// C5F5  A2 00     LDX #$00                        A:00 X:00 Y:00 P:24 SP:FD
	// C5F7  86 00     STX $00 = 00                    A:00 X:00 Y:00 P:26 SP:FD
	// C5F9  86 10     STX $10 = 00                    A:00 X:00 Y:00 P:26 SP:FD
	// C5FB  86 11     STX $11 = 00                    A:00 X:00 Y:00 P:26 SP:FD
	// C5FD  20 2D C7  JSR $C72D                       A:00 X:00 Y:00 P:26 SP:FD
	// C72D  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C72E  38        SEC                             A:00 X:00 Y:00 P:26 SP:FB
	// C72F  B0 04     BCS $C735                       A:00 X:00 Y:00 P:27 SP:FB
	// C735  EA        NOP                             A:00 X:00 Y:00 P:27 SP:FB
	// C736  18        CLC                             A:00 X:00 Y:00 P:27 SP:FB
	// C737  B0 03     BCS $C73C                       A:00 X:00 Y:00 P:26 SP:FB
	// C739  4C 40 C7  JMP $C740                       A:00 X:00 Y:00 P:26 SP:FB
	// C740  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C741  38        SEC                             A:00 X:00 Y:00 P:26 SP:FB
	// C742  90 03     BCC $C747                       A:00 X:00 Y:00 P:27 SP:FB
	// C744  4C 4B C7  JMP $C74B                       A:00 X:00 Y:00 P:27 SP:FB
	// C74B  EA        NOP                             A:00 X:00 Y:00 P:27 SP:FB
	// C74C  18        CLC                             A:00 X:00 Y:00 P:27 SP:FB
	// C74D  90 04     BCC $C753                       A:00 X:00 Y:00 P:26 SP:FB
	// C753  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C754  A9 00     LDA #$00                        A:00 X:00 Y:00 P:26 SP:FB
	// C756  F0 04     BEQ $C75C                       A:00 X:00 Y:00 P:26 SP:FB
	// C75C  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C75D  A9 40     LDA #$40                        A:00 X:00 Y:00 P:26 SP:FB
	// C75F  F0 03     BEQ $C764                       A:40 X:00 Y:00 P:24 SP:FB
	// C761  4C 68 C7  JMP $C768                       A:40 X:00 Y:00 P:24 SP:FB
	// C768  EA        NOP                             A:40 X:00 Y:00 P:24 SP:FB
	// C769  A9 40     LDA #$40                        A:40 X:00 Y:00 P:24 SP:FB
	// C76B  D0 04     BNE $C771                       A:40 X:00 Y:00 P:24 SP:FB
	// C771  EA        NOP                             A:40 X:00 Y:00 P:24 SP:FB
	// C772  A9 00     LDA #$00                        A:40 X:00 Y:00 P:24 SP:FB
	// C774  D0 03     BNE $C779                       A:00 X:00 Y:00 P:26 SP:FB
	// C776  4C 7D C7  JMP $C77D                       A:00 X:00 Y:00 P:26 SP:FB
	// C77D  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C77E  A9 FF     LDA #$FF                        A:00 X:00 Y:00 P:26 SP:FB
	// C780  85 01     STA $01 = 00                    A:FF X:00 Y:00 P:A4 SP:FB
	// C782  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:A4 SP:FB
	// C784  70 04     BVS $C78A                       A:FF X:00 Y:00 P:E4 SP:FB
	// C78A  EA        NOP                             A:FF X:00 Y:00 P:E4 SP:FB
	// C78B  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:E4 SP:FB
	// C78D  50 03     BVC $C792                       A:FF X:00 Y:00 P:E4 SP:FB
	// C78F  4C 96 C7  JMP $C796                       A:FF X:00 Y:00 P:E4 SP:FB
	// C796  EA        NOP                             A:FF X:00 Y:00 P:E4 SP:FB
	// C797  A9 00     LDA #$00                        A:FF X:00 Y:00 P:E4 SP:FB
	// C799  85 01     STA $01 = FF                    A:00 X:00 Y:00 P:66 SP:FB
	// C79B  24 01     BIT $01 = 00                    A:00 X:00 Y:00 P:66 SP:FB
	// C79D  50 04     BVC $C7A3                       A:00 X:00 Y:00 P:26 SP:FB
	// C7A3  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C7A4  24 01     BIT $01 = 00                    A:00 X:00 Y:00 P:26 SP:FB
	// C7A6  70 03     BVS $C7AB                       A:00 X:00 Y:00 P:26 SP:FB
	// C7A8  4C AF C7  JMP $C7AF                       A:00 X:00 Y:00 P:26 SP:FB
	// C7AF  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C7B0  A9 00     LDA #$00                        A:00 X:00 Y:00 P:26 SP:FB
	// C7B2  10 04     BPL $C7B8                       A:00 X:00 Y:00 P:26 SP:FB
	// C7B8  EA        NOP                             A:00 X:00 Y:00 P:26 SP:FB
	// C7B9  A9 80     LDA #$80                        A:00 X:00 Y:00 P:26 SP:FB
	// C7BB  10 03     BPL $C7C0                       A:80 X:00 Y:00 P:A4 SP:FB
	// C7BD  4C D9 C7  JMP $C7D9                       A:80 X:00 Y:00 P:A4 SP:FB
	// C7D9  EA        NOP                             A:80 X:00 Y:00 P:A4 SP:FB
	// C7DA  60        RTS                             A:80 X:00 Y:00 P:A4 SP:FB
	// C600  20 DB C7  JSR $C7DB                       A:80 X:00 Y:00 P:A4 SP:FD
	// C7DB  EA        NOP                             A:80 X:00 Y:00 P:A4 SP:FB
	// C7DC  A9 FF     LDA #$FF                        A:80 X:00 Y:00 P:A4 SP:FB
	// C7DE  85 01     STA $01 = 00                    A:FF X:00 Y:00 P:A4 SP:FB
	// C7E0  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:A4 SP:FB
	// C7E2  A9 00     LDA #$00                        A:FF X:00 Y:00 P:E4 SP:FB
	// C7E4  38        SEC                             A:00 X:00 Y:00 P:66 SP:FB
	// C7E5  78        SEI                             A:00 X:00 Y:00 P:67 SP:FB
	// C7E6  F8        SED                             A:00 X:00 Y:00 P:67 SP:FB
	// C7E7  08        PHP                             A:00 X:00 Y:00 P:6F SP:FB
	// C7E8  68        PLA                             A:00 X:00 Y:00 P:6F SP:FA
	// C7E9  29 EF     AND #$EF                        A:7F X:00 Y:00 P:6D SP:FB
	// C7EB  C9 6F     CMP #$6F                        A:6F X:00 Y:00 P:6D SP:FB
	// C7ED  F0 04     BEQ $C7F3                       A:6F X:00 Y:00 P:6F SP:FB
	// C7F3  EA        NOP                             A:6F X:00 Y:00 P:6F SP:FB
	// C7F4  A9 40     LDA #$40                        A:6F X:00 Y:00 P:6F SP:FB
	// C7F6  85 01     STA $01 = FF                    A:40 X:00 Y:00 P:6D SP:FB
	// C7F8  24 01     BIT $01 = 40                    A:40 X:00 Y:00 P:6D SP:FB
	// C7FA  D8        CLD                             A:40 X:00 Y:00 P:6D SP:FB
	// C7FB  A9 10     LDA #$10                        A:40 X:00 Y:00 P:65 SP:FB
	// C7FD  18        CLC                             A:10 X:00 Y:00 P:65 SP:FB
	// C7FE  08        PHP                             A:10 X:00 Y:00 P:64 SP:FB
	// C7FF  68        PLA                             A:10 X:00 Y:00 P:64 SP:FA
	// C800  29 EF     AND #$EF                        A:74 X:00 Y:00 P:64 SP:FB
	// C802  C9 64     CMP #$64                        A:64 X:00 Y:00 P:64 SP:FB
	// C804  F0 04     BEQ $C80A                       A:64 X:00 Y:00 P:67 SP:FB
	// C80A  EA        NOP                             A:64 X:00 Y:00 P:67 SP:FB
	// C80B  A9 80     LDA #$80                        A:64 X:00 Y:00 P:67 SP:FB
	// C80D  85 01     STA $01 = 40                    A:80 X:00 Y:00 P:E5 SP:FB
	// C80F  24 01     BIT $01 = 80                    A:80 X:00 Y:00 P:E5 SP:FB
	// C811  F8        SED                             A:80 X:00 Y:00 P:A5 SP:FB
	// C812  A9 00     LDA #$00                        A:80 X:00 Y:00 P:AD SP:FB
	// C814  38        SEC                             A:00 X:00 Y:00 P:2F SP:FB
	// C815  08        PHP                             A:00 X:00 Y:00 P:2F SP:FB
	// C816  68        PLA                             A:00 X:00 Y:00 P:2F SP:FA
	// C817  29 EF     AND #$EF                        A:3F X:00 Y:00 P:2D SP:FB
	// C819  C9 2F     CMP #$2F                        A:2F X:00 Y:00 P:2D SP:FB
	// C81B  F0 04     BEQ $C821                       A:2F X:00 Y:00 P:2F SP:FB
	// C821  EA        NOP                             A:2F X:00 Y:00 P:2F SP:FB
	// C822  A9 FF     LDA #$FF                        A:2F X:00 Y:00 P:2F SP:FB
	// C824  48        PHA                             A:FF X:00 Y:00 P:AD SP:FB
	// C825  28        PLP                             A:FF X:00 Y:00 P:AD SP:FA
	// C826  D0 09     BNE $C831                       A:FF X:00 Y:00 P:EF SP:FB
	// C828  10 07     BPL $C831                       A:FF X:00 Y:00 P:EF SP:FB
	// C82A  50 05     BVC $C831                       A:FF X:00 Y:00 P:EF SP:FB
	// C82C  90 03     BCC $C831                       A:FF X:00 Y:00 P:EF SP:FB
	// C82E  4C 35 C8  JMP $C835                       A:FF X:00 Y:00 P:EF SP:FB
	// C835  EA        NOP                             A:FF X:00 Y:00 P:EF SP:FB
	// C836  A9 04     LDA #$04                        A:FF X:00 Y:00 P:EF SP:FB
	// C838  48        PHA                             A:04 X:00 Y:00 P:6D SP:FB
	// C839  28        PLP                             A:04 X:00 Y:00 P:6D SP:FA
	// C83A  F0 09     BEQ $C845                       A:04 X:00 Y:00 P:24 SP:FB
	// C83C  30 07     BMI $C845                       A:04 X:00 Y:00 P:24 SP:FB
	// C83E  70 05     BVS $C845                       A:04 X:00 Y:00 P:24 SP:FB
	// C840  B0 03     BCS $C845                       A:04 X:00 Y:00 P:24 SP:FB
	// C842  4C 49 C8  JMP $C849                       A:04 X:00 Y:00 P:24 SP:FB
	// C849  EA        NOP                             A:04 X:00 Y:00 P:24 SP:FB
	// C84A  F8        SED                             A:04 X:00 Y:00 P:24 SP:FB
	// C84B  A9 FF     LDA #$FF                        A:04 X:00 Y:00 P:2C SP:FB
	// C84D  85 01     STA $01 = 80                    A:FF X:00 Y:00 P:AC SP:FB
	// C84F  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:AC SP:FB
	// C851  18        CLC                             A:FF X:00 Y:00 P:EC SP:FB
	// C852  A9 00     LDA #$00                        A:FF X:00 Y:00 P:EC SP:FB
	// C854  48        PHA                             A:00 X:00 Y:00 P:6E SP:FB
	// C855  A9 FF     LDA #$FF                        A:00 X:00 Y:00 P:6E SP:FA
	// C857  68        PLA                             A:FF X:00 Y:00 P:EC SP:FA
	// C858  D0 09     BNE $C863                       A:00 X:00 Y:00 P:6E SP:FB
	// C85A  30 07     BMI $C863                       A:00 X:00 Y:00 P:6E SP:FB
	// C85C  50 05     BVC $C863                       A:00 X:00 Y:00 P:6E SP:FB
	// C85E  B0 03     BCS $C863                       A:00 X:00 Y:00 P:6E SP:FB
	// C860  4C 67 C8  JMP $C867                       A:00 X:00 Y:00 P:6E SP:FB
	// C867  EA        NOP                             A:00 X:00 Y:00 P:6E SP:FB
	// C868  A9 00     LDA #$00                        A:00 X:00 Y:00 P:6E SP:FB
	// C86A  85 01     STA $01 = FF                    A:00 X:00 Y:00 P:6E SP:FB
	// C86C  24 01     BIT $01 = 00                    A:00 X:00 Y:00 P:6E SP:FB
	// C86E  38        SEC                             A:00 X:00 Y:00 P:2E SP:FB
	// C86F  A9 FF     LDA #$FF                        A:00 X:00 Y:00 P:2F SP:FB
	// C871  48        PHA                             A:FF X:00 Y:00 P:AD SP:FB
	// C872  A9 00     LDA #$00                        A:FF X:00 Y:00 P:AD SP:FA
	// C874  68        PLA                             A:00 X:00 Y:00 P:2F SP:FA
	// C875  F0 09     BEQ $C880                       A:FF X:00 Y:00 P:AD SP:FB
	// C877  10 07     BPL $C880                       A:FF X:00 Y:00 P:AD SP:FB
	// C879  70 05     BVS $C880                       A:FF X:00 Y:00 P:AD SP:FB
	// C87B  90 03     BCC $C880                       A:FF X:00 Y:00 P:AD SP:FB
	// C87D  4C 84 C8  JMP $C884                       A:FF X:00 Y:00 P:AD SP:FB
	// C884  60        RTS                             A:FF X:00 Y:00 P:AD SP:FB
	// C603  20 85 C8  JSR $C885                       A:FF X:00 Y:00 P:AD SP:FD
	// C885  EA        NOP                             A:FF X:00 Y:00 P:AD SP:FB
	// C886  18        CLC                             A:FF X:00 Y:00 P:AD SP:FB
	// C887  A9 FF     LDA #$FF                        A:FF X:00 Y:00 P:AC SP:FB
	// C889  85 01     STA $01 = 00                    A:FF X:00 Y:00 P:AC SP:FB
	// C88B  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:AC SP:FB
	// C88D  A9 55     LDA #$55                        A:FF X:00 Y:00 P:EC SP:FB
	// C88F  09 AA     ORA #$AA                        A:55 X:00 Y:00 P:6C SP:FB
	// C891  B0 0B     BCS $C89E                       A:FF X:00 Y:00 P:EC SP:FB
	// C893  10 09     BPL $C89E                       A:FF X:00 Y:00 P:EC SP:FB
	// C895  C9 FF     CMP #$FF                        A:FF X:00 Y:00 P:EC SP:FB
	// C897  D0 05     BNE $C89E                       A:FF X:00 Y:00 P:6F SP:FB
	// C899  50 03     BVC $C89E                       A:FF X:00 Y:00 P:6F SP:FB
	// C89B  4C A2 C8  JMP $C8A2                       A:FF X:00 Y:00 P:6F SP:FB
	// C8A2  EA        NOP                             A:FF X:00 Y:00 P:6F SP:FB
	// C8A3  38        SEC                             A:FF X:00 Y:00 P:6F SP:FB
	// C8A4  B8        CLV                             A:FF X:00 Y:00 P:6F SP:FB
	// C8A5  A9 00     LDA #$00                        A:FF X:00 Y:00 P:2F SP:FB
	// C8A7  09 00     ORA #$00                        A:00 X:00 Y:00 P:2F SP:FB
	// C8A9  D0 09     BNE $C8B4                       A:00 X:00 Y:00 P:2F SP:FB
	// C8AB  70 07     BVS $C8B4                       A:00 X:00 Y:00 P:2F SP:FB
	// C8AD  90 05     BCC $C8B4                       A:00 X:00 Y:00 P:2F SP:FB
	// C8AF  30 03     BMI $C8B4                       A:00 X:00 Y:00 P:2F SP:FB
	// C8B1  4C B8 C8  JMP $C8B8                       A:00 X:00 Y:00 P:2F SP:FB
	// C8B8  EA        NOP                             A:00 X:00 Y:00 P:2F SP:FB
	// C8B9  18        CLC                             A:00 X:00 Y:00 P:2F SP:FB
	// C8BA  24 01     BIT $01 = FF                    A:00 X:00 Y:00 P:2E SP:FB
	// C8BC  A9 55     LDA #$55                        A:00 X:00 Y:00 P:EE SP:FB
	// C8BE  29 AA     AND #$AA                        A:55 X:00 Y:00 P:6C SP:FB
	// C8C0  D0 09     BNE $C8CB                       A:00 X:00 Y:00 P:6E SP:FB
	// C8C2  50 07     BVC $C8CB                       A:00 X:00 Y:00 P:6E SP:FB
	// C8C4  B0 05     BCS $C8CB                       A:00 X:00 Y:00 P:6E SP:FB
	// C8C6  30 03     BMI $C8CB                       A:00 X:00 Y:00 P:6E SP:FB
	// C8C8  4C CF C8  JMP $C8CF                       A:00 X:00 Y:00 P:6E SP:FB
	// C8CF  EA        NOP                             A:00 X:00 Y:00 P:6E SP:FB
	// C8D0  38        SEC                             A:00 X:00 Y:00 P:6E SP:FB
	// C8D1  B8        CLV                             A:00 X:00 Y:00 P:6F SP:FB
	// C8D2  A9 F8     LDA #$F8                        A:00 X:00 Y:00 P:2F SP:FB
	// C8D4  29 EF     AND #$EF                        A:F8 X:00 Y:00 P:AD SP:FB
	// C8D6  90 0B     BCC $C8E3                       A:E8 X:00 Y:00 P:AD SP:FB
	// C8D8  10 09     BPL $C8E3                       A:E8 X:00 Y:00 P:AD SP:FB
	// C8DA  C9 E8     CMP #$E8                        A:E8 X:00 Y:00 P:AD SP:FB
	// C8DC  D0 05     BNE $C8E3                       A:E8 X:00 Y:00 P:2F SP:FB
	// C8DE  70 03     BVS $C8E3                       A:E8 X:00 Y:00 P:2F SP:FB
	// C8E0  4C E7 C8  JMP $C8E7                       A:E8 X:00 Y:00 P:2F SP:FB
	// C8E7  EA        NOP                             A:E8 X:00 Y:00 P:2F SP:FB
	// C8E8  18        CLC                             A:E8 X:00 Y:00 P:2F SP:FB
	// C8E9  24 01     BIT $01 = FF                    A:E8 X:00 Y:00 P:2E SP:FB
	// C8EB  A9 5F     LDA #$5F                        A:E8 X:00 Y:00 P:EC SP:FB
	// C8ED  49 AA     EOR #$AA                        A:5F X:00 Y:00 P:6C SP:FB
	// C8EF  B0 0B     BCS $C8FC                       A:F5 X:00 Y:00 P:EC SP:FB
	// C8F1  10 09     BPL $C8FC                       A:F5 X:00 Y:00 P:EC SP:FB
	// C8F3  C9 F5     CMP #$F5                        A:F5 X:00 Y:00 P:EC SP:FB
	// C8F5  D0 05     BNE $C8FC                       A:F5 X:00 Y:00 P:6F SP:FB
	// C8F7  50 03     BVC $C8FC                       A:F5 X:00 Y:00 P:6F SP:FB
	// C8F9  4C 00 C9  JMP $C900                       A:F5 X:00 Y:00 P:6F SP:FB
	// C900  EA        NOP                             A:F5 X:00 Y:00 P:6F SP:FB
	// C901  38        SEC                             A:F5 X:00 Y:00 P:6F SP:FB
	// C902  B8        CLV                             A:F5 X:00 Y:00 P:6F SP:FB
	// C903  A9 70     LDA #$70                        A:F5 X:00 Y:00 P:2F SP:FB
	// C905  49 70     EOR #$70                        A:70 X:00 Y:00 P:2D SP:FB
	// C907  D0 09     BNE $C912                       A:00 X:00 Y:00 P:2F SP:FB
	// C909  70 07     BVS $C912                       A:00 X:00 Y:00 P:2F SP:FB
	// C90B  90 05     BCC $C912                       A:00 X:00 Y:00 P:2F SP:FB
	// C90D  30 03     BMI $C912                       A:00 X:00 Y:00 P:2F SP:FB
	// C90F  4C 16 C9  JMP $C916                       A:00 X:00 Y:00 P:2F SP:FB
	// C916  EA        NOP                             A:00 X:00 Y:00 P:2F SP:FB
	// C917  18        CLC                             A:00 X:00 Y:00 P:2F SP:FB
	// C918  24 01     BIT $01 = FF                    A:00 X:00 Y:00 P:2E SP:FB
	// C91A  A9 00     LDA #$00                        A:00 X:00 Y:00 P:EE SP:FB
	// C91C  69 69     ADC #$69                        A:00 X:00 Y:00 P:6E SP:FB
	// C91E  30 0B     BMI $C92B                       A:69 X:00 Y:00 P:2C SP:FB
	// C920  B0 09     BCS $C92B                       A:69 X:00 Y:00 P:2C SP:FB
	// C922  C9 69     CMP #$69                        A:69 X:00 Y:00 P:2C SP:FB
	// C924  D0 05     BNE $C92B                       A:69 X:00 Y:00 P:2F SP:FB
	// C926  70 03     BVS $C92B                       A:69 X:00 Y:00 P:2F SP:FB
	// C928  4C 2F C9  JMP $C92F                       A:69 X:00 Y:00 P:2F SP:FB
	// C92F  EA        NOP                             A:69 X:00 Y:00 P:2F SP:FB
	// C930  38        SEC                             A:69 X:00 Y:00 P:2F SP:FB
	// C931  F8        SED                             A:69 X:00 Y:00 P:2F SP:FB
	// C932  24 01     BIT $01 = FF                    A:69 X:00 Y:00 P:2F SP:FB
	// C934  A9 01     LDA #$01                        A:69 X:00 Y:00 P:ED SP:FB
	// C936  69 69     ADC #$69                        A:01 X:00 Y:00 P:6D SP:FB
	// C938  30 0B     BMI $C945                       A:6B X:00 Y:00 P:2C SP:FB
	// C93A  B0 09     BCS $C945                       A:6B X:00 Y:00 P:2C SP:FB
	// C93C  C9 6B     CMP #$6B                        A:6B X:00 Y:00 P:2C SP:FB
	// C93E  D0 05     BNE $C945                       A:6B X:00 Y:00 P:2F SP:FB
	// C940  70 03     BVS $C945                       A:6B X:00 Y:00 P:2F SP:FB
	// C942  4C 49 C9  JMP $C949                       A:6B X:00 Y:00 P:2F SP:FB
	// C949  EA        NOP                             A:6B X:00 Y:00 P:2F SP:FB
	// C94A  D8        CLD                             A:6B X:00 Y:00 P:2F SP:FB
	// C94B  38        SEC                             A:6B X:00 Y:00 P:27 SP:FB
	// C94C  B8        CLV                             A:6B X:00 Y:00 P:27 SP:FB
	// C94D  A9 7F     LDA #$7F                        A:6B X:00 Y:00 P:27 SP:FB
	// C94F  69 7F     ADC #$7F                        A:7F X:00 Y:00 P:25 SP:FB
	// C951  10 0B     BPL $C95E                       A:FF X:00 Y:00 P:E4 SP:FB
	// C953  B0 09     BCS $C95E                       A:FF X:00 Y:00 P:E4 SP:FB
	// C955  C9 FF     CMP #$FF                        A:FF X:00 Y:00 P:E4 SP:FB
	// C957  D0 05     BNE $C95E                       A:FF X:00 Y:00 P:67 SP:FB
	// C959  50 03     BVC $C95E                       A:FF X:00 Y:00 P:67 SP:FB
	// C95B  4C 62 C9  JMP $C962                       A:FF X:00 Y:00 P:67 SP:FB
	// C962  EA        NOP                             A:FF X:00 Y:00 P:67 SP:FB
	// C963  18        CLC                             A:FF X:00 Y:00 P:67 SP:FB
	// C964  24 01     BIT $01 = FF                    A:FF X:00 Y:00 P:66 SP:FB
	// C966  A9 7F     LDA #$7F                        A:FF X:00 Y:00 P:E4 SP:FB
	// C968  69 80     ADC #$80                        A:7F X:00 Y:00 P:64 SP:FB
	// C96A  10 0B     BPL $C977                       A:FF X:00 Y:00 P:A4 SP:FB
	// C96C  B0 09     BCS $C977                       A:FF X:00 Y:00 P:A4 SP:FB
	// C96E  C9 FF     CMP #$FF                        A:FF X:00 Y:00 P:A4 SP:FB
	// C970  D0 05     BNE $C977                       A:FF X:00 Y:00 P:27 SP:FB
	// C972  70 03     BVS $C977                       A:FF X:00 Y:00 P:27 SP:FB
	// C974  4C 7B C9  JMP $C97B                       A:FF X:00 Y:00 P:27 SP:FB
	// C97B  EA        NOP                             A:FF X:00 Y:00 P:27 SP:FB
	// C97C  38        SEC                             A:FF X:00 Y:00 P:27 SP:FB
	// C97D  B8        CLV                             A:FF X:00 Y:00 P:27 SP:FB
	// C97E  A9 7F     LDA #$7F                        A:FF X:00 Y:00 P:27 SP:FB
	// C980  69 80     ADC #$80                        A:7F X:00 Y:00 P:25 SP:FB
	// C982  D0 09     BNE $C98D                       A:00 X:00 Y:00 P:27 SP:FB
	// C984  30 07     BMI $C98D                       A:00 X:00 Y:00 P:27 SP:FB
	// C986  70 05     BVS $C98D                       A:00 X:00 Y:00 P:27 SP:FB
	// C988  90 03     BCC $C98D                       A:00 X:00 Y:00 P:27 SP:FB
	// C98A  4C 91 C9  JMP $C991                       A:00 X:00 Y:00 P:27 SP:FB
	// C991  EA        NOP                             A:00 X:00 Y:00 P:27 SP:FB
	// C992  38        SEC                             A:00 X:00 Y:00 P:27 SP:FB
	// C993  B8        CLV                             A:00 X:00 Y:00 P:27 SP:FB
	// C994  A9 9F     LDA #$9F                        A:00 X:00 Y:00 P:27 SP:FB
	// C996  F0 09     BEQ $C9A1                       A:9F X:00 Y:00 P:A5 SP:FB
	// C998  10 07     BPL $C9A1                       A:9F X:00 Y:00 P:A5 SP:FB
	// C99A  70 05     BVS $C9A1                       A:9F X:00 Y:00 P:A5 SP:FB
	// C99C  90 03     BCC $C9A1                       A:9F X:00 Y:00 P:A5 SP:FB
	// C99E  4C A5 C9  JMP $C9A5                       A:9F X:00 Y:00 P:A5 SP:FB
	// C9A5  EA        NOP                             A:9F X:00 Y:00 P:A5 SP:FB
	// C9A6  18        CLC                             A:9F X:00 Y:00 P:A5 SP:FB
	// C9A7  24 01     BIT $01 = FF                    A:9F X:00 Y:00 P:A4 SP:FB
	// C9A9  A9 00     LDA #$00                        A:9F X:00 Y:00 P:E4 SP:FB
	// C9AB  D0 09     BNE $C9B6                       A:00 X:00 Y:00 P:66 SP:FB
	// C9AD  30 07     BMI $C9B6                       A:00 X:00 Y:00 P:66 SP:FB
	// C9AF  50 05     BVC $C9B6                       A:00 X:00 Y:00 P:66 SP:FB
	// C9B1  B0 03     BCS $C9B6                       A:00 X:00 Y:00 P:66 SP:FB
	// C9B3  4C BA C9  JMP $C9BA                       A:00 X:00 Y:00 P:66 SP:FB
	// C9BA  EA        NOP                             A:00 X:00 Y:00 P:66 SP:FB
	// C9BB  24 01     BIT $01 = FF                    A:00 X:00 Y:00 P:66 SP:FB
	// C9BD  A9 40     LDA #$40                        A:00 X:00 Y:00 P:E6 SP:FB
	// C9BF  C9 40     CMP #$40                        A:40 X:00 Y:00 P:64 SP:FB
	// C9C1  30 09     BMI $C9CC                       A:40 X:00 Y:00 P:67 SP:FB
	// C9C3  90 07     BCC $C9CC                       A:40 X:00 Y:00 P:67 SP:FB
	// C9C5  D0 05     BNE $C9CC                       A:40 X:00 Y:00 P:67 SP:FB
	// C9C7  50 03     BVC $C9CC                       A:40 X:00 Y:00 P:67 SP:FB
	// C9C9  4C D0 C9  JMP $C9D0                       A:40 X:00 Y:00 P:67 SP:FB
	// C9D0  EA        NOP                             A:40 X:00 Y:00 P:67 SP:FB
	// C9D1  B8        CLV                             A:40 X:00 Y:00 P:67 SP:FB
	// C9D2  C9 3F     CMP #$3F                        A:40 X:00 Y:00 P:27 SP:FB
	// C9D4  F0 09     BEQ $C9DF                       A:40 X:00 Y:00 P:25 SP:FB
	// C9D6  30 07     BMI $C9DF                       A:40 X:00 Y:00 P:25 SP:FB
	// C9D8  90 05     BCC $C9DF                       A:40 X:00 Y:00 P:25 SP:FB
	// C9DA  70 03     BVS $C9DF                       A:40 X:00 Y:00 P:25 SP:FB
	// C9DC  4C E3 C9  JMP $C9E3                       A:40 X:00 Y:00 P:25 SP:FB
	// C9E3  EA        NOP                             A:40 X:00 Y:00 P:25 SP:FB
	// C9E4  C9 41     CMP #$41                        A:40 X:00 Y:00 P:25 SP:FB
	// C9E6  F0 07     BEQ $C9EF                       A:40 X:00 Y:00 P:A4 SP:FB
	// C9E8  10 05     BPL $C9EF                       A:40 X:00 Y:00 P:A4 SP:FB
	// C9EA  10 03     BPL $C9EF                       A:40 X:00 Y:00 P:A4 SP:FB
	// C9EC  4C F3 C9  JMP $C9F3                       A:40 X:00 Y:00 P:A4 SP:FB
	// C9F3  EA        NOP                             A:40 X:00 Y:00 P:A4 SP:FB
	// C9F4  A9 80     LDA #$80                        A:40 X:00 Y:00 P:A4 SP:FB
	// C9F6  C9 00     CMP #$00                        A:80 X:00 Y:00 P:A4 SP:FB
	// C9F8  F0 07     BEQ $CA01                       A:80 X:00 Y:00 P:A5 SP:FB
	// C9FA  10 05     BPL $CA01                       A:80 X:00 Y:00 P:A5 SP:FB
	// C9FC  90 03     BCC $CA01                       A:80 X:00 Y:00 P:A5 SP:FB
	// C9FE  4C 05 CA  JMP $CA05                       A:80 X:00 Y:00 P:A5 SP:FB
	// CA05  EA        NOP                             A:80 X:00 Y:00 P:A5 SP:FB
	// CA06  C9 80     CMP #$80                        A:80 X:00 Y:00 P:A5 SP:FB
	// CA08  D0 07     BNE $CA11                       A:80 X:00 Y:00 P:27 SP:FB
	// CA0A  30 05     BMI $CA11                       A:80 X:00 Y:00 P:27 SP:FB
	// CA0C  90 03     BCC $CA11                       A:80 X:00 Y:00 P:27 SP:FB
	// CA0E  4C 15 CA  JMP $CA15                       A:80 X:00 Y:00 P:27 SP:FB
	// CA15  EA        NOP                             A:80 X:00 Y:00 P:27 SP:FB
	// CA16  C9 81     CMP #$81                        A:80 X:00 Y:00 P:27 SP:FB
	// CA18  B0 07     BCS $CA21                       A:80 X:00 Y:00 P:A4 SP:FB
	// CA1A  F0 05     BEQ $CA21                       A:80 X:00 Y:00 P:A4 SP:FB
	// CA1C  10 03     BPL $CA21                       A:80 X:00 Y:00 P:A4 SP:FB
	// CA1E  4C 25 CA  JMP $CA25                       A:80 X:00 Y:00 P:A4 SP:FB
	// CA25  EA        NOP                             A:80 X:00 Y:00 P:A4 SP:FB
	// CA26  C9 7F     CMP #$7F                        A:80 X:00 Y:00 P:A4 SP:FB
	// CA28  90 07     BCC $CA31                       A:80 X:00 Y:00 P:25 SP:FB
	// CA2A  F0 05     BEQ $CA31                       A:80 X:00 Y:00 P:25 SP:FB
	// CA2C  30 03     BMI $CA31                       A:80 X:00 Y:00 P:25 SP:FB
	// CA2E  4C 35 CA  JMP $CA35                       A:80 X:00 Y:00 P:25 SP:FB
	// CA35  EA        NOP                             A:80 X:00 Y:00 P:25 SP:FB
	// CA36  24 01     BIT $01 = FF                    A:80 X:00 Y:00 P:25 SP:FB
	// CA38  A0 40     LDY #$40                        A:80 X:00 Y:00 P:E5 SP:FB
	// CA3A  C0 40     CPY #$40                        A:80 X:00 Y:40 P:65 SP:FB
	// CA3C  D0 09     BNE $CA47                       A:80 X:00 Y:40 P:67 SP:FB
	// CA3E  30 07     BMI $CA47                       A:80 X:00 Y:40 P:67 SP:FB
	// CA40  90 05     BCC $CA47                       A:80 X:00 Y:40 P:67 SP:FB
	// CA42  50 03     BVC $CA47                       A:80 X:00 Y:40 P:67 SP:FB
	// CA44  4C 4B CA  JMP $CA4B                       A:80 X:00 Y:40 P:67 SP:FB
	// CA4B  EA        NOP                             A:80 X:00 Y:40 P:67 SP:FB
	// CA4C  B8        CLV                             A:80 X:00 Y:40 P:67 SP:FB
	// CA4D  C0 3F     CPY #$3F                        A:80 X:00 Y:40 P:27 SP:FB
	// CA4F  F0 09     BEQ $CA5A                       A:80 X:00 Y:40 P:25 SP:FB
	// CA51  30 07     BMI $CA5A                       A:80 X:00 Y:40 P:25 SP:FB
	// CA53  90 05     BCC $CA5A                       A:80 X:00 Y:40 P:25 SP:FB
	// CA55  70 03     BVS $CA5A                       A:80 X:00 Y:40 P:25 SP:FB
	// CA57  4C 5E CA  JMP $CA5E                       A:80 X:00 Y:40 P:25 SP:FB
	// CA5E  EA        NOP                             A:80 X:00 Y:40 P:25 SP:FB
	// CA5F  C0 41     CPY #$41                        A:80 X:00 Y:40 P:25 SP:FB
	// CA61  F0 07     BEQ $CA6A                       A:80 X:00 Y:40 P:A4 SP:FB
	// CA63  10 05     BPL $CA6A                       A:80 X:00 Y:40 P:A4 SP:FB
	// CA65  10 03     BPL $CA6A                       A:80 X:00 Y:40 P:A4 SP:FB
	// CA67  4C 6E CA  JMP $CA6E                       A:80 X:00 Y:40 P:A4 SP:FB
	// CA6E  EA        NOP                             A:80 X:00 Y:40 P:A4 SP:FB
	// CA6F  A0 80     LDY #$80                        A:80 X:00 Y:40 P:A4 SP:FB
	// CA71  C0 00     CPY #$00                        A:80 X:00 Y:80 P:A4 SP:FB
	// CA73  F0 07     BEQ $CA7C                       A:80 X:00 Y:80 P:A5 SP:FB
	// CA75  10 05     BPL $CA7C                       A:80 X:00 Y:80 P:A5 SP:FB
	// CA77  90 03     BCC $CA7C                       A:80 X:00 Y:80 P:A5 SP:FB
	// CA79  4C 80 CA  JMP $CA80                       A:80 X:00 Y:80 P:A5 SP:FB
	// CA80  EA        NOP                             A:80 X:00 Y:80 P:A5 SP:FB
	// CA81  C0 80     CPY #$80                        A:80 X:00 Y:80 P:A5 SP:FB
	// CA83  D0 07     BNE $CA8C                       A:80 X:00 Y:80 P:27 SP:FB
	// CA85  30 05     BMI $CA8C                       A:80 X:00 Y:80 P:27 SP:FB
	// CA87  90 03     BCC $CA8C                       A:80 X:00 Y:80 P:27 SP:FB
	// CA89  4C 90 CA  JMP $CA90                       A:80 X:00 Y:80 P:27 SP:FB
	// CA90  EA        NOP                             A:80 X:00 Y:80 P:27 SP:FB
	// CA91  C0 81     CPY #$81                        A:80 X:00 Y:80 P:27 SP:FB
	// CA93  B0 07     BCS $CA9C                       A:80 X:00 Y:80 P:A4 SP:FB
	// CA95  F0 05     BEQ $CA9C                       A:80 X:00 Y:80 P:A4 SP:FB
	// CA97  10 03     BPL $CA9C                       A:80 X:00 Y:80 P:A4 SP:FB
	// CA99  4C A0 CA  JMP $CAA0                       A:80 X:00 Y:80 P:A4 SP:FB
	// CAA0  EA        NOP                             A:80 X:00 Y:80 P:A4 SP:FB
	// CAA1  C0 7F     CPY #$7F                        A:80 X:00 Y:80 P:A4 SP:FB
	// CAA3  90 07     BCC $CAAC                       A:80 X:00 Y:80 P:25 SP:FB
	// CAA5  F0 05     BEQ $CAAC                       A:80 X:00 Y:80 P:25 SP:FB
	// CAA7  30 03     BMI $CAAC                       A:80 X:00 Y:80 P:25 SP:FB
	// CAA9  4C B0 CA  JMP $CAB0                       A:80 X:00 Y:80 P:25 SP:FB
	// CAB0  EA        NOP                             A:80 X:00 Y:80 P:25 SP:FB
	// CAB1  24 01     BIT $01 = FF                    A:80 X:00 Y:80 P:25 SP:FB
	// CAB3  A2 40     LDX #$40                        A:80 X:00 Y:80 P:E5 SP:FB
	// CAB5  E0 40     CPX #$40                        A:80 X:40 Y:80 P:65 SP:FB
	// CAB7  D0 09     BNE $CAC2                       A:80 X:40 Y:80 P:67 SP:FB
	// CAB9  30 07     BMI $CAC2                       A:80 X:40 Y:80 P:67 SP:FB
	// CABB  90 05     BCC $CAC2                       A:80 X:40 Y:80 P:67 SP:FB
	// CABD  50 03     BVC $CAC2                       A:80 X:40 Y:80 P:67 SP:FB
	// CABF  4C C6 CA  JMP $CAC6                       A:80 X:40 Y:80 P:67 SP:FB
	// CAC6  EA        NOP                             A:80 X:40 Y:80 P:67 SP:FB
	// CAC7  B8        CLV                             A:80 X:40 Y:80 P:67 SP:FB
	// CAC8  E0 3F     CPX #$3F                        A:80 X:40 Y:80 P:27 SP:FB
	// CACA  F0 09     BEQ $CAD5                       A:80 X:40 Y:80 P:25 SP:FB
	// CACC  30 07     BMI $CAD5                       A:80 X:40 Y:80 P:25 SP:FB
	// CACE  90 05     BCC $CAD5                       A:80 X:40 Y:80 P:25 SP:FB
	// CAD0  70 03     BVS $CAD5                       A:80 X:40 Y:80 P:25 SP:FB
	// CAD2  4C D9 CA  JMP $CAD9                       A:80 X:40 Y:80 P:25 SP:FB
	// CAD9  EA        NOP                             A:80 X:40 Y:80 P:25 SP:FB
	// CADA  E0 41     CPX #$41                        A:80 X:40 Y:80 P:25 SP:FB
	// CADC  F0 07     BEQ $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB
	// CADE  10 05     BPL $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB
	// CAE0  10 03     BPL $CAE5                       A:80 X:40 Y:80 P:A4 SP:FB
	// CAE2  4C E9 CA  JMP $CAE9                       A:80 X:40 Y:80 P:A4 SP:FB
	// CAE9  EA        NOP                             A:80 X:40 Y:80 P:A4 SP:FB
	// CAEA  A2 80     LDX #$80                        A:80 X:40 Y:80 P:A4 SP:FB
	// CAEC  E0 00     CPX #$00                        A:80 X:80 Y:80 P:A4 SP:FB
	// CAEE  F0 07     BEQ $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB
	// CAF0  10 05     BPL $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB
	// CAF2  90 03     BCC $CAF7                       A:80 X:80 Y:80 P:A5 SP:FB
	// CAF4  4C FB CA  JMP $CAFB                       A:80 X:80 Y:80 P:A5 SP:FB
	// CAFB  EA        NOP                             A:80 X:80 Y:80 P:A5 SP:FB
	// CAFC  E0 80     CPX #$80                        A:80 X:80 Y:80 P:A5 SP:FB
	// CAFE  D0 07     BNE $CB07                       A:80 X:80 Y:80 P:27 SP:FB
	// CB00  30 05     BMI $CB07                       A:80 X:80 Y:80 P:27 SP:FB
	// CB02  90 03     BCC $CB07                       A:80 X:80 Y:80 P:27 SP:FB
	// CB04  4C 0B CB  JMP $CB0B                       A:80 X:80 Y:80 P:27 SP:FB
	// CB0B  EA        NOP                             A:80 X:80 Y:80 P:27 SP:FB
	// CB0C  E0 81     CPX #$81                        A:80 X:80 Y:80 P:27 SP:FB
	// CB0E  B0 07     BCS $CB17                       A:80 X:80 Y:80 P:A4 SP:FB
	// CB10  F0 05     BEQ $CB17                       A:80 X:80 Y:80 P:A4 SP:FB
	// CB12  10 03     BPL $CB17                       A:80 X:80 Y:80 P:A4 SP:FB
	// CB14  4C 1B CB  JMP $CB1B                       A:80 X:80 Y:80 P:A4 SP:FB
	// CB1B  EA        NOP                             A:80 X:80 Y:80 P:A4 SP:FB
	// CB1C  E0 7F     CPX #$7F                        A:80 X:80 Y:80 P:A4 SP:FB
	// CB1E  90 07     BCC $CB27                       A:80 X:80 Y:80 P:25 SP:FB
	// CB20  F0 05     BEQ $CB27                       A:80 X:80 Y:80 P:25 SP:FB
	// CB22  30 03     BMI $CB27                       A:80 X:80 Y:80 P:25 SP:FB
	// CB24  4C 2B CB  JMP $CB2B                       A:80 X:80 Y:80 P:25 SP:FB
	// CB2B  EA        NOP                             A:80 X:80 Y:80 P:25 SP:FB
	// CB2C  38        SEC                             A:80 X:80 Y:80 P:25 SP:FB
	// CB2D  B8        CLV                             A:80 X:80 Y:80 P:25 SP:FB
	// CB2E  A2 9F     LDX #$9F                        A:80 X:80 Y:80 P:25 SP:FB
	// CB30  F0 09     BEQ $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB
	// CB32  10 07     BPL $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB
	// CB34  70 05     BVS $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB
	// CB36  90 03     BCC $CB3B                       A:80 X:9F Y:80 P:A5 SP:FB
	// CB38  4C 3F CB  JMP $CB3F                       A:80 X:9F Y:80 P:A5 SP:FB
	// CB3F  EA        NOP                             A:80 X:9F Y:80 P:A5 SP:FB
	// CB40  18        CLC                             A:80 X:9F Y:80 P:A5 SP:FB
	// CB41  24 01     BIT $01 = FF                    A:80 X:9F Y:80 P:A4 SP:FB
	// CB43  A2 00     LDX #$00                        A:80 X:9F Y:80 P:E4 SP:FB
	// CB45  D0 09     BNE $CB50                       A:80 X:00 Y:80 P:66 SP:FB
	// CB47  30 07     BMI $CB50                       A:80 X:00 Y:80 P:66 SP:FB
	// CB49  50 05     BVC $CB50                       A:80 X:00 Y:80 P:66 SP:FB
	// CB4B  B0 03     BCS $CB50                       A:80 X:00 Y:80 P:66 SP:FB
	// CB4D  4C 54 CB  JMP $CB54                       A:80 X:00 Y:80 P:66 SP:FB
	// CB54  EA        NOP                             A:80 X:00 Y:80 P:66 SP:FB
	// CB55  38        SEC                             A:80 X:00 Y:80 P:66 SP:FB
	// CB56  B8        CLV                             A:80 X:00 Y:80 P:67 SP:FB
	// CB57  A0 9F     LDY #$9F                        A:80 X:00 Y:80 P:27 SP:FB
	// CB59  F0 09     BEQ $CB64                       A:80 X:00 Y:9F P:A5 SP:FB
	// CB5B  10 07     BPL $CB64                       A:80 X:00 Y:9F P:A5 SP:FB
	// CB5D  70 05     BVS $CB64                       A:80 X:00 Y:9F P:A5 SP:FB
	// CB5F  90 03     BCC $CB64                       A:80 X:00 Y:9F P:A5 SP:FB
	// CB61  4C 68 CB  JMP $CB68                       A:80 X:00 Y:9F P:A5 SP:FB
	// CB68  EA        NOP                             A:80 X:00 Y:9F P:A5 SP:FB
	// CB69  18        CLC                             A:80 X:00 Y:9F P:A5 SP:FB
	// CB6A  24 01     BIT $01 = FF                    A:80 X:00 Y:9F P:A4 SP:FB
	// CB6C  A0 00     LDY #$00                        A:80 X:00 Y:9F P:E4 SP:FB
	// CB6E  D0 09     BNE $CB79                       A:80 X:00 Y:00 P:66 SP:FB
	// CB70  30 07     BMI $CB79                       A:80 X:00 Y:00 P:66 SP:FB
	// CB72  50 05     BVC $CB79                       A:80 X:00 Y:00 P:66 SP:FB
	// CB74  B0 03     BCS $CB79                       A:80 X:00 Y:00 P:66 SP:FB
	// CB76  4C 7D CB  JMP $CB7D                       A:80 X:00 Y:00 P:66 SP:FB
	// CB7D  EA        NOP                             A:80 X:00 Y:00 P:66 SP:FB
	// CB7E  A9 55     LDA #$55                        A:80 X:00 Y:00 P:66 SP:FB
	// CB80  A2 AA     LDX #$AA                        A:55 X:00 Y:00 P:64 SP:FB
	// CB82  A0 33     LDY #$33                        A:55 X:AA Y:00 P:E4 SP:FB
	// CB84  C9 55     CMP #$55                        A:55 X:AA Y:33 P:64 SP:FB
	// CB86  D0 23     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB88  E0 AA     CPX #$AA                        A:55 X:AA Y:33 P:67 SP:FB
	// CB8A  D0 1F     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB8C  C0 33     CPY #$33                        A:55 X:AA Y:33 P:67 SP:FB
	// CB8E  D0 1B     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB90  C9 55     CMP #$55                        A:55 X:AA Y:33 P:67 SP:FB
	// CB92  D0 17     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB94  E0 AA     CPX #$AA                        A:55 X:AA Y:33 P:67 SP:FB
	// CB96  D0 13     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB98  C0 33     CPY #$33                        A:55 X:AA Y:33 P:67 SP:FB
	// CB9A  D0 0F     BNE $CBAB                       A:55 X:AA Y:33 P:67 SP:FB
	// CB9C  C9 56     CMP #$56                        A:55 X:AA Y:33 P:67 SP:FB
	// CB9E  F0 0B     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB
	// CBA0  E0 AB     CPX #$AB                        A:55 X:AA Y:33 P:E4 SP:FB
	// CBA2  F0 07     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB
	// CBA4  C0 34     CPY #$34                        A:55 X:AA Y:33 P:E4 SP:FB
	// CBA6  F0 03     BEQ $CBAB                       A:55 X:AA Y:33 P:E4 SP:FB
	// CBA8  4C AF CB  JMP $CBAF                       A:55 X:AA Y:33 P:E4 SP:FB
	// CBAF  A0 71     LDY #$71                        A:55 X:AA Y:33 P:E4 SP:FB
	// CBB1  20 31 F9  JSR $F931                       A:55 X:AA Y:71 P:64 SP:FB
	// F931  24 01     BIT $01 = FF                    A:55 X:AA Y:71 P:64 SP:F9
	// F933  A9 40     LDA #$40                        A:55 X:AA Y:71 P:E4 SP:F9
	// F935  38        SEC                             A:40 X:AA Y:71 P:64 SP:F9
	// F936  60        RTS                             A:40 X:AA Y:71 P:65 SP:F9
	// CBB4  E9 40     SBC #$40                        A:40 X:AA Y:71 P:65 SP:FB
	// CBB6  20 37 F9  JSR $F937                       A:00 X:AA Y:71 P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:AA Y:71 P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:AA Y:71 P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:AA Y:71 P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:AA Y:71 P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:AA Y:71 P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:AA Y:71 P:27 SP:F9
	// F943  60        RTS                             A:00 X:AA Y:71 P:27 SP:F9
	// CBB9  C8        INY                             A:00 X:AA Y:71 P:27 SP:FB
	// CBBA  20 47 F9  JSR $F947                       A:00 X:AA Y:72 P:25 SP:FB
	// F947  B8        CLV                             A:00 X:AA Y:72 P:25 SP:F9
	// F948  38        SEC                             A:00 X:AA Y:72 P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:00 X:AA Y:72 P:25 SP:F9
	// F94B  60        RTS                             A:40 X:AA Y:72 P:25 SP:F9
	// CBBD  E9 3F     SBC #$3F                        A:40 X:AA Y:72 P:25 SP:FB
	// CBBF  20 4C F9  JSR $F94C                       A:01 X:AA Y:72 P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:AA Y:72 P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:AA Y:72 P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:AA Y:72 P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:AA Y:72 P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:AA Y:72 P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:AA Y:72 P:27 SP:F9
	// F958  60        RTS                             A:01 X:AA Y:72 P:27 SP:F9
	// CBC2  C8        INY                             A:01 X:AA Y:72 P:27 SP:FB
	// CBC3  20 5C F9  JSR $F95C                       A:01 X:AA Y:73 P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:01 X:AA Y:73 P:25 SP:F9
	// F95E  38        SEC                             A:40 X:AA Y:73 P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:AA Y:73 P:25 SP:F9
	// F961  60        RTS                             A:40 X:AA Y:73 P:E5 SP:F9
	// CBC6  E9 41     SBC #$41                        A:40 X:AA Y:73 P:E5 SP:FB
	// CBC8  20 62 F9  JSR $F962                       A:FF X:AA Y:73 P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:AA Y:73 P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:AA Y:73 P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:AA Y:73 P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:AA Y:73 P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:AA Y:73 P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:AA Y:73 P:27 SP:F9
	// F96E  60        RTS                             A:FF X:AA Y:73 P:27 SP:F9
	// CBCB  C8        INY                             A:FF X:AA Y:73 P:27 SP:FB
	// CBCC  20 72 F9  JSR $F972                       A:FF X:AA Y:74 P:25 SP:FB
	// F972  18        CLC                             A:FF X:AA Y:74 P:25 SP:F9
	// F973  A9 80     LDA #$80                        A:FF X:AA Y:74 P:24 SP:F9
	// F975  60        RTS                             A:80 X:AA Y:74 P:A4 SP:F9
	// CBCF  E9 00     SBC #$00                        A:80 X:AA Y:74 P:A4 SP:FB
	// CBD1  20 76 F9  JSR $F976                       A:7F X:AA Y:74 P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:AA Y:74 P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:AA Y:74 P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:AA Y:74 P:67 SP:F9
	// F97C  60        RTS                             A:7F X:AA Y:74 P:67 SP:F9
	// CBD4  C8        INY                             A:7F X:AA Y:74 P:67 SP:FB
	// CBD5  20 80 F9  JSR $F980                       A:7F X:AA Y:75 P:65 SP:FB
	// F980  38        SEC                             A:7F X:AA Y:75 P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:AA Y:75 P:65 SP:F9
	// F983  60        RTS                             A:81 X:AA Y:75 P:E5 SP:F9
	// CBD8  E9 7F     SBC #$7F                        A:81 X:AA Y:75 P:E5 SP:FB
	// CBDA  20 84 F9  JSR $F984                       A:02 X:AA Y:75 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:AA Y:75 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:AA Y:75 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:AA Y:75 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:AA Y:75 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:AA Y:75 P:67 SP:F9
	// CBDD  60        RTS                             A:02 X:AA Y:75 P:67 SP:FB
	// C606  20 DE CB  JSR $CBDE                       A:02 X:AA Y:75 P:67 SP:FD
	// CBDE  EA        NOP                             A:02 X:AA Y:75 P:67 SP:FB
	// CBDF  A9 FF     LDA #$FF                        A:02 X:AA Y:75 P:67 SP:FB
	// CBE1  85 01     STA $01 = FF                    A:FF X:AA Y:75 P:E5 SP:FB
	// CBE3  A9 44     LDA #$44                        A:FF X:AA Y:75 P:E5 SP:FB
	// CBE5  A2 55     LDX #$55                        A:44 X:AA Y:75 P:65 SP:FB
	// CBE7  A0 66     LDY #$66                        A:44 X:55 Y:75 P:65 SP:FB
	// CBE9  E8        INX                             A:44 X:55 Y:66 P:65 SP:FB
	// CBEA  88        DEY                             A:44 X:56 Y:66 P:65 SP:FB
	// CBEB  E0 56     CPX #$56                        A:44 X:56 Y:65 P:65 SP:FB
	// CBED  D0 21     BNE $CC10                       A:44 X:56 Y:65 P:67 SP:FB
	// CBEF  C0 65     CPY #$65                        A:44 X:56 Y:65 P:67 SP:FB
	// CBF1  D0 1D     BNE $CC10                       A:44 X:56 Y:65 P:67 SP:FB
	// CBF3  E8        INX                             A:44 X:56 Y:65 P:67 SP:FB
	// CBF4  E8        INX                             A:44 X:57 Y:65 P:65 SP:FB
	// CBF5  88        DEY                             A:44 X:58 Y:65 P:65 SP:FB
	// CBF6  88        DEY                             A:44 X:58 Y:64 P:65 SP:FB
	// CBF7  E0 58     CPX #$58                        A:44 X:58 Y:63 P:65 SP:FB
	// CBF9  D0 15     BNE $CC10                       A:44 X:58 Y:63 P:67 SP:FB
	// CBFB  C0 63     CPY #$63                        A:44 X:58 Y:63 P:67 SP:FB
	// CBFD  D0 11     BNE $CC10                       A:44 X:58 Y:63 P:67 SP:FB
	// CBFF  CA        DEX                             A:44 X:58 Y:63 P:67 SP:FB
	// CC00  C8        INY                             A:44 X:57 Y:63 P:65 SP:FB
	// CC01  E0 57     CPX #$57                        A:44 X:57 Y:64 P:65 SP:FB
	// CC03  D0 0B     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB
	// CC05  C0 64     CPY #$64                        A:44 X:57 Y:64 P:67 SP:FB
	// CC07  D0 07     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB
	// CC09  C9 44     CMP #$44                        A:44 X:57 Y:64 P:67 SP:FB
	// CC0B  D0 03     BNE $CC10                       A:44 X:57 Y:64 P:67 SP:FB
	// CC0D  4C 14 CC  JMP $CC14                       A:44 X:57 Y:64 P:67 SP:FB
	// CC14  EA        NOP                             A:44 X:57 Y:64 P:67 SP:FB
	// CC15  38        SEC                             A:44 X:57 Y:64 P:67 SP:FB
	// CC16  A2 69     LDX #$69                        A:44 X:57 Y:64 P:67 SP:FB
	// CC18  A9 96     LDA #$96                        A:44 X:69 Y:64 P:65 SP:FB
	// CC1A  24 01     BIT $01 = FF                    A:96 X:69 Y:64 P:E5 SP:FB
	// CC1C  A0 FF     LDY #$FF                        A:96 X:69 Y:64 P:E5 SP:FB
	// CC1E  C8        INY                             A:96 X:69 Y:FF P:E5 SP:FB
	// CC1F  D0 3D     BNE $CC5E                       A:96 X:69 Y:00 P:67 SP:FB
	// CC21  30 3B     BMI $CC5E                       A:96 X:69 Y:00 P:67 SP:FB
	// CC23  90 39     BCC $CC5E                       A:96 X:69 Y:00 P:67 SP:FB
	// CC25  50 37     BVC $CC5E                       A:96 X:69 Y:00 P:67 SP:FB
	// CC27  C0 00     CPY #$00                        A:96 X:69 Y:00 P:67 SP:FB
	// CC29  D0 33     BNE $CC5E                       A:96 X:69 Y:00 P:67 SP:FB
	// CC2B  C8        INY                             A:96 X:69 Y:00 P:67 SP:FB
	// CC2C  F0 30     BEQ $CC5E                       A:96 X:69 Y:01 P:65 SP:FB
	// CC2E  30 2E     BMI $CC5E                       A:96 X:69 Y:01 P:65 SP:FB
	// CC30  90 2C     BCC $CC5E                       A:96 X:69 Y:01 P:65 SP:FB
	// CC32  50 2A     BVC $CC5E                       A:96 X:69 Y:01 P:65 SP:FB
	// CC34  18        CLC                             A:96 X:69 Y:01 P:65 SP:FB
	// CC35  B8        CLV                             A:96 X:69 Y:01 P:64 SP:FB
	// CC36  A0 00     LDY #$00                        A:96 X:69 Y:01 P:24 SP:FB
	// CC38  88        DEY                             A:96 X:69 Y:00 P:26 SP:FB
	// CC39  F0 23     BEQ $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB
	// CC3B  10 21     BPL $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB
	// CC3D  B0 1F     BCS $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB
	// CC3F  70 1D     BVS $CC5E                       A:96 X:69 Y:FF P:A4 SP:FB
	// CC41  C0 FF     CPY #$FF                        A:96 X:69 Y:FF P:A4 SP:FB
	// CC43  D0 19     BNE $CC5E                       A:96 X:69 Y:FF P:27 SP:FB
	// CC45  18        CLC                             A:96 X:69 Y:FF P:27 SP:FB
	// CC46  88        DEY                             A:96 X:69 Y:FF P:26 SP:FB
	// CC47  F0 15     BEQ $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB
	// CC49  10 13     BPL $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB
	// CC4B  B0 11     BCS $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB
	// CC4D  70 0F     BVS $CC5E                       A:96 X:69 Y:FE P:A4 SP:FB
	// CC4F  C0 FE     CPY #$FE                        A:96 X:69 Y:FE P:A4 SP:FB
	// CC51  D0 0B     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB
	// CC53  C9 96     CMP #$96                        A:96 X:69 Y:FE P:27 SP:FB
	// CC55  D0 07     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB
	// CC57  E0 69     CPX #$69                        A:96 X:69 Y:FE P:27 SP:FB
	// CC59  D0 03     BNE $CC5E                       A:96 X:69 Y:FE P:27 SP:FB
	// CC5B  4C 62 CC  JMP $CC62                       A:96 X:69 Y:FE P:27 SP:FB
	// CC62  EA        NOP                             A:96 X:69 Y:FE P:27 SP:FB
	// CC63  38        SEC                             A:96 X:69 Y:FE P:27 SP:FB
	// CC64  A0 69     LDY #$69                        A:96 X:69 Y:FE P:27 SP:FB
	// CC66  A9 96     LDA #$96                        A:96 X:69 Y:69 P:25 SP:FB
	// CC68  24 01     BIT $01 = FF                    A:96 X:69 Y:69 P:A5 SP:FB
	// CC6A  A2 FF     LDX #$FF                        A:96 X:69 Y:69 P:E5 SP:FB
	// CC6C  E8        INX                             A:96 X:FF Y:69 P:E5 SP:FB
	// CC6D  D0 3D     BNE $CCAC                       A:96 X:00 Y:69 P:67 SP:FB
	// CC6F  30 3B     BMI $CCAC                       A:96 X:00 Y:69 P:67 SP:FB
	// CC71  90 39     BCC $CCAC                       A:96 X:00 Y:69 P:67 SP:FB
	// CC73  50 37     BVC $CCAC                       A:96 X:00 Y:69 P:67 SP:FB
	// CC75  E0 00     CPX #$00                        A:96 X:00 Y:69 P:67 SP:FB
	// CC77  D0 33     BNE $CCAC                       A:96 X:00 Y:69 P:67 SP:FB
	// CC79  E8        INX                             A:96 X:00 Y:69 P:67 SP:FB
	// CC7A  F0 30     BEQ $CCAC                       A:96 X:01 Y:69 P:65 SP:FB
	// CC7C  30 2E     BMI $CCAC                       A:96 X:01 Y:69 P:65 SP:FB
	// CC7E  90 2C     BCC $CCAC                       A:96 X:01 Y:69 P:65 SP:FB
	// CC80  50 2A     BVC $CCAC                       A:96 X:01 Y:69 P:65 SP:FB
	// CC82  18        CLC                             A:96 X:01 Y:69 P:65 SP:FB
	// CC83  B8        CLV                             A:96 X:01 Y:69 P:64 SP:FB
	// CC84  A2 00     LDX #$00                        A:96 X:01 Y:69 P:24 SP:FB
	// CC86  CA        DEX                             A:96 X:00 Y:69 P:26 SP:FB
	// CC87  F0 23     BEQ $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB
	// CC89  10 21     BPL $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB
	// CC8B  B0 1F     BCS $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB
	// CC8D  70 1D     BVS $CCAC                       A:96 X:FF Y:69 P:A4 SP:FB
	// CC8F  E0 FF     CPX #$FF                        A:96 X:FF Y:69 P:A4 SP:FB
	// CC91  D0 19     BNE $CCAC                       A:96 X:FF Y:69 P:27 SP:FB
	// CC93  18        CLC                             A:96 X:FF Y:69 P:27 SP:FB
	// CC94  CA        DEX                             A:96 X:FF Y:69 P:26 SP:FB
	// CC95  F0 15     BEQ $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB
	// CC97  10 13     BPL $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB
	// CC99  B0 11     BCS $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB
	// CC9B  70 0F     BVS $CCAC                       A:96 X:FE Y:69 P:A4 SP:FB
	// CC9D  E0 FE     CPX #$FE                        A:96 X:FE Y:69 P:A4 SP:FB
	// CC9F  D0 0B     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB
	// CCA1  C9 96     CMP #$96                        A:96 X:FE Y:69 P:27 SP:FB
	// CCA3  D0 07     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB
	// CCA5  C0 69     CPY #$69                        A:96 X:FE Y:69 P:27 SP:FB
	// CCA7  D0 03     BNE $CCAC                       A:96 X:FE Y:69 P:27 SP:FB
	// CCA9  4C B0 CC  JMP $CCB0                       A:96 X:FE Y:69 P:27 SP:FB
	// CCB0  EA        NOP                             A:96 X:FE Y:69 P:27 SP:FB
	// CCB1  A9 85     LDA #$85                        A:96 X:FE Y:69 P:27 SP:FB
	// CCB3  A2 34     LDX #$34                        A:85 X:FE Y:69 P:A5 SP:FB
	// CCB5  A0 99     LDY #$99                        A:85 X:34 Y:69 P:25 SP:FB
	// CCB7  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB
	// CCB8  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB
	// CCBA  A8        TAY                             A:85 X:34 Y:99 P:E4 SP:FB
	// CCBB  F0 2E     BEQ $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB
	// CCBD  B0 2C     BCS $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB
	// CCBF  50 2A     BVC $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB
	// CCC1  10 28     BPL $CCEB                       A:85 X:34 Y:85 P:E4 SP:FB
	// CCC3  C9 85     CMP #$85                        A:85 X:34 Y:85 P:E4 SP:FB
	// CCC5  D0 24     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB
	// CCC7  E0 34     CPX #$34                        A:85 X:34 Y:85 P:67 SP:FB
	// CCC9  D0 20     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB
	// CCCB  C0 85     CPY #$85                        A:85 X:34 Y:85 P:67 SP:FB
	// CCCD  D0 1C     BNE $CCEB                       A:85 X:34 Y:85 P:67 SP:FB
	// CCCF  A9 00     LDA #$00                        A:85 X:34 Y:85 P:67 SP:FB
	// CCD1  38        SEC                             A:00 X:34 Y:85 P:67 SP:FB
	// CCD2  B8        CLV                             A:00 X:34 Y:85 P:67 SP:FB
	// CCD3  A8        TAY                             A:00 X:34 Y:85 P:27 SP:FB
	// CCD4  D0 15     BNE $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCD6  90 13     BCC $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCD8  70 11     BVS $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCDA  30 0F     BMI $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCDC  C9 00     CMP #$00                        A:00 X:34 Y:00 P:27 SP:FB
	// CCDE  D0 0B     BNE $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCE0  E0 34     CPX #$34                        A:00 X:34 Y:00 P:27 SP:FB
	// CCE2  D0 07     BNE $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCE4  C0 00     CPY #$00                        A:00 X:34 Y:00 P:27 SP:FB
	// CCE6  D0 03     BNE $CCEB                       A:00 X:34 Y:00 P:27 SP:FB
	// CCE8  4C EF CC  JMP $CCEF                       A:00 X:34 Y:00 P:27 SP:FB
	// CCEF  EA        NOP                             A:00 X:34 Y:00 P:27 SP:FB
	// CCF0  A9 85     LDA #$85                        A:00 X:34 Y:00 P:27 SP:FB
	// CCF2  A2 34     LDX #$34                        A:85 X:34 Y:00 P:A5 SP:FB
	// CCF4  A0 99     LDY #$99                        A:85 X:34 Y:00 P:25 SP:FB
	// CCF6  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB
	// CCF7  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB
	// CCF9  AA        TAX                             A:85 X:34 Y:99 P:E4 SP:FB
	// CCFA  F0 2E     BEQ $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB
	// CCFC  B0 2C     BCS $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB
	// CCFE  50 2A     BVC $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB
	// CD00  10 28     BPL $CD2A                       A:85 X:85 Y:99 P:E4 SP:FB
	// CD02  C9 85     CMP #$85                        A:85 X:85 Y:99 P:E4 SP:FB
	// CD04  D0 24     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB
	// CD06  E0 85     CPX #$85                        A:85 X:85 Y:99 P:67 SP:FB
	// CD08  D0 20     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB
	// CD0A  C0 99     CPY #$99                        A:85 X:85 Y:99 P:67 SP:FB
	// CD0C  D0 1C     BNE $CD2A                       A:85 X:85 Y:99 P:67 SP:FB
	// CD0E  A9 00     LDA #$00                        A:85 X:85 Y:99 P:67 SP:FB
	// CD10  38        SEC                             A:00 X:85 Y:99 P:67 SP:FB
	// CD11  B8        CLV                             A:00 X:85 Y:99 P:67 SP:FB
	// CD12  AA        TAX                             A:00 X:85 Y:99 P:27 SP:FB
	// CD13  D0 15     BNE $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD15  90 13     BCC $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD17  70 11     BVS $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD19  30 0F     BMI $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD1B  C9 00     CMP #$00                        A:00 X:00 Y:99 P:27 SP:FB
	// CD1D  D0 0B     BNE $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD1F  E0 00     CPX #$00                        A:00 X:00 Y:99 P:27 SP:FB
	// CD21  D0 07     BNE $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD23  C0 99     CPY #$99                        A:00 X:00 Y:99 P:27 SP:FB
	// CD25  D0 03     BNE $CD2A                       A:00 X:00 Y:99 P:27 SP:FB
	// CD27  4C 2E CD  JMP $CD2E                       A:00 X:00 Y:99 P:27 SP:FB
	// CD2E  EA        NOP                             A:00 X:00 Y:99 P:27 SP:FB
	// CD2F  A9 85     LDA #$85                        A:00 X:00 Y:99 P:27 SP:FB
	// CD31  A2 34     LDX #$34                        A:85 X:00 Y:99 P:A5 SP:FB
	// CD33  A0 99     LDY #$99                        A:85 X:34 Y:99 P:25 SP:FB
	// CD35  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB
	// CD36  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB
	// CD38  98        TYA                             A:85 X:34 Y:99 P:E4 SP:FB
	// CD39  F0 2E     BEQ $CD69                       A:99 X:34 Y:99 P:E4 SP:FB
	// CD3B  B0 2C     BCS $CD69                       A:99 X:34 Y:99 P:E4 SP:FB
	// CD3D  50 2A     BVC $CD69                       A:99 X:34 Y:99 P:E4 SP:FB
	// CD3F  10 28     BPL $CD69                       A:99 X:34 Y:99 P:E4 SP:FB
	// CD41  C9 99     CMP #$99                        A:99 X:34 Y:99 P:E4 SP:FB
	// CD43  D0 24     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB
	// CD45  E0 34     CPX #$34                        A:99 X:34 Y:99 P:67 SP:FB
	// CD47  D0 20     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB
	// CD49  C0 99     CPY #$99                        A:99 X:34 Y:99 P:67 SP:FB
	// CD4B  D0 1C     BNE $CD69                       A:99 X:34 Y:99 P:67 SP:FB
	// CD4D  A0 00     LDY #$00                        A:99 X:34 Y:99 P:67 SP:FB
	// CD4F  38        SEC                             A:99 X:34 Y:00 P:67 SP:FB
	// CD50  B8        CLV                             A:99 X:34 Y:00 P:67 SP:FB
	// CD51  98        TYA                             A:99 X:34 Y:00 P:27 SP:FB
	// CD52  D0 15     BNE $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD54  90 13     BCC $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD56  70 11     BVS $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD58  30 0F     BMI $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD5A  C9 00     CMP #$00                        A:00 X:34 Y:00 P:27 SP:FB
	// CD5C  D0 0B     BNE $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD5E  E0 34     CPX #$34                        A:00 X:34 Y:00 P:27 SP:FB
	// CD60  D0 07     BNE $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD62  C0 00     CPY #$00                        A:00 X:34 Y:00 P:27 SP:FB
	// CD64  D0 03     BNE $CD69                       A:00 X:34 Y:00 P:27 SP:FB
	// CD66  4C 6D CD  JMP $CD6D                       A:00 X:34 Y:00 P:27 SP:FB
	// CD6D  EA        NOP                             A:00 X:34 Y:00 P:27 SP:FB
	// CD6E  A9 85     LDA #$85                        A:00 X:34 Y:00 P:27 SP:FB
	// CD70  A2 34     LDX #$34                        A:85 X:34 Y:00 P:A5 SP:FB
	// CD72  A0 99     LDY #$99                        A:85 X:34 Y:00 P:25 SP:FB
	// CD74  18        CLC                             A:85 X:34 Y:99 P:A5 SP:FB
	// CD75  24 01     BIT $01 = FF                    A:85 X:34 Y:99 P:A4 SP:FB
	// CD77  8A        TXA                             A:85 X:34 Y:99 P:E4 SP:FB
	// CD78  F0 2E     BEQ $CDA8                       A:34 X:34 Y:99 P:64 SP:FB
	// CD7A  B0 2C     BCS $CDA8                       A:34 X:34 Y:99 P:64 SP:FB
	// CD7C  50 2A     BVC $CDA8                       A:34 X:34 Y:99 P:64 SP:FB
	// CD7E  30 28     BMI $CDA8                       A:34 X:34 Y:99 P:64 SP:FB
	// CD80  C9 34     CMP #$34                        A:34 X:34 Y:99 P:64 SP:FB
	// CD82  D0 24     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB
	// CD84  E0 34     CPX #$34                        A:34 X:34 Y:99 P:67 SP:FB
	// CD86  D0 20     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB
	// CD88  C0 99     CPY #$99                        A:34 X:34 Y:99 P:67 SP:FB
	// CD8A  D0 1C     BNE $CDA8                       A:34 X:34 Y:99 P:67 SP:FB
	// CD8C  A2 00     LDX #$00                        A:34 X:34 Y:99 P:67 SP:FB
	// CD8E  38        SEC                             A:34 X:00 Y:99 P:67 SP:FB
	// CD8F  B8        CLV                             A:34 X:00 Y:99 P:67 SP:FB
	// CD90  8A        TXA                             A:34 X:00 Y:99 P:27 SP:FB
	// CD91  D0 15     BNE $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CD93  90 13     BCC $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CD95  70 11     BVS $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CD97  30 0F     BMI $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CD99  C9 00     CMP #$00                        A:00 X:00 Y:99 P:27 SP:FB
	// CD9B  D0 0B     BNE $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CD9D  E0 00     CPX #$00                        A:00 X:00 Y:99 P:27 SP:FB
	// CD9F  D0 07     BNE $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CDA1  C0 99     CPY #$99                        A:00 X:00 Y:99 P:27 SP:FB
	// CDA3  D0 03     BNE $CDA8                       A:00 X:00 Y:99 P:27 SP:FB
	// CDA5  4C AC CD  JMP $CDAC                       A:00 X:00 Y:99 P:27 SP:FB
	// CDAC  EA        NOP                             A:00 X:00 Y:99 P:27 SP:FB
	// CDAD  BA        TSX                             A:00 X:00 Y:99 P:27 SP:FB
	// CDAE  8E FF 07  STX $07FF = 00                  A:00 X:FB Y:99 P:A5 SP:FB
	// CDB1  A0 33     LDY #$33                        A:00 X:FB Y:99 P:A5 SP:FB
	// CDB3  A2 69     LDX #$69                        A:00 X:FB Y:33 P:25 SP:FB
	// CDB5  A9 84     LDA #$84                        A:00 X:69 Y:33 P:25 SP:FB
	// CDB7  18        CLC                             A:84 X:69 Y:33 P:A5 SP:FB
	// CDB8  24 01     BIT $01 = FF                    A:84 X:69 Y:33 P:A4 SP:FB
	// CDBA  9A        TXS                             A:84 X:69 Y:33 P:E4 SP:FB
	// CDBB  F0 32     BEQ $CDEF                       A:84 X:69 Y:33 P:E4 SP:69
	// CDBD  10 30     BPL $CDEF                       A:84 X:69 Y:33 P:E4 SP:69
	// CDBF  B0 2E     BCS $CDEF                       A:84 X:69 Y:33 P:E4 SP:69
	// CDC1  50 2C     BVC $CDEF                       A:84 X:69 Y:33 P:E4 SP:69
	// CDC3  C9 84     CMP #$84                        A:84 X:69 Y:33 P:E4 SP:69
	// CDC5  D0 28     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69
	// CDC7  E0 69     CPX #$69                        A:84 X:69 Y:33 P:67 SP:69
	// CDC9  D0 24     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69
	// CDCB  C0 33     CPY #$33                        A:84 X:69 Y:33 P:67 SP:69
	// CDCD  D0 20     BNE $CDEF                       A:84 X:69 Y:33 P:67 SP:69
	// CDCF  A0 01     LDY #$01                        A:84 X:69 Y:33 P:67 SP:69
	// CDD1  A9 04     LDA #$04                        A:84 X:69 Y:01 P:65 SP:69
	// CDD3  38        SEC                             A:04 X:69 Y:01 P:65 SP:69
	// CDD4  B8        CLV                             A:04 X:69 Y:01 P:65 SP:69
	// CDD5  A2 00     LDX #$00                        A:04 X:69 Y:01 P:25 SP:69
	// CDD7  BA        TSX                             A:04 X:00 Y:01 P:27 SP:69
	// CDD8  F0 15     BEQ $CDEF                       A:04 X:69 Y:01 P:25 SP:69
	// CDDA  30 13     BMI $CDEF                       A:04 X:69 Y:01 P:25 SP:69
	// CDDC  90 11     BCC $CDEF                       A:04 X:69 Y:01 P:25 SP:69
	// CDDE  70 0F     BVS $CDEF                       A:04 X:69 Y:01 P:25 SP:69
	// CDE0  E0 69     CPX #$69                        A:04 X:69 Y:01 P:25 SP:69
	// CDE2  D0 0B     BNE $CDEF                       A:04 X:69 Y:01 P:27 SP:69
	// CDE4  C9 04     CMP #$04                        A:04 X:69 Y:01 P:27 SP:69
	// CDE6  D0 07     BNE $CDEF                       A:04 X:69 Y:01 P:27 SP:69
	// CDE8  C0 01     CPY #$01                        A:04 X:69 Y:01 P:27 SP:69
	// CDEA  D0 03     BNE $CDEF                       A:04 X:69 Y:01 P:27 SP:69
	// CDEC  4C F3 CD  JMP $CDF3                       A:04 X:69 Y:01 P:27 SP:69
	// CDF3  AE FF 07  LDX $07FF = FB                  A:04 X:69 Y:01 P:27 SP:69
	// CDF6  9A        TXS                             A:04 X:FB Y:01 P:A5 SP:69
	// CDF7  60        RTS                             A:04 X:FB Y:01 P:A5 SP:FB
	// C609  20 F8 CD  JSR $CDF8                       A:04 X:FB Y:01 P:A5 SP:FD
	// CDF8  A9 FF     LDA #$FF                        A:04 X:FB Y:01 P:A5 SP:FB
	// CDFA  85 01     STA $01 = FF                    A:FF X:FB Y:01 P:A5 SP:FB
	// CDFC  BA        TSX                             A:FF X:FB Y:01 P:A5 SP:FB
	// CDFD  8E FF 07  STX $07FF = FB                  A:FF X:FB Y:01 P:A5 SP:FB
	// CE00  EA        NOP                             A:FF X:FB Y:01 P:A5 SP:FB
	// CE01  A2 80     LDX #$80                        A:FF X:FB Y:01 P:A5 SP:FB
	// CE03  9A        TXS                             A:FF X:80 Y:01 P:A5 SP:FB
	// CE04  A9 33     LDA #$33                        A:FF X:80 Y:01 P:A5 SP:80
	// CE06  48        PHA                             A:33 X:80 Y:01 P:25 SP:80
	// CE07  A9 69     LDA #$69                        A:33 X:80 Y:01 P:25 SP:7F
	// CE09  48        PHA                             A:69 X:80 Y:01 P:25 SP:7F
	// CE0A  BA        TSX                             A:69 X:80 Y:01 P:25 SP:7E
	// CE0B  E0 7E     CPX #$7E                        A:69 X:7E Y:01 P:25 SP:7E
	// CE0D  D0 20     BNE $CE2F                       A:69 X:7E Y:01 P:27 SP:7E
	// CE0F  68        PLA                             A:69 X:7E Y:01 P:27 SP:7E
	// CE10  C9 69     CMP #$69                        A:69 X:7E Y:01 P:25 SP:7F
	// CE12  D0 1B     BNE $CE2F                       A:69 X:7E Y:01 P:27 SP:7F
	// CE14  68        PLA                             A:69 X:7E Y:01 P:27 SP:7F
	// CE15  C9 33     CMP #$33                        A:33 X:7E Y:01 P:25 SP:80
	// CE17  D0 16     BNE $CE2F                       A:33 X:7E Y:01 P:27 SP:80
	// CE19  BA        TSX                             A:33 X:7E Y:01 P:27 SP:80
	// CE1A  E0 80     CPX #$80                        A:33 X:80 Y:01 P:A5 SP:80
	// CE1C  D0 11     BNE $CE2F                       A:33 X:80 Y:01 P:27 SP:80
	// CE1E  AD 80 01  LDA $0180 = 33                  A:33 X:80 Y:01 P:27 SP:80
	// CE21  C9 33     CMP #$33                        A:33 X:80 Y:01 P:25 SP:80
	// CE23  D0 0A     BNE $CE2F                       A:33 X:80 Y:01 P:27 SP:80
	// CE25  AD 7F 01  LDA $017F = 69                  A:33 X:80 Y:01 P:27 SP:80
	// CE28  C9 69     CMP #$69                        A:69 X:80 Y:01 P:25 SP:80
	// CE2A  D0 03     BNE $CE2F                       A:69 X:80 Y:01 P:27 SP:80
	// CE2C  4C 33 CE  JMP $CE33                       A:69 X:80 Y:01 P:27 SP:80
	// CE33  EA        NOP                             A:69 X:80 Y:01 P:27 SP:80
	// CE34  A2 80     LDX #$80                        A:69 X:80 Y:01 P:27 SP:80
	// CE36  9A        TXS                             A:69 X:80 Y:01 P:A5 SP:80
	// CE37  20 3D CE  JSR $CE3D                       A:69 X:80 Y:01 P:A5 SP:80
	// CE3D  BA        TSX                             A:69 X:80 Y:01 P:A5 SP:7E
	// CE3E  E0 7E     CPX #$7E                        A:69 X:7E Y:01 P:25 SP:7E
	// CE40  D0 19     BNE $CE5B                       A:69 X:7E Y:01 P:27 SP:7E
	// CE42  68        PLA                             A:69 X:7E Y:01 P:27 SP:7E
	// CE43  68        PLA                             A:39 X:7E Y:01 P:25 SP:7F
	// CE44  BA        TSX                             A:CE X:7E Y:01 P:A5 SP:80
	// CE45  E0 80     CPX #$80                        A:CE X:80 Y:01 P:A5 SP:80
	// CE47  D0 12     BNE $CE5B                       A:CE X:80 Y:01 P:27 SP:80
	// CE49  A9 00     LDA #$00                        A:CE X:80 Y:01 P:27 SP:80
	// CE4B  20 4E CE  JSR $CE4E                       A:00 X:80 Y:01 P:27 SP:80
	// CE4E  68        PLA                             A:00 X:80 Y:01 P:27 SP:7E
	// CE4F  C9 4D     CMP #$4D                        A:4D X:80 Y:01 P:25 SP:7F
	// CE51  D0 08     BNE $CE5B                       A:4D X:80 Y:01 P:27 SP:7F
	// CE53  68        PLA                             A:4D X:80 Y:01 P:27 SP:7F
	// CE54  C9 CE     CMP #$CE                        A:CE X:80 Y:01 P:A5 SP:80
	// CE56  D0 03     BNE $CE5B                       A:CE X:80 Y:01 P:27 SP:80
	// CE58  4C 5F CE  JMP $CE5F                       A:CE X:80 Y:01 P:27 SP:80
	// CE5F  EA        NOP                             A:CE X:80 Y:01 P:27 SP:80
	// CE60  A9 CE     LDA #$CE                        A:CE X:80 Y:01 P:27 SP:80
	// CE62  48        PHA                             A:CE X:80 Y:01 P:A5 SP:80
	// CE63  A9 66     LDA #$66                        A:CE X:80 Y:01 P:A5 SP:7F
	// CE65  48        PHA                             A:66 X:80 Y:01 P:25 SP:7F
	// CE66  60        RTS                             A:66 X:80 Y:01 P:25 SP:7E
	// CE67  A2 77     LDX #$77                        A:66 X:80 Y:01 P:25 SP:80
	// CE69  A0 69     LDY #$69                        A:66 X:77 Y:01 P:25 SP:80
	// CE6B  18        CLC                             A:66 X:77 Y:69 P:25 SP:80
	// CE6C  24 01     BIT $01 = FF                    A:66 X:77 Y:69 P:24 SP:80
	// CE6E  A9 83     LDA #$83                        A:66 X:77 Y:69 P:E4 SP:80
	// CE70  20 66 CE  JSR $CE66                       A:83 X:77 Y:69 P:E4 SP:80
	// CE66  60        RTS                             A:83 X:77 Y:69 P:E4 SP:7E
	// CE73  F0 24     BEQ $CE99                       A:83 X:77 Y:69 P:E4 SP:80
	// CE75  10 22     BPL $CE99                       A:83 X:77 Y:69 P:E4 SP:80
	// CE77  B0 20     BCS $CE99                       A:83 X:77 Y:69 P:E4 SP:80
	// CE79  50 1E     BVC $CE99                       A:83 X:77 Y:69 P:E4 SP:80
	// CE7B  C9 83     CMP #$83                        A:83 X:77 Y:69 P:E4 SP:80
	// CE7D  D0 1A     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80
	// CE7F  C0 69     CPY #$69                        A:83 X:77 Y:69 P:67 SP:80
	// CE81  D0 16     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80
	// CE83  E0 77     CPX #$77                        A:83 X:77 Y:69 P:67 SP:80
	// CE85  D0 12     BNE $CE99                       A:83 X:77 Y:69 P:67 SP:80
	// CE87  38        SEC                             A:83 X:77 Y:69 P:67 SP:80
	// CE88  B8        CLV                             A:83 X:77 Y:69 P:67 SP:80
	// CE89  A9 00     LDA #$00                        A:83 X:77 Y:69 P:27 SP:80
	// CE8B  20 66 CE  JSR $CE66                       A:00 X:77 Y:69 P:27 SP:80
	// CE66  60        RTS                             A:00 X:77 Y:69 P:27 SP:7E
	// CE8E  D0 09     BNE $CE99                       A:00 X:77 Y:69 P:27 SP:80
	// CE90  30 07     BMI $CE99                       A:00 X:77 Y:69 P:27 SP:80
	// CE92  90 05     BCC $CE99                       A:00 X:77 Y:69 P:27 SP:80
	// CE94  70 03     BVS $CE99                       A:00 X:77 Y:69 P:27 SP:80
	// CE96  4C 9D CE  JMP $CE9D                       A:00 X:77 Y:69 P:27 SP:80
	// CE9D  EA        NOP                             A:00 X:77 Y:69 P:27 SP:80
	// CE9E  A9 CE     LDA #$CE                        A:00 X:77 Y:69 P:27 SP:80
	// CEA0  48        PHA                             A:CE X:77 Y:69 P:A5 SP:80
	// CEA1  A9 AE     LDA #$AE                        A:CE X:77 Y:69 P:A5 SP:7F
	// CEA3  48        PHA                             A:AE X:77 Y:69 P:A5 SP:7F
	// CEA4  A9 65     LDA #$65                        A:AE X:77 Y:69 P:A5 SP:7E
	// CEA6  48        PHA                             A:65 X:77 Y:69 P:25 SP:7E
	// CEA7  A9 55     LDA #$55                        A:65 X:77 Y:69 P:25 SP:7D
	// CEA9  A0 88     LDY #$88                        A:55 X:77 Y:69 P:25 SP:7D
	// CEAB  A2 99     LDX #$99                        A:55 X:77 Y:88 P:A5 SP:7D
	// CEAD  40        RTI                             A:55 X:99 Y:88 P:A5 SP:7D
	// CEAE  30 35     BMI $CEE5                       A:55 X:99 Y:88 P:65 SP:80
	// CEB0  50 33     BVC $CEE5                       A:55 X:99 Y:88 P:65 SP:80
	// CEB2  F0 31     BEQ $CEE5                       A:55 X:99 Y:88 P:65 SP:80
	// CEB4  90 2F     BCC $CEE5                       A:55 X:99 Y:88 P:65 SP:80
	// CEB6  C9 55     CMP #$55                        A:55 X:99 Y:88 P:65 SP:80
	// CEB8  D0 2B     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80
	// CEBA  C0 88     CPY #$88                        A:55 X:99 Y:88 P:67 SP:80
	// CEBC  D0 27     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80
	// CEBE  E0 99     CPX #$99                        A:55 X:99 Y:88 P:67 SP:80
	// CEC0  D0 23     BNE $CEE5                       A:55 X:99 Y:88 P:67 SP:80
	// CEC2  A9 CE     LDA #$CE                        A:55 X:99 Y:88 P:67 SP:80
	// CEC4  48        PHA                             A:CE X:99 Y:88 P:E5 SP:80
	// CEC5  A9 CE     LDA #$CE                        A:CE X:99 Y:88 P:E5 SP:7F
	// CEC7  48        PHA                             A:CE X:99 Y:88 P:E5 SP:7F
	// CEC8  A9 87     LDA #$87                        A:CE X:99 Y:88 P:E5 SP:7E
	// CECA  48        PHA                             A:87 X:99 Y:88 P:E5 SP:7E
	// CECB  A9 55     LDA #$55                        A:87 X:99 Y:88 P:E5 SP:7D
	// CECD  40        RTI                             A:55 X:99 Y:88 P:65 SP:7D
	// CECE  10 15     BPL $CEE5                       A:55 X:99 Y:88 P:A7 SP:80
	// CED0  70 13     BVS $CEE5                       A:55 X:99 Y:88 P:A7 SP:80
	// CED2  D0 11     BNE $CEE5                       A:55 X:99 Y:88 P:A7 SP:80
	// CED4  90 0F     BCC $CEE5                       A:55 X:99 Y:88 P:A7 SP:80
	// CED6  C9 55     CMP #$55                        A:55 X:99 Y:88 P:A7 SP:80
	// CED8  D0 0B     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80
	// CEDA  C0 88     CPY #$88                        A:55 X:99 Y:88 P:27 SP:80
	// CEDC  D0 07     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80
	// CEDE  E0 99     CPX #$99                        A:55 X:99 Y:88 P:27 SP:80
	// CEE0  D0 03     BNE $CEE5                       A:55 X:99 Y:88 P:27 SP:80
	// CEE2  4C E9 CE  JMP $CEE9                       A:55 X:99 Y:88 P:27 SP:80
	// CEE9  AE FF 07  LDX $07FF = FB                  A:55 X:99 Y:88 P:27 SP:80
	// CEEC  9A        TXS                             A:55 X:FB Y:88 P:A5 SP:80
	// CEED  60        RTS                             A:55 X:FB Y:88 P:A5 SP:FB
	// C60C  20 EE CE  JSR $CEEE                       A:55 X:FB Y:88 P:A5 SP:FD
	// CEEE  A2 55     LDX #$55                        A:55 X:FB Y:88 P:A5 SP:FB
	// CEF0  A0 69     LDY #$69                        A:55 X:55 Y:88 P:25 SP:FB
	// CEF2  A9 FF     LDA #$FF                        A:55 X:55 Y:69 P:25 SP:FB
	// CEF4  85 01     STA $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB
	// CEF6  EA        NOP                             A:FF X:55 Y:69 P:A5 SP:FB
	// CEF7  24 01     BIT $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB
	// CEF9  38        SEC                             A:FF X:55 Y:69 P:E5 SP:FB
	// CEFA  A9 01     LDA #$01                        A:FF X:55 Y:69 P:E5 SP:FB
	// CEFC  4A        LSR A                           A:01 X:55 Y:69 P:65 SP:FB
	// CEFD  90 1D     BCC $CF1C                       A:00 X:55 Y:69 P:67 SP:FB
	// CEFF  D0 1B     BNE $CF1C                       A:00 X:55 Y:69 P:67 SP:FB
	// CF01  30 19     BMI $CF1C                       A:00 X:55 Y:69 P:67 SP:FB
	// CF03  50 17     BVC $CF1C                       A:00 X:55 Y:69 P:67 SP:FB
	// CF05  C9 00     CMP #$00                        A:00 X:55 Y:69 P:67 SP:FB
	// CF07  D0 13     BNE $CF1C                       A:00 X:55 Y:69 P:67 SP:FB
	// CF09  B8        CLV                             A:00 X:55 Y:69 P:67 SP:FB
	// CF0A  A9 AA     LDA #$AA                        A:00 X:55 Y:69 P:27 SP:FB
	// CF0C  4A        LSR A                           A:AA X:55 Y:69 P:A5 SP:FB
	// CF0D  B0 0D     BCS $CF1C                       A:55 X:55 Y:69 P:24 SP:FB
	// CF0F  F0 0B     BEQ $CF1C                       A:55 X:55 Y:69 P:24 SP:FB
	// CF11  30 09     BMI $CF1C                       A:55 X:55 Y:69 P:24 SP:FB
	// CF13  70 07     BVS $CF1C                       A:55 X:55 Y:69 P:24 SP:FB
	// CF15  C9 55     CMP #$55                        A:55 X:55 Y:69 P:24 SP:FB
	// CF17  D0 03     BNE $CF1C                       A:55 X:55 Y:69 P:27 SP:FB
	// CF19  4C 20 CF  JMP $CF20                       A:55 X:55 Y:69 P:27 SP:FB
	// CF20  EA        NOP                             A:55 X:55 Y:69 P:27 SP:FB
	// CF21  24 01     BIT $01 = FF                    A:55 X:55 Y:69 P:27 SP:FB
	// CF23  38        SEC                             A:55 X:55 Y:69 P:E5 SP:FB
	// CF24  A9 80     LDA #$80                        A:55 X:55 Y:69 P:E5 SP:FB
	// CF26  0A        ASL A                           A:80 X:55 Y:69 P:E5 SP:FB
	// CF27  90 1E     BCC $CF47                       A:00 X:55 Y:69 P:67 SP:FB
	// CF29  D0 1C     BNE $CF47                       A:00 X:55 Y:69 P:67 SP:FB
	// CF2B  30 1A     BMI $CF47                       A:00 X:55 Y:69 P:67 SP:FB
	// CF2D  50 18     BVC $CF47                       A:00 X:55 Y:69 P:67 SP:FB
	// CF2F  C9 00     CMP #$00                        A:00 X:55 Y:69 P:67 SP:FB
	// CF31  D0 14     BNE $CF47                       A:00 X:55 Y:69 P:67 SP:FB
	// CF33  B8        CLV                             A:00 X:55 Y:69 P:67 SP:FB
	// CF34  38        SEC                             A:00 X:55 Y:69 P:27 SP:FB
	// CF35  A9 55     LDA #$55                        A:00 X:55 Y:69 P:27 SP:FB
	// CF37  0A        ASL A                           A:55 X:55 Y:69 P:25 SP:FB
	// CF38  B0 0D     BCS $CF47                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF3A  F0 0B     BEQ $CF47                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF3C  10 09     BPL $CF47                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF3E  70 07     BVS $CF47                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF40  C9 AA     CMP #$AA                        A:AA X:55 Y:69 P:A4 SP:FB
	// CF42  D0 03     BNE $CF47                       A:AA X:55 Y:69 P:27 SP:FB
	// CF44  4C 4B CF  JMP $CF4B                       A:AA X:55 Y:69 P:27 SP:FB
	// CF4B  EA        NOP                             A:AA X:55 Y:69 P:27 SP:FB
	// CF4C  24 01     BIT $01 = FF                    A:AA X:55 Y:69 P:27 SP:FB
	// CF4E  38        SEC                             A:AA X:55 Y:69 P:E5 SP:FB
	// CF4F  A9 01     LDA #$01                        A:AA X:55 Y:69 P:E5 SP:FB
	// CF51  6A        ROR A                           A:01 X:55 Y:69 P:65 SP:FB
	// CF52  90 1E     BCC $CF72                       A:80 X:55 Y:69 P:E5 SP:FB
	// CF54  F0 1C     BEQ $CF72                       A:80 X:55 Y:69 P:E5 SP:FB
	// CF56  10 1A     BPL $CF72                       A:80 X:55 Y:69 P:E5 SP:FB
	// CF58  50 18     BVC $CF72                       A:80 X:55 Y:69 P:E5 SP:FB
	// CF5A  C9 80     CMP #$80                        A:80 X:55 Y:69 P:E5 SP:FB
	// CF5C  D0 14     BNE $CF72                       A:80 X:55 Y:69 P:67 SP:FB
	// CF5E  B8        CLV                             A:80 X:55 Y:69 P:67 SP:FB
	// CF5F  18        CLC                             A:80 X:55 Y:69 P:27 SP:FB
	// CF60  A9 55     LDA #$55                        A:80 X:55 Y:69 P:26 SP:FB
	// CF62  6A        ROR A                           A:55 X:55 Y:69 P:24 SP:FB
	// CF63  90 0D     BCC $CF72                       A:2A X:55 Y:69 P:25 SP:FB
	// CF65  F0 0B     BEQ $CF72                       A:2A X:55 Y:69 P:25 SP:FB
	// CF67  30 09     BMI $CF72                       A:2A X:55 Y:69 P:25 SP:FB
	// CF69  70 07     BVS $CF72                       A:2A X:55 Y:69 P:25 SP:FB
	// CF6B  C9 2A     CMP #$2A                        A:2A X:55 Y:69 P:25 SP:FB
	// CF6D  D0 03     BNE $CF72                       A:2A X:55 Y:69 P:27 SP:FB
	// CF6F  4C 76 CF  JMP $CF76                       A:2A X:55 Y:69 P:27 SP:FB
	// CF76  EA        NOP                             A:2A X:55 Y:69 P:27 SP:FB
	// CF77  24 01     BIT $01 = FF                    A:2A X:55 Y:69 P:27 SP:FB
	// CF79  38        SEC                             A:2A X:55 Y:69 P:E5 SP:FB
	// CF7A  A9 80     LDA #$80                        A:2A X:55 Y:69 P:E5 SP:FB
	// CF7C  2A        ROL A                           A:80 X:55 Y:69 P:E5 SP:FB
	// CF7D  90 1E     BCC $CF9D                       A:01 X:55 Y:69 P:65 SP:FB
	// CF7F  F0 1C     BEQ $CF9D                       A:01 X:55 Y:69 P:65 SP:FB
	// CF81  30 1A     BMI $CF9D                       A:01 X:55 Y:69 P:65 SP:FB
	// CF83  50 18     BVC $CF9D                       A:01 X:55 Y:69 P:65 SP:FB
	// CF85  C9 01     CMP #$01                        A:01 X:55 Y:69 P:65 SP:FB
	// CF87  D0 14     BNE $CF9D                       A:01 X:55 Y:69 P:67 SP:FB
	// CF89  B8        CLV                             A:01 X:55 Y:69 P:67 SP:FB
	// CF8A  18        CLC                             A:01 X:55 Y:69 P:27 SP:FB
	// CF8B  A9 55     LDA #$55                        A:01 X:55 Y:69 P:26 SP:FB
	// CF8D  2A        ROL A                           A:55 X:55 Y:69 P:24 SP:FB
	// CF8E  B0 0D     BCS $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF90  F0 0B     BEQ $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF92  10 09     BPL $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF94  70 07     BVS $CF9D                       A:AA X:55 Y:69 P:A4 SP:FB
	// CF96  C9 AA     CMP #$AA                        A:AA X:55 Y:69 P:A4 SP:FB
	// CF98  D0 03     BNE $CF9D                       A:AA X:55 Y:69 P:27 SP:FB
	// CF9A  4C A1 CF  JMP $CFA1                       A:AA X:55 Y:69 P:27 SP:FB
	// CFA1  60        RTS                             A:AA X:55 Y:69 P:27 SP:FB
	// C60F  20 A2 CF  JSR $CFA2                       A:AA X:55 Y:69 P:27 SP:FD
	// CFA2  A5 00     LDA $00 = 00                    A:AA X:55 Y:69 P:27 SP:FB
	// CFA4  8D FF 07  STA $07FF = FB                  A:00 X:55 Y:69 P:27 SP:FB
	// CFA7  A9 00     LDA #$00                        A:00 X:55 Y:69 P:27 SP:FB
	// CFA9  85 80     STA $80 = 00                    A:00 X:55 Y:69 P:27 SP:FB
	// CFAB  A9 02     LDA #$02                        A:00 X:55 Y:69 P:27 SP:FB
	// CFAD  85 81     STA $81 = 00                    A:02 X:55 Y:69 P:25 SP:FB
	// CFAF  A9 FF     LDA #$FF                        A:02 X:55 Y:69 P:25 SP:FB
	// CFB1  85 01     STA $01 = FF                    A:FF X:55 Y:69 P:A5 SP:FB
	// CFB3  A9 00     LDA #$00                        A:FF X:55 Y:69 P:A5 SP:FB
	// CFB5  85 82     STA $82 = 00                    A:00 X:55 Y:69 P:27 SP:FB
	// CFB7  A9 03     LDA #$03                        A:00 X:55 Y:69 P:27 SP:FB
	// CFB9  85 83     STA $83 = 00                    A:03 X:55 Y:69 P:25 SP:FB
	// CFBB  85 84     STA $84 = 00                    A:03 X:55 Y:69 P:25 SP:FB
	// CFBD  A9 00     LDA #$00                        A:03 X:55 Y:69 P:25 SP:FB
	// CFBF  85 FF     STA $FF = 00                    A:00 X:55 Y:69 P:27 SP:FB
	// CFC1  A9 04     LDA #$04                        A:00 X:55 Y:69 P:27 SP:FB
	// CFC3  85 00     STA $00 = 00                    A:04 X:55 Y:69 P:25 SP:FB
	// CFC5  A9 5A     LDA #$5A                        A:04 X:55 Y:69 P:25 SP:FB
	// CFC7  8D 00 02  STA $0200 = 00                  A:5A X:55 Y:69 P:25 SP:FB
	// CFCA  A9 5B     LDA #$5B                        A:5A X:55 Y:69 P:25 SP:FB
	// CFCC  8D 00 03  STA $0300 = 00                  A:5B X:55 Y:69 P:25 SP:FB
	// CFCF  A9 5C     LDA #$5C                        A:5B X:55 Y:69 P:25 SP:FB
	// CFD1  8D 03 03  STA $0303 = 00                  A:5C X:55 Y:69 P:25 SP:FB
	// CFD4  A9 5D     LDA #$5D                        A:5C X:55 Y:69 P:25 SP:FB
	// CFD6  8D 00 04  STA $0400 = 00                  A:5D X:55 Y:69 P:25 SP:FB
	// CFD9  A2 00     LDX #$00                        A:5D X:55 Y:69 P:25 SP:FB
	// CFDB  A1 80     LDA ($80,X) @ 80 = 0200 = 5A    A:5D X:00 Y:69 P:27 SP:FB
	// CFDD  C9 5A     CMP #$5A                        A:5A X:00 Y:69 P:25 SP:FB
	// CFDF  D0 1F     BNE $D000                       A:5A X:00 Y:69 P:27 SP:FB
	// CFE1  E8        INX                             A:5A X:00 Y:69 P:27 SP:FB
	// CFE2  E8        INX                             A:5A X:01 Y:69 P:25 SP:FB
	// CFE3  A1 80     LDA ($80,X) @ 82 = 0300 = 5B    A:5A X:02 Y:69 P:25 SP:FB
	// CFE5  C9 5B     CMP #$5B                        A:5B X:02 Y:69 P:25 SP:FB
	// CFE7  D0 17     BNE $D000                       A:5B X:02 Y:69 P:27 SP:FB
	// CFE9  E8        INX                             A:5B X:02 Y:69 P:27 SP:FB
	// CFEA  A1 80     LDA ($80,X) @ 83 = 0303 = 5C    A:5B X:03 Y:69 P:25 SP:FB
	// CFEC  C9 5C     CMP #$5C                        A:5C X:03 Y:69 P:25 SP:FB
	// CFEE  D0 10     BNE $D000                       A:5C X:03 Y:69 P:27 SP:FB
	// CFF0  A2 00     LDX #$00                        A:5C X:03 Y:69 P:27 SP:FB
	// CFF2  A1 FF     LDA ($FF,X) @ FF = 0400 = 5D    A:5C X:00 Y:69 P:27 SP:FB
	// CFF4  C9 5D     CMP #$5D                        A:5D X:00 Y:69 P:25 SP:FB
	// CFF6  D0 08     BNE $D000                       A:5D X:00 Y:69 P:27 SP:FB
	// CFF8  A2 81     LDX #$81                        A:5D X:00 Y:69 P:27 SP:FB
	// CFFA  A1 FF     LDA ($FF,X) @ 80 = 0200 = 5A    A:5D X:81 Y:69 P:A5 SP:FB
	// CFFC  C9 5A     CMP #$5A                        A:5A X:81 Y:69 P:25 SP:FB
	// CFFE  F0 05     BEQ $D005                       A:5A X:81 Y:69 P:27 SP:FB
	// D005  A9 AA     LDA #$AA                        A:5A X:81 Y:69 P:27 SP:FB
	// D007  A2 00     LDX #$00                        A:AA X:81 Y:69 P:A5 SP:FB
	// D009  81 80     STA ($80,X) @ 80 = 0200 = 5A    A:AA X:00 Y:69 P:27 SP:FB
	// D00B  E8        INX                             A:AA X:00 Y:69 P:27 SP:FB
	// D00C  E8        INX                             A:AA X:01 Y:69 P:25 SP:FB
	// D00D  A9 AB     LDA #$AB                        A:AA X:02 Y:69 P:25 SP:FB
	// D00F  81 80     STA ($80,X) @ 82 = 0300 = 5B    A:AB X:02 Y:69 P:A5 SP:FB
	// D011  E8        INX                             A:AB X:02 Y:69 P:A5 SP:FB
	// D012  A9 AC     LDA #$AC                        A:AB X:03 Y:69 P:25 SP:FB
	// D014  81 80     STA ($80,X) @ 83 = 0303 = 5C    A:AC X:03 Y:69 P:A5 SP:FB
	// D016  A2 00     LDX #$00                        A:AC X:03 Y:69 P:A5 SP:FB
	// D018  A9 AD     LDA #$AD                        A:AC X:00 Y:69 P:27 SP:FB
	// D01A  81 FF     STA ($FF,X) @ FF = 0400 = 5D    A:AD X:00 Y:69 P:A5 SP:FB
	// D01C  AD 00 02  LDA $0200 = AA                  A:AD X:00 Y:69 P:A5 SP:FB
	// D01F  C9 AA     CMP #$AA                        A:AA X:00 Y:69 P:A5 SP:FB
	// D021  D0 15     BNE $D038                       A:AA X:00 Y:69 P:27 SP:FB
	// D023  AD 00 03  LDA $0300 = AB                  A:AA X:00 Y:69 P:27 SP:FB
	// D026  C9 AB     CMP #$AB                        A:AB X:00 Y:69 P:A5 SP:FB
	// D028  D0 0E     BNE $D038                       A:AB X:00 Y:69 P:27 SP:FB
	// D02A  AD 03 03  LDA $0303 = AC                  A:AB X:00 Y:69 P:27 SP:FB
	// D02D  C9 AC     CMP #$AC                        A:AC X:00 Y:69 P:A5 SP:FB
	// D02F  D0 07     BNE $D038                       A:AC X:00 Y:69 P:27 SP:FB
	// D031  AD 00 04  LDA $0400 = AD                  A:AC X:00 Y:69 P:27 SP:FB
	// D034  C9 AD     CMP #$AD                        A:AD X:00 Y:69 P:A5 SP:FB
	// D036  F0 05     BEQ $D03D                       A:AD X:00 Y:69 P:27 SP:FB
	// D03D  AD FF 07  LDA $07FF = 00                  A:AD X:00 Y:69 P:27 SP:FB
	// D040  85 00     STA $00 = 04                    A:00 X:00 Y:69 P:27 SP:FB
	// D042  A9 00     LDA #$00                        A:00 X:00 Y:69 P:27 SP:FB
	// D044  8D 00 03  STA $0300 = AB                  A:00 X:00 Y:69 P:27 SP:FB
	// D047  A9 AA     LDA #$AA                        A:00 X:00 Y:69 P:27 SP:FB
	// D049  8D 00 02  STA $0200 = AA                  A:AA X:00 Y:69 P:A5 SP:FB
	// D04C  A2 00     LDX #$00                        A:AA X:00 Y:69 P:A5 SP:FB
	// D04E  A0 5A     LDY #$5A                        A:AA X:00 Y:69 P:27 SP:FB
	// D050  20 B6 F7  JSR $F7B6                       A:AA X:00 Y:5A P:25 SP:FB
	// F7B6  18        CLC                             A:AA X:00 Y:5A P:25 SP:F9
	// F7B7  A9 FF     LDA #$FF                        A:AA X:00 Y:5A P:24 SP:F9
	// F7B9  85 01     STA $01 = FF                    A:FF X:00 Y:5A P:A4 SP:F9
	// F7BB  24 01     BIT $01 = FF                    A:FF X:00 Y:5A P:A4 SP:F9
	// F7BD  A9 55     LDA #$55                        A:FF X:00 Y:5A P:E4 SP:F9
	// F7BF  60        RTS                             A:55 X:00 Y:5A P:64 SP:F9
	// D053  01 80     ORA ($80,X) @ 80 = 0200 = AA    A:55 X:00 Y:5A P:64 SP:FB
	// D055  20 C0 F7  JSR $F7C0                       A:FF X:00 Y:5A P:E4 SP:FB
	// F7C0  B0 09     BCS $F7CB                       A:FF X:00 Y:5A P:E4 SP:F9
	// F7C2  10 07     BPL $F7CB                       A:FF X:00 Y:5A P:E4 SP:F9
	// F7C4  C9 FF     CMP #$FF                        A:FF X:00 Y:5A P:E4 SP:F9
	// F7C6  D0 03     BNE $F7CB                       A:FF X:00 Y:5A P:67 SP:F9
	// F7C8  50 01     BVC $F7CB                       A:FF X:00 Y:5A P:67 SP:F9
	// F7CA  60        RTS                             A:FF X:00 Y:5A P:67 SP:F9
	// D058  C8        INY                             A:FF X:00 Y:5A P:67 SP:FB
	// D059  20 CE F7  JSR $F7CE                       A:FF X:00 Y:5B P:65 SP:FB
	// F7CE  38        SEC                             A:FF X:00 Y:5B P:65 SP:F9
	// F7CF  B8        CLV                             A:FF X:00 Y:5B P:65 SP:F9
	// F7D0  A9 00     LDA #$00                        A:FF X:00 Y:5B P:25 SP:F9
	// F7D2  60        RTS                             A:00 X:00 Y:5B P:27 SP:F9
	// D05C  01 82     ORA ($82,X) @ 82 = 0300 = 00    A:00 X:00 Y:5B P:27 SP:FB
	// D05E  20 D3 F7  JSR $F7D3                       A:00 X:00 Y:5B P:27 SP:FB
	// F7D3  D0 07     BNE $F7DC                       A:00 X:00 Y:5B P:27 SP:F9
	// F7D5  70 05     BVS $F7DC                       A:00 X:00 Y:5B P:27 SP:F9
	// F7D7  90 03     BCC $F7DC                       A:00 X:00 Y:5B P:27 SP:F9
	// F7D9  30 01     BMI $F7DC                       A:00 X:00 Y:5B P:27 SP:F9
	// F7DB  60        RTS                             A:00 X:00 Y:5B P:27 SP:F9
	// D061  C8        INY                             A:00 X:00 Y:5B P:27 SP:FB
	// D062  20 DF F7  JSR $F7DF                       A:00 X:00 Y:5C P:25 SP:FB
	// F7DF  18        CLC                             A:00 X:00 Y:5C P:25 SP:F9
	// F7E0  24 01     BIT $01 = FF                    A:00 X:00 Y:5C P:24 SP:F9
	// F7E2  A9 55     LDA #$55                        A:00 X:00 Y:5C P:E6 SP:F9
	// F7E4  60        RTS                             A:55 X:00 Y:5C P:64 SP:F9
	// D065  21 80     AND ($80,X) @ 80 = 0200 = AA    A:55 X:00 Y:5C P:64 SP:FB
	// D067  20 E5 F7  JSR $F7E5                       A:00 X:00 Y:5C P:66 SP:FB
	// F7E5  D0 07     BNE $F7EE                       A:00 X:00 Y:5C P:66 SP:F9
	// F7E7  50 05     BVC $F7EE                       A:00 X:00 Y:5C P:66 SP:F9
	// F7E9  B0 03     BCS $F7EE                       A:00 X:00 Y:5C P:66 SP:F9
	// F7EB  30 01     BMI $F7EE                       A:00 X:00 Y:5C P:66 SP:F9
	// F7ED  60        RTS                             A:00 X:00 Y:5C P:66 SP:F9
	// D06A  C8        INY                             A:00 X:00 Y:5C P:66 SP:FB
	// D06B  A9 EF     LDA #$EF                        A:00 X:00 Y:5D P:64 SP:FB
	// D06D  8D 00 03  STA $0300 = 00                  A:EF X:00 Y:5D P:E4 SP:FB
	// D070  20 F1 F7  JSR $F7F1                       A:EF X:00 Y:5D P:E4 SP:FB
	// F7F1  38        SEC                             A:EF X:00 Y:5D P:E4 SP:F9
	// F7F2  B8        CLV                             A:EF X:00 Y:5D P:E5 SP:F9
	// F7F3  A9 F8     LDA #$F8                        A:EF X:00 Y:5D P:A5 SP:F9
	// F7F5  60        RTS                             A:F8 X:00 Y:5D P:A5 SP:F9
	// D073  21 82     AND ($82,X) @ 82 = 0300 = EF    A:F8 X:00 Y:5D P:A5 SP:FB
	// D075  20 F6 F7  JSR $F7F6                       A:E8 X:00 Y:5D P:A5 SP:FB
	// F7F6  90 09     BCC $F801                       A:E8 X:00 Y:5D P:A5 SP:F9
	// F7F8  10 07     BPL $F801                       A:E8 X:00 Y:5D P:A5 SP:F9
	// F7FA  C9 E8     CMP #$E8                        A:E8 X:00 Y:5D P:A5 SP:F9
	// F7FC  D0 03     BNE $F801                       A:E8 X:00 Y:5D P:27 SP:F9
	// F7FE  70 01     BVS $F801                       A:E8 X:00 Y:5D P:27 SP:F9
	// F800  60        RTS                             A:E8 X:00 Y:5D P:27 SP:F9
	// D078  C8        INY                             A:E8 X:00 Y:5D P:27 SP:FB
	// D079  20 04 F8  JSR $F804                       A:E8 X:00 Y:5E P:25 SP:FB
	// F804  18        CLC                             A:E8 X:00 Y:5E P:25 SP:F9
	// F805  24 01     BIT $01 = FF                    A:E8 X:00 Y:5E P:24 SP:F9
	// F807  A9 5F     LDA #$5F                        A:E8 X:00 Y:5E P:E4 SP:F9
	// F809  60        RTS                             A:5F X:00 Y:5E P:64 SP:F9
	// D07C  41 80     EOR ($80,X) @ 80 = 0200 = AA    A:5F X:00 Y:5E P:64 SP:FB
	// D07E  20 0A F8  JSR $F80A                       A:F5 X:00 Y:5E P:E4 SP:FB
	// F80A  B0 09     BCS $F815                       A:F5 X:00 Y:5E P:E4 SP:F9
	// F80C  10 07     BPL $F815                       A:F5 X:00 Y:5E P:E4 SP:F9
	// F80E  C9 F5     CMP #$F5                        A:F5 X:00 Y:5E P:E4 SP:F9
	// F810  D0 03     BNE $F815                       A:F5 X:00 Y:5E P:67 SP:F9
	// F812  50 01     BVC $F815                       A:F5 X:00 Y:5E P:67 SP:F9
	// F814  60        RTS                             A:F5 X:00 Y:5E P:67 SP:F9
	// D081  C8        INY                             A:F5 X:00 Y:5E P:67 SP:FB
	// D082  A9 70     LDA #$70                        A:F5 X:00 Y:5F P:65 SP:FB
	// D084  8D 00 03  STA $0300 = EF                  A:70 X:00 Y:5F P:65 SP:FB
	// D087  20 18 F8  JSR $F818                       A:70 X:00 Y:5F P:65 SP:FB
	// F818  38        SEC                             A:70 X:00 Y:5F P:65 SP:F9
	// F819  B8        CLV                             A:70 X:00 Y:5F P:65 SP:F9
	// F81A  A9 70     LDA #$70                        A:70 X:00 Y:5F P:25 SP:F9
	// F81C  60        RTS                             A:70 X:00 Y:5F P:25 SP:F9
	// D08A  41 82     EOR ($82,X) @ 82 = 0300 = 70    A:70 X:00 Y:5F P:25 SP:FB
	// D08C  20 1D F8  JSR $F81D                       A:00 X:00 Y:5F P:27 SP:FB
	// F81D  D0 07     BNE $F826                       A:00 X:00 Y:5F P:27 SP:F9
	// F81F  70 05     BVS $F826                       A:00 X:00 Y:5F P:27 SP:F9
	// F821  90 03     BCC $F826                       A:00 X:00 Y:5F P:27 SP:F9
	// F823  30 01     BMI $F826                       A:00 X:00 Y:5F P:27 SP:F9
	// F825  60        RTS                             A:00 X:00 Y:5F P:27 SP:F9
	// D08F  C8        INY                             A:00 X:00 Y:5F P:27 SP:FB
	// D090  A9 69     LDA #$69                        A:00 X:00 Y:60 P:25 SP:FB
	// D092  8D 00 02  STA $0200 = AA                  A:69 X:00 Y:60 P:25 SP:FB
	// D095  20 29 F8  JSR $F829                       A:69 X:00 Y:60 P:25 SP:FB
	// F829  18        CLC                             A:69 X:00 Y:60 P:25 SP:F9
	// F82A  24 01     BIT $01 = FF                    A:69 X:00 Y:60 P:24 SP:F9
	// F82C  A9 00     LDA #$00                        A:69 X:00 Y:60 P:E4 SP:F9
	// F82E  60        RTS                             A:00 X:00 Y:60 P:66 SP:F9
	// D098  61 80     ADC ($80,X) @ 80 = 0200 = 69    A:00 X:00 Y:60 P:66 SP:FB
	// D09A  20 2F F8  JSR $F82F                       A:69 X:00 Y:60 P:24 SP:FB
	// F82F  30 09     BMI $F83A                       A:69 X:00 Y:60 P:24 SP:F9
	// F831  B0 07     BCS $F83A                       A:69 X:00 Y:60 P:24 SP:F9
	// F833  C9 69     CMP #$69                        A:69 X:00 Y:60 P:24 SP:F9
	// F835  D0 03     BNE $F83A                       A:69 X:00 Y:60 P:27 SP:F9
	// F837  70 01     BVS $F83A                       A:69 X:00 Y:60 P:27 SP:F9
	// F839  60        RTS                             A:69 X:00 Y:60 P:27 SP:F9
	// D09D  C8        INY                             A:69 X:00 Y:60 P:27 SP:FB
	// D09E  20 3D F8  JSR $F83D                       A:69 X:00 Y:61 P:25 SP:FB
	// F83D  38        SEC                             A:69 X:00 Y:61 P:25 SP:F9
	// F83E  24 01     BIT $01 = FF                    A:69 X:00 Y:61 P:25 SP:F9
	// F840  A9 00     LDA #$00                        A:69 X:00 Y:61 P:E5 SP:F9
	// F842  60        RTS                             A:00 X:00 Y:61 P:67 SP:F9
	// D0A1  61 80     ADC ($80,X) @ 80 = 0200 = 69    A:00 X:00 Y:61 P:67 SP:FB
	// D0A3  20 43 F8  JSR $F843                       A:6A X:00 Y:61 P:24 SP:FB
	// F843  30 09     BMI $F84E                       A:6A X:00 Y:61 P:24 SP:F9
	// F845  B0 07     BCS $F84E                       A:6A X:00 Y:61 P:24 SP:F9
	// F847  C9 6A     CMP #$6A                        A:6A X:00 Y:61 P:24 SP:F9
	// F849  D0 03     BNE $F84E                       A:6A X:00 Y:61 P:27 SP:F9
	// F84B  70 01     BVS $F84E                       A:6A X:00 Y:61 P:27 SP:F9
	// F84D  60        RTS                             A:6A X:00 Y:61 P:27 SP:F9
	// D0A6  C8        INY                             A:6A X:00 Y:61 P:27 SP:FB
	// D0A7  A9 7F     LDA #$7F                        A:6A X:00 Y:62 P:25 SP:FB
	// D0A9  8D 00 02  STA $0200 = 69                  A:7F X:00 Y:62 P:25 SP:FB
	// D0AC  20 51 F8  JSR $F851                       A:7F X:00 Y:62 P:25 SP:FB
	// F851  38        SEC                             A:7F X:00 Y:62 P:25 SP:F9
	// F852  B8        CLV                             A:7F X:00 Y:62 P:25 SP:F9
	// F853  A9 7F     LDA #$7F                        A:7F X:00 Y:62 P:25 SP:F9
	// F855  60        RTS                             A:7F X:00 Y:62 P:25 SP:F9
	// D0AF  61 80     ADC ($80,X) @ 80 = 0200 = 7F    A:7F X:00 Y:62 P:25 SP:FB
	// D0B1  20 56 F8  JSR $F856                       A:FF X:00 Y:62 P:E4 SP:FB
	// F856  10 09     BPL $F861                       A:FF X:00 Y:62 P:E4 SP:F9
	// F858  B0 07     BCS $F861                       A:FF X:00 Y:62 P:E4 SP:F9
	// F85A  C9 FF     CMP #$FF                        A:FF X:00 Y:62 P:E4 SP:F9
	// F85C  D0 03     BNE $F861                       A:FF X:00 Y:62 P:67 SP:F9
	// F85E  50 01     BVC $F861                       A:FF X:00 Y:62 P:67 SP:F9
	// F860  60        RTS                             A:FF X:00 Y:62 P:67 SP:F9
	// D0B4  C8        INY                             A:FF X:00 Y:62 P:67 SP:FB
	// D0B5  A9 80     LDA #$80                        A:FF X:00 Y:63 P:65 SP:FB
	// D0B7  8D 00 02  STA $0200 = 7F                  A:80 X:00 Y:63 P:E5 SP:FB
	// D0BA  20 64 F8  JSR $F864                       A:80 X:00 Y:63 P:E5 SP:FB
	// F864  18        CLC                             A:80 X:00 Y:63 P:E5 SP:F9
	// F865  24 01     BIT $01 = FF                    A:80 X:00 Y:63 P:E4 SP:F9
	// F867  A9 7F     LDA #$7F                        A:80 X:00 Y:63 P:E4 SP:F9
	// F869  60        RTS                             A:7F X:00 Y:63 P:64 SP:F9
	// D0BD  61 80     ADC ($80,X) @ 80 = 0200 = 80    A:7F X:00 Y:63 P:64 SP:FB
	// D0BF  20 6A F8  JSR $F86A                       A:FF X:00 Y:63 P:A4 SP:FB
	// F86A  10 09     BPL $F875                       A:FF X:00 Y:63 P:A4 SP:F9
	// F86C  B0 07     BCS $F875                       A:FF X:00 Y:63 P:A4 SP:F9
	// F86E  C9 FF     CMP #$FF                        A:FF X:00 Y:63 P:A4 SP:F9
	// F870  D0 03     BNE $F875                       A:FF X:00 Y:63 P:27 SP:F9
	// F872  70 01     BVS $F875                       A:FF X:00 Y:63 P:27 SP:F9
	// F874  60        RTS                             A:FF X:00 Y:63 P:27 SP:F9
	// D0C2  C8        INY                             A:FF X:00 Y:63 P:27 SP:FB
	// D0C3  20 78 F8  JSR $F878                       A:FF X:00 Y:64 P:25 SP:FB
	// F878  38        SEC                             A:FF X:00 Y:64 P:25 SP:F9
	// F879  B8        CLV                             A:FF X:00 Y:64 P:25 SP:F9
	// F87A  A9 7F     LDA #$7F                        A:FF X:00 Y:64 P:25 SP:F9
	// F87C  60        RTS                             A:7F X:00 Y:64 P:25 SP:F9
	// D0C6  61 80     ADC ($80,X) @ 80 = 0200 = 80    A:7F X:00 Y:64 P:25 SP:FB
	// D0C8  20 7D F8  JSR $F87D                       A:00 X:00 Y:64 P:27 SP:FB
	// F87D  D0 07     BNE $F886                       A:00 X:00 Y:64 P:27 SP:F9
	// F87F  30 05     BMI $F886                       A:00 X:00 Y:64 P:27 SP:F9
	// F881  70 03     BVS $F886                       A:00 X:00 Y:64 P:27 SP:F9
	// F883  90 01     BCC $F886                       A:00 X:00 Y:64 P:27 SP:F9
	// F885  60        RTS                             A:00 X:00 Y:64 P:27 SP:F9
	// D0CB  C8        INY                             A:00 X:00 Y:64 P:27 SP:FB
	// D0CC  A9 40     LDA #$40                        A:00 X:00 Y:65 P:25 SP:FB
	// D0CE  8D 00 02  STA $0200 = 80                  A:40 X:00 Y:65 P:25 SP:FB
	// D0D1  20 89 F8  JSR $F889                       A:40 X:00 Y:65 P:25 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:00 Y:65 P:25 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:00 Y:65 P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:00 Y:65 P:65 SP:F9
	// D0D4  C1 80     CMP ($80,X) @ 80 = 0200 = 40    A:40 X:00 Y:65 P:65 SP:FB
	// D0D6  20 8E F8  JSR $F88E                       A:40 X:00 Y:65 P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:00 Y:65 P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:00 Y:65 P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:00 Y:65 P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:00 Y:65 P:67 SP:F9
	// F896  60        RTS                             A:40 X:00 Y:65 P:67 SP:F9
	// D0D9  C8        INY                             A:40 X:00 Y:65 P:67 SP:FB
	// D0DA  48        PHA                             A:40 X:00 Y:66 P:65 SP:FB
	// D0DB  A9 3F     LDA #$3F                        A:40 X:00 Y:66 P:65 SP:FA
	// D0DD  8D 00 02  STA $0200 = 40                  A:3F X:00 Y:66 P:65 SP:FA
	// D0E0  68        PLA                             A:3F X:00 Y:66 P:65 SP:FA
	// D0E1  20 9A F8  JSR $F89A                       A:40 X:00 Y:66 P:65 SP:FB
	// F89A  B8        CLV                             A:40 X:00 Y:66 P:65 SP:F9
	// F89B  60        RTS                             A:40 X:00 Y:66 P:25 SP:F9
	// D0E4  C1 80     CMP ($80,X) @ 80 = 0200 = 3F    A:40 X:00 Y:66 P:25 SP:FB
	// D0E6  20 9C F8  JSR $F89C                       A:40 X:00 Y:66 P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:40 X:00 Y:66 P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:40 X:00 Y:66 P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:40 X:00 Y:66 P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:40 X:00 Y:66 P:25 SP:F9
	// F8A4  60        RTS                             A:40 X:00 Y:66 P:25 SP:F9
	// D0E9  C8        INY                             A:40 X:00 Y:66 P:25 SP:FB
	// D0EA  48        PHA                             A:40 X:00 Y:67 P:25 SP:FB
	// D0EB  A9 41     LDA #$41                        A:40 X:00 Y:67 P:25 SP:FA
	// D0ED  8D 00 02  STA $0200 = 3F                  A:41 X:00 Y:67 P:25 SP:FA
	// D0F0  68        PLA                             A:41 X:00 Y:67 P:25 SP:FA
	// D0F1  C1 80     CMP ($80,X) @ 80 = 0200 = 41    A:40 X:00 Y:67 P:25 SP:FB
	// D0F3  20 A8 F8  JSR $F8A8                       A:40 X:00 Y:67 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:40 X:00 Y:67 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:40 X:00 Y:67 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:40 X:00 Y:67 P:A4 SP:F9
	// F8AE  60        RTS                             A:40 X:00 Y:67 P:A4 SP:F9
	// D0F6  C8        INY                             A:40 X:00 Y:67 P:A4 SP:FB
	// D0F7  48        PHA                             A:40 X:00 Y:68 P:24 SP:FB
	// D0F8  A9 00     LDA #$00                        A:40 X:00 Y:68 P:24 SP:FA
	// D0FA  8D 00 02  STA $0200 = 41                  A:00 X:00 Y:68 P:26 SP:FA
	// D0FD  68        PLA                             A:00 X:00 Y:68 P:26 SP:FA
	// D0FE  20 B2 F8  JSR $F8B2                       A:40 X:00 Y:68 P:24 SP:FB
	// F8B2  A9 80     LDA #$80                        A:40 X:00 Y:68 P:24 SP:F9
	// F8B4  60        RTS                             A:80 X:00 Y:68 P:A4 SP:F9
	// D101  C1 80     CMP ($80,X) @ 80 = 0200 = 00    A:80 X:00 Y:68 P:A4 SP:FB
	// D103  20 B5 F8  JSR $F8B5                       A:80 X:00 Y:68 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:00 Y:68 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:00 Y:68 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:00 Y:68 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:00 Y:68 P:A5 SP:F9
	// D106  C8        INY                             A:80 X:00 Y:68 P:A5 SP:FB
	// D107  48        PHA                             A:80 X:00 Y:69 P:25 SP:FB
	// D108  A9 80     LDA #$80                        A:80 X:00 Y:69 P:25 SP:FA
	// D10A  8D 00 02  STA $0200 = 00                  A:80 X:00 Y:69 P:A5 SP:FA
	// D10D  68        PLA                             A:80 X:00 Y:69 P:A5 SP:FA
	// D10E  C1 80     CMP ($80,X) @ 80 = 0200 = 80    A:80 X:00 Y:69 P:A5 SP:FB
	// D110  20 BF F8  JSR $F8BF                       A:80 X:00 Y:69 P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:00 Y:69 P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:00 Y:69 P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:00 Y:69 P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:00 Y:69 P:27 SP:F9
	// D113  C8        INY                             A:80 X:00 Y:69 P:27 SP:FB
	// D114  48        PHA                             A:80 X:00 Y:6A P:25 SP:FB
	// D115  A9 81     LDA #$81                        A:80 X:00 Y:6A P:25 SP:FA
	// D117  8D 00 02  STA $0200 = 80                  A:81 X:00 Y:6A P:A5 SP:FA
	// D11A  68        PLA                             A:81 X:00 Y:6A P:A5 SP:FA
	// D11B  C1 80     CMP ($80,X) @ 80 = 0200 = 81    A:80 X:00 Y:6A P:A5 SP:FB
	// D11D  20 C9 F8  JSR $F8C9                       A:80 X:00 Y:6A P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:80 X:00 Y:6A P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:80 X:00 Y:6A P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:80 X:00 Y:6A P:A4 SP:F9
	// F8CF  60        RTS                             A:80 X:00 Y:6A P:A4 SP:F9
	// D120  C8        INY                             A:80 X:00 Y:6A P:A4 SP:FB
	// D121  48        PHA                             A:80 X:00 Y:6B P:24 SP:FB
	// D122  A9 7F     LDA #$7F                        A:80 X:00 Y:6B P:24 SP:FA
	// D124  8D 00 02  STA $0200 = 81                  A:7F X:00 Y:6B P:24 SP:FA
	// D127  68        PLA                             A:7F X:00 Y:6B P:24 SP:FA
	// D128  C1 80     CMP ($80,X) @ 80 = 0200 = 7F    A:80 X:00 Y:6B P:A4 SP:FB
	// D12A  20 D3 F8  JSR $F8D3                       A:80 X:00 Y:6B P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:80 X:00 Y:6B P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:80 X:00 Y:6B P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:80 X:00 Y:6B P:25 SP:F9
	// F8D9  60        RTS                             A:80 X:00 Y:6B P:25 SP:F9
	// D12D  C8        INY                             A:80 X:00 Y:6B P:25 SP:FB
	// D12E  A9 40     LDA #$40                        A:80 X:00 Y:6C P:25 SP:FB
	// D130  8D 00 02  STA $0200 = 7F                  A:40 X:00 Y:6C P:25 SP:FB
	// D133  20 31 F9  JSR $F931                       A:40 X:00 Y:6C P:25 SP:FB
	// F931  24 01     BIT $01 = FF                    A:40 X:00 Y:6C P:25 SP:F9
	// F933  A9 40     LDA #$40                        A:40 X:00 Y:6C P:E5 SP:F9
	// F935  38        SEC                             A:40 X:00 Y:6C P:65 SP:F9
	// F936  60        RTS                             A:40 X:00 Y:6C P:65 SP:F9
	// D136  E1 80     SBC ($80,X) @ 80 = 0200 = 40    A:40 X:00 Y:6C P:65 SP:FB
	// D138  20 37 F9  JSR $F937                       A:00 X:00 Y:6C P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:00 Y:6C P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:00 Y:6C P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:00 Y:6C P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:00 Y:6C P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:00 Y:6C P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:00 Y:6C P:27 SP:F9
	// F943  60        RTS                             A:00 X:00 Y:6C P:27 SP:F9
	// D13B  C8        INY                             A:00 X:00 Y:6C P:27 SP:FB
	// D13C  A9 3F     LDA #$3F                        A:00 X:00 Y:6D P:25 SP:FB
	// D13E  8D 00 02  STA $0200 = 40                  A:3F X:00 Y:6D P:25 SP:FB
	// D141  20 47 F9  JSR $F947                       A:3F X:00 Y:6D P:25 SP:FB
	// F947  B8        CLV                             A:3F X:00 Y:6D P:25 SP:F9
	// F948  38        SEC                             A:3F X:00 Y:6D P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:3F X:00 Y:6D P:25 SP:F9
	// F94B  60        RTS                             A:40 X:00 Y:6D P:25 SP:F9
	// D144  E1 80     SBC ($80,X) @ 80 = 0200 = 3F    A:40 X:00 Y:6D P:25 SP:FB
	// D146  20 4C F9  JSR $F94C                       A:01 X:00 Y:6D P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:00 Y:6D P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:00 Y:6D P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:00 Y:6D P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:00 Y:6D P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:00 Y:6D P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:00 Y:6D P:27 SP:F9
	// F958  60        RTS                             A:01 X:00 Y:6D P:27 SP:F9
	// D149  C8        INY                             A:01 X:00 Y:6D P:27 SP:FB
	// D14A  A9 41     LDA #$41                        A:01 X:00 Y:6E P:25 SP:FB
	// D14C  8D 00 02  STA $0200 = 3F                  A:41 X:00 Y:6E P:25 SP:FB
	// D14F  20 5C F9  JSR $F95C                       A:41 X:00 Y:6E P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:41 X:00 Y:6E P:25 SP:F9
	// F95E  38        SEC                             A:40 X:00 Y:6E P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:00 Y:6E P:25 SP:F9
	// F961  60        RTS                             A:40 X:00 Y:6E P:E5 SP:F9
	// D152  E1 80     SBC ($80,X) @ 80 = 0200 = 41    A:40 X:00 Y:6E P:E5 SP:FB
	// D154  20 62 F9  JSR $F962                       A:FF X:00 Y:6E P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:00 Y:6E P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:00 Y:6E P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:00 Y:6E P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:00 Y:6E P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:00 Y:6E P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:00 Y:6E P:27 SP:F9
	// F96E  60        RTS                             A:FF X:00 Y:6E P:27 SP:F9
	// D157  C8        INY                             A:FF X:00 Y:6E P:27 SP:FB
	// D158  A9 00     LDA #$00                        A:FF X:00 Y:6F P:25 SP:FB
	// D15A  8D 00 02  STA $0200 = 41                  A:00 X:00 Y:6F P:27 SP:FB
	// D15D  20 72 F9  JSR $F972                       A:00 X:00 Y:6F P:27 SP:FB
	// F972  18        CLC                             A:00 X:00 Y:6F P:27 SP:F9
	// F973  A9 80     LDA #$80                        A:00 X:00 Y:6F P:26 SP:F9
	// F975  60        RTS                             A:80 X:00 Y:6F P:A4 SP:F9
	// D160  E1 80     SBC ($80,X) @ 80 = 0200 = 00    A:80 X:00 Y:6F P:A4 SP:FB
	// D162  20 76 F9  JSR $F976                       A:7F X:00 Y:6F P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:00 Y:6F P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:00 Y:6F P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:00 Y:6F P:67 SP:F9
	// F97C  60        RTS                             A:7F X:00 Y:6F P:67 SP:F9
	// D165  C8        INY                             A:7F X:00 Y:6F P:67 SP:FB
	// D166  A9 7F     LDA #$7F                        A:7F X:00 Y:70 P:65 SP:FB
	// D168  8D 00 02  STA $0200 = 00                  A:7F X:00 Y:70 P:65 SP:FB
	// D16B  20 80 F9  JSR $F980                       A:7F X:00 Y:70 P:65 SP:FB
	// F980  38        SEC                             A:7F X:00 Y:70 P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:00 Y:70 P:65 SP:F9
	// F983  60        RTS                             A:81 X:00 Y:70 P:E5 SP:F9
	// D16E  E1 80     SBC ($80,X) @ 80 = 0200 = 7F    A:81 X:00 Y:70 P:E5 SP:FB
	// D170  20 84 F9  JSR $F984                       A:02 X:00 Y:70 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:00 Y:70 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:00 Y:70 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:00 Y:70 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:00 Y:70 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:00 Y:70 P:67 SP:F9
	// D173  60        RTS                             A:02 X:00 Y:70 P:67 SP:FB
	// C612  20 74 D1  JSR $D174                       A:02 X:00 Y:70 P:67 SP:FD
	// D174  A9 55     LDA #$55                        A:02 X:00 Y:70 P:67 SP:FB
	// D176  85 78     STA $78 = 00                    A:55 X:00 Y:70 P:65 SP:FB
	// D178  A9 FF     LDA #$FF                        A:55 X:00 Y:70 P:65 SP:FB
	// D17A  85 01     STA $01 = FF                    A:FF X:00 Y:70 P:E5 SP:FB
	// D17C  24 01     BIT $01 = FF                    A:FF X:00 Y:70 P:E5 SP:FB
	// D17E  A0 11     LDY #$11                        A:FF X:00 Y:70 P:E5 SP:FB
	// D180  A2 23     LDX #$23                        A:FF X:00 Y:11 P:65 SP:FB
	// D182  A9 00     LDA #$00                        A:FF X:23 Y:11 P:65 SP:FB
	// D184  A5 78     LDA $78 = 55                    A:00 X:23 Y:11 P:67 SP:FB
	// D186  F0 10     BEQ $D198                       A:55 X:23 Y:11 P:65 SP:FB
	// D188  30 0E     BMI $D198                       A:55 X:23 Y:11 P:65 SP:FB
	// D18A  C9 55     CMP #$55                        A:55 X:23 Y:11 P:65 SP:FB
	// D18C  D0 0A     BNE $D198                       A:55 X:23 Y:11 P:67 SP:FB
	// D18E  C0 11     CPY #$11                        A:55 X:23 Y:11 P:67 SP:FB
	// D190  D0 06     BNE $D198                       A:55 X:23 Y:11 P:67 SP:FB
	// D192  E0 23     CPX #$23                        A:55 X:23 Y:11 P:67 SP:FB
	// D194  50 02     BVC $D198                       A:55 X:23 Y:11 P:67 SP:FB
	// D196  F0 04     BEQ $D19C                       A:55 X:23 Y:11 P:67 SP:FB
	// D19C  A9 46     LDA #$46                        A:55 X:23 Y:11 P:67 SP:FB
	// D19E  24 01     BIT $01 = FF                    A:46 X:23 Y:11 P:65 SP:FB
	// D1A0  85 78     STA $78 = 55                    A:46 X:23 Y:11 P:E5 SP:FB
	// D1A2  F0 0A     BEQ $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB
	// D1A4  10 08     BPL $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB
	// D1A6  50 06     BVC $D1AE                       A:46 X:23 Y:11 P:E5 SP:FB
	// D1A8  A5 78     LDA $78 = 46                    A:46 X:23 Y:11 P:E5 SP:FB
	// D1AA  C9 46     CMP #$46                        A:46 X:23 Y:11 P:65 SP:FB
	// D1AC  F0 04     BEQ $D1B2                       A:46 X:23 Y:11 P:67 SP:FB
	// D1B2  A9 55     LDA #$55                        A:46 X:23 Y:11 P:67 SP:FB
	// D1B4  85 78     STA $78 = 46                    A:55 X:23 Y:11 P:65 SP:FB
	// D1B6  24 01     BIT $01 = FF                    A:55 X:23 Y:11 P:65 SP:FB
	// D1B8  A9 11     LDA #$11                        A:55 X:23 Y:11 P:E5 SP:FB
	// D1BA  A2 23     LDX #$23                        A:11 X:23 Y:11 P:65 SP:FB
	// D1BC  A0 00     LDY #$00                        A:11 X:23 Y:11 P:65 SP:FB
	// D1BE  A4 78     LDY $78 = 55                    A:11 X:23 Y:00 P:67 SP:FB
	// D1C0  F0 10     BEQ $D1D2                       A:11 X:23 Y:55 P:65 SP:FB
	// D1C2  30 0E     BMI $D1D2                       A:11 X:23 Y:55 P:65 SP:FB
	// D1C4  C0 55     CPY #$55                        A:11 X:23 Y:55 P:65 SP:FB
	// D1C6  D0 0A     BNE $D1D2                       A:11 X:23 Y:55 P:67 SP:FB
	// D1C8  C9 11     CMP #$11                        A:11 X:23 Y:55 P:67 SP:FB
	// D1CA  D0 06     BNE $D1D2                       A:11 X:23 Y:55 P:67 SP:FB
	// D1CC  E0 23     CPX #$23                        A:11 X:23 Y:55 P:67 SP:FB
	// D1CE  50 02     BVC $D1D2                       A:11 X:23 Y:55 P:67 SP:FB
	// D1D0  F0 04     BEQ $D1D6                       A:11 X:23 Y:55 P:67 SP:FB
	// D1D6  A0 46     LDY #$46                        A:11 X:23 Y:55 P:67 SP:FB
	// D1D8  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:65 SP:FB
	// D1DA  84 78     STY $78 = 55                    A:11 X:23 Y:46 P:E5 SP:FB
	// D1DC  F0 0A     BEQ $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB
	// D1DE  10 08     BPL $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB
	// D1E0  50 06     BVC $D1E8                       A:11 X:23 Y:46 P:E5 SP:FB
	// D1E2  A4 78     LDY $78 = 46                    A:11 X:23 Y:46 P:E5 SP:FB
	// D1E4  C0 46     CPY #$46                        A:11 X:23 Y:46 P:65 SP:FB
	// D1E6  F0 04     BEQ $D1EC                       A:11 X:23 Y:46 P:67 SP:FB
	// D1EC  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:67 SP:FB
	// D1EE  A9 55     LDA #$55                        A:11 X:23 Y:46 P:E5 SP:FB
	// D1F0  85 78     STA $78 = 46                    A:55 X:23 Y:46 P:65 SP:FB
	// D1F2  A0 11     LDY #$11                        A:55 X:23 Y:46 P:65 SP:FB
	// D1F4  A9 23     LDA #$23                        A:55 X:23 Y:11 P:65 SP:FB
	// D1F6  A2 00     LDX #$00                        A:23 X:23 Y:11 P:65 SP:FB
	// D1F8  A6 78     LDX $78 = 55                    A:23 X:00 Y:11 P:67 SP:FB
	// D1FA  F0 10     BEQ $D20C                       A:23 X:55 Y:11 P:65 SP:FB
	// D1FC  30 0E     BMI $D20C                       A:23 X:55 Y:11 P:65 SP:FB
	// D1FE  E0 55     CPX #$55                        A:23 X:55 Y:11 P:65 SP:FB
	// D200  D0 0A     BNE $D20C                       A:23 X:55 Y:11 P:67 SP:FB
	// D202  C0 11     CPY #$11                        A:23 X:55 Y:11 P:67 SP:FB
	// D204  D0 06     BNE $D20C                       A:23 X:55 Y:11 P:67 SP:FB
	// D206  C9 23     CMP #$23                        A:23 X:55 Y:11 P:67 SP:FB
	// D208  50 02     BVC $D20C                       A:23 X:55 Y:11 P:67 SP:FB
	// D20A  F0 04     BEQ $D210                       A:23 X:55 Y:11 P:67 SP:FB
	// D210  A2 46     LDX #$46                        A:23 X:55 Y:11 P:67 SP:FB
	// D212  24 01     BIT $01 = FF                    A:23 X:46 Y:11 P:65 SP:FB
	// D214  86 78     STX $78 = 55                    A:23 X:46 Y:11 P:E5 SP:FB
	// D216  F0 0A     BEQ $D222                       A:23 X:46 Y:11 P:E5 SP:FB
	// D218  10 08     BPL $D222                       A:23 X:46 Y:11 P:E5 SP:FB
	// D21A  50 06     BVC $D222                       A:23 X:46 Y:11 P:E5 SP:FB
	// D21C  A6 78     LDX $78 = 46                    A:23 X:46 Y:11 P:E5 SP:FB
	// D21E  E0 46     CPX #$46                        A:23 X:46 Y:11 P:65 SP:FB
	// D220  F0 04     BEQ $D226                       A:23 X:46 Y:11 P:67 SP:FB
	// D226  A9 C0     LDA #$C0                        A:23 X:46 Y:11 P:67 SP:FB
	// D228  85 78     STA $78 = 46                    A:C0 X:46 Y:11 P:E5 SP:FB
	// D22A  A2 33     LDX #$33                        A:C0 X:46 Y:11 P:E5 SP:FB
	// D22C  A0 88     LDY #$88                        A:C0 X:33 Y:11 P:65 SP:FB
	// D22E  A9 05     LDA #$05                        A:C0 X:33 Y:88 P:E5 SP:FB
	// D230  24 78     BIT $78 = C0                    A:05 X:33 Y:88 P:65 SP:FB
	// D232  10 10     BPL $D244                       A:05 X:33 Y:88 P:E7 SP:FB
	// D234  50 0E     BVC $D244                       A:05 X:33 Y:88 P:E7 SP:FB
	// D236  D0 0C     BNE $D244                       A:05 X:33 Y:88 P:E7 SP:FB
	// D238  C9 05     CMP #$05                        A:05 X:33 Y:88 P:E7 SP:FB
	// D23A  D0 08     BNE $D244                       A:05 X:33 Y:88 P:67 SP:FB
	// D23C  E0 33     CPX #$33                        A:05 X:33 Y:88 P:67 SP:FB
	// D23E  D0 04     BNE $D244                       A:05 X:33 Y:88 P:67 SP:FB
	// D240  C0 88     CPY #$88                        A:05 X:33 Y:88 P:67 SP:FB
	// D242  F0 04     BEQ $D248                       A:05 X:33 Y:88 P:67 SP:FB
	// D248  A9 03     LDA #$03                        A:05 X:33 Y:88 P:67 SP:FB
	// D24A  85 78     STA $78 = C0                    A:03 X:33 Y:88 P:65 SP:FB
	// D24C  A9 01     LDA #$01                        A:03 X:33 Y:88 P:65 SP:FB
	// D24E  24 78     BIT $78 = 03                    A:01 X:33 Y:88 P:65 SP:FB
	// D250  30 08     BMI $D25A                       A:01 X:33 Y:88 P:25 SP:FB
	// D252  70 06     BVS $D25A                       A:01 X:33 Y:88 P:25 SP:FB
	// D254  F0 04     BEQ $D25A                       A:01 X:33 Y:88 P:25 SP:FB
	// D256  C9 01     CMP #$01                        A:01 X:33 Y:88 P:25 SP:FB
	// D258  F0 04     BEQ $D25E                       A:01 X:33 Y:88 P:27 SP:FB
	// D25E  A0 7E     LDY #$7E                        A:01 X:33 Y:88 P:27 SP:FB
	// D260  A9 AA     LDA #$AA                        A:01 X:33 Y:7E P:25 SP:FB
	// D262  85 78     STA $78 = 03                    A:AA X:33 Y:7E P:A5 SP:FB
	// D264  20 B6 F7  JSR $F7B6                       A:AA X:33 Y:7E P:A5 SP:FB
	// F7B6  18        CLC                             A:AA X:33 Y:7E P:A5 SP:F9
	// F7B7  A9 FF     LDA #$FF                        A:AA X:33 Y:7E P:A4 SP:F9
	// F7B9  85 01     STA $01 = FF                    A:FF X:33 Y:7E P:A4 SP:F9
	// F7BB  24 01     BIT $01 = FF                    A:FF X:33 Y:7E P:A4 SP:F9
	// F7BD  A9 55     LDA #$55                        A:FF X:33 Y:7E P:E4 SP:F9
	// F7BF  60        RTS                             A:55 X:33 Y:7E P:64 SP:F9
	// D267  05 78     ORA $78 = AA                    A:55 X:33 Y:7E P:64 SP:FB
	// D269  20 C0 F7  JSR $F7C0                       A:FF X:33 Y:7E P:E4 SP:FB
	// F7C0  B0 09     BCS $F7CB                       A:FF X:33 Y:7E P:E4 SP:F9
	// F7C2  10 07     BPL $F7CB                       A:FF X:33 Y:7E P:E4 SP:F9
	// F7C4  C9 FF     CMP #$FF                        A:FF X:33 Y:7E P:E4 SP:F9
	// F7C6  D0 03     BNE $F7CB                       A:FF X:33 Y:7E P:67 SP:F9
	// F7C8  50 01     BVC $F7CB                       A:FF X:33 Y:7E P:67 SP:F9
	// F7CA  60        RTS                             A:FF X:33 Y:7E P:67 SP:F9
	// D26C  C8        INY                             A:FF X:33 Y:7E P:67 SP:FB
	// D26D  A9 00     LDA #$00                        A:FF X:33 Y:7F P:65 SP:FB
	// D26F  85 78     STA $78 = AA                    A:00 X:33 Y:7F P:67 SP:FB
	// D271  20 CE F7  JSR $F7CE                       A:00 X:33 Y:7F P:67 SP:FB
	// F7CE  38        SEC                             A:00 X:33 Y:7F P:67 SP:F9
	// F7CF  B8        CLV                             A:00 X:33 Y:7F P:67 SP:F9
	// F7D0  A9 00     LDA #$00                        A:00 X:33 Y:7F P:27 SP:F9
	// F7D2  60        RTS                             A:00 X:33 Y:7F P:27 SP:F9
	// D274  05 78     ORA $78 = 00                    A:00 X:33 Y:7F P:27 SP:FB
	// D276  20 D3 F7  JSR $F7D3                       A:00 X:33 Y:7F P:27 SP:FB
	// F7D3  D0 07     BNE $F7DC                       A:00 X:33 Y:7F P:27 SP:F9
	// F7D5  70 05     BVS $F7DC                       A:00 X:33 Y:7F P:27 SP:F9
	// F7D7  90 03     BCC $F7DC                       A:00 X:33 Y:7F P:27 SP:F9
	// F7D9  30 01     BMI $F7DC                       A:00 X:33 Y:7F P:27 SP:F9
	// F7DB  60        RTS                             A:00 X:33 Y:7F P:27 SP:F9
	// D279  C8        INY                             A:00 X:33 Y:7F P:27 SP:FB
	// D27A  A9 AA     LDA #$AA                        A:00 X:33 Y:80 P:A5 SP:FB
	// D27C  85 78     STA $78 = 00                    A:AA X:33 Y:80 P:A5 SP:FB
	// D27E  20 DF F7  JSR $F7DF                       A:AA X:33 Y:80 P:A5 SP:FB
	// F7DF  18        CLC                             A:AA X:33 Y:80 P:A5 SP:F9
	// F7E0  24 01     BIT $01 = FF                    A:AA X:33 Y:80 P:A4 SP:F9
	// F7E2  A9 55     LDA #$55                        A:AA X:33 Y:80 P:E4 SP:F9
	// F7E4  60        RTS                             A:55 X:33 Y:80 P:64 SP:F9
	// D281  25 78     AND $78 = AA                    A:55 X:33 Y:80 P:64 SP:FB
	// D283  20 E5 F7  JSR $F7E5                       A:00 X:33 Y:80 P:66 SP:FB
	// F7E5  D0 07     BNE $F7EE                       A:00 X:33 Y:80 P:66 SP:F9
	// F7E7  50 05     BVC $F7EE                       A:00 X:33 Y:80 P:66 SP:F9
	// F7E9  B0 03     BCS $F7EE                       A:00 X:33 Y:80 P:66 SP:F9
	// F7EB  30 01     BMI $F7EE                       A:00 X:33 Y:80 P:66 SP:F9
	// F7ED  60        RTS                             A:00 X:33 Y:80 P:66 SP:F9
	// D286  C8        INY                             A:00 X:33 Y:80 P:66 SP:FB
	// D287  A9 EF     LDA #$EF                        A:00 X:33 Y:81 P:E4 SP:FB
	// D289  85 78     STA $78 = AA                    A:EF X:33 Y:81 P:E4 SP:FB
	// D28B  20 F1 F7  JSR $F7F1                       A:EF X:33 Y:81 P:E4 SP:FB
	// F7F1  38        SEC                             A:EF X:33 Y:81 P:E4 SP:F9
	// F7F2  B8        CLV                             A:EF X:33 Y:81 P:E5 SP:F9
	// F7F3  A9 F8     LDA #$F8                        A:EF X:33 Y:81 P:A5 SP:F9
	// F7F5  60        RTS                             A:F8 X:33 Y:81 P:A5 SP:F9
	// D28E  25 78     AND $78 = EF                    A:F8 X:33 Y:81 P:A5 SP:FB
	// D290  20 F6 F7  JSR $F7F6                       A:E8 X:33 Y:81 P:A5 SP:FB
	// F7F6  90 09     BCC $F801                       A:E8 X:33 Y:81 P:A5 SP:F9
	// F7F8  10 07     BPL $F801                       A:E8 X:33 Y:81 P:A5 SP:F9
	// F7FA  C9 E8     CMP #$E8                        A:E8 X:33 Y:81 P:A5 SP:F9
	// F7FC  D0 03     BNE $F801                       A:E8 X:33 Y:81 P:27 SP:F9
	// F7FE  70 01     BVS $F801                       A:E8 X:33 Y:81 P:27 SP:F9
	// F800  60        RTS                             A:E8 X:33 Y:81 P:27 SP:F9
	// D293  C8        INY                             A:E8 X:33 Y:81 P:27 SP:FB
	// D294  A9 AA     LDA #$AA                        A:E8 X:33 Y:82 P:A5 SP:FB
	// D296  85 78     STA $78 = EF                    A:AA X:33 Y:82 P:A5 SP:FB
	// D298  20 04 F8  JSR $F804                       A:AA X:33 Y:82 P:A5 SP:FB
	// F804  18        CLC                             A:AA X:33 Y:82 P:A5 SP:F9
	// F805  24 01     BIT $01 = FF                    A:AA X:33 Y:82 P:A4 SP:F9
	// F807  A9 5F     LDA #$5F                        A:AA X:33 Y:82 P:E4 SP:F9
	// F809  60        RTS                             A:5F X:33 Y:82 P:64 SP:F9
	// D29B  45 78     EOR $78 = AA                    A:5F X:33 Y:82 P:64 SP:FB
	// D29D  20 0A F8  JSR $F80A                       A:F5 X:33 Y:82 P:E4 SP:FB
	// F80A  B0 09     BCS $F815                       A:F5 X:33 Y:82 P:E4 SP:F9
	// F80C  10 07     BPL $F815                       A:F5 X:33 Y:82 P:E4 SP:F9
	// F80E  C9 F5     CMP #$F5                        A:F5 X:33 Y:82 P:E4 SP:F9
	// F810  D0 03     BNE $F815                       A:F5 X:33 Y:82 P:67 SP:F9
	// F812  50 01     BVC $F815                       A:F5 X:33 Y:82 P:67 SP:F9
	// F814  60        RTS                             A:F5 X:33 Y:82 P:67 SP:F9
	// D2A0  C8        INY                             A:F5 X:33 Y:82 P:67 SP:FB
	// D2A1  A9 70     LDA #$70                        A:F5 X:33 Y:83 P:E5 SP:FB
	// D2A3  85 78     STA $78 = AA                    A:70 X:33 Y:83 P:65 SP:FB
	// D2A5  20 18 F8  JSR $F818                       A:70 X:33 Y:83 P:65 SP:FB
	// F818  38        SEC                             A:70 X:33 Y:83 P:65 SP:F9
	// F819  B8        CLV                             A:70 X:33 Y:83 P:65 SP:F9
	// F81A  A9 70     LDA #$70                        A:70 X:33 Y:83 P:25 SP:F9
	// F81C  60        RTS                             A:70 X:33 Y:83 P:25 SP:F9
	// D2A8  45 78     EOR $78 = 70                    A:70 X:33 Y:83 P:25 SP:FB
	// D2AA  20 1D F8  JSR $F81D                       A:00 X:33 Y:83 P:27 SP:FB
	// F81D  D0 07     BNE $F826                       A:00 X:33 Y:83 P:27 SP:F9
	// F81F  70 05     BVS $F826                       A:00 X:33 Y:83 P:27 SP:F9
	// F821  90 03     BCC $F826                       A:00 X:33 Y:83 P:27 SP:F9
	// F823  30 01     BMI $F826                       A:00 X:33 Y:83 P:27 SP:F9
	// F825  60        RTS                             A:00 X:33 Y:83 P:27 SP:F9
	// D2AD  C8        INY                             A:00 X:33 Y:83 P:27 SP:FB
	// D2AE  A9 69     LDA #$69                        A:00 X:33 Y:84 P:A5 SP:FB
	// D2B0  85 78     STA $78 = 70                    A:69 X:33 Y:84 P:25 SP:FB
	// D2B2  20 29 F8  JSR $F829                       A:69 X:33 Y:84 P:25 SP:FB
	// F829  18        CLC                             A:69 X:33 Y:84 P:25 SP:F9
	// F82A  24 01     BIT $01 = FF                    A:69 X:33 Y:84 P:24 SP:F9
	// F82C  A9 00     LDA #$00                        A:69 X:33 Y:84 P:E4 SP:F9
	// F82E  60        RTS                             A:00 X:33 Y:84 P:66 SP:F9
	// D2B5  65 78     ADC $78 = 69                    A:00 X:33 Y:84 P:66 SP:FB
	// D2B7  20 2F F8  JSR $F82F                       A:69 X:33 Y:84 P:24 SP:FB
	// F82F  30 09     BMI $F83A                       A:69 X:33 Y:84 P:24 SP:F9
	// F831  B0 07     BCS $F83A                       A:69 X:33 Y:84 P:24 SP:F9
	// F833  C9 69     CMP #$69                        A:69 X:33 Y:84 P:24 SP:F9
	// F835  D0 03     BNE $F83A                       A:69 X:33 Y:84 P:27 SP:F9
	// F837  70 01     BVS $F83A                       A:69 X:33 Y:84 P:27 SP:F9
	// F839  60        RTS                             A:69 X:33 Y:84 P:27 SP:F9
	// D2BA  C8        INY                             A:69 X:33 Y:84 P:27 SP:FB
	// D2BB  20 3D F8  JSR $F83D                       A:69 X:33 Y:85 P:A5 SP:FB
	// F83D  38        SEC                             A:69 X:33 Y:85 P:A5 SP:F9
	// F83E  24 01     BIT $01 = FF                    A:69 X:33 Y:85 P:A5 SP:F9
	// F840  A9 00     LDA #$00                        A:69 X:33 Y:85 P:E5 SP:F9
	// F842  60        RTS                             A:00 X:33 Y:85 P:67 SP:F9
	// D2BE  65 78     ADC $78 = 69                    A:00 X:33 Y:85 P:67 SP:FB
	// D2C0  20 43 F8  JSR $F843                       A:6A X:33 Y:85 P:24 SP:FB
	// F843  30 09     BMI $F84E                       A:6A X:33 Y:85 P:24 SP:F9
	// F845  B0 07     BCS $F84E                       A:6A X:33 Y:85 P:24 SP:F9
	// F847  C9 6A     CMP #$6A                        A:6A X:33 Y:85 P:24 SP:F9
	// F849  D0 03     BNE $F84E                       A:6A X:33 Y:85 P:27 SP:F9
	// F84B  70 01     BVS $F84E                       A:6A X:33 Y:85 P:27 SP:F9
	// F84D  60        RTS                             A:6A X:33 Y:85 P:27 SP:F9
	// D2C3  C8        INY                             A:6A X:33 Y:85 P:27 SP:FB
	// D2C4  A9 7F     LDA #$7F                        A:6A X:33 Y:86 P:A5 SP:FB
	// D2C6  85 78     STA $78 = 69                    A:7F X:33 Y:86 P:25 SP:FB
	// D2C8  20 51 F8  JSR $F851                       A:7F X:33 Y:86 P:25 SP:FB
	// F851  38        SEC                             A:7F X:33 Y:86 P:25 SP:F9
	// F852  B8        CLV                             A:7F X:33 Y:86 P:25 SP:F9
	// F853  A9 7F     LDA #$7F                        A:7F X:33 Y:86 P:25 SP:F9
	// F855  60        RTS                             A:7F X:33 Y:86 P:25 SP:F9
	// D2CB  65 78     ADC $78 = 7F                    A:7F X:33 Y:86 P:25 SP:FB
	// D2CD  20 56 F8  JSR $F856                       A:FF X:33 Y:86 P:E4 SP:FB
	// F856  10 09     BPL $F861                       A:FF X:33 Y:86 P:E4 SP:F9
	// F858  B0 07     BCS $F861                       A:FF X:33 Y:86 P:E4 SP:F9
	// F85A  C9 FF     CMP #$FF                        A:FF X:33 Y:86 P:E4 SP:F9
	// F85C  D0 03     BNE $F861                       A:FF X:33 Y:86 P:67 SP:F9
	// F85E  50 01     BVC $F861                       A:FF X:33 Y:86 P:67 SP:F9
	// F860  60        RTS                             A:FF X:33 Y:86 P:67 SP:F9
	// D2D0  C8        INY                             A:FF X:33 Y:86 P:67 SP:FB
	// D2D1  A9 80     LDA #$80                        A:FF X:33 Y:87 P:E5 SP:FB
	// D2D3  85 78     STA $78 = 7F                    A:80 X:33 Y:87 P:E5 SP:FB
	// D2D5  20 64 F8  JSR $F864                       A:80 X:33 Y:87 P:E5 SP:FB
	// F864  18        CLC                             A:80 X:33 Y:87 P:E5 SP:F9
	// F865  24 01     BIT $01 = FF                    A:80 X:33 Y:87 P:E4 SP:F9
	// F867  A9 7F     LDA #$7F                        A:80 X:33 Y:87 P:E4 SP:F9
	// F869  60        RTS                             A:7F X:33 Y:87 P:64 SP:F9
	// D2D8  65 78     ADC $78 = 80                    A:7F X:33 Y:87 P:64 SP:FB
	// D2DA  20 6A F8  JSR $F86A                       A:FF X:33 Y:87 P:A4 SP:FB
	// F86A  10 09     BPL $F875                       A:FF X:33 Y:87 P:A4 SP:F9
	// F86C  B0 07     BCS $F875                       A:FF X:33 Y:87 P:A4 SP:F9
	// F86E  C9 FF     CMP #$FF                        A:FF X:33 Y:87 P:A4 SP:F9
	// F870  D0 03     BNE $F875                       A:FF X:33 Y:87 P:27 SP:F9
	// F872  70 01     BVS $F875                       A:FF X:33 Y:87 P:27 SP:F9
	// F874  60        RTS                             A:FF X:33 Y:87 P:27 SP:F9
	// D2DD  C8        INY                             A:FF X:33 Y:87 P:27 SP:FB
	// D2DE  20 78 F8  JSR $F878                       A:FF X:33 Y:88 P:A5 SP:FB
	// F878  38        SEC                             A:FF X:33 Y:88 P:A5 SP:F9
	// F879  B8        CLV                             A:FF X:33 Y:88 P:A5 SP:F9
	// F87A  A9 7F     LDA #$7F                        A:FF X:33 Y:88 P:A5 SP:F9
	// F87C  60        RTS                             A:7F X:33 Y:88 P:25 SP:F9
	// D2E1  65 78     ADC $78 = 80                    A:7F X:33 Y:88 P:25 SP:FB
	// D2E3  20 7D F8  JSR $F87D                       A:00 X:33 Y:88 P:27 SP:FB
	// F87D  D0 07     BNE $F886                       A:00 X:33 Y:88 P:27 SP:F9
	// F87F  30 05     BMI $F886                       A:00 X:33 Y:88 P:27 SP:F9
	// F881  70 03     BVS $F886                       A:00 X:33 Y:88 P:27 SP:F9
	// F883  90 01     BCC $F886                       A:00 X:33 Y:88 P:27 SP:F9
	// F885  60        RTS                             A:00 X:33 Y:88 P:27 SP:F9
	// D2E6  C8        INY                             A:00 X:33 Y:88 P:27 SP:FB
	// D2E7  A9 40     LDA #$40                        A:00 X:33 Y:89 P:A5 SP:FB
	// D2E9  85 78     STA $78 = 80                    A:40 X:33 Y:89 P:25 SP:FB
	// D2EB  20 89 F8  JSR $F889                       A:40 X:33 Y:89 P:25 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:33 Y:89 P:25 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:33 Y:89 P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:33 Y:89 P:65 SP:F9
	// D2EE  C5 78     CMP $78 = 40                    A:40 X:33 Y:89 P:65 SP:FB
	// D2F0  20 8E F8  JSR $F88E                       A:40 X:33 Y:89 P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:33 Y:89 P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:33 Y:89 P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:33 Y:89 P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:33 Y:89 P:67 SP:F9
	// F896  60        RTS                             A:40 X:33 Y:89 P:67 SP:F9
	// D2F3  C8        INY                             A:40 X:33 Y:89 P:67 SP:FB
	// D2F4  48        PHA                             A:40 X:33 Y:8A P:E5 SP:FB
	// D2F5  A9 3F     LDA #$3F                        A:40 X:33 Y:8A P:E5 SP:FA
	// D2F7  85 78     STA $78 = 40                    A:3F X:33 Y:8A P:65 SP:FA
	// D2F9  68        PLA                             A:3F X:33 Y:8A P:65 SP:FA
	// D2FA  20 9A F8  JSR $F89A                       A:40 X:33 Y:8A P:65 SP:FB
	// F89A  B8        CLV                             A:40 X:33 Y:8A P:65 SP:F9
	// F89B  60        RTS                             A:40 X:33 Y:8A P:25 SP:F9
	// D2FD  C5 78     CMP $78 = 3F                    A:40 X:33 Y:8A P:25 SP:FB
	// D2FF  20 9C F8  JSR $F89C                       A:40 X:33 Y:8A P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:40 X:33 Y:8A P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:40 X:33 Y:8A P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:40 X:33 Y:8A P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:40 X:33 Y:8A P:25 SP:F9
	// F8A4  60        RTS                             A:40 X:33 Y:8A P:25 SP:F9
	// D302  C8        INY                             A:40 X:33 Y:8A P:25 SP:FB
	// D303  48        PHA                             A:40 X:33 Y:8B P:A5 SP:FB
	// D304  A9 41     LDA #$41                        A:40 X:33 Y:8B P:A5 SP:FA
	// D306  85 78     STA $78 = 3F                    A:41 X:33 Y:8B P:25 SP:FA
	// D308  68        PLA                             A:41 X:33 Y:8B P:25 SP:FA
	// D309  C5 78     CMP $78 = 41                    A:40 X:33 Y:8B P:25 SP:FB
	// D30B  20 A8 F8  JSR $F8A8                       A:40 X:33 Y:8B P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:40 X:33 Y:8B P:A4 SP:F9
	// F8AE  60        RTS                             A:40 X:33 Y:8B P:A4 SP:F9
	// D30E  C8        INY                             A:40 X:33 Y:8B P:A4 SP:FB
	// D30F  48        PHA                             A:40 X:33 Y:8C P:A4 SP:FB
	// D310  A9 00     LDA #$00                        A:40 X:33 Y:8C P:A4 SP:FA
	// D312  85 78     STA $78 = 41                    A:00 X:33 Y:8C P:26 SP:FA
	// D314  68        PLA                             A:00 X:33 Y:8C P:26 SP:FA
	// D315  20 B2 F8  JSR $F8B2                       A:40 X:33 Y:8C P:24 SP:FB
	// F8B2  A9 80     LDA #$80                        A:40 X:33 Y:8C P:24 SP:F9
	// F8B4  60        RTS                             A:80 X:33 Y:8C P:A4 SP:F9
	// D318  C5 78     CMP $78 = 00                    A:80 X:33 Y:8C P:A4 SP:FB
	// D31A  20 B5 F8  JSR $F8B5                       A:80 X:33 Y:8C P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:33 Y:8C P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:33 Y:8C P:A5 SP:F9
	// D31D  C8        INY                             A:80 X:33 Y:8C P:A5 SP:FB
	// D31E  48        PHA                             A:80 X:33 Y:8D P:A5 SP:FB
	// D31F  A9 80     LDA #$80                        A:80 X:33 Y:8D P:A5 SP:FA
	// D321  85 78     STA $78 = 00                    A:80 X:33 Y:8D P:A5 SP:FA
	// D323  68        PLA                             A:80 X:33 Y:8D P:A5 SP:FA
	// D324  C5 78     CMP $78 = 80                    A:80 X:33 Y:8D P:A5 SP:FB
	// D326  20 BF F8  JSR $F8BF                       A:80 X:33 Y:8D P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:33 Y:8D P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:33 Y:8D P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:33 Y:8D P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:33 Y:8D P:27 SP:F9
	// D329  C8        INY                             A:80 X:33 Y:8D P:27 SP:FB
	// D32A  48        PHA                             A:80 X:33 Y:8E P:A5 SP:FB
	// D32B  A9 81     LDA #$81                        A:80 X:33 Y:8E P:A5 SP:FA
	// D32D  85 78     STA $78 = 80                    A:81 X:33 Y:8E P:A5 SP:FA
	// D32F  68        PLA                             A:81 X:33 Y:8E P:A5 SP:FA
	// D330  C5 78     CMP $78 = 81                    A:80 X:33 Y:8E P:A5 SP:FB
	// D332  20 C9 F8  JSR $F8C9                       A:80 X:33 Y:8E P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:80 X:33 Y:8E P:A4 SP:F9
	// F8CF  60        RTS                             A:80 X:33 Y:8E P:A4 SP:F9
	// D335  C8        INY                             A:80 X:33 Y:8E P:A4 SP:FB
	// D336  48        PHA                             A:80 X:33 Y:8F P:A4 SP:FB
	// D337  A9 7F     LDA #$7F                        A:80 X:33 Y:8F P:A4 SP:FA
	// D339  85 78     STA $78 = 81                    A:7F X:33 Y:8F P:24 SP:FA
	// D33B  68        PLA                             A:7F X:33 Y:8F P:24 SP:FA
	// D33C  C5 78     CMP $78 = 7F                    A:80 X:33 Y:8F P:A4 SP:FB
	// D33E  20 D3 F8  JSR $F8D3                       A:80 X:33 Y:8F P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:80 X:33 Y:8F P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:80 X:33 Y:8F P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:80 X:33 Y:8F P:25 SP:F9
	// F8D9  60        RTS                             A:80 X:33 Y:8F P:25 SP:F9
	// D341  C8        INY                             A:80 X:33 Y:8F P:25 SP:FB
	// D342  A9 40     LDA #$40                        A:80 X:33 Y:90 P:A5 SP:FB
	// D344  85 78     STA $78 = 7F                    A:40 X:33 Y:90 P:25 SP:FB
	// D346  20 31 F9  JSR $F931                       A:40 X:33 Y:90 P:25 SP:FB
	// F931  24 01     BIT $01 = FF                    A:40 X:33 Y:90 P:25 SP:F9
	// F933  A9 40     LDA #$40                        A:40 X:33 Y:90 P:E5 SP:F9
	// F935  38        SEC                             A:40 X:33 Y:90 P:65 SP:F9
	// F936  60        RTS                             A:40 X:33 Y:90 P:65 SP:F9
	// D349  E5 78     SBC $78 = 40                    A:40 X:33 Y:90 P:65 SP:FB
	// D34B  20 37 F9  JSR $F937                       A:00 X:33 Y:90 P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:33 Y:90 P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:33 Y:90 P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:33 Y:90 P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:33 Y:90 P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:33 Y:90 P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:33 Y:90 P:27 SP:F9
	// F943  60        RTS                             A:00 X:33 Y:90 P:27 SP:F9
	// D34E  C8        INY                             A:00 X:33 Y:90 P:27 SP:FB
	// D34F  A9 3F     LDA #$3F                        A:00 X:33 Y:91 P:A5 SP:FB
	// D351  85 78     STA $78 = 40                    A:3F X:33 Y:91 P:25 SP:FB
	// D353  20 47 F9  JSR $F947                       A:3F X:33 Y:91 P:25 SP:FB
	// F947  B8        CLV                             A:3F X:33 Y:91 P:25 SP:F9
	// F948  38        SEC                             A:3F X:33 Y:91 P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:3F X:33 Y:91 P:25 SP:F9
	// F94B  60        RTS                             A:40 X:33 Y:91 P:25 SP:F9
	// D356  E5 78     SBC $78 = 3F                    A:40 X:33 Y:91 P:25 SP:FB
	// D358  20 4C F9  JSR $F94C                       A:01 X:33 Y:91 P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:33 Y:91 P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:33 Y:91 P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:33 Y:91 P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:33 Y:91 P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:33 Y:91 P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:33 Y:91 P:27 SP:F9
	// F958  60        RTS                             A:01 X:33 Y:91 P:27 SP:F9
	// D35B  C8        INY                             A:01 X:33 Y:91 P:27 SP:FB
	// D35C  A9 41     LDA #$41                        A:01 X:33 Y:92 P:A5 SP:FB
	// D35E  85 78     STA $78 = 3F                    A:41 X:33 Y:92 P:25 SP:FB
	// D360  20 5C F9  JSR $F95C                       A:41 X:33 Y:92 P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:41 X:33 Y:92 P:25 SP:F9
	// F95E  38        SEC                             A:40 X:33 Y:92 P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:33 Y:92 P:25 SP:F9
	// F961  60        RTS                             A:40 X:33 Y:92 P:E5 SP:F9
	// D363  E5 78     SBC $78 = 41                    A:40 X:33 Y:92 P:E5 SP:FB
	// D365  20 62 F9  JSR $F962                       A:FF X:33 Y:92 P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:33 Y:92 P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:33 Y:92 P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:33 Y:92 P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:33 Y:92 P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:33 Y:92 P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:33 Y:92 P:27 SP:F9
	// F96E  60        RTS                             A:FF X:33 Y:92 P:27 SP:F9
	// D368  C8        INY                             A:FF X:33 Y:92 P:27 SP:FB
	// D369  A9 00     LDA #$00                        A:FF X:33 Y:93 P:A5 SP:FB
	// D36B  85 78     STA $78 = 41                    A:00 X:33 Y:93 P:27 SP:FB
	// D36D  20 72 F9  JSR $F972                       A:00 X:33 Y:93 P:27 SP:FB
	// F972  18        CLC                             A:00 X:33 Y:93 P:27 SP:F9
	// F973  A9 80     LDA #$80                        A:00 X:33 Y:93 P:26 SP:F9
	// F975  60        RTS                             A:80 X:33 Y:93 P:A4 SP:F9
	// D370  E5 78     SBC $78 = 00                    A:80 X:33 Y:93 P:A4 SP:FB
	// D372  20 76 F9  JSR $F976                       A:7F X:33 Y:93 P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:33 Y:93 P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:33 Y:93 P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:33 Y:93 P:67 SP:F9
	// F97C  60        RTS                             A:7F X:33 Y:93 P:67 SP:F9
	// D375  C8        INY                             A:7F X:33 Y:93 P:67 SP:FB
	// D376  A9 7F     LDA #$7F                        A:7F X:33 Y:94 P:E5 SP:FB
	// D378  85 78     STA $78 = 00                    A:7F X:33 Y:94 P:65 SP:FB
	// D37A  20 80 F9  JSR $F980                       A:7F X:33 Y:94 P:65 SP:FB
	// F980  38        SEC                             A:7F X:33 Y:94 P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:33 Y:94 P:65 SP:F9
	// F983  60        RTS                             A:81 X:33 Y:94 P:E5 SP:F9
	// D37D  E5 78     SBC $78 = 7F                    A:81 X:33 Y:94 P:E5 SP:FB
	// D37F  20 84 F9  JSR $F984                       A:02 X:33 Y:94 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:33 Y:94 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:33 Y:94 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:33 Y:94 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:33 Y:94 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:33 Y:94 P:67 SP:F9
	// D382  C8        INY                             A:02 X:33 Y:94 P:67 SP:FB
	// D383  A9 40     LDA #$40                        A:02 X:33 Y:95 P:E5 SP:FB
	// D385  85 78     STA $78 = 7F                    A:40 X:33 Y:95 P:65 SP:FB
	// D387  20 89 F8  JSR $F889                       A:40 X:33 Y:95 P:65 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:33 Y:95 P:65 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:33 Y:95 P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:33 Y:95 P:65 SP:F9
	// D38A  AA        TAX                             A:40 X:33 Y:95 P:65 SP:FB
	// D38B  E4 78     CPX $78 = 40                    A:40 X:40 Y:95 P:65 SP:FB
	// D38D  20 8E F8  JSR $F88E                       A:40 X:40 Y:95 P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:40 Y:95 P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:40 Y:95 P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:40 Y:95 P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:40 Y:95 P:67 SP:F9
	// F896  60        RTS                             A:40 X:40 Y:95 P:67 SP:F9
	// D390  C8        INY                             A:40 X:40 Y:95 P:67 SP:FB
	// D391  A9 3F     LDA #$3F                        A:40 X:40 Y:96 P:E5 SP:FB
	// D393  85 78     STA $78 = 40                    A:3F X:40 Y:96 P:65 SP:FB
	// D395  20 9A F8  JSR $F89A                       A:3F X:40 Y:96 P:65 SP:FB
	// F89A  B8        CLV                             A:3F X:40 Y:96 P:65 SP:F9
	// F89B  60        RTS                             A:3F X:40 Y:96 P:25 SP:F9
	// D398  E4 78     CPX $78 = 3F                    A:3F X:40 Y:96 P:25 SP:FB
	// D39A  20 9C F8  JSR $F89C                       A:3F X:40 Y:96 P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:3F X:40 Y:96 P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:3F X:40 Y:96 P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:3F X:40 Y:96 P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:3F X:40 Y:96 P:25 SP:F9
	// F8A4  60        RTS                             A:3F X:40 Y:96 P:25 SP:F9
	// D39D  C8        INY                             A:3F X:40 Y:96 P:25 SP:FB
	// D39E  A9 41     LDA #$41                        A:3F X:40 Y:97 P:A5 SP:FB
	// D3A0  85 78     STA $78 = 3F                    A:41 X:40 Y:97 P:25 SP:FB
	// D3A2  E4 78     CPX $78 = 41                    A:41 X:40 Y:97 P:25 SP:FB
	// D3A4  20 A8 F8  JSR $F8A8                       A:41 X:40 Y:97 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:41 X:40 Y:97 P:A4 SP:F9
	// F8AE  60        RTS                             A:41 X:40 Y:97 P:A4 SP:F9
	// D3A7  C8        INY                             A:41 X:40 Y:97 P:A4 SP:FB
	// D3A8  A9 00     LDA #$00                        A:41 X:40 Y:98 P:A4 SP:FB
	// D3AA  85 78     STA $78 = 41                    A:00 X:40 Y:98 P:26 SP:FB
	// D3AC  20 B2 F8  JSR $F8B2                       A:00 X:40 Y:98 P:26 SP:FB
	// F8B2  A9 80     LDA #$80                        A:00 X:40 Y:98 P:26 SP:F9
	// F8B4  60        RTS                             A:80 X:40 Y:98 P:A4 SP:F9
	// D3AF  AA        TAX                             A:80 X:40 Y:98 P:A4 SP:FB
	// D3B0  E4 78     CPX $78 = 00                    A:80 X:80 Y:98 P:A4 SP:FB
	// D3B2  20 B5 F8  JSR $F8B5                       A:80 X:80 Y:98 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:80 Y:98 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:80 Y:98 P:A5 SP:F9
	// D3B5  C8        INY                             A:80 X:80 Y:98 P:A5 SP:FB
	// D3B6  A9 80     LDA #$80                        A:80 X:80 Y:99 P:A5 SP:FB
	// D3B8  85 78     STA $78 = 00                    A:80 X:80 Y:99 P:A5 SP:FB
	// D3BA  E4 78     CPX $78 = 80                    A:80 X:80 Y:99 P:A5 SP:FB
	// D3BC  20 BF F8  JSR $F8BF                       A:80 X:80 Y:99 P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:80 Y:99 P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:80 Y:99 P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:80 Y:99 P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:80 Y:99 P:27 SP:F9
	// D3BF  C8        INY                             A:80 X:80 Y:99 P:27 SP:FB
	// D3C0  A9 81     LDA #$81                        A:80 X:80 Y:9A P:A5 SP:FB
	// D3C2  85 78     STA $78 = 80                    A:81 X:80 Y:9A P:A5 SP:FB
	// D3C4  E4 78     CPX $78 = 81                    A:81 X:80 Y:9A P:A5 SP:FB
	// D3C6  20 C9 F8  JSR $F8C9                       A:81 X:80 Y:9A P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:81 X:80 Y:9A P:A4 SP:F9
	// F8CF  60        RTS                             A:81 X:80 Y:9A P:A4 SP:F9
	// D3C9  C8        INY                             A:81 X:80 Y:9A P:A4 SP:FB
	// D3CA  A9 7F     LDA #$7F                        A:81 X:80 Y:9B P:A4 SP:FB
	// D3CC  85 78     STA $78 = 81                    A:7F X:80 Y:9B P:24 SP:FB
	// D3CE  E4 78     CPX $78 = 7F                    A:7F X:80 Y:9B P:24 SP:FB
	// D3D0  20 D3 F8  JSR $F8D3                       A:7F X:80 Y:9B P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:7F X:80 Y:9B P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:7F X:80 Y:9B P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:7F X:80 Y:9B P:25 SP:F9
	// F8D9  60        RTS                             A:7F X:80 Y:9B P:25 SP:F9
	// D3D3  C8        INY                             A:7F X:80 Y:9B P:25 SP:FB
	// D3D4  98        TYA                             A:7F X:80 Y:9C P:A5 SP:FB
	// D3D5  AA        TAX                             A:9C X:80 Y:9C P:A5 SP:FB
	// D3D6  A9 40     LDA #$40                        A:9C X:9C Y:9C P:A5 SP:FB
	// D3D8  85 78     STA $78 = 7F                    A:40 X:9C Y:9C P:25 SP:FB
	// D3DA  20 DD F8  JSR $F8DD                       A:40 X:9C Y:9C P:25 SP:FB
	// F8DD  24 01     BIT $01 = FF                    A:40 X:9C Y:9C P:25 SP:F9
	// F8DF  A0 40     LDY #$40                        A:40 X:9C Y:9C P:E5 SP:F9
	// F8E1  60        RTS                             A:40 X:9C Y:40 P:65 SP:F9
	// D3DD  C4 78     CPY $78 = 40                    A:40 X:9C Y:40 P:65 SP:FB
	// D3DF  20 E2 F8  JSR $F8E2                       A:40 X:9C Y:40 P:67 SP:FB
	// F8E2  30 07     BMI $F8EB                       A:40 X:9C Y:40 P:67 SP:F9
	// F8E4  90 05     BCC $F8EB                       A:40 X:9C Y:40 P:67 SP:F9
	// F8E6  D0 03     BNE $F8EB                       A:40 X:9C Y:40 P:67 SP:F9
	// F8E8  50 01     BVC $F8EB                       A:40 X:9C Y:40 P:67 SP:F9
	// F8EA  60        RTS                             A:40 X:9C Y:40 P:67 SP:F9
	// D3E2  E8        INX                             A:40 X:9C Y:40 P:67 SP:FB
	// D3E3  A9 3F     LDA #$3F                        A:40 X:9D Y:40 P:E5 SP:FB
	// D3E5  85 78     STA $78 = 40                    A:3F X:9D Y:40 P:65 SP:FB
	// D3E7  20 EE F8  JSR $F8EE                       A:3F X:9D Y:40 P:65 SP:FB
	// F8EE  B8        CLV                             A:3F X:9D Y:40 P:65 SP:F9
	// F8EF  60        RTS                             A:3F X:9D Y:40 P:25 SP:F9
	// D3EA  C4 78     CPY $78 = 3F                    A:3F X:9D Y:40 P:25 SP:FB
	// D3EC  20 F0 F8  JSR $F8F0                       A:3F X:9D Y:40 P:25 SP:FB
	// F8F0  F0 07     BEQ $F8F9                       A:3F X:9D Y:40 P:25 SP:F9
	// F8F2  30 05     BMI $F8F9                       A:3F X:9D Y:40 P:25 SP:F9
	// F8F4  90 03     BCC $F8F9                       A:3F X:9D Y:40 P:25 SP:F9
	// F8F6  70 01     BVS $F8F9                       A:3F X:9D Y:40 P:25 SP:F9
	// F8F8  60        RTS                             A:3F X:9D Y:40 P:25 SP:F9
	// D3EF  E8        INX                             A:3F X:9D Y:40 P:25 SP:FB
	// D3F0  A9 41     LDA #$41                        A:3F X:9E Y:40 P:A5 SP:FB
	// D3F2  85 78     STA $78 = 3F                    A:41 X:9E Y:40 P:25 SP:FB
	// D3F4  C4 78     CPY $78 = 41                    A:41 X:9E Y:40 P:25 SP:FB
	// D3F6  20 FC F8  JSR $F8FC                       A:41 X:9E Y:40 P:A4 SP:FB
	// F8FC  F0 05     BEQ $F903                       A:41 X:9E Y:40 P:A4 SP:F9
	// F8FE  10 03     BPL $F903                       A:41 X:9E Y:40 P:A4 SP:F9
	// F900  10 01     BPL $F903                       A:41 X:9E Y:40 P:A4 SP:F9
	// F902  60        RTS                             A:41 X:9E Y:40 P:A4 SP:F9
	// D3F9  E8        INX                             A:41 X:9E Y:40 P:A4 SP:FB
	// D3FA  A9 00     LDA #$00                        A:41 X:9F Y:40 P:A4 SP:FB
	// D3FC  85 78     STA $78 = 41                    A:00 X:9F Y:40 P:26 SP:FB
	// D3FE  20 06 F9  JSR $F906                       A:00 X:9F Y:40 P:26 SP:FB
	// F906  A0 80     LDY #$80                        A:00 X:9F Y:40 P:26 SP:F9
	// F908  60        RTS                             A:00 X:9F Y:80 P:A4 SP:F9
	// D401  C4 78     CPY $78 = 00                    A:00 X:9F Y:80 P:A4 SP:FB
	// D403  20 09 F9  JSR $F909                       A:00 X:9F Y:80 P:A5 SP:FB
	// F909  F0 05     BEQ $F910                       A:00 X:9F Y:80 P:A5 SP:F9
	// F90B  10 03     BPL $F910                       A:00 X:9F Y:80 P:A5 SP:F9
	// F90D  90 01     BCC $F910                       A:00 X:9F Y:80 P:A5 SP:F9
	// F90F  60        RTS                             A:00 X:9F Y:80 P:A5 SP:F9
	// D406  E8        INX                             A:00 X:9F Y:80 P:A5 SP:FB
	// D407  A9 80     LDA #$80                        A:00 X:A0 Y:80 P:A5 SP:FB
	// D409  85 78     STA $78 = 00                    A:80 X:A0 Y:80 P:A5 SP:FB
	// D40B  C4 78     CPY $78 = 80                    A:80 X:A0 Y:80 P:A5 SP:FB
	// D40D  20 13 F9  JSR $F913                       A:80 X:A0 Y:80 P:27 SP:FB
	// F913  D0 05     BNE $F91A                       A:80 X:A0 Y:80 P:27 SP:F9
	// F915  30 03     BMI $F91A                       A:80 X:A0 Y:80 P:27 SP:F9
	// F917  90 01     BCC $F91A                       A:80 X:A0 Y:80 P:27 SP:F9
	// F919  60        RTS                             A:80 X:A0 Y:80 P:27 SP:F9
	// D410  E8        INX                             A:80 X:A0 Y:80 P:27 SP:FB
	// D411  A9 81     LDA #$81                        A:80 X:A1 Y:80 P:A5 SP:FB
	// D413  85 78     STA $78 = 80                    A:81 X:A1 Y:80 P:A5 SP:FB
	// D415  C4 78     CPY $78 = 81                    A:81 X:A1 Y:80 P:A5 SP:FB
	// D417  20 1D F9  JSR $F91D                       A:81 X:A1 Y:80 P:A4 SP:FB
	// F91D  B0 05     BCS $F924                       A:81 X:A1 Y:80 P:A4 SP:F9
	// F91F  F0 03     BEQ $F924                       A:81 X:A1 Y:80 P:A4 SP:F9
	// F921  10 01     BPL $F924                       A:81 X:A1 Y:80 P:A4 SP:F9
	// F923  60        RTS                             A:81 X:A1 Y:80 P:A4 SP:F9
	// D41A  E8        INX                             A:81 X:A1 Y:80 P:A4 SP:FB
	// D41B  A9 7F     LDA #$7F                        A:81 X:A2 Y:80 P:A4 SP:FB
	// D41D  85 78     STA $78 = 81                    A:7F X:A2 Y:80 P:24 SP:FB
	// D41F  C4 78     CPY $78 = 7F                    A:7F X:A2 Y:80 P:24 SP:FB
	// D421  20 27 F9  JSR $F927                       A:7F X:A2 Y:80 P:25 SP:FB
	// F927  90 05     BCC $F92E                       A:7F X:A2 Y:80 P:25 SP:F9
	// F929  F0 03     BEQ $F92E                       A:7F X:A2 Y:80 P:25 SP:F9
	// F92B  30 01     BMI $F92E                       A:7F X:A2 Y:80 P:25 SP:F9
	// F92D  60        RTS                             A:7F X:A2 Y:80 P:25 SP:F9
	// D424  E8        INX                             A:7F X:A2 Y:80 P:25 SP:FB
	// D425  8A        TXA                             A:7F X:A3 Y:80 P:A5 SP:FB
	// D426  A8        TAY                             A:A3 X:A3 Y:80 P:A5 SP:FB
	// D427  20 90 F9  JSR $F990                       A:A3 X:A3 Y:A3 P:A5 SP:FB
	// F990  A2 55     LDX #$55                        A:A3 X:A3 Y:A3 P:A5 SP:F9
	// F992  A9 FF     LDA #$FF                        A:A3 X:55 Y:A3 P:25 SP:F9
	// F994  85 01     STA $01 = FF                    A:FF X:55 Y:A3 P:A5 SP:F9
	// F996  EA        NOP                             A:FF X:55 Y:A3 P:A5 SP:F9
	// F997  24 01     BIT $01 = FF                    A:FF X:55 Y:A3 P:A5 SP:F9
	// F999  38        SEC                             A:FF X:55 Y:A3 P:E5 SP:F9
	// F99A  A9 01     LDA #$01                        A:FF X:55 Y:A3 P:E5 SP:F9
	// F99C  60        RTS                             A:01 X:55 Y:A3 P:65 SP:F9
	// D42A  85 78     STA $78 = 7F                    A:01 X:55 Y:A3 P:65 SP:FB
	// D42C  46 78     LSR $78 = 01                    A:01 X:55 Y:A3 P:65 SP:FB
	// D42E  A5 78     LDA $78 = 00                    A:01 X:55 Y:A3 P:67 SP:FB
	// D430  20 9D F9  JSR $F99D                       A:00 X:55 Y:A3 P:67 SP:FB
	// F99D  90 1B     BCC $F9BA                       A:00 X:55 Y:A3 P:67 SP:F9
	// F99F  D0 19     BNE $F9BA                       A:00 X:55 Y:A3 P:67 SP:F9
	// F9A1  30 17     BMI $F9BA                       A:00 X:55 Y:A3 P:67 SP:F9
	// F9A3  50 15     BVC $F9BA                       A:00 X:55 Y:A3 P:67 SP:F9
	// F9A5  C9 00     CMP #$00                        A:00 X:55 Y:A3 P:67 SP:F9
	// F9A7  D0 11     BNE $F9BA                       A:00 X:55 Y:A3 P:67 SP:F9
	// F9A9  B8        CLV                             A:00 X:55 Y:A3 P:67 SP:F9
	// F9AA  A9 AA     LDA #$AA                        A:00 X:55 Y:A3 P:27 SP:F9
	// F9AC  60        RTS                             A:AA X:55 Y:A3 P:A5 SP:F9
	// D433  C8        INY                             A:AA X:55 Y:A3 P:A5 SP:FB
	// D434  85 78     STA $78 = 00                    A:AA X:55 Y:A4 P:A5 SP:FB
	// D436  46 78     LSR $78 = AA                    A:AA X:55 Y:A4 P:A5 SP:FB
	// D438  A5 78     LDA $78 = 55                    A:AA X:55 Y:A4 P:24 SP:FB
	// D43A  20 AD F9  JSR $F9AD                       A:55 X:55 Y:A4 P:24 SP:FB
	// F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9
	// F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9
	// F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9
	// F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:A4 P:24 SP:F9
	// F9B5  C9 55     CMP #$55                        A:55 X:55 Y:A4 P:24 SP:F9
	// F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:A4 P:27 SP:F9
	// F9B9  60        RTS                             A:55 X:55 Y:A4 P:27 SP:F9
	// D43D  C8        INY                             A:55 X:55 Y:A4 P:27 SP:FB
	// D43E  20 BD F9  JSR $F9BD                       A:55 X:55 Y:A5 P:A5 SP:FB
	// F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:A5 P:A5 SP:F9
	// F9BF  38        SEC                             A:55 X:55 Y:A5 P:E5 SP:F9
	// F9C0  A9 80     LDA #$80                        A:55 X:55 Y:A5 P:E5 SP:F9
	// F9C2  60        RTS                             A:80 X:55 Y:A5 P:E5 SP:F9
	// D441  85 78     STA $78 = 55                    A:80 X:55 Y:A5 P:E5 SP:FB
	// D443  06 78     ASL $78 = 80                    A:80 X:55 Y:A5 P:E5 SP:FB
	// D445  A5 78     LDA $78 = 00                    A:80 X:55 Y:A5 P:67 SP:FB
	// D447  20 C3 F9  JSR $F9C3                       A:00 X:55 Y:A5 P:67 SP:FB
	// F9C3  90 1C     BCC $F9E1                       A:00 X:55 Y:A5 P:67 SP:F9
	// F9C5  D0 1A     BNE $F9E1                       A:00 X:55 Y:A5 P:67 SP:F9
	// F9C7  30 18     BMI $F9E1                       A:00 X:55 Y:A5 P:67 SP:F9
	// F9C9  50 16     BVC $F9E1                       A:00 X:55 Y:A5 P:67 SP:F9
	// F9CB  C9 00     CMP #$00                        A:00 X:55 Y:A5 P:67 SP:F9
	// F9CD  D0 12     BNE $F9E1                       A:00 X:55 Y:A5 P:67 SP:F9
	// F9CF  B8        CLV                             A:00 X:55 Y:A5 P:67 SP:F9
	// F9D0  A9 55     LDA #$55                        A:00 X:55 Y:A5 P:27 SP:F9
	// F9D2  38        SEC                             A:55 X:55 Y:A5 P:25 SP:F9
	// F9D3  60        RTS                             A:55 X:55 Y:A5 P:25 SP:F9
	// D44A  C8        INY                             A:55 X:55 Y:A5 P:25 SP:FB
	// D44B  85 78     STA $78 = 00                    A:55 X:55 Y:A6 P:A5 SP:FB
	// D44D  06 78     ASL $78 = 55                    A:55 X:55 Y:A6 P:A5 SP:FB
	// D44F  A5 78     LDA $78 = AA                    A:55 X:55 Y:A6 P:A4 SP:FB
	// D451  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:A6 P:A4 SP:FB
	// F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9
	// F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9
	// F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9
	// F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:A6 P:A4 SP:F9
	// F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:A6 P:A4 SP:F9
	// F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:A6 P:27 SP:F9
	// F9E0  60        RTS                             A:AA X:55 Y:A6 P:27 SP:F9
	// D454  C8        INY                             A:AA X:55 Y:A6 P:27 SP:FB
	// D455  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:A7 P:A5 SP:FB
	// F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:A7 P:A5 SP:F9
	// F9E6  38        SEC                             A:AA X:55 Y:A7 P:E5 SP:F9
	// F9E7  A9 01     LDA #$01                        A:AA X:55 Y:A7 P:E5 SP:F9
	// F9E9  60        RTS                             A:01 X:55 Y:A7 P:65 SP:F9
	// D458  85 78     STA $78 = AA                    A:01 X:55 Y:A7 P:65 SP:FB
	// D45A  66 78     ROR $78 = 01                    A:01 X:55 Y:A7 P:65 SP:FB
	// D45C  A5 78     LDA $78 = 80                    A:01 X:55 Y:A7 P:E5 SP:FB
	// D45E  20 EA F9  JSR $F9EA                       A:80 X:55 Y:A7 P:E5 SP:FB
	// F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9
	// F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9
	// F9EE  10 18     BPL $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9
	// F9F0  50 16     BVC $FA08                       A:80 X:55 Y:A7 P:E5 SP:F9
	// F9F2  C9 80     CMP #$80                        A:80 X:55 Y:A7 P:E5 SP:F9
	// F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:A7 P:67 SP:F9
	// F9F6  B8        CLV                             A:80 X:55 Y:A7 P:67 SP:F9
	// F9F7  18        CLC                             A:80 X:55 Y:A7 P:27 SP:F9
	// F9F8  A9 55     LDA #$55                        A:80 X:55 Y:A7 P:26 SP:F9
	// F9FA  60        RTS                             A:55 X:55 Y:A7 P:24 SP:F9
	// D461  C8        INY                             A:55 X:55 Y:A7 P:24 SP:FB
	// D462  85 78     STA $78 = 80                    A:55 X:55 Y:A8 P:A4 SP:FB
	// D464  66 78     ROR $78 = 55                    A:55 X:55 Y:A8 P:A4 SP:FB
	// D466  A5 78     LDA $78 = 2A                    A:55 X:55 Y:A8 P:25 SP:FB
	// D468  20 FB F9  JSR $F9FB                       A:2A X:55 Y:A8 P:25 SP:FB
	// F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:A8 P:25 SP:F9
	// F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:A8 P:25 SP:F9
	// F9FF  30 07     BMI $FA08                       A:2A X:55 Y:A8 P:25 SP:F9
	// FA01  70 05     BVS $FA08                       A:2A X:55 Y:A8 P:25 SP:F9
	// FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:A8 P:25 SP:F9
	// FA05  D0 01     BNE $FA08                       A:2A X:55 Y:A8 P:27 SP:F9
	// FA07  60        RTS                             A:2A X:55 Y:A8 P:27 SP:F9
	// D46B  C8        INY                             A:2A X:55 Y:A8 P:27 SP:FB
	// D46C  20 0A FA  JSR $FA0A                       A:2A X:55 Y:A9 P:A5 SP:FB
	// FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:A9 P:A5 SP:F9
	// FA0C  38        SEC                             A:2A X:55 Y:A9 P:E5 SP:F9
	// FA0D  A9 80     LDA #$80                        A:2A X:55 Y:A9 P:E5 SP:F9
	// FA0F  60        RTS                             A:80 X:55 Y:A9 P:E5 SP:F9
	// D46F  85 78     STA $78 = 2A                    A:80 X:55 Y:A9 P:E5 SP:FB
	// D471  26 78     ROL $78 = 80                    A:80 X:55 Y:A9 P:E5 SP:FB
	// D473  A5 78     LDA $78 = 01                    A:80 X:55 Y:A9 P:65 SP:FB
	// D475  20 10 FA  JSR $FA10                       A:01 X:55 Y:A9 P:65 SP:FB
	// FA10  90 1C     BCC $FA2E                       A:01 X:55 Y:A9 P:65 SP:F9
	// FA12  F0 1A     BEQ $FA2E                       A:01 X:55 Y:A9 P:65 SP:F9
	// FA14  30 18     BMI $FA2E                       A:01 X:55 Y:A9 P:65 SP:F9
	// FA16  50 16     BVC $FA2E                       A:01 X:55 Y:A9 P:65 SP:F9
	// FA18  C9 01     CMP #$01                        A:01 X:55 Y:A9 P:65 SP:F9
	// FA1A  D0 12     BNE $FA2E                       A:01 X:55 Y:A9 P:67 SP:F9
	// FA1C  B8        CLV                             A:01 X:55 Y:A9 P:67 SP:F9
	// FA1D  18        CLC                             A:01 X:55 Y:A9 P:27 SP:F9
	// FA1E  A9 55     LDA #$55                        A:01 X:55 Y:A9 P:26 SP:F9
	// FA20  60        RTS                             A:55 X:55 Y:A9 P:24 SP:F9
	// D478  C8        INY                             A:55 X:55 Y:A9 P:24 SP:FB
	// D479  85 78     STA $78 = 01                    A:55 X:55 Y:AA P:A4 SP:FB
	// D47B  26 78     ROL $78 = 55                    A:55 X:55 Y:AA P:A4 SP:FB
	// D47D  A5 78     LDA $78 = AA                    A:55 X:55 Y:AA P:A4 SP:FB
	// D47F  20 21 FA  JSR $FA21                       A:AA X:55 Y:AA P:A4 SP:FB
	// FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9
	// FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9
	// FA25  10 07     BPL $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9
	// FA27  70 05     BVS $FA2E                       A:AA X:55 Y:AA P:A4 SP:F9
	// FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:AA P:A4 SP:F9
	// FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:AA P:27 SP:F9
	// FA2D  60        RTS                             A:AA X:55 Y:AA P:27 SP:F9
	// D482  A9 FF     LDA #$FF                        A:AA X:55 Y:AA P:27 SP:FB
	// D484  85 78     STA $78 = AA                    A:FF X:55 Y:AA P:A5 SP:FB
	// D486  85 01     STA $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB
	// D488  24 01     BIT $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB
	// D48A  38        SEC                             A:FF X:55 Y:AA P:E5 SP:FB
	// D48B  E6 78     INC $78 = FF                    A:FF X:55 Y:AA P:E5 SP:FB
	// D48D  D0 0C     BNE $D49B                       A:FF X:55 Y:AA P:67 SP:FB
	// D48F  30 0A     BMI $D49B                       A:FF X:55 Y:AA P:67 SP:FB
	// D491  50 08     BVC $D49B                       A:FF X:55 Y:AA P:67 SP:FB
	// D493  90 06     BCC $D49B                       A:FF X:55 Y:AA P:67 SP:FB
	// D495  A5 78     LDA $78 = 00                    A:FF X:55 Y:AA P:67 SP:FB
	// D497  C9 00     CMP #$00                        A:00 X:55 Y:AA P:67 SP:FB
	// D499  F0 04     BEQ $D49F                       A:00 X:55 Y:AA P:67 SP:FB
	// D49F  A9 7F     LDA #$7F                        A:00 X:55 Y:AA P:67 SP:FB
	// D4A1  85 78     STA $78 = 00                    A:7F X:55 Y:AA P:65 SP:FB
	// D4A3  B8        CLV                             A:7F X:55 Y:AA P:65 SP:FB
	// D4A4  18        CLC                             A:7F X:55 Y:AA P:25 SP:FB
	// D4A5  E6 78     INC $78 = 7F                    A:7F X:55 Y:AA P:24 SP:FB
	// D4A7  F0 0C     BEQ $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB
	// D4A9  10 0A     BPL $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB
	// D4AB  70 08     BVS $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB
	// D4AD  B0 06     BCS $D4B5                       A:7F X:55 Y:AA P:A4 SP:FB
	// D4AF  A5 78     LDA $78 = 80                    A:7F X:55 Y:AA P:A4 SP:FB
	// D4B1  C9 80     CMP #$80                        A:80 X:55 Y:AA P:A4 SP:FB
	// D4B3  F0 04     BEQ $D4B9                       A:80 X:55 Y:AA P:27 SP:FB
	// D4B9  A9 00     LDA #$00                        A:80 X:55 Y:AA P:27 SP:FB
	// D4BB  85 78     STA $78 = 80                    A:00 X:55 Y:AA P:27 SP:FB
	// D4BD  24 01     BIT $01 = FF                    A:00 X:55 Y:AA P:27 SP:FB
	// D4BF  38        SEC                             A:00 X:55 Y:AA P:E7 SP:FB
	// D4C0  C6 78     DEC $78 = 00                    A:00 X:55 Y:AA P:E7 SP:FB
	// D4C2  F0 0C     BEQ $D4D0                       A:00 X:55 Y:AA P:E5 SP:FB
	// D4C4  10 0A     BPL $D4D0                       A:00 X:55 Y:AA P:E5 SP:FB
	// D4C6  50 08     BVC $D4D0                       A:00 X:55 Y:AA P:E5 SP:FB
	// D4C8  90 06     BCC $D4D0                       A:00 X:55 Y:AA P:E5 SP:FB
	// D4CA  A5 78     LDA $78 = FF                    A:00 X:55 Y:AA P:E5 SP:FB
	// D4CC  C9 FF     CMP #$FF                        A:FF X:55 Y:AA P:E5 SP:FB
	// D4CE  F0 04     BEQ $D4D4                       A:FF X:55 Y:AA P:67 SP:FB
	// D4D4  A9 80     LDA #$80                        A:FF X:55 Y:AA P:67 SP:FB
	// D4D6  85 78     STA $78 = FF                    A:80 X:55 Y:AA P:E5 SP:FB
	// D4D8  B8        CLV                             A:80 X:55 Y:AA P:E5 SP:FB
	// D4D9  18        CLC                             A:80 X:55 Y:AA P:A5 SP:FB
	// D4DA  C6 78     DEC $78 = 80                    A:80 X:55 Y:AA P:A4 SP:FB
	// D4DC  F0 0C     BEQ $D4EA                       A:80 X:55 Y:AA P:24 SP:FB
	// D4DE  30 0A     BMI $D4EA                       A:80 X:55 Y:AA P:24 SP:FB
	// D4E0  70 08     BVS $D4EA                       A:80 X:55 Y:AA P:24 SP:FB
	// D4E2  B0 06     BCS $D4EA                       A:80 X:55 Y:AA P:24 SP:FB
	// D4E4  A5 78     LDA $78 = 7F                    A:80 X:55 Y:AA P:24 SP:FB
	// D4E6  C9 7F     CMP #$7F                        A:7F X:55 Y:AA P:24 SP:FB
	// D4E8  F0 04     BEQ $D4EE                       A:7F X:55 Y:AA P:27 SP:FB
	// D4EE  A9 01     LDA #$01                        A:7F X:55 Y:AA P:27 SP:FB
	// D4F0  85 78     STA $78 = 7F                    A:01 X:55 Y:AA P:25 SP:FB
	// D4F2  C6 78     DEC $78 = 01                    A:01 X:55 Y:AA P:25 SP:FB
	// D4F4  F0 04     BEQ $D4FA                       A:01 X:55 Y:AA P:27 SP:FB
	// D4FA  60        RTS                             A:01 X:55 Y:AA P:27 SP:FB
	// C615  20 FB D4  JSR $D4FB                       A:01 X:55 Y:AA P:27 SP:FD
	// D4FB  A9 55     LDA #$55                        A:01 X:55 Y:AA P:27 SP:FB
	// D4FD  8D 78 06  STA $0678 = 00                  A:55 X:55 Y:AA P:25 SP:FB
	// D500  A9 FF     LDA #$FF                        A:55 X:55 Y:AA P:25 SP:FB
	// D502  85 01     STA $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB
	// D504  24 01     BIT $01 = FF                    A:FF X:55 Y:AA P:A5 SP:FB
	// D506  A0 11     LDY #$11                        A:FF X:55 Y:AA P:E5 SP:FB
	// D508  A2 23     LDX #$23                        A:FF X:55 Y:11 P:65 SP:FB
	// D50A  A9 00     LDA #$00                        A:FF X:23 Y:11 P:65 SP:FB
	// D50C  AD 78 06  LDA $0678 = 55                  A:00 X:23 Y:11 P:67 SP:FB
	// D50F  F0 10     BEQ $D521                       A:55 X:23 Y:11 P:65 SP:FB
	// D511  30 0E     BMI $D521                       A:55 X:23 Y:11 P:65 SP:FB
	// D513  C9 55     CMP #$55                        A:55 X:23 Y:11 P:65 SP:FB
	// D515  D0 0A     BNE $D521                       A:55 X:23 Y:11 P:67 SP:FB
	// D517  C0 11     CPY #$11                        A:55 X:23 Y:11 P:67 SP:FB
	// D519  D0 06     BNE $D521                       A:55 X:23 Y:11 P:67 SP:FB
	// D51B  E0 23     CPX #$23                        A:55 X:23 Y:11 P:67 SP:FB
	// D51D  50 02     BVC $D521                       A:55 X:23 Y:11 P:67 SP:FB
	// D51F  F0 04     BEQ $D525                       A:55 X:23 Y:11 P:67 SP:FB
	// D525  A9 46     LDA #$46                        A:55 X:23 Y:11 P:67 SP:FB
	// D527  24 01     BIT $01 = FF                    A:46 X:23 Y:11 P:65 SP:FB
	// D529  8D 78 06  STA $0678 = 55                  A:46 X:23 Y:11 P:E5 SP:FB
	// D52C  F0 0B     BEQ $D539                       A:46 X:23 Y:11 P:E5 SP:FB
	// D52E  10 09     BPL $D539                       A:46 X:23 Y:11 P:E5 SP:FB
	// D530  50 07     BVC $D539                       A:46 X:23 Y:11 P:E5 SP:FB
	// D532  AD 78 06  LDA $0678 = 46                  A:46 X:23 Y:11 P:E5 SP:FB
	// D535  C9 46     CMP #$46                        A:46 X:23 Y:11 P:65 SP:FB
	// D537  F0 04     BEQ $D53D                       A:46 X:23 Y:11 P:67 SP:FB
	// D53D  A9 55     LDA #$55                        A:46 X:23 Y:11 P:67 SP:FB
	// D53F  8D 78 06  STA $0678 = 46                  A:55 X:23 Y:11 P:65 SP:FB
	// D542  24 01     BIT $01 = FF                    A:55 X:23 Y:11 P:65 SP:FB
	// D544  A9 11     LDA #$11                        A:55 X:23 Y:11 P:E5 SP:FB
	// D546  A2 23     LDX #$23                        A:11 X:23 Y:11 P:65 SP:FB
	// D548  A0 00     LDY #$00                        A:11 X:23 Y:11 P:65 SP:FB
	// D54A  AC 78 06  LDY $0678 = 55                  A:11 X:23 Y:00 P:67 SP:FB
	// D54D  F0 10     BEQ $D55F                       A:11 X:23 Y:55 P:65 SP:FB
	// D54F  30 0E     BMI $D55F                       A:11 X:23 Y:55 P:65 SP:FB
	// D551  C0 55     CPY #$55                        A:11 X:23 Y:55 P:65 SP:FB
	// D553  D0 0A     BNE $D55F                       A:11 X:23 Y:55 P:67 SP:FB
	// D555  C9 11     CMP #$11                        A:11 X:23 Y:55 P:67 SP:FB
	// D557  D0 06     BNE $D55F                       A:11 X:23 Y:55 P:67 SP:FB
	// D559  E0 23     CPX #$23                        A:11 X:23 Y:55 P:67 SP:FB
	// D55B  50 02     BVC $D55F                       A:11 X:23 Y:55 P:67 SP:FB
	// D55D  F0 04     BEQ $D563                       A:11 X:23 Y:55 P:67 SP:FB
	// D563  A0 46     LDY #$46                        A:11 X:23 Y:55 P:67 SP:FB
	// D565  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:65 SP:FB
	// D567  8C 78 06  STY $0678 = 55                  A:11 X:23 Y:46 P:E5 SP:FB
	// D56A  F0 0B     BEQ $D577                       A:11 X:23 Y:46 P:E5 SP:FB
	// D56C  10 09     BPL $D577                       A:11 X:23 Y:46 P:E5 SP:FB
	// D56E  50 07     BVC $D577                       A:11 X:23 Y:46 P:E5 SP:FB
	// D570  AC 78 06  LDY $0678 = 46                  A:11 X:23 Y:46 P:E5 SP:FB
	// D573  C0 46     CPY #$46                        A:11 X:23 Y:46 P:65 SP:FB
	// D575  F0 04     BEQ $D57B                       A:11 X:23 Y:46 P:67 SP:FB
	// D57B  24 01     BIT $01 = FF                    A:11 X:23 Y:46 P:67 SP:FB
	// D57D  A9 55     LDA #$55                        A:11 X:23 Y:46 P:E5 SP:FB
	// D57F  8D 78 06  STA $0678 = 46                  A:55 X:23 Y:46 P:65 SP:FB
	// D582  A0 11     LDY #$11                        A:55 X:23 Y:46 P:65 SP:FB
	// D584  A9 23     LDA #$23                        A:55 X:23 Y:11 P:65 SP:FB
	// D586  A2 00     LDX #$00                        A:23 X:23 Y:11 P:65 SP:FB
	// D588  AE 78 06  LDX $0678 = 55                  A:23 X:00 Y:11 P:67 SP:FB
	// D58B  F0 10     BEQ $D59D                       A:23 X:55 Y:11 P:65 SP:FB
	// D58D  30 0E     BMI $D59D                       A:23 X:55 Y:11 P:65 SP:FB
	// D58F  E0 55     CPX #$55                        A:23 X:55 Y:11 P:65 SP:FB
	// D591  D0 0A     BNE $D59D                       A:23 X:55 Y:11 P:67 SP:FB
	// D593  C0 11     CPY #$11                        A:23 X:55 Y:11 P:67 SP:FB
	// D595  D0 06     BNE $D59D                       A:23 X:55 Y:11 P:67 SP:FB
	// D597  C9 23     CMP #$23                        A:23 X:55 Y:11 P:67 SP:FB
	// D599  50 02     BVC $D59D                       A:23 X:55 Y:11 P:67 SP:FB
	// D59B  F0 04     BEQ $D5A1                       A:23 X:55 Y:11 P:67 SP:FB
	// D5A1  A2 46     LDX #$46                        A:23 X:55 Y:11 P:67 SP:FB
	// D5A3  24 01     BIT $01 = FF                    A:23 X:46 Y:11 P:65 SP:FB
	// D5A5  8E 78 06  STX $0678 = 55                  A:23 X:46 Y:11 P:E5 SP:FB
	// D5A8  F0 0B     BEQ $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB
	// D5AA  10 09     BPL $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB
	// D5AC  50 07     BVC $D5B5                       A:23 X:46 Y:11 P:E5 SP:FB
	// D5AE  AE 78 06  LDX $0678 = 46                  A:23 X:46 Y:11 P:E5 SP:FB
	// D5B1  E0 46     CPX #$46                        A:23 X:46 Y:11 P:65 SP:FB
	// D5B3  F0 04     BEQ $D5B9                       A:23 X:46 Y:11 P:67 SP:FB
	// D5B9  A9 C0     LDA #$C0                        A:23 X:46 Y:11 P:67 SP:FB
	// D5BB  8D 78 06  STA $0678 = 46                  A:C0 X:46 Y:11 P:E5 SP:FB
	// D5BE  A2 33     LDX #$33                        A:C0 X:46 Y:11 P:E5 SP:FB
	// D5C0  A0 88     LDY #$88                        A:C0 X:33 Y:11 P:65 SP:FB
	// D5C2  A9 05     LDA #$05                        A:C0 X:33 Y:88 P:E5 SP:FB
	// D5C4  2C 78 06  BIT $0678 = C0                  A:05 X:33 Y:88 P:65 SP:FB
	// D5C7  10 10     BPL $D5D9                       A:05 X:33 Y:88 P:E7 SP:FB
	// D5C9  50 0E     BVC $D5D9                       A:05 X:33 Y:88 P:E7 SP:FB
	// D5CB  D0 0C     BNE $D5D9                       A:05 X:33 Y:88 P:E7 SP:FB
	// D5CD  C9 05     CMP #$05                        A:05 X:33 Y:88 P:E7 SP:FB
	// D5CF  D0 08     BNE $D5D9                       A:05 X:33 Y:88 P:67 SP:FB
	// D5D1  E0 33     CPX #$33                        A:05 X:33 Y:88 P:67 SP:FB
	// D5D3  D0 04     BNE $D5D9                       A:05 X:33 Y:88 P:67 SP:FB
	// D5D5  C0 88     CPY #$88                        A:05 X:33 Y:88 P:67 SP:FB
	// D5D7  F0 04     BEQ $D5DD                       A:05 X:33 Y:88 P:67 SP:FB
	// D5DD  A9 03     LDA #$03                        A:05 X:33 Y:88 P:67 SP:FB
	// D5DF  8D 78 06  STA $0678 = C0                  A:03 X:33 Y:88 P:65 SP:FB
	// D5E2  A9 01     LDA #$01                        A:03 X:33 Y:88 P:65 SP:FB
	// D5E4  2C 78 06  BIT $0678 = 03                  A:01 X:33 Y:88 P:65 SP:FB
	// D5E7  30 08     BMI $D5F1                       A:01 X:33 Y:88 P:25 SP:FB
	// D5E9  70 06     BVS $D5F1                       A:01 X:33 Y:88 P:25 SP:FB
	// D5EB  F0 04     BEQ $D5F1                       A:01 X:33 Y:88 P:25 SP:FB
	// D5ED  C9 01     CMP #$01                        A:01 X:33 Y:88 P:25 SP:FB
	// D5EF  F0 04     BEQ $D5F5                       A:01 X:33 Y:88 P:27 SP:FB
	// D5F5  A0 B8     LDY #$B8                        A:01 X:33 Y:88 P:27 SP:FB
	// D5F7  A9 AA     LDA #$AA                        A:01 X:33 Y:B8 P:A5 SP:FB
	// D5F9  8D 78 06  STA $0678 = 03                  A:AA X:33 Y:B8 P:A5 SP:FB
	// D5FC  20 B6 F7  JSR $F7B6                       A:AA X:33 Y:B8 P:A5 SP:FB
	// F7B6  18        CLC                             A:AA X:33 Y:B8 P:A5 SP:F9
	// F7B7  A9 FF     LDA #$FF                        A:AA X:33 Y:B8 P:A4 SP:F9
	// F7B9  85 01     STA $01 = FF                    A:FF X:33 Y:B8 P:A4 SP:F9
	// F7BB  24 01     BIT $01 = FF                    A:FF X:33 Y:B8 P:A4 SP:F9
	// F7BD  A9 55     LDA #$55                        A:FF X:33 Y:B8 P:E4 SP:F9
	// F7BF  60        RTS                             A:55 X:33 Y:B8 P:64 SP:F9
	// D5FF  0D 78 06  ORA $0678 = AA                  A:55 X:33 Y:B8 P:64 SP:FB
	// D602  20 C0 F7  JSR $F7C0                       A:FF X:33 Y:B8 P:E4 SP:FB
	// F7C0  B0 09     BCS $F7CB                       A:FF X:33 Y:B8 P:E4 SP:F9
	// F7C2  10 07     BPL $F7CB                       A:FF X:33 Y:B8 P:E4 SP:F9
	// F7C4  C9 FF     CMP #$FF                        A:FF X:33 Y:B8 P:E4 SP:F9
	// F7C6  D0 03     BNE $F7CB                       A:FF X:33 Y:B8 P:67 SP:F9
	// F7C8  50 01     BVC $F7CB                       A:FF X:33 Y:B8 P:67 SP:F9
	// F7CA  60        RTS                             A:FF X:33 Y:B8 P:67 SP:F9
	// D605  C8        INY                             A:FF X:33 Y:B8 P:67 SP:FB
	// D606  A9 00     LDA #$00                        A:FF X:33 Y:B9 P:E5 SP:FB
	// D608  8D 78 06  STA $0678 = AA                  A:00 X:33 Y:B9 P:67 SP:FB
	// D60B  20 CE F7  JSR $F7CE                       A:00 X:33 Y:B9 P:67 SP:FB
	// F7CE  38        SEC                             A:00 X:33 Y:B9 P:67 SP:F9
	// F7CF  B8        CLV                             A:00 X:33 Y:B9 P:67 SP:F9
	// F7D0  A9 00     LDA #$00                        A:00 X:33 Y:B9 P:27 SP:F9
	// F7D2  60        RTS                             A:00 X:33 Y:B9 P:27 SP:F9
	// D60E  0D 78 06  ORA $0678 = 00                  A:00 X:33 Y:B9 P:27 SP:FB
	// D611  20 D3 F7  JSR $F7D3                       A:00 X:33 Y:B9 P:27 SP:FB
	// F7D3  D0 07     BNE $F7DC                       A:00 X:33 Y:B9 P:27 SP:F9
	// F7D5  70 05     BVS $F7DC                       A:00 X:33 Y:B9 P:27 SP:F9
	// F7D7  90 03     BCC $F7DC                       A:00 X:33 Y:B9 P:27 SP:F9
	// F7D9  30 01     BMI $F7DC                       A:00 X:33 Y:B9 P:27 SP:F9
	// F7DB  60        RTS                             A:00 X:33 Y:B9 P:27 SP:F9
	// D614  C8        INY                             A:00 X:33 Y:B9 P:27 SP:FB
	// D615  A9 AA     LDA #$AA                        A:00 X:33 Y:BA P:A5 SP:FB
	// D617  8D 78 06  STA $0678 = 00                  A:AA X:33 Y:BA P:A5 SP:FB
	// D61A  20 DF F7  JSR $F7DF                       A:AA X:33 Y:BA P:A5 SP:FB
	// F7DF  18        CLC                             A:AA X:33 Y:BA P:A5 SP:F9
	// F7E0  24 01     BIT $01 = FF                    A:AA X:33 Y:BA P:A4 SP:F9
	// F7E2  A9 55     LDA #$55                        A:AA X:33 Y:BA P:E4 SP:F9
	// F7E4  60        RTS                             A:55 X:33 Y:BA P:64 SP:F9
	// D61D  2D 78 06  AND $0678 = AA                  A:55 X:33 Y:BA P:64 SP:FB
	// D620  20 E5 F7  JSR $F7E5                       A:00 X:33 Y:BA P:66 SP:FB
	// F7E5  D0 07     BNE $F7EE                       A:00 X:33 Y:BA P:66 SP:F9
	// F7E7  50 05     BVC $F7EE                       A:00 X:33 Y:BA P:66 SP:F9
	// F7E9  B0 03     BCS $F7EE                       A:00 X:33 Y:BA P:66 SP:F9
	// F7EB  30 01     BMI $F7EE                       A:00 X:33 Y:BA P:66 SP:F9
	// F7ED  60        RTS                             A:00 X:33 Y:BA P:66 SP:F9
	// D623  C8        INY                             A:00 X:33 Y:BA P:66 SP:FB
	// D624  A9 EF     LDA #$EF                        A:00 X:33 Y:BB P:E4 SP:FB
	// D626  8D 78 06  STA $0678 = AA                  A:EF X:33 Y:BB P:E4 SP:FB
	// D629  20 F1 F7  JSR $F7F1                       A:EF X:33 Y:BB P:E4 SP:FB
	// F7F1  38        SEC                             A:EF X:33 Y:BB P:E4 SP:F9
	// F7F2  B8        CLV                             A:EF X:33 Y:BB P:E5 SP:F9
	// F7F3  A9 F8     LDA #$F8                        A:EF X:33 Y:BB P:A5 SP:F9
	// F7F5  60        RTS                             A:F8 X:33 Y:BB P:A5 SP:F9
	// D62C  2D 78 06  AND $0678 = EF                  A:F8 X:33 Y:BB P:A5 SP:FB
	// D62F  20 F6 F7  JSR $F7F6                       A:E8 X:33 Y:BB P:A5 SP:FB
	// F7F6  90 09     BCC $F801                       A:E8 X:33 Y:BB P:A5 SP:F9
	// F7F8  10 07     BPL $F801                       A:E8 X:33 Y:BB P:A5 SP:F9
	// F7FA  C9 E8     CMP #$E8                        A:E8 X:33 Y:BB P:A5 SP:F9
	// F7FC  D0 03     BNE $F801                       A:E8 X:33 Y:BB P:27 SP:F9
	// F7FE  70 01     BVS $F801                       A:E8 X:33 Y:BB P:27 SP:F9
	// F800  60        RTS                             A:E8 X:33 Y:BB P:27 SP:F9
	// D632  C8        INY                             A:E8 X:33 Y:BB P:27 SP:FB
	// D633  A9 AA     LDA #$AA                        A:E8 X:33 Y:BC P:A5 SP:FB
	// D635  8D 78 06  STA $0678 = EF                  A:AA X:33 Y:BC P:A5 SP:FB
	// D638  20 04 F8  JSR $F804                       A:AA X:33 Y:BC P:A5 SP:FB
	// F804  18        CLC                             A:AA X:33 Y:BC P:A5 SP:F9
	// F805  24 01     BIT $01 = FF                    A:AA X:33 Y:BC P:A4 SP:F9
	// F807  A9 5F     LDA #$5F                        A:AA X:33 Y:BC P:E4 SP:F9
	// F809  60        RTS                             A:5F X:33 Y:BC P:64 SP:F9
	// D63B  4D 78 06  EOR $0678 = AA                  A:5F X:33 Y:BC P:64 SP:FB
	// D63E  20 0A F8  JSR $F80A                       A:F5 X:33 Y:BC P:E4 SP:FB
	// F80A  B0 09     BCS $F815                       A:F5 X:33 Y:BC P:E4 SP:F9
	// F80C  10 07     BPL $F815                       A:F5 X:33 Y:BC P:E4 SP:F9
	// F80E  C9 F5     CMP #$F5                        A:F5 X:33 Y:BC P:E4 SP:F9
	// F810  D0 03     BNE $F815                       A:F5 X:33 Y:BC P:67 SP:F9
	// F812  50 01     BVC $F815                       A:F5 X:33 Y:BC P:67 SP:F9
	// F814  60        RTS                             A:F5 X:33 Y:BC P:67 SP:F9
	// D641  C8        INY                             A:F5 X:33 Y:BC P:67 SP:FB
	// D642  A9 70     LDA #$70                        A:F5 X:33 Y:BD P:E5 SP:FB
	// D644  8D 78 06  STA $0678 = AA                  A:70 X:33 Y:BD P:65 SP:FB
	// D647  20 18 F8  JSR $F818                       A:70 X:33 Y:BD P:65 SP:FB
	// F818  38        SEC                             A:70 X:33 Y:BD P:65 SP:F9
	// F819  B8        CLV                             A:70 X:33 Y:BD P:65 SP:F9
	// F81A  A9 70     LDA #$70                        A:70 X:33 Y:BD P:25 SP:F9
	// F81C  60        RTS                             A:70 X:33 Y:BD P:25 SP:F9
	// D64A  4D 78 06  EOR $0678 = 70                  A:70 X:33 Y:BD P:25 SP:FB
	// D64D  20 1D F8  JSR $F81D                       A:00 X:33 Y:BD P:27 SP:FB
	// F81D  D0 07     BNE $F826                       A:00 X:33 Y:BD P:27 SP:F9
	// F81F  70 05     BVS $F826                       A:00 X:33 Y:BD P:27 SP:F9
	// F821  90 03     BCC $F826                       A:00 X:33 Y:BD P:27 SP:F9
	// F823  30 01     BMI $F826                       A:00 X:33 Y:BD P:27 SP:F9
	// F825  60        RTS                             A:00 X:33 Y:BD P:27 SP:F9
	// D650  C8        INY                             A:00 X:33 Y:BD P:27 SP:FB
	// D651  A9 69     LDA #$69                        A:00 X:33 Y:BE P:A5 SP:FB
	// D653  8D 78 06  STA $0678 = 70                  A:69 X:33 Y:BE P:25 SP:FB
	// D656  20 29 F8  JSR $F829                       A:69 X:33 Y:BE P:25 SP:FB
	// F829  18        CLC                             A:69 X:33 Y:BE P:25 SP:F9
	// F82A  24 01     BIT $01 = FF                    A:69 X:33 Y:BE P:24 SP:F9
	// F82C  A9 00     LDA #$00                        A:69 X:33 Y:BE P:E4 SP:F9
	// F82E  60        RTS                             A:00 X:33 Y:BE P:66 SP:F9
	// D659  6D 78 06  ADC $0678 = 69                  A:00 X:33 Y:BE P:66 SP:FB
	// D65C  20 2F F8  JSR $F82F                       A:69 X:33 Y:BE P:24 SP:FB
	// F82F  30 09     BMI $F83A                       A:69 X:33 Y:BE P:24 SP:F9
	// F831  B0 07     BCS $F83A                       A:69 X:33 Y:BE P:24 SP:F9
	// F833  C9 69     CMP #$69                        A:69 X:33 Y:BE P:24 SP:F9
	// F835  D0 03     BNE $F83A                       A:69 X:33 Y:BE P:27 SP:F9
	// F837  70 01     BVS $F83A                       A:69 X:33 Y:BE P:27 SP:F9
	// F839  60        RTS                             A:69 X:33 Y:BE P:27 SP:F9
	// D65F  C8        INY                             A:69 X:33 Y:BE P:27 SP:FB
	// D660  20 3D F8  JSR $F83D                       A:69 X:33 Y:BF P:A5 SP:FB
	// F83D  38        SEC                             A:69 X:33 Y:BF P:A5 SP:F9
	// F83E  24 01     BIT $01 = FF                    A:69 X:33 Y:BF P:A5 SP:F9
	// F840  A9 00     LDA #$00                        A:69 X:33 Y:BF P:E5 SP:F9
	// F842  60        RTS                             A:00 X:33 Y:BF P:67 SP:F9
	// D663  6D 78 06  ADC $0678 = 69                  A:00 X:33 Y:BF P:67 SP:FB
	// D666  20 43 F8  JSR $F843                       A:6A X:33 Y:BF P:24 SP:FB
	// F843  30 09     BMI $F84E                       A:6A X:33 Y:BF P:24 SP:F9
	// F845  B0 07     BCS $F84E                       A:6A X:33 Y:BF P:24 SP:F9
	// F847  C9 6A     CMP #$6A                        A:6A X:33 Y:BF P:24 SP:F9
	// F849  D0 03     BNE $F84E                       A:6A X:33 Y:BF P:27 SP:F9
	// F84B  70 01     BVS $F84E                       A:6A X:33 Y:BF P:27 SP:F9
	// F84D  60        RTS                             A:6A X:33 Y:BF P:27 SP:F9
	// D669  C8        INY                             A:6A X:33 Y:BF P:27 SP:FB
	// D66A  A9 7F     LDA #$7F                        A:6A X:33 Y:C0 P:A5 SP:FB
	// D66C  8D 78 06  STA $0678 = 69                  A:7F X:33 Y:C0 P:25 SP:FB
	// D66F  20 51 F8  JSR $F851                       A:7F X:33 Y:C0 P:25 SP:FB
	// F851  38        SEC                             A:7F X:33 Y:C0 P:25 SP:F9
	// F852  B8        CLV                             A:7F X:33 Y:C0 P:25 SP:F9
	// F853  A9 7F     LDA #$7F                        A:7F X:33 Y:C0 P:25 SP:F9
	// F855  60        RTS                             A:7F X:33 Y:C0 P:25 SP:F9
	// D672  6D 78 06  ADC $0678 = 7F                  A:7F X:33 Y:C0 P:25 SP:FB
	// D675  20 56 F8  JSR $F856                       A:FF X:33 Y:C0 P:E4 SP:FB
	// F856  10 09     BPL $F861                       A:FF X:33 Y:C0 P:E4 SP:F9
	// F858  B0 07     BCS $F861                       A:FF X:33 Y:C0 P:E4 SP:F9
	// F85A  C9 FF     CMP #$FF                        A:FF X:33 Y:C0 P:E4 SP:F9
	// F85C  D0 03     BNE $F861                       A:FF X:33 Y:C0 P:67 SP:F9
	// F85E  50 01     BVC $F861                       A:FF X:33 Y:C0 P:67 SP:F9
	// F860  60        RTS                             A:FF X:33 Y:C0 P:67 SP:F9
	// D678  C8        INY                             A:FF X:33 Y:C0 P:67 SP:FB
	// D679  A9 80     LDA #$80                        A:FF X:33 Y:C1 P:E5 SP:FB
	// D67B  8D 78 06  STA $0678 = 7F                  A:80 X:33 Y:C1 P:E5 SP:FB
	// D67E  20 64 F8  JSR $F864                       A:80 X:33 Y:C1 P:E5 SP:FB
	// F864  18        CLC                             A:80 X:33 Y:C1 P:E5 SP:F9
	// F865  24 01     BIT $01 = FF                    A:80 X:33 Y:C1 P:E4 SP:F9
	// F867  A9 7F     LDA #$7F                        A:80 X:33 Y:C1 P:E4 SP:F9
	// F869  60        RTS                             A:7F X:33 Y:C1 P:64 SP:F9
	// D681  6D 78 06  ADC $0678 = 80                  A:7F X:33 Y:C1 P:64 SP:FB
	// D684  20 6A F8  JSR $F86A                       A:FF X:33 Y:C1 P:A4 SP:FB
	// F86A  10 09     BPL $F875                       A:FF X:33 Y:C1 P:A4 SP:F9
	// F86C  B0 07     BCS $F875                       A:FF X:33 Y:C1 P:A4 SP:F9
	// F86E  C9 FF     CMP #$FF                        A:FF X:33 Y:C1 P:A4 SP:F9
	// F870  D0 03     BNE $F875                       A:FF X:33 Y:C1 P:27 SP:F9
	// F872  70 01     BVS $F875                       A:FF X:33 Y:C1 P:27 SP:F9
	// F874  60        RTS                             A:FF X:33 Y:C1 P:27 SP:F9
	// D687  C8        INY                             A:FF X:33 Y:C1 P:27 SP:FB
	// D688  20 78 F8  JSR $F878                       A:FF X:33 Y:C2 P:A5 SP:FB
	// F878  38        SEC                             A:FF X:33 Y:C2 P:A5 SP:F9
	// F879  B8        CLV                             A:FF X:33 Y:C2 P:A5 SP:F9
	// F87A  A9 7F     LDA #$7F                        A:FF X:33 Y:C2 P:A5 SP:F9
	// F87C  60        RTS                             A:7F X:33 Y:C2 P:25 SP:F9
	// D68B  6D 78 06  ADC $0678 = 80                  A:7F X:33 Y:C2 P:25 SP:FB
	// D68E  20 7D F8  JSR $F87D                       A:00 X:33 Y:C2 P:27 SP:FB
	// F87D  D0 07     BNE $F886                       A:00 X:33 Y:C2 P:27 SP:F9
	// F87F  30 05     BMI $F886                       A:00 X:33 Y:C2 P:27 SP:F9
	// F881  70 03     BVS $F886                       A:00 X:33 Y:C2 P:27 SP:F9
	// F883  90 01     BCC $F886                       A:00 X:33 Y:C2 P:27 SP:F9
	// F885  60        RTS                             A:00 X:33 Y:C2 P:27 SP:F9
	// D691  C8        INY                             A:00 X:33 Y:C2 P:27 SP:FB
	// D692  A9 40     LDA #$40                        A:00 X:33 Y:C3 P:A5 SP:FB
	// D694  8D 78 06  STA $0678 = 80                  A:40 X:33 Y:C3 P:25 SP:FB
	// D697  20 89 F8  JSR $F889                       A:40 X:33 Y:C3 P:25 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:33 Y:C3 P:25 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:33 Y:C3 P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:33 Y:C3 P:65 SP:F9
	// D69A  CD 78 06  CMP $0678 = 40                  A:40 X:33 Y:C3 P:65 SP:FB
	// D69D  20 8E F8  JSR $F88E                       A:40 X:33 Y:C3 P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:33 Y:C3 P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:33 Y:C3 P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:33 Y:C3 P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:33 Y:C3 P:67 SP:F9
	// F896  60        RTS                             A:40 X:33 Y:C3 P:67 SP:F9
	// D6A0  C8        INY                             A:40 X:33 Y:C3 P:67 SP:FB
	// D6A1  48        PHA                             A:40 X:33 Y:C4 P:E5 SP:FB
	// D6A2  A9 3F     LDA #$3F                        A:40 X:33 Y:C4 P:E5 SP:FA
	// D6A4  8D 78 06  STA $0678 = 40                  A:3F X:33 Y:C4 P:65 SP:FA
	// D6A7  68        PLA                             A:3F X:33 Y:C4 P:65 SP:FA
	// D6A8  20 9A F8  JSR $F89A                       A:40 X:33 Y:C4 P:65 SP:FB
	// F89A  B8        CLV                             A:40 X:33 Y:C4 P:65 SP:F9
	// F89B  60        RTS                             A:40 X:33 Y:C4 P:25 SP:F9
	// D6AB  CD 78 06  CMP $0678 = 3F                  A:40 X:33 Y:C4 P:25 SP:FB
	// D6AE  20 9C F8  JSR $F89C                       A:40 X:33 Y:C4 P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:40 X:33 Y:C4 P:25 SP:F9
	// F8A4  60        RTS                             A:40 X:33 Y:C4 P:25 SP:F9
	// D6B1  C8        INY                             A:40 X:33 Y:C4 P:25 SP:FB
	// D6B2  48        PHA                             A:40 X:33 Y:C5 P:A5 SP:FB
	// D6B3  A9 41     LDA #$41                        A:40 X:33 Y:C5 P:A5 SP:FA
	// D6B5  8D 78 06  STA $0678 = 3F                  A:41 X:33 Y:C5 P:25 SP:FA
	// D6B8  68        PLA                             A:41 X:33 Y:C5 P:25 SP:FA
	// D6B9  CD 78 06  CMP $0678 = 41                  A:40 X:33 Y:C5 P:25 SP:FB
	// D6BC  20 A8 F8  JSR $F8A8                       A:40 X:33 Y:C5 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:40 X:33 Y:C5 P:A4 SP:F9
	// F8AE  60        RTS                             A:40 X:33 Y:C5 P:A4 SP:F9
	// D6BF  C8        INY                             A:40 X:33 Y:C5 P:A4 SP:FB
	// D6C0  48        PHA                             A:40 X:33 Y:C6 P:A4 SP:FB
	// D6C1  A9 00     LDA #$00                        A:40 X:33 Y:C6 P:A4 SP:FA
	// D6C3  8D 78 06  STA $0678 = 41                  A:00 X:33 Y:C6 P:26 SP:FA
	// D6C6  68        PLA                             A:00 X:33 Y:C6 P:26 SP:FA
	// D6C7  20 B2 F8  JSR $F8B2                       A:40 X:33 Y:C6 P:24 SP:FB
	// F8B2  A9 80     LDA #$80                        A:40 X:33 Y:C6 P:24 SP:F9
	// F8B4  60        RTS                             A:80 X:33 Y:C6 P:A4 SP:F9
	// D6CA  CD 78 06  CMP $0678 = 00                  A:80 X:33 Y:C6 P:A4 SP:FB
	// D6CD  20 B5 F8  JSR $F8B5                       A:80 X:33 Y:C6 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:33 Y:C6 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:33 Y:C6 P:A5 SP:F9
	// D6D0  C8        INY                             A:80 X:33 Y:C6 P:A5 SP:FB
	// D6D1  48        PHA                             A:80 X:33 Y:C7 P:A5 SP:FB
	// D6D2  A9 80     LDA #$80                        A:80 X:33 Y:C7 P:A5 SP:FA
	// D6D4  8D 78 06  STA $0678 = 00                  A:80 X:33 Y:C7 P:A5 SP:FA
	// D6D7  68        PLA                             A:80 X:33 Y:C7 P:A5 SP:FA
	// D6D8  CD 78 06  CMP $0678 = 80                  A:80 X:33 Y:C7 P:A5 SP:FB
	// D6DB  20 BF F8  JSR $F8BF                       A:80 X:33 Y:C7 P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:33 Y:C7 P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:33 Y:C7 P:27 SP:F9
	// D6DE  C8        INY                             A:80 X:33 Y:C7 P:27 SP:FB
	// D6DF  48        PHA                             A:80 X:33 Y:C8 P:A5 SP:FB
	// D6E0  A9 81     LDA #$81                        A:80 X:33 Y:C8 P:A5 SP:FA
	// D6E2  8D 78 06  STA $0678 = 80                  A:81 X:33 Y:C8 P:A5 SP:FA
	// D6E5  68        PLA                             A:81 X:33 Y:C8 P:A5 SP:FA
	// D6E6  CD 78 06  CMP $0678 = 81                  A:80 X:33 Y:C8 P:A5 SP:FB
	// D6E9  20 C9 F8  JSR $F8C9                       A:80 X:33 Y:C8 P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:80 X:33 Y:C8 P:A4 SP:F9
	// F8CF  60        RTS                             A:80 X:33 Y:C8 P:A4 SP:F9
	// D6EC  C8        INY                             A:80 X:33 Y:C8 P:A4 SP:FB
	// D6ED  48        PHA                             A:80 X:33 Y:C9 P:A4 SP:FB
	// D6EE  A9 7F     LDA #$7F                        A:80 X:33 Y:C9 P:A4 SP:FA
	// D6F0  8D 78 06  STA $0678 = 81                  A:7F X:33 Y:C9 P:24 SP:FA
	// D6F3  68        PLA                             A:7F X:33 Y:C9 P:24 SP:FA
	// D6F4  CD 78 06  CMP $0678 = 7F                  A:80 X:33 Y:C9 P:A4 SP:FB
	// D6F7  20 D3 F8  JSR $F8D3                       A:80 X:33 Y:C9 P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:80 X:33 Y:C9 P:25 SP:F9
	// F8D9  60        RTS                             A:80 X:33 Y:C9 P:25 SP:F9
	// D6FA  C8        INY                             A:80 X:33 Y:C9 P:25 SP:FB
	// D6FB  A9 40     LDA #$40                        A:80 X:33 Y:CA P:A5 SP:FB
	// D6FD  8D 78 06  STA $0678 = 7F                  A:40 X:33 Y:CA P:25 SP:FB
	// D700  20 31 F9  JSR $F931                       A:40 X:33 Y:CA P:25 SP:FB
	// F931  24 01     BIT $01 = FF                    A:40 X:33 Y:CA P:25 SP:F9
	// F933  A9 40     LDA #$40                        A:40 X:33 Y:CA P:E5 SP:F9
	// F935  38        SEC                             A:40 X:33 Y:CA P:65 SP:F9
	// F936  60        RTS                             A:40 X:33 Y:CA P:65 SP:F9
	// D703  ED 78 06  SBC $0678 = 40                  A:40 X:33 Y:CA P:65 SP:FB
	// D706  20 37 F9  JSR $F937                       A:00 X:33 Y:CA P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:33 Y:CA P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:33 Y:CA P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:33 Y:CA P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:33 Y:CA P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:33 Y:CA P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:33 Y:CA P:27 SP:F9
	// F943  60        RTS                             A:00 X:33 Y:CA P:27 SP:F9
	// D709  C8        INY                             A:00 X:33 Y:CA P:27 SP:FB
	// D70A  A9 3F     LDA #$3F                        A:00 X:33 Y:CB P:A5 SP:FB
	// D70C  8D 78 06  STA $0678 = 40                  A:3F X:33 Y:CB P:25 SP:FB
	// D70F  20 47 F9  JSR $F947                       A:3F X:33 Y:CB P:25 SP:FB
	// F947  B8        CLV                             A:3F X:33 Y:CB P:25 SP:F9
	// F948  38        SEC                             A:3F X:33 Y:CB P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:3F X:33 Y:CB P:25 SP:F9
	// F94B  60        RTS                             A:40 X:33 Y:CB P:25 SP:F9
	// D712  ED 78 06  SBC $0678 = 3F                  A:40 X:33 Y:CB P:25 SP:FB
	// D715  20 4C F9  JSR $F94C                       A:01 X:33 Y:CB P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:33 Y:CB P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:33 Y:CB P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:33 Y:CB P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:33 Y:CB P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:33 Y:CB P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:33 Y:CB P:27 SP:F9
	// F958  60        RTS                             A:01 X:33 Y:CB P:27 SP:F9
	// D718  C8        INY                             A:01 X:33 Y:CB P:27 SP:FB
	// D719  A9 41     LDA #$41                        A:01 X:33 Y:CC P:A5 SP:FB
	// D71B  8D 78 06  STA $0678 = 3F                  A:41 X:33 Y:CC P:25 SP:FB
	// D71E  20 5C F9  JSR $F95C                       A:41 X:33 Y:CC P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:41 X:33 Y:CC P:25 SP:F9
	// F95E  38        SEC                             A:40 X:33 Y:CC P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:33 Y:CC P:25 SP:F9
	// F961  60        RTS                             A:40 X:33 Y:CC P:E5 SP:F9
	// D721  ED 78 06  SBC $0678 = 41                  A:40 X:33 Y:CC P:E5 SP:FB
	// D724  20 62 F9  JSR $F962                       A:FF X:33 Y:CC P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:33 Y:CC P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:33 Y:CC P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:33 Y:CC P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:33 Y:CC P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:33 Y:CC P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:33 Y:CC P:27 SP:F9
	// F96E  60        RTS                             A:FF X:33 Y:CC P:27 SP:F9
	// D727  C8        INY                             A:FF X:33 Y:CC P:27 SP:FB
	// D728  A9 00     LDA #$00                        A:FF X:33 Y:CD P:A5 SP:FB
	// D72A  8D 78 06  STA $0678 = 41                  A:00 X:33 Y:CD P:27 SP:FB
	// D72D  20 72 F9  JSR $F972                       A:00 X:33 Y:CD P:27 SP:FB
	// F972  18        CLC                             A:00 X:33 Y:CD P:27 SP:F9
	// F973  A9 80     LDA #$80                        A:00 X:33 Y:CD P:26 SP:F9
	// F975  60        RTS                             A:80 X:33 Y:CD P:A4 SP:F9
	// D730  ED 78 06  SBC $0678 = 00                  A:80 X:33 Y:CD P:A4 SP:FB
	// D733  20 76 F9  JSR $F976                       A:7F X:33 Y:CD P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:33 Y:CD P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:33 Y:CD P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:33 Y:CD P:67 SP:F9
	// F97C  60        RTS                             A:7F X:33 Y:CD P:67 SP:F9
	// D736  C8        INY                             A:7F X:33 Y:CD P:67 SP:FB
	// D737  A9 7F     LDA #$7F                        A:7F X:33 Y:CE P:E5 SP:FB
	// D739  8D 78 06  STA $0678 = 00                  A:7F X:33 Y:CE P:65 SP:FB
	// D73C  20 80 F9  JSR $F980                       A:7F X:33 Y:CE P:65 SP:FB
	// F980  38        SEC                             A:7F X:33 Y:CE P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:33 Y:CE P:65 SP:F9
	// F983  60        RTS                             A:81 X:33 Y:CE P:E5 SP:F9
	// D73F  ED 78 06  SBC $0678 = 7F                  A:81 X:33 Y:CE P:E5 SP:FB
	// D742  20 84 F9  JSR $F984                       A:02 X:33 Y:CE P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:33 Y:CE P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:33 Y:CE P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:33 Y:CE P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:33 Y:CE P:67 SP:F9
	// F98C  60        RTS                             A:02 X:33 Y:CE P:67 SP:F9
	// D745  C8        INY                             A:02 X:33 Y:CE P:67 SP:FB
	// D746  A9 40     LDA #$40                        A:02 X:33 Y:CF P:E5 SP:FB
	// D748  8D 78 06  STA $0678 = 7F                  A:40 X:33 Y:CF P:65 SP:FB
	// D74B  20 89 F8  JSR $F889                       A:40 X:33 Y:CF P:65 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:33 Y:CF P:65 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:33 Y:CF P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:33 Y:CF P:65 SP:F9
	// D74E  AA        TAX                             A:40 X:33 Y:CF P:65 SP:FB
	// D74F  EC 78 06  CPX $0678 = 40                  A:40 X:40 Y:CF P:65 SP:FB
	// D752  20 8E F8  JSR $F88E                       A:40 X:40 Y:CF P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:40 Y:CF P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:40 Y:CF P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:40 Y:CF P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:40 Y:CF P:67 SP:F9
	// F896  60        RTS                             A:40 X:40 Y:CF P:67 SP:F9
	// D755  C8        INY                             A:40 X:40 Y:CF P:67 SP:FB
	// D756  A9 3F     LDA #$3F                        A:40 X:40 Y:D0 P:E5 SP:FB
	// D758  8D 78 06  STA $0678 = 40                  A:3F X:40 Y:D0 P:65 SP:FB
	// D75B  20 9A F8  JSR $F89A                       A:3F X:40 Y:D0 P:65 SP:FB
	// F89A  B8        CLV                             A:3F X:40 Y:D0 P:65 SP:F9
	// F89B  60        RTS                             A:3F X:40 Y:D0 P:25 SP:F9
	// D75E  EC 78 06  CPX $0678 = 3F                  A:3F X:40 Y:D0 P:25 SP:FB
	// D761  20 9C F8  JSR $F89C                       A:3F X:40 Y:D0 P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:3F X:40 Y:D0 P:25 SP:F9
	// F8A4  60        RTS                             A:3F X:40 Y:D0 P:25 SP:F9
	// D764  C8        INY                             A:3F X:40 Y:D0 P:25 SP:FB
	// D765  A9 41     LDA #$41                        A:3F X:40 Y:D1 P:A5 SP:FB
	// D767  8D 78 06  STA $0678 = 3F                  A:41 X:40 Y:D1 P:25 SP:FB
	// D76A  EC 78 06  CPX $0678 = 41                  A:41 X:40 Y:D1 P:25 SP:FB
	// D76D  20 A8 F8  JSR $F8A8                       A:41 X:40 Y:D1 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:41 X:40 Y:D1 P:A4 SP:F9
	// F8AE  60        RTS                             A:41 X:40 Y:D1 P:A4 SP:F9
	// D770  C8        INY                             A:41 X:40 Y:D1 P:A4 SP:FB
	// D771  A9 00     LDA #$00                        A:41 X:40 Y:D2 P:A4 SP:FB
	// D773  8D 78 06  STA $0678 = 41                  A:00 X:40 Y:D2 P:26 SP:FB
	// D776  20 B2 F8  JSR $F8B2                       A:00 X:40 Y:D2 P:26 SP:FB
	// F8B2  A9 80     LDA #$80                        A:00 X:40 Y:D2 P:26 SP:F9
	// F8B4  60        RTS                             A:80 X:40 Y:D2 P:A4 SP:F9
	// D779  AA        TAX                             A:80 X:40 Y:D2 P:A4 SP:FB
	// D77A  EC 78 06  CPX $0678 = 00                  A:80 X:80 Y:D2 P:A4 SP:FB
	// D77D  20 B5 F8  JSR $F8B5                       A:80 X:80 Y:D2 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:80 Y:D2 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:80 Y:D2 P:A5 SP:F9
	// D780  C8        INY                             A:80 X:80 Y:D2 P:A5 SP:FB
	// D781  A9 80     LDA #$80                        A:80 X:80 Y:D3 P:A5 SP:FB
	// D783  8D 78 06  STA $0678 = 00                  A:80 X:80 Y:D3 P:A5 SP:FB
	// D786  EC 78 06  CPX $0678 = 80                  A:80 X:80 Y:D3 P:A5 SP:FB
	// D789  20 BF F8  JSR $F8BF                       A:80 X:80 Y:D3 P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:80 Y:D3 P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:80 Y:D3 P:27 SP:F9
	// D78C  C8        INY                             A:80 X:80 Y:D3 P:27 SP:FB
	// D78D  A9 81     LDA #$81                        A:80 X:80 Y:D4 P:A5 SP:FB
	// D78F  8D 78 06  STA $0678 = 80                  A:81 X:80 Y:D4 P:A5 SP:FB
	// D792  EC 78 06  CPX $0678 = 81                  A:81 X:80 Y:D4 P:A5 SP:FB
	// D795  20 C9 F8  JSR $F8C9                       A:81 X:80 Y:D4 P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:81 X:80 Y:D4 P:A4 SP:F9
	// F8CF  60        RTS                             A:81 X:80 Y:D4 P:A4 SP:F9
	// D798  C8        INY                             A:81 X:80 Y:D4 P:A4 SP:FB
	// D799  A9 7F     LDA #$7F                        A:81 X:80 Y:D5 P:A4 SP:FB
	// D79B  8D 78 06  STA $0678 = 81                  A:7F X:80 Y:D5 P:24 SP:FB
	// D79E  EC 78 06  CPX $0678 = 7F                  A:7F X:80 Y:D5 P:24 SP:FB
	// D7A1  20 D3 F8  JSR $F8D3                       A:7F X:80 Y:D5 P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:7F X:80 Y:D5 P:25 SP:F9
	// F8D9  60        RTS                             A:7F X:80 Y:D5 P:25 SP:F9
	// D7A4  C8        INY                             A:7F X:80 Y:D5 P:25 SP:FB
	// D7A5  98        TYA                             A:7F X:80 Y:D6 P:A5 SP:FB
	// D7A6  AA        TAX                             A:D6 X:80 Y:D6 P:A5 SP:FB
	// D7A7  A9 40     LDA #$40                        A:D6 X:D6 Y:D6 P:A5 SP:FB
	// D7A9  8D 78 06  STA $0678 = 7F                  A:40 X:D6 Y:D6 P:25 SP:FB
	// D7AC  20 DD F8  JSR $F8DD                       A:40 X:D6 Y:D6 P:25 SP:FB
	// F8DD  24 01     BIT $01 = FF                    A:40 X:D6 Y:D6 P:25 SP:F9
	// F8DF  A0 40     LDY #$40                        A:40 X:D6 Y:D6 P:E5 SP:F9
	// F8E1  60        RTS                             A:40 X:D6 Y:40 P:65 SP:F9
	// D7AF  CC 78 06  CPY $0678 = 40                  A:40 X:D6 Y:40 P:65 SP:FB
	// D7B2  20 E2 F8  JSR $F8E2                       A:40 X:D6 Y:40 P:67 SP:FB
	// F8E2  30 07     BMI $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9
	// F8E4  90 05     BCC $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9
	// F8E6  D0 03     BNE $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9
	// F8E8  50 01     BVC $F8EB                       A:40 X:D6 Y:40 P:67 SP:F9
	// F8EA  60        RTS                             A:40 X:D6 Y:40 P:67 SP:F9
	// D7B5  E8        INX                             A:40 X:D6 Y:40 P:67 SP:FB
	// D7B6  A9 3F     LDA #$3F                        A:40 X:D7 Y:40 P:E5 SP:FB
	// D7B8  8D 78 06  STA $0678 = 40                  A:3F X:D7 Y:40 P:65 SP:FB
	// D7BB  20 EE F8  JSR $F8EE                       A:3F X:D7 Y:40 P:65 SP:FB
	// F8EE  B8        CLV                             A:3F X:D7 Y:40 P:65 SP:F9
	// F8EF  60        RTS                             A:3F X:D7 Y:40 P:25 SP:F9
	// D7BE  CC 78 06  CPY $0678 = 3F                  A:3F X:D7 Y:40 P:25 SP:FB
	// D7C1  20 F0 F8  JSR $F8F0                       A:3F X:D7 Y:40 P:25 SP:FB
	// F8F0  F0 07     BEQ $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9
	// F8F2  30 05     BMI $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9
	// F8F4  90 03     BCC $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9
	// F8F6  70 01     BVS $F8F9                       A:3F X:D7 Y:40 P:25 SP:F9
	// F8F8  60        RTS                             A:3F X:D7 Y:40 P:25 SP:F9
	// D7C4  E8        INX                             A:3F X:D7 Y:40 P:25 SP:FB
	// D7C5  A9 41     LDA #$41                        A:3F X:D8 Y:40 P:A5 SP:FB
	// D7C7  8D 78 06  STA $0678 = 3F                  A:41 X:D8 Y:40 P:25 SP:FB
	// D7CA  CC 78 06  CPY $0678 = 41                  A:41 X:D8 Y:40 P:25 SP:FB
	// D7CD  20 FC F8  JSR $F8FC                       A:41 X:D8 Y:40 P:A4 SP:FB
	// F8FC  F0 05     BEQ $F903                       A:41 X:D8 Y:40 P:A4 SP:F9
	// F8FE  10 03     BPL $F903                       A:41 X:D8 Y:40 P:A4 SP:F9
	// F900  10 01     BPL $F903                       A:41 X:D8 Y:40 P:A4 SP:F9
	// F902  60        RTS                             A:41 X:D8 Y:40 P:A4 SP:F9
	// D7D0  E8        INX                             A:41 X:D8 Y:40 P:A4 SP:FB
	// D7D1  A9 00     LDA #$00                        A:41 X:D9 Y:40 P:A4 SP:FB
	// D7D3  8D 78 06  STA $0678 = 41                  A:00 X:D9 Y:40 P:26 SP:FB
	// D7D6  20 06 F9  JSR $F906                       A:00 X:D9 Y:40 P:26 SP:FB
	// F906  A0 80     LDY #$80                        A:00 X:D9 Y:40 P:26 SP:F9
	// F908  60        RTS                             A:00 X:D9 Y:80 P:A4 SP:F9
	// D7D9  CC 78 06  CPY $0678 = 00                  A:00 X:D9 Y:80 P:A4 SP:FB
	// D7DC  20 09 F9  JSR $F909                       A:00 X:D9 Y:80 P:A5 SP:FB
	// F909  F0 05     BEQ $F910                       A:00 X:D9 Y:80 P:A5 SP:F9
	// F90B  10 03     BPL $F910                       A:00 X:D9 Y:80 P:A5 SP:F9
	// F90D  90 01     BCC $F910                       A:00 X:D9 Y:80 P:A5 SP:F9
	// F90F  60        RTS                             A:00 X:D9 Y:80 P:A5 SP:F9
	// D7DF  E8        INX                             A:00 X:D9 Y:80 P:A5 SP:FB
	// D7E0  A9 80     LDA #$80                        A:00 X:DA Y:80 P:A5 SP:FB
	// D7E2  8D 78 06  STA $0678 = 00                  A:80 X:DA Y:80 P:A5 SP:FB
	// D7E5  CC 78 06  CPY $0678 = 80                  A:80 X:DA Y:80 P:A5 SP:FB
	// D7E8  20 13 F9  JSR $F913                       A:80 X:DA Y:80 P:27 SP:FB
	// F913  D0 05     BNE $F91A                       A:80 X:DA Y:80 P:27 SP:F9
	// F915  30 03     BMI $F91A                       A:80 X:DA Y:80 P:27 SP:F9
	// F917  90 01     BCC $F91A                       A:80 X:DA Y:80 P:27 SP:F9
	// F919  60        RTS                             A:80 X:DA Y:80 P:27 SP:F9
	// D7EB  E8        INX                             A:80 X:DA Y:80 P:27 SP:FB
	// D7EC  A9 81     LDA #$81                        A:80 X:DB Y:80 P:A5 SP:FB
	// D7EE  8D 78 06  STA $0678 = 80                  A:81 X:DB Y:80 P:A5 SP:FB
	// D7F1  CC 78 06  CPY $0678 = 81                  A:81 X:DB Y:80 P:A5 SP:FB
	// D7F4  20 1D F9  JSR $F91D                       A:81 X:DB Y:80 P:A4 SP:FB
	// F91D  B0 05     BCS $F924                       A:81 X:DB Y:80 P:A4 SP:F9
	// F91F  F0 03     BEQ $F924                       A:81 X:DB Y:80 P:A4 SP:F9
	// F921  10 01     BPL $F924                       A:81 X:DB Y:80 P:A4 SP:F9
	// F923  60        RTS                             A:81 X:DB Y:80 P:A4 SP:F9
	// D7F7  E8        INX                             A:81 X:DB Y:80 P:A4 SP:FB
	// D7F8  A9 7F     LDA #$7F                        A:81 X:DC Y:80 P:A4 SP:FB
	// D7FA  8D 78 06  STA $0678 = 81                  A:7F X:DC Y:80 P:24 SP:FB
	// D7FD  CC 78 06  CPY $0678 = 7F                  A:7F X:DC Y:80 P:24 SP:FB
	// D800  20 27 F9  JSR $F927                       A:7F X:DC Y:80 P:25 SP:FB
	// F927  90 05     BCC $F92E                       A:7F X:DC Y:80 P:25 SP:F9
	// F929  F0 03     BEQ $F92E                       A:7F X:DC Y:80 P:25 SP:F9
	// F92B  30 01     BMI $F92E                       A:7F X:DC Y:80 P:25 SP:F9
	// F92D  60        RTS                             A:7F X:DC Y:80 P:25 SP:F9
	// D803  E8        INX                             A:7F X:DC Y:80 P:25 SP:FB
	// D804  8A        TXA                             A:7F X:DD Y:80 P:A5 SP:FB
	// D805  A8        TAY                             A:DD X:DD Y:80 P:A5 SP:FB
	// D806  20 90 F9  JSR $F990                       A:DD X:DD Y:DD P:A5 SP:FB
	// F990  A2 55     LDX #$55                        A:DD X:DD Y:DD P:A5 SP:F9
	// F992  A9 FF     LDA #$FF                        A:DD X:55 Y:DD P:25 SP:F9
	// F994  85 01     STA $01 = FF                    A:FF X:55 Y:DD P:A5 SP:F9
	// F996  EA        NOP                             A:FF X:55 Y:DD P:A5 SP:F9
	// F997  24 01     BIT $01 = FF                    A:FF X:55 Y:DD P:A5 SP:F9
	// F999  38        SEC                             A:FF X:55 Y:DD P:E5 SP:F9
	// F99A  A9 01     LDA #$01                        A:FF X:55 Y:DD P:E5 SP:F9
	// F99C  60        RTS                             A:01 X:55 Y:DD P:65 SP:F9
	// D809  8D 78 06  STA $0678 = 7F                  A:01 X:55 Y:DD P:65 SP:FB
	// D80C  4E 78 06  LSR $0678 = 01                  A:01 X:55 Y:DD P:65 SP:FB
	// D80F  AD 78 06  LDA $0678 = 00                  A:01 X:55 Y:DD P:67 SP:FB
	// D812  20 9D F9  JSR $F99D                       A:00 X:55 Y:DD P:67 SP:FB
	// F99D  90 1B     BCC $F9BA                       A:00 X:55 Y:DD P:67 SP:F9
	// F99F  D0 19     BNE $F9BA                       A:00 X:55 Y:DD P:67 SP:F9
	// F9A1  30 17     BMI $F9BA                       A:00 X:55 Y:DD P:67 SP:F9
	// F9A3  50 15     BVC $F9BA                       A:00 X:55 Y:DD P:67 SP:F9
	// F9A5  C9 00     CMP #$00                        A:00 X:55 Y:DD P:67 SP:F9
	// F9A7  D0 11     BNE $F9BA                       A:00 X:55 Y:DD P:67 SP:F9
	// F9A9  B8        CLV                             A:00 X:55 Y:DD P:67 SP:F9
	// F9AA  A9 AA     LDA #$AA                        A:00 X:55 Y:DD P:27 SP:F9
	// F9AC  60        RTS                             A:AA X:55 Y:DD P:A5 SP:F9
	// D815  C8        INY                             A:AA X:55 Y:DD P:A5 SP:FB
	// D816  8D 78 06  STA $0678 = 00                  A:AA X:55 Y:DE P:A5 SP:FB
	// D819  4E 78 06  LSR $0678 = AA                  A:AA X:55 Y:DE P:A5 SP:FB
	// D81C  AD 78 06  LDA $0678 = 55                  A:AA X:55 Y:DE P:24 SP:FB
	// D81F  20 AD F9  JSR $F9AD                       A:55 X:55 Y:DE P:24 SP:FB
	// F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:DE P:24 SP:F9
	// F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:DE P:24 SP:F9
	// F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:DE P:24 SP:F9
	// F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:DE P:24 SP:F9
	// F9B5  C9 55     CMP #$55                        A:55 X:55 Y:DE P:24 SP:F9
	// F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:DE P:27 SP:F9
	// F9B9  60        RTS                             A:55 X:55 Y:DE P:27 SP:F9
	// D822  C8        INY                             A:55 X:55 Y:DE P:27 SP:FB
	// D823  20 BD F9  JSR $F9BD                       A:55 X:55 Y:DF P:A5 SP:FB
	// F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:DF P:A5 SP:F9
	// F9BF  38        SEC                             A:55 X:55 Y:DF P:E5 SP:F9
	// F9C0  A9 80     LDA #$80                        A:55 X:55 Y:DF P:E5 SP:F9
	// F9C2  60        RTS                             A:80 X:55 Y:DF P:E5 SP:F9
	// D826  8D 78 06  STA $0678 = 55                  A:80 X:55 Y:DF P:E5 SP:FB
	// D829  0E 78 06  ASL $0678 = 80                  A:80 X:55 Y:DF P:E5 SP:FB
	// D82C  AD 78 06  LDA $0678 = 00                  A:80 X:55 Y:DF P:67 SP:FB
	// D82F  20 C3 F9  JSR $F9C3                       A:00 X:55 Y:DF P:67 SP:FB
	// F9C3  90 1C     BCC $F9E1                       A:00 X:55 Y:DF P:67 SP:F9
	// F9C5  D0 1A     BNE $F9E1                       A:00 X:55 Y:DF P:67 SP:F9
	// F9C7  30 18     BMI $F9E1                       A:00 X:55 Y:DF P:67 SP:F9
	// F9C9  50 16     BVC $F9E1                       A:00 X:55 Y:DF P:67 SP:F9
	// F9CB  C9 00     CMP #$00                        A:00 X:55 Y:DF P:67 SP:F9
	// F9CD  D0 12     BNE $F9E1                       A:00 X:55 Y:DF P:67 SP:F9
	// F9CF  B8        CLV                             A:00 X:55 Y:DF P:67 SP:F9
	// F9D0  A9 55     LDA #$55                        A:00 X:55 Y:DF P:27 SP:F9
	// F9D2  38        SEC                             A:55 X:55 Y:DF P:25 SP:F9
	// F9D3  60        RTS                             A:55 X:55 Y:DF P:25 SP:F9
	// D832  C8        INY                             A:55 X:55 Y:DF P:25 SP:FB
	// D833  8D 78 06  STA $0678 = 00                  A:55 X:55 Y:E0 P:A5 SP:FB
	// D836  0E 78 06  ASL $0678 = 55                  A:55 X:55 Y:E0 P:A5 SP:FB
	// D839  AD 78 06  LDA $0678 = AA                  A:55 X:55 Y:E0 P:A4 SP:FB
	// D83C  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:E0 P:A4 SP:FB
	// F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9
	// F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9
	// F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9
	// F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:E0 P:A4 SP:F9
	// F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:E0 P:A4 SP:F9
	// F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:E0 P:27 SP:F9
	// F9E0  60        RTS                             A:AA X:55 Y:E0 P:27 SP:F9
	// D83F  C8        INY                             A:AA X:55 Y:E0 P:27 SP:FB
	// D840  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:E1 P:A5 SP:FB
	// F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:E1 P:A5 SP:F9
	// F9E6  38        SEC                             A:AA X:55 Y:E1 P:E5 SP:F9
	// F9E7  A9 01     LDA #$01                        A:AA X:55 Y:E1 P:E5 SP:F9
	// F9E9  60        RTS                             A:01 X:55 Y:E1 P:65 SP:F9
	// D843  8D 78 06  STA $0678 = AA                  A:01 X:55 Y:E1 P:65 SP:FB
	// D846  6E 78 06  ROR $0678 = 01                  A:01 X:55 Y:E1 P:65 SP:FB
	// D849  AD 78 06  LDA $0678 = 80                  A:01 X:55 Y:E1 P:E5 SP:FB
	// D84C  20 EA F9  JSR $F9EA                       A:80 X:55 Y:E1 P:E5 SP:FB
	// F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9
	// F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9
	// F9EE  10 18     BPL $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9
	// F9F0  50 16     BVC $FA08                       A:80 X:55 Y:E1 P:E5 SP:F9
	// F9F2  C9 80     CMP #$80                        A:80 X:55 Y:E1 P:E5 SP:F9
	// F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:E1 P:67 SP:F9
	// F9F6  B8        CLV                             A:80 X:55 Y:E1 P:67 SP:F9
	// F9F7  18        CLC                             A:80 X:55 Y:E1 P:27 SP:F9
	// F9F8  A9 55     LDA #$55                        A:80 X:55 Y:E1 P:26 SP:F9
	// F9FA  60        RTS                             A:55 X:55 Y:E1 P:24 SP:F9
	// D84F  C8        INY                             A:55 X:55 Y:E1 P:24 SP:FB
	// D850  8D 78 06  STA $0678 = 80                  A:55 X:55 Y:E2 P:A4 SP:FB
	// D853  6E 78 06  ROR $0678 = 55                  A:55 X:55 Y:E2 P:A4 SP:FB
	// D856  AD 78 06  LDA $0678 = 2A                  A:55 X:55 Y:E2 P:25 SP:FB
	// D859  20 FB F9  JSR $F9FB                       A:2A X:55 Y:E2 P:25 SP:FB
	// F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:E2 P:25 SP:F9
	// F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:E2 P:25 SP:F9
	// F9FF  30 07     BMI $FA08                       A:2A X:55 Y:E2 P:25 SP:F9
	// FA01  70 05     BVS $FA08                       A:2A X:55 Y:E2 P:25 SP:F9
	// FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:E2 P:25 SP:F9
	// FA05  D0 01     BNE $FA08                       A:2A X:55 Y:E2 P:27 SP:F9
	// FA07  60        RTS                             A:2A X:55 Y:E2 P:27 SP:F9
	// D85C  C8        INY                             A:2A X:55 Y:E2 P:27 SP:FB
	// D85D  20 0A FA  JSR $FA0A                       A:2A X:55 Y:E3 P:A5 SP:FB
	// FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:E3 P:A5 SP:F9
	// FA0C  38        SEC                             A:2A X:55 Y:E3 P:E5 SP:F9
	// FA0D  A9 80     LDA #$80                        A:2A X:55 Y:E3 P:E5 SP:F9
	// FA0F  60        RTS                             A:80 X:55 Y:E3 P:E5 SP:F9
	// D860  8D 78 06  STA $0678 = 2A                  A:80 X:55 Y:E3 P:E5 SP:FB
	// D863  2E 78 06  ROL $0678 = 80                  A:80 X:55 Y:E3 P:E5 SP:FB
	// D866  AD 78 06  LDA $0678 = 01                  A:80 X:55 Y:E3 P:65 SP:FB
	// D869  20 10 FA  JSR $FA10                       A:01 X:55 Y:E3 P:65 SP:FB
	// FA10  90 1C     BCC $FA2E                       A:01 X:55 Y:E3 P:65 SP:F9
	// FA12  F0 1A     BEQ $FA2E                       A:01 X:55 Y:E3 P:65 SP:F9
	// FA14  30 18     BMI $FA2E                       A:01 X:55 Y:E3 P:65 SP:F9
	// FA16  50 16     BVC $FA2E                       A:01 X:55 Y:E3 P:65 SP:F9
	// FA18  C9 01     CMP #$01                        A:01 X:55 Y:E3 P:65 SP:F9
	// FA1A  D0 12     BNE $FA2E                       A:01 X:55 Y:E3 P:67 SP:F9
	// FA1C  B8        CLV                             A:01 X:55 Y:E3 P:67 SP:F9
	// FA1D  18        CLC                             A:01 X:55 Y:E3 P:27 SP:F9
	// FA1E  A9 55     LDA #$55                        A:01 X:55 Y:E3 P:26 SP:F9
	// FA20  60        RTS                             A:55 X:55 Y:E3 P:24 SP:F9
	// D86C  C8        INY                             A:55 X:55 Y:E3 P:24 SP:FB
	// D86D  8D 78 06  STA $0678 = 01                  A:55 X:55 Y:E4 P:A4 SP:FB
	// D870  2E 78 06  ROL $0678 = 55                  A:55 X:55 Y:E4 P:A4 SP:FB
	// D873  AD 78 06  LDA $0678 = AA                  A:55 X:55 Y:E4 P:A4 SP:FB
	// D876  20 21 FA  JSR $FA21                       A:AA X:55 Y:E4 P:A4 SP:FB
	// FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9
	// FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9
	// FA25  10 07     BPL $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9
	// FA27  70 05     BVS $FA2E                       A:AA X:55 Y:E4 P:A4 SP:F9
	// FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:E4 P:A4 SP:F9
	// FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:E4 P:27 SP:F9
	// FA2D  60        RTS                             A:AA X:55 Y:E4 P:27 SP:F9
	// D879  A9 FF     LDA #$FF                        A:AA X:55 Y:E4 P:27 SP:FB
	// D87B  8D 78 06  STA $0678 = AA                  A:FF X:55 Y:E4 P:A5 SP:FB
	// D87E  85 01     STA $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB
	// D880  24 01     BIT $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB
	// D882  38        SEC                             A:FF X:55 Y:E4 P:E5 SP:FB
	// D883  EE 78 06  INC $0678 = FF                  A:FF X:55 Y:E4 P:E5 SP:FB
	// D886  D0 0D     BNE $D895                       A:FF X:55 Y:E4 P:67 SP:FB
	// D888  30 0B     BMI $D895                       A:FF X:55 Y:E4 P:67 SP:FB
	// D88A  50 09     BVC $D895                       A:FF X:55 Y:E4 P:67 SP:FB
	// D88C  90 07     BCC $D895                       A:FF X:55 Y:E4 P:67 SP:FB
	// D88E  AD 78 06  LDA $0678 = 00                  A:FF X:55 Y:E4 P:67 SP:FB
	// D891  C9 00     CMP #$00                        A:00 X:55 Y:E4 P:67 SP:FB
	// D893  F0 04     BEQ $D899                       A:00 X:55 Y:E4 P:67 SP:FB
	// D899  A9 7F     LDA #$7F                        A:00 X:55 Y:E4 P:67 SP:FB
	// D89B  8D 78 06  STA $0678 = 00                  A:7F X:55 Y:E4 P:65 SP:FB
	// D89E  B8        CLV                             A:7F X:55 Y:E4 P:65 SP:FB
	// D89F  18        CLC                             A:7F X:55 Y:E4 P:25 SP:FB
	// D8A0  EE 78 06  INC $0678 = 7F                  A:7F X:55 Y:E4 P:24 SP:FB
	// D8A3  F0 0D     BEQ $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB
	// D8A5  10 0B     BPL $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB
	// D8A7  70 09     BVS $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB
	// D8A9  B0 07     BCS $D8B2                       A:7F X:55 Y:E4 P:A4 SP:FB
	// D8AB  AD 78 06  LDA $0678 = 80                  A:7F X:55 Y:E4 P:A4 SP:FB
	// D8AE  C9 80     CMP #$80                        A:80 X:55 Y:E4 P:A4 SP:FB
	// D8B0  F0 04     BEQ $D8B6                       A:80 X:55 Y:E4 P:27 SP:FB
	// D8B6  A9 00     LDA #$00                        A:80 X:55 Y:E4 P:27 SP:FB
	// D8B8  8D 78 06  STA $0678 = 80                  A:00 X:55 Y:E4 P:27 SP:FB
	// D8BB  24 01     BIT $01 = FF                    A:00 X:55 Y:E4 P:27 SP:FB
	// D8BD  38        SEC                             A:00 X:55 Y:E4 P:E7 SP:FB
	// D8BE  CE 78 06  DEC $0678 = 00                  A:00 X:55 Y:E4 P:E7 SP:FB
	// D8C1  F0 0D     BEQ $D8D0                       A:00 X:55 Y:E4 P:E5 SP:FB
	// D8C3  10 0B     BPL $D8D0                       A:00 X:55 Y:E4 P:E5 SP:FB
	// D8C5  50 09     BVC $D8D0                       A:00 X:55 Y:E4 P:E5 SP:FB
	// D8C7  90 07     BCC $D8D0                       A:00 X:55 Y:E4 P:E5 SP:FB
	// D8C9  AD 78 06  LDA $0678 = FF                  A:00 X:55 Y:E4 P:E5 SP:FB
	// D8CC  C9 FF     CMP #$FF                        A:FF X:55 Y:E4 P:E5 SP:FB
	// D8CE  F0 04     BEQ $D8D4                       A:FF X:55 Y:E4 P:67 SP:FB
	// D8D4  A9 80     LDA #$80                        A:FF X:55 Y:E4 P:67 SP:FB
	// D8D6  8D 78 06  STA $0678 = FF                  A:80 X:55 Y:E4 P:E5 SP:FB
	// D8D9  B8        CLV                             A:80 X:55 Y:E4 P:E5 SP:FB
	// D8DA  18        CLC                             A:80 X:55 Y:E4 P:A5 SP:FB
	// D8DB  CE 78 06  DEC $0678 = 80                  A:80 X:55 Y:E4 P:A4 SP:FB
	// D8DE  F0 0D     BEQ $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB
	// D8E0  30 0B     BMI $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB
	// D8E2  70 09     BVS $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB
	// D8E4  B0 07     BCS $D8ED                       A:80 X:55 Y:E4 P:24 SP:FB
	// D8E6  AD 78 06  LDA $0678 = 7F                  A:80 X:55 Y:E4 P:24 SP:FB
	// D8E9  C9 7F     CMP #$7F                        A:7F X:55 Y:E4 P:24 SP:FB
	// D8EB  F0 04     BEQ $D8F1                       A:7F X:55 Y:E4 P:27 SP:FB
	// D8F1  A9 01     LDA #$01                        A:7F X:55 Y:E4 P:27 SP:FB
	// D8F3  8D 78 06  STA $0678 = 7F                  A:01 X:55 Y:E4 P:25 SP:FB
	// D8F6  CE 78 06  DEC $0678 = 01                  A:01 X:55 Y:E4 P:25 SP:FB
	// D8F9  F0 04     BEQ $D8FF                       A:01 X:55 Y:E4 P:27 SP:FB
	// D8FF  60        RTS                             A:01 X:55 Y:E4 P:27 SP:FB
	// C618  20 00 D9  JSR $D900                       A:01 X:55 Y:E4 P:27 SP:FD
	// D900  A9 A3     LDA #$A3                        A:01 X:55 Y:E4 P:27 SP:FB
	// D902  85 33     STA $33 = 00                    A:A3 X:55 Y:E4 P:A5 SP:FB
	// D904  A9 89     LDA #$89                        A:A3 X:55 Y:E4 P:A5 SP:FB
	// D906  8D 00 03  STA $0300 = 70                  A:89 X:55 Y:E4 P:A5 SP:FB
	// D909  A9 12     LDA #$12                        A:89 X:55 Y:E4 P:A5 SP:FB
	// D90B  8D 45 02  STA $0245 = 00                  A:12 X:55 Y:E4 P:25 SP:FB
	// D90E  A9 FF     LDA #$FF                        A:12 X:55 Y:E4 P:25 SP:FB
	// D910  85 01     STA $01 = FF                    A:FF X:55 Y:E4 P:A5 SP:FB
	// D912  A2 65     LDX #$65                        A:FF X:55 Y:E4 P:A5 SP:FB
	// D914  A9 00     LDA #$00                        A:FF X:65 Y:E4 P:25 SP:FB
	// D916  85 89     STA $89 = 00                    A:00 X:65 Y:E4 P:27 SP:FB
	// D918  A9 03     LDA #$03                        A:00 X:65 Y:E4 P:27 SP:FB
	// D91A  85 8A     STA $8A = 00                    A:03 X:65 Y:E4 P:25 SP:FB
	// D91C  A0 00     LDY #$00                        A:03 X:65 Y:E4 P:25 SP:FB
	// D91E  38        SEC                             A:03 X:65 Y:00 P:27 SP:FB
	// D91F  A9 00     LDA #$00                        A:03 X:65 Y:00 P:27 SP:FB
	// D921  B8        CLV                             A:00 X:65 Y:00 P:27 SP:FB
	// D922  B1 89     LDA ($89),Y = 0300 @ 0300 = 89  A:00 X:65 Y:00 P:27 SP:FB
	// D924  F0 0C     BEQ $D932                       A:89 X:65 Y:00 P:A5 SP:FB
	// D926  90 0A     BCC $D932                       A:89 X:65 Y:00 P:A5 SP:FB
	// D928  70 08     BVS $D932                       A:89 X:65 Y:00 P:A5 SP:FB
	// D92A  C9 89     CMP #$89                        A:89 X:65 Y:00 P:A5 SP:FB
	// D92C  D0 04     BNE $D932                       A:89 X:65 Y:00 P:27 SP:FB
	// D92E  E0 65     CPX #$65                        A:89 X:65 Y:00 P:27 SP:FB
	// D930  F0 04     BEQ $D936                       A:89 X:65 Y:00 P:27 SP:FB
	// D936  A9 FF     LDA #$FF                        A:89 X:65 Y:00 P:27 SP:FB
	// D938  85 97     STA $97 = 00                    A:FF X:65 Y:00 P:A5 SP:FB
	// D93A  85 98     STA $98 = 00                    A:FF X:65 Y:00 P:A5 SP:FB
	// D93C  24 98     BIT $98 = FF                    A:FF X:65 Y:00 P:A5 SP:FB
	// D93E  A0 34     LDY #$34                        A:FF X:65 Y:00 P:E5 SP:FB
	// D940  B1 97     LDA ($97),Y = FFFF @ 0033 = A3  A:FF X:65 Y:34 P:65 SP:FB
	// D942  C9 A3     CMP #$A3                        A:A3 X:65 Y:34 P:E5 SP:FB
	// D944  D0 02     BNE $D948                       A:A3 X:65 Y:34 P:67 SP:FB
	// D946  B0 04     BCS $D94C                       A:A3 X:65 Y:34 P:67 SP:FB
	// D94C  A5 00     LDA $00 = 00                    A:A3 X:65 Y:34 P:67 SP:FB
	// D94E  48        PHA                             A:00 X:65 Y:34 P:67 SP:FB
	// D94F  A9 46     LDA #$46                        A:00 X:65 Y:34 P:67 SP:FA
	// D951  85 FF     STA $FF = 00                    A:46 X:65 Y:34 P:65 SP:FA
	// D953  A9 01     LDA #$01                        A:46 X:65 Y:34 P:65 SP:FA
	// D955  85 00     STA $00 = 00                    A:01 X:65 Y:34 P:65 SP:FA
	// D957  A0 FF     LDY #$FF                        A:01 X:65 Y:34 P:65 SP:FA
	// D959  B1 FF     LDA ($FF),Y = 0146 @ 0245 = 12  A:01 X:65 Y:FF P:E5 SP:FA
	// D95B  C9 12     CMP #$12                        A:12 X:65 Y:FF P:65 SP:FA
	// D95D  F0 04     BEQ $D963                       A:12 X:65 Y:FF P:67 SP:FA
	// D963  68        PLA                             A:12 X:65 Y:FF P:67 SP:FA
	// D964  85 00     STA $00 = 01                    A:00 X:65 Y:FF P:67 SP:FB
	// D966  A2 ED     LDX #$ED                        A:00 X:65 Y:FF P:67 SP:FB
	// D968  A9 00     LDA #$00                        A:00 X:ED Y:FF P:E5 SP:FB
	// D96A  85 33     STA $33 = A3                    A:00 X:ED Y:FF P:67 SP:FB
	// D96C  A9 04     LDA #$04                        A:00 X:ED Y:FF P:67 SP:FB
	// D96E  85 34     STA $34 = 00                    A:04 X:ED Y:FF P:65 SP:FB
	// D970  A0 00     LDY #$00                        A:04 X:ED Y:FF P:65 SP:FB
	// D972  18        CLC                             A:04 X:ED Y:00 P:67 SP:FB
	// D973  A9 FF     LDA #$FF                        A:04 X:ED Y:00 P:66 SP:FB
	// D975  85 01     STA $01 = FF                    A:FF X:ED Y:00 P:E4 SP:FB
	// D977  24 01     BIT $01 = FF                    A:FF X:ED Y:00 P:E4 SP:FB
	// D979  A9 AA     LDA #$AA                        A:FF X:ED Y:00 P:E4 SP:FB
	// D97B  8D 00 04  STA $0400 = AD                  A:AA X:ED Y:00 P:E4 SP:FB
	// D97E  A9 55     LDA #$55                        A:AA X:ED Y:00 P:E4 SP:FB
	// D980  11 33     ORA ($33),Y = 0400 @ 0400 = AA  A:55 X:ED Y:00 P:64 SP:FB
	// D982  B0 08     BCS $D98C                       A:FF X:ED Y:00 P:E4 SP:FB
	// D984  10 06     BPL $D98C                       A:FF X:ED Y:00 P:E4 SP:FB
	// D986  C9 FF     CMP #$FF                        A:FF X:ED Y:00 P:E4 SP:FB
	// D988  D0 02     BNE $D98C                       A:FF X:ED Y:00 P:67 SP:FB
	// D98A  70 02     BVS $D98E                       A:FF X:ED Y:00 P:67 SP:FB
	// D98E  E8        INX                             A:FF X:ED Y:00 P:67 SP:FB
	// D98F  38        SEC                             A:FF X:EE Y:00 P:E5 SP:FB
	// D990  B8        CLV                             A:FF X:EE Y:00 P:E5 SP:FB
	// D991  A9 00     LDA #$00                        A:FF X:EE Y:00 P:A5 SP:FB
	// D993  11 33     ORA ($33),Y = 0400 @ 0400 = AA  A:00 X:EE Y:00 P:27 SP:FB
	// D995  F0 06     BEQ $D99D                       A:AA X:EE Y:00 P:A5 SP:FB
	// D997  70 04     BVS $D99D                       A:AA X:EE Y:00 P:A5 SP:FB
	// D999  90 02     BCC $D99D                       A:AA X:EE Y:00 P:A5 SP:FB
	// D99B  30 02     BMI $D99F                       A:AA X:EE Y:00 P:A5 SP:FB
	// D99F  E8        INX                             A:AA X:EE Y:00 P:A5 SP:FB
	// D9A0  18        CLC                             A:AA X:EF Y:00 P:A5 SP:FB
	// D9A1  24 01     BIT $01 = FF                    A:AA X:EF Y:00 P:A4 SP:FB
	// D9A3  A9 55     LDA #$55                        A:AA X:EF Y:00 P:E4 SP:FB
	// D9A5  31 33     AND ($33),Y = 0400 @ 0400 = AA  A:55 X:EF Y:00 P:64 SP:FB
	// D9A7  D0 06     BNE $D9AF                       A:00 X:EF Y:00 P:66 SP:FB
	// D9A9  50 04     BVC $D9AF                       A:00 X:EF Y:00 P:66 SP:FB
	// D9AB  B0 02     BCS $D9AF                       A:00 X:EF Y:00 P:66 SP:FB
	// D9AD  10 02     BPL $D9B1                       A:00 X:EF Y:00 P:66 SP:FB
	// D9B1  E8        INX                             A:00 X:EF Y:00 P:66 SP:FB
	// D9B2  38        SEC                             A:00 X:F0 Y:00 P:E4 SP:FB
	// D9B3  B8        CLV                             A:00 X:F0 Y:00 P:E5 SP:FB
	// D9B4  A9 EF     LDA #$EF                        A:00 X:F0 Y:00 P:A5 SP:FB
	// D9B6  8D 00 04  STA $0400 = AA                  A:EF X:F0 Y:00 P:A5 SP:FB
	// D9B9  A9 F8     LDA #$F8                        A:EF X:F0 Y:00 P:A5 SP:FB
	// D9BB  31 33     AND ($33),Y = 0400 @ 0400 = EF  A:F8 X:F0 Y:00 P:A5 SP:FB
	// D9BD  90 08     BCC $D9C7                       A:E8 X:F0 Y:00 P:A5 SP:FB
	// D9BF  10 06     BPL $D9C7                       A:E8 X:F0 Y:00 P:A5 SP:FB
	// D9C1  C9 E8     CMP #$E8                        A:E8 X:F0 Y:00 P:A5 SP:FB
	// D9C3  D0 02     BNE $D9C7                       A:E8 X:F0 Y:00 P:27 SP:FB
	// D9C5  50 02     BVC $D9C9                       A:E8 X:F0 Y:00 P:27 SP:FB
	// D9C9  E8        INX                             A:E8 X:F0 Y:00 P:27 SP:FB
	// D9CA  18        CLC                             A:E8 X:F1 Y:00 P:A5 SP:FB
	// D9CB  24 01     BIT $01 = FF                    A:E8 X:F1 Y:00 P:A4 SP:FB
	// D9CD  A9 AA     LDA #$AA                        A:E8 X:F1 Y:00 P:E4 SP:FB
	// D9CF  8D 00 04  STA $0400 = EF                  A:AA X:F1 Y:00 P:E4 SP:FB
	// D9D2  A9 5F     LDA #$5F                        A:AA X:F1 Y:00 P:E4 SP:FB
	// D9D4  51 33     EOR ($33),Y = 0400 @ 0400 = AA  A:5F X:F1 Y:00 P:64 SP:FB
	// D9D6  B0 08     BCS $D9E0                       A:F5 X:F1 Y:00 P:E4 SP:FB
	// D9D8  10 06     BPL $D9E0                       A:F5 X:F1 Y:00 P:E4 SP:FB
	// D9DA  C9 F5     CMP #$F5                        A:F5 X:F1 Y:00 P:E4 SP:FB
	// D9DC  D0 02     BNE $D9E0                       A:F5 X:F1 Y:00 P:67 SP:FB
	// D9DE  70 02     BVS $D9E2                       A:F5 X:F1 Y:00 P:67 SP:FB
	// D9E2  E8        INX                             A:F5 X:F1 Y:00 P:67 SP:FB
	// D9E3  38        SEC                             A:F5 X:F2 Y:00 P:E5 SP:FB
	// D9E4  B8        CLV                             A:F5 X:F2 Y:00 P:E5 SP:FB
	// D9E5  A9 70     LDA #$70                        A:F5 X:F2 Y:00 P:A5 SP:FB
	// D9E7  8D 00 04  STA $0400 = AA                  A:70 X:F2 Y:00 P:25 SP:FB
	// D9EA  51 33     EOR ($33),Y = 0400 @ 0400 = 70  A:70 X:F2 Y:00 P:25 SP:FB
	// D9EC  D0 06     BNE $D9F4                       A:00 X:F2 Y:00 P:27 SP:FB
	// D9EE  70 04     BVS $D9F4                       A:00 X:F2 Y:00 P:27 SP:FB
	// D9F0  90 02     BCC $D9F4                       A:00 X:F2 Y:00 P:27 SP:FB
	// D9F2  10 02     BPL $D9F6                       A:00 X:F2 Y:00 P:27 SP:FB
	// D9F6  E8        INX                             A:00 X:F2 Y:00 P:27 SP:FB
	// D9F7  18        CLC                             A:00 X:F3 Y:00 P:A5 SP:FB
	// D9F8  24 01     BIT $01 = FF                    A:00 X:F3 Y:00 P:A4 SP:FB
	// D9FA  A9 69     LDA #$69                        A:00 X:F3 Y:00 P:E6 SP:FB
	// D9FC  8D 00 04  STA $0400 = 70                  A:69 X:F3 Y:00 P:64 SP:FB
	// D9FF  A9 00     LDA #$00                        A:69 X:F3 Y:00 P:64 SP:FB
	// DA01  71 33     ADC ($33),Y = 0400 @ 0400 = 69  A:00 X:F3 Y:00 P:66 SP:FB
	// DA03  30 08     BMI $DA0D                       A:69 X:F3 Y:00 P:24 SP:FB
	// DA05  B0 06     BCS $DA0D                       A:69 X:F3 Y:00 P:24 SP:FB
	// DA07  C9 69     CMP #$69                        A:69 X:F3 Y:00 P:24 SP:FB
	// DA09  D0 02     BNE $DA0D                       A:69 X:F3 Y:00 P:27 SP:FB
	// DA0B  50 02     BVC $DA0F                       A:69 X:F3 Y:00 P:27 SP:FB
	// DA0F  E8        INX                             A:69 X:F3 Y:00 P:27 SP:FB
	// DA10  38        SEC                             A:69 X:F4 Y:00 P:A5 SP:FB
	// DA11  24 01     BIT $01 = FF                    A:69 X:F4 Y:00 P:A5 SP:FB
	// DA13  A9 00     LDA #$00                        A:69 X:F4 Y:00 P:E5 SP:FB
	// DA15  71 33     ADC ($33),Y = 0400 @ 0400 = 69  A:00 X:F4 Y:00 P:67 SP:FB
	// DA17  30 08     BMI $DA21                       A:6A X:F4 Y:00 P:24 SP:FB
	// DA19  B0 06     BCS $DA21                       A:6A X:F4 Y:00 P:24 SP:FB
	// DA1B  C9 6A     CMP #$6A                        A:6A X:F4 Y:00 P:24 SP:FB
	// DA1D  D0 02     BNE $DA21                       A:6A X:F4 Y:00 P:27 SP:FB
	// DA1F  50 02     BVC $DA23                       A:6A X:F4 Y:00 P:27 SP:FB
	// DA23  E8        INX                             A:6A X:F4 Y:00 P:27 SP:FB
	// DA24  38        SEC                             A:6A X:F5 Y:00 P:A5 SP:FB
	// DA25  B8        CLV                             A:6A X:F5 Y:00 P:A5 SP:FB
	// DA26  A9 7F     LDA #$7F                        A:6A X:F5 Y:00 P:A5 SP:FB
	// DA28  8D 00 04  STA $0400 = 69                  A:7F X:F5 Y:00 P:25 SP:FB
	// DA2B  71 33     ADC ($33),Y = 0400 @ 0400 = 7F  A:7F X:F5 Y:00 P:25 SP:FB
	// DA2D  10 08     BPL $DA37                       A:FF X:F5 Y:00 P:E4 SP:FB
	// DA2F  B0 06     BCS $DA37                       A:FF X:F5 Y:00 P:E4 SP:FB
	// DA31  C9 FF     CMP #$FF                        A:FF X:F5 Y:00 P:E4 SP:FB
	// DA33  D0 02     BNE $DA37                       A:FF X:F5 Y:00 P:67 SP:FB
	// DA35  70 02     BVS $DA39                       A:FF X:F5 Y:00 P:67 SP:FB
	// DA39  E8        INX                             A:FF X:F5 Y:00 P:67 SP:FB
	// DA3A  18        CLC                             A:FF X:F6 Y:00 P:E5 SP:FB
	// DA3B  24 01     BIT $01 = FF                    A:FF X:F6 Y:00 P:E4 SP:FB
	// DA3D  A9 80     LDA #$80                        A:FF X:F6 Y:00 P:E4 SP:FB
	// DA3F  8D 00 04  STA $0400 = 7F                  A:80 X:F6 Y:00 P:E4 SP:FB
	// DA42  A9 7F     LDA #$7F                        A:80 X:F6 Y:00 P:E4 SP:FB
	// DA44  71 33     ADC ($33),Y = 0400 @ 0400 = 80  A:7F X:F6 Y:00 P:64 SP:FB
	// DA46  10 08     BPL $DA50                       A:FF X:F6 Y:00 P:A4 SP:FB
	// DA48  B0 06     BCS $DA50                       A:FF X:F6 Y:00 P:A4 SP:FB
	// DA4A  C9 FF     CMP #$FF                        A:FF X:F6 Y:00 P:A4 SP:FB
	// DA4C  D0 02     BNE $DA50                       A:FF X:F6 Y:00 P:27 SP:FB
	// DA4E  50 02     BVC $DA52                       A:FF X:F6 Y:00 P:27 SP:FB
	// DA52  E8        INX                             A:FF X:F6 Y:00 P:27 SP:FB
	// DA53  38        SEC                             A:FF X:F7 Y:00 P:A5 SP:FB
	// DA54  B8        CLV                             A:FF X:F7 Y:00 P:A5 SP:FB
	// DA55  A9 80     LDA #$80                        A:FF X:F7 Y:00 P:A5 SP:FB
	// DA57  8D 00 04  STA $0400 = 80                  A:80 X:F7 Y:00 P:A5 SP:FB
	// DA5A  A9 7F     LDA #$7F                        A:80 X:F7 Y:00 P:A5 SP:FB
	// DA5C  71 33     ADC ($33),Y = 0400 @ 0400 = 80  A:7F X:F7 Y:00 P:25 SP:FB
	// DA5E  D0 06     BNE $DA66                       A:00 X:F7 Y:00 P:27 SP:FB
	// DA60  30 04     BMI $DA66                       A:00 X:F7 Y:00 P:27 SP:FB
	// DA62  70 02     BVS $DA66                       A:00 X:F7 Y:00 P:27 SP:FB
	// DA64  B0 02     BCS $DA68                       A:00 X:F7 Y:00 P:27 SP:FB
	// DA68  E8        INX                             A:00 X:F7 Y:00 P:27 SP:FB
	// DA69  24 01     BIT $01 = FF                    A:00 X:F8 Y:00 P:A5 SP:FB
	// DA6B  A9 40     LDA #$40                        A:00 X:F8 Y:00 P:E7 SP:FB
	// DA6D  8D 00 04  STA $0400 = 80                  A:40 X:F8 Y:00 P:65 SP:FB
	// DA70  D1 33     CMP ($33),Y = 0400 @ 0400 = 40  A:40 X:F8 Y:00 P:65 SP:FB
	// DA72  30 06     BMI $DA7A                       A:40 X:F8 Y:00 P:67 SP:FB
	// DA74  90 04     BCC $DA7A                       A:40 X:F8 Y:00 P:67 SP:FB
	// DA76  D0 02     BNE $DA7A                       A:40 X:F8 Y:00 P:67 SP:FB
	// DA78  70 02     BVS $DA7C                       A:40 X:F8 Y:00 P:67 SP:FB
	// DA7C  E8        INX                             A:40 X:F8 Y:00 P:67 SP:FB
	// DA7D  B8        CLV                             A:40 X:F9 Y:00 P:E5 SP:FB
	// DA7E  CE 00 04  DEC $0400 = 40                  A:40 X:F9 Y:00 P:A5 SP:FB
	// DA81  D1 33     CMP ($33),Y = 0400 @ 0400 = 3F  A:40 X:F9 Y:00 P:25 SP:FB
	// DA83  F0 06     BEQ $DA8B                       A:40 X:F9 Y:00 P:25 SP:FB
	// DA85  30 04     BMI $DA8B                       A:40 X:F9 Y:00 P:25 SP:FB
	// DA87  90 02     BCC $DA8B                       A:40 X:F9 Y:00 P:25 SP:FB
	// DA89  50 02     BVC $DA8D                       A:40 X:F9 Y:00 P:25 SP:FB
	// DA8D  E8        INX                             A:40 X:F9 Y:00 P:25 SP:FB
	// DA8E  EE 00 04  INC $0400 = 3F                  A:40 X:FA Y:00 P:A5 SP:FB
	// DA91  EE 00 04  INC $0400 = 40                  A:40 X:FA Y:00 P:25 SP:FB
	// DA94  D1 33     CMP ($33),Y = 0400 @ 0400 = 41  A:40 X:FA Y:00 P:25 SP:FB
	// DA96  F0 02     BEQ $DA9A                       A:40 X:FA Y:00 P:A4 SP:FB
	// DA98  30 02     BMI $DA9C                       A:40 X:FA Y:00 P:A4 SP:FB
	// DA9C  E8        INX                             A:40 X:FA Y:00 P:A4 SP:FB
	// DA9D  A9 00     LDA #$00                        A:40 X:FB Y:00 P:A4 SP:FB
	// DA9F  8D 00 04  STA $0400 = 41                  A:00 X:FB Y:00 P:26 SP:FB
	// DAA2  A9 80     LDA #$80                        A:00 X:FB Y:00 P:26 SP:FB
	// DAA4  D1 33     CMP ($33),Y = 0400 @ 0400 = 00  A:80 X:FB Y:00 P:A4 SP:FB
	// DAA6  F0 04     BEQ $DAAC                       A:80 X:FB Y:00 P:A5 SP:FB
	// DAA8  10 02     BPL $DAAC                       A:80 X:FB Y:00 P:A5 SP:FB
	// DAAA  B0 02     BCS $DAAE                       A:80 X:FB Y:00 P:A5 SP:FB
	// DAAE  E8        INX                             A:80 X:FB Y:00 P:A5 SP:FB
	// DAAF  A0 80     LDY #$80                        A:80 X:FC Y:00 P:A5 SP:FB
	// DAB1  8C 00 04  STY $0400 = 00                  A:80 X:FC Y:80 P:A5 SP:FB
	// DAB4  A0 00     LDY #$00                        A:80 X:FC Y:80 P:A5 SP:FB
	// DAB6  D1 33     CMP ($33),Y = 0400 @ 0400 = 80  A:80 X:FC Y:00 P:27 SP:FB
	// DAB8  D0 04     BNE $DABE                       A:80 X:FC Y:00 P:27 SP:FB
	// DABA  30 02     BMI $DABE                       A:80 X:FC Y:00 P:27 SP:FB
	// DABC  B0 02     BCS $DAC0                       A:80 X:FC Y:00 P:27 SP:FB
	// DAC0  E8        INX                             A:80 X:FC Y:00 P:27 SP:FB
	// DAC1  EE 00 04  INC $0400 = 80                  A:80 X:FD Y:00 P:A5 SP:FB
	// DAC4  D1 33     CMP ($33),Y = 0400 @ 0400 = 81  A:80 X:FD Y:00 P:A5 SP:FB
	// DAC6  B0 04     BCS $DACC                       A:80 X:FD Y:00 P:A4 SP:FB
	// DAC8  F0 02     BEQ $DACC                       A:80 X:FD Y:00 P:A4 SP:FB
	// DACA  30 02     BMI $DACE                       A:80 X:FD Y:00 P:A4 SP:FB
	// DACE  E8        INX                             A:80 X:FD Y:00 P:A4 SP:FB
	// DACF  CE 00 04  DEC $0400 = 81                  A:80 X:FE Y:00 P:A4 SP:FB
	// DAD2  CE 00 04  DEC $0400 = 80                  A:80 X:FE Y:00 P:A4 SP:FB
	// DAD5  D1 33     CMP ($33),Y = 0400 @ 0400 = 7F  A:80 X:FE Y:00 P:24 SP:FB
	// DAD7  90 04     BCC $DADD                       A:80 X:FE Y:00 P:25 SP:FB
	// DAD9  F0 02     BEQ $DADD                       A:80 X:FE Y:00 P:25 SP:FB
	// DADB  10 02     BPL $DADF                       A:80 X:FE Y:00 P:25 SP:FB
	// DADF  60        RTS                             A:80 X:FE Y:00 P:25 SP:FB
	// C61B  A5 00     LDA $00 = 00                    A:80 X:FE Y:00 P:25 SP:FD
	// C61D  85 10     STA $10 = 00                    A:00 X:FE Y:00 P:27 SP:FD
	// C61F  A9 00     LDA #$00                        A:00 X:FE Y:00 P:27 SP:FD
	// C621  85 00     STA $00 = 00                    A:00 X:FE Y:00 P:27 SP:FD
	// C623  20 E0 DA  JSR $DAE0                       A:00 X:FE Y:00 P:27 SP:FD
	// DAE0  A9 00     LDA #$00                        A:00 X:FE Y:00 P:27 SP:FB
	// DAE2  85 33     STA $33 = 00                    A:00 X:FE Y:00 P:27 SP:FB
	// DAE4  A9 04     LDA #$04                        A:00 X:FE Y:00 P:27 SP:FB
	// DAE6  85 34     STA $34 = 04                    A:04 X:FE Y:00 P:25 SP:FB
	// DAE8  A0 00     LDY #$00                        A:04 X:FE Y:00 P:25 SP:FB
	// DAEA  A2 01     LDX #$01                        A:04 X:FE Y:00 P:27 SP:FB
	// DAEC  24 01     BIT $01 = FF                    A:04 X:01 Y:00 P:25 SP:FB
	// DAEE  A9 40     LDA #$40                        A:04 X:01 Y:00 P:E5 SP:FB
	// DAF0  8D 00 04  STA $0400 = 7F                  A:40 X:01 Y:00 P:65 SP:FB
	// DAF3  38        SEC                             A:40 X:01 Y:00 P:65 SP:FB
	// DAF4  F1 33     SBC ($33),Y = 0400 @ 0400 = 40  A:40 X:01 Y:00 P:65 SP:FB
	// DAF6  30 0A     BMI $DB02                       A:00 X:01 Y:00 P:27 SP:FB
	// DAF8  90 08     BCC $DB02                       A:00 X:01 Y:00 P:27 SP:FB
	// DAFA  D0 06     BNE $DB02                       A:00 X:01 Y:00 P:27 SP:FB
	// DAFC  70 04     BVS $DB02                       A:00 X:01 Y:00 P:27 SP:FB
	// DAFE  C9 00     CMP #$00                        A:00 X:01 Y:00 P:27 SP:FB
	// DB00  F0 02     BEQ $DB04                       A:00 X:01 Y:00 P:27 SP:FB
	// DB04  E8        INX                             A:00 X:01 Y:00 P:27 SP:FB
	// DB05  B8        CLV                             A:00 X:02 Y:00 P:25 SP:FB
	// DB06  38        SEC                             A:00 X:02 Y:00 P:25 SP:FB
	// DB07  A9 40     LDA #$40                        A:00 X:02 Y:00 P:25 SP:FB
	// DB09  CE 00 04  DEC $0400 = 40                  A:40 X:02 Y:00 P:25 SP:FB
	// DB0C  F1 33     SBC ($33),Y = 0400 @ 0400 = 3F  A:40 X:02 Y:00 P:25 SP:FB
	// DB0E  F0 0A     BEQ $DB1A                       A:01 X:02 Y:00 P:25 SP:FB
	// DB10  30 08     BMI $DB1A                       A:01 X:02 Y:00 P:25 SP:FB
	// DB12  90 06     BCC $DB1A                       A:01 X:02 Y:00 P:25 SP:FB
	// DB14  70 04     BVS $DB1A                       A:01 X:02 Y:00 P:25 SP:FB
	// DB16  C9 01     CMP #$01                        A:01 X:02 Y:00 P:25 SP:FB
	// DB18  F0 02     BEQ $DB1C                       A:01 X:02 Y:00 P:27 SP:FB
	// DB1C  E8        INX                             A:01 X:02 Y:00 P:27 SP:FB
	// DB1D  A9 40     LDA #$40                        A:01 X:03 Y:00 P:25 SP:FB
	// DB1F  38        SEC                             A:40 X:03 Y:00 P:25 SP:FB
	// DB20  24 01     BIT $01 = FF                    A:40 X:03 Y:00 P:25 SP:FB
	// DB22  EE 00 04  INC $0400 = 3F                  A:40 X:03 Y:00 P:E5 SP:FB
	// DB25  EE 00 04  INC $0400 = 40                  A:40 X:03 Y:00 P:65 SP:FB
	// DB28  F1 33     SBC ($33),Y = 0400 @ 0400 = 41  A:40 X:03 Y:00 P:65 SP:FB
	// DB2A  B0 0A     BCS $DB36                       A:FF X:03 Y:00 P:A4 SP:FB
	// DB2C  F0 08     BEQ $DB36                       A:FF X:03 Y:00 P:A4 SP:FB
	// DB2E  10 06     BPL $DB36                       A:FF X:03 Y:00 P:A4 SP:FB
	// DB30  70 04     BVS $DB36                       A:FF X:03 Y:00 P:A4 SP:FB
	// DB32  C9 FF     CMP #$FF                        A:FF X:03 Y:00 P:A4 SP:FB
	// DB34  F0 02     BEQ $DB38                       A:FF X:03 Y:00 P:27 SP:FB
	// DB38  E8        INX                             A:FF X:03 Y:00 P:27 SP:FB
	// DB39  18        CLC                             A:FF X:04 Y:00 P:25 SP:FB
	// DB3A  A9 00     LDA #$00                        A:FF X:04 Y:00 P:24 SP:FB
	// DB3C  8D 00 04  STA $0400 = 41                  A:00 X:04 Y:00 P:26 SP:FB
	// DB3F  A9 80     LDA #$80                        A:00 X:04 Y:00 P:26 SP:FB
	// DB41  F1 33     SBC ($33),Y = 0400 @ 0400 = 00  A:80 X:04 Y:00 P:A4 SP:FB
	// DB43  90 04     BCC $DB49                       A:7F X:04 Y:00 P:65 SP:FB
	// DB45  C9 7F     CMP #$7F                        A:7F X:04 Y:00 P:65 SP:FB
	// DB47  F0 02     BEQ $DB4B                       A:7F X:04 Y:00 P:67 SP:FB
	// DB4B  E8        INX                             A:7F X:04 Y:00 P:67 SP:FB
	// DB4C  38        SEC                             A:7F X:05 Y:00 P:65 SP:FB
	// DB4D  A9 7F     LDA #$7F                        A:7F X:05 Y:00 P:65 SP:FB
	// DB4F  8D 00 04  STA $0400 = 00                  A:7F X:05 Y:00 P:65 SP:FB
	// DB52  A9 81     LDA #$81                        A:7F X:05 Y:00 P:65 SP:FB
	// DB54  F1 33     SBC ($33),Y = 0400 @ 0400 = 7F  A:81 X:05 Y:00 P:E5 SP:FB
	// DB56  50 06     BVC $DB5E                       A:02 X:05 Y:00 P:65 SP:FB
	// DB58  90 04     BCC $DB5E                       A:02 X:05 Y:00 P:65 SP:FB
	// DB5A  C9 02     CMP #$02                        A:02 X:05 Y:00 P:65 SP:FB
	// DB5C  F0 02     BEQ $DB60                       A:02 X:05 Y:00 P:67 SP:FB
	// DB60  E8        INX                             A:02 X:05 Y:00 P:67 SP:FB
	// DB61  A9 00     LDA #$00                        A:02 X:06 Y:00 P:65 SP:FB
	// DB63  A9 87     LDA #$87                        A:00 X:06 Y:00 P:67 SP:FB
	// DB65  91 33     STA ($33),Y = 0400 @ 0400 = 7F  A:87 X:06 Y:00 P:E5 SP:FB
	// DB67  AD 00 04  LDA $0400 = 87                  A:87 X:06 Y:00 P:E5 SP:FB
	// DB6A  C9 87     CMP #$87                        A:87 X:06 Y:00 P:E5 SP:FB
	// DB6C  F0 02     BEQ $DB70                       A:87 X:06 Y:00 P:67 SP:FB
	// DB70  E8        INX                             A:87 X:06 Y:00 P:67 SP:FB
	// DB71  A9 7E     LDA #$7E                        A:87 X:07 Y:00 P:65 SP:FB
	// DB73  8D 00 02  STA $0200 = 7F                  A:7E X:07 Y:00 P:65 SP:FB
	// DB76  A9 DB     LDA #$DB                        A:7E X:07 Y:00 P:65 SP:FB
	// DB78  8D 01 02  STA $0201 = 00                  A:DB X:07 Y:00 P:E5 SP:FB
	// DB7B  6C 00 02  JMP ($0200) = DB7E              A:DB X:07 Y:00 P:E5 SP:FB
	// DB7E  A9 00     LDA #$00                        A:DB X:07 Y:00 P:E5 SP:FB
	// DB80  8D FF 02  STA $02FF = 00                  A:00 X:07 Y:00 P:67 SP:FB
	// DB83  A9 01     LDA #$01                        A:00 X:07 Y:00 P:67 SP:FB
	// DB85  8D 00 03  STA $0300 = 89                  A:01 X:07 Y:00 P:65 SP:FB
	// DB88  A9 03     LDA #$03                        A:01 X:07 Y:00 P:65 SP:FB
	// DB8A  8D 00 02  STA $0200 = 7E                  A:03 X:07 Y:00 P:65 SP:FB
	// DB8D  A9 A9     LDA #$A9                        A:03 X:07 Y:00 P:65 SP:FB
	// DB8F  8D 00 01  STA $0100 = 00                  A:A9 X:07 Y:00 P:E5 SP:FB
	// DB92  A9 55     LDA #$55                        A:A9 X:07 Y:00 P:E5 SP:FB
	// DB94  8D 01 01  STA $0101 = 00                  A:55 X:07 Y:00 P:65 SP:FB
	// DB97  A9 60     LDA #$60                        A:55 X:07 Y:00 P:65 SP:FB
	// DB99  8D 02 01  STA $0102 = 00                  A:60 X:07 Y:00 P:65 SP:FB
	// DB9C  A9 A9     LDA #$A9                        A:60 X:07 Y:00 P:65 SP:FB
	// DB9E  8D 00 03  STA $0300 = 01                  A:A9 X:07 Y:00 P:E5 SP:FB
	// DBA1  A9 AA     LDA #$AA                        A:A9 X:07 Y:00 P:E5 SP:FB
	// DBA3  8D 01 03  STA $0301 = 00                  A:AA X:07 Y:00 P:E5 SP:FB
	// DBA6  A9 60     LDA #$60                        A:AA X:07 Y:00 P:E5 SP:FB
	// DBA8  8D 02 03  STA $0302 = 00                  A:60 X:07 Y:00 P:65 SP:FB
	// DBAB  20 B5 DB  JSR $DBB5                       A:60 X:07 Y:00 P:65 SP:FB
	// DBB5  6C FF 02  JMP ($02FF) = A900              A:60 X:07 Y:00 P:65 SP:F9
	// 0300  A9 AA     LDA #$AA                        A:60 X:07 Y:00 P:65 SP:F9
	// 0302  60        RTS                             A:AA X:07 Y:00 P:E5 SP:F9
	// DBAE  C9 AA     CMP #$AA                        A:AA X:07 Y:00 P:E5 SP:FB
	// DBB0  F0 02     BEQ $DBB4                       A:AA X:07 Y:00 P:67 SP:FB
	// DBB4  60        RTS                             A:AA X:07 Y:00 P:67 SP:FB
	// C626  20 4A DF  JSR $DF4A                       A:AA X:07 Y:00 P:67 SP:FD
	// DF4A  A9 89     LDA #$89                        A:AA X:07 Y:00 P:67 SP:FB
	// DF4C  8D 00 03  STA $0300 = A9                  A:89 X:07 Y:00 P:E5 SP:FB
	// DF4F  A9 A3     LDA #$A3                        A:89 X:07 Y:00 P:E5 SP:FB
	// DF51  85 33     STA $33 = 00                    A:A3 X:07 Y:00 P:E5 SP:FB
	// DF53  A9 12     LDA #$12                        A:A3 X:07 Y:00 P:E5 SP:FB
	// DF55  8D 45 02  STA $0245 = 12                  A:12 X:07 Y:00 P:65 SP:FB
	// DF58  A2 65     LDX #$65                        A:12 X:07 Y:00 P:65 SP:FB
	// DF5A  A0 00     LDY #$00                        A:12 X:65 Y:00 P:65 SP:FB
	// DF5C  38        SEC                             A:12 X:65 Y:00 P:67 SP:FB
	// DF5D  A9 00     LDA #$00                        A:12 X:65 Y:00 P:67 SP:FB
	// DF5F  B8        CLV                             A:00 X:65 Y:00 P:67 SP:FB
	// DF60  B9 00 03  LDA $0300,Y @ 0300 = 89         A:00 X:65 Y:00 P:27 SP:FB
	// DF63  F0 0C     BEQ $DF71                       A:89 X:65 Y:00 P:A5 SP:FB
	// DF65  90 0A     BCC $DF71                       A:89 X:65 Y:00 P:A5 SP:FB
	// DF67  70 08     BVS $DF71                       A:89 X:65 Y:00 P:A5 SP:FB
	// DF69  C9 89     CMP #$89                        A:89 X:65 Y:00 P:A5 SP:FB
	// DF6B  D0 04     BNE $DF71                       A:89 X:65 Y:00 P:27 SP:FB
	// DF6D  E0 65     CPX #$65                        A:89 X:65 Y:00 P:27 SP:FB
	// DF6F  F0 04     BEQ $DF75                       A:89 X:65 Y:00 P:27 SP:FB
	// DF75  A9 FF     LDA #$FF                        A:89 X:65 Y:00 P:27 SP:FB
	// DF77  85 01     STA $01 = FF                    A:FF X:65 Y:00 P:A5 SP:FB
	// DF79  24 01     BIT $01 = FF                    A:FF X:65 Y:00 P:A5 SP:FB
	// DF7B  A0 34     LDY #$34                        A:FF X:65 Y:00 P:E5 SP:FB
	// DF7D  B9 FF FF  LDA $FFFF,Y @ 0033 = A3         A:FF X:65 Y:34 P:65 SP:FB
	// DF80  C9 A3     CMP #$A3                        A:A3 X:65 Y:34 P:E5 SP:FB
	// DF82  D0 02     BNE $DF86                       A:A3 X:65 Y:34 P:67 SP:FB
	// DF84  B0 04     BCS $DF8A                       A:A3 X:65 Y:34 P:67 SP:FB
	// DF8A  A9 46     LDA #$46                        A:A3 X:65 Y:34 P:67 SP:FB
	// DF8C  85 FF     STA $FF = 46                    A:46 X:65 Y:34 P:65 SP:FB
	// DF8E  A0 FF     LDY #$FF                        A:46 X:65 Y:34 P:65 SP:FB
	// DF90  B9 46 01  LDA $0146,Y @ 0245 = 12         A:46 X:65 Y:FF P:E5 SP:FB
	// DF93  C9 12     CMP #$12                        A:12 X:65 Y:FF P:65 SP:FB
	// DF95  F0 04     BEQ $DF9B                       A:12 X:65 Y:FF P:67 SP:FB
	// DF9B  A2 39     LDX #$39                        A:12 X:65 Y:FF P:67 SP:FB
	// DF9D  18        CLC                             A:12 X:39 Y:FF P:65 SP:FB
	// DF9E  A9 FF     LDA #$FF                        A:12 X:39 Y:FF P:64 SP:FB
	// DFA0  85 01     STA $01 = FF                    A:FF X:39 Y:FF P:E4 SP:FB
	// DFA2  24 01     BIT $01 = FF                    A:FF X:39 Y:FF P:E4 SP:FB
	// DFA4  A9 AA     LDA #$AA                        A:FF X:39 Y:FF P:E4 SP:FB
	// DFA6  8D 00 04  STA $0400 = 87                  A:AA X:39 Y:FF P:E4 SP:FB
	// DFA9  A9 55     LDA #$55                        A:AA X:39 Y:FF P:E4 SP:FB
	// DFAB  A0 00     LDY #$00                        A:55 X:39 Y:FF P:64 SP:FB
	// DFAD  19 00 04  ORA $0400,Y @ 0400 = AA         A:55 X:39 Y:00 P:66 SP:FB
	// DFB0  B0 08     BCS $DFBA                       A:FF X:39 Y:00 P:E4 SP:FB
	// DFB2  10 06     BPL $DFBA                       A:FF X:39 Y:00 P:E4 SP:FB
	// DFB4  C9 FF     CMP #$FF                        A:FF X:39 Y:00 P:E4 SP:FB
	// DFB6  D0 02     BNE $DFBA                       A:FF X:39 Y:00 P:67 SP:FB
	// DFB8  70 02     BVS $DFBC                       A:FF X:39 Y:00 P:67 SP:FB
	// DFBC  E8        INX                             A:FF X:39 Y:00 P:67 SP:FB
	// DFBD  38        SEC                             A:FF X:3A Y:00 P:65 SP:FB
	// DFBE  B8        CLV                             A:FF X:3A Y:00 P:65 SP:FB
	// DFBF  A9 00     LDA #$00                        A:FF X:3A Y:00 P:25 SP:FB
	// DFC1  19 00 04  ORA $0400,Y @ 0400 = AA         A:00 X:3A Y:00 P:27 SP:FB
	// DFC4  F0 06     BEQ $DFCC                       A:AA X:3A Y:00 P:A5 SP:FB
	// DFC6  70 04     BVS $DFCC                       A:AA X:3A Y:00 P:A5 SP:FB
	// DFC8  90 02     BCC $DFCC                       A:AA X:3A Y:00 P:A5 SP:FB
	// DFCA  30 02     BMI $DFCE                       A:AA X:3A Y:00 P:A5 SP:FB
	// DFCE  E8        INX                             A:AA X:3A Y:00 P:A5 SP:FB
	// DFCF  18        CLC                             A:AA X:3B Y:00 P:25 SP:FB
	// DFD0  24 01     BIT $01 = FF                    A:AA X:3B Y:00 P:24 SP:FB
	// DFD2  A9 55     LDA #$55                        A:AA X:3B Y:00 P:E4 SP:FB
	// DFD4  39 00 04  AND $0400,Y @ 0400 = AA         A:55 X:3B Y:00 P:64 SP:FB
	// DFD7  D0 06     BNE $DFDF                       A:00 X:3B Y:00 P:66 SP:FB
	// DFD9  50 04     BVC $DFDF                       A:00 X:3B Y:00 P:66 SP:FB
	// DFDB  B0 02     BCS $DFDF                       A:00 X:3B Y:00 P:66 SP:FB
	// DFDD  10 02     BPL $DFE1                       A:00 X:3B Y:00 P:66 SP:FB
	// DFE1  E8        INX                             A:00 X:3B Y:00 P:66 SP:FB
	// DFE2  38        SEC                             A:00 X:3C Y:00 P:64 SP:FB
	// DFE3  B8        CLV                             A:00 X:3C Y:00 P:65 SP:FB
	// DFE4  A9 EF     LDA #$EF                        A:00 X:3C Y:00 P:25 SP:FB
	// DFE6  8D 00 04  STA $0400 = AA                  A:EF X:3C Y:00 P:A5 SP:FB
	// DFE9  A9 F8     LDA #$F8                        A:EF X:3C Y:00 P:A5 SP:FB
	// DFEB  39 00 04  AND $0400,Y @ 0400 = EF         A:F8 X:3C Y:00 P:A5 SP:FB
	// DFEE  90 08     BCC $DFF8                       A:E8 X:3C Y:00 P:A5 SP:FB
	// DFF0  10 06     BPL $DFF8                       A:E8 X:3C Y:00 P:A5 SP:FB
	// DFF2  C9 E8     CMP #$E8                        A:E8 X:3C Y:00 P:A5 SP:FB
	// DFF4  D0 02     BNE $DFF8                       A:E8 X:3C Y:00 P:27 SP:FB
	// DFF6  50 02     BVC $DFFA                       A:E8 X:3C Y:00 P:27 SP:FB
	// DFFA  E8        INX                             A:E8 X:3C Y:00 P:27 SP:FB
	// DFFB  18        CLC                             A:E8 X:3D Y:00 P:25 SP:FB
	// DFFC  24 01     BIT $01 = FF                    A:E8 X:3D Y:00 P:24 SP:FB
	// DFFE  A9 AA     LDA #$AA                        A:E8 X:3D Y:00 P:E4 SP:FB
	// E000  8D 00 04  STA $0400 = EF                  A:AA X:3D Y:00 P:E4 SP:FB
	// E003  A9 5F     LDA #$5F                        A:AA X:3D Y:00 P:E4 SP:FB
	// E005  59 00 04  EOR $0400,Y @ 0400 = AA         A:5F X:3D Y:00 P:64 SP:FB
	// E008  B0 08     BCS $E012                       A:F5 X:3D Y:00 P:E4 SP:FB
	// E00A  10 06     BPL $E012                       A:F5 X:3D Y:00 P:E4 SP:FB
	// E00C  C9 F5     CMP #$F5                        A:F5 X:3D Y:00 P:E4 SP:FB
	// E00E  D0 02     BNE $E012                       A:F5 X:3D Y:00 P:67 SP:FB
	// E010  70 02     BVS $E014                       A:F5 X:3D Y:00 P:67 SP:FB
	// E014  E8        INX                             A:F5 X:3D Y:00 P:67 SP:FB
	// E015  38        SEC                             A:F5 X:3E Y:00 P:65 SP:FB
	// E016  B8        CLV                             A:F5 X:3E Y:00 P:65 SP:FB
	// E017  A9 70     LDA #$70                        A:F5 X:3E Y:00 P:25 SP:FB
	// E019  8D 00 04  STA $0400 = AA                  A:70 X:3E Y:00 P:25 SP:FB
	// E01C  59 00 04  EOR $0400,Y @ 0400 = 70         A:70 X:3E Y:00 P:25 SP:FB
	// E01F  D0 06     BNE $E027                       A:00 X:3E Y:00 P:27 SP:FB
	// E021  70 04     BVS $E027                       A:00 X:3E Y:00 P:27 SP:FB
	// E023  90 02     BCC $E027                       A:00 X:3E Y:00 P:27 SP:FB
	// E025  10 02     BPL $E029                       A:00 X:3E Y:00 P:27 SP:FB
	// E029  E8        INX                             A:00 X:3E Y:00 P:27 SP:FB
	// E02A  18        CLC                             A:00 X:3F Y:00 P:25 SP:FB
	// E02B  24 01     BIT $01 = FF                    A:00 X:3F Y:00 P:24 SP:FB
	// E02D  A9 69     LDA #$69                        A:00 X:3F Y:00 P:E6 SP:FB
	// E02F  8D 00 04  STA $0400 = 70                  A:69 X:3F Y:00 P:64 SP:FB
	// E032  A9 00     LDA #$00                        A:69 X:3F Y:00 P:64 SP:FB
	// E034  79 00 04  ADC $0400,Y @ 0400 = 69         A:00 X:3F Y:00 P:66 SP:FB
	// E037  30 08     BMI $E041                       A:69 X:3F Y:00 P:24 SP:FB
	// E039  B0 06     BCS $E041                       A:69 X:3F Y:00 P:24 SP:FB
	// E03B  C9 69     CMP #$69                        A:69 X:3F Y:00 P:24 SP:FB
	// E03D  D0 02     BNE $E041                       A:69 X:3F Y:00 P:27 SP:FB
	// E03F  50 02     BVC $E043                       A:69 X:3F Y:00 P:27 SP:FB
	// E043  E8        INX                             A:69 X:3F Y:00 P:27 SP:FB
	// E044  38        SEC                             A:69 X:40 Y:00 P:25 SP:FB
	// E045  24 01     BIT $01 = FF                    A:69 X:40 Y:00 P:25 SP:FB
	// E047  A9 00     LDA #$00                        A:69 X:40 Y:00 P:E5 SP:FB
	// E049  79 00 04  ADC $0400,Y @ 0400 = 69         A:00 X:40 Y:00 P:67 SP:FB
	// E04C  30 08     BMI $E056                       A:6A X:40 Y:00 P:24 SP:FB
	// E04E  B0 06     BCS $E056                       A:6A X:40 Y:00 P:24 SP:FB
	// E050  C9 6A     CMP #$6A                        A:6A X:40 Y:00 P:24 SP:FB
	// E052  D0 02     BNE $E056                       A:6A X:40 Y:00 P:27 SP:FB
	// E054  50 02     BVC $E058                       A:6A X:40 Y:00 P:27 SP:FB
	// E058  E8        INX                             A:6A X:40 Y:00 P:27 SP:FB
	// E059  38        SEC                             A:6A X:41 Y:00 P:25 SP:FB
	// E05A  B8        CLV                             A:6A X:41 Y:00 P:25 SP:FB
	// E05B  A9 7F     LDA #$7F                        A:6A X:41 Y:00 P:25 SP:FB
	// E05D  8D 00 04  STA $0400 = 69                  A:7F X:41 Y:00 P:25 SP:FB
	// E060  79 00 04  ADC $0400,Y @ 0400 = 7F         A:7F X:41 Y:00 P:25 SP:FB
	// E063  10 08     BPL $E06D                       A:FF X:41 Y:00 P:E4 SP:FB
	// E065  B0 06     BCS $E06D                       A:FF X:41 Y:00 P:E4 SP:FB
	// E067  C9 FF     CMP #$FF                        A:FF X:41 Y:00 P:E4 SP:FB
	// E069  D0 02     BNE $E06D                       A:FF X:41 Y:00 P:67 SP:FB
	// E06B  70 02     BVS $E06F                       A:FF X:41 Y:00 P:67 SP:FB
	// E06F  E8        INX                             A:FF X:41 Y:00 P:67 SP:FB
	// E070  18        CLC                             A:FF X:42 Y:00 P:65 SP:FB
	// E071  24 01     BIT $01 = FF                    A:FF X:42 Y:00 P:64 SP:FB
	// E073  A9 80     LDA #$80                        A:FF X:42 Y:00 P:E4 SP:FB
	// E075  8D 00 04  STA $0400 = 7F                  A:80 X:42 Y:00 P:E4 SP:FB
	// E078  A9 7F     LDA #$7F                        A:80 X:42 Y:00 P:E4 SP:FB
	// E07A  79 00 04  ADC $0400,Y @ 0400 = 80         A:7F X:42 Y:00 P:64 SP:FB
	// E07D  10 08     BPL $E087                       A:FF X:42 Y:00 P:A4 SP:FB
	// E07F  B0 06     BCS $E087                       A:FF X:42 Y:00 P:A4 SP:FB
	// E081  C9 FF     CMP #$FF                        A:FF X:42 Y:00 P:A4 SP:FB
	// E083  D0 02     BNE $E087                       A:FF X:42 Y:00 P:27 SP:FB
	// E085  50 02     BVC $E089                       A:FF X:42 Y:00 P:27 SP:FB
	// E089  E8        INX                             A:FF X:42 Y:00 P:27 SP:FB
	// E08A  38        SEC                             A:FF X:43 Y:00 P:25 SP:FB
	// E08B  B8        CLV                             A:FF X:43 Y:00 P:25 SP:FB
	// E08C  A9 80     LDA #$80                        A:FF X:43 Y:00 P:25 SP:FB
	// E08E  8D 00 04  STA $0400 = 80                  A:80 X:43 Y:00 P:A5 SP:FB
	// E091  A9 7F     LDA #$7F                        A:80 X:43 Y:00 P:A5 SP:FB
	// E093  79 00 04  ADC $0400,Y @ 0400 = 80         A:7F X:43 Y:00 P:25 SP:FB
	// E096  D0 06     BNE $E09E                       A:00 X:43 Y:00 P:27 SP:FB
	// E098  30 04     BMI $E09E                       A:00 X:43 Y:00 P:27 SP:FB
	// E09A  70 02     BVS $E09E                       A:00 X:43 Y:00 P:27 SP:FB
	// E09C  B0 02     BCS $E0A0                       A:00 X:43 Y:00 P:27 SP:FB
	// E0A0  E8        INX                             A:00 X:43 Y:00 P:27 SP:FB
	// E0A1  24 01     BIT $01 = FF                    A:00 X:44 Y:00 P:25 SP:FB
	// E0A3  A9 40     LDA #$40                        A:00 X:44 Y:00 P:E7 SP:FB
	// E0A5  8D 00 04  STA $0400 = 80                  A:40 X:44 Y:00 P:65 SP:FB
	// E0A8  D9 00 04  CMP $0400,Y @ 0400 = 40         A:40 X:44 Y:00 P:65 SP:FB
	// E0AB  30 06     BMI $E0B3                       A:40 X:44 Y:00 P:67 SP:FB
	// E0AD  90 04     BCC $E0B3                       A:40 X:44 Y:00 P:67 SP:FB
	// E0AF  D0 02     BNE $E0B3                       A:40 X:44 Y:00 P:67 SP:FB
	// E0B1  70 02     BVS $E0B5                       A:40 X:44 Y:00 P:67 SP:FB
	// E0B5  E8        INX                             A:40 X:44 Y:00 P:67 SP:FB
	// E0B6  B8        CLV                             A:40 X:45 Y:00 P:65 SP:FB
	// E0B7  CE 00 04  DEC $0400 = 40                  A:40 X:45 Y:00 P:25 SP:FB
	// E0BA  D9 00 04  CMP $0400,Y @ 0400 = 3F         A:40 X:45 Y:00 P:25 SP:FB
	// E0BD  F0 06     BEQ $E0C5                       A:40 X:45 Y:00 P:25 SP:FB
	// E0BF  30 04     BMI $E0C5                       A:40 X:45 Y:00 P:25 SP:FB
	// E0C1  90 02     BCC $E0C5                       A:40 X:45 Y:00 P:25 SP:FB
	// E0C3  50 02     BVC $E0C7                       A:40 X:45 Y:00 P:25 SP:FB
	// E0C7  E8        INX                             A:40 X:45 Y:00 P:25 SP:FB
	// E0C8  EE 00 04  INC $0400 = 3F                  A:40 X:46 Y:00 P:25 SP:FB
	// E0CB  EE 00 04  INC $0400 = 40                  A:40 X:46 Y:00 P:25 SP:FB
	// E0CE  D9 00 04  CMP $0400,Y @ 0400 = 41         A:40 X:46 Y:00 P:25 SP:FB
	// E0D1  F0 02     BEQ $E0D5                       A:40 X:46 Y:00 P:A4 SP:FB
	// E0D3  30 02     BMI $E0D7                       A:40 X:46 Y:00 P:A4 SP:FB
	// E0D7  E8        INX                             A:40 X:46 Y:00 P:A4 SP:FB
	// E0D8  A9 00     LDA #$00                        A:40 X:47 Y:00 P:24 SP:FB
	// E0DA  8D 00 04  STA $0400 = 41                  A:00 X:47 Y:00 P:26 SP:FB
	// E0DD  A9 80     LDA #$80                        A:00 X:47 Y:00 P:26 SP:FB
	// E0DF  D9 00 04  CMP $0400,Y @ 0400 = 00         A:80 X:47 Y:00 P:A4 SP:FB
	// E0E2  F0 04     BEQ $E0E8                       A:80 X:47 Y:00 P:A5 SP:FB
	// E0E4  10 02     BPL $E0E8                       A:80 X:47 Y:00 P:A5 SP:FB
	// E0E6  B0 02     BCS $E0EA                       A:80 X:47 Y:00 P:A5 SP:FB
	// E0EA  E8        INX                             A:80 X:47 Y:00 P:A5 SP:FB
	// E0EB  A0 80     LDY #$80                        A:80 X:48 Y:00 P:25 SP:FB
	// E0ED  8C 00 04  STY $0400 = 00                  A:80 X:48 Y:80 P:A5 SP:FB
	// E0F0  A0 00     LDY #$00                        A:80 X:48 Y:80 P:A5 SP:FB
	// E0F2  D9 00 04  CMP $0400,Y @ 0400 = 80         A:80 X:48 Y:00 P:27 SP:FB
	// E0F5  D0 04     BNE $E0FB                       A:80 X:48 Y:00 P:27 SP:FB
	// E0F7  30 02     BMI $E0FB                       A:80 X:48 Y:00 P:27 SP:FB
	// E0F9  B0 02     BCS $E0FD                       A:80 X:48 Y:00 P:27 SP:FB
	// E0FD  E8        INX                             A:80 X:48 Y:00 P:27 SP:FB
	// E0FE  EE 00 04  INC $0400 = 80                  A:80 X:49 Y:00 P:25 SP:FB
	// E101  D9 00 04  CMP $0400,Y @ 0400 = 81         A:80 X:49 Y:00 P:A5 SP:FB
	// E104  B0 04     BCS $E10A                       A:80 X:49 Y:00 P:A4 SP:FB
	// E106  F0 02     BEQ $E10A                       A:80 X:49 Y:00 P:A4 SP:FB
	// E108  30 02     BMI $E10C                       A:80 X:49 Y:00 P:A4 SP:FB
	// E10C  E8        INX                             A:80 X:49 Y:00 P:A4 SP:FB
	// E10D  CE 00 04  DEC $0400 = 81                  A:80 X:4A Y:00 P:24 SP:FB
	// E110  CE 00 04  DEC $0400 = 80                  A:80 X:4A Y:00 P:A4 SP:FB
	// E113  D9 00 04  CMP $0400,Y @ 0400 = 7F         A:80 X:4A Y:00 P:24 SP:FB
	// E116  90 04     BCC $E11C                       A:80 X:4A Y:00 P:25 SP:FB
	// E118  F0 02     BEQ $E11C                       A:80 X:4A Y:00 P:25 SP:FB
	// E11A  10 02     BPL $E11E                       A:80 X:4A Y:00 P:25 SP:FB
	// E11E  E8        INX                             A:80 X:4A Y:00 P:25 SP:FB
	// E11F  24 01     BIT $01 = FF                    A:80 X:4B Y:00 P:25 SP:FB
	// E121  A9 40     LDA #$40                        A:80 X:4B Y:00 P:E5 SP:FB
	// E123  8D 00 04  STA $0400 = 7F                  A:40 X:4B Y:00 P:65 SP:FB
	// E126  38        SEC                             A:40 X:4B Y:00 P:65 SP:FB
	// E127  F9 00 04  SBC $0400,Y @ 0400 = 40         A:40 X:4B Y:00 P:65 SP:FB
	// E12A  30 0A     BMI $E136                       A:00 X:4B Y:00 P:27 SP:FB
	// E12C  90 08     BCC $E136                       A:00 X:4B Y:00 P:27 SP:FB
	// E12E  D0 06     BNE $E136                       A:00 X:4B Y:00 P:27 SP:FB
	// E130  70 04     BVS $E136                       A:00 X:4B Y:00 P:27 SP:FB
	// E132  C9 00     CMP #$00                        A:00 X:4B Y:00 P:27 SP:FB
	// E134  F0 02     BEQ $E138                       A:00 X:4B Y:00 P:27 SP:FB
	// E138  E8        INX                             A:00 X:4B Y:00 P:27 SP:FB
	// E139  B8        CLV                             A:00 X:4C Y:00 P:25 SP:FB
	// E13A  38        SEC                             A:00 X:4C Y:00 P:25 SP:FB
	// E13B  A9 40     LDA #$40                        A:00 X:4C Y:00 P:25 SP:FB
	// E13D  CE 00 04  DEC $0400 = 40                  A:40 X:4C Y:00 P:25 SP:FB
	// E140  F9 00 04  SBC $0400,Y @ 0400 = 3F         A:40 X:4C Y:00 P:25 SP:FB
	// E143  F0 0A     BEQ $E14F                       A:01 X:4C Y:00 P:25 SP:FB
	// E145  30 08     BMI $E14F                       A:01 X:4C Y:00 P:25 SP:FB
	// E147  90 06     BCC $E14F                       A:01 X:4C Y:00 P:25 SP:FB
	// E149  70 04     BVS $E14F                       A:01 X:4C Y:00 P:25 SP:FB
	// E14B  C9 01     CMP #$01                        A:01 X:4C Y:00 P:25 SP:FB
	// E14D  F0 02     BEQ $E151                       A:01 X:4C Y:00 P:27 SP:FB
	// E151  E8        INX                             A:01 X:4C Y:00 P:27 SP:FB
	// E152  A9 40     LDA #$40                        A:01 X:4D Y:00 P:25 SP:FB
	// E154  38        SEC                             A:40 X:4D Y:00 P:25 SP:FB
	// E155  24 01     BIT $01 = FF                    A:40 X:4D Y:00 P:25 SP:FB
	// E157  EE 00 04  INC $0400 = 3F                  A:40 X:4D Y:00 P:E5 SP:FB
	// E15A  EE 00 04  INC $0400 = 40                  A:40 X:4D Y:00 P:65 SP:FB
	// E15D  F9 00 04  SBC $0400,Y @ 0400 = 41         A:40 X:4D Y:00 P:65 SP:FB
	// E160  B0 0A     BCS $E16C                       A:FF X:4D Y:00 P:A4 SP:FB
	// E162  F0 08     BEQ $E16C                       A:FF X:4D Y:00 P:A4 SP:FB
	// E164  10 06     BPL $E16C                       A:FF X:4D Y:00 P:A4 SP:FB
	// E166  70 04     BVS $E16C                       A:FF X:4D Y:00 P:A4 SP:FB
	// E168  C9 FF     CMP #$FF                        A:FF X:4D Y:00 P:A4 SP:FB
	// E16A  F0 02     BEQ $E16E                       A:FF X:4D Y:00 P:27 SP:FB
	// E16E  E8        INX                             A:FF X:4D Y:00 P:27 SP:FB
	// E16F  18        CLC                             A:FF X:4E Y:00 P:25 SP:FB
	// E170  A9 00     LDA #$00                        A:FF X:4E Y:00 P:24 SP:FB
	// E172  8D 00 04  STA $0400 = 41                  A:00 X:4E Y:00 P:26 SP:FB
	// E175  A9 80     LDA #$80                        A:00 X:4E Y:00 P:26 SP:FB
	// E177  F9 00 04  SBC $0400,Y @ 0400 = 00         A:80 X:4E Y:00 P:A4 SP:FB
	// E17A  90 04     BCC $E180                       A:7F X:4E Y:00 P:65 SP:FB
	// E17C  C9 7F     CMP #$7F                        A:7F X:4E Y:00 P:65 SP:FB
	// E17E  F0 02     BEQ $E182                       A:7F X:4E Y:00 P:67 SP:FB
	// E182  E8        INX                             A:7F X:4E Y:00 P:67 SP:FB
	// E183  38        SEC                             A:7F X:4F Y:00 P:65 SP:FB
	// E184  A9 7F     LDA #$7F                        A:7F X:4F Y:00 P:65 SP:FB
	// E186  8D 00 04  STA $0400 = 00                  A:7F X:4F Y:00 P:65 SP:FB
	// E189  A9 81     LDA #$81                        A:7F X:4F Y:00 P:65 SP:FB
	// E18B  F9 00 04  SBC $0400,Y @ 0400 = 7F         A:81 X:4F Y:00 P:E5 SP:FB
	// E18E  50 06     BVC $E196                       A:02 X:4F Y:00 P:65 SP:FB
	// E190  90 04     BCC $E196                       A:02 X:4F Y:00 P:65 SP:FB
	// E192  C9 02     CMP #$02                        A:02 X:4F Y:00 P:65 SP:FB
	// E194  F0 02     BEQ $E198                       A:02 X:4F Y:00 P:67 SP:FB
	// E198  E8        INX                             A:02 X:4F Y:00 P:67 SP:FB
	// E199  A9 00     LDA #$00                        A:02 X:50 Y:00 P:65 SP:FB
	// E19B  A9 87     LDA #$87                        A:00 X:50 Y:00 P:67 SP:FB
	// E19D  99 00 04  STA $0400,Y @ 0400 = 7F         A:87 X:50 Y:00 P:E5 SP:FB
	// E1A0  AD 00 04  LDA $0400 = 87                  A:87 X:50 Y:00 P:E5 SP:FB
	// E1A3  C9 87     CMP #$87                        A:87 X:50 Y:00 P:E5 SP:FB
	// E1A5  F0 02     BEQ $E1A9                       A:87 X:50 Y:00 P:67 SP:FB
	// E1A9  60        RTS                             A:87 X:50 Y:00 P:67 SP:FB
	// C629  20 B8 DB  JSR $DBB8                       A:87 X:50 Y:00 P:67 SP:FD
	// DBB8  A9 FF     LDA #$FF                        A:87 X:50 Y:00 P:67 SP:FB
	// DBBA  85 01     STA $01 = FF                    A:FF X:50 Y:00 P:E5 SP:FB
	// DBBC  A9 AA     LDA #$AA                        A:FF X:50 Y:00 P:E5 SP:FB
	// DBBE  85 33     STA $33 = A3                    A:AA X:50 Y:00 P:E5 SP:FB
	// DBC0  A9 BB     LDA #$BB                        A:AA X:50 Y:00 P:E5 SP:FB
	// DBC2  85 89     STA $89 = 00                    A:BB X:50 Y:00 P:E5 SP:FB
	// DBC4  A2 00     LDX #$00                        A:BB X:50 Y:00 P:E5 SP:FB
	// DBC6  A9 66     LDA #$66                        A:BB X:00 Y:00 P:67 SP:FB
	// DBC8  24 01     BIT $01 = FF                    A:66 X:00 Y:00 P:65 SP:FB
	// DBCA  38        SEC                             A:66 X:00 Y:00 P:E5 SP:FB
	// DBCB  A0 00     LDY #$00                        A:66 X:00 Y:00 P:E5 SP:FB
	// DBCD  B4 33     LDY $33,X @ 33 = AA             A:66 X:00 Y:00 P:67 SP:FB
	// DBCF  10 12     BPL $DBE3                       A:66 X:00 Y:AA P:E5 SP:FB
	// DBD1  F0 10     BEQ $DBE3                       A:66 X:00 Y:AA P:E5 SP:FB
	// DBD3  50 0E     BVC $DBE3                       A:66 X:00 Y:AA P:E5 SP:FB
	// DBD5  90 0C     BCC $DBE3                       A:66 X:00 Y:AA P:E5 SP:FB
	// DBD7  C9 66     CMP #$66                        A:66 X:00 Y:AA P:E5 SP:FB
	// DBD9  D0 08     BNE $DBE3                       A:66 X:00 Y:AA P:67 SP:FB
	// DBDB  E0 00     CPX #$00                        A:66 X:00 Y:AA P:67 SP:FB
	// DBDD  D0 04     BNE $DBE3                       A:66 X:00 Y:AA P:67 SP:FB
	// DBDF  C0 AA     CPY #$AA                        A:66 X:00 Y:AA P:67 SP:FB
	// DBE1  F0 04     BEQ $DBE7                       A:66 X:00 Y:AA P:67 SP:FB
	// DBE7  A2 8A     LDX #$8A                        A:66 X:00 Y:AA P:67 SP:FB
	// DBE9  A9 66     LDA #$66                        A:66 X:8A Y:AA P:E5 SP:FB
	// DBEB  B8        CLV                             A:66 X:8A Y:AA P:65 SP:FB
	// DBEC  18        CLC                             A:66 X:8A Y:AA P:25 SP:FB
	// DBED  A0 00     LDY #$00                        A:66 X:8A Y:AA P:24 SP:FB
	// DBEF  B4 FF     LDY $FF,X @ 89 = BB             A:66 X:8A Y:00 P:26 SP:FB
	// DBF1  10 12     BPL $DC05                       A:66 X:8A Y:BB P:A4 SP:FB
	// DBF3  F0 10     BEQ $DC05                       A:66 X:8A Y:BB P:A4 SP:FB
	// DBF5  70 0E     BVS $DC05                       A:66 X:8A Y:BB P:A4 SP:FB
	// DBF7  B0 0C     BCS $DC05                       A:66 X:8A Y:BB P:A4 SP:FB
	// DBF9  C0 BB     CPY #$BB                        A:66 X:8A Y:BB P:A4 SP:FB
	// DBFB  D0 08     BNE $DC05                       A:66 X:8A Y:BB P:27 SP:FB
	// DBFD  C9 66     CMP #$66                        A:66 X:8A Y:BB P:27 SP:FB
	// DBFF  D0 04     BNE $DC05                       A:66 X:8A Y:BB P:27 SP:FB
	// DC01  E0 8A     CPX #$8A                        A:66 X:8A Y:BB P:27 SP:FB
	// DC03  F0 04     BEQ $DC09                       A:66 X:8A Y:BB P:27 SP:FB
	// DC09  24 01     BIT $01 = FF                    A:66 X:8A Y:BB P:27 SP:FB
	// DC0B  38        SEC                             A:66 X:8A Y:BB P:E5 SP:FB
	// DC0C  A0 44     LDY #$44                        A:66 X:8A Y:BB P:E5 SP:FB
	// DC0E  A2 00     LDX #$00                        A:66 X:8A Y:44 P:65 SP:FB
	// DC10  94 33     STY $33,X @ 33 = AA             A:66 X:00 Y:44 P:67 SP:FB
	// DC12  A5 33     LDA $33 = 44                    A:66 X:00 Y:44 P:67 SP:FB
	// DC14  90 18     BCC $DC2E                       A:44 X:00 Y:44 P:65 SP:FB
	// DC16  C9 44     CMP #$44                        A:44 X:00 Y:44 P:65 SP:FB
	// DC18  D0 14     BNE $DC2E                       A:44 X:00 Y:44 P:67 SP:FB
	// DC1A  50 12     BVC $DC2E                       A:44 X:00 Y:44 P:67 SP:FB
	// DC1C  18        CLC                             A:44 X:00 Y:44 P:67 SP:FB
	// DC1D  B8        CLV                             A:44 X:00 Y:44 P:66 SP:FB
	// DC1E  A0 99     LDY #$99                        A:44 X:00 Y:44 P:26 SP:FB
	// DC20  A2 80     LDX #$80                        A:44 X:00 Y:99 P:A4 SP:FB
	// DC22  94 85     STY $85,X @ 05 = 00             A:44 X:80 Y:99 P:A4 SP:FB
	// DC24  A5 05     LDA $05 = 99                    A:44 X:80 Y:99 P:A4 SP:FB
	// DC26  B0 06     BCS $DC2E                       A:99 X:80 Y:99 P:A4 SP:FB
	// DC28  C9 99     CMP #$99                        A:99 X:80 Y:99 P:A4 SP:FB
	// DC2A  D0 02     BNE $DC2E                       A:99 X:80 Y:99 P:27 SP:FB
	// DC2C  50 04     BVC $DC32                       A:99 X:80 Y:99 P:27 SP:FB
	// DC32  A0 0B     LDY #$0B                        A:99 X:80 Y:99 P:27 SP:FB
	// DC34  A9 AA     LDA #$AA                        A:99 X:80 Y:0B P:25 SP:FB
	// DC36  A2 78     LDX #$78                        A:AA X:80 Y:0B P:A5 SP:FB
	// DC38  85 78     STA $78 = 00                    A:AA X:78 Y:0B P:25 SP:FB
	// DC3A  20 B6 F7  JSR $F7B6                       A:AA X:78 Y:0B P:25 SP:FB
	// F7B6  18        CLC                             A:AA X:78 Y:0B P:25 SP:F9
	// F7B7  A9 FF     LDA #$FF                        A:AA X:78 Y:0B P:24 SP:F9
	// F7B9  85 01     STA $01 = FF                    A:FF X:78 Y:0B P:A4 SP:F9
	// F7BB  24 01     BIT $01 = FF                    A:FF X:78 Y:0B P:A4 SP:F9
	// F7BD  A9 55     LDA #$55                        A:FF X:78 Y:0B P:E4 SP:F9
	// F7BF  60        RTS                             A:55 X:78 Y:0B P:64 SP:F9
	// DC3D  15 00     ORA $00,X @ 78 = AA             A:55 X:78 Y:0B P:64 SP:FB
	// DC3F  20 C0 F7  JSR $F7C0                       A:FF X:78 Y:0B P:E4 SP:FB
	// F7C0  B0 09     BCS $F7CB                       A:FF X:78 Y:0B P:E4 SP:F9
	// F7C2  10 07     BPL $F7CB                       A:FF X:78 Y:0B P:E4 SP:F9
	// F7C4  C9 FF     CMP #$FF                        A:FF X:78 Y:0B P:E4 SP:F9
	// F7C6  D0 03     BNE $F7CB                       A:FF X:78 Y:0B P:67 SP:F9
	// F7C8  50 01     BVC $F7CB                       A:FF X:78 Y:0B P:67 SP:F9
	// F7CA  60        RTS                             A:FF X:78 Y:0B P:67 SP:F9
	// DC42  C8        INY                             A:FF X:78 Y:0B P:67 SP:FB
	// DC43  A9 00     LDA #$00                        A:FF X:78 Y:0C P:65 SP:FB
	// DC45  85 78     STA $78 = AA                    A:00 X:78 Y:0C P:67 SP:FB
	// DC47  20 CE F7  JSR $F7CE                       A:00 X:78 Y:0C P:67 SP:FB
	// F7CE  38        SEC                             A:00 X:78 Y:0C P:67 SP:F9
	// F7CF  B8        CLV                             A:00 X:78 Y:0C P:67 SP:F9
	// F7D0  A9 00     LDA #$00                        A:00 X:78 Y:0C P:27 SP:F9
	// F7D2  60        RTS                             A:00 X:78 Y:0C P:27 SP:F9
	// DC4A  15 00     ORA $00,X @ 78 = 00             A:00 X:78 Y:0C P:27 SP:FB
	// DC4C  20 D3 F7  JSR $F7D3                       A:00 X:78 Y:0C P:27 SP:FB
	// F7D3  D0 07     BNE $F7DC                       A:00 X:78 Y:0C P:27 SP:F9
	// F7D5  70 05     BVS $F7DC                       A:00 X:78 Y:0C P:27 SP:F9
	// F7D7  90 03     BCC $F7DC                       A:00 X:78 Y:0C P:27 SP:F9
	// F7D9  30 01     BMI $F7DC                       A:00 X:78 Y:0C P:27 SP:F9
	// F7DB  60        RTS                             A:00 X:78 Y:0C P:27 SP:F9
	// DC4F  C8        INY                             A:00 X:78 Y:0C P:27 SP:FB
	// DC50  A9 AA     LDA #$AA                        A:00 X:78 Y:0D P:25 SP:FB
	// DC52  85 78     STA $78 = 00                    A:AA X:78 Y:0D P:A5 SP:FB
	// DC54  20 DF F7  JSR $F7DF                       A:AA X:78 Y:0D P:A5 SP:FB
	// F7DF  18        CLC                             A:AA X:78 Y:0D P:A5 SP:F9
	// F7E0  24 01     BIT $01 = FF                    A:AA X:78 Y:0D P:A4 SP:F9
	// F7E2  A9 55     LDA #$55                        A:AA X:78 Y:0D P:E4 SP:F9
	// F7E4  60        RTS                             A:55 X:78 Y:0D P:64 SP:F9
	// DC57  35 00     AND $00,X @ 78 = AA             A:55 X:78 Y:0D P:64 SP:FB
	// DC59  20 E5 F7  JSR $F7E5                       A:00 X:78 Y:0D P:66 SP:FB
	// F7E5  D0 07     BNE $F7EE                       A:00 X:78 Y:0D P:66 SP:F9
	// F7E7  50 05     BVC $F7EE                       A:00 X:78 Y:0D P:66 SP:F9
	// F7E9  B0 03     BCS $F7EE                       A:00 X:78 Y:0D P:66 SP:F9
	// F7EB  30 01     BMI $F7EE                       A:00 X:78 Y:0D P:66 SP:F9
	// F7ED  60        RTS                             A:00 X:78 Y:0D P:66 SP:F9
	// DC5C  C8        INY                             A:00 X:78 Y:0D P:66 SP:FB
	// DC5D  A9 EF     LDA #$EF                        A:00 X:78 Y:0E P:64 SP:FB
	// DC5F  85 78     STA $78 = AA                    A:EF X:78 Y:0E P:E4 SP:FB
	// DC61  20 F1 F7  JSR $F7F1                       A:EF X:78 Y:0E P:E4 SP:FB
	// F7F1  38        SEC                             A:EF X:78 Y:0E P:E4 SP:F9
	// F7F2  B8        CLV                             A:EF X:78 Y:0E P:E5 SP:F9
	// F7F3  A9 F8     LDA #$F8                        A:EF X:78 Y:0E P:A5 SP:F9
	// F7F5  60        RTS                             A:F8 X:78 Y:0E P:A5 SP:F9
	// DC64  35 00     AND $00,X @ 78 = EF             A:F8 X:78 Y:0E P:A5 SP:FB
	// DC66  20 F6 F7  JSR $F7F6                       A:E8 X:78 Y:0E P:A5 SP:FB
	// F7F6  90 09     BCC $F801                       A:E8 X:78 Y:0E P:A5 SP:F9
	// F7F8  10 07     BPL $F801                       A:E8 X:78 Y:0E P:A5 SP:F9
	// F7FA  C9 E8     CMP #$E8                        A:E8 X:78 Y:0E P:A5 SP:F9
	// F7FC  D0 03     BNE $F801                       A:E8 X:78 Y:0E P:27 SP:F9
	// F7FE  70 01     BVS $F801                       A:E8 X:78 Y:0E P:27 SP:F9
	// F800  60        RTS                             A:E8 X:78 Y:0E P:27 SP:F9
	// DC69  C8        INY                             A:E8 X:78 Y:0E P:27 SP:FB
	// DC6A  A9 AA     LDA #$AA                        A:E8 X:78 Y:0F P:25 SP:FB
	// DC6C  85 78     STA $78 = EF                    A:AA X:78 Y:0F P:A5 SP:FB
	// DC6E  20 04 F8  JSR $F804                       A:AA X:78 Y:0F P:A5 SP:FB
	// F804  18        CLC                             A:AA X:78 Y:0F P:A5 SP:F9
	// F805  24 01     BIT $01 = FF                    A:AA X:78 Y:0F P:A4 SP:F9
	// F807  A9 5F     LDA #$5F                        A:AA X:78 Y:0F P:E4 SP:F9
	// F809  60        RTS                             A:5F X:78 Y:0F P:64 SP:F9
	// DC71  55 00     EOR $00,X @ 78 = AA             A:5F X:78 Y:0F P:64 SP:FB
	// DC73  20 0A F8  JSR $F80A                       A:F5 X:78 Y:0F P:E4 SP:FB
	// F80A  B0 09     BCS $F815                       A:F5 X:78 Y:0F P:E4 SP:F9
	// F80C  10 07     BPL $F815                       A:F5 X:78 Y:0F P:E4 SP:F9
	// F80E  C9 F5     CMP #$F5                        A:F5 X:78 Y:0F P:E4 SP:F9
	// F810  D0 03     BNE $F815                       A:F5 X:78 Y:0F P:67 SP:F9
	// F812  50 01     BVC $F815                       A:F5 X:78 Y:0F P:67 SP:F9
	// F814  60        RTS                             A:F5 X:78 Y:0F P:67 SP:F9
	// DC76  C8        INY                             A:F5 X:78 Y:0F P:67 SP:FB
	// DC77  A9 70     LDA #$70                        A:F5 X:78 Y:10 P:65 SP:FB
	// DC79  85 78     STA $78 = AA                    A:70 X:78 Y:10 P:65 SP:FB
	// DC7B  20 18 F8  JSR $F818                       A:70 X:78 Y:10 P:65 SP:FB
	// F818  38        SEC                             A:70 X:78 Y:10 P:65 SP:F9
	// F819  B8        CLV                             A:70 X:78 Y:10 P:65 SP:F9
	// F81A  A9 70     LDA #$70                        A:70 X:78 Y:10 P:25 SP:F9
	// F81C  60        RTS                             A:70 X:78 Y:10 P:25 SP:F9
	// DC7E  55 00     EOR $00,X @ 78 = 70             A:70 X:78 Y:10 P:25 SP:FB
	// DC80  20 1D F8  JSR $F81D                       A:00 X:78 Y:10 P:27 SP:FB
	// F81D  D0 07     BNE $F826                       A:00 X:78 Y:10 P:27 SP:F9
	// F81F  70 05     BVS $F826                       A:00 X:78 Y:10 P:27 SP:F9
	// F821  90 03     BCC $F826                       A:00 X:78 Y:10 P:27 SP:F9
	// F823  30 01     BMI $F826                       A:00 X:78 Y:10 P:27 SP:F9
	// F825  60        RTS                             A:00 X:78 Y:10 P:27 SP:F9
	// DC83  C8        INY                             A:00 X:78 Y:10 P:27 SP:FB
	// DC84  A9 69     LDA #$69                        A:00 X:78 Y:11 P:25 SP:FB
	// DC86  85 78     STA $78 = 70                    A:69 X:78 Y:11 P:25 SP:FB
	// DC88  20 29 F8  JSR $F829                       A:69 X:78 Y:11 P:25 SP:FB
	// F829  18        CLC                             A:69 X:78 Y:11 P:25 SP:F9
	// F82A  24 01     BIT $01 = FF                    A:69 X:78 Y:11 P:24 SP:F9
	// F82C  A9 00     LDA #$00                        A:69 X:78 Y:11 P:E4 SP:F9
	// F82E  60        RTS                             A:00 X:78 Y:11 P:66 SP:F9
	// DC8B  75 00     ADC $00,X @ 78 = 69             A:00 X:78 Y:11 P:66 SP:FB
	// DC8D  20 2F F8  JSR $F82F                       A:69 X:78 Y:11 P:24 SP:FB
	// F82F  30 09     BMI $F83A                       A:69 X:78 Y:11 P:24 SP:F9
	// F831  B0 07     BCS $F83A                       A:69 X:78 Y:11 P:24 SP:F9
	// F833  C9 69     CMP #$69                        A:69 X:78 Y:11 P:24 SP:F9
	// F835  D0 03     BNE $F83A                       A:69 X:78 Y:11 P:27 SP:F9
	// F837  70 01     BVS $F83A                       A:69 X:78 Y:11 P:27 SP:F9
	// F839  60        RTS                             A:69 X:78 Y:11 P:27 SP:F9
	// DC90  C8        INY                             A:69 X:78 Y:11 P:27 SP:FB
	// DC91  20 3D F8  JSR $F83D                       A:69 X:78 Y:12 P:25 SP:FB
	// F83D  38        SEC                             A:69 X:78 Y:12 P:25 SP:F9
	// F83E  24 01     BIT $01 = FF                    A:69 X:78 Y:12 P:25 SP:F9
	// F840  A9 00     LDA #$00                        A:69 X:78 Y:12 P:E5 SP:F9
	// F842  60        RTS                             A:00 X:78 Y:12 P:67 SP:F9
	// DC94  75 00     ADC $00,X @ 78 = 69             A:00 X:78 Y:12 P:67 SP:FB
	// DC96  20 43 F8  JSR $F843                       A:6A X:78 Y:12 P:24 SP:FB
	// F843  30 09     BMI $F84E                       A:6A X:78 Y:12 P:24 SP:F9
	// F845  B0 07     BCS $F84E                       A:6A X:78 Y:12 P:24 SP:F9
	// F847  C9 6A     CMP #$6A                        A:6A X:78 Y:12 P:24 SP:F9
	// F849  D0 03     BNE $F84E                       A:6A X:78 Y:12 P:27 SP:F9
	// F84B  70 01     BVS $F84E                       A:6A X:78 Y:12 P:27 SP:F9
	// F84D  60        RTS                             A:6A X:78 Y:12 P:27 SP:F9
	// DC99  C8        INY                             A:6A X:78 Y:12 P:27 SP:FB
	// DC9A  A9 7F     LDA #$7F                        A:6A X:78 Y:13 P:25 SP:FB
	// DC9C  85 78     STA $78 = 69                    A:7F X:78 Y:13 P:25 SP:FB
	// DC9E  20 51 F8  JSR $F851                       A:7F X:78 Y:13 P:25 SP:FB
	// F851  38        SEC                             A:7F X:78 Y:13 P:25 SP:F9
	// F852  B8        CLV                             A:7F X:78 Y:13 P:25 SP:F9
	// F853  A9 7F     LDA #$7F                        A:7F X:78 Y:13 P:25 SP:F9
	// F855  60        RTS                             A:7F X:78 Y:13 P:25 SP:F9
	// DCA1  75 00     ADC $00,X @ 78 = 7F             A:7F X:78 Y:13 P:25 SP:FB
	// DCA3  20 56 F8  JSR $F856                       A:FF X:78 Y:13 P:E4 SP:FB
	// F856  10 09     BPL $F861                       A:FF X:78 Y:13 P:E4 SP:F9
	// F858  B0 07     BCS $F861                       A:FF X:78 Y:13 P:E4 SP:F9
	// F85A  C9 FF     CMP #$FF                        A:FF X:78 Y:13 P:E4 SP:F9
	// F85C  D0 03     BNE $F861                       A:FF X:78 Y:13 P:67 SP:F9
	// F85E  50 01     BVC $F861                       A:FF X:78 Y:13 P:67 SP:F9
	// F860  60        RTS                             A:FF X:78 Y:13 P:67 SP:F9
	// DCA6  C8        INY                             A:FF X:78 Y:13 P:67 SP:FB
	// DCA7  A9 80     LDA #$80                        A:FF X:78 Y:14 P:65 SP:FB
	// DCA9  85 78     STA $78 = 7F                    A:80 X:78 Y:14 P:E5 SP:FB
	// DCAB  20 64 F8  JSR $F864                       A:80 X:78 Y:14 P:E5 SP:FB
	// F864  18        CLC                             A:80 X:78 Y:14 P:E5 SP:F9
	// F865  24 01     BIT $01 = FF                    A:80 X:78 Y:14 P:E4 SP:F9
	// F867  A9 7F     LDA #$7F                        A:80 X:78 Y:14 P:E4 SP:F9
	// F869  60        RTS                             A:7F X:78 Y:14 P:64 SP:F9
	// DCAE  75 00     ADC $00,X @ 78 = 80             A:7F X:78 Y:14 P:64 SP:FB
	// DCB0  20 6A F8  JSR $F86A                       A:FF X:78 Y:14 P:A4 SP:FB
	// F86A  10 09     BPL $F875                       A:FF X:78 Y:14 P:A4 SP:F9
	// F86C  B0 07     BCS $F875                       A:FF X:78 Y:14 P:A4 SP:F9
	// F86E  C9 FF     CMP #$FF                        A:FF X:78 Y:14 P:A4 SP:F9
	// F870  D0 03     BNE $F875                       A:FF X:78 Y:14 P:27 SP:F9
	// F872  70 01     BVS $F875                       A:FF X:78 Y:14 P:27 SP:F9
	// F874  60        RTS                             A:FF X:78 Y:14 P:27 SP:F9
	// DCB3  C8        INY                             A:FF X:78 Y:14 P:27 SP:FB
	// DCB4  20 78 F8  JSR $F878                       A:FF X:78 Y:15 P:25 SP:FB
	// F878  38        SEC                             A:FF X:78 Y:15 P:25 SP:F9
	// F879  B8        CLV                             A:FF X:78 Y:15 P:25 SP:F9
	// F87A  A9 7F     LDA #$7F                        A:FF X:78 Y:15 P:25 SP:F9
	// F87C  60        RTS                             A:7F X:78 Y:15 P:25 SP:F9
	// DCB7  75 00     ADC $00,X @ 78 = 80             A:7F X:78 Y:15 P:25 SP:FB
	// DCB9  20 7D F8  JSR $F87D                       A:00 X:78 Y:15 P:27 SP:FB
	// F87D  D0 07     BNE $F886                       A:00 X:78 Y:15 P:27 SP:F9
	// F87F  30 05     BMI $F886                       A:00 X:78 Y:15 P:27 SP:F9
	// F881  70 03     BVS $F886                       A:00 X:78 Y:15 P:27 SP:F9
	// F883  90 01     BCC $F886                       A:00 X:78 Y:15 P:27 SP:F9
	// F885  60        RTS                             A:00 X:78 Y:15 P:27 SP:F9
	// DCBC  C8        INY                             A:00 X:78 Y:15 P:27 SP:FB
	// DCBD  A9 40     LDA #$40                        A:00 X:78 Y:16 P:25 SP:FB
	// DCBF  85 78     STA $78 = 80                    A:40 X:78 Y:16 P:25 SP:FB
	// DCC1  20 89 F8  JSR $F889                       A:40 X:78 Y:16 P:25 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:78 Y:16 P:25 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:78 Y:16 P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:78 Y:16 P:65 SP:F9
	// DCC4  D5 00     CMP $00,X @ 78 = 40             A:40 X:78 Y:16 P:65 SP:FB
	// DCC6  20 8E F8  JSR $F88E                       A:40 X:78 Y:16 P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:78 Y:16 P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:78 Y:16 P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:78 Y:16 P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:78 Y:16 P:67 SP:F9
	// F896  60        RTS                             A:40 X:78 Y:16 P:67 SP:F9
	// DCC9  C8        INY                             A:40 X:78 Y:16 P:67 SP:FB
	// DCCA  48        PHA                             A:40 X:78 Y:17 P:65 SP:FB
	// DCCB  A9 3F     LDA #$3F                        A:40 X:78 Y:17 P:65 SP:FA
	// DCCD  85 78     STA $78 = 40                    A:3F X:78 Y:17 P:65 SP:FA
	// DCCF  68        PLA                             A:3F X:78 Y:17 P:65 SP:FA
	// DCD0  20 9A F8  JSR $F89A                       A:40 X:78 Y:17 P:65 SP:FB
	// F89A  B8        CLV                             A:40 X:78 Y:17 P:65 SP:F9
	// F89B  60        RTS                             A:40 X:78 Y:17 P:25 SP:F9
	// DCD3  D5 00     CMP $00,X @ 78 = 3F             A:40 X:78 Y:17 P:25 SP:FB
	// DCD5  20 9C F8  JSR $F89C                       A:40 X:78 Y:17 P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:40 X:78 Y:17 P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:40 X:78 Y:17 P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:40 X:78 Y:17 P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:40 X:78 Y:17 P:25 SP:F9
	// F8A4  60        RTS                             A:40 X:78 Y:17 P:25 SP:F9
	// DCD8  C8        INY                             A:40 X:78 Y:17 P:25 SP:FB
	// DCD9  48        PHA                             A:40 X:78 Y:18 P:25 SP:FB
	// DCDA  A9 41     LDA #$41                        A:40 X:78 Y:18 P:25 SP:FA
	// DCDC  85 78     STA $78 = 3F                    A:41 X:78 Y:18 P:25 SP:FA
	// DCDE  68        PLA                             A:41 X:78 Y:18 P:25 SP:FA
	// DCDF  D5 00     CMP $00,X @ 78 = 41             A:40 X:78 Y:18 P:25 SP:FB
	// DCE1  20 A8 F8  JSR $F8A8                       A:40 X:78 Y:18 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:40 X:78 Y:18 P:A4 SP:F9
	// F8AE  60        RTS                             A:40 X:78 Y:18 P:A4 SP:F9
	// DCE4  C8        INY                             A:40 X:78 Y:18 P:A4 SP:FB
	// DCE5  48        PHA                             A:40 X:78 Y:19 P:24 SP:FB
	// DCE6  A9 00     LDA #$00                        A:40 X:78 Y:19 P:24 SP:FA
	// DCE8  85 78     STA $78 = 41                    A:00 X:78 Y:19 P:26 SP:FA
	// DCEA  68        PLA                             A:00 X:78 Y:19 P:26 SP:FA
	// DCEB  20 B2 F8  JSR $F8B2                       A:40 X:78 Y:19 P:24 SP:FB
	// F8B2  A9 80     LDA #$80                        A:40 X:78 Y:19 P:24 SP:F9
	// F8B4  60        RTS                             A:80 X:78 Y:19 P:A4 SP:F9
	// DCEE  D5 00     CMP $00,X @ 78 = 00             A:80 X:78 Y:19 P:A4 SP:FB
	// DCF0  20 B5 F8  JSR $F8B5                       A:80 X:78 Y:19 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:78 Y:19 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:78 Y:19 P:A5 SP:F9
	// DCF3  C8        INY                             A:80 X:78 Y:19 P:A5 SP:FB
	// DCF4  48        PHA                             A:80 X:78 Y:1A P:25 SP:FB
	// DCF5  A9 80     LDA #$80                        A:80 X:78 Y:1A P:25 SP:FA
	// DCF7  85 78     STA $78 = 00                    A:80 X:78 Y:1A P:A5 SP:FA
	// DCF9  68        PLA                             A:80 X:78 Y:1A P:A5 SP:FA
	// DCFA  D5 00     CMP $00,X @ 78 = 80             A:80 X:78 Y:1A P:A5 SP:FB
	// DCFC  20 BF F8  JSR $F8BF                       A:80 X:78 Y:1A P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:78 Y:1A P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:78 Y:1A P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:78 Y:1A P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:78 Y:1A P:27 SP:F9
	// DCFF  C8        INY                             A:80 X:78 Y:1A P:27 SP:FB
	// DD00  48        PHA                             A:80 X:78 Y:1B P:25 SP:FB
	// DD01  A9 81     LDA #$81                        A:80 X:78 Y:1B P:25 SP:FA
	// DD03  85 78     STA $78 = 80                    A:81 X:78 Y:1B P:A5 SP:FA
	// DD05  68        PLA                             A:81 X:78 Y:1B P:A5 SP:FA
	// DD06  D5 00     CMP $00,X @ 78 = 81             A:80 X:78 Y:1B P:A5 SP:FB
	// DD08  20 C9 F8  JSR $F8C9                       A:80 X:78 Y:1B P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:80 X:78 Y:1B P:A4 SP:F9
	// F8CF  60        RTS                             A:80 X:78 Y:1B P:A4 SP:F9
	// DD0B  C8        INY                             A:80 X:78 Y:1B P:A4 SP:FB
	// DD0C  48        PHA                             A:80 X:78 Y:1C P:24 SP:FB
	// DD0D  A9 7F     LDA #$7F                        A:80 X:78 Y:1C P:24 SP:FA
	// DD0F  85 78     STA $78 = 81                    A:7F X:78 Y:1C P:24 SP:FA
	// DD11  68        PLA                             A:7F X:78 Y:1C P:24 SP:FA
	// DD12  D5 00     CMP $00,X @ 78 = 7F             A:80 X:78 Y:1C P:A4 SP:FB
	// DD14  20 D3 F8  JSR $F8D3                       A:80 X:78 Y:1C P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:80 X:78 Y:1C P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:80 X:78 Y:1C P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:80 X:78 Y:1C P:25 SP:F9
	// F8D9  60        RTS                             A:80 X:78 Y:1C P:25 SP:F9
	// DD17  C8        INY                             A:80 X:78 Y:1C P:25 SP:FB
	// DD18  A9 40     LDA #$40                        A:80 X:78 Y:1D P:25 SP:FB
	// DD1A  85 78     STA $78 = 7F                    A:40 X:78 Y:1D P:25 SP:FB
	// DD1C  20 31 F9  JSR $F931                       A:40 X:78 Y:1D P:25 SP:FB
	// F931  24 01     BIT $01 = FF                    A:40 X:78 Y:1D P:25 SP:F9
	// F933  A9 40     LDA #$40                        A:40 X:78 Y:1D P:E5 SP:F9
	// F935  38        SEC                             A:40 X:78 Y:1D P:65 SP:F9
	// F936  60        RTS                             A:40 X:78 Y:1D P:65 SP:F9
	// DD1F  F5 00     SBC $00,X @ 78 = 40             A:40 X:78 Y:1D P:65 SP:FB
	// DD21  20 37 F9  JSR $F937                       A:00 X:78 Y:1D P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:78 Y:1D P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:78 Y:1D P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:78 Y:1D P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:78 Y:1D P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:78 Y:1D P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:78 Y:1D P:27 SP:F9
	// F943  60        RTS                             A:00 X:78 Y:1D P:27 SP:F9
	// DD24  C8        INY                             A:00 X:78 Y:1D P:27 SP:FB
	// DD25  A9 3F     LDA #$3F                        A:00 X:78 Y:1E P:25 SP:FB
	// DD27  85 78     STA $78 = 40                    A:3F X:78 Y:1E P:25 SP:FB
	// DD29  20 47 F9  JSR $F947                       A:3F X:78 Y:1E P:25 SP:FB
	// F947  B8        CLV                             A:3F X:78 Y:1E P:25 SP:F9
	// F948  38        SEC                             A:3F X:78 Y:1E P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:3F X:78 Y:1E P:25 SP:F9
	// F94B  60        RTS                             A:40 X:78 Y:1E P:25 SP:F9
	// DD2C  F5 00     SBC $00,X @ 78 = 3F             A:40 X:78 Y:1E P:25 SP:FB
	// DD2E  20 4C F9  JSR $F94C                       A:01 X:78 Y:1E P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:78 Y:1E P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:78 Y:1E P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:78 Y:1E P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:78 Y:1E P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:78 Y:1E P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:78 Y:1E P:27 SP:F9
	// F958  60        RTS                             A:01 X:78 Y:1E P:27 SP:F9
	// DD31  C8        INY                             A:01 X:78 Y:1E P:27 SP:FB
	// DD32  A9 41     LDA #$41                        A:01 X:78 Y:1F P:25 SP:FB
	// DD34  85 78     STA $78 = 3F                    A:41 X:78 Y:1F P:25 SP:FB
	// DD36  20 5C F9  JSR $F95C                       A:41 X:78 Y:1F P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:41 X:78 Y:1F P:25 SP:F9
	// F95E  38        SEC                             A:40 X:78 Y:1F P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:78 Y:1F P:25 SP:F9
	// F961  60        RTS                             A:40 X:78 Y:1F P:E5 SP:F9
	// DD39  F5 00     SBC $00,X @ 78 = 41             A:40 X:78 Y:1F P:E5 SP:FB
	// DD3B  20 62 F9  JSR $F962                       A:FF X:78 Y:1F P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:78 Y:1F P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:78 Y:1F P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:78 Y:1F P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:78 Y:1F P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:78 Y:1F P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:78 Y:1F P:27 SP:F9
	// F96E  60        RTS                             A:FF X:78 Y:1F P:27 SP:F9
	// DD3E  C8        INY                             A:FF X:78 Y:1F P:27 SP:FB
	// DD3F  A9 00     LDA #$00                        A:FF X:78 Y:20 P:25 SP:FB
	// DD41  85 78     STA $78 = 41                    A:00 X:78 Y:20 P:27 SP:FB
	// DD43  20 72 F9  JSR $F972                       A:00 X:78 Y:20 P:27 SP:FB
	// F972  18        CLC                             A:00 X:78 Y:20 P:27 SP:F9
	// F973  A9 80     LDA #$80                        A:00 X:78 Y:20 P:26 SP:F9
	// F975  60        RTS                             A:80 X:78 Y:20 P:A4 SP:F9
	// DD46  F5 00     SBC $00,X @ 78 = 00             A:80 X:78 Y:20 P:A4 SP:FB
	// DD48  20 76 F9  JSR $F976                       A:7F X:78 Y:20 P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:78 Y:20 P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:78 Y:20 P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:78 Y:20 P:67 SP:F9
	// F97C  60        RTS                             A:7F X:78 Y:20 P:67 SP:F9
	// DD4B  C8        INY                             A:7F X:78 Y:20 P:67 SP:FB
	// DD4C  A9 7F     LDA #$7F                        A:7F X:78 Y:21 P:65 SP:FB
	// DD4E  85 78     STA $78 = 00                    A:7F X:78 Y:21 P:65 SP:FB
	// DD50  20 80 F9  JSR $F980                       A:7F X:78 Y:21 P:65 SP:FB
	// F980  38        SEC                             A:7F X:78 Y:21 P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:78 Y:21 P:65 SP:F9
	// F983  60        RTS                             A:81 X:78 Y:21 P:E5 SP:F9
	// DD53  F5 00     SBC $00,X @ 78 = 7F             A:81 X:78 Y:21 P:E5 SP:FB
	// DD55  20 84 F9  JSR $F984                       A:02 X:78 Y:21 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:78 Y:21 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:78 Y:21 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:78 Y:21 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:78 Y:21 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:78 Y:21 P:67 SP:F9
	// DD58  A9 AA     LDA #$AA                        A:02 X:78 Y:21 P:67 SP:FB
	// DD5A  85 33     STA $33 = 44                    A:AA X:78 Y:21 P:E5 SP:FB
	// DD5C  A9 BB     LDA #$BB                        A:AA X:78 Y:21 P:E5 SP:FB
	// DD5E  85 89     STA $89 = BB                    A:BB X:78 Y:21 P:E5 SP:FB
	// DD60  A2 00     LDX #$00                        A:BB X:78 Y:21 P:E5 SP:FB
	// DD62  A0 66     LDY #$66                        A:BB X:00 Y:21 P:67 SP:FB
	// DD64  24 01     BIT $01 = FF                    A:BB X:00 Y:66 P:65 SP:FB
	// DD66  38        SEC                             A:BB X:00 Y:66 P:E5 SP:FB
	// DD67  A9 00     LDA #$00                        A:BB X:00 Y:66 P:E5 SP:FB
	// DD69  B5 33     LDA $33,X @ 33 = AA             A:00 X:00 Y:66 P:67 SP:FB
	// DD6B  10 12     BPL $DD7F                       A:AA X:00 Y:66 P:E5 SP:FB
	// DD6D  F0 10     BEQ $DD7F                       A:AA X:00 Y:66 P:E5 SP:FB
	// DD6F  50 0E     BVC $DD7F                       A:AA X:00 Y:66 P:E5 SP:FB
	// DD71  90 0C     BCC $DD7F                       A:AA X:00 Y:66 P:E5 SP:FB
	// DD73  C0 66     CPY #$66                        A:AA X:00 Y:66 P:E5 SP:FB
	// DD75  D0 08     BNE $DD7F                       A:AA X:00 Y:66 P:67 SP:FB
	// DD77  E0 00     CPX #$00                        A:AA X:00 Y:66 P:67 SP:FB
	// DD79  D0 04     BNE $DD7F                       A:AA X:00 Y:66 P:67 SP:FB
	// DD7B  C9 AA     CMP #$AA                        A:AA X:00 Y:66 P:67 SP:FB
	// DD7D  F0 04     BEQ $DD83                       A:AA X:00 Y:66 P:67 SP:FB
	// DD83  A2 8A     LDX #$8A                        A:AA X:00 Y:66 P:67 SP:FB
	// DD85  A0 66     LDY #$66                        A:AA X:8A Y:66 P:E5 SP:FB
	// DD87  B8        CLV                             A:AA X:8A Y:66 P:65 SP:FB
	// DD88  18        CLC                             A:AA X:8A Y:66 P:25 SP:FB
	// DD89  A9 00     LDA #$00                        A:AA X:8A Y:66 P:24 SP:FB
	// DD8B  B5 FF     LDA $FF,X @ 89 = BB             A:00 X:8A Y:66 P:26 SP:FB
	// DD8D  10 12     BPL $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB
	// DD8F  F0 10     BEQ $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB
	// DD91  70 0E     BVS $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB
	// DD93  B0 0C     BCS $DDA1                       A:BB X:8A Y:66 P:A4 SP:FB
	// DD95  C9 BB     CMP #$BB                        A:BB X:8A Y:66 P:A4 SP:FB
	// DD97  D0 08     BNE $DDA1                       A:BB X:8A Y:66 P:27 SP:FB
	// DD99  C0 66     CPY #$66                        A:BB X:8A Y:66 P:27 SP:FB
	// DD9B  D0 04     BNE $DDA1                       A:BB X:8A Y:66 P:27 SP:FB
	// DD9D  E0 8A     CPX #$8A                        A:BB X:8A Y:66 P:27 SP:FB
	// DD9F  F0 04     BEQ $DDA5                       A:BB X:8A Y:66 P:27 SP:FB
	// DDA5  24 01     BIT $01 = FF                    A:BB X:8A Y:66 P:27 SP:FB
	// DDA7  38        SEC                             A:BB X:8A Y:66 P:E5 SP:FB
	// DDA8  A9 44     LDA #$44                        A:BB X:8A Y:66 P:E5 SP:FB
	// DDAA  A2 00     LDX #$00                        A:44 X:8A Y:66 P:65 SP:FB
	// DDAC  95 33     STA $33,X @ 33 = AA             A:44 X:00 Y:66 P:67 SP:FB
	// DDAE  A5 33     LDA $33 = 44                    A:44 X:00 Y:66 P:67 SP:FB
	// DDB0  90 18     BCC $DDCA                       A:44 X:00 Y:66 P:65 SP:FB
	// DDB2  C9 44     CMP #$44                        A:44 X:00 Y:66 P:65 SP:FB
	// DDB4  D0 14     BNE $DDCA                       A:44 X:00 Y:66 P:67 SP:FB
	// DDB6  50 12     BVC $DDCA                       A:44 X:00 Y:66 P:67 SP:FB
	// DDB8  18        CLC                             A:44 X:00 Y:66 P:67 SP:FB
	// DDB9  B8        CLV                             A:44 X:00 Y:66 P:66 SP:FB
	// DDBA  A9 99     LDA #$99                        A:44 X:00 Y:66 P:26 SP:FB
	// DDBC  A2 80     LDX #$80                        A:99 X:00 Y:66 P:A4 SP:FB
	// DDBE  95 85     STA $85,X @ 05 = 99             A:99 X:80 Y:66 P:A4 SP:FB
	// DDC0  A5 05     LDA $05 = 99                    A:99 X:80 Y:66 P:A4 SP:FB
	// DDC2  B0 06     BCS $DDCA                       A:99 X:80 Y:66 P:A4 SP:FB
	// DDC4  C9 99     CMP #$99                        A:99 X:80 Y:66 P:A4 SP:FB
	// DDC6  D0 02     BNE $DDCA                       A:99 X:80 Y:66 P:27 SP:FB
	// DDC8  50 04     BVC $DDCE                       A:99 X:80 Y:66 P:27 SP:FB
	// DDCE  A0 25     LDY #$25                        A:99 X:80 Y:66 P:27 SP:FB
	// DDD0  A2 78     LDX #$78                        A:99 X:80 Y:25 P:25 SP:FB
	// DDD2  20 90 F9  JSR $F990                       A:99 X:78 Y:25 P:25 SP:FB
	// F990  A2 55     LDX #$55                        A:99 X:78 Y:25 P:25 SP:F9
	// F992  A9 FF     LDA #$FF                        A:99 X:55 Y:25 P:25 SP:F9
	// F994  85 01     STA $01 = FF                    A:FF X:55 Y:25 P:A5 SP:F9
	// F996  EA        NOP                             A:FF X:55 Y:25 P:A5 SP:F9
	// F997  24 01     BIT $01 = FF                    A:FF X:55 Y:25 P:A5 SP:F9
	// F999  38        SEC                             A:FF X:55 Y:25 P:E5 SP:F9
	// F99A  A9 01     LDA #$01                        A:FF X:55 Y:25 P:E5 SP:F9
	// F99C  60        RTS                             A:01 X:55 Y:25 P:65 SP:F9
	// DDD5  95 00     STA $00,X @ 55 = 00             A:01 X:55 Y:25 P:65 SP:FB
	// DDD7  56 00     LSR $00,X @ 55 = 01             A:01 X:55 Y:25 P:65 SP:FB
	// DDD9  B5 00     LDA $00,X @ 55 = 00             A:01 X:55 Y:25 P:67 SP:FB
	// DDDB  20 9D F9  JSR $F99D                       A:00 X:55 Y:25 P:67 SP:FB
	// F99D  90 1B     BCC $F9BA                       A:00 X:55 Y:25 P:67 SP:F9
	// F99F  D0 19     BNE $F9BA                       A:00 X:55 Y:25 P:67 SP:F9
	// F9A1  30 17     BMI $F9BA                       A:00 X:55 Y:25 P:67 SP:F9
	// F9A3  50 15     BVC $F9BA                       A:00 X:55 Y:25 P:67 SP:F9
	// F9A5  C9 00     CMP #$00                        A:00 X:55 Y:25 P:67 SP:F9
	// F9A7  D0 11     BNE $F9BA                       A:00 X:55 Y:25 P:67 SP:F9
	// F9A9  B8        CLV                             A:00 X:55 Y:25 P:67 SP:F9
	// F9AA  A9 AA     LDA #$AA                        A:00 X:55 Y:25 P:27 SP:F9
	// F9AC  60        RTS                             A:AA X:55 Y:25 P:A5 SP:F9
	// DDDE  C8        INY                             A:AA X:55 Y:25 P:A5 SP:FB
	// DDDF  95 00     STA $00,X @ 55 = 00             A:AA X:55 Y:26 P:25 SP:FB
	// DDE1  56 00     LSR $00,X @ 55 = AA             A:AA X:55 Y:26 P:25 SP:FB
	// DDE3  B5 00     LDA $00,X @ 55 = 55             A:AA X:55 Y:26 P:24 SP:FB
	// DDE5  20 AD F9  JSR $F9AD                       A:55 X:55 Y:26 P:24 SP:FB
	// F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:26 P:24 SP:F9
	// F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:26 P:24 SP:F9
	// F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:26 P:24 SP:F9
	// F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:26 P:24 SP:F9
	// F9B5  C9 55     CMP #$55                        A:55 X:55 Y:26 P:24 SP:F9
	// F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:26 P:27 SP:F9
	// F9B9  60        RTS                             A:55 X:55 Y:26 P:27 SP:F9
	// DDE8  C8        INY                             A:55 X:55 Y:26 P:27 SP:FB
	// DDE9  20 BD F9  JSR $F9BD                       A:55 X:55 Y:27 P:25 SP:FB
	// F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:27 P:25 SP:F9
	// F9BF  38        SEC                             A:55 X:55 Y:27 P:E5 SP:F9
	// F9C0  A9 80     LDA #$80                        A:55 X:55 Y:27 P:E5 SP:F9
	// F9C2  60        RTS                             A:80 X:55 Y:27 P:E5 SP:F9
	// DDEC  95 00     STA $00,X @ 55 = 55             A:80 X:55 Y:27 P:E5 SP:FB
	// DDEE  16 00     ASL $00,X @ 55 = 80             A:80 X:55 Y:27 P:E5 SP:FB
	// DDF0  B5 00     LDA $00,X @ 55 = 00             A:80 X:55 Y:27 P:67 SP:FB
	// DDF2  20 C3 F9  JSR $F9C3                       A:00 X:55 Y:27 P:67 SP:FB
	// F9C3  90 1C     BCC $F9E1                       A:00 X:55 Y:27 P:67 SP:F9
	// F9C5  D0 1A     BNE $F9E1                       A:00 X:55 Y:27 P:67 SP:F9
	// F9C7  30 18     BMI $F9E1                       A:00 X:55 Y:27 P:67 SP:F9
	// F9C9  50 16     BVC $F9E1                       A:00 X:55 Y:27 P:67 SP:F9
	// F9CB  C9 00     CMP #$00                        A:00 X:55 Y:27 P:67 SP:F9
	// F9CD  D0 12     BNE $F9E1                       A:00 X:55 Y:27 P:67 SP:F9
	// F9CF  B8        CLV                             A:00 X:55 Y:27 P:67 SP:F9
	// F9D0  A9 55     LDA #$55                        A:00 X:55 Y:27 P:27 SP:F9
	// F9D2  38        SEC                             A:55 X:55 Y:27 P:25 SP:F9
	// F9D3  60        RTS                             A:55 X:55 Y:27 P:25 SP:F9
	// DDF5  C8        INY                             A:55 X:55 Y:27 P:25 SP:FB
	// DDF6  95 00     STA $00,X @ 55 = 00             A:55 X:55 Y:28 P:25 SP:FB
	// DDF8  16 00     ASL $00,X @ 55 = 55             A:55 X:55 Y:28 P:25 SP:FB
	// DDFA  B5 00     LDA $00,X @ 55 = AA             A:55 X:55 Y:28 P:A4 SP:FB
	// DDFC  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:28 P:A4 SP:FB
	// F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9
	// F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9
	// F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9
	// F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:28 P:A4 SP:F9
	// F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:28 P:A4 SP:F9
	// F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:28 P:27 SP:F9
	// F9E0  60        RTS                             A:AA X:55 Y:28 P:27 SP:F9
	// DDFF  C8        INY                             A:AA X:55 Y:28 P:27 SP:FB
	// DE00  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:29 P:25 SP:FB
	// F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:29 P:25 SP:F9
	// F9E6  38        SEC                             A:AA X:55 Y:29 P:E5 SP:F9
	// F9E7  A9 01     LDA #$01                        A:AA X:55 Y:29 P:E5 SP:F9
	// F9E9  60        RTS                             A:01 X:55 Y:29 P:65 SP:F9
	// DE03  95 00     STA $00,X @ 55 = AA             A:01 X:55 Y:29 P:65 SP:FB
	// DE05  76 00     ROR $00,X @ 55 = 01             A:01 X:55 Y:29 P:65 SP:FB
	// DE07  B5 00     LDA $00,X @ 55 = 80             A:01 X:55 Y:29 P:E5 SP:FB
	// DE09  20 EA F9  JSR $F9EA                       A:80 X:55 Y:29 P:E5 SP:FB
	// F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:29 P:E5 SP:F9
	// F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:29 P:E5 SP:F9
	// F9EE  10 18     BPL $FA08                       A:80 X:55 Y:29 P:E5 SP:F9
	// F9F0  50 16     BVC $FA08                       A:80 X:55 Y:29 P:E5 SP:F9
	// F9F2  C9 80     CMP #$80                        A:80 X:55 Y:29 P:E5 SP:F9
	// F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:29 P:67 SP:F9
	// F9F6  B8        CLV                             A:80 X:55 Y:29 P:67 SP:F9
	// F9F7  18        CLC                             A:80 X:55 Y:29 P:27 SP:F9
	// F9F8  A9 55     LDA #$55                        A:80 X:55 Y:29 P:26 SP:F9
	// F9FA  60        RTS                             A:55 X:55 Y:29 P:24 SP:F9
	// DE0C  C8        INY                             A:55 X:55 Y:29 P:24 SP:FB
	// DE0D  95 00     STA $00,X @ 55 = 80             A:55 X:55 Y:2A P:24 SP:FB
	// DE0F  76 00     ROR $00,X @ 55 = 55             A:55 X:55 Y:2A P:24 SP:FB
	// DE11  B5 00     LDA $00,X @ 55 = 2A             A:55 X:55 Y:2A P:25 SP:FB
	// DE13  20 FB F9  JSR $F9FB                       A:2A X:55 Y:2A P:25 SP:FB
	// F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:2A P:25 SP:F9
	// F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:2A P:25 SP:F9
	// F9FF  30 07     BMI $FA08                       A:2A X:55 Y:2A P:25 SP:F9
	// FA01  70 05     BVS $FA08                       A:2A X:55 Y:2A P:25 SP:F9
	// FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:2A P:25 SP:F9
	// FA05  D0 01     BNE $FA08                       A:2A X:55 Y:2A P:27 SP:F9
	// FA07  60        RTS                             A:2A X:55 Y:2A P:27 SP:F9
	// DE16  C8        INY                             A:2A X:55 Y:2A P:27 SP:FB
	// DE17  20 0A FA  JSR $FA0A                       A:2A X:55 Y:2B P:25 SP:FB
	// FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:2B P:25 SP:F9
	// FA0C  38        SEC                             A:2A X:55 Y:2B P:E5 SP:F9
	// FA0D  A9 80     LDA #$80                        A:2A X:55 Y:2B P:E5 SP:F9
	// FA0F  60        RTS                             A:80 X:55 Y:2B P:E5 SP:F9
	// DE1A  95 00     STA $00,X @ 55 = 2A             A:80 X:55 Y:2B P:E5 SP:FB
	// DE1C  36 00     ROL $00,X @ 55 = 80             A:80 X:55 Y:2B P:E5 SP:FB
	// DE1E  B5 00     LDA $00,X @ 55 = 01             A:80 X:55 Y:2B P:65 SP:FB
	// DE20  20 10 FA  JSR $FA10                       A:01 X:55 Y:2B P:65 SP:FB
	// FA10  90 1C     BCC $FA2E                       A:01 X:55 Y:2B P:65 SP:F9
	// FA12  F0 1A     BEQ $FA2E                       A:01 X:55 Y:2B P:65 SP:F9
	// FA14  30 18     BMI $FA2E                       A:01 X:55 Y:2B P:65 SP:F9
	// FA16  50 16     BVC $FA2E                       A:01 X:55 Y:2B P:65 SP:F9
	// FA18  C9 01     CMP #$01                        A:01 X:55 Y:2B P:65 SP:F9
	// FA1A  D0 12     BNE $FA2E                       A:01 X:55 Y:2B P:67 SP:F9
	// FA1C  B8        CLV                             A:01 X:55 Y:2B P:67 SP:F9
	// FA1D  18        CLC                             A:01 X:55 Y:2B P:27 SP:F9
	// FA1E  A9 55     LDA #$55                        A:01 X:55 Y:2B P:26 SP:F9
	// FA20  60        RTS                             A:55 X:55 Y:2B P:24 SP:F9
	// DE23  C8        INY                             A:55 X:55 Y:2B P:24 SP:FB
	// DE24  95 00     STA $00,X @ 55 = 01             A:55 X:55 Y:2C P:24 SP:FB
	// DE26  36 00     ROL $00,X @ 55 = 55             A:55 X:55 Y:2C P:24 SP:FB
	// DE28  B5 00     LDA $00,X @ 55 = AA             A:55 X:55 Y:2C P:A4 SP:FB
	// DE2A  20 21 FA  JSR $FA21                       A:AA X:55 Y:2C P:A4 SP:FB
	// FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9
	// FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9
	// FA25  10 07     BPL $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9
	// FA27  70 05     BVS $FA2E                       A:AA X:55 Y:2C P:A4 SP:F9
	// FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:2C P:A4 SP:F9
	// FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:2C P:27 SP:F9
	// FA2D  60        RTS                             A:AA X:55 Y:2C P:27 SP:F9
	// DE2D  A9 FF     LDA #$FF                        A:AA X:55 Y:2C P:27 SP:FB
	// DE2F  95 00     STA $00,X @ 55 = AA             A:FF X:55 Y:2C P:A5 SP:FB
	// DE31  85 01     STA $01 = FF                    A:FF X:55 Y:2C P:A5 SP:FB
	// DE33  24 01     BIT $01 = FF                    A:FF X:55 Y:2C P:A5 SP:FB
	// DE35  38        SEC                             A:FF X:55 Y:2C P:E5 SP:FB
	// DE36  F6 00     INC $00,X @ 55 = FF             A:FF X:55 Y:2C P:E5 SP:FB
	// DE38  D0 0C     BNE $DE46                       A:FF X:55 Y:2C P:67 SP:FB
	// DE3A  30 0A     BMI $DE46                       A:FF X:55 Y:2C P:67 SP:FB
	// DE3C  50 08     BVC $DE46                       A:FF X:55 Y:2C P:67 SP:FB
	// DE3E  90 06     BCC $DE46                       A:FF X:55 Y:2C P:67 SP:FB
	// DE40  B5 00     LDA $00,X @ 55 = 00             A:FF X:55 Y:2C P:67 SP:FB
	// DE42  C9 00     CMP #$00                        A:00 X:55 Y:2C P:67 SP:FB
	// DE44  F0 04     BEQ $DE4A                       A:00 X:55 Y:2C P:67 SP:FB
	// DE4A  A9 7F     LDA #$7F                        A:00 X:55 Y:2C P:67 SP:FB
	// DE4C  95 00     STA $00,X @ 55 = 00             A:7F X:55 Y:2C P:65 SP:FB
	// DE4E  B8        CLV                             A:7F X:55 Y:2C P:65 SP:FB
	// DE4F  18        CLC                             A:7F X:55 Y:2C P:25 SP:FB
	// DE50  F6 00     INC $00,X @ 55 = 7F             A:7F X:55 Y:2C P:24 SP:FB
	// DE52  F0 0C     BEQ $DE60                       A:7F X:55 Y:2C P:A4 SP:FB
	// DE54  10 0A     BPL $DE60                       A:7F X:55 Y:2C P:A4 SP:FB
	// DE56  70 08     BVS $DE60                       A:7F X:55 Y:2C P:A4 SP:FB
	// DE58  B0 06     BCS $DE60                       A:7F X:55 Y:2C P:A4 SP:FB
	// DE5A  B5 00     LDA $00,X @ 55 = 80             A:7F X:55 Y:2C P:A4 SP:FB
	// DE5C  C9 80     CMP #$80                        A:80 X:55 Y:2C P:A4 SP:FB
	// DE5E  F0 04     BEQ $DE64                       A:80 X:55 Y:2C P:27 SP:FB
	// DE64  A9 00     LDA #$00                        A:80 X:55 Y:2C P:27 SP:FB
	// DE66  95 00     STA $00,X @ 55 = 80             A:00 X:55 Y:2C P:27 SP:FB
	// DE68  24 01     BIT $01 = FF                    A:00 X:55 Y:2C P:27 SP:FB
	// DE6A  38        SEC                             A:00 X:55 Y:2C P:E7 SP:FB
	// DE6B  D6 00     DEC $00,X @ 55 = 00             A:00 X:55 Y:2C P:E7 SP:FB
	// DE6D  F0 0C     BEQ $DE7B                       A:00 X:55 Y:2C P:E5 SP:FB
	// DE6F  10 0A     BPL $DE7B                       A:00 X:55 Y:2C P:E5 SP:FB
	// DE71  50 08     BVC $DE7B                       A:00 X:55 Y:2C P:E5 SP:FB
	// DE73  90 06     BCC $DE7B                       A:00 X:55 Y:2C P:E5 SP:FB
	// DE75  B5 00     LDA $00,X @ 55 = FF             A:00 X:55 Y:2C P:E5 SP:FB
	// DE77  C9 FF     CMP #$FF                        A:FF X:55 Y:2C P:E5 SP:FB
	// DE79  F0 04     BEQ $DE7F                       A:FF X:55 Y:2C P:67 SP:FB
	// DE7F  A9 80     LDA #$80                        A:FF X:55 Y:2C P:67 SP:FB
	// DE81  95 00     STA $00,X @ 55 = FF             A:80 X:55 Y:2C P:E5 SP:FB
	// DE83  B8        CLV                             A:80 X:55 Y:2C P:E5 SP:FB
	// DE84  18        CLC                             A:80 X:55 Y:2C P:A5 SP:FB
	// DE85  D6 00     DEC $00,X @ 55 = 80             A:80 X:55 Y:2C P:A4 SP:FB
	// DE87  F0 0C     BEQ $DE95                       A:80 X:55 Y:2C P:24 SP:FB
	// DE89  30 0A     BMI $DE95                       A:80 X:55 Y:2C P:24 SP:FB
	// DE8B  70 08     BVS $DE95                       A:80 X:55 Y:2C P:24 SP:FB
	// DE8D  B0 06     BCS $DE95                       A:80 X:55 Y:2C P:24 SP:FB
	// DE8F  B5 00     LDA $00,X @ 55 = 7F             A:80 X:55 Y:2C P:24 SP:FB
	// DE91  C9 7F     CMP #$7F                        A:7F X:55 Y:2C P:24 SP:FB
	// DE93  F0 04     BEQ $DE99                       A:7F X:55 Y:2C P:27 SP:FB
	// DE99  A9 01     LDA #$01                        A:7F X:55 Y:2C P:27 SP:FB
	// DE9B  95 00     STA $00,X @ 55 = 7F             A:01 X:55 Y:2C P:25 SP:FB
	// DE9D  D6 00     DEC $00,X @ 55 = 01             A:01 X:55 Y:2C P:25 SP:FB
	// DE9F  F0 04     BEQ $DEA5                       A:01 X:55 Y:2C P:27 SP:FB
	// DEA5  A9 33     LDA #$33                        A:01 X:55 Y:2C P:27 SP:FB
	// DEA7  85 78     STA $78 = 7F                    A:33 X:55 Y:2C P:25 SP:FB
	// DEA9  A9 44     LDA #$44                        A:33 X:55 Y:2C P:25 SP:FB
	// DEAB  A0 78     LDY #$78                        A:44 X:55 Y:2C P:25 SP:FB
	// DEAD  A2 00     LDX #$00                        A:44 X:55 Y:78 P:25 SP:FB
	// DEAF  38        SEC                             A:44 X:00 Y:78 P:27 SP:FB
	// DEB0  24 01     BIT $01 = FF                    A:44 X:00 Y:78 P:27 SP:FB
	// DEB2  B6 00     LDX $00,Y @ 78 = 33             A:44 X:00 Y:78 P:E5 SP:FB
	// DEB4  90 12     BCC $DEC8                       A:44 X:33 Y:78 P:65 SP:FB
	// DEB6  50 10     BVC $DEC8                       A:44 X:33 Y:78 P:65 SP:FB
	// DEB8  30 0E     BMI $DEC8                       A:44 X:33 Y:78 P:65 SP:FB
	// DEBA  F0 0C     BEQ $DEC8                       A:44 X:33 Y:78 P:65 SP:FB
	// DEBC  E0 33     CPX #$33                        A:44 X:33 Y:78 P:65 SP:FB
	// DEBE  D0 08     BNE $DEC8                       A:44 X:33 Y:78 P:67 SP:FB
	// DEC0  C0 78     CPY #$78                        A:44 X:33 Y:78 P:67 SP:FB
	// DEC2  D0 04     BNE $DEC8                       A:44 X:33 Y:78 P:67 SP:FB
	// DEC4  C9 44     CMP #$44                        A:44 X:33 Y:78 P:67 SP:FB
	// DEC6  F0 04     BEQ $DECC                       A:44 X:33 Y:78 P:67 SP:FB
	// DECC  A9 97     LDA #$97                        A:44 X:33 Y:78 P:67 SP:FB
	// DECE  85 7F     STA $7F = 00                    A:97 X:33 Y:78 P:E5 SP:FB
	// DED0  A9 47     LDA #$47                        A:97 X:33 Y:78 P:E5 SP:FB
	// DED2  A0 FF     LDY #$FF                        A:47 X:33 Y:78 P:65 SP:FB
	// DED4  A2 00     LDX #$00                        A:47 X:33 Y:FF P:E5 SP:FB
	// DED6  18        CLC                             A:47 X:00 Y:FF P:67 SP:FB
	// DED7  B8        CLV                             A:47 X:00 Y:FF P:66 SP:FB
	// DED8  B6 80     LDX $80,Y @ 7F = 97             A:47 X:00 Y:FF P:26 SP:FB
	// DEDA  B0 12     BCS $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB
	// DEDC  70 10     BVS $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB
	// DEDE  10 0E     BPL $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB
	// DEE0  F0 0C     BEQ $DEEE                       A:47 X:97 Y:FF P:A4 SP:FB
	// DEE2  E0 97     CPX #$97                        A:47 X:97 Y:FF P:A4 SP:FB
	// DEE4  D0 08     BNE $DEEE                       A:47 X:97 Y:FF P:27 SP:FB
	// DEE6  C0 FF     CPY #$FF                        A:47 X:97 Y:FF P:27 SP:FB
	// DEE8  D0 04     BNE $DEEE                       A:47 X:97 Y:FF P:27 SP:FB
	// DEEA  C9 47     CMP #$47                        A:47 X:97 Y:FF P:27 SP:FB
	// DEEC  F0 04     BEQ $DEF2                       A:47 X:97 Y:FF P:27 SP:FB
	// DEF2  A9 00     LDA #$00                        A:47 X:97 Y:FF P:27 SP:FB
	// DEF4  85 7F     STA $7F = 97                    A:00 X:97 Y:FF P:27 SP:FB
	// DEF6  A9 47     LDA #$47                        A:00 X:97 Y:FF P:27 SP:FB
	// DEF8  A0 FF     LDY #$FF                        A:47 X:97 Y:FF P:25 SP:FB
	// DEFA  A2 69     LDX #$69                        A:47 X:97 Y:FF P:A5 SP:FB
	// DEFC  18        CLC                             A:47 X:69 Y:FF P:25 SP:FB
	// DEFD  B8        CLV                             A:47 X:69 Y:FF P:24 SP:FB
	// DEFE  96 80     STX $80,Y @ 7F = 00             A:47 X:69 Y:FF P:24 SP:FB
	// DF00  B0 18     BCS $DF1A                       A:47 X:69 Y:FF P:24 SP:FB
	// DF02  70 16     BVS $DF1A                       A:47 X:69 Y:FF P:24 SP:FB
	// DF04  30 14     BMI $DF1A                       A:47 X:69 Y:FF P:24 SP:FB
	// DF06  F0 12     BEQ $DF1A                       A:47 X:69 Y:FF P:24 SP:FB
	// DF08  E0 69     CPX #$69                        A:47 X:69 Y:FF P:24 SP:FB
	// DF0A  D0 0E     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB
	// DF0C  C0 FF     CPY #$FF                        A:47 X:69 Y:FF P:27 SP:FB
	// DF0E  D0 0A     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB
	// DF10  C9 47     CMP #$47                        A:47 X:69 Y:FF P:27 SP:FB
	// DF12  D0 06     BNE $DF1A                       A:47 X:69 Y:FF P:27 SP:FB
	// DF14  A5 7F     LDA $7F = 69                    A:47 X:69 Y:FF P:27 SP:FB
	// DF16  C9 69     CMP #$69                        A:69 X:69 Y:FF P:25 SP:FB
	// DF18  F0 04     BEQ $DF1E                       A:69 X:69 Y:FF P:27 SP:FB
	// DF1E  A9 F5     LDA #$F5                        A:69 X:69 Y:FF P:27 SP:FB
	// DF20  85 4F     STA $4F = 00                    A:F5 X:69 Y:FF P:A5 SP:FB
	// DF22  A9 47     LDA #$47                        A:F5 X:69 Y:FF P:A5 SP:FB
	// DF24  A0 4F     LDY #$4F                        A:47 X:69 Y:FF P:25 SP:FB
	// DF26  24 01     BIT $01 = FF                    A:47 X:69 Y:4F P:25 SP:FB
	// DF28  A2 00     LDX #$00                        A:47 X:69 Y:4F P:E5 SP:FB
	// DF2A  38        SEC                             A:47 X:00 Y:4F P:67 SP:FB
	// DF2B  96 00     STX $00,Y @ 4F = F5             A:47 X:00 Y:4F P:67 SP:FB
	// DF2D  90 16     BCC $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF2F  50 14     BVC $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF31  30 12     BMI $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF33  D0 10     BNE $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF35  E0 00     CPX #$00                        A:47 X:00 Y:4F P:67 SP:FB
	// DF37  D0 0C     BNE $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF39  C0 4F     CPY #$4F                        A:47 X:00 Y:4F P:67 SP:FB
	// DF3B  D0 08     BNE $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF3D  C9 47     CMP #$47                        A:47 X:00 Y:4F P:67 SP:FB
	// DF3F  D0 04     BNE $DF45                       A:47 X:00 Y:4F P:67 SP:FB
	// DF41  A5 4F     LDA $4F = 00                    A:47 X:00 Y:4F P:67 SP:FB
	// DF43  F0 04     BEQ $DF49                       A:00 X:00 Y:4F P:67 SP:FB
	// DF49  60        RTS                             A:00 X:00 Y:4F P:67 SP:FB
	// C62C  20 AA E1  JSR $E1AA                       A:00 X:00 Y:4F P:67 SP:FD
	// E1AA  A9 FF     LDA #$FF                        A:00 X:00 Y:4F P:67 SP:FB
	// E1AC  85 01     STA $01 = FF                    A:FF X:00 Y:4F P:E5 SP:FB
	// E1AE  A9 AA     LDA #$AA                        A:FF X:00 Y:4F P:E5 SP:FB
	// E1B0  8D 33 06  STA $0633 = 00                  A:AA X:00 Y:4F P:E5 SP:FB
	// E1B3  A9 BB     LDA #$BB                        A:AA X:00 Y:4F P:E5 SP:FB
	// E1B5  8D 89 06  STA $0689 = 00                  A:BB X:00 Y:4F P:E5 SP:FB
	// E1B8  A2 00     LDX #$00                        A:BB X:00 Y:4F P:E5 SP:FB
	// E1BA  A9 66     LDA #$66                        A:BB X:00 Y:4F P:67 SP:FB
	// E1BC  24 01     BIT $01 = FF                    A:66 X:00 Y:4F P:65 SP:FB
	// E1BE  38        SEC                             A:66 X:00 Y:4F P:E5 SP:FB
	// E1BF  A0 00     LDY #$00                        A:66 X:00 Y:4F P:E5 SP:FB
	// E1C1  BC 33 06  LDY $0633,X @ 0633 = AA         A:66 X:00 Y:00 P:67 SP:FB
	// E1C4  10 12     BPL $E1D8                       A:66 X:00 Y:AA P:E5 SP:FB
	// E1C6  F0 10     BEQ $E1D8                       A:66 X:00 Y:AA P:E5 SP:FB
	// E1C8  50 0E     BVC $E1D8                       A:66 X:00 Y:AA P:E5 SP:FB
	// E1CA  90 0C     BCC $E1D8                       A:66 X:00 Y:AA P:E5 SP:FB
	// E1CC  C9 66     CMP #$66                        A:66 X:00 Y:AA P:E5 SP:FB
	// E1CE  D0 08     BNE $E1D8                       A:66 X:00 Y:AA P:67 SP:FB
	// E1D0  E0 00     CPX #$00                        A:66 X:00 Y:AA P:67 SP:FB
	// E1D2  D0 04     BNE $E1D8                       A:66 X:00 Y:AA P:67 SP:FB
	// E1D4  C0 AA     CPY #$AA                        A:66 X:00 Y:AA P:67 SP:FB
	// E1D6  F0 04     BEQ $E1DC                       A:66 X:00 Y:AA P:67 SP:FB
	// E1DC  A2 8A     LDX #$8A                        A:66 X:00 Y:AA P:67 SP:FB
	// E1DE  A9 66     LDA #$66                        A:66 X:8A Y:AA P:E5 SP:FB
	// E1E0  B8        CLV                             A:66 X:8A Y:AA P:65 SP:FB
	// E1E1  18        CLC                             A:66 X:8A Y:AA P:25 SP:FB
	// E1E2  A0 00     LDY #$00                        A:66 X:8A Y:AA P:24 SP:FB
	// E1E4  BC FF 05  LDY $05FF,X @ 0689 = BB         A:66 X:8A Y:00 P:26 SP:FB
	// E1E7  10 12     BPL $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB
	// E1E9  F0 10     BEQ $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB
	// E1EB  70 0E     BVS $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB
	// E1ED  B0 0C     BCS $E1FB                       A:66 X:8A Y:BB P:A4 SP:FB
	// E1EF  C0 BB     CPY #$BB                        A:66 X:8A Y:BB P:A4 SP:FB
	// E1F1  D0 08     BNE $E1FB                       A:66 X:8A Y:BB P:27 SP:FB
	// E1F3  C9 66     CMP #$66                        A:66 X:8A Y:BB P:27 SP:FB
	// E1F5  D0 04     BNE $E1FB                       A:66 X:8A Y:BB P:27 SP:FB
	// E1F7  E0 8A     CPX #$8A                        A:66 X:8A Y:BB P:27 SP:FB
	// E1F9  F0 04     BEQ $E1FF                       A:66 X:8A Y:BB P:27 SP:FB
	// E1FF  A0 53     LDY #$53                        A:66 X:8A Y:BB P:27 SP:FB
	// E201  A9 AA     LDA #$AA                        A:66 X:8A Y:53 P:25 SP:FB
	// E203  A2 78     LDX #$78                        A:AA X:8A Y:53 P:A5 SP:FB
	// E205  8D 78 06  STA $0678 = 00                  A:AA X:78 Y:53 P:25 SP:FB
	// E208  20 B6 F7  JSR $F7B6                       A:AA X:78 Y:53 P:25 SP:FB
	// F7B6  18        CLC                             A:AA X:78 Y:53 P:25 SP:F9
	// F7B7  A9 FF     LDA #$FF                        A:AA X:78 Y:53 P:24 SP:F9
	// F7B9  85 01     STA $01 = FF                    A:FF X:78 Y:53 P:A4 SP:F9
	// F7BB  24 01     BIT $01 = FF                    A:FF X:78 Y:53 P:A4 SP:F9
	// F7BD  A9 55     LDA #$55                        A:FF X:78 Y:53 P:E4 SP:F9
	// F7BF  60        RTS                             A:55 X:78 Y:53 P:64 SP:F9
	// E20B  1D 00 06  ORA $0600,X @ 0678 = AA         A:55 X:78 Y:53 P:64 SP:FB
	// E20E  20 C0 F7  JSR $F7C0                       A:FF X:78 Y:53 P:E4 SP:FB
	// F7C0  B0 09     BCS $F7CB                       A:FF X:78 Y:53 P:E4 SP:F9
	// F7C2  10 07     BPL $F7CB                       A:FF X:78 Y:53 P:E4 SP:F9
	// F7C4  C9 FF     CMP #$FF                        A:FF X:78 Y:53 P:E4 SP:F9
	// F7C6  D0 03     BNE $F7CB                       A:FF X:78 Y:53 P:67 SP:F9
	// F7C8  50 01     BVC $F7CB                       A:FF X:78 Y:53 P:67 SP:F9
	// F7CA  60        RTS                             A:FF X:78 Y:53 P:67 SP:F9
	// E211  C8        INY                             A:FF X:78 Y:53 P:67 SP:FB
	// E212  A9 00     LDA #$00                        A:FF X:78 Y:54 P:65 SP:FB
	// E214  8D 78 06  STA $0678 = AA                  A:00 X:78 Y:54 P:67 SP:FB
	// E217  20 CE F7  JSR $F7CE                       A:00 X:78 Y:54 P:67 SP:FB
	// F7CE  38        SEC                             A:00 X:78 Y:54 P:67 SP:F9
	// F7CF  B8        CLV                             A:00 X:78 Y:54 P:67 SP:F9
	// F7D0  A9 00     LDA #$00                        A:00 X:78 Y:54 P:27 SP:F9
	// F7D2  60        RTS                             A:00 X:78 Y:54 P:27 SP:F9
	// E21A  1D 00 06  ORA $0600,X @ 0678 = 00         A:00 X:78 Y:54 P:27 SP:FB
	// E21D  20 D3 F7  JSR $F7D3                       A:00 X:78 Y:54 P:27 SP:FB
	// F7D3  D0 07     BNE $F7DC                       A:00 X:78 Y:54 P:27 SP:F9
	// F7D5  70 05     BVS $F7DC                       A:00 X:78 Y:54 P:27 SP:F9
	// F7D7  90 03     BCC $F7DC                       A:00 X:78 Y:54 P:27 SP:F9
	// F7D9  30 01     BMI $F7DC                       A:00 X:78 Y:54 P:27 SP:F9
	// F7DB  60        RTS                             A:00 X:78 Y:54 P:27 SP:F9
	// E220  C8        INY                             A:00 X:78 Y:54 P:27 SP:FB
	// E221  A9 AA     LDA #$AA                        A:00 X:78 Y:55 P:25 SP:FB
	// E223  8D 78 06  STA $0678 = 00                  A:AA X:78 Y:55 P:A5 SP:FB
	// E226  20 DF F7  JSR $F7DF                       A:AA X:78 Y:55 P:A5 SP:FB
	// F7DF  18        CLC                             A:AA X:78 Y:55 P:A5 SP:F9
	// F7E0  24 01     BIT $01 = FF                    A:AA X:78 Y:55 P:A4 SP:F9
	// F7E2  A9 55     LDA #$55                        A:AA X:78 Y:55 P:E4 SP:F9
	// F7E4  60        RTS                             A:55 X:78 Y:55 P:64 SP:F9
	// E229  3D 00 06  AND $0600,X @ 0678 = AA         A:55 X:78 Y:55 P:64 SP:FB
	// E22C  20 E5 F7  JSR $F7E5                       A:00 X:78 Y:55 P:66 SP:FB
	// F7E5  D0 07     BNE $F7EE                       A:00 X:78 Y:55 P:66 SP:F9
	// F7E7  50 05     BVC $F7EE                       A:00 X:78 Y:55 P:66 SP:F9
	// F7E9  B0 03     BCS $F7EE                       A:00 X:78 Y:55 P:66 SP:F9
	// F7EB  30 01     BMI $F7EE                       A:00 X:78 Y:55 P:66 SP:F9
	// F7ED  60        RTS                             A:00 X:78 Y:55 P:66 SP:F9
	// E22F  C8        INY                             A:00 X:78 Y:55 P:66 SP:FB
	// E230  A9 EF     LDA #$EF                        A:00 X:78 Y:56 P:64 SP:FB
	// E232  8D 78 06  STA $0678 = AA                  A:EF X:78 Y:56 P:E4 SP:FB
	// E235  20 F1 F7  JSR $F7F1                       A:EF X:78 Y:56 P:E4 SP:FB
	// F7F1  38        SEC                             A:EF X:78 Y:56 P:E4 SP:F9
	// F7F2  B8        CLV                             A:EF X:78 Y:56 P:E5 SP:F9
	// F7F3  A9 F8     LDA #$F8                        A:EF X:78 Y:56 P:A5 SP:F9
	// F7F5  60        RTS                             A:F8 X:78 Y:56 P:A5 SP:F9
	// E238  3D 00 06  AND $0600,X @ 0678 = EF         A:F8 X:78 Y:56 P:A5 SP:FB
	// E23B  20 F6 F7  JSR $F7F6                       A:E8 X:78 Y:56 P:A5 SP:FB
	// F7F6  90 09     BCC $F801                       A:E8 X:78 Y:56 P:A5 SP:F9
	// F7F8  10 07     BPL $F801                       A:E8 X:78 Y:56 P:A5 SP:F9
	// F7FA  C9 E8     CMP #$E8                        A:E8 X:78 Y:56 P:A5 SP:F9
	// F7FC  D0 03     BNE $F801                       A:E8 X:78 Y:56 P:27 SP:F9
	// F7FE  70 01     BVS $F801                       A:E8 X:78 Y:56 P:27 SP:F9
	// F800  60        RTS                             A:E8 X:78 Y:56 P:27 SP:F9
	// E23E  C8        INY                             A:E8 X:78 Y:56 P:27 SP:FB
	// E23F  A9 AA     LDA #$AA                        A:E8 X:78 Y:57 P:25 SP:FB
	// E241  8D 78 06  STA $0678 = EF                  A:AA X:78 Y:57 P:A5 SP:FB
	// E244  20 04 F8  JSR $F804                       A:AA X:78 Y:57 P:A5 SP:FB
	// F804  18        CLC                             A:AA X:78 Y:57 P:A5 SP:F9
	// F805  24 01     BIT $01 = FF                    A:AA X:78 Y:57 P:A4 SP:F9
	// F807  A9 5F     LDA #$5F                        A:AA X:78 Y:57 P:E4 SP:F9
	// F809  60        RTS                             A:5F X:78 Y:57 P:64 SP:F9
	// E247  5D 00 06  EOR $0600,X @ 0678 = AA         A:5F X:78 Y:57 P:64 SP:FB
	// E24A  20 0A F8  JSR $F80A                       A:F5 X:78 Y:57 P:E4 SP:FB
	// F80A  B0 09     BCS $F815                       A:F5 X:78 Y:57 P:E4 SP:F9
	// F80C  10 07     BPL $F815                       A:F5 X:78 Y:57 P:E4 SP:F9
	// F80E  C9 F5     CMP #$F5                        A:F5 X:78 Y:57 P:E4 SP:F9
	// F810  D0 03     BNE $F815                       A:F5 X:78 Y:57 P:67 SP:F9
	// F812  50 01     BVC $F815                       A:F5 X:78 Y:57 P:67 SP:F9
	// F814  60        RTS                             A:F5 X:78 Y:57 P:67 SP:F9
	// E24D  C8        INY                             A:F5 X:78 Y:57 P:67 SP:FB
	// E24E  A9 70     LDA #$70                        A:F5 X:78 Y:58 P:65 SP:FB
	// E250  8D 78 06  STA $0678 = AA                  A:70 X:78 Y:58 P:65 SP:FB
	// E253  20 18 F8  JSR $F818                       A:70 X:78 Y:58 P:65 SP:FB
	// F818  38        SEC                             A:70 X:78 Y:58 P:65 SP:F9
	// F819  B8        CLV                             A:70 X:78 Y:58 P:65 SP:F9
	// F81A  A9 70     LDA #$70                        A:70 X:78 Y:58 P:25 SP:F9
	// F81C  60        RTS                             A:70 X:78 Y:58 P:25 SP:F9
	// E256  5D 00 06  EOR $0600,X @ 0678 = 70         A:70 X:78 Y:58 P:25 SP:FB
	// E259  20 1D F8  JSR $F81D                       A:00 X:78 Y:58 P:27 SP:FB
	// F81D  D0 07     BNE $F826                       A:00 X:78 Y:58 P:27 SP:F9
	// F81F  70 05     BVS $F826                       A:00 X:78 Y:58 P:27 SP:F9
	// F821  90 03     BCC $F826                       A:00 X:78 Y:58 P:27 SP:F9
	// F823  30 01     BMI $F826                       A:00 X:78 Y:58 P:27 SP:F9
	// F825  60        RTS                             A:00 X:78 Y:58 P:27 SP:F9
	// E25C  C8        INY                             A:00 X:78 Y:58 P:27 SP:FB
	// E25D  A9 69     LDA #$69                        A:00 X:78 Y:59 P:25 SP:FB
	// E25F  8D 78 06  STA $0678 = 70                  A:69 X:78 Y:59 P:25 SP:FB
	// E262  20 29 F8  JSR $F829                       A:69 X:78 Y:59 P:25 SP:FB
	// F829  18        CLC                             A:69 X:78 Y:59 P:25 SP:F9
	// F82A  24 01     BIT $01 = FF                    A:69 X:78 Y:59 P:24 SP:F9
	// F82C  A9 00     LDA #$00                        A:69 X:78 Y:59 P:E4 SP:F9
	// F82E  60        RTS                             A:00 X:78 Y:59 P:66 SP:F9
	// E265  7D 00 06  ADC $0600,X @ 0678 = 69         A:00 X:78 Y:59 P:66 SP:FB
	// E268  20 2F F8  JSR $F82F                       A:69 X:78 Y:59 P:24 SP:FB
	// F82F  30 09     BMI $F83A                       A:69 X:78 Y:59 P:24 SP:F9
	// F831  B0 07     BCS $F83A                       A:69 X:78 Y:59 P:24 SP:F9
	// F833  C9 69     CMP #$69                        A:69 X:78 Y:59 P:24 SP:F9
	// F835  D0 03     BNE $F83A                       A:69 X:78 Y:59 P:27 SP:F9
	// F837  70 01     BVS $F83A                       A:69 X:78 Y:59 P:27 SP:F9
	// F839  60        RTS                             A:69 X:78 Y:59 P:27 SP:F9
	// E26B  C8        INY                             A:69 X:78 Y:59 P:27 SP:FB
	// E26C  20 3D F8  JSR $F83D                       A:69 X:78 Y:5A P:25 SP:FB
	// F83D  38        SEC                             A:69 X:78 Y:5A P:25 SP:F9
	// F83E  24 01     BIT $01 = FF                    A:69 X:78 Y:5A P:25 SP:F9
	// F840  A9 00     LDA #$00                        A:69 X:78 Y:5A P:E5 SP:F9
	// F842  60        RTS                             A:00 X:78 Y:5A P:67 SP:F9
	// E26F  7D 00 06  ADC $0600,X @ 0678 = 69         A:00 X:78 Y:5A P:67 SP:FB
	// E272  20 43 F8  JSR $F843                       A:6A X:78 Y:5A P:24 SP:FB
	// F843  30 09     BMI $F84E                       A:6A X:78 Y:5A P:24 SP:F9
	// F845  B0 07     BCS $F84E                       A:6A X:78 Y:5A P:24 SP:F9
	// F847  C9 6A     CMP #$6A                        A:6A X:78 Y:5A P:24 SP:F9
	// F849  D0 03     BNE $F84E                       A:6A X:78 Y:5A P:27 SP:F9
	// F84B  70 01     BVS $F84E                       A:6A X:78 Y:5A P:27 SP:F9
	// F84D  60        RTS                             A:6A X:78 Y:5A P:27 SP:F9
	// E275  C8        INY                             A:6A X:78 Y:5A P:27 SP:FB
	// E276  A9 7F     LDA #$7F                        A:6A X:78 Y:5B P:25 SP:FB
	// E278  8D 78 06  STA $0678 = 69                  A:7F X:78 Y:5B P:25 SP:FB
	// E27B  20 51 F8  JSR $F851                       A:7F X:78 Y:5B P:25 SP:FB
	// F851  38        SEC                             A:7F X:78 Y:5B P:25 SP:F9
	// F852  B8        CLV                             A:7F X:78 Y:5B P:25 SP:F9
	// F853  A9 7F     LDA #$7F                        A:7F X:78 Y:5B P:25 SP:F9
	// F855  60        RTS                             A:7F X:78 Y:5B P:25 SP:F9
	// E27E  7D 00 06  ADC $0600,X @ 0678 = 7F         A:7F X:78 Y:5B P:25 SP:FB
	// E281  20 56 F8  JSR $F856                       A:FF X:78 Y:5B P:E4 SP:FB
	// F856  10 09     BPL $F861                       A:FF X:78 Y:5B P:E4 SP:F9
	// F858  B0 07     BCS $F861                       A:FF X:78 Y:5B P:E4 SP:F9
	// F85A  C9 FF     CMP #$FF                        A:FF X:78 Y:5B P:E4 SP:F9
	// F85C  D0 03     BNE $F861                       A:FF X:78 Y:5B P:67 SP:F9
	// F85E  50 01     BVC $F861                       A:FF X:78 Y:5B P:67 SP:F9
	// F860  60        RTS                             A:FF X:78 Y:5B P:67 SP:F9
	// E284  C8        INY                             A:FF X:78 Y:5B P:67 SP:FB
	// E285  A9 80     LDA #$80                        A:FF X:78 Y:5C P:65 SP:FB
	// E287  8D 78 06  STA $0678 = 7F                  A:80 X:78 Y:5C P:E5 SP:FB
	// E28A  20 64 F8  JSR $F864                       A:80 X:78 Y:5C P:E5 SP:FB
	// F864  18        CLC                             A:80 X:78 Y:5C P:E5 SP:F9
	// F865  24 01     BIT $01 = FF                    A:80 X:78 Y:5C P:E4 SP:F9
	// F867  A9 7F     LDA #$7F                        A:80 X:78 Y:5C P:E4 SP:F9
	// F869  60        RTS                             A:7F X:78 Y:5C P:64 SP:F9
	// E28D  7D 00 06  ADC $0600,X @ 0678 = 80         A:7F X:78 Y:5C P:64 SP:FB
	// E290  20 6A F8  JSR $F86A                       A:FF X:78 Y:5C P:A4 SP:FB
	// F86A  10 09     BPL $F875                       A:FF X:78 Y:5C P:A4 SP:F9
	// F86C  B0 07     BCS $F875                       A:FF X:78 Y:5C P:A4 SP:F9
	// F86E  C9 FF     CMP #$FF                        A:FF X:78 Y:5C P:A4 SP:F9
	// F870  D0 03     BNE $F875                       A:FF X:78 Y:5C P:27 SP:F9
	// F872  70 01     BVS $F875                       A:FF X:78 Y:5C P:27 SP:F9
	// F874  60        RTS                             A:FF X:78 Y:5C P:27 SP:F9
	// E293  C8        INY                             A:FF X:78 Y:5C P:27 SP:FB
	// E294  20 78 F8  JSR $F878                       A:FF X:78 Y:5D P:25 SP:FB
	// F878  38        SEC                             A:FF X:78 Y:5D P:25 SP:F9
	// F879  B8        CLV                             A:FF X:78 Y:5D P:25 SP:F9
	// F87A  A9 7F     LDA #$7F                        A:FF X:78 Y:5D P:25 SP:F9
	// F87C  60        RTS                             A:7F X:78 Y:5D P:25 SP:F9
	// E297  7D 00 06  ADC $0600,X @ 0678 = 80         A:7F X:78 Y:5D P:25 SP:FB
	// E29A  20 7D F8  JSR $F87D                       A:00 X:78 Y:5D P:27 SP:FB
	// F87D  D0 07     BNE $F886                       A:00 X:78 Y:5D P:27 SP:F9
	// F87F  30 05     BMI $F886                       A:00 X:78 Y:5D P:27 SP:F9
	// F881  70 03     BVS $F886                       A:00 X:78 Y:5D P:27 SP:F9
	// F883  90 01     BCC $F886                       A:00 X:78 Y:5D P:27 SP:F9
	// F885  60        RTS                             A:00 X:78 Y:5D P:27 SP:F9
	// E29D  C8        INY                             A:00 X:78 Y:5D P:27 SP:FB
	// E29E  A9 40     LDA #$40                        A:00 X:78 Y:5E P:25 SP:FB
	// E2A0  8D 78 06  STA $0678 = 80                  A:40 X:78 Y:5E P:25 SP:FB
	// E2A3  20 89 F8  JSR $F889                       A:40 X:78 Y:5E P:25 SP:FB
	// F889  24 01     BIT $01 = FF                    A:40 X:78 Y:5E P:25 SP:F9
	// F88B  A9 40     LDA #$40                        A:40 X:78 Y:5E P:E5 SP:F9
	// F88D  60        RTS                             A:40 X:78 Y:5E P:65 SP:F9
	// E2A6  DD 00 06  CMP $0600,X @ 0678 = 40         A:40 X:78 Y:5E P:65 SP:FB
	// E2A9  20 8E F8  JSR $F88E                       A:40 X:78 Y:5E P:67 SP:FB
	// F88E  30 07     BMI $F897                       A:40 X:78 Y:5E P:67 SP:F9
	// F890  90 05     BCC $F897                       A:40 X:78 Y:5E P:67 SP:F9
	// F892  D0 03     BNE $F897                       A:40 X:78 Y:5E P:67 SP:F9
	// F894  50 01     BVC $F897                       A:40 X:78 Y:5E P:67 SP:F9
	// F896  60        RTS                             A:40 X:78 Y:5E P:67 SP:F9
	// E2AC  C8        INY                             A:40 X:78 Y:5E P:67 SP:FB
	// E2AD  48        PHA                             A:40 X:78 Y:5F P:65 SP:FB
	// E2AE  A9 3F     LDA #$3F                        A:40 X:78 Y:5F P:65 SP:FA
	// E2B0  8D 78 06  STA $0678 = 40                  A:3F X:78 Y:5F P:65 SP:FA
	// E2B3  68        PLA                             A:3F X:78 Y:5F P:65 SP:FA
	// E2B4  20 9A F8  JSR $F89A                       A:40 X:78 Y:5F P:65 SP:FB
	// F89A  B8        CLV                             A:40 X:78 Y:5F P:65 SP:F9
	// F89B  60        RTS                             A:40 X:78 Y:5F P:25 SP:F9
	// E2B7  DD 00 06  CMP $0600,X @ 0678 = 3F         A:40 X:78 Y:5F P:25 SP:FB
	// E2BA  20 9C F8  JSR $F89C                       A:40 X:78 Y:5F P:25 SP:FB
	// F89C  F0 07     BEQ $F8A5                       A:40 X:78 Y:5F P:25 SP:F9
	// F89E  30 05     BMI $F8A5                       A:40 X:78 Y:5F P:25 SP:F9
	// F8A0  90 03     BCC $F8A5                       A:40 X:78 Y:5F P:25 SP:F9
	// F8A2  70 01     BVS $F8A5                       A:40 X:78 Y:5F P:25 SP:F9
	// F8A4  60        RTS                             A:40 X:78 Y:5F P:25 SP:F9
	// E2BD  C8        INY                             A:40 X:78 Y:5F P:25 SP:FB
	// E2BE  48        PHA                             A:40 X:78 Y:60 P:25 SP:FB
	// E2BF  A9 41     LDA #$41                        A:40 X:78 Y:60 P:25 SP:FA
	// E2C1  8D 78 06  STA $0678 = 3F                  A:41 X:78 Y:60 P:25 SP:FA
	// E2C4  68        PLA                             A:41 X:78 Y:60 P:25 SP:FA
	// E2C5  DD 00 06  CMP $0600,X @ 0678 = 41         A:40 X:78 Y:60 P:25 SP:FB
	// E2C8  20 A8 F8  JSR $F8A8                       A:40 X:78 Y:60 P:A4 SP:FB
	// F8A8  F0 05     BEQ $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9
	// F8AA  10 03     BPL $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9
	// F8AC  10 01     BPL $F8AF                       A:40 X:78 Y:60 P:A4 SP:F9
	// F8AE  60        RTS                             A:40 X:78 Y:60 P:A4 SP:F9
	// E2CB  C8        INY                             A:40 X:78 Y:60 P:A4 SP:FB
	// E2CC  48        PHA                             A:40 X:78 Y:61 P:24 SP:FB
	// E2CD  A9 00     LDA #$00                        A:40 X:78 Y:61 P:24 SP:FA
	// E2CF  8D 78 06  STA $0678 = 41                  A:00 X:78 Y:61 P:26 SP:FA
	// E2D2  68        PLA                             A:00 X:78 Y:61 P:26 SP:FA
	// E2D3  20 B2 F8  JSR $F8B2                       A:40 X:78 Y:61 P:24 SP:FB
	// F8B2  A9 80     LDA #$80                        A:40 X:78 Y:61 P:24 SP:F9
	// F8B4  60        RTS                             A:80 X:78 Y:61 P:A4 SP:F9
	// E2D6  DD 00 06  CMP $0600,X @ 0678 = 00         A:80 X:78 Y:61 P:A4 SP:FB
	// E2D9  20 B5 F8  JSR $F8B5                       A:80 X:78 Y:61 P:A5 SP:FB
	// F8B5  F0 05     BEQ $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9
	// F8B7  10 03     BPL $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9
	// F8B9  90 01     BCC $F8BC                       A:80 X:78 Y:61 P:A5 SP:F9
	// F8BB  60        RTS                             A:80 X:78 Y:61 P:A5 SP:F9
	// E2DC  C8        INY                             A:80 X:78 Y:61 P:A5 SP:FB
	// E2DD  48        PHA                             A:80 X:78 Y:62 P:25 SP:FB
	// E2DE  A9 80     LDA #$80                        A:80 X:78 Y:62 P:25 SP:FA
	// E2E0  8D 78 06  STA $0678 = 00                  A:80 X:78 Y:62 P:A5 SP:FA
	// E2E3  68        PLA                             A:80 X:78 Y:62 P:A5 SP:FA
	// E2E4  DD 00 06  CMP $0600,X @ 0678 = 80         A:80 X:78 Y:62 P:A5 SP:FB
	// E2E7  20 BF F8  JSR $F8BF                       A:80 X:78 Y:62 P:27 SP:FB
	// F8BF  D0 05     BNE $F8C6                       A:80 X:78 Y:62 P:27 SP:F9
	// F8C1  30 03     BMI $F8C6                       A:80 X:78 Y:62 P:27 SP:F9
	// F8C3  90 01     BCC $F8C6                       A:80 X:78 Y:62 P:27 SP:F9
	// F8C5  60        RTS                             A:80 X:78 Y:62 P:27 SP:F9
	// E2EA  C8        INY                             A:80 X:78 Y:62 P:27 SP:FB
	// E2EB  48        PHA                             A:80 X:78 Y:63 P:25 SP:FB
	// E2EC  A9 81     LDA #$81                        A:80 X:78 Y:63 P:25 SP:FA
	// E2EE  8D 78 06  STA $0678 = 80                  A:81 X:78 Y:63 P:A5 SP:FA
	// E2F1  68        PLA                             A:81 X:78 Y:63 P:A5 SP:FA
	// E2F2  DD 00 06  CMP $0600,X @ 0678 = 81         A:80 X:78 Y:63 P:A5 SP:FB
	// E2F5  20 C9 F8  JSR $F8C9                       A:80 X:78 Y:63 P:A4 SP:FB
	// F8C9  B0 05     BCS $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9
	// F8CB  F0 03     BEQ $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9
	// F8CD  10 01     BPL $F8D0                       A:80 X:78 Y:63 P:A4 SP:F9
	// F8CF  60        RTS                             A:80 X:78 Y:63 P:A4 SP:F9
	// E2F8  C8        INY                             A:80 X:78 Y:63 P:A4 SP:FB
	// E2F9  48        PHA                             A:80 X:78 Y:64 P:24 SP:FB
	// E2FA  A9 7F     LDA #$7F                        A:80 X:78 Y:64 P:24 SP:FA
	// E2FC  8D 78 06  STA $0678 = 81                  A:7F X:78 Y:64 P:24 SP:FA
	// E2FF  68        PLA                             A:7F X:78 Y:64 P:24 SP:FA
	// E300  DD 00 06  CMP $0600,X @ 0678 = 7F         A:80 X:78 Y:64 P:A4 SP:FB
	// E303  20 D3 F8  JSR $F8D3                       A:80 X:78 Y:64 P:25 SP:FB
	// F8D3  90 05     BCC $F8DA                       A:80 X:78 Y:64 P:25 SP:F9
	// F8D5  F0 03     BEQ $F8DA                       A:80 X:78 Y:64 P:25 SP:F9
	// F8D7  30 01     BMI $F8DA                       A:80 X:78 Y:64 P:25 SP:F9
	// F8D9  60        RTS                             A:80 X:78 Y:64 P:25 SP:F9
	// E306  C8        INY                             A:80 X:78 Y:64 P:25 SP:FB
	// E307  A9 40     LDA #$40                        A:80 X:78 Y:65 P:25 SP:FB
	// E309  8D 78 06  STA $0678 = 7F                  A:40 X:78 Y:65 P:25 SP:FB
	// E30C  20 31 F9  JSR $F931                       A:40 X:78 Y:65 P:25 SP:FB
	// F931  24 01     BIT $01 = FF                    A:40 X:78 Y:65 P:25 SP:F9
	// F933  A9 40     LDA #$40                        A:40 X:78 Y:65 P:E5 SP:F9
	// F935  38        SEC                             A:40 X:78 Y:65 P:65 SP:F9
	// F936  60        RTS                             A:40 X:78 Y:65 P:65 SP:F9
	// E30F  FD 00 06  SBC $0600,X @ 0678 = 40         A:40 X:78 Y:65 P:65 SP:FB
	// E312  20 37 F9  JSR $F937                       A:00 X:78 Y:65 P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:78 Y:65 P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:78 Y:65 P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:78 Y:65 P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:78 Y:65 P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:78 Y:65 P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:78 Y:65 P:27 SP:F9
	// F943  60        RTS                             A:00 X:78 Y:65 P:27 SP:F9
	// E315  C8        INY                             A:00 X:78 Y:65 P:27 SP:FB
	// E316  A9 3F     LDA #$3F                        A:00 X:78 Y:66 P:25 SP:FB
	// E318  8D 78 06  STA $0678 = 40                  A:3F X:78 Y:66 P:25 SP:FB
	// E31B  20 47 F9  JSR $F947                       A:3F X:78 Y:66 P:25 SP:FB
	// F947  B8        CLV                             A:3F X:78 Y:66 P:25 SP:F9
	// F948  38        SEC                             A:3F X:78 Y:66 P:25 SP:F9
	// F949  A9 40     LDA #$40                        A:3F X:78 Y:66 P:25 SP:F9
	// F94B  60        RTS                             A:40 X:78 Y:66 P:25 SP:F9
	// E31E  FD 00 06  SBC $0600,X @ 0678 = 3F         A:40 X:78 Y:66 P:25 SP:FB
	// E321  20 4C F9  JSR $F94C                       A:01 X:78 Y:66 P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:78 Y:66 P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:78 Y:66 P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:78 Y:66 P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:78 Y:66 P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:78 Y:66 P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:78 Y:66 P:27 SP:F9
	// F958  60        RTS                             A:01 X:78 Y:66 P:27 SP:F9
	// E324  C8        INY                             A:01 X:78 Y:66 P:27 SP:FB
	// E325  A9 41     LDA #$41                        A:01 X:78 Y:67 P:25 SP:FB
	// E327  8D 78 06  STA $0678 = 3F                  A:41 X:78 Y:67 P:25 SP:FB
	// E32A  20 5C F9  JSR $F95C                       A:41 X:78 Y:67 P:25 SP:FB
	// F95C  A9 40     LDA #$40                        A:41 X:78 Y:67 P:25 SP:F9
	// F95E  38        SEC                             A:40 X:78 Y:67 P:25 SP:F9
	// F95F  24 01     BIT $01 = FF                    A:40 X:78 Y:67 P:25 SP:F9
	// F961  60        RTS                             A:40 X:78 Y:67 P:E5 SP:F9
	// E32D  FD 00 06  SBC $0600,X @ 0678 = 41         A:40 X:78 Y:67 P:E5 SP:FB
	// E330  20 62 F9  JSR $F962                       A:FF X:78 Y:67 P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:78 Y:67 P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:78 Y:67 P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:78 Y:67 P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:78 Y:67 P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:78 Y:67 P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:78 Y:67 P:27 SP:F9
	// F96E  60        RTS                             A:FF X:78 Y:67 P:27 SP:F9
	// E333  C8        INY                             A:FF X:78 Y:67 P:27 SP:FB
	// E334  A9 00     LDA #$00                        A:FF X:78 Y:68 P:25 SP:FB
	// E336  8D 78 06  STA $0678 = 41                  A:00 X:78 Y:68 P:27 SP:FB
	// E339  20 72 F9  JSR $F972                       A:00 X:78 Y:68 P:27 SP:FB
	// F972  18        CLC                             A:00 X:78 Y:68 P:27 SP:F9
	// F973  A9 80     LDA #$80                        A:00 X:78 Y:68 P:26 SP:F9
	// F975  60        RTS                             A:80 X:78 Y:68 P:A4 SP:F9
	// E33C  FD 00 06  SBC $0600,X @ 0678 = 00         A:80 X:78 Y:68 P:A4 SP:FB
	// E33F  20 76 F9  JSR $F976                       A:7F X:78 Y:68 P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:78 Y:68 P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:78 Y:68 P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:78 Y:68 P:67 SP:F9
	// F97C  60        RTS                             A:7F X:78 Y:68 P:67 SP:F9
	// E342  C8        INY                             A:7F X:78 Y:68 P:67 SP:FB
	// E343  A9 7F     LDA #$7F                        A:7F X:78 Y:69 P:65 SP:FB
	// E345  8D 78 06  STA $0678 = 00                  A:7F X:78 Y:69 P:65 SP:FB
	// E348  20 80 F9  JSR $F980                       A:7F X:78 Y:69 P:65 SP:FB
	// F980  38        SEC                             A:7F X:78 Y:69 P:65 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:78 Y:69 P:65 SP:F9
	// F983  60        RTS                             A:81 X:78 Y:69 P:E5 SP:F9
	// E34B  FD 00 06  SBC $0600,X @ 0678 = 7F         A:81 X:78 Y:69 P:E5 SP:FB
	// E34E  20 84 F9  JSR $F984                       A:02 X:78 Y:69 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:78 Y:69 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:78 Y:69 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:78 Y:69 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:78 Y:69 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:78 Y:69 P:67 SP:F9
	// E351  A9 AA     LDA #$AA                        A:02 X:78 Y:69 P:67 SP:FB
	// E353  8D 33 06  STA $0633 = AA                  A:AA X:78 Y:69 P:E5 SP:FB
	// E356  A9 BB     LDA #$BB                        A:AA X:78 Y:69 P:E5 SP:FB
	// E358  8D 89 06  STA $0689 = BB                  A:BB X:78 Y:69 P:E5 SP:FB
	// E35B  A2 00     LDX #$00                        A:BB X:78 Y:69 P:E5 SP:FB
	// E35D  A0 66     LDY #$66                        A:BB X:00 Y:69 P:67 SP:FB
	// E35F  24 01     BIT $01 = FF                    A:BB X:00 Y:66 P:65 SP:FB
	// E361  38        SEC                             A:BB X:00 Y:66 P:E5 SP:FB
	// E362  A9 00     LDA #$00                        A:BB X:00 Y:66 P:E5 SP:FB
	// E364  BD 33 06  LDA $0633,X @ 0633 = AA         A:00 X:00 Y:66 P:67 SP:FB
	// E367  10 12     BPL $E37B                       A:AA X:00 Y:66 P:E5 SP:FB
	// E369  F0 10     BEQ $E37B                       A:AA X:00 Y:66 P:E5 SP:FB
	// E36B  50 0E     BVC $E37B                       A:AA X:00 Y:66 P:E5 SP:FB
	// E36D  90 0C     BCC $E37B                       A:AA X:00 Y:66 P:E5 SP:FB
	// E36F  C0 66     CPY #$66                        A:AA X:00 Y:66 P:E5 SP:FB
	// E371  D0 08     BNE $E37B                       A:AA X:00 Y:66 P:67 SP:FB
	// E373  E0 00     CPX #$00                        A:AA X:00 Y:66 P:67 SP:FB
	// E375  D0 04     BNE $E37B                       A:AA X:00 Y:66 P:67 SP:FB
	// E377  C9 AA     CMP #$AA                        A:AA X:00 Y:66 P:67 SP:FB
	// E379  F0 04     BEQ $E37F                       A:AA X:00 Y:66 P:67 SP:FB
	// E37F  A2 8A     LDX #$8A                        A:AA X:00 Y:66 P:67 SP:FB
	// E381  A0 66     LDY #$66                        A:AA X:8A Y:66 P:E5 SP:FB
	// E383  B8        CLV                             A:AA X:8A Y:66 P:65 SP:FB
	// E384  18        CLC                             A:AA X:8A Y:66 P:25 SP:FB
	// E385  A9 00     LDA #$00                        A:AA X:8A Y:66 P:24 SP:FB
	// E387  BD FF 05  LDA $05FF,X @ 0689 = BB         A:00 X:8A Y:66 P:26 SP:FB
	// E38A  10 12     BPL $E39E                       A:BB X:8A Y:66 P:A4 SP:FB
	// E38C  F0 10     BEQ $E39E                       A:BB X:8A Y:66 P:A4 SP:FB
	// E38E  70 0E     BVS $E39E                       A:BB X:8A Y:66 P:A4 SP:FB
	// E390  B0 0C     BCS $E39E                       A:BB X:8A Y:66 P:A4 SP:FB
	// E392  C9 BB     CMP #$BB                        A:BB X:8A Y:66 P:A4 SP:FB
	// E394  D0 08     BNE $E39E                       A:BB X:8A Y:66 P:27 SP:FB
	// E396  C0 66     CPY #$66                        A:BB X:8A Y:66 P:27 SP:FB
	// E398  D0 04     BNE $E39E                       A:BB X:8A Y:66 P:27 SP:FB
	// E39A  E0 8A     CPX #$8A                        A:BB X:8A Y:66 P:27 SP:FB
	// E39C  F0 04     BEQ $E3A2                       A:BB X:8A Y:66 P:27 SP:FB
	// E3A2  24 01     BIT $01 = FF                    A:BB X:8A Y:66 P:27 SP:FB
	// E3A4  38        SEC                             A:BB X:8A Y:66 P:E5 SP:FB
	// E3A5  A9 44     LDA #$44                        A:BB X:8A Y:66 P:E5 SP:FB
	// E3A7  A2 00     LDX #$00                        A:44 X:8A Y:66 P:65 SP:FB
	// E3A9  9D 33 06  STA $0633,X @ 0633 = AA         A:44 X:00 Y:66 P:67 SP:FB
	// E3AC  AD 33 06  LDA $0633 = 44                  A:44 X:00 Y:66 P:67 SP:FB
	// E3AF  90 1A     BCC $E3CB                       A:44 X:00 Y:66 P:65 SP:FB
	// E3B1  C9 44     CMP #$44                        A:44 X:00 Y:66 P:65 SP:FB
	// E3B3  D0 16     BNE $E3CB                       A:44 X:00 Y:66 P:67 SP:FB
	// E3B5  50 14     BVC $E3CB                       A:44 X:00 Y:66 P:67 SP:FB
	// E3B7  18        CLC                             A:44 X:00 Y:66 P:67 SP:FB
	// E3B8  B8        CLV                             A:44 X:00 Y:66 P:66 SP:FB
	// E3B9  A9 99     LDA #$99                        A:44 X:00 Y:66 P:26 SP:FB
	// E3BB  A2 80     LDX #$80                        A:99 X:00 Y:66 P:A4 SP:FB
	// E3BD  9D 85 05  STA $0585,X @ 0605 = 00         A:99 X:80 Y:66 P:A4 SP:FB
	// E3C0  AD 05 06  LDA $0605 = 99                  A:99 X:80 Y:66 P:A4 SP:FB
	// E3C3  B0 06     BCS $E3CB                       A:99 X:80 Y:66 P:A4 SP:FB
	// E3C5  C9 99     CMP #$99                        A:99 X:80 Y:66 P:A4 SP:FB
	// E3C7  D0 02     BNE $E3CB                       A:99 X:80 Y:66 P:27 SP:FB
	// E3C9  50 04     BVC $E3CF                       A:99 X:80 Y:66 P:27 SP:FB
	// E3CF  A0 6D     LDY #$6D                        A:99 X:80 Y:66 P:27 SP:FB
	// E3D1  A2 6D     LDX #$6D                        A:99 X:80 Y:6D P:25 SP:FB
	// E3D3  20 90 F9  JSR $F990                       A:99 X:6D Y:6D P:25 SP:FB
	// F990  A2 55     LDX #$55                        A:99 X:6D Y:6D P:25 SP:F9
	// F992  A9 FF     LDA #$FF                        A:99 X:55 Y:6D P:25 SP:F9
	// F994  85 01     STA $01 = FF                    A:FF X:55 Y:6D P:A5 SP:F9
	// F996  EA        NOP                             A:FF X:55 Y:6D P:A5 SP:F9
	// F997  24 01     BIT $01 = FF                    A:FF X:55 Y:6D P:A5 SP:F9
	// F999  38        SEC                             A:FF X:55 Y:6D P:E5 SP:F9
	// F99A  A9 01     LDA #$01                        A:FF X:55 Y:6D P:E5 SP:F9
	// F99C  60        RTS                             A:01 X:55 Y:6D P:65 SP:F9
	// E3D6  9D 00 06  STA $0600,X @ 0655 = 00         A:01 X:55 Y:6D P:65 SP:FB
	// E3D9  5E 00 06  LSR $0600,X @ 0655 = 01         A:01 X:55 Y:6D P:65 SP:FB
	// E3DC  BD 00 06  LDA $0600,X @ 0655 = 00         A:01 X:55 Y:6D P:67 SP:FB
	// E3DF  20 9D F9  JSR $F99D                       A:00 X:55 Y:6D P:67 SP:FB
	// F99D  90 1B     BCC $F9BA                       A:00 X:55 Y:6D P:67 SP:F9
	// F99F  D0 19     BNE $F9BA                       A:00 X:55 Y:6D P:67 SP:F9
	// F9A1  30 17     BMI $F9BA                       A:00 X:55 Y:6D P:67 SP:F9
	// F9A3  50 15     BVC $F9BA                       A:00 X:55 Y:6D P:67 SP:F9
	// F9A5  C9 00     CMP #$00                        A:00 X:55 Y:6D P:67 SP:F9
	// F9A7  D0 11     BNE $F9BA                       A:00 X:55 Y:6D P:67 SP:F9
	// F9A9  B8        CLV                             A:00 X:55 Y:6D P:67 SP:F9
	// F9AA  A9 AA     LDA #$AA                        A:00 X:55 Y:6D P:27 SP:F9
	// F9AC  60        RTS                             A:AA X:55 Y:6D P:A5 SP:F9
	// E3E2  C8        INY                             A:AA X:55 Y:6D P:A5 SP:FB
	// E3E3  9D 00 06  STA $0600,X @ 0655 = 00         A:AA X:55 Y:6E P:25 SP:FB
	// E3E6  5E 00 06  LSR $0600,X @ 0655 = AA         A:AA X:55 Y:6E P:25 SP:FB
	// E3E9  BD 00 06  LDA $0600,X @ 0655 = 55         A:AA X:55 Y:6E P:24 SP:FB
	// E3EC  20 AD F9  JSR $F9AD                       A:55 X:55 Y:6E P:24 SP:FB
	// F9AD  B0 0B     BCS $F9BA                       A:55 X:55 Y:6E P:24 SP:F9
	// F9AF  F0 09     BEQ $F9BA                       A:55 X:55 Y:6E P:24 SP:F9
	// F9B1  30 07     BMI $F9BA                       A:55 X:55 Y:6E P:24 SP:F9
	// F9B3  70 05     BVS $F9BA                       A:55 X:55 Y:6E P:24 SP:F9
	// F9B5  C9 55     CMP #$55                        A:55 X:55 Y:6E P:24 SP:F9
	// F9B7  D0 01     BNE $F9BA                       A:55 X:55 Y:6E P:27 SP:F9
	// F9B9  60        RTS                             A:55 X:55 Y:6E P:27 SP:F9
	// E3EF  C8        INY                             A:55 X:55 Y:6E P:27 SP:FB
	// E3F0  20 BD F9  JSR $F9BD                       A:55 X:55 Y:6F P:25 SP:FB
	// F9BD  24 01     BIT $01 = FF                    A:55 X:55 Y:6F P:25 SP:F9
	// F9BF  38        SEC                             A:55 X:55 Y:6F P:E5 SP:F9
	// F9C0  A9 80     LDA #$80                        A:55 X:55 Y:6F P:E5 SP:F9
	// F9C2  60        RTS                             A:80 X:55 Y:6F P:E5 SP:F9
	// E3F3  9D 00 06  STA $0600,X @ 0655 = 55         A:80 X:55 Y:6F P:E5 SP:FB
	// E3F6  1E 00 06  ASL $0600,X @ 0655 = 80         A:80 X:55 Y:6F P:E5 SP:FB
	// E3F9  BD 00 06  LDA $0600,X @ 0655 = 00         A:80 X:55 Y:6F P:67 SP:FB
	// E3FC  20 C3 F9  JSR $F9C3                       A:00 X:55 Y:6F P:67 SP:FB
	// F9C3  90 1C     BCC $F9E1                       A:00 X:55 Y:6F P:67 SP:F9
	// F9C5  D0 1A     BNE $F9E1                       A:00 X:55 Y:6F P:67 SP:F9
	// F9C7  30 18     BMI $F9E1                       A:00 X:55 Y:6F P:67 SP:F9
	// F9C9  50 16     BVC $F9E1                       A:00 X:55 Y:6F P:67 SP:F9
	// F9CB  C9 00     CMP #$00                        A:00 X:55 Y:6F P:67 SP:F9
	// F9CD  D0 12     BNE $F9E1                       A:00 X:55 Y:6F P:67 SP:F9
	// F9CF  B8        CLV                             A:00 X:55 Y:6F P:67 SP:F9
	// F9D0  A9 55     LDA #$55                        A:00 X:55 Y:6F P:27 SP:F9
	// F9D2  38        SEC                             A:55 X:55 Y:6F P:25 SP:F9
	// F9D3  60        RTS                             A:55 X:55 Y:6F P:25 SP:F9
	// E3FF  C8        INY                             A:55 X:55 Y:6F P:25 SP:FB
	// E400  9D 00 06  STA $0600,X @ 0655 = 00         A:55 X:55 Y:70 P:25 SP:FB
	// E403  1E 00 06  ASL $0600,X @ 0655 = 55         A:55 X:55 Y:70 P:25 SP:FB
	// E406  BD 00 06  LDA $0600,X @ 0655 = AA         A:55 X:55 Y:70 P:A4 SP:FB
	// E409  20 D4 F9  JSR $F9D4                       A:AA X:55 Y:70 P:A4 SP:FB
	// F9D4  B0 0B     BCS $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9
	// F9D6  F0 09     BEQ $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9
	// F9D8  10 07     BPL $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9
	// F9DA  70 05     BVS $F9E1                       A:AA X:55 Y:70 P:A4 SP:F9
	// F9DC  C9 AA     CMP #$AA                        A:AA X:55 Y:70 P:A4 SP:F9
	// F9DE  D0 01     BNE $F9E1                       A:AA X:55 Y:70 P:27 SP:F9
	// F9E0  60        RTS                             A:AA X:55 Y:70 P:27 SP:F9
	// E40C  C8        INY                             A:AA X:55 Y:70 P:27 SP:FB
	// E40D  20 E4 F9  JSR $F9E4                       A:AA X:55 Y:71 P:25 SP:FB
	// F9E4  24 01     BIT $01 = FF                    A:AA X:55 Y:71 P:25 SP:F9
	// F9E6  38        SEC                             A:AA X:55 Y:71 P:E5 SP:F9
	// F9E7  A9 01     LDA #$01                        A:AA X:55 Y:71 P:E5 SP:F9
	// F9E9  60        RTS                             A:01 X:55 Y:71 P:65 SP:F9
	// E410  9D 00 06  STA $0600,X @ 0655 = AA         A:01 X:55 Y:71 P:65 SP:FB
	// E413  7E 00 06  ROR $0600,X @ 0655 = 01         A:01 X:55 Y:71 P:65 SP:FB
	// E416  BD 00 06  LDA $0600,X @ 0655 = 80         A:01 X:55 Y:71 P:E5 SP:FB
	// E419  20 EA F9  JSR $F9EA                       A:80 X:55 Y:71 P:E5 SP:FB
	// F9EA  90 1C     BCC $FA08                       A:80 X:55 Y:71 P:E5 SP:F9
	// F9EC  F0 1A     BEQ $FA08                       A:80 X:55 Y:71 P:E5 SP:F9
	// F9EE  10 18     BPL $FA08                       A:80 X:55 Y:71 P:E5 SP:F9
	// F9F0  50 16     BVC $FA08                       A:80 X:55 Y:71 P:E5 SP:F9
	// F9F2  C9 80     CMP #$80                        A:80 X:55 Y:71 P:E5 SP:F9
	// F9F4  D0 12     BNE $FA08                       A:80 X:55 Y:71 P:67 SP:F9
	// F9F6  B8        CLV                             A:80 X:55 Y:71 P:67 SP:F9
	// F9F7  18        CLC                             A:80 X:55 Y:71 P:27 SP:F9
	// F9F8  A9 55     LDA #$55                        A:80 X:55 Y:71 P:26 SP:F9
	// F9FA  60        RTS                             A:55 X:55 Y:71 P:24 SP:F9
	// E41C  C8        INY                             A:55 X:55 Y:71 P:24 SP:FB
	// E41D  9D 00 06  STA $0600,X @ 0655 = 80         A:55 X:55 Y:72 P:24 SP:FB
	// E420  7E 00 06  ROR $0600,X @ 0655 = 55         A:55 X:55 Y:72 P:24 SP:FB
	// E423  BD 00 06  LDA $0600,X @ 0655 = 2A         A:55 X:55 Y:72 P:25 SP:FB
	// E426  20 FB F9  JSR $F9FB                       A:2A X:55 Y:72 P:25 SP:FB
	// F9FB  90 0B     BCC $FA08                       A:2A X:55 Y:72 P:25 SP:F9
	// F9FD  F0 09     BEQ $FA08                       A:2A X:55 Y:72 P:25 SP:F9
	// F9FF  30 07     BMI $FA08                       A:2A X:55 Y:72 P:25 SP:F9
	// FA01  70 05     BVS $FA08                       A:2A X:55 Y:72 P:25 SP:F9
	// FA03  C9 2A     CMP #$2A                        A:2A X:55 Y:72 P:25 SP:F9
	// FA05  D0 01     BNE $FA08                       A:2A X:55 Y:72 P:27 SP:F9
	// FA07  60        RTS                             A:2A X:55 Y:72 P:27 SP:F9
	// E429  C8        INY                             A:2A X:55 Y:72 P:27 SP:FB
	// E42A  20 0A FA  JSR $FA0A                       A:2A X:55 Y:73 P:25 SP:FB
	// FA0A  24 01     BIT $01 = FF                    A:2A X:55 Y:73 P:25 SP:F9
	// FA0C  38        SEC                             A:2A X:55 Y:73 P:E5 SP:F9
	// FA0D  A9 80     LDA #$80                        A:2A X:55 Y:73 P:E5 SP:F9
	// FA0F  60        RTS                             A:80 X:55 Y:73 P:E5 SP:F9
	// E42D  9D 00 06  STA $0600,X @ 0655 = 2A         A:80 X:55 Y:73 P:E5 SP:FB
	// E430  3E 00 06  ROL $0600,X @ 0655 = 80         A:80 X:55 Y:73 P:E5 SP:FB
	// E433  BD 00 06  LDA $0600,X @ 0655 = 01         A:80 X:55 Y:73 P:65 SP:FB
	// E436  20 10 FA  JSR $FA10                       A:01 X:55 Y:73 P:65 SP:FB
	// FA10  90 1C     BCC $FA2E                       A:01 X:55 Y:73 P:65 SP:F9
	// FA12  F0 1A     BEQ $FA2E                       A:01 X:55 Y:73 P:65 SP:F9
	// FA14  30 18     BMI $FA2E                       A:01 X:55 Y:73 P:65 SP:F9
	// FA16  50 16     BVC $FA2E                       A:01 X:55 Y:73 P:65 SP:F9
	// FA18  C9 01     CMP #$01                        A:01 X:55 Y:73 P:65 SP:F9
	// FA1A  D0 12     BNE $FA2E                       A:01 X:55 Y:73 P:67 SP:F9
	// FA1C  B8        CLV                             A:01 X:55 Y:73 P:67 SP:F9
	// FA1D  18        CLC                             A:01 X:55 Y:73 P:27 SP:F9
	// FA1E  A9 55     LDA #$55                        A:01 X:55 Y:73 P:26 SP:F9
	// FA20  60        RTS                             A:55 X:55 Y:73 P:24 SP:F9
	// E439  C8        INY                             A:55 X:55 Y:73 P:24 SP:FB
	// E43A  9D 00 06  STA $0600,X @ 0655 = 01         A:55 X:55 Y:74 P:24 SP:FB
	// E43D  3E 00 06  ROL $0600,X @ 0655 = 55         A:55 X:55 Y:74 P:24 SP:FB
	// E440  BD 00 06  LDA $0600,X @ 0655 = AA         A:55 X:55 Y:74 P:A4 SP:FB
	// E443  20 21 FA  JSR $FA21                       A:AA X:55 Y:74 P:A4 SP:FB
	// FA21  B0 0B     BCS $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9
	// FA23  F0 09     BEQ $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9
	// FA25  10 07     BPL $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9
	// FA27  70 05     BVS $FA2E                       A:AA X:55 Y:74 P:A4 SP:F9
	// FA29  C9 AA     CMP #$AA                        A:AA X:55 Y:74 P:A4 SP:F9
	// FA2B  D0 01     BNE $FA2E                       A:AA X:55 Y:74 P:27 SP:F9
	// FA2D  60        RTS                             A:AA X:55 Y:74 P:27 SP:F9
	// E446  A9 FF     LDA #$FF                        A:AA X:55 Y:74 P:27 SP:FB
	// E448  9D 00 06  STA $0600,X @ 0655 = AA         A:FF X:55 Y:74 P:A5 SP:FB
	// E44B  85 01     STA $01 = FF                    A:FF X:55 Y:74 P:A5 SP:FB
	// E44D  24 01     BIT $01 = FF                    A:FF X:55 Y:74 P:A5 SP:FB
	// E44F  38        SEC                             A:FF X:55 Y:74 P:E5 SP:FB
	// E450  FE 00 06  INC $0600,X @ 0655 = FF         A:FF X:55 Y:74 P:E5 SP:FB
	// E453  D0 0D     BNE $E462                       A:FF X:55 Y:74 P:67 SP:FB
	// E455  30 0B     BMI $E462                       A:FF X:55 Y:74 P:67 SP:FB
	// E457  50 09     BVC $E462                       A:FF X:55 Y:74 P:67 SP:FB
	// E459  90 07     BCC $E462                       A:FF X:55 Y:74 P:67 SP:FB
	// E45B  BD 00 06  LDA $0600,X @ 0655 = 00         A:FF X:55 Y:74 P:67 SP:FB
	// E45E  C9 00     CMP #$00                        A:00 X:55 Y:74 P:67 SP:FB
	// E460  F0 04     BEQ $E466                       A:00 X:55 Y:74 P:67 SP:FB
	// E466  A9 7F     LDA #$7F                        A:00 X:55 Y:74 P:67 SP:FB
	// E468  9D 00 06  STA $0600,X @ 0655 = 00         A:7F X:55 Y:74 P:65 SP:FB
	// E46B  B8        CLV                             A:7F X:55 Y:74 P:65 SP:FB
	// E46C  18        CLC                             A:7F X:55 Y:74 P:25 SP:FB
	// E46D  FE 00 06  INC $0600,X @ 0655 = 7F         A:7F X:55 Y:74 P:24 SP:FB
	// E470  F0 0D     BEQ $E47F                       A:7F X:55 Y:74 P:A4 SP:FB
	// E472  10 0B     BPL $E47F                       A:7F X:55 Y:74 P:A4 SP:FB
	// E474  70 09     BVS $E47F                       A:7F X:55 Y:74 P:A4 SP:FB
	// E476  B0 07     BCS $E47F                       A:7F X:55 Y:74 P:A4 SP:FB
	// E478  BD 00 06  LDA $0600,X @ 0655 = 80         A:7F X:55 Y:74 P:A4 SP:FB
	// E47B  C9 80     CMP #$80                        A:80 X:55 Y:74 P:A4 SP:FB
	// E47D  F0 04     BEQ $E483                       A:80 X:55 Y:74 P:27 SP:FB
	// E483  A9 00     LDA #$00                        A:80 X:55 Y:74 P:27 SP:FB
	// E485  9D 00 06  STA $0600,X @ 0655 = 80         A:00 X:55 Y:74 P:27 SP:FB
	// E488  24 01     BIT $01 = FF                    A:00 X:55 Y:74 P:27 SP:FB
	// E48A  38        SEC                             A:00 X:55 Y:74 P:E7 SP:FB
	// E48B  DE 00 06  DEC $0600,X @ 0655 = 00         A:00 X:55 Y:74 P:E7 SP:FB
	// E48E  F0 0D     BEQ $E49D                       A:00 X:55 Y:74 P:E5 SP:FB
	// E490  10 0B     BPL $E49D                       A:00 X:55 Y:74 P:E5 SP:FB
	// E492  50 09     BVC $E49D                       A:00 X:55 Y:74 P:E5 SP:FB
	// E494  90 07     BCC $E49D                       A:00 X:55 Y:74 P:E5 SP:FB
	// E496  BD 00 06  LDA $0600,X @ 0655 = FF         A:00 X:55 Y:74 P:E5 SP:FB
	// E499  C9 FF     CMP #$FF                        A:FF X:55 Y:74 P:E5 SP:FB
	// E49B  F0 04     BEQ $E4A1                       A:FF X:55 Y:74 P:67 SP:FB
	// E4A1  A9 80     LDA #$80                        A:FF X:55 Y:74 P:67 SP:FB
	// E4A3  9D 00 06  STA $0600,X @ 0655 = FF         A:80 X:55 Y:74 P:E5 SP:FB
	// E4A6  B8        CLV                             A:80 X:55 Y:74 P:E5 SP:FB
	// E4A7  18        CLC                             A:80 X:55 Y:74 P:A5 SP:FB
	// E4A8  DE 00 06  DEC $0600,X @ 0655 = 80         A:80 X:55 Y:74 P:A4 SP:FB
	// E4AB  F0 0D     BEQ $E4BA                       A:80 X:55 Y:74 P:24 SP:FB
	// E4AD  30 0B     BMI $E4BA                       A:80 X:55 Y:74 P:24 SP:FB
	// E4AF  70 09     BVS $E4BA                       A:80 X:55 Y:74 P:24 SP:FB
	// E4B1  B0 07     BCS $E4BA                       A:80 X:55 Y:74 P:24 SP:FB
	// E4B3  BD 00 06  LDA $0600,X @ 0655 = 7F         A:80 X:55 Y:74 P:24 SP:FB
	// E4B6  C9 7F     CMP #$7F                        A:7F X:55 Y:74 P:24 SP:FB
	// E4B8  F0 04     BEQ $E4BE                       A:7F X:55 Y:74 P:27 SP:FB
	// E4BE  A9 01     LDA #$01                        A:7F X:55 Y:74 P:27 SP:FB
	// E4C0  9D 00 06  STA $0600,X @ 0655 = 7F         A:01 X:55 Y:74 P:25 SP:FB
	// E4C3  DE 00 06  DEC $0600,X @ 0655 = 01         A:01 X:55 Y:74 P:25 SP:FB
	// E4C6  F0 04     BEQ $E4CC                       A:01 X:55 Y:74 P:27 SP:FB
	// E4CC  A9 33     LDA #$33                        A:01 X:55 Y:74 P:27 SP:FB
	// E4CE  8D 78 06  STA $0678 = 7F                  A:33 X:55 Y:74 P:25 SP:FB
	// E4D1  A9 44     LDA #$44                        A:33 X:55 Y:74 P:25 SP:FB
	// E4D3  A0 78     LDY #$78                        A:44 X:55 Y:74 P:25 SP:FB
	// E4D5  A2 00     LDX #$00                        A:44 X:55 Y:78 P:25 SP:FB
	// E4D7  38        SEC                             A:44 X:00 Y:78 P:27 SP:FB
	// E4D8  24 01     BIT $01 = FF                    A:44 X:00 Y:78 P:27 SP:FB
	// E4DA  BE 00 06  LDX $0600,Y @ 0678 = 33         A:44 X:00 Y:78 P:E5 SP:FB
	// E4DD  90 12     BCC $E4F1                       A:44 X:33 Y:78 P:65 SP:FB
	// E4DF  50 10     BVC $E4F1                       A:44 X:33 Y:78 P:65 SP:FB
	// E4E1  30 0E     BMI $E4F1                       A:44 X:33 Y:78 P:65 SP:FB
	// E4E3  F0 0C     BEQ $E4F1                       A:44 X:33 Y:78 P:65 SP:FB
	// E4E5  E0 33     CPX #$33                        A:44 X:33 Y:78 P:65 SP:FB
	// E4E7  D0 08     BNE $E4F1                       A:44 X:33 Y:78 P:67 SP:FB
	// E4E9  C0 78     CPY #$78                        A:44 X:33 Y:78 P:67 SP:FB
	// E4EB  D0 04     BNE $E4F1                       A:44 X:33 Y:78 P:67 SP:FB
	// E4ED  C9 44     CMP #$44                        A:44 X:33 Y:78 P:67 SP:FB
	// E4EF  F0 04     BEQ $E4F5                       A:44 X:33 Y:78 P:67 SP:FB
	// E4F5  A9 97     LDA #$97                        A:44 X:33 Y:78 P:67 SP:FB
	// E4F7  8D 7F 06  STA $067F = 00                  A:97 X:33 Y:78 P:E5 SP:FB
	// E4FA  A9 47     LDA #$47                        A:97 X:33 Y:78 P:E5 SP:FB
	// E4FC  A0 FF     LDY #$FF                        A:47 X:33 Y:78 P:65 SP:FB
	// E4FE  A2 00     LDX #$00                        A:47 X:33 Y:FF P:E5 SP:FB
	// E500  18        CLC                             A:47 X:00 Y:FF P:67 SP:FB
	// E501  B8        CLV                             A:47 X:00 Y:FF P:66 SP:FB
	// E502  BE 80 05  LDX $0580,Y @ 067F = 97         A:47 X:00 Y:FF P:26 SP:FB
	// E505  B0 12     BCS $E519                       A:47 X:97 Y:FF P:A4 SP:FB
	// E507  70 10     BVS $E519                       A:47 X:97 Y:FF P:A4 SP:FB
	// E509  10 0E     BPL $E519                       A:47 X:97 Y:FF P:A4 SP:FB
	// E50B  F0 0C     BEQ $E519                       A:47 X:97 Y:FF P:A4 SP:FB
	// E50D  E0 97     CPX #$97                        A:47 X:97 Y:FF P:A4 SP:FB
	// E50F  D0 08     BNE $E519                       A:47 X:97 Y:FF P:27 SP:FB
	// E511  C0 FF     CPY #$FF                        A:47 X:97 Y:FF P:27 SP:FB
	// E513  D0 04     BNE $E519                       A:47 X:97 Y:FF P:27 SP:FB
	// E515  C9 47     CMP #$47                        A:47 X:97 Y:FF P:27 SP:FB
	// E517  F0 04     BEQ $E51D                       A:47 X:97 Y:FF P:27 SP:FB
	// E51D  60        RTS                             A:47 X:97 Y:FF P:27 SP:FB
	// C62F  20 A3 C6  JSR $C6A3                       A:47 X:97 Y:FF P:27 SP:FD
	// C6A3  A0 4E     LDY #$4E                        A:47 X:97 Y:FF P:27 SP:FB
	// C6A5  A9 FF     LDA #$FF                        A:47 X:97 Y:4E P:25 SP:FB
	// C6A7  85 01     STA $01 = FF                    A:FF X:97 Y:4E P:A5 SP:FB
	// C6A9  20 B0 C6  JSR $C6B0                       A:FF X:97 Y:4E P:A5 SP:FB
	// C6B0  A9 FF     LDA #$FF                        A:FF X:97 Y:4E P:A5 SP:F9
	// C6B2  48        PHA                             A:FF X:97 Y:4E P:A5 SP:F9
	// C6B3  A9 AA     LDA #$AA                        A:FF X:97 Y:4E P:A5 SP:F8
	// C6B5  D0 05     BNE $C6BC                       A:AA X:97 Y:4E P:A5 SP:F8
	// C6BC  28        PLP                             A:AA X:97 Y:4E P:A5 SP:F8
	// C6BD  04 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9
	// C6BF  44 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9
	// C6C1  64 A9    *NOP $A9 = 00                    A:AA X:97 Y:4E P:EF SP:F9
	// C6C3  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9
	// C6C4  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9
	// C6C5  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9
	// C6C6  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F9
	// C6C7  08        PHP                             A:AA X:97 Y:4E P:EF SP:F9
	// C6C8  48        PHA                             A:AA X:97 Y:4E P:EF SP:F8
	// C6C9  0C A9 A9 *NOP $A9A9 = 00                  A:AA X:97 Y:4E P:EF SP:F7
	// C6CC  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7
	// C6CD  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7
	// C6CE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7
	// C6CF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F7
	// C6D0  08        PHP                             A:AA X:97 Y:4E P:EF SP:F7
	// C6D1  48        PHA                             A:AA X:97 Y:4E P:EF SP:F6
	// C6D2  14 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6D4  34 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6D6  54 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6D8  74 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6DA  D4 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6DC  F4 A9    *NOP $A9,X @ 40 = 00             A:AA X:97 Y:4E P:EF SP:F5
	// C6DE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5
	// C6DF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5
	// C6E0  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5
	// C6E1  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F5
	// C6E2  08        PHP                             A:AA X:97 Y:4E P:EF SP:F5
	// C6E3  48        PHA                             A:AA X:97 Y:4E P:EF SP:F4
	// C6E4  1A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6E5  3A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6E6  5A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6E7  7A       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6E8  DA       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6E9  FA       *NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6EA  80 89    *NOP #$89                        A:AA X:97 Y:4E P:EF SP:F3
	// C6EC  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6ED  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6EE  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6EF  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6F0  08        PHP                             A:AA X:97 Y:4E P:EF SP:F3
	// C6F1  48        PHA                             A:AA X:97 Y:4E P:EF SP:F2
	// C6F2  1C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C6F5  3C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C6F8  5C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C6FB  7C A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C6FE  DC A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C701  FC A9 A9 *NOP $A9A9,X @ AA40 = 00         A:AA X:97 Y:4E P:EF SP:F1
	// C704  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1
	// C705  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1
	// C706  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1
	// C707  EA        NOP                             A:AA X:97 Y:4E P:EF SP:F1
	// C708  08        PHP                             A:AA X:97 Y:4E P:EF SP:F1
	// C709  48        PHA                             A:AA X:97 Y:4E P:EF SP:F0
	// C70A  A2 05     LDX #$05                        A:AA X:97 Y:4E P:EF SP:EF
	// C70C  68        PLA                             A:AA X:05 Y:4E P:6D SP:EF
	// C70D  C9 55     CMP #$55                        A:AA X:05 Y:4E P:ED SP:F0
	// C70F  F0 0A     BEQ $C71B                       A:AA X:05 Y:4E P:6D SP:F0
	// C711  C9 AA     CMP #$AA                        A:AA X:05 Y:4E P:6D SP:F0
	// C713  F0 06     BEQ $C71B                       A:AA X:05 Y:4E P:6F SP:F0
	// C71B  68        PLA                             A:AA X:05 Y:4E P:6F SP:F0
	// C71C  29 CB     AND #$CB                        A:FF X:05 Y:4E P:ED SP:F1
	// C71E  C9 00     CMP #$00                        A:CB X:05 Y:4E P:ED SP:F1
	// C720  F0 06     BEQ $C728                       A:CB X:05 Y:4E P:ED SP:F1
	// C722  C9 CB     CMP #$CB                        A:CB X:05 Y:4E P:ED SP:F1
	// C724  F0 02     BEQ $C728                       A:CB X:05 Y:4E P:6F SP:F1
	// C728  C8        INY                             A:CB X:05 Y:4E P:6F SP:F1
	// C729  CA        DEX                             A:CB X:05 Y:4F P:6D SP:F1
	// C72A  D0 E0     BNE $C70C                       A:CB X:04 Y:4F P:6D SP:F1
	// C70C  68        PLA                             A:CB X:04 Y:4F P:6D SP:F1
	// C70D  C9 55     CMP #$55                        A:AA X:04 Y:4F P:ED SP:F2
	// C70F  F0 0A     BEQ $C71B                       A:AA X:04 Y:4F P:6D SP:F2
	// C711  C9 AA     CMP #$AA                        A:AA X:04 Y:4F P:6D SP:F2
	// C713  F0 06     BEQ $C71B                       A:AA X:04 Y:4F P:6F SP:F2
	// C71B  68        PLA                             A:AA X:04 Y:4F P:6F SP:F2
	// C71C  29 CB     AND #$CB                        A:FF X:04 Y:4F P:ED SP:F3
	// C71E  C9 00     CMP #$00                        A:CB X:04 Y:4F P:ED SP:F3
	// C720  F0 06     BEQ $C728                       A:CB X:04 Y:4F P:ED SP:F3
	// C722  C9 CB     CMP #$CB                        A:CB X:04 Y:4F P:ED SP:F3
	// C724  F0 02     BEQ $C728                       A:CB X:04 Y:4F P:6F SP:F3
	// C728  C8        INY                             A:CB X:04 Y:4F P:6F SP:F3
	// C729  CA        DEX                             A:CB X:04 Y:50 P:6D SP:F3
	// C72A  D0 E0     BNE $C70C                       A:CB X:03 Y:50 P:6D SP:F3
	// C70C  68        PLA                             A:CB X:03 Y:50 P:6D SP:F3
	// C70D  C9 55     CMP #$55                        A:AA X:03 Y:50 P:ED SP:F4
	// C70F  F0 0A     BEQ $C71B                       A:AA X:03 Y:50 P:6D SP:F4
	// C711  C9 AA     CMP #$AA                        A:AA X:03 Y:50 P:6D SP:F4
	// C713  F0 06     BEQ $C71B                       A:AA X:03 Y:50 P:6F SP:F4
	// C71B  68        PLA                             A:AA X:03 Y:50 P:6F SP:F4
	// C71C  29 CB     AND #$CB                        A:FF X:03 Y:50 P:ED SP:F5
	// C71E  C9 00     CMP #$00                        A:CB X:03 Y:50 P:ED SP:F5
	// C720  F0 06     BEQ $C728                       A:CB X:03 Y:50 P:ED SP:F5
	// C722  C9 CB     CMP #$CB                        A:CB X:03 Y:50 P:ED SP:F5
	// C724  F0 02     BEQ $C728                       A:CB X:03 Y:50 P:6F SP:F5
	// C728  C8        INY                             A:CB X:03 Y:50 P:6F SP:F5
	// C729  CA        DEX                             A:CB X:03 Y:51 P:6D SP:F5
	// C72A  D0 E0     BNE $C70C                       A:CB X:02 Y:51 P:6D SP:F5
	// C70C  68        PLA                             A:CB X:02 Y:51 P:6D SP:F5
	// C70D  C9 55     CMP #$55                        A:AA X:02 Y:51 P:ED SP:F6
	// C70F  F0 0A     BEQ $C71B                       A:AA X:02 Y:51 P:6D SP:F6
	// C711  C9 AA     CMP #$AA                        A:AA X:02 Y:51 P:6D SP:F6
	// C713  F0 06     BEQ $C71B                       A:AA X:02 Y:51 P:6F SP:F6
	// C71B  68        PLA                             A:AA X:02 Y:51 P:6F SP:F6
	// C71C  29 CB     AND #$CB                        A:FF X:02 Y:51 P:ED SP:F7
	// C71E  C9 00     CMP #$00                        A:CB X:02 Y:51 P:ED SP:F7
	// C720  F0 06     BEQ $C728                       A:CB X:02 Y:51 P:ED SP:F7
	// C722  C9 CB     CMP #$CB                        A:CB X:02 Y:51 P:ED SP:F7
	// C724  F0 02     BEQ $C728                       A:CB X:02 Y:51 P:6F SP:F7
	// C728  C8        INY                             A:CB X:02 Y:51 P:6F SP:F7
	// C729  CA        DEX                             A:CB X:02 Y:52 P:6D SP:F7
	// C72A  D0 E0     BNE $C70C                       A:CB X:01 Y:52 P:6D SP:F7
	// C70C  68        PLA                             A:CB X:01 Y:52 P:6D SP:F7
	// C70D  C9 55     CMP #$55                        A:AA X:01 Y:52 P:ED SP:F8
	// C70F  F0 0A     BEQ $C71B                       A:AA X:01 Y:52 P:6D SP:F8
	// C711  C9 AA     CMP #$AA                        A:AA X:01 Y:52 P:6D SP:F8
	// C713  F0 06     BEQ $C71B                       A:AA X:01 Y:52 P:6F SP:F8
	// C71B  68        PLA                             A:AA X:01 Y:52 P:6F SP:F8
	// C71C  29 CB     AND #$CB                        A:FF X:01 Y:52 P:ED SP:F9
	// C71E  C9 00     CMP #$00                        A:CB X:01 Y:52 P:ED SP:F9
	// C720  F0 06     BEQ $C728                       A:CB X:01 Y:52 P:ED SP:F9
	// C722  C9 CB     CMP #$CB                        A:CB X:01 Y:52 P:ED SP:F9
	// C724  F0 02     BEQ $C728                       A:CB X:01 Y:52 P:6F SP:F9
	// C728  C8        INY                             A:CB X:01 Y:52 P:6F SP:F9
	// C729  CA        DEX                             A:CB X:01 Y:53 P:6D SP:F9
	// C72A  D0 E0     BNE $C70C                       A:CB X:00 Y:53 P:6F SP:F9
	// C72C  60        RTS                             A:CB X:00 Y:53 P:6F SP:F9
	// C6AC  20 B7 C6  JSR $C6B7                       A:CB X:00 Y:53 P:6F SP:FB
	// C6B7  A9 34     LDA #$34                        A:CB X:00 Y:53 P:6F SP:F9
	// C6B9  48        PHA                             A:34 X:00 Y:53 P:6D SP:F9
	// C6BA  A9 55     LDA #$55                        A:34 X:00 Y:53 P:6D SP:F8
	// C6BC  28        PLP                             A:55 X:00 Y:53 P:6D SP:F8
	// C6BD  04 A9    *NOP $A9 = 00                    A:55 X:00 Y:53 P:24 SP:F9
	// C6BF  44 A9    *NOP $A9 = 00                    A:55 X:00 Y:53 P:24 SP:F9
	// C6C1  64 A9    *NOP $A9 = 00                    A:55 X:00 Y:53 P:24 SP:F9
	// C6C3  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F9
	// C6C4  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F9
	// C6C5  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F9
	// C6C6  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F9
	// C6C7  08        PHP                             A:55 X:00 Y:53 P:24 SP:F9
	// C6C8  48        PHA                             A:55 X:00 Y:53 P:24 SP:F8
	// C6C9  0C A9 A9 *NOP $A9A9 = 00                  A:55 X:00 Y:53 P:24 SP:F7
	// C6CC  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F7
	// C6CD  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F7
	// C6CE  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F7
	// C6CF  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F7
	// C6D0  08        PHP                             A:55 X:00 Y:53 P:24 SP:F7
	// C6D1  48        PHA                             A:55 X:00 Y:53 P:24 SP:F6
	// C6D2  14 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6D4  34 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6D6  54 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6D8  74 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6DA  D4 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6DC  F4 A9    *NOP $A9,X @ A9 = 00             A:55 X:00 Y:53 P:24 SP:F5
	// C6DE  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F5
	// C6DF  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F5
	// C6E0  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F5
	// C6E1  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F5
	// C6E2  08        PHP                             A:55 X:00 Y:53 P:24 SP:F5
	// C6E3  48        PHA                             A:55 X:00 Y:53 P:24 SP:F4
	// C6E4  1A       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6E5  3A       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6E6  5A       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6E7  7A       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6E8  DA       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6E9  FA       *NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6EA  80 89    *NOP #$89                        A:55 X:00 Y:53 P:24 SP:F3
	// C6EC  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6ED  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6EE  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6EF  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6F0  08        PHP                             A:55 X:00 Y:53 P:24 SP:F3
	// C6F1  48        PHA                             A:55 X:00 Y:53 P:24 SP:F2
	// C6F2  1C A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C6F5  3C A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C6F8  5C A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C6FB  7C A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C6FE  DC A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C701  FC A9 A9 *NOP $A9A9,X @ A9A9 = 00         A:55 X:00 Y:53 P:24 SP:F1
	// C704  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F1
	// C705  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F1
	// C706  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F1
	// C707  EA        NOP                             A:55 X:00 Y:53 P:24 SP:F1
	// C708  08        PHP                             A:55 X:00 Y:53 P:24 SP:F1
	// C709  48        PHA                             A:55 X:00 Y:53 P:24 SP:F0
	// C70A  A2 05     LDX #$05                        A:55 X:00 Y:53 P:24 SP:EF
	// C70C  68        PLA                             A:55 X:05 Y:53 P:24 SP:EF
	// C70D  C9 55     CMP #$55                        A:55 X:05 Y:53 P:24 SP:F0
	// C70F  F0 0A     BEQ $C71B                       A:55 X:05 Y:53 P:27 SP:F0
	// C71B  68        PLA                             A:55 X:05 Y:53 P:27 SP:F0
	// C71C  29 CB     AND #$CB                        A:34 X:05 Y:53 P:25 SP:F1
	// C71E  C9 00     CMP #$00                        A:00 X:05 Y:53 P:27 SP:F1
	// C720  F0 06     BEQ $C728                       A:00 X:05 Y:53 P:27 SP:F1
	// C728  C8        INY                             A:00 X:05 Y:53 P:27 SP:F1
	// C729  CA        DEX                             A:00 X:05 Y:54 P:25 SP:F1
	// C72A  D0 E0     BNE $C70C                       A:00 X:04 Y:54 P:25 SP:F1
	// C70C  68        PLA                             A:00 X:04 Y:54 P:25 SP:F1
	// C70D  C9 55     CMP #$55                        A:55 X:04 Y:54 P:25 SP:F2
	// C70F  F0 0A     BEQ $C71B                       A:55 X:04 Y:54 P:27 SP:F2
	// C71B  68        PLA                             A:55 X:04 Y:54 P:27 SP:F2
	// C71C  29 CB     AND #$CB                        A:34 X:04 Y:54 P:25 SP:F3
	// C71E  C9 00     CMP #$00                        A:00 X:04 Y:54 P:27 SP:F3
	// C720  F0 06     BEQ $C728                       A:00 X:04 Y:54 P:27 SP:F3
	// C728  C8        INY                             A:00 X:04 Y:54 P:27 SP:F3
	// C729  CA        DEX                             A:00 X:04 Y:55 P:25 SP:F3
	// C72A  D0 E0     BNE $C70C                       A:00 X:03 Y:55 P:25 SP:F3
	// C70C  68        PLA                             A:00 X:03 Y:55 P:25 SP:F3
	// C70D  C9 55     CMP #$55                        A:55 X:03 Y:55 P:25 SP:F4
	// C70F  F0 0A     BEQ $C71B                       A:55 X:03 Y:55 P:27 SP:F4
	// C71B  68        PLA                             A:55 X:03 Y:55 P:27 SP:F4
	// C71C  29 CB     AND #$CB                        A:34 X:03 Y:55 P:25 SP:F5
	// C71E  C9 00     CMP #$00                        A:00 X:03 Y:55 P:27 SP:F5
	// C720  F0 06     BEQ $C728                       A:00 X:03 Y:55 P:27 SP:F5
	// C728  C8        INY                             A:00 X:03 Y:55 P:27 SP:F5
	// C729  CA        DEX                             A:00 X:03 Y:56 P:25 SP:F5
	// C72A  D0 E0     BNE $C70C                       A:00 X:02 Y:56 P:25 SP:F5
	// C70C  68        PLA                             A:00 X:02 Y:56 P:25 SP:F5
	// C70D  C9 55     CMP #$55                        A:55 X:02 Y:56 P:25 SP:F6
	// C70F  F0 0A     BEQ $C71B                       A:55 X:02 Y:56 P:27 SP:F6
	// C71B  68        PLA                             A:55 X:02 Y:56 P:27 SP:F6
	// C71C  29 CB     AND #$CB                        A:34 X:02 Y:56 P:25 SP:F7
	// C71E  C9 00     CMP #$00                        A:00 X:02 Y:56 P:27 SP:F7
	// C720  F0 06     BEQ $C728                       A:00 X:02 Y:56 P:27 SP:F7
	// C728  C8        INY                             A:00 X:02 Y:56 P:27 SP:F7
	// C729  CA        DEX                             A:00 X:02 Y:57 P:25 SP:F7
	// C72A  D0 E0     BNE $C70C                       A:00 X:01 Y:57 P:25 SP:F7
	// C70C  68        PLA                             A:00 X:01 Y:57 P:25 SP:F7
	// C70D  C9 55     CMP #$55                        A:55 X:01 Y:57 P:25 SP:F8
	// C70F  F0 0A     BEQ $C71B                       A:55 X:01 Y:57 P:27 SP:F8
	// C71B  68        PLA                             A:55 X:01 Y:57 P:27 SP:F8
	// C71C  29 CB     AND #$CB                        A:34 X:01 Y:57 P:25 SP:F9
	// C71E  C9 00     CMP #$00                        A:00 X:01 Y:57 P:27 SP:F9
	// C720  F0 06     BEQ $C728                       A:00 X:01 Y:57 P:27 SP:F9
	// C728  C8        INY                             A:00 X:01 Y:57 P:27 SP:F9
	// C729  CA        DEX                             A:00 X:01 Y:58 P:25 SP:F9
	// C72A  D0 E0     BNE $C70C                       A:00 X:00 Y:58 P:27 SP:F9
	// C72C  60        RTS                             A:00 X:00 Y:58 P:27 SP:F9
	// C6AF  60        RTS                             A:00 X:00 Y:58 P:27 SP:FB
	// C632  20 1E E5  JSR $E51E                       A:00 X:00 Y:58 P:27 SP:FD
	// E51E  A9 55     LDA #$55                        A:00 X:00 Y:58 P:27 SP:FB
	// E520  8D 80 05  STA $0580 = 00                  A:55 X:00 Y:58 P:25 SP:FB
	// E523  A9 AA     LDA #$AA                        A:55 X:00 Y:58 P:25 SP:FB
	// E525  8D 32 04  STA $0432 = 00                  A:AA X:00 Y:58 P:A5 SP:FB
	// E528  A9 80     LDA #$80                        A:AA X:00 Y:58 P:A5 SP:FB
	// E52A  85 43     STA $43 = 00                    A:80 X:00 Y:58 P:A5 SP:FB
	// E52C  A9 05     LDA #$05                        A:80 X:00 Y:58 P:A5 SP:FB
	// E52E  85 44     STA $44 = 00                    A:05 X:00 Y:58 P:25 SP:FB
	// E530  A9 32     LDA #$32                        A:05 X:00 Y:58 P:25 SP:FB
	// E532  85 45     STA $45 = 00                    A:32 X:00 Y:58 P:25 SP:FB
	// E534  A9 04     LDA #$04                        A:32 X:00 Y:58 P:25 SP:FB
	// E536  85 46     STA $46 = 00                    A:04 X:00 Y:58 P:25 SP:FB
	// E538  A2 03     LDX #$03                        A:04 X:00 Y:58 P:25 SP:FB
	// E53A  A0 77     LDY #$77                        A:04 X:03 Y:58 P:25 SP:FB
	// E53C  A9 FF     LDA #$FF                        A:04 X:03 Y:77 P:25 SP:FB
	// E53E  85 01     STA $01 = FF                    A:FF X:03 Y:77 P:A5 SP:FB
	// E540  24 01     BIT $01 = FF                    A:FF X:03 Y:77 P:A5 SP:FB
	// E542  38        SEC                             A:FF X:03 Y:77 P:E5 SP:FB
	// E543  A9 00     LDA #$00                        A:FF X:03 Y:77 P:E5 SP:FB
	// E545  A3 40    *LAX ($40,X) @ 43 = 0580 = 55    A:00 X:03 Y:77 P:67 SP:FB
	// E547  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB
	// E548  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB
	// E549  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB
	// E54A  EA        NOP                             A:55 X:55 Y:77 P:65 SP:FB
	// E54B  F0 12     BEQ $E55F                       A:55 X:55 Y:77 P:65 SP:FB
	// E54D  30 10     BMI $E55F                       A:55 X:55 Y:77 P:65 SP:FB
	// E54F  50 0E     BVC $E55F                       A:55 X:55 Y:77 P:65 SP:FB
	// E551  90 0C     BCC $E55F                       A:55 X:55 Y:77 P:65 SP:FB
	// E553  C9 55     CMP #$55                        A:55 X:55 Y:77 P:65 SP:FB
	// E555  D0 08     BNE $E55F                       A:55 X:55 Y:77 P:67 SP:FB
	// E557  E0 55     CPX #$55                        A:55 X:55 Y:77 P:67 SP:FB
	// E559  D0 04     BNE $E55F                       A:55 X:55 Y:77 P:67 SP:FB
	// E55B  C0 77     CPY #$77                        A:55 X:55 Y:77 P:67 SP:FB
	// E55D  F0 04     BEQ $E563                       A:55 X:55 Y:77 P:67 SP:FB
	// E563  A2 05     LDX #$05                        A:55 X:55 Y:77 P:67 SP:FB
	// E565  A0 33     LDY #$33                        A:55 X:05 Y:77 P:65 SP:FB
	// E567  B8        CLV                             A:55 X:05 Y:33 P:65 SP:FB
	// E568  18        CLC                             A:55 X:05 Y:33 P:25 SP:FB
	// E569  A9 00     LDA #$00                        A:55 X:05 Y:33 P:24 SP:FB
	// E56B  A3 40    *LAX ($40,X) @ 45 = 0432 = AA    A:00 X:05 Y:33 P:26 SP:FB
	// E56D  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB
	// E56E  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB
	// E56F  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB
	// E570  EA        NOP                             A:AA X:AA Y:33 P:A4 SP:FB
	// E571  F0 12     BEQ $E585                       A:AA X:AA Y:33 P:A4 SP:FB
	// E573  10 10     BPL $E585                       A:AA X:AA Y:33 P:A4 SP:FB
	// E575  70 0E     BVS $E585                       A:AA X:AA Y:33 P:A4 SP:FB
	// E577  B0 0C     BCS $E585                       A:AA X:AA Y:33 P:A4 SP:FB
	// E579  C9 AA     CMP #$AA                        A:AA X:AA Y:33 P:A4 SP:FB
	// E57B  D0 08     BNE $E585                       A:AA X:AA Y:33 P:27 SP:FB
	// E57D  E0 AA     CPX #$AA                        A:AA X:AA Y:33 P:27 SP:FB
	// E57F  D0 04     BNE $E585                       A:AA X:AA Y:33 P:27 SP:FB
	// E581  C0 33     CPY #$33                        A:AA X:AA Y:33 P:27 SP:FB
	// E583  F0 04     BEQ $E589                       A:AA X:AA Y:33 P:27 SP:FB
	// E589  A9 87     LDA #$87                        A:AA X:AA Y:33 P:27 SP:FB
	// E58B  85 67     STA $67 = 00                    A:87 X:AA Y:33 P:A5 SP:FB
	// E58D  A9 32     LDA #$32                        A:87 X:AA Y:33 P:A5 SP:FB
	// E58F  85 68     STA $68 = 00                    A:32 X:AA Y:33 P:25 SP:FB
	// E591  A0 57     LDY #$57                        A:32 X:AA Y:33 P:25 SP:FB
	// E593  24 01     BIT $01 = FF                    A:32 X:AA Y:57 P:25 SP:FB
	// E595  38        SEC                             A:32 X:AA Y:57 P:E5 SP:FB
	// E596  A9 00     LDA #$00                        A:32 X:AA Y:57 P:E5 SP:FB
	// E598  A7 67    *LAX $67 = 87                    A:00 X:AA Y:57 P:67 SP:FB
	// E59A  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E59B  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E59C  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E59D  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E59E  F0 12     BEQ $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5A0  10 10     BPL $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5A2  50 0E     BVC $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5A4  90 0C     BCC $E5B2                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5A6  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB
	// E5A8  D0 08     BNE $E5B2                       A:87 X:87 Y:57 P:67 SP:FB
	// E5AA  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB
	// E5AC  D0 04     BNE $E5B2                       A:87 X:87 Y:57 P:67 SP:FB
	// E5AE  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB
	// E5B0  F0 04     BEQ $E5B6                       A:87 X:87 Y:57 P:67 SP:FB
	// E5B6  A0 53     LDY #$53                        A:87 X:87 Y:57 P:67 SP:FB
	// E5B8  B8        CLV                             A:87 X:87 Y:53 P:65 SP:FB
	// E5B9  18        CLC                             A:87 X:87 Y:53 P:25 SP:FB
	// E5BA  A9 00     LDA #$00                        A:87 X:87 Y:53 P:24 SP:FB
	// E5BC  A7 68    *LAX $68 = 32                    A:00 X:87 Y:53 P:26 SP:FB
	// E5BE  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E5BF  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E5C0  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E5C1  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E5C2  F0 12     BEQ $E5D6                       A:32 X:32 Y:53 P:24 SP:FB
	// E5C4  30 10     BMI $E5D6                       A:32 X:32 Y:53 P:24 SP:FB
	// E5C6  70 0E     BVS $E5D6                       A:32 X:32 Y:53 P:24 SP:FB
	// E5C8  B0 0C     BCS $E5D6                       A:32 X:32 Y:53 P:24 SP:FB
	// E5CA  C9 32     CMP #$32                        A:32 X:32 Y:53 P:24 SP:FB
	// E5CC  D0 08     BNE $E5D6                       A:32 X:32 Y:53 P:27 SP:FB
	// E5CE  E0 32     CPX #$32                        A:32 X:32 Y:53 P:27 SP:FB
	// E5D0  D0 04     BNE $E5D6                       A:32 X:32 Y:53 P:27 SP:FB
	// E5D2  C0 53     CPY #$53                        A:32 X:32 Y:53 P:27 SP:FB
	// E5D4  F0 04     BEQ $E5DA                       A:32 X:32 Y:53 P:27 SP:FB
	// E5DA  A9 87     LDA #$87                        A:32 X:32 Y:53 P:27 SP:FB
	// E5DC  8D 77 05  STA $0577 = 00                  A:87 X:32 Y:53 P:A5 SP:FB
	// E5DF  A9 32     LDA #$32                        A:87 X:32 Y:53 P:A5 SP:FB
	// E5E1  8D 78 05  STA $0578 = 00                  A:32 X:32 Y:53 P:25 SP:FB
	// E5E4  A0 57     LDY #$57                        A:32 X:32 Y:53 P:25 SP:FB
	// E5E6  24 01     BIT $01 = FF                    A:32 X:32 Y:57 P:25 SP:FB
	// E5E8  38        SEC                             A:32 X:32 Y:57 P:E5 SP:FB
	// E5E9  A9 00     LDA #$00                        A:32 X:32 Y:57 P:E5 SP:FB
	// E5EB  AF 77 05 *LAX $0577 = 87                  A:00 X:32 Y:57 P:67 SP:FB
	// E5EE  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E5EF  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E5F0  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E5F1  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E5F2  F0 12     BEQ $E606                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5F4  10 10     BPL $E606                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5F6  50 0E     BVC $E606                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5F8  90 0C     BCC $E606                       A:87 X:87 Y:57 P:E5 SP:FB
	// E5FA  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB
	// E5FC  D0 08     BNE $E606                       A:87 X:87 Y:57 P:67 SP:FB
	// E5FE  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB
	// E600  D0 04     BNE $E606                       A:87 X:87 Y:57 P:67 SP:FB
	// E602  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB
	// E604  F0 04     BEQ $E60A                       A:87 X:87 Y:57 P:67 SP:FB
	// E60A  A0 53     LDY #$53                        A:87 X:87 Y:57 P:67 SP:FB
	// E60C  B8        CLV                             A:87 X:87 Y:53 P:65 SP:FB
	// E60D  18        CLC                             A:87 X:87 Y:53 P:25 SP:FB
	// E60E  A9 00     LDA #$00                        A:87 X:87 Y:53 P:24 SP:FB
	// E610  AF 78 05 *LAX $0578 = 32                  A:00 X:87 Y:53 P:26 SP:FB
	// E613  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E614  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E615  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E616  EA        NOP                             A:32 X:32 Y:53 P:24 SP:FB
	// E617  F0 12     BEQ $E62B                       A:32 X:32 Y:53 P:24 SP:FB
	// E619  30 10     BMI $E62B                       A:32 X:32 Y:53 P:24 SP:FB
	// E61B  70 0E     BVS $E62B                       A:32 X:32 Y:53 P:24 SP:FB
	// E61D  B0 0C     BCS $E62B                       A:32 X:32 Y:53 P:24 SP:FB
	// E61F  C9 32     CMP #$32                        A:32 X:32 Y:53 P:24 SP:FB
	// E621  D0 08     BNE $E62B                       A:32 X:32 Y:53 P:27 SP:FB
	// E623  E0 32     CPX #$32                        A:32 X:32 Y:53 P:27 SP:FB
	// E625  D0 04     BNE $E62B                       A:32 X:32 Y:53 P:27 SP:FB
	// E627  C0 53     CPY #$53                        A:32 X:32 Y:53 P:27 SP:FB
	// E629  F0 04     BEQ $E62F                       A:32 X:32 Y:53 P:27 SP:FB
	// E62F  A9 FF     LDA #$FF                        A:32 X:32 Y:53 P:27 SP:FB
	// E631  85 43     STA $43 = 80                    A:FF X:32 Y:53 P:A5 SP:FB
	// E633  A9 04     LDA #$04                        A:FF X:32 Y:53 P:A5 SP:FB
	// E635  85 44     STA $44 = 05                    A:04 X:32 Y:53 P:25 SP:FB
	// E637  A9 32     LDA #$32                        A:04 X:32 Y:53 P:25 SP:FB
	// E639  85 45     STA $45 = 32                    A:32 X:32 Y:53 P:25 SP:FB
	// E63B  A9 04     LDA #$04                        A:32 X:32 Y:53 P:25 SP:FB
	// E63D  85 46     STA $46 = 04                    A:04 X:32 Y:53 P:25 SP:FB
	// E63F  A9 55     LDA #$55                        A:04 X:32 Y:53 P:25 SP:FB
	// E641  8D 80 05  STA $0580 = 55                  A:55 X:32 Y:53 P:25 SP:FB
	// E644  A9 AA     LDA #$AA                        A:55 X:32 Y:53 P:25 SP:FB
	// E646  8D 32 04  STA $0432 = AA                  A:AA X:32 Y:53 P:A5 SP:FB
	// E649  A2 03     LDX #$03                        A:AA X:32 Y:53 P:A5 SP:FB
	// E64B  A0 81     LDY #$81                        A:AA X:03 Y:53 P:25 SP:FB
	// E64D  24 01     BIT $01 = FF                    A:AA X:03 Y:81 P:A5 SP:FB
	// E64F  38        SEC                             A:AA X:03 Y:81 P:E5 SP:FB
	// E650  A9 00     LDA #$00                        A:AA X:03 Y:81 P:E5 SP:FB
	// E652  B3 43    *LAX ($43),Y = 04FF @ 0580 = 55  A:00 X:03 Y:81 P:67 SP:FB
	// E654  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB
	// E655  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB
	// E656  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB
	// E657  EA        NOP                             A:55 X:55 Y:81 P:65 SP:FB
	// E658  F0 12     BEQ $E66C                       A:55 X:55 Y:81 P:65 SP:FB
	// E65A  30 10     BMI $E66C                       A:55 X:55 Y:81 P:65 SP:FB
	// E65C  50 0E     BVC $E66C                       A:55 X:55 Y:81 P:65 SP:FB
	// E65E  90 0C     BCC $E66C                       A:55 X:55 Y:81 P:65 SP:FB
	// E660  C9 55     CMP #$55                        A:55 X:55 Y:81 P:65 SP:FB
	// E662  D0 08     BNE $E66C                       A:55 X:55 Y:81 P:67 SP:FB
	// E664  E0 55     CPX #$55                        A:55 X:55 Y:81 P:67 SP:FB
	// E666  D0 04     BNE $E66C                       A:55 X:55 Y:81 P:67 SP:FB
	// E668  C0 81     CPY #$81                        A:55 X:55 Y:81 P:67 SP:FB
	// E66A  F0 04     BEQ $E670                       A:55 X:55 Y:81 P:67 SP:FB
	// E670  A2 05     LDX #$05                        A:55 X:55 Y:81 P:67 SP:FB
	// E672  A0 00     LDY #$00                        A:55 X:05 Y:81 P:65 SP:FB
	// E674  B8        CLV                             A:55 X:05 Y:00 P:67 SP:FB
	// E675  18        CLC                             A:55 X:05 Y:00 P:27 SP:FB
	// E676  A9 00     LDA #$00                        A:55 X:05 Y:00 P:26 SP:FB
	// E678  B3 45    *LAX ($45),Y = 0432 @ 0432 = AA  A:00 X:05 Y:00 P:26 SP:FB
	// E67A  EA        NOP                             A:AA X:AA Y:00 P:A4 SP:FB
	// E67B  EA        NOP                             A:AA X:AA Y:00 P:A4 SP:FB
	// E67C  EA        NOP                             A:AA X:AA Y:00 P:A4 SP:FB
	// E67D  EA        NOP                             A:AA X:AA Y:00 P:A4 SP:FB
	// E67E  F0 12     BEQ $E692                       A:AA X:AA Y:00 P:A4 SP:FB
	// E680  10 10     BPL $E692                       A:AA X:AA Y:00 P:A4 SP:FB
	// E682  70 0E     BVS $E692                       A:AA X:AA Y:00 P:A4 SP:FB
	// E684  B0 0C     BCS $E692                       A:AA X:AA Y:00 P:A4 SP:FB
	// E686  C9 AA     CMP #$AA                        A:AA X:AA Y:00 P:A4 SP:FB
	// E688  D0 08     BNE $E692                       A:AA X:AA Y:00 P:27 SP:FB
	// E68A  E0 AA     CPX #$AA                        A:AA X:AA Y:00 P:27 SP:FB
	// E68C  D0 04     BNE $E692                       A:AA X:AA Y:00 P:27 SP:FB
	// E68E  C0 00     CPY #$00                        A:AA X:AA Y:00 P:27 SP:FB
	// E690  F0 04     BEQ $E696                       A:AA X:AA Y:00 P:27 SP:FB
	// E696  A9 87     LDA #$87                        A:AA X:AA Y:00 P:27 SP:FB
	// E698  85 67     STA $67 = 87                    A:87 X:AA Y:00 P:A5 SP:FB
	// E69A  A9 32     LDA #$32                        A:87 X:AA Y:00 P:A5 SP:FB
	// E69C  85 68     STA $68 = 32                    A:32 X:AA Y:00 P:25 SP:FB
	// E69E  A0 57     LDY #$57                        A:32 X:AA Y:00 P:25 SP:FB
	// E6A0  24 01     BIT $01 = FF                    A:32 X:AA Y:57 P:25 SP:FB
	// E6A2  38        SEC                             A:32 X:AA Y:57 P:E5 SP:FB
	// E6A3  A9 00     LDA #$00                        A:32 X:AA Y:57 P:E5 SP:FB
	// E6A5  B7 10    *LAX $10,Y @ 67 = 87             A:00 X:AA Y:57 P:67 SP:FB
	// E6A7  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E6A8  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E6A9  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E6AA  EA        NOP                             A:87 X:87 Y:57 P:E5 SP:FB
	// E6AB  F0 12     BEQ $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB
	// E6AD  10 10     BPL $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB
	// E6AF  50 0E     BVC $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB
	// E6B1  90 0C     BCC $E6BF                       A:87 X:87 Y:57 P:E5 SP:FB
	// E6B3  C9 87     CMP #$87                        A:87 X:87 Y:57 P:E5 SP:FB
	// E6B5  D0 08     BNE $E6BF                       A:87 X:87 Y:57 P:67 SP:FB
	// E6B7  E0 87     CPX #$87                        A:87 X:87 Y:57 P:67 SP:FB
	// E6B9  D0 04     BNE $E6BF                       A:87 X:87 Y:57 P:67 SP:FB
	// E6BB  C0 57     CPY #$57                        A:87 X:87 Y:57 P:67 SP:FB
	// E6BD  F0 04     BEQ $E6C3                       A:87 X:87 Y:57 P:67 SP:FB
	// E6C3  A0 FF     LDY #$FF                        A:87 X:87 Y:57 P:67 SP:FB
	// E6C5  B8        CLV                             A:87 X:87 Y:FF P:E5 SP:FB
	// E6C6  18        CLC                             A:87 X:87 Y:FF P:A5 SP:FB
	// E6C7  A9 00     LDA #$00                        A:87 X:87 Y:FF P:A4 SP:FB
	// E6C9  B7 69    *LAX $69,Y @ 68 = 32             A:00 X:87 Y:FF P:26 SP:FB
	// E6CB  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB
	// E6CC  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB
	// E6CD  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB
	// E6CE  EA        NOP                             A:32 X:32 Y:FF P:24 SP:FB
	// E6CF  F0 12     BEQ $E6E3                       A:32 X:32 Y:FF P:24 SP:FB
	// E6D1  30 10     BMI $E6E3                       A:32 X:32 Y:FF P:24 SP:FB
	// E6D3  70 0E     BVS $E6E3                       A:32 X:32 Y:FF P:24 SP:FB
	// E6D5  B0 0C     BCS $E6E3                       A:32 X:32 Y:FF P:24 SP:FB
	// E6D7  C9 32     CMP #$32                        A:32 X:32 Y:FF P:24 SP:FB
	// E6D9  D0 08     BNE $E6E3                       A:32 X:32 Y:FF P:27 SP:FB
	// E6DB  E0 32     CPX #$32                        A:32 X:32 Y:FF P:27 SP:FB
	// E6DD  D0 04     BNE $E6E3                       A:32 X:32 Y:FF P:27 SP:FB
	// E6DF  C0 FF     CPY #$FF                        A:32 X:32 Y:FF P:27 SP:FB
	// E6E1  F0 04     BEQ $E6E7                       A:32 X:32 Y:FF P:27 SP:FB
	// E6E7  A9 87     LDA #$87                        A:32 X:32 Y:FF P:27 SP:FB
	// E6E9  8D 87 05  STA $0587 = 00                  A:87 X:32 Y:FF P:A5 SP:FB
	// E6EC  A9 32     LDA #$32                        A:87 X:32 Y:FF P:A5 SP:FB
	// E6EE  8D 88 05  STA $0588 = 00                  A:32 X:32 Y:FF P:25 SP:FB
	// E6F1  A0 30     LDY #$30                        A:32 X:32 Y:FF P:25 SP:FB
	// E6F3  24 01     BIT $01 = FF                    A:32 X:32 Y:30 P:25 SP:FB
	// E6F5  38        SEC                             A:32 X:32 Y:30 P:E5 SP:FB
	// E6F6  A9 00     LDA #$00                        A:32 X:32 Y:30 P:E5 SP:FB
	// E6F8  BF 57 05 *LAX $0557,Y @ 0587 = 87         A:00 X:32 Y:30 P:67 SP:FB
	// E6FB  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB
	// E6FC  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB
	// E6FD  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB
	// E6FE  EA        NOP                             A:87 X:87 Y:30 P:E5 SP:FB
	// E6FF  F0 12     BEQ $E713                       A:87 X:87 Y:30 P:E5 SP:FB
	// E701  10 10     BPL $E713                       A:87 X:87 Y:30 P:E5 SP:FB
	// E703  50 0E     BVC $E713                       A:87 X:87 Y:30 P:E5 SP:FB
	// E705  90 0C     BCC $E713                       A:87 X:87 Y:30 P:E5 SP:FB
	// E707  C9 87     CMP #$87                        A:87 X:87 Y:30 P:E5 SP:FB
	// E709  D0 08     BNE $E713                       A:87 X:87 Y:30 P:67 SP:FB
	// E70B  E0 87     CPX #$87                        A:87 X:87 Y:30 P:67 SP:FB
	// E70D  D0 04     BNE $E713                       A:87 X:87 Y:30 P:67 SP:FB
	// E70F  C0 30     CPY #$30                        A:87 X:87 Y:30 P:67 SP:FB
	// E711  F0 04     BEQ $E717                       A:87 X:87 Y:30 P:67 SP:FB
	// E717  A0 40     LDY #$40                        A:87 X:87 Y:30 P:67 SP:FB
	// E719  B8        CLV                             A:87 X:87 Y:40 P:65 SP:FB
	// E71A  18        CLC                             A:87 X:87 Y:40 P:25 SP:FB
	// E71B  A9 00     LDA #$00                        A:87 X:87 Y:40 P:24 SP:FB
	// E71D  BF 48 05 *LAX $0548,Y @ 0588 = 32         A:00 X:87 Y:40 P:26 SP:FB
	// E720  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB
	// E721  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB
	// E722  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB
	// E723  EA        NOP                             A:32 X:32 Y:40 P:24 SP:FB
	// E724  F0 12     BEQ $E738                       A:32 X:32 Y:40 P:24 SP:FB
	// E726  30 10     BMI $E738                       A:32 X:32 Y:40 P:24 SP:FB
	// E728  70 0E     BVS $E738                       A:32 X:32 Y:40 P:24 SP:FB
	// E72A  B0 0C     BCS $E738                       A:32 X:32 Y:40 P:24 SP:FB
	// E72C  C9 32     CMP #$32                        A:32 X:32 Y:40 P:24 SP:FB
	// E72E  D0 08     BNE $E738                       A:32 X:32 Y:40 P:27 SP:FB
	// E730  E0 32     CPX #$32                        A:32 X:32 Y:40 P:27 SP:FB
	// E732  D0 04     BNE $E738                       A:32 X:32 Y:40 P:27 SP:FB
	// E734  C0 40     CPY #$40                        A:32 X:32 Y:40 P:27 SP:FB
	// E736  F0 04     BEQ $E73C                       A:32 X:32 Y:40 P:27 SP:FB
	// E73C  60        RTS                             A:32 X:32 Y:40 P:27 SP:FB
	// C635  20 3D E7  JSR $E73D                       A:32 X:32 Y:40 P:27 SP:FD
	// E73D  A9 C0     LDA #$C0                        A:32 X:32 Y:40 P:27 SP:FB
	// E73F  85 01     STA $01 = FF                    A:C0 X:32 Y:40 P:A5 SP:FB
	// E741  A9 00     LDA #$00                        A:C0 X:32 Y:40 P:A5 SP:FB
	// E743  8D 89 04  STA $0489 = 00                  A:00 X:32 Y:40 P:27 SP:FB
	// E746  A9 89     LDA #$89                        A:00 X:32 Y:40 P:27 SP:FB
	// E748  85 60     STA $60 = 00                    A:89 X:32 Y:40 P:A5 SP:FB
	// E74A  A9 04     LDA #$04                        A:89 X:32 Y:40 P:A5 SP:FB
	// E74C  85 61     STA $61 = 00                    A:04 X:32 Y:40 P:25 SP:FB
	// E74E  A0 44     LDY #$44                        A:04 X:32 Y:40 P:25 SP:FB
	// E750  A2 17     LDX #$17                        A:04 X:32 Y:44 P:25 SP:FB
	// E752  A9 3E     LDA #$3E                        A:04 X:17 Y:44 P:25 SP:FB
	// E754  24 01     BIT $01 = C0                    A:3E X:17 Y:44 P:25 SP:FB
	// E756  18        CLC                             A:3E X:17 Y:44 P:E7 SP:FB
	// E757  83 49    *SAX ($49,X) @ 60 = 0489 = 00    A:3E X:17 Y:44 P:E6 SP:FB
	// E759  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB
	// E75A  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB
	// E75B  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB
	// E75C  EA        NOP                             A:3E X:17 Y:44 P:E6 SP:FB
	// E75D  D0 19     BNE $E778                       A:3E X:17 Y:44 P:E6 SP:FB
	// E75F  B0 17     BCS $E778                       A:3E X:17 Y:44 P:E6 SP:FB
	// E761  50 15     BVC $E778                       A:3E X:17 Y:44 P:E6 SP:FB
	// E763  10 13     BPL $E778                       A:3E X:17 Y:44 P:E6 SP:FB
	// E765  C9 3E     CMP #$3E                        A:3E X:17 Y:44 P:E6 SP:FB
	// E767  D0 0F     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB
	// E769  C0 44     CPY #$44                        A:3E X:17 Y:44 P:67 SP:FB
	// E76B  D0 0B     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB
	// E76D  E0 17     CPX #$17                        A:3E X:17 Y:44 P:67 SP:FB
	// E76F  D0 07     BNE $E778                       A:3E X:17 Y:44 P:67 SP:FB
	// E771  AD 89 04  LDA $0489 = 16                  A:3E X:17 Y:44 P:67 SP:FB
	// E774  C9 16     CMP #$16                        A:16 X:17 Y:44 P:65 SP:FB
	// E776  F0 04     BEQ $E77C                       A:16 X:17 Y:44 P:67 SP:FB
	// E77C  A0 44     LDY #$44                        A:16 X:17 Y:44 P:67 SP:FB
	// E77E  A2 7A     LDX #$7A                        A:16 X:17 Y:44 P:65 SP:FB
	// E780  A9 66     LDA #$66                        A:16 X:7A Y:44 P:65 SP:FB
	// E782  38        SEC                             A:66 X:7A Y:44 P:65 SP:FB
	// E783  B8        CLV                             A:66 X:7A Y:44 P:65 SP:FB
	// E784  83 E6    *SAX ($E6,X) @ 60 = 0489 = 16    A:66 X:7A Y:44 P:25 SP:FB
	// E786  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB
	// E787  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB
	// E788  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB
	// E789  EA        NOP                             A:66 X:7A Y:44 P:25 SP:FB
	// E78A  F0 19     BEQ $E7A5                       A:66 X:7A Y:44 P:25 SP:FB
	// E78C  90 17     BCC $E7A5                       A:66 X:7A Y:44 P:25 SP:FB
	// E78E  70 15     BVS $E7A5                       A:66 X:7A Y:44 P:25 SP:FB
	// E790  30 13     BMI $E7A5                       A:66 X:7A Y:44 P:25 SP:FB
	// E792  C9 66     CMP #$66                        A:66 X:7A Y:44 P:25 SP:FB
	// E794  D0 0F     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB
	// E796  C0 44     CPY #$44                        A:66 X:7A Y:44 P:27 SP:FB
	// E798  D0 0B     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB
	// E79A  E0 7A     CPX #$7A                        A:66 X:7A Y:44 P:27 SP:FB
	// E79C  D0 07     BNE $E7A5                       A:66 X:7A Y:44 P:27 SP:FB
	// E79E  AD 89 04  LDA $0489 = 62                  A:66 X:7A Y:44 P:27 SP:FB
	// E7A1  C9 62     CMP #$62                        A:62 X:7A Y:44 P:25 SP:FB
	// E7A3  F0 04     BEQ $E7A9                       A:62 X:7A Y:44 P:27 SP:FB
	// E7A9  A9 FF     LDA #$FF                        A:62 X:7A Y:44 P:27 SP:FB
	// E7AB  85 49     STA $49 = 00                    A:FF X:7A Y:44 P:A5 SP:FB
	// E7AD  A0 44     LDY #$44                        A:FF X:7A Y:44 P:A5 SP:FB
	// E7AF  A2 AA     LDX #$AA                        A:FF X:7A Y:44 P:25 SP:FB
	// E7B1  A9 55     LDA #$55                        A:FF X:AA Y:44 P:A5 SP:FB
	// E7B3  24 01     BIT $01 = C0                    A:55 X:AA Y:44 P:25 SP:FB
	// E7B5  18        CLC                             A:55 X:AA Y:44 P:E5 SP:FB
	// E7B6  87 49    *SAX $49 = FF                    A:55 X:AA Y:44 P:E4 SP:FB
	// E7B8  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB
	// E7B9  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB
	// E7BA  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB
	// E7BB  EA        NOP                             A:55 X:AA Y:44 P:E4 SP:FB
	// E7BC  F0 18     BEQ $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB
	// E7BE  B0 16     BCS $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB
	// E7C0  50 14     BVC $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB
	// E7C2  10 12     BPL $E7D6                       A:55 X:AA Y:44 P:E4 SP:FB
	// E7C4  C9 55     CMP #$55                        A:55 X:AA Y:44 P:E4 SP:FB
	// E7C6  D0 0E     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB
	// E7C8  C0 44     CPY #$44                        A:55 X:AA Y:44 P:67 SP:FB
	// E7CA  D0 0A     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB
	// E7CC  E0 AA     CPX #$AA                        A:55 X:AA Y:44 P:67 SP:FB
	// E7CE  D0 06     BNE $E7D6                       A:55 X:AA Y:44 P:67 SP:FB
	// E7D0  A5 49     LDA $49 = 00                    A:55 X:AA Y:44 P:67 SP:FB
	// E7D2  C9 00     CMP #$00                        A:00 X:AA Y:44 P:67 SP:FB
	// E7D4  F0 04     BEQ $E7DA                       A:00 X:AA Y:44 P:67 SP:FB
	// E7DA  A9 00     LDA #$00                        A:00 X:AA Y:44 P:67 SP:FB
	// E7DC  85 56     STA $56 = 00                    A:00 X:AA Y:44 P:67 SP:FB
	// E7DE  A0 58     LDY #$58                        A:00 X:AA Y:44 P:67 SP:FB
	// E7E0  A2 EF     LDX #$EF                        A:00 X:AA Y:58 P:65 SP:FB
	// E7E2  A9 66     LDA #$66                        A:00 X:EF Y:58 P:E5 SP:FB
	// E7E4  38        SEC                             A:66 X:EF Y:58 P:65 SP:FB
	// E7E5  B8        CLV                             A:66 X:EF Y:58 P:65 SP:FB
	// E7E6  87 56    *SAX $56 = 00                    A:66 X:EF Y:58 P:25 SP:FB
	// E7E8  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB
	// E7E9  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB
	// E7EA  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB
	// E7EB  EA        NOP                             A:66 X:EF Y:58 P:25 SP:FB
	// E7EC  F0 18     BEQ $E806                       A:66 X:EF Y:58 P:25 SP:FB
	// E7EE  90 16     BCC $E806                       A:66 X:EF Y:58 P:25 SP:FB
	// E7F0  70 14     BVS $E806                       A:66 X:EF Y:58 P:25 SP:FB
	// E7F2  30 12     BMI $E806                       A:66 X:EF Y:58 P:25 SP:FB
	// E7F4  C9 66     CMP #$66                        A:66 X:EF Y:58 P:25 SP:FB
	// E7F6  D0 0E     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB
	// E7F8  C0 58     CPY #$58                        A:66 X:EF Y:58 P:27 SP:FB
	// E7FA  D0 0A     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB
	// E7FC  E0 EF     CPX #$EF                        A:66 X:EF Y:58 P:27 SP:FB
	// E7FE  D0 06     BNE $E806                       A:66 X:EF Y:58 P:27 SP:FB
	// E800  A5 56     LDA $56 = 66                    A:66 X:EF Y:58 P:27 SP:FB
	// E802  C9 66     CMP #$66                        A:66 X:EF Y:58 P:25 SP:FB
	// E804  F0 04     BEQ $E80A                       A:66 X:EF Y:58 P:27 SP:FB
	// E80A  A9 FF     LDA #$FF                        A:66 X:EF Y:58 P:27 SP:FB
	// E80C  8D 49 05  STA $0549 = 00                  A:FF X:EF Y:58 P:A5 SP:FB
	// E80F  A0 E5     LDY #$E5                        A:FF X:EF Y:58 P:A5 SP:FB
	// E811  A2 AF     LDX #$AF                        A:FF X:EF Y:E5 P:A5 SP:FB
	// E813  A9 F5     LDA #$F5                        A:FF X:AF Y:E5 P:A5 SP:FB
	// E815  24 01     BIT $01 = C0                    A:F5 X:AF Y:E5 P:A5 SP:FB
	// E817  18        CLC                             A:F5 X:AF Y:E5 P:E5 SP:FB
	// E818  8F 49 05 *SAX $0549 = FF                  A:F5 X:AF Y:E5 P:E4 SP:FB
	// E81B  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB
	// E81C  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB
	// E81D  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB
	// E81E  EA        NOP                             A:F5 X:AF Y:E5 P:E4 SP:FB
	// E81F  F0 19     BEQ $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB
	// E821  B0 17     BCS $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB
	// E823  50 15     BVC $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB
	// E825  10 13     BPL $E83A                       A:F5 X:AF Y:E5 P:E4 SP:FB
	// E827  C9 F5     CMP #$F5                        A:F5 X:AF Y:E5 P:E4 SP:FB
	// E829  D0 0F     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB
	// E82B  C0 E5     CPY #$E5                        A:F5 X:AF Y:E5 P:67 SP:FB
	// E82D  D0 0B     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB
	// E82F  E0 AF     CPX #$AF                        A:F5 X:AF Y:E5 P:67 SP:FB
	// E831  D0 07     BNE $E83A                       A:F5 X:AF Y:E5 P:67 SP:FB
	// E833  AD 49 05  LDA $0549 = A5                  A:F5 X:AF Y:E5 P:67 SP:FB
	// E836  C9 A5     CMP #$A5                        A:A5 X:AF Y:E5 P:E5 SP:FB
	// E838  F0 04     BEQ $E83E                       A:A5 X:AF Y:E5 P:67 SP:FB
	// E83E  A9 00     LDA #$00                        A:A5 X:AF Y:E5 P:67 SP:FB
	// E840  8D 56 05  STA $0556 = 00                  A:00 X:AF Y:E5 P:67 SP:FB
	// E843  A0 58     LDY #$58                        A:00 X:AF Y:E5 P:67 SP:FB
	// E845  A2 B3     LDX #$B3                        A:00 X:AF Y:58 P:65 SP:FB
	// E847  A9 97     LDA #$97                        A:00 X:B3 Y:58 P:E5 SP:FB
	// E849  38        SEC                             A:97 X:B3 Y:58 P:E5 SP:FB
	// E84A  B8        CLV                             A:97 X:B3 Y:58 P:E5 SP:FB
	// E84B  8F 56 05 *SAX $0556 = 00                  A:97 X:B3 Y:58 P:A5 SP:FB
	// E84E  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB
	// E84F  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB
	// E850  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB
	// E851  EA        NOP                             A:97 X:B3 Y:58 P:A5 SP:FB
	// E852  F0 19     BEQ $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB
	// E854  90 17     BCC $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB
	// E856  70 15     BVS $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB
	// E858  10 13     BPL $E86D                       A:97 X:B3 Y:58 P:A5 SP:FB
	// E85A  C9 97     CMP #$97                        A:97 X:B3 Y:58 P:A5 SP:FB
	// E85C  D0 0F     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB
	// E85E  C0 58     CPY #$58                        A:97 X:B3 Y:58 P:27 SP:FB
	// E860  D0 0B     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB
	// E862  E0 B3     CPX #$B3                        A:97 X:B3 Y:58 P:27 SP:FB
	// E864  D0 07     BNE $E86D                       A:97 X:B3 Y:58 P:27 SP:FB
	// E866  AD 56 05  LDA $0556 = 93                  A:97 X:B3 Y:58 P:27 SP:FB
	// E869  C9 93     CMP #$93                        A:93 X:B3 Y:58 P:A5 SP:FB
	// E86B  F0 04     BEQ $E871                       A:93 X:B3 Y:58 P:27 SP:FB
	// E871  A9 FF     LDA #$FF                        A:93 X:B3 Y:58 P:27 SP:FB
	// E873  85 49     STA $49 = 00                    A:FF X:B3 Y:58 P:A5 SP:FB
	// E875  A0 FF     LDY #$FF                        A:FF X:B3 Y:58 P:A5 SP:FB
	// E877  A2 AA     LDX #$AA                        A:FF X:B3 Y:FF P:A5 SP:FB
	// E879  A9 55     LDA #$55                        A:FF X:AA Y:FF P:A5 SP:FB
	// E87B  24 01     BIT $01 = C0                    A:55 X:AA Y:FF P:25 SP:FB
	// E87D  18        CLC                             A:55 X:AA Y:FF P:E5 SP:FB
	// E87E  97 4A    *SAX $4A,Y @ 49 = FF             A:55 X:AA Y:FF P:E4 SP:FB
	// E880  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB
	// E881  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB
	// E882  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB
	// E883  EA        NOP                             A:55 X:AA Y:FF P:E4 SP:FB
	// E884  F0 18     BEQ $E89E                       A:55 X:AA Y:FF P:E4 SP:FB
	// E886  B0 16     BCS $E89E                       A:55 X:AA Y:FF P:E4 SP:FB
	// E888  50 14     BVC $E89E                       A:55 X:AA Y:FF P:E4 SP:FB
	// E88A  10 12     BPL $E89E                       A:55 X:AA Y:FF P:E4 SP:FB
	// E88C  C9 55     CMP #$55                        A:55 X:AA Y:FF P:E4 SP:FB
	// E88E  D0 0E     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB
	// E890  C0 FF     CPY #$FF                        A:55 X:AA Y:FF P:67 SP:FB
	// E892  D0 0A     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB
	// E894  E0 AA     CPX #$AA                        A:55 X:AA Y:FF P:67 SP:FB
	// E896  D0 06     BNE $E89E                       A:55 X:AA Y:FF P:67 SP:FB
	// E898  A5 49     LDA $49 = 00                    A:55 X:AA Y:FF P:67 SP:FB
	// E89A  C9 00     CMP #$00                        A:00 X:AA Y:FF P:67 SP:FB
	// E89C  F0 04     BEQ $E8A2                       A:00 X:AA Y:FF P:67 SP:FB
	// E8A2  A9 00     LDA #$00                        A:00 X:AA Y:FF P:67 SP:FB
	// E8A4  85 56     STA $56 = 66                    A:00 X:AA Y:FF P:67 SP:FB
	// E8A6  A0 06     LDY #$06                        A:00 X:AA Y:FF P:67 SP:FB
	// E8A8  A2 EF     LDX #$EF                        A:00 X:AA Y:06 P:65 SP:FB
	// E8AA  A9 66     LDA #$66                        A:00 X:EF Y:06 P:E5 SP:FB
	// E8AC  38        SEC                             A:66 X:EF Y:06 P:65 SP:FB
	// E8AD  B8        CLV                             A:66 X:EF Y:06 P:65 SP:FB
	// E8AE  97 50    *SAX $50,Y @ 56 = 00             A:66 X:EF Y:06 P:25 SP:FB
	// E8B0  EA        NOP                             A:66 X:EF Y:06 P:25 SP:FB
	// E8B1  EA        NOP                             A:66 X:EF Y:06 P:25 SP:FB
	// E8B2  EA        NOP                             A:66 X:EF Y:06 P:25 SP:FB
	// E8B3  EA        NOP                             A:66 X:EF Y:06 P:25 SP:FB
	// E8B4  F0 18     BEQ $E8CE                       A:66 X:EF Y:06 P:25 SP:FB
	// E8B6  90 16     BCC $E8CE                       A:66 X:EF Y:06 P:25 SP:FB
	// E8B8  70 14     BVS $E8CE                       A:66 X:EF Y:06 P:25 SP:FB
	// E8BA  30 12     BMI $E8CE                       A:66 X:EF Y:06 P:25 SP:FB
	// E8BC  C9 66     CMP #$66                        A:66 X:EF Y:06 P:25 SP:FB
	// E8BE  D0 0E     BNE $E8CE                       A:66 X:EF Y:06 P:27 SP:FB
	// E8C0  C0 06     CPY #$06                        A:66 X:EF Y:06 P:27 SP:FB
	// E8C2  D0 0A     BNE $E8CE                       A:66 X:EF Y:06 P:27 SP:FB
	// E8C4  E0 EF     CPX #$EF                        A:66 X:EF Y:06 P:27 SP:FB
	// E8C6  D0 06     BNE $E8CE                       A:66 X:EF Y:06 P:27 SP:FB
	// E8C8  A5 56     LDA $56 = 66                    A:66 X:EF Y:06 P:27 SP:FB
	// E8CA  C9 66     CMP #$66                        A:66 X:EF Y:06 P:25 SP:FB
	// E8CC  F0 04     BEQ $E8D2                       A:66 X:EF Y:06 P:27 SP:FB
	// E8D2  60        RTS                             A:66 X:EF Y:06 P:27 SP:FB
	// C638  20 D3 E8  JSR $E8D3                       A:66 X:EF Y:06 P:27 SP:FD
	// E8D3  A0 90     LDY #$90                        A:66 X:EF Y:06 P:27 SP:FB
	// E8D5  20 31 F9  JSR $F931                       A:66 X:EF Y:90 P:A5 SP:FB
	// F931  24 01     BIT $01 = C0                    A:66 X:EF Y:90 P:A5 SP:F9
	// F933  A9 40     LDA #$40                        A:66 X:EF Y:90 P:E5 SP:F9
	// F935  38        SEC                             A:40 X:EF Y:90 P:65 SP:F9
	// F936  60        RTS                             A:40 X:EF Y:90 P:65 SP:F9
	// E8D8  EB 40    *SBC #$40                        A:40 X:EF Y:90 P:65 SP:FB
	// E8DA  EA        NOP                             A:00 X:EF Y:90 P:27 SP:FB
	// E8DB  EA        NOP                             A:00 X:EF Y:90 P:27 SP:FB
	// E8DC  EA        NOP                             A:00 X:EF Y:90 P:27 SP:FB
	// E8DD  EA        NOP                             A:00 X:EF Y:90 P:27 SP:FB
	// E8DE  20 37 F9  JSR $F937                       A:00 X:EF Y:90 P:27 SP:FB
	// F937  30 0B     BMI $F944                       A:00 X:EF Y:90 P:27 SP:F9
	// F939  90 09     BCC $F944                       A:00 X:EF Y:90 P:27 SP:F9
	// F93B  D0 07     BNE $F944                       A:00 X:EF Y:90 P:27 SP:F9
	// F93D  70 05     BVS $F944                       A:00 X:EF Y:90 P:27 SP:F9
	// F93F  C9 00     CMP #$00                        A:00 X:EF Y:90 P:27 SP:F9
	// F941  D0 01     BNE $F944                       A:00 X:EF Y:90 P:27 SP:F9
	// F943  60        RTS                             A:00 X:EF Y:90 P:27 SP:F9
	// E8E1  C8        INY                             A:00 X:EF Y:90 P:27 SP:FB
	// E8E2  20 47 F9  JSR $F947                       A:00 X:EF Y:91 P:A5 SP:FB
	// F947  B8        CLV                             A:00 X:EF Y:91 P:A5 SP:F9
	// F948  38        SEC                             A:00 X:EF Y:91 P:A5 SP:F9
	// F949  A9 40     LDA #$40                        A:00 X:EF Y:91 P:A5 SP:F9
	// F94B  60        RTS                             A:40 X:EF Y:91 P:25 SP:F9
	// E8E5  EB 3F    *SBC #$3F                        A:40 X:EF Y:91 P:25 SP:FB
	// E8E7  EA        NOP                             A:01 X:EF Y:91 P:25 SP:FB
	// E8E8  EA        NOP                             A:01 X:EF Y:91 P:25 SP:FB
	// E8E9  EA        NOP                             A:01 X:EF Y:91 P:25 SP:FB
	// E8EA  EA        NOP                             A:01 X:EF Y:91 P:25 SP:FB
	// E8EB  20 4C F9  JSR $F94C                       A:01 X:EF Y:91 P:25 SP:FB
	// F94C  F0 0B     BEQ $F959                       A:01 X:EF Y:91 P:25 SP:F9
	// F94E  30 09     BMI $F959                       A:01 X:EF Y:91 P:25 SP:F9
	// F950  90 07     BCC $F959                       A:01 X:EF Y:91 P:25 SP:F9
	// F952  70 05     BVS $F959                       A:01 X:EF Y:91 P:25 SP:F9
	// F954  C9 01     CMP #$01                        A:01 X:EF Y:91 P:25 SP:F9
	// F956  D0 01     BNE $F959                       A:01 X:EF Y:91 P:27 SP:F9
	// F958  60        RTS                             A:01 X:EF Y:91 P:27 SP:F9
	// E8EE  C8        INY                             A:01 X:EF Y:91 P:27 SP:FB
	// E8EF  20 5C F9  JSR $F95C                       A:01 X:EF Y:92 P:A5 SP:FB
	// F95C  A9 40     LDA #$40                        A:01 X:EF Y:92 P:A5 SP:F9
	// F95E  38        SEC                             A:40 X:EF Y:92 P:25 SP:F9
	// F95F  24 01     BIT $01 = C0                    A:40 X:EF Y:92 P:25 SP:F9
	// F961  60        RTS                             A:40 X:EF Y:92 P:E5 SP:F9
	// E8F2  EB 41    *SBC #$41                        A:40 X:EF Y:92 P:E5 SP:FB
	// E8F4  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB
	// E8F5  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB
	// E8F6  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB
	// E8F7  EA        NOP                             A:FF X:EF Y:92 P:A4 SP:FB
	// E8F8  20 62 F9  JSR $F962                       A:FF X:EF Y:92 P:A4 SP:FB
	// F962  B0 0B     BCS $F96F                       A:FF X:EF Y:92 P:A4 SP:F9
	// F964  F0 09     BEQ $F96F                       A:FF X:EF Y:92 P:A4 SP:F9
	// F966  10 07     BPL $F96F                       A:FF X:EF Y:92 P:A4 SP:F9
	// F968  70 05     BVS $F96F                       A:FF X:EF Y:92 P:A4 SP:F9
	// F96A  C9 FF     CMP #$FF                        A:FF X:EF Y:92 P:A4 SP:F9
	// F96C  D0 01     BNE $F96F                       A:FF X:EF Y:92 P:27 SP:F9
	// F96E  60        RTS                             A:FF X:EF Y:92 P:27 SP:F9
	// E8FB  C8        INY                             A:FF X:EF Y:92 P:27 SP:FB
	// E8FC  20 72 F9  JSR $F972                       A:FF X:EF Y:93 P:A5 SP:FB
	// F972  18        CLC                             A:FF X:EF Y:93 P:A5 SP:F9
	// F973  A9 80     LDA #$80                        A:FF X:EF Y:93 P:A4 SP:F9
	// F975  60        RTS                             A:80 X:EF Y:93 P:A4 SP:F9
	// E8FF  EB 00    *SBC #$00                        A:80 X:EF Y:93 P:A4 SP:FB
	// E901  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB
	// E902  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB
	// E903  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB
	// E904  EA        NOP                             A:7F X:EF Y:93 P:65 SP:FB
	// E905  20 76 F9  JSR $F976                       A:7F X:EF Y:93 P:65 SP:FB
	// F976  90 05     BCC $F97D                       A:7F X:EF Y:93 P:65 SP:F9
	// F978  C9 7F     CMP #$7F                        A:7F X:EF Y:93 P:65 SP:F9
	// F97A  D0 01     BNE $F97D                       A:7F X:EF Y:93 P:67 SP:F9
	// F97C  60        RTS                             A:7F X:EF Y:93 P:67 SP:F9
	// E908  C8        INY                             A:7F X:EF Y:93 P:67 SP:FB
	// E909  20 80 F9  JSR $F980                       A:7F X:EF Y:94 P:E5 SP:FB
	// F980  38        SEC                             A:7F X:EF Y:94 P:E5 SP:F9
	// F981  A9 81     LDA #$81                        A:7F X:EF Y:94 P:E5 SP:F9
	// F983  60        RTS                             A:81 X:EF Y:94 P:E5 SP:F9
	// E90C  EB 7F    *SBC #$7F                        A:81 X:EF Y:94 P:E5 SP:FB
	// E90E  EA        NOP                             A:02 X:EF Y:94 P:65 SP:FB
	// E90F  EA        NOP                             A:02 X:EF Y:94 P:65 SP:FB
	// E910  EA        NOP                             A:02 X:EF Y:94 P:65 SP:FB
	// E911  EA        NOP                             A:02 X:EF Y:94 P:65 SP:FB
	// E912  20 84 F9  JSR $F984                       A:02 X:EF Y:94 P:65 SP:FB
	// F984  50 07     BVC $F98D                       A:02 X:EF Y:94 P:65 SP:F9
	// F986  90 05     BCC $F98D                       A:02 X:EF Y:94 P:65 SP:F9
	// F988  C9 02     CMP #$02                        A:02 X:EF Y:94 P:65 SP:F9
	// F98A  D0 01     BNE $F98D                       A:02 X:EF Y:94 P:67 SP:F9
	// F98C  60        RTS                             A:02 X:EF Y:94 P:67 SP:F9
	// E915  60        RTS                             A:02 X:EF Y:94 P:67 SP:FB
	// C63B  20 16 E9  JSR $E916                       A:02 X:EF Y:94 P:67 SP:FD
	// E916  A9 FF     LDA #$FF                        A:02 X:EF Y:94 P:67 SP:FB
	// E918  85 01     STA $01 = C0                    A:FF X:EF Y:94 P:E5 SP:FB
	// E91A  A0 95     LDY #$95                        A:FF X:EF Y:94 P:E5 SP:FB
	// E91C  A2 02     LDX #$02                        A:FF X:EF Y:95 P:E5 SP:FB
	// E91E  A9 47     LDA #$47                        A:FF X:02 Y:95 P:65 SP:FB
	// E920  85 47     STA $47 = 00                    A:47 X:02 Y:95 P:65 SP:FB
	// E922  A9 06     LDA #$06                        A:47 X:02 Y:95 P:65 SP:FB
	// E924  85 48     STA $48 = 00                    A:06 X:02 Y:95 P:65 SP:FB
	// E926  A9 EB     LDA #$EB                        A:06 X:02 Y:95 P:65 SP:FB
	// E928  8D 47 06  STA $0647 = 00                  A:EB X:02 Y:95 P:E5 SP:FB
	// E92B  20 31 FA  JSR $FA31                       A:EB X:02 Y:95 P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:02 Y:95 P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:02 Y:95 P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:02 Y:95 P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:02 Y:95 P:64 SP:F9
	// E92E  C3 45    *DCP ($45,X) @ 47 = 0647 = EB    A:40 X:02 Y:95 P:64 SP:FB
	// E930  EA        NOP                             A:40 X:02 Y:95 P:64 SP:FB
	// E931  EA        NOP                             A:40 X:02 Y:95 P:64 SP:FB
	// E932  EA        NOP                             A:40 X:02 Y:95 P:64 SP:FB
	// E933  EA        NOP                             A:40 X:02 Y:95 P:64 SP:FB
	// E934  20 37 FA  JSR $FA37                       A:40 X:02 Y:95 P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:02 Y:95 P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:02 Y:95 P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:02 Y:95 P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:02 Y:95 P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:02 Y:95 P:67 SP:F9
	// FA41  60        RTS                             A:40 X:02 Y:95 P:67 SP:F9
	// E937  AD 47 06  LDA $0647 = EA                  A:40 X:02 Y:95 P:67 SP:FB
	// E93A  C9 EA     CMP #$EA                        A:EA X:02 Y:95 P:E5 SP:FB
	// E93C  F0 02     BEQ $E940                       A:EA X:02 Y:95 P:67 SP:FB
	// E940  C8        INY                             A:EA X:02 Y:95 P:67 SP:FB
	// E941  A9 00     LDA #$00                        A:EA X:02 Y:96 P:E5 SP:FB
	// E943  8D 47 06  STA $0647 = EA                  A:00 X:02 Y:96 P:67 SP:FB
	// E946  20 42 FA  JSR $FA42                       A:00 X:02 Y:96 P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:02 Y:96 P:67 SP:F9
	// FA43  38        SEC                             A:00 X:02 Y:96 P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:02 Y:96 P:27 SP:F9
	// FA46  60        RTS                             A:FF X:02 Y:96 P:A5 SP:F9
	// E949  C3 45    *DCP ($45,X) @ 47 = 0647 = 00    A:FF X:02 Y:96 P:A5 SP:FB
	// E94B  EA        NOP                             A:FF X:02 Y:96 P:27 SP:FB
	// E94C  EA        NOP                             A:FF X:02 Y:96 P:27 SP:FB
	// E94D  EA        NOP                             A:FF X:02 Y:96 P:27 SP:FB
	// E94E  EA        NOP                             A:FF X:02 Y:96 P:27 SP:FB
	// E94F  20 47 FA  JSR $FA47                       A:FF X:02 Y:96 P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:02 Y:96 P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:02 Y:96 P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:02 Y:96 P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:02 Y:96 P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:02 Y:96 P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:02 Y:96 P:27 SP:F9
	// FA53  60        RTS                             A:FF X:02 Y:96 P:27 SP:F9
	// E952  AD 47 06  LDA $0647 = FF                  A:FF X:02 Y:96 P:27 SP:FB
	// E955  C9 FF     CMP #$FF                        A:FF X:02 Y:96 P:A5 SP:FB
	// E957  F0 02     BEQ $E95B                       A:FF X:02 Y:96 P:27 SP:FB
	// E95B  C8        INY                             A:FF X:02 Y:96 P:27 SP:FB
	// E95C  A9 37     LDA #$37                        A:FF X:02 Y:97 P:A5 SP:FB
	// E95E  8D 47 06  STA $0647 = FF                  A:37 X:02 Y:97 P:25 SP:FB
	// E961  20 54 FA  JSR $FA54                       A:37 X:02 Y:97 P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:02 Y:97 P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:02 Y:97 P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:02 Y:97 P:E5 SP:F9
	// E964  C3 45    *DCP ($45,X) @ 47 = 0647 = 37    A:F0 X:02 Y:97 P:E5 SP:FB
	// E966  EA        NOP                             A:F0 X:02 Y:97 P:E5 SP:FB
	// E967  EA        NOP                             A:F0 X:02 Y:97 P:E5 SP:FB
	// E968  EA        NOP                             A:F0 X:02 Y:97 P:E5 SP:FB
	// E969  EA        NOP                             A:F0 X:02 Y:97 P:E5 SP:FB
	// E96A  20 59 FA  JSR $FA59                       A:F0 X:02 Y:97 P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:02 Y:97 P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:02 Y:97 P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:02 Y:97 P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:02 Y:97 P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:02 Y:97 P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:02 Y:97 P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:02 Y:97 P:67 SP:F9
	// E96D  AD 47 06  LDA $0647 = 36                  A:F0 X:02 Y:97 P:67 SP:FB
	// E970  C9 36     CMP #$36                        A:36 X:02 Y:97 P:65 SP:FB
	// E972  F0 02     BEQ $E976                       A:36 X:02 Y:97 P:67 SP:FB
	// E976  C8        INY                             A:36 X:02 Y:97 P:67 SP:FB
	// E977  A9 EB     LDA #$EB                        A:36 X:02 Y:98 P:E5 SP:FB
	// E979  85 47     STA $47 = 47                    A:EB X:02 Y:98 P:E5 SP:FB
	// E97B  20 31 FA  JSR $FA31                       A:EB X:02 Y:98 P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:02 Y:98 P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:02 Y:98 P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:02 Y:98 P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:02 Y:98 P:64 SP:F9
	// E97E  C7 47    *DCP $47 = EB                    A:40 X:02 Y:98 P:64 SP:FB
	// E980  EA        NOP                             A:40 X:02 Y:98 P:64 SP:FB
	// E981  EA        NOP                             A:40 X:02 Y:98 P:64 SP:FB
	// E982  EA        NOP                             A:40 X:02 Y:98 P:64 SP:FB
	// E983  EA        NOP                             A:40 X:02 Y:98 P:64 SP:FB
	// E984  20 37 FA  JSR $FA37                       A:40 X:02 Y:98 P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:02 Y:98 P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:02 Y:98 P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:02 Y:98 P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:02 Y:98 P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:02 Y:98 P:67 SP:F9
	// FA41  60        RTS                             A:40 X:02 Y:98 P:67 SP:F9
	// E987  A5 47     LDA $47 = EA                    A:40 X:02 Y:98 P:67 SP:FB
	// E989  C9 EA     CMP #$EA                        A:EA X:02 Y:98 P:E5 SP:FB
	// E98B  F0 02     BEQ $E98F                       A:EA X:02 Y:98 P:67 SP:FB
	// E98F  C8        INY                             A:EA X:02 Y:98 P:67 SP:FB
	// E990  A9 00     LDA #$00                        A:EA X:02 Y:99 P:E5 SP:FB
	// E992  85 47     STA $47 = EA                    A:00 X:02 Y:99 P:67 SP:FB
	// E994  20 42 FA  JSR $FA42                       A:00 X:02 Y:99 P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:02 Y:99 P:67 SP:F9
	// FA43  38        SEC                             A:00 X:02 Y:99 P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:02 Y:99 P:27 SP:F9
	// FA46  60        RTS                             A:FF X:02 Y:99 P:A5 SP:F9
	// E997  C7 47    *DCP $47 = 00                    A:FF X:02 Y:99 P:A5 SP:FB
	// E999  EA        NOP                             A:FF X:02 Y:99 P:27 SP:FB
	// E99A  EA        NOP                             A:FF X:02 Y:99 P:27 SP:FB
	// E99B  EA        NOP                             A:FF X:02 Y:99 P:27 SP:FB
	// E99C  EA        NOP                             A:FF X:02 Y:99 P:27 SP:FB
	// E99D  20 47 FA  JSR $FA47                       A:FF X:02 Y:99 P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:02 Y:99 P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:02 Y:99 P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:02 Y:99 P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:02 Y:99 P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:02 Y:99 P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:02 Y:99 P:27 SP:F9
	// FA53  60        RTS                             A:FF X:02 Y:99 P:27 SP:F9
	// E9A0  A5 47     LDA $47 = FF                    A:FF X:02 Y:99 P:27 SP:FB
	// E9A2  C9 FF     CMP #$FF                        A:FF X:02 Y:99 P:A5 SP:FB
	// E9A4  F0 02     BEQ $E9A8                       A:FF X:02 Y:99 P:27 SP:FB
	// E9A8  C8        INY                             A:FF X:02 Y:99 P:27 SP:FB
	// E9A9  A9 37     LDA #$37                        A:FF X:02 Y:9A P:A5 SP:FB
	// E9AB  85 47     STA $47 = FF                    A:37 X:02 Y:9A P:25 SP:FB
	// E9AD  20 54 FA  JSR $FA54                       A:37 X:02 Y:9A P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:02 Y:9A P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:02 Y:9A P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:02 Y:9A P:E5 SP:F9
	// E9B0  C7 47    *DCP $47 = 37                    A:F0 X:02 Y:9A P:E5 SP:FB
	// E9B2  EA        NOP                             A:F0 X:02 Y:9A P:E5 SP:FB
	// E9B3  EA        NOP                             A:F0 X:02 Y:9A P:E5 SP:FB
	// E9B4  EA        NOP                             A:F0 X:02 Y:9A P:E5 SP:FB
	// E9B5  EA        NOP                             A:F0 X:02 Y:9A P:E5 SP:FB
	// E9B6  20 59 FA  JSR $FA59                       A:F0 X:02 Y:9A P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:02 Y:9A P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:02 Y:9A P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:02 Y:9A P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:02 Y:9A P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:02 Y:9A P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:02 Y:9A P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:02 Y:9A P:67 SP:F9
	// E9B9  A5 47     LDA $47 = 36                    A:F0 X:02 Y:9A P:67 SP:FB
	// E9BB  C9 36     CMP #$36                        A:36 X:02 Y:9A P:65 SP:FB
	// E9BD  F0 02     BEQ $E9C1                       A:36 X:02 Y:9A P:67 SP:FB
	// E9C1  C8        INY                             A:36 X:02 Y:9A P:67 SP:FB
	// E9C2  A9 EB     LDA #$EB                        A:36 X:02 Y:9B P:E5 SP:FB
	// E9C4  8D 47 06  STA $0647 = 36                  A:EB X:02 Y:9B P:E5 SP:FB
	// E9C7  20 31 FA  JSR $FA31                       A:EB X:02 Y:9B P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:02 Y:9B P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:02 Y:9B P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:02 Y:9B P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:02 Y:9B P:64 SP:F9
	// E9CA  CF 47 06 *DCP $0647 = EB                  A:40 X:02 Y:9B P:64 SP:FB
	// E9CD  EA        NOP                             A:40 X:02 Y:9B P:64 SP:FB
	// E9CE  EA        NOP                             A:40 X:02 Y:9B P:64 SP:FB
	// E9CF  EA        NOP                             A:40 X:02 Y:9B P:64 SP:FB
	// E9D0  EA        NOP                             A:40 X:02 Y:9B P:64 SP:FB
	// E9D1  20 37 FA  JSR $FA37                       A:40 X:02 Y:9B P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:02 Y:9B P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:02 Y:9B P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:02 Y:9B P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:02 Y:9B P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:02 Y:9B P:67 SP:F9
	// FA41  60        RTS                             A:40 X:02 Y:9B P:67 SP:F9
	// E9D4  AD 47 06  LDA $0647 = EA                  A:40 X:02 Y:9B P:67 SP:FB
	// E9D7  C9 EA     CMP #$EA                        A:EA X:02 Y:9B P:E5 SP:FB
	// E9D9  F0 02     BEQ $E9DD                       A:EA X:02 Y:9B P:67 SP:FB
	// E9DD  C8        INY                             A:EA X:02 Y:9B P:67 SP:FB
	// E9DE  A9 00     LDA #$00                        A:EA X:02 Y:9C P:E5 SP:FB
	// E9E0  8D 47 06  STA $0647 = EA                  A:00 X:02 Y:9C P:67 SP:FB
	// E9E3  20 42 FA  JSR $FA42                       A:00 X:02 Y:9C P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:02 Y:9C P:67 SP:F9
	// FA43  38        SEC                             A:00 X:02 Y:9C P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:02 Y:9C P:27 SP:F9
	// FA46  60        RTS                             A:FF X:02 Y:9C P:A5 SP:F9
	// E9E6  CF 47 06 *DCP $0647 = 00                  A:FF X:02 Y:9C P:A5 SP:FB
	// E9E9  EA        NOP                             A:FF X:02 Y:9C P:27 SP:FB
	// E9EA  EA        NOP                             A:FF X:02 Y:9C P:27 SP:FB
	// E9EB  EA        NOP                             A:FF X:02 Y:9C P:27 SP:FB
	// E9EC  EA        NOP                             A:FF X:02 Y:9C P:27 SP:FB
	// E9ED  20 47 FA  JSR $FA47                       A:FF X:02 Y:9C P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:02 Y:9C P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:02 Y:9C P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:02 Y:9C P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:02 Y:9C P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:02 Y:9C P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:02 Y:9C P:27 SP:F9
	// FA53  60        RTS                             A:FF X:02 Y:9C P:27 SP:F9
	// E9F0  AD 47 06  LDA $0647 = FF                  A:FF X:02 Y:9C P:27 SP:FB
	// E9F3  C9 FF     CMP #$FF                        A:FF X:02 Y:9C P:A5 SP:FB
	// E9F5  F0 02     BEQ $E9F9                       A:FF X:02 Y:9C P:27 SP:FB
	// E9F9  C8        INY                             A:FF X:02 Y:9C P:27 SP:FB
	// E9FA  A9 37     LDA #$37                        A:FF X:02 Y:9D P:A5 SP:FB
	// E9FC  8D 47 06  STA $0647 = FF                  A:37 X:02 Y:9D P:25 SP:FB
	// E9FF  20 54 FA  JSR $FA54                       A:37 X:02 Y:9D P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:02 Y:9D P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:02 Y:9D P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:02 Y:9D P:E5 SP:F9
	// EA02  CF 47 06 *DCP $0647 = 37                  A:F0 X:02 Y:9D P:E5 SP:FB
	// EA05  EA        NOP                             A:F0 X:02 Y:9D P:E5 SP:FB
	// EA06  EA        NOP                             A:F0 X:02 Y:9D P:E5 SP:FB
	// EA07  EA        NOP                             A:F0 X:02 Y:9D P:E5 SP:FB
	// EA08  EA        NOP                             A:F0 X:02 Y:9D P:E5 SP:FB
	// EA09  20 59 FA  JSR $FA59                       A:F0 X:02 Y:9D P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:02 Y:9D P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:02 Y:9D P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:02 Y:9D P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:02 Y:9D P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:02 Y:9D P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:02 Y:9D P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:02 Y:9D P:67 SP:F9
	// EA0C  AD 47 06  LDA $0647 = 36                  A:F0 X:02 Y:9D P:67 SP:FB
	// EA0F  C9 36     CMP #$36                        A:36 X:02 Y:9D P:65 SP:FB
	// EA11  F0 02     BEQ $EA15                       A:36 X:02 Y:9D P:67 SP:FB
	// EA15  A9 EB     LDA #$EB                        A:36 X:02 Y:9D P:67 SP:FB
	// EA17  8D 47 06  STA $0647 = 36                  A:EB X:02 Y:9D P:E5 SP:FB
	// EA1A  A9 48     LDA #$48                        A:EB X:02 Y:9D P:E5 SP:FB
	// EA1C  85 45     STA $45 = 32                    A:48 X:02 Y:9D P:65 SP:FB
	// EA1E  A9 05     LDA #$05                        A:48 X:02 Y:9D P:65 SP:FB
	// EA20  85 46     STA $46 = 04                    A:05 X:02 Y:9D P:65 SP:FB
	// EA22  A0 FF     LDY #$FF                        A:05 X:02 Y:9D P:65 SP:FB
	// EA24  20 31 FA  JSR $FA31                       A:05 X:02 Y:FF P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:E5 SP:F9
	// FA33  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:05 X:02 Y:FF P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:02 Y:FF P:64 SP:F9
	// EA27  D3 45    *DCP ($45),Y = 0548 @ 0647 = EB  A:40 X:02 Y:FF P:64 SP:FB
	// EA29  EA        NOP                             A:40 X:02 Y:FF P:64 SP:FB
	// EA2A  EA        NOP                             A:40 X:02 Y:FF P:64 SP:FB
	// EA2B  08        PHP                             A:40 X:02 Y:FF P:64 SP:FB
	// EA2C  48        PHA                             A:40 X:02 Y:FF P:64 SP:FA
	// EA2D  A0 9E     LDY #$9E                        A:40 X:02 Y:FF P:64 SP:F9
	// EA2F  68        PLA                             A:40 X:02 Y:9E P:E4 SP:F9
	// EA30  28        PLP                             A:40 X:02 Y:9E P:64 SP:FA
	// EA31  20 37 FA  JSR $FA37                       A:40 X:02 Y:9E P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:02 Y:9E P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:02 Y:9E P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:02 Y:9E P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:02 Y:9E P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:02 Y:9E P:67 SP:F9
	// FA41  60        RTS                             A:40 X:02 Y:9E P:67 SP:F9
	// EA34  AD 47 06  LDA $0647 = EA                  A:40 X:02 Y:9E P:67 SP:FB
	// EA37  C9 EA     CMP #$EA                        A:EA X:02 Y:9E P:E5 SP:FB
	// EA39  F0 02     BEQ $EA3D                       A:EA X:02 Y:9E P:67 SP:FB
	// EA3D  A0 FF     LDY #$FF                        A:EA X:02 Y:9E P:67 SP:FB
	// EA3F  A9 00     LDA #$00                        A:EA X:02 Y:FF P:E5 SP:FB
	// EA41  8D 47 06  STA $0647 = EA                  A:00 X:02 Y:FF P:67 SP:FB
	// EA44  20 42 FA  JSR $FA42                       A:00 X:02 Y:FF P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:02 Y:FF P:67 SP:F9
	// FA43  38        SEC                             A:00 X:02 Y:FF P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:02 Y:FF P:27 SP:F9
	// FA46  60        RTS                             A:FF X:02 Y:FF P:A5 SP:F9
	// EA47  D3 45    *DCP ($45),Y = 0548 @ 0647 = 00  A:FF X:02 Y:FF P:A5 SP:FB
	// EA49  EA        NOP                             A:FF X:02 Y:FF P:27 SP:FB
	// EA4A  EA        NOP                             A:FF X:02 Y:FF P:27 SP:FB
	// EA4B  08        PHP                             A:FF X:02 Y:FF P:27 SP:FB
	// EA4C  48        PHA                             A:FF X:02 Y:FF P:27 SP:FA
	// EA4D  A0 9F     LDY #$9F                        A:FF X:02 Y:FF P:27 SP:F9
	// EA4F  68        PLA                             A:FF X:02 Y:9F P:A5 SP:F9
	// EA50  28        PLP                             A:FF X:02 Y:9F P:A5 SP:FA
	// EA51  20 47 FA  JSR $FA47                       A:FF X:02 Y:9F P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:02 Y:9F P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:02 Y:9F P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:02 Y:9F P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:02 Y:9F P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:02 Y:9F P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:02 Y:9F P:27 SP:F9
	// FA53  60        RTS                             A:FF X:02 Y:9F P:27 SP:F9
	// EA54  AD 47 06  LDA $0647 = FF                  A:FF X:02 Y:9F P:27 SP:FB
	// EA57  C9 FF     CMP #$FF                        A:FF X:02 Y:9F P:A5 SP:FB
	// EA59  F0 02     BEQ $EA5D                       A:FF X:02 Y:9F P:27 SP:FB
	// EA5D  A0 FF     LDY #$FF                        A:FF X:02 Y:9F P:27 SP:FB
	// EA5F  A9 37     LDA #$37                        A:FF X:02 Y:FF P:A5 SP:FB
	// EA61  8D 47 06  STA $0647 = FF                  A:37 X:02 Y:FF P:25 SP:FB
	// EA64  20 54 FA  JSR $FA54                       A:37 X:02 Y:FF P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:02 Y:FF P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:02 Y:FF P:E5 SP:F9
	// EA67  D3 45    *DCP ($45),Y = 0548 @ 0647 = 37  A:F0 X:02 Y:FF P:E5 SP:FB
	// EA69  EA        NOP                             A:F0 X:02 Y:FF P:E5 SP:FB
	// EA6A  EA        NOP                             A:F0 X:02 Y:FF P:E5 SP:FB
	// EA6B  08        PHP                             A:F0 X:02 Y:FF P:E5 SP:FB
	// EA6C  48        PHA                             A:F0 X:02 Y:FF P:E5 SP:FA
	// EA6D  A0 A0     LDY #$A0                        A:F0 X:02 Y:FF P:E5 SP:F9
	// EA6F  68        PLA                             A:F0 X:02 Y:A0 P:E5 SP:F9
	// EA70  28        PLP                             A:F0 X:02 Y:A0 P:E5 SP:FA
	// EA71  20 59 FA  JSR $FA59                       A:F0 X:02 Y:A0 P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:02 Y:A0 P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:02 Y:A0 P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:02 Y:A0 P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:02 Y:A0 P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:02 Y:A0 P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:02 Y:A0 P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:02 Y:A0 P:67 SP:F9
	// EA74  AD 47 06  LDA $0647 = 36                  A:F0 X:02 Y:A0 P:67 SP:FB
	// EA77  C9 36     CMP #$36                        A:36 X:02 Y:A0 P:65 SP:FB
	// EA79  F0 02     BEQ $EA7D                       A:36 X:02 Y:A0 P:67 SP:FB
	// EA7D  A0 A1     LDY #$A1                        A:36 X:02 Y:A0 P:67 SP:FB
	// EA7F  A2 FF     LDX #$FF                        A:36 X:02 Y:A1 P:E5 SP:FB
	// EA81  A9 EB     LDA #$EB                        A:36 X:FF Y:A1 P:E5 SP:FB
	// EA83  85 47     STA $47 = 36                    A:EB X:FF Y:A1 P:E5 SP:FB
	// EA85  20 31 FA  JSR $FA31                       A:EB X:FF Y:A1 P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:A1 P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:FF Y:A1 P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:FF Y:A1 P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:FF Y:A1 P:64 SP:F9
	// EA88  D7 48    *DCP $48,X @ 47 = EB             A:40 X:FF Y:A1 P:64 SP:FB
	// EA8A  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB
	// EA8B  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB
	// EA8C  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB
	// EA8D  EA        NOP                             A:40 X:FF Y:A1 P:64 SP:FB
	// EA8E  20 37 FA  JSR $FA37                       A:40 X:FF Y:A1 P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A1 P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A1 P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A1 P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A1 P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A1 P:67 SP:F9
	// FA41  60        RTS                             A:40 X:FF Y:A1 P:67 SP:F9
	// EA91  A5 47     LDA $47 = EA                    A:40 X:FF Y:A1 P:67 SP:FB
	// EA93  C9 EA     CMP #$EA                        A:EA X:FF Y:A1 P:E5 SP:FB
	// EA95  F0 02     BEQ $EA99                       A:EA X:FF Y:A1 P:67 SP:FB
	// EA99  C8        INY                             A:EA X:FF Y:A1 P:67 SP:FB
	// EA9A  A9 00     LDA #$00                        A:EA X:FF Y:A2 P:E5 SP:FB
	// EA9C  85 47     STA $47 = EA                    A:00 X:FF Y:A2 P:67 SP:FB
	// EA9E  20 42 FA  JSR $FA42                       A:00 X:FF Y:A2 P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:FF Y:A2 P:67 SP:F9
	// FA43  38        SEC                             A:00 X:FF Y:A2 P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:FF Y:A2 P:27 SP:F9
	// FA46  60        RTS                             A:FF X:FF Y:A2 P:A5 SP:F9
	// EAA1  D7 48    *DCP $48,X @ 47 = 00             A:FF X:FF Y:A2 P:A5 SP:FB
	// EAA3  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB
	// EAA4  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB
	// EAA5  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB
	// EAA6  EA        NOP                             A:FF X:FF Y:A2 P:27 SP:FB
	// EAA7  20 47 FA  JSR $FA47                       A:FF X:FF Y:A2 P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A2 P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A2 P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A2 P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A2 P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A2 P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A2 P:27 SP:F9
	// FA53  60        RTS                             A:FF X:FF Y:A2 P:27 SP:F9
	// EAAA  A5 47     LDA $47 = FF                    A:FF X:FF Y:A2 P:27 SP:FB
	// EAAC  C9 FF     CMP #$FF                        A:FF X:FF Y:A2 P:A5 SP:FB
	// EAAE  F0 02     BEQ $EAB2                       A:FF X:FF Y:A2 P:27 SP:FB
	// EAB2  C8        INY                             A:FF X:FF Y:A2 P:27 SP:FB
	// EAB3  A9 37     LDA #$37                        A:FF X:FF Y:A3 P:A5 SP:FB
	// EAB5  85 47     STA $47 = FF                    A:37 X:FF Y:A3 P:25 SP:FB
	// EAB7  20 54 FA  JSR $FA54                       A:37 X:FF Y:A3 P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:A3 P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:A3 P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:FF Y:A3 P:E5 SP:F9
	// EABA  D7 48    *DCP $48,X @ 47 = 37             A:F0 X:FF Y:A3 P:E5 SP:FB
	// EABC  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB
	// EABD  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB
	// EABE  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB
	// EABF  EA        NOP                             A:F0 X:FF Y:A3 P:E5 SP:FB
	// EAC0  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A3 P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A3 P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A3 P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A3 P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:FF Y:A3 P:67 SP:F9
	// EAC3  A5 47     LDA $47 = 36                    A:F0 X:FF Y:A3 P:67 SP:FB
	// EAC5  C9 36     CMP #$36                        A:36 X:FF Y:A3 P:65 SP:FB
	// EAC7  F0 02     BEQ $EACB                       A:36 X:FF Y:A3 P:67 SP:FB
	// EACB  A9 EB     LDA #$EB                        A:36 X:FF Y:A3 P:67 SP:FB
	// EACD  8D 47 06  STA $0647 = 36                  A:EB X:FF Y:A3 P:E5 SP:FB
	// EAD0  A0 FF     LDY #$FF                        A:EB X:FF Y:A3 P:E5 SP:FB
	// EAD2  20 31 FA  JSR $FA31                       A:EB X:FF Y:FF P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:FF P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:FF Y:FF P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:FF Y:FF P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:FF Y:FF P:64 SP:F9
	// EAD5  DB 48 05 *DCP $0548,Y @ 0647 = EB         A:40 X:FF Y:FF P:64 SP:FB
	// EAD8  EA        NOP                             A:40 X:FF Y:FF P:64 SP:FB
	// EAD9  EA        NOP                             A:40 X:FF Y:FF P:64 SP:FB
	// EADA  08        PHP                             A:40 X:FF Y:FF P:64 SP:FB
	// EADB  48        PHA                             A:40 X:FF Y:FF P:64 SP:FA
	// EADC  A0 A4     LDY #$A4                        A:40 X:FF Y:FF P:64 SP:F9
	// EADE  68        PLA                             A:40 X:FF Y:A4 P:E4 SP:F9
	// EADF  28        PLP                             A:40 X:FF Y:A4 P:64 SP:FA
	// EAE0  20 37 FA  JSR $FA37                       A:40 X:FF Y:A4 P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A4 P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A4 P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A4 P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A4 P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A4 P:67 SP:F9
	// FA41  60        RTS                             A:40 X:FF Y:A4 P:67 SP:F9
	// EAE3  AD 47 06  LDA $0647 = EA                  A:40 X:FF Y:A4 P:67 SP:FB
	// EAE6  C9 EA     CMP #$EA                        A:EA X:FF Y:A4 P:E5 SP:FB
	// EAE8  F0 02     BEQ $EAEC                       A:EA X:FF Y:A4 P:67 SP:FB
	// EAEC  A0 FF     LDY #$FF                        A:EA X:FF Y:A4 P:67 SP:FB
	// EAEE  A9 00     LDA #$00                        A:EA X:FF Y:FF P:E5 SP:FB
	// EAF0  8D 47 06  STA $0647 = EA                  A:00 X:FF Y:FF P:67 SP:FB
	// EAF3  20 42 FA  JSR $FA42                       A:00 X:FF Y:FF P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:FF Y:FF P:67 SP:F9
	// FA43  38        SEC                             A:00 X:FF Y:FF P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:FF Y:FF P:27 SP:F9
	// FA46  60        RTS                             A:FF X:FF Y:FF P:A5 SP:F9
	// EAF6  DB 48 05 *DCP $0548,Y @ 0647 = 00         A:FF X:FF Y:FF P:A5 SP:FB
	// EAF9  EA        NOP                             A:FF X:FF Y:FF P:27 SP:FB
	// EAFA  EA        NOP                             A:FF X:FF Y:FF P:27 SP:FB
	// EAFB  08        PHP                             A:FF X:FF Y:FF P:27 SP:FB
	// EAFC  48        PHA                             A:FF X:FF Y:FF P:27 SP:FA
	// EAFD  A0 A5     LDY #$A5                        A:FF X:FF Y:FF P:27 SP:F9
	// EAFF  68        PLA                             A:FF X:FF Y:A5 P:A5 SP:F9
	// EB00  28        PLP                             A:FF X:FF Y:A5 P:A5 SP:FA
	// EB01  20 47 FA  JSR $FA47                       A:FF X:FF Y:A5 P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A5 P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A5 P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A5 P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A5 P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A5 P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A5 P:27 SP:F9
	// FA53  60        RTS                             A:FF X:FF Y:A5 P:27 SP:F9
	// EB04  AD 47 06  LDA $0647 = FF                  A:FF X:FF Y:A5 P:27 SP:FB
	// EB07  C9 FF     CMP #$FF                        A:FF X:FF Y:A5 P:A5 SP:FB
	// EB09  F0 02     BEQ $EB0D                       A:FF X:FF Y:A5 P:27 SP:FB
	// EB0D  A0 FF     LDY #$FF                        A:FF X:FF Y:A5 P:27 SP:FB
	// EB0F  A9 37     LDA #$37                        A:FF X:FF Y:FF P:A5 SP:FB
	// EB11  8D 47 06  STA $0647 = FF                  A:37 X:FF Y:FF P:25 SP:FB
	// EB14  20 54 FA  JSR $FA54                       A:37 X:FF Y:FF P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:FF P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:FF Y:FF P:E5 SP:F9
	// EB17  DB 48 05 *DCP $0548,Y @ 0647 = 37         A:F0 X:FF Y:FF P:E5 SP:FB
	// EB1A  EA        NOP                             A:F0 X:FF Y:FF P:E5 SP:FB
	// EB1B  EA        NOP                             A:F0 X:FF Y:FF P:E5 SP:FB
	// EB1C  08        PHP                             A:F0 X:FF Y:FF P:E5 SP:FB
	// EB1D  48        PHA                             A:F0 X:FF Y:FF P:E5 SP:FA
	// EB1E  A0 A6     LDY #$A6                        A:F0 X:FF Y:FF P:E5 SP:F9
	// EB20  68        PLA                             A:F0 X:FF Y:A6 P:E5 SP:F9
	// EB21  28        PLP                             A:F0 X:FF Y:A6 P:E5 SP:FA
	// EB22  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A6 P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A6 P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A6 P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A6 P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:FF Y:A6 P:67 SP:F9
	// EB25  AD 47 06  LDA $0647 = 36                  A:F0 X:FF Y:A6 P:67 SP:FB
	// EB28  C9 36     CMP #$36                        A:36 X:FF Y:A6 P:65 SP:FB
	// EB2A  F0 02     BEQ $EB2E                       A:36 X:FF Y:A6 P:67 SP:FB
	// EB2E  A0 A7     LDY #$A7                        A:36 X:FF Y:A6 P:67 SP:FB
	// EB30  A2 FF     LDX #$FF                        A:36 X:FF Y:A7 P:E5 SP:FB
	// EB32  A9 EB     LDA #$EB                        A:36 X:FF Y:A7 P:E5 SP:FB
	// EB34  8D 47 06  STA $0647 = 36                  A:EB X:FF Y:A7 P:E5 SP:FB
	// EB37  20 31 FA  JSR $FA31                       A:EB X:FF Y:A7 P:E5 SP:FB
	// FA31  24 01     BIT $01 = FF                    A:EB X:FF Y:A7 P:E5 SP:F9
	// FA33  18        CLC                             A:EB X:FF Y:A7 P:E5 SP:F9
	// FA34  A9 40     LDA #$40                        A:EB X:FF Y:A7 P:E4 SP:F9
	// FA36  60        RTS                             A:40 X:FF Y:A7 P:64 SP:F9
	// EB3A  DF 48 05 *DCP $0548,X @ 0647 = EB         A:40 X:FF Y:A7 P:64 SP:FB
	// EB3D  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB
	// EB3E  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB
	// EB3F  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB
	// EB40  EA        NOP                             A:40 X:FF Y:A7 P:64 SP:FB
	// EB41  20 37 FA  JSR $FA37                       A:40 X:FF Y:A7 P:64 SP:FB
	// FA37  50 2C     BVC $FA65                       A:40 X:FF Y:A7 P:64 SP:F9
	// FA39  B0 2A     BCS $FA65                       A:40 X:FF Y:A7 P:64 SP:F9
	// FA3B  30 28     BMI $FA65                       A:40 X:FF Y:A7 P:64 SP:F9
	// FA3D  C9 40     CMP #$40                        A:40 X:FF Y:A7 P:64 SP:F9
	// FA3F  D0 24     BNE $FA65                       A:40 X:FF Y:A7 P:67 SP:F9
	// FA41  60        RTS                             A:40 X:FF Y:A7 P:67 SP:F9
	// EB44  AD 47 06  LDA $0647 = EA                  A:40 X:FF Y:A7 P:67 SP:FB
	// EB47  C9 EA     CMP #$EA                        A:EA X:FF Y:A7 P:E5 SP:FB
	// EB49  F0 02     BEQ $EB4D                       A:EA X:FF Y:A7 P:67 SP:FB
	// EB4D  C8        INY                             A:EA X:FF Y:A7 P:67 SP:FB
	// EB4E  A9 00     LDA #$00                        A:EA X:FF Y:A8 P:E5 SP:FB
	// EB50  8D 47 06  STA $0647 = EA                  A:00 X:FF Y:A8 P:67 SP:FB
	// EB53  20 42 FA  JSR $FA42                       A:00 X:FF Y:A8 P:67 SP:FB
	// FA42  B8        CLV                             A:00 X:FF Y:A8 P:67 SP:F9
	// FA43  38        SEC                             A:00 X:FF Y:A8 P:27 SP:F9
	// FA44  A9 FF     LDA #$FF                        A:00 X:FF Y:A8 P:27 SP:F9
	// FA46  60        RTS                             A:FF X:FF Y:A8 P:A5 SP:F9
	// EB56  DF 48 05 *DCP $0548,X @ 0647 = 00         A:FF X:FF Y:A8 P:A5 SP:FB
	// EB59  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB
	// EB5A  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB
	// EB5B  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB
	// EB5C  EA        NOP                             A:FF X:FF Y:A8 P:27 SP:FB
	// EB5D  20 47 FA  JSR $FA47                       A:FF X:FF Y:A8 P:27 SP:FB
	// FA47  70 1C     BVS $FA65                       A:FF X:FF Y:A8 P:27 SP:F9
	// FA49  D0 1A     BNE $FA65                       A:FF X:FF Y:A8 P:27 SP:F9
	// FA4B  30 18     BMI $FA65                       A:FF X:FF Y:A8 P:27 SP:F9
	// FA4D  90 16     BCC $FA65                       A:FF X:FF Y:A8 P:27 SP:F9
	// FA4F  C9 FF     CMP #$FF                        A:FF X:FF Y:A8 P:27 SP:F9
	// FA51  D0 12     BNE $FA65                       A:FF X:FF Y:A8 P:27 SP:F9
	// FA53  60        RTS                             A:FF X:FF Y:A8 P:27 SP:F9
	// EB60  AD 47 06  LDA $0647 = FF                  A:FF X:FF Y:A8 P:27 SP:FB
	// EB63  C9 FF     CMP #$FF                        A:FF X:FF Y:A8 P:A5 SP:FB
	// EB65  F0 02     BEQ $EB69                       A:FF X:FF Y:A8 P:27 SP:FB
	// EB69  C8        INY                             A:FF X:FF Y:A8 P:27 SP:FB
	// EB6A  A9 37     LDA #$37                        A:FF X:FF Y:A9 P:A5 SP:FB
	// EB6C  8D 47 06  STA $0647 = FF                  A:37 X:FF Y:A9 P:25 SP:FB
	// EB6F  20 54 FA  JSR $FA54                       A:37 X:FF Y:A9 P:25 SP:FB
	// FA54  24 01     BIT $01 = FF                    A:37 X:FF Y:A9 P:25 SP:F9
	// FA56  A9 F0     LDA #$F0                        A:37 X:FF Y:A9 P:E5 SP:F9
	// FA58  60        RTS                             A:F0 X:FF Y:A9 P:E5 SP:F9
	// EB72  DF 48 05 *DCP $0548,X @ 0647 = 37         A:F0 X:FF Y:A9 P:E5 SP:FB
	// EB75  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB
	// EB76  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB
	// EB77  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB
	// EB78  EA        NOP                             A:F0 X:FF Y:A9 P:E5 SP:FB
	// EB79  20 59 FA  JSR $FA59                       A:F0 X:FF Y:A9 P:E5 SP:FB
	// FA59  50 0A     BVC $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9
	// FA5B  F0 08     BEQ $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9
	// FA5D  10 06     BPL $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9
	// FA5F  90 04     BCC $FA65                       A:F0 X:FF Y:A9 P:E5 SP:F9
	// FA61  C9 F0     CMP #$F0                        A:F0 X:FF Y:A9 P:E5 SP:F9
	// FA63  F0 02     BEQ $FA67                       A:F0 X:FF Y:A9 P:67 SP:F9
	// FA67  60        RTS                             A:F0 X:FF Y:A9 P:67 SP:F9
	// EB7C  AD 47 06  LDA $0647 = 36                  A:F0 X:FF Y:A9 P:67 SP:FB
	// EB7F  C9 36     CMP #$36                        A:36 X:FF Y:A9 P:65 SP:FB
	// EB81  F0 02     BEQ $EB85                       A:36 X:FF Y:A9 P:67 SP:FB
	// EB85  60        RTS                             A:36 X:FF Y:A9 P:67 SP:FB
	// C63E  20 86 EB  JSR $EB86                       A:36 X:FF Y:A9 P:67 SP:FD
	// EB86  A9 FF     LDA #$FF                        A:36 X:FF Y:A9 P:67 SP:FB
	// EB88  85 01     STA $01 = FF                    A:FF X:FF Y:A9 P:E5 SP:FB
	// EB8A  A0 AA     LDY #$AA                        A:FF X:FF Y:A9 P:E5 SP:FB
	// EB8C  A2 02     LDX #$02                        A:FF X:FF Y:AA P:E5 SP:FB
	// EB8E  A9 47     LDA #$47                        A:FF X:02 Y:AA P:65 SP:FB
	// EB90  85 47     STA $47 = 36                    A:47 X:02 Y:AA P:65 SP:FB
	// EB92  A9 06     LDA #$06                        A:47 X:02 Y:AA P:65 SP:FB
	// EB94  85 48     STA $48 = 06                    A:06 X:02 Y:AA P:65 SP:FB
	// EB96  A9 EB     LDA #$EB                        A:06 X:02 Y:AA P:65 SP:FB
	// EB98  8D 47 06  STA $0647 = 36                  A:EB X:02 Y:AA P:E5 SP:FB
	// EB9B  20 B1 FA  JSR $FAB1                       A:EB X:02 Y:AA P:E5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:02 Y:AA P:E5 SP:F9
	// FAB3  18        CLC                             A:EB X:02 Y:AA P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:02 Y:AA P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:02 Y:AA P:64 SP:F9
	// EB9E  E3 45    *ISB ($45,X) @ 47 = 0647 = EB    A:40 X:02 Y:AA P:64 SP:FB
	// EBA0  EA        NOP                             A:53 X:02 Y:AA P:24 SP:FB
	// EBA1  EA        NOP                             A:53 X:02 Y:AA P:24 SP:FB
	// EBA2  EA        NOP                             A:53 X:02 Y:AA P:24 SP:FB
	// EBA3  EA        NOP                             A:53 X:02 Y:AA P:24 SP:FB
	// EBA4  20 B7 FA  JSR $FAB7                       A:53 X:02 Y:AA P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:02 Y:AA P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:02 Y:AA P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:02 Y:AA P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:02 Y:AA P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:02 Y:AA P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:02 Y:AA P:27 SP:F9
	// EBA7  AD 47 06  LDA $0647 = EC                  A:53 X:02 Y:AA P:27 SP:FB
	// EBAA  C9 EC     CMP #$EC                        A:EC X:02 Y:AA P:A5 SP:FB
	// EBAC  F0 02     BEQ $EBB0                       A:EC X:02 Y:AA P:27 SP:FB
	// EBB0  C8        INY                             A:EC X:02 Y:AA P:27 SP:FB
	// EBB1  A9 FF     LDA #$FF                        A:EC X:02 Y:AB P:A5 SP:FB
	// EBB3  8D 47 06  STA $0647 = EC                  A:FF X:02 Y:AB P:A5 SP:FB
	// EBB6  20 C2 FA  JSR $FAC2                       A:FF X:02 Y:AB P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:02 Y:AB P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:02 Y:AB P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:02 Y:AB P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:02 Y:AB P:A5 SP:F9
	// EBB9  E3 45    *ISB ($45,X) @ 47 = 0647 = FF    A:FF X:02 Y:AB P:A5 SP:FB
	// EBBB  EA        NOP                             A:FF X:02 Y:AB P:A5 SP:FB
	// EBBC  EA        NOP                             A:FF X:02 Y:AB P:A5 SP:FB
	// EBBD  EA        NOP                             A:FF X:02 Y:AB P:A5 SP:FB
	// EBBE  EA        NOP                             A:FF X:02 Y:AB P:A5 SP:FB
	// EBBF  20 C7 FA  JSR $FAC7                       A:FF X:02 Y:AB P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:02 Y:AB P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:02 Y:AB P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:02 Y:AB P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:02 Y:AB P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:02 Y:AB P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:02 Y:AB P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:02 Y:AB P:27 SP:F9
	// EBC2  AD 47 06  LDA $0647 = 00                  A:FF X:02 Y:AB P:27 SP:FB
	// EBC5  C9 00     CMP #$00                        A:00 X:02 Y:AB P:27 SP:FB
	// EBC7  F0 02     BEQ $EBCB                       A:00 X:02 Y:AB P:27 SP:FB
	// EBCB  C8        INY                             A:00 X:02 Y:AB P:27 SP:FB
	// EBCC  A9 37     LDA #$37                        A:00 X:02 Y:AC P:A5 SP:FB
	// EBCE  8D 47 06  STA $0647 = 00                  A:37 X:02 Y:AC P:25 SP:FB
	// EBD1  20 D4 FA  JSR $FAD4                       A:37 X:02 Y:AC P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:02 Y:AC P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:02 Y:AC P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:02 Y:AC P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:02 Y:AC P:E5 SP:F9
	// EBD4  E3 45    *ISB ($45,X) @ 47 = 0647 = 37    A:F0 X:02 Y:AC P:E5 SP:FB
	// EBD6  EA        NOP                             A:B8 X:02 Y:AC P:A5 SP:FB
	// EBD7  EA        NOP                             A:B8 X:02 Y:AC P:A5 SP:FB
	// EBD8  EA        NOP                             A:B8 X:02 Y:AC P:A5 SP:FB
	// EBD9  EA        NOP                             A:B8 X:02 Y:AC P:A5 SP:FB
	// EBDA  20 DA FA  JSR $FADA                       A:B8 X:02 Y:AC P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:02 Y:AC P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:02 Y:AC P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:02 Y:AC P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:02 Y:AC P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:02 Y:AC P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:02 Y:AC P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:02 Y:AC P:27 SP:F9
	// EBDD  AD 47 06  LDA $0647 = 38                  A:B8 X:02 Y:AC P:27 SP:FB
	// EBE0  C9 38     CMP #$38                        A:38 X:02 Y:AC P:25 SP:FB
	// EBE2  F0 02     BEQ $EBE6                       A:38 X:02 Y:AC P:27 SP:FB
	// EBE6  C8        INY                             A:38 X:02 Y:AC P:27 SP:FB
	// EBE7  A9 EB     LDA #$EB                        A:38 X:02 Y:AD P:A5 SP:FB
	// EBE9  85 47     STA $47 = 47                    A:EB X:02 Y:AD P:A5 SP:FB
	// EBEB  20 B1 FA  JSR $FAB1                       A:EB X:02 Y:AD P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:02 Y:AD P:A5 SP:F9
	// FAB3  18        CLC                             A:EB X:02 Y:AD P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:02 Y:AD P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:02 Y:AD P:64 SP:F9
	// EBEE  E7 47    *ISB $47 = EB                    A:40 X:02 Y:AD P:64 SP:FB
	// EBF0  EA        NOP                             A:53 X:02 Y:AD P:24 SP:FB
	// EBF1  EA        NOP                             A:53 X:02 Y:AD P:24 SP:FB
	// EBF2  EA        NOP                             A:53 X:02 Y:AD P:24 SP:FB
	// EBF3  EA        NOP                             A:53 X:02 Y:AD P:24 SP:FB
	// EBF4  20 B7 FA  JSR $FAB7                       A:53 X:02 Y:AD P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:02 Y:AD P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:02 Y:AD P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:02 Y:AD P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:02 Y:AD P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:02 Y:AD P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:02 Y:AD P:27 SP:F9
	// EBF7  A5 47     LDA $47 = EC                    A:53 X:02 Y:AD P:27 SP:FB
	// EBF9  C9 EC     CMP #$EC                        A:EC X:02 Y:AD P:A5 SP:FB
	// EBFB  F0 02     BEQ $EBFF                       A:EC X:02 Y:AD P:27 SP:FB
	// EBFF  C8        INY                             A:EC X:02 Y:AD P:27 SP:FB
	// EC00  A9 FF     LDA #$FF                        A:EC X:02 Y:AE P:A5 SP:FB
	// EC02  85 47     STA $47 = EC                    A:FF X:02 Y:AE P:A5 SP:FB
	// EC04  20 C2 FA  JSR $FAC2                       A:FF X:02 Y:AE P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:02 Y:AE P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:02 Y:AE P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:02 Y:AE P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:02 Y:AE P:A5 SP:F9
	// EC07  E7 47    *ISB $47 = FF                    A:FF X:02 Y:AE P:A5 SP:FB
	// EC09  EA        NOP                             A:FF X:02 Y:AE P:A5 SP:FB
	// EC0A  EA        NOP                             A:FF X:02 Y:AE P:A5 SP:FB
	// EC0B  EA        NOP                             A:FF X:02 Y:AE P:A5 SP:FB
	// EC0C  EA        NOP                             A:FF X:02 Y:AE P:A5 SP:FB
	// EC0D  20 C7 FA  JSR $FAC7                       A:FF X:02 Y:AE P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:02 Y:AE P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:02 Y:AE P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:02 Y:AE P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:02 Y:AE P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:02 Y:AE P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:02 Y:AE P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:02 Y:AE P:27 SP:F9
	// EC10  A5 47     LDA $47 = 00                    A:FF X:02 Y:AE P:27 SP:FB
	// EC12  C9 00     CMP #$00                        A:00 X:02 Y:AE P:27 SP:FB
	// EC14  F0 02     BEQ $EC18                       A:00 X:02 Y:AE P:27 SP:FB
	// EC18  C8        INY                             A:00 X:02 Y:AE P:27 SP:FB
	// EC19  A9 37     LDA #$37                        A:00 X:02 Y:AF P:A5 SP:FB
	// EC1B  85 47     STA $47 = 00                    A:37 X:02 Y:AF P:25 SP:FB
	// EC1D  20 D4 FA  JSR $FAD4                       A:37 X:02 Y:AF P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:02 Y:AF P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:02 Y:AF P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:02 Y:AF P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:02 Y:AF P:E5 SP:F9
	// EC20  E7 47    *ISB $47 = 37                    A:F0 X:02 Y:AF P:E5 SP:FB
	// EC22  EA        NOP                             A:B8 X:02 Y:AF P:A5 SP:FB
	// EC23  EA        NOP                             A:B8 X:02 Y:AF P:A5 SP:FB
	// EC24  EA        NOP                             A:B8 X:02 Y:AF P:A5 SP:FB
	// EC25  EA        NOP                             A:B8 X:02 Y:AF P:A5 SP:FB
	// EC26  20 DA FA  JSR $FADA                       A:B8 X:02 Y:AF P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:02 Y:AF P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:02 Y:AF P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:02 Y:AF P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:02 Y:AF P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:02 Y:AF P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:02 Y:AF P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:02 Y:AF P:27 SP:F9
	// EC29  A5 47     LDA $47 = 38                    A:B8 X:02 Y:AF P:27 SP:FB
	// EC2B  C9 38     CMP #$38                        A:38 X:02 Y:AF P:25 SP:FB
	// EC2D  F0 02     BEQ $EC31                       A:38 X:02 Y:AF P:27 SP:FB
	// EC31  C8        INY                             A:38 X:02 Y:AF P:27 SP:FB
	// EC32  A9 EB     LDA #$EB                        A:38 X:02 Y:B0 P:A5 SP:FB
	// EC34  8D 47 06  STA $0647 = 38                  A:EB X:02 Y:B0 P:A5 SP:FB
	// EC37  20 B1 FA  JSR $FAB1                       A:EB X:02 Y:B0 P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:02 Y:B0 P:A5 SP:F9
	// FAB3  18        CLC                             A:EB X:02 Y:B0 P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:02 Y:B0 P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:02 Y:B0 P:64 SP:F9
	// EC3A  EF 47 06 *ISB $0647 = EB                  A:40 X:02 Y:B0 P:64 SP:FB
	// EC3D  EA        NOP                             A:53 X:02 Y:B0 P:24 SP:FB
	// EC3E  EA        NOP                             A:53 X:02 Y:B0 P:24 SP:FB
	// EC3F  EA        NOP                             A:53 X:02 Y:B0 P:24 SP:FB
	// EC40  EA        NOP                             A:53 X:02 Y:B0 P:24 SP:FB
	// EC41  20 B7 FA  JSR $FAB7                       A:53 X:02 Y:B0 P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:02 Y:B0 P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:02 Y:B0 P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:02 Y:B0 P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:02 Y:B0 P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:02 Y:B0 P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:02 Y:B0 P:27 SP:F9
	// EC44  AD 47 06  LDA $0647 = EC                  A:53 X:02 Y:B0 P:27 SP:FB
	// EC47  C9 EC     CMP #$EC                        A:EC X:02 Y:B0 P:A5 SP:FB
	// EC49  F0 02     BEQ $EC4D                       A:EC X:02 Y:B0 P:27 SP:FB
	// EC4D  C8        INY                             A:EC X:02 Y:B0 P:27 SP:FB
	// EC4E  A9 FF     LDA #$FF                        A:EC X:02 Y:B1 P:A5 SP:FB
	// EC50  8D 47 06  STA $0647 = EC                  A:FF X:02 Y:B1 P:A5 SP:FB
	// EC53  20 C2 FA  JSR $FAC2                       A:FF X:02 Y:B1 P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:02 Y:B1 P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:02 Y:B1 P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:02 Y:B1 P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:02 Y:B1 P:A5 SP:F9
	// EC56  EF 47 06 *ISB $0647 = FF                  A:FF X:02 Y:B1 P:A5 SP:FB
	// EC59  EA        NOP                             A:FF X:02 Y:B1 P:A5 SP:FB
	// EC5A  EA        NOP                             A:FF X:02 Y:B1 P:A5 SP:FB
	// EC5B  EA        NOP                             A:FF X:02 Y:B1 P:A5 SP:FB
	// EC5C  EA        NOP                             A:FF X:02 Y:B1 P:A5 SP:FB
	// EC5D  20 C7 FA  JSR $FAC7                       A:FF X:02 Y:B1 P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:02 Y:B1 P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:02 Y:B1 P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:02 Y:B1 P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:02 Y:B1 P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:02 Y:B1 P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:02 Y:B1 P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:02 Y:B1 P:27 SP:F9
	// EC60  AD 47 06  LDA $0647 = 00                  A:FF X:02 Y:B1 P:27 SP:FB
	// EC63  C9 00     CMP #$00                        A:00 X:02 Y:B1 P:27 SP:FB
	// EC65  F0 02     BEQ $EC69                       A:00 X:02 Y:B1 P:27 SP:FB
	// EC69  C8        INY                             A:00 X:02 Y:B1 P:27 SP:FB
	// EC6A  A9 37     LDA #$37                        A:00 X:02 Y:B2 P:A5 SP:FB
	// EC6C  8D 47 06  STA $0647 = 00                  A:37 X:02 Y:B2 P:25 SP:FB
	// EC6F  20 D4 FA  JSR $FAD4                       A:37 X:02 Y:B2 P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:02 Y:B2 P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:02 Y:B2 P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:02 Y:B2 P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:02 Y:B2 P:E5 SP:F9
	// EC72  EF 47 06 *ISB $0647 = 37                  A:F0 X:02 Y:B2 P:E5 SP:FB
	// EC75  EA        NOP                             A:B8 X:02 Y:B2 P:A5 SP:FB
	// EC76  EA        NOP                             A:B8 X:02 Y:B2 P:A5 SP:FB
	// EC77  EA        NOP                             A:B8 X:02 Y:B2 P:A5 SP:FB
	// EC78  EA        NOP                             A:B8 X:02 Y:B2 P:A5 SP:FB
	// EC79  20 DA FA  JSR $FADA                       A:B8 X:02 Y:B2 P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:02 Y:B2 P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:02 Y:B2 P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:02 Y:B2 P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:02 Y:B2 P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:02 Y:B2 P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:02 Y:B2 P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:02 Y:B2 P:27 SP:F9
	// EC7C  AD 47 06  LDA $0647 = 38                  A:B8 X:02 Y:B2 P:27 SP:FB
	// EC7F  C9 38     CMP #$38                        A:38 X:02 Y:B2 P:25 SP:FB
	// EC81  F0 02     BEQ $EC85                       A:38 X:02 Y:B2 P:27 SP:FB
	// EC85  A9 EB     LDA #$EB                        A:38 X:02 Y:B2 P:27 SP:FB
	// EC87  8D 47 06  STA $0647 = 38                  A:EB X:02 Y:B2 P:A5 SP:FB
	// EC8A  A9 48     LDA #$48                        A:EB X:02 Y:B2 P:A5 SP:FB
	// EC8C  85 45     STA $45 = 48                    A:48 X:02 Y:B2 P:25 SP:FB
	// EC8E  A9 05     LDA #$05                        A:48 X:02 Y:B2 P:25 SP:FB
	// EC90  85 46     STA $46 = 05                    A:05 X:02 Y:B2 P:25 SP:FB
	// EC92  A0 FF     LDY #$FF                        A:05 X:02 Y:B2 P:25 SP:FB
	// EC94  20 B1 FA  JSR $FAB1                       A:05 X:02 Y:FF P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:A5 SP:F9
	// FAB3  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:05 X:02 Y:FF P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:02 Y:FF P:64 SP:F9
	// EC97  F3 45    *ISB ($45),Y = 0548 @ 0647 = EB  A:40 X:02 Y:FF P:64 SP:FB
	// EC99  EA        NOP                             A:53 X:02 Y:FF P:24 SP:FB
	// EC9A  EA        NOP                             A:53 X:02 Y:FF P:24 SP:FB
	// EC9B  08        PHP                             A:53 X:02 Y:FF P:24 SP:FB
	// EC9C  48        PHA                             A:53 X:02 Y:FF P:24 SP:FA
	// EC9D  A0 B3     LDY #$B3                        A:53 X:02 Y:FF P:24 SP:F9
	// EC9F  68        PLA                             A:53 X:02 Y:B3 P:A4 SP:F9
	// ECA0  28        PLP                             A:53 X:02 Y:B3 P:24 SP:FA
	// ECA1  20 B7 FA  JSR $FAB7                       A:53 X:02 Y:B3 P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:02 Y:B3 P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:02 Y:B3 P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:02 Y:B3 P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:02 Y:B3 P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:02 Y:B3 P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:02 Y:B3 P:27 SP:F9
	// ECA4  AD 47 06  LDA $0647 = EC                  A:53 X:02 Y:B3 P:27 SP:FB
	// ECA7  C9 EC     CMP #$EC                        A:EC X:02 Y:B3 P:A5 SP:FB
	// ECA9  F0 02     BEQ $ECAD                       A:EC X:02 Y:B3 P:27 SP:FB
	// ECAD  A0 FF     LDY #$FF                        A:EC X:02 Y:B3 P:27 SP:FB
	// ECAF  A9 FF     LDA #$FF                        A:EC X:02 Y:FF P:A5 SP:FB
	// ECB1  8D 47 06  STA $0647 = EC                  A:FF X:02 Y:FF P:A5 SP:FB
	// ECB4  20 C2 FA  JSR $FAC2                       A:FF X:02 Y:FF P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:02 Y:FF P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:02 Y:FF P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:02 Y:FF P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:02 Y:FF P:A5 SP:F9
	// ECB7  F3 45    *ISB ($45),Y = 0548 @ 0647 = FF  A:FF X:02 Y:FF P:A5 SP:FB
	// ECB9  EA        NOP                             A:FF X:02 Y:FF P:A5 SP:FB
	// ECBA  EA        NOP                             A:FF X:02 Y:FF P:A5 SP:FB
	// ECBB  08        PHP                             A:FF X:02 Y:FF P:A5 SP:FB
	// ECBC  48        PHA                             A:FF X:02 Y:FF P:A5 SP:FA
	// ECBD  A0 B4     LDY #$B4                        A:FF X:02 Y:FF P:A5 SP:F9
	// ECBF  68        PLA                             A:FF X:02 Y:B4 P:A5 SP:F9
	// ECC0  28        PLP                             A:FF X:02 Y:B4 P:A5 SP:FA
	// ECC1  20 C7 FA  JSR $FAC7                       A:FF X:02 Y:B4 P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:02 Y:B4 P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:02 Y:B4 P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:02 Y:B4 P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:02 Y:B4 P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:02 Y:B4 P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:02 Y:B4 P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:02 Y:B4 P:27 SP:F9
	// ECC4  AD 47 06  LDA $0647 = 00                  A:FF X:02 Y:B4 P:27 SP:FB
	// ECC7  C9 00     CMP #$00                        A:00 X:02 Y:B4 P:27 SP:FB
	// ECC9  F0 02     BEQ $ECCD                       A:00 X:02 Y:B4 P:27 SP:FB
	// ECCD  A0 FF     LDY #$FF                        A:00 X:02 Y:B4 P:27 SP:FB
	// ECCF  A9 37     LDA #$37                        A:00 X:02 Y:FF P:A5 SP:FB
	// ECD1  8D 47 06  STA $0647 = 00                  A:37 X:02 Y:FF P:25 SP:FB
	// ECD4  20 D4 FA  JSR $FAD4                       A:37 X:02 Y:FF P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:02 Y:FF P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:02 Y:FF P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:02 Y:FF P:E5 SP:F9
	// ECD7  F3 45    *ISB ($45),Y = 0548 @ 0647 = 37  A:F0 X:02 Y:FF P:E5 SP:FB
	// ECD9  EA        NOP                             A:B8 X:02 Y:FF P:A5 SP:FB
	// ECDA  EA        NOP                             A:B8 X:02 Y:FF P:A5 SP:FB
	// ECDB  08        PHP                             A:B8 X:02 Y:FF P:A5 SP:FB
	// ECDC  48        PHA                             A:B8 X:02 Y:FF P:A5 SP:FA
	// ECDD  A0 B5     LDY #$B5                        A:B8 X:02 Y:FF P:A5 SP:F9
	// ECDF  68        PLA                             A:B8 X:02 Y:B5 P:A5 SP:F9
	// ECE0  28        PLP                             A:B8 X:02 Y:B5 P:A5 SP:FA
	// ECE1  20 DA FA  JSR $FADA                       A:B8 X:02 Y:B5 P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:02 Y:B5 P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:02 Y:B5 P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:02 Y:B5 P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:02 Y:B5 P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:02 Y:B5 P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:02 Y:B5 P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:02 Y:B5 P:27 SP:F9
	// ECE4  AD 47 06  LDA $0647 = 38                  A:B8 X:02 Y:B5 P:27 SP:FB
	// ECE7  C9 38     CMP #$38                        A:38 X:02 Y:B5 P:25 SP:FB
	// ECE9  F0 02     BEQ $ECED                       A:38 X:02 Y:B5 P:27 SP:FB
	// ECED  A0 B6     LDY #$B6                        A:38 X:02 Y:B5 P:27 SP:FB
	// ECEF  A2 FF     LDX #$FF                        A:38 X:02 Y:B6 P:A5 SP:FB
	// ECF1  A9 EB     LDA #$EB                        A:38 X:FF Y:B6 P:A5 SP:FB
	// ECF3  85 47     STA $47 = 38                    A:EB X:FF Y:B6 P:A5 SP:FB
	// ECF5  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:B6 P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:B6 P:A5 SP:F9
	// FAB3  18        CLC                             A:EB X:FF Y:B6 P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:FF Y:B6 P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:FF Y:B6 P:64 SP:F9
	// ECF8  F7 48    *ISB $48,X @ 47 = EB             A:40 X:FF Y:B6 P:64 SP:FB
	// ECFA  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB
	// ECFB  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB
	// ECFC  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB
	// ECFD  EA        NOP                             A:53 X:FF Y:B6 P:24 SP:FB
	// ECFE  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:B6 P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:FF Y:B6 P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:FF Y:B6 P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:B6 P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:FF Y:B6 P:27 SP:F9
	// ED01  A5 47     LDA $47 = EC                    A:53 X:FF Y:B6 P:27 SP:FB
	// ED03  C9 EC     CMP #$EC                        A:EC X:FF Y:B6 P:A5 SP:FB
	// ED05  F0 02     BEQ $ED09                       A:EC X:FF Y:B6 P:27 SP:FB
	// ED09  C8        INY                             A:EC X:FF Y:B6 P:27 SP:FB
	// ED0A  A9 FF     LDA #$FF                        A:EC X:FF Y:B7 P:A5 SP:FB
	// ED0C  85 47     STA $47 = EC                    A:FF X:FF Y:B7 P:A5 SP:FB
	// ED0E  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:B7 P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:FF Y:B7 P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:FF Y:B7 P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:B7 P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:FF Y:B7 P:A5 SP:F9
	// ED11  F7 48    *ISB $48,X @ 47 = FF             A:FF X:FF Y:B7 P:A5 SP:FB
	// ED13  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB
	// ED14  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB
	// ED15  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB
	// ED16  EA        NOP                             A:FF X:FF Y:B7 P:A5 SP:FB
	// ED17  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:B7 P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:FF Y:B7 P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:B7 P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:B7 P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:FF Y:B7 P:27 SP:F9
	// ED1A  A5 47     LDA $47 = 00                    A:FF X:FF Y:B7 P:27 SP:FB
	// ED1C  C9 00     CMP #$00                        A:00 X:FF Y:B7 P:27 SP:FB
	// ED1E  F0 02     BEQ $ED22                       A:00 X:FF Y:B7 P:27 SP:FB
	// ED22  C8        INY                             A:00 X:FF Y:B7 P:27 SP:FB
	// ED23  A9 37     LDA #$37                        A:00 X:FF Y:B8 P:A5 SP:FB
	// ED25  85 47     STA $47 = 00                    A:37 X:FF Y:B8 P:25 SP:FB
	// ED27  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:B8 P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:B8 P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:FF Y:B8 P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:B8 P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:FF Y:B8 P:E5 SP:F9
	// ED2A  F7 48    *ISB $48,X @ 47 = 37             A:F0 X:FF Y:B8 P:E5 SP:FB
	// ED2C  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB
	// ED2D  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB
	// ED2E  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB
	// ED2F  EA        NOP                             A:B8 X:FF Y:B8 P:A5 SP:FB
	// ED30  20 DA FA  JSR $FADA                       A:B8 X:FF Y:B8 P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:B8 P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:B8 P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:B8 P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:FF Y:B8 P:27 SP:F9
	// ED33  A5 47     LDA $47 = 38                    A:B8 X:FF Y:B8 P:27 SP:FB
	// ED35  C9 38     CMP #$38                        A:38 X:FF Y:B8 P:25 SP:FB
	// ED37  F0 02     BEQ $ED3B                       A:38 X:FF Y:B8 P:27 SP:FB
	// ED3B  A9 EB     LDA #$EB                        A:38 X:FF Y:B8 P:27 SP:FB
	// ED3D  8D 47 06  STA $0647 = 38                  A:EB X:FF Y:B8 P:A5 SP:FB
	// ED40  A0 FF     LDY #$FF                        A:EB X:FF Y:B8 P:A5 SP:FB
	// ED42  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:FF P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:FF P:A5 SP:F9
	// FAB3  18        CLC                             A:EB X:FF Y:FF P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:FF Y:FF P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:FF Y:FF P:64 SP:F9
	// ED45  FB 48 05 *ISB $0548,Y @ 0647 = EB         A:40 X:FF Y:FF P:64 SP:FB
	// ED48  EA        NOP                             A:53 X:FF Y:FF P:24 SP:FB
	// ED49  EA        NOP                             A:53 X:FF Y:FF P:24 SP:FB
	// ED4A  08        PHP                             A:53 X:FF Y:FF P:24 SP:FB
	// ED4B  48        PHA                             A:53 X:FF Y:FF P:24 SP:FA
	// ED4C  A0 B9     LDY #$B9                        A:53 X:FF Y:FF P:24 SP:F9
	// ED4E  68        PLA                             A:53 X:FF Y:B9 P:A4 SP:F9
	// ED4F  28        PLP                             A:53 X:FF Y:B9 P:24 SP:FA
	// ED50  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:B9 P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:FF Y:B9 P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:FF Y:B9 P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:B9 P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:FF Y:B9 P:27 SP:F9
	// ED53  AD 47 06  LDA $0647 = EC                  A:53 X:FF Y:B9 P:27 SP:FB
	// ED56  C9 EC     CMP #$EC                        A:EC X:FF Y:B9 P:A5 SP:FB
	// ED58  F0 02     BEQ $ED5C                       A:EC X:FF Y:B9 P:27 SP:FB
	// ED5C  A0 FF     LDY #$FF                        A:EC X:FF Y:B9 P:27 SP:FB
	// ED5E  A9 FF     LDA #$FF                        A:EC X:FF Y:FF P:A5 SP:FB
	// ED60  8D 47 06  STA $0647 = EC                  A:FF X:FF Y:FF P:A5 SP:FB
	// ED63  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:FF P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:FF Y:FF P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:FF Y:FF P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:FF P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:FF Y:FF P:A5 SP:F9
	// ED66  FB 48 05 *ISB $0548,Y @ 0647 = FF         A:FF X:FF Y:FF P:A5 SP:FB
	// ED69  EA        NOP                             A:FF X:FF Y:FF P:A5 SP:FB
	// ED6A  EA        NOP                             A:FF X:FF Y:FF P:A5 SP:FB
	// ED6B  08        PHP                             A:FF X:FF Y:FF P:A5 SP:FB
	// ED6C  48        PHA                             A:FF X:FF Y:FF P:A5 SP:FA
	// ED6D  A0 BA     LDY #$BA                        A:FF X:FF Y:FF P:A5 SP:F9
	// ED6F  68        PLA                             A:FF X:FF Y:BA P:A5 SP:F9
	// ED70  28        PLP                             A:FF X:FF Y:BA P:A5 SP:FA
	// ED71  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:BA P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:FF Y:BA P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:BA P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:BA P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:FF Y:BA P:27 SP:F9
	// ED74  AD 47 06  LDA $0647 = 00                  A:FF X:FF Y:BA P:27 SP:FB
	// ED77  C9 00     CMP #$00                        A:00 X:FF Y:BA P:27 SP:FB
	// ED79  F0 02     BEQ $ED7D                       A:00 X:FF Y:BA P:27 SP:FB
	// ED7D  A0 FF     LDY #$FF                        A:00 X:FF Y:BA P:27 SP:FB
	// ED7F  A9 37     LDA #$37                        A:00 X:FF Y:FF P:A5 SP:FB
	// ED81  8D 47 06  STA $0647 = 00                  A:37 X:FF Y:FF P:25 SP:FB
	// ED84  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:FF P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:FF P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:FF Y:FF P:E5 SP:F9
	// ED87  FB 48 05 *ISB $0548,Y @ 0647 = 37         A:F0 X:FF Y:FF P:E5 SP:FB
	// ED8A  EA        NOP                             A:B8 X:FF Y:FF P:A5 SP:FB
	// ED8B  EA        NOP                             A:B8 X:FF Y:FF P:A5 SP:FB
	// ED8C  08        PHP                             A:B8 X:FF Y:FF P:A5 SP:FB
	// ED8D  48        PHA                             A:B8 X:FF Y:FF P:A5 SP:FA
	// ED8E  A0 BB     LDY #$BB                        A:B8 X:FF Y:FF P:A5 SP:F9
	// ED90  68        PLA                             A:B8 X:FF Y:BB P:A5 SP:F9
	// ED91  28        PLP                             A:B8 X:FF Y:BB P:A5 SP:FA
	// ED92  20 DA FA  JSR $FADA                       A:B8 X:FF Y:BB P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:BB P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:BB P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:BB P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:FF Y:BB P:27 SP:F9
	// ED95  AD 47 06  LDA $0647 = 38                  A:B8 X:FF Y:BB P:27 SP:FB
	// ED98  C9 38     CMP #$38                        A:38 X:FF Y:BB P:25 SP:FB
	// ED9A  F0 02     BEQ $ED9E                       A:38 X:FF Y:BB P:27 SP:FB
	// ED9E  A0 BC     LDY #$BC                        A:38 X:FF Y:BB P:27 SP:FB
	// EDA0  A2 FF     LDX #$FF                        A:38 X:FF Y:BC P:A5 SP:FB
	// EDA2  A9 EB     LDA #$EB                        A:38 X:FF Y:BC P:A5 SP:FB
	// EDA4  8D 47 06  STA $0647 = 38                  A:EB X:FF Y:BC P:A5 SP:FB
	// EDA7  20 B1 FA  JSR $FAB1                       A:EB X:FF Y:BC P:A5 SP:FB
	// FAB1  24 01     BIT $01 = FF                    A:EB X:FF Y:BC P:A5 SP:F9
	// FAB3  18        CLC                             A:EB X:FF Y:BC P:E5 SP:F9
	// FAB4  A9 40     LDA #$40                        A:EB X:FF Y:BC P:E4 SP:F9
	// FAB6  60        RTS                             A:40 X:FF Y:BC P:64 SP:F9
	// EDAA  FF 48 05 *ISB $0548,X @ 0647 = EB         A:40 X:FF Y:BC P:64 SP:FB
	// EDAD  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB
	// EDAE  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB
	// EDAF  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB
	// EDB0  EA        NOP                             A:53 X:FF Y:BC P:24 SP:FB
	// EDB1  20 B7 FA  JSR $FAB7                       A:53 X:FF Y:BC P:24 SP:FB
	// FAB7  70 2D     BVS $FAE6                       A:53 X:FF Y:BC P:24 SP:F9
	// FAB9  B0 2B     BCS $FAE6                       A:53 X:FF Y:BC P:24 SP:F9
	// FABB  30 29     BMI $FAE6                       A:53 X:FF Y:BC P:24 SP:F9
	// FABD  C9 53     CMP #$53                        A:53 X:FF Y:BC P:24 SP:F9
	// FABF  D0 25     BNE $FAE6                       A:53 X:FF Y:BC P:27 SP:F9
	// FAC1  60        RTS                             A:53 X:FF Y:BC P:27 SP:F9
	// EDB4  AD 47 06  LDA $0647 = EC                  A:53 X:FF Y:BC P:27 SP:FB
	// EDB7  C9 EC     CMP #$EC                        A:EC X:FF Y:BC P:A5 SP:FB
	// EDB9  F0 02     BEQ $EDBD                       A:EC X:FF Y:BC P:27 SP:FB
	// EDBD  C8        INY                             A:EC X:FF Y:BC P:27 SP:FB
	// EDBE  A9 FF     LDA #$FF                        A:EC X:FF Y:BD P:A5 SP:FB
	// EDC0  8D 47 06  STA $0647 = EC                  A:FF X:FF Y:BD P:A5 SP:FB
	// EDC3  20 C2 FA  JSR $FAC2                       A:FF X:FF Y:BD P:A5 SP:FB
	// FAC2  B8        CLV                             A:FF X:FF Y:BD P:A5 SP:F9
	// FAC3  38        SEC                             A:FF X:FF Y:BD P:A5 SP:F9
	// FAC4  A9 FF     LDA #$FF                        A:FF X:FF Y:BD P:A5 SP:F9
	// FAC6  60        RTS                             A:FF X:FF Y:BD P:A5 SP:F9
	// EDC6  FF 48 05 *ISB $0548,X @ 0647 = FF         A:FF X:FF Y:BD P:A5 SP:FB
	// EDC9  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB
	// EDCA  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB
	// EDCB  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB
	// EDCC  EA        NOP                             A:FF X:FF Y:BD P:A5 SP:FB
	// EDCD  20 C7 FA  JSR $FAC7                       A:FF X:FF Y:BD P:A5 SP:FB
	// FAC7  70 1D     BVS $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9
	// FAC9  F0 1B     BEQ $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9
	// FACB  10 19     BPL $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9
	// FACD  90 17     BCC $FAE6                       A:FF X:FF Y:BD P:A5 SP:F9
	// FACF  C9 FF     CMP #$FF                        A:FF X:FF Y:BD P:A5 SP:F9
	// FAD1  D0 13     BNE $FAE6                       A:FF X:FF Y:BD P:27 SP:F9
	// FAD3  60        RTS                             A:FF X:FF Y:BD P:27 SP:F9
	// EDD0  AD 47 06  LDA $0647 = 00                  A:FF X:FF Y:BD P:27 SP:FB
	// EDD3  C9 00     CMP #$00                        A:00 X:FF Y:BD P:27 SP:FB
	// EDD5  F0 02     BEQ $EDD9                       A:00 X:FF Y:BD P:27 SP:FB
	// EDD9  C8        INY                             A:00 X:FF Y:BD P:27 SP:FB
	// EDDA  A9 37     LDA #$37                        A:00 X:FF Y:BE P:A5 SP:FB
	// EDDC  8D 47 06  STA $0647 = 00                  A:37 X:FF Y:BE P:25 SP:FB
	// EDDF  20 D4 FA  JSR $FAD4                       A:37 X:FF Y:BE P:25 SP:FB
	// FAD4  24 01     BIT $01 = FF                    A:37 X:FF Y:BE P:25 SP:F9
	// FAD6  38        SEC                             A:37 X:FF Y:BE P:E5 SP:F9
	// FAD7  A9 F0     LDA #$F0                        A:37 X:FF Y:BE P:E5 SP:F9
	// FAD9  60        RTS                             A:F0 X:FF Y:BE P:E5 SP:F9
	// EDE2  FF 48 05 *ISB $0548,X @ 0647 = 37         A:F0 X:FF Y:BE P:E5 SP:FB
	// EDE5  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB
	// EDE6  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB
	// EDE7  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB
	// EDE8  EA        NOP                             A:B8 X:FF Y:BE P:A5 SP:FB
	// EDE9  20 DA FA  JSR $FADA                       A:B8 X:FF Y:BE P:A5 SP:FB
	// FADA  70 0A     BVS $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9
	// FADC  F0 08     BEQ $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9
	// FADE  10 06     BPL $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9
	// FAE0  90 04     BCC $FAE6                       A:B8 X:FF Y:BE P:A5 SP:F9
	// FAE2  C9 B8     CMP #$B8                        A:B8 X:FF Y:BE P:A5 SP:F9
	// FAE4  F0 02     BEQ $FAE8                       A:B8 X:FF Y:BE P:27 SP:F9
	// FAE8  60        RTS                             A:B8 X:FF Y:BE P:27 SP:F9
	// EDEC  AD 47 06  LDA $0647 = 38                  A:B8 X:FF Y:BE P:27 SP:FB
	// EDEF  C9 38     CMP #$38                        A:38 X:FF Y:BE P:25 SP:FB
	// EDF1  F0 02     BEQ $EDF5                       A:38 X:FF Y:BE P:27 SP:FB
	// EDF5  60        RTS                             A:38 X:FF Y:BE P:27 SP:FB
	// C641  20 F6 ED  JSR $EDF6                       A:38 X:FF Y:BE P:27 SP:FD
	// EDF6  A9 FF     LDA #$FF                        A:38 X:FF Y:BE P:27 SP:FB
	// EDF8  85 01     STA $01 = FF                    A:FF X:FF Y:BE P:A5 SP:FB
	// EDFA  A0 BF     LDY #$BF                        A:FF X:FF Y:BE P:A5 SP:FB
	// EDFC  A2 02     LDX #$02                        A:FF X:FF Y:BF P:A5 SP:FB
	// EDFE  A9 47     LDA #$47                        A:FF X:02 Y:BF P:25 SP:FB
	// EE00  85 47     STA $47 = 38                    A:47 X:02 Y:BF P:25 SP:FB
	// EE02  A9 06     LDA #$06                        A:47 X:02 Y:BF P:25 SP:FB
	// EE04  85 48     STA $48 = 06                    A:06 X:02 Y:BF P:25 SP:FB
	// EE06  A9 A5     LDA #$A5                        A:06 X:02 Y:BF P:25 SP:FB
	// EE08  8D 47 06  STA $0647 = 38                  A:A5 X:02 Y:BF P:A5 SP:FB
	// EE0B  20 7B FA  JSR $FA7B                       A:A5 X:02 Y:BF P:A5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:02 Y:BF P:A5 SP:F9
	// FA7D  18        CLC                             A:A5 X:02 Y:BF P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:02 Y:BF P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:02 Y:BF P:E4 SP:F9
	// EE0E  03 45    *SLO ($45,X) @ 47 = 0647 = A5    A:B3 X:02 Y:BF P:E4 SP:FB
	// EE10  EA        NOP                             A:FB X:02 Y:BF P:E5 SP:FB
	// EE11  EA        NOP                             A:FB X:02 Y:BF P:E5 SP:FB
	// EE12  EA        NOP                             A:FB X:02 Y:BF P:E5 SP:FB
	// EE13  EA        NOP                             A:FB X:02 Y:BF P:E5 SP:FB
	// EE14  20 81 FA  JSR $FA81                       A:FB X:02 Y:BF P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:02 Y:BF P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:02 Y:BF P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:02 Y:BF P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:02 Y:BF P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:02 Y:BF P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:02 Y:BF P:67 SP:F9
	// EE17  AD 47 06  LDA $0647 = 4A                  A:FB X:02 Y:BF P:67 SP:FB
	// EE1A  C9 4A     CMP #$4A                        A:4A X:02 Y:BF P:65 SP:FB
	// EE1C  F0 02     BEQ $EE20                       A:4A X:02 Y:BF P:67 SP:FB
	// EE20  C8        INY                             A:4A X:02 Y:BF P:67 SP:FB
	// EE21  A9 29     LDA #$29                        A:4A X:02 Y:C0 P:E5 SP:FB
	// EE23  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:C0 P:65 SP:FB
	// EE26  20 8C FA  JSR $FA8C                       A:29 X:02 Y:C0 P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:02 Y:C0 P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:02 Y:C0 P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:02 Y:C0 P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:02 Y:C0 P:A4 SP:F9
	// EE29  03 45    *SLO ($45,X) @ 47 = 0647 = 29    A:C3 X:02 Y:C0 P:A4 SP:FB
	// EE2B  EA        NOP                             A:D3 X:02 Y:C0 P:A4 SP:FB
	// EE2C  EA        NOP                             A:D3 X:02 Y:C0 P:A4 SP:FB
	// EE2D  EA        NOP                             A:D3 X:02 Y:C0 P:A4 SP:FB
	// EE2E  EA        NOP                             A:D3 X:02 Y:C0 P:A4 SP:FB
	// EE2F  20 91 FA  JSR $FA91                       A:D3 X:02 Y:C0 P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:02 Y:C0 P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:02 Y:C0 P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:02 Y:C0 P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:02 Y:C0 P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:02 Y:C0 P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:02 Y:C0 P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:02 Y:C0 P:27 SP:F9
	// EE32  AD 47 06  LDA $0647 = 52                  A:D3 X:02 Y:C0 P:27 SP:FB
	// EE35  C9 52     CMP #$52                        A:52 X:02 Y:C0 P:25 SP:FB
	// EE37  F0 02     BEQ $EE3B                       A:52 X:02 Y:C0 P:27 SP:FB
	// EE3B  C8        INY                             A:52 X:02 Y:C0 P:27 SP:FB
	// EE3C  A9 37     LDA #$37                        A:52 X:02 Y:C1 P:A5 SP:FB
	// EE3E  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:C1 P:25 SP:FB
	// EE41  20 9E FA  JSR $FA9E                       A:37 X:02 Y:C1 P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:02 Y:C1 P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:02 Y:C1 P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:02 Y:C1 P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:02 Y:C1 P:65 SP:F9
	// EE44  03 45    *SLO ($45,X) @ 47 = 0647 = 37    A:10 X:02 Y:C1 P:65 SP:FB
	// EE46  EA        NOP                             A:7E X:02 Y:C1 P:64 SP:FB
	// EE47  EA        NOP                             A:7E X:02 Y:C1 P:64 SP:FB
	// EE48  EA        NOP                             A:7E X:02 Y:C1 P:64 SP:FB
	// EE49  EA        NOP                             A:7E X:02 Y:C1 P:64 SP:FB
	// EE4A  20 A4 FA  JSR $FAA4                       A:7E X:02 Y:C1 P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:02 Y:C1 P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:02 Y:C1 P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:02 Y:C1 P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:02 Y:C1 P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:02 Y:C1 P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:02 Y:C1 P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:02 Y:C1 P:67 SP:F9
	// EE4D  AD 47 06  LDA $0647 = 6E                  A:7E X:02 Y:C1 P:67 SP:FB
	// EE50  C9 6E     CMP #$6E                        A:6E X:02 Y:C1 P:65 SP:FB
	// EE52  F0 02     BEQ $EE56                       A:6E X:02 Y:C1 P:67 SP:FB
	// EE56  C8        INY                             A:6E X:02 Y:C1 P:67 SP:FB
	// EE57  A9 A5     LDA #$A5                        A:6E X:02 Y:C2 P:E5 SP:FB
	// EE59  85 47     STA $47 = 47                    A:A5 X:02 Y:C2 P:E5 SP:FB
	// EE5B  20 7B FA  JSR $FA7B                       A:A5 X:02 Y:C2 P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:02 Y:C2 P:E5 SP:F9
	// FA7D  18        CLC                             A:A5 X:02 Y:C2 P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:02 Y:C2 P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:02 Y:C2 P:E4 SP:F9
	// EE5E  07 47    *SLO $47 = A5                    A:B3 X:02 Y:C2 P:E4 SP:FB
	// EE60  EA        NOP                             A:FB X:02 Y:C2 P:E5 SP:FB
	// EE61  EA        NOP                             A:FB X:02 Y:C2 P:E5 SP:FB
	// EE62  EA        NOP                             A:FB X:02 Y:C2 P:E5 SP:FB
	// EE63  EA        NOP                             A:FB X:02 Y:C2 P:E5 SP:FB
	// EE64  20 81 FA  JSR $FA81                       A:FB X:02 Y:C2 P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:02 Y:C2 P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:02 Y:C2 P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:02 Y:C2 P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:02 Y:C2 P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:02 Y:C2 P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:02 Y:C2 P:67 SP:F9
	// EE67  A5 47     LDA $47 = 4A                    A:FB X:02 Y:C2 P:67 SP:FB
	// EE69  C9 4A     CMP #$4A                        A:4A X:02 Y:C2 P:65 SP:FB
	// EE6B  F0 02     BEQ $EE6F                       A:4A X:02 Y:C2 P:67 SP:FB
	// EE6F  C8        INY                             A:4A X:02 Y:C2 P:67 SP:FB
	// EE70  A9 29     LDA #$29                        A:4A X:02 Y:C3 P:E5 SP:FB
	// EE72  85 47     STA $47 = 4A                    A:29 X:02 Y:C3 P:65 SP:FB
	// EE74  20 8C FA  JSR $FA8C                       A:29 X:02 Y:C3 P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:02 Y:C3 P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:02 Y:C3 P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:02 Y:C3 P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:02 Y:C3 P:A4 SP:F9
	// EE77  07 47    *SLO $47 = 29                    A:C3 X:02 Y:C3 P:A4 SP:FB
	// EE79  EA        NOP                             A:D3 X:02 Y:C3 P:A4 SP:FB
	// EE7A  EA        NOP                             A:D3 X:02 Y:C3 P:A4 SP:FB
	// EE7B  EA        NOP                             A:D3 X:02 Y:C3 P:A4 SP:FB
	// EE7C  EA        NOP                             A:D3 X:02 Y:C3 P:A4 SP:FB
	// EE7D  20 91 FA  JSR $FA91                       A:D3 X:02 Y:C3 P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:02 Y:C3 P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:02 Y:C3 P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:02 Y:C3 P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:02 Y:C3 P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:02 Y:C3 P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:02 Y:C3 P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:02 Y:C3 P:27 SP:F9
	// EE80  A5 47     LDA $47 = 52                    A:D3 X:02 Y:C3 P:27 SP:FB
	// EE82  C9 52     CMP #$52                        A:52 X:02 Y:C3 P:25 SP:FB
	// EE84  F0 02     BEQ $EE88                       A:52 X:02 Y:C3 P:27 SP:FB
	// EE88  C8        INY                             A:52 X:02 Y:C3 P:27 SP:FB
	// EE89  A9 37     LDA #$37                        A:52 X:02 Y:C4 P:A5 SP:FB
	// EE8B  85 47     STA $47 = 52                    A:37 X:02 Y:C4 P:25 SP:FB
	// EE8D  20 9E FA  JSR $FA9E                       A:37 X:02 Y:C4 P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:02 Y:C4 P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:02 Y:C4 P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:02 Y:C4 P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:02 Y:C4 P:65 SP:F9
	// EE90  07 47    *SLO $47 = 37                    A:10 X:02 Y:C4 P:65 SP:FB
	// EE92  EA        NOP                             A:7E X:02 Y:C4 P:64 SP:FB
	// EE93  EA        NOP                             A:7E X:02 Y:C4 P:64 SP:FB
	// EE94  EA        NOP                             A:7E X:02 Y:C4 P:64 SP:FB
	// EE95  EA        NOP                             A:7E X:02 Y:C4 P:64 SP:FB
	// EE96  20 A4 FA  JSR $FAA4                       A:7E X:02 Y:C4 P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:02 Y:C4 P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:02 Y:C4 P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:02 Y:C4 P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:02 Y:C4 P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:02 Y:C4 P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:02 Y:C4 P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:02 Y:C4 P:67 SP:F9
	// EE99  A5 47     LDA $47 = 6E                    A:7E X:02 Y:C4 P:67 SP:FB
	// EE9B  C9 6E     CMP #$6E                        A:6E X:02 Y:C4 P:65 SP:FB
	// EE9D  F0 02     BEQ $EEA1                       A:6E X:02 Y:C4 P:67 SP:FB
	// EEA1  C8        INY                             A:6E X:02 Y:C4 P:67 SP:FB
	// EEA2  A9 A5     LDA #$A5                        A:6E X:02 Y:C5 P:E5 SP:FB
	// EEA4  8D 47 06  STA $0647 = 6E                  A:A5 X:02 Y:C5 P:E5 SP:FB
	// EEA7  20 7B FA  JSR $FA7B                       A:A5 X:02 Y:C5 P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:02 Y:C5 P:E5 SP:F9
	// FA7D  18        CLC                             A:A5 X:02 Y:C5 P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:02 Y:C5 P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:02 Y:C5 P:E4 SP:F9
	// EEAA  0F 47 06 *SLO $0647 = A5                  A:B3 X:02 Y:C5 P:E4 SP:FB
	// EEAD  EA        NOP                             A:FB X:02 Y:C5 P:E5 SP:FB
	// EEAE  EA        NOP                             A:FB X:02 Y:C5 P:E5 SP:FB
	// EEAF  EA        NOP                             A:FB X:02 Y:C5 P:E5 SP:FB
	// EEB0  EA        NOP                             A:FB X:02 Y:C5 P:E5 SP:FB
	// EEB1  20 81 FA  JSR $FA81                       A:FB X:02 Y:C5 P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:02 Y:C5 P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:02 Y:C5 P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:02 Y:C5 P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:02 Y:C5 P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:02 Y:C5 P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:02 Y:C5 P:67 SP:F9
	// EEB4  AD 47 06  LDA $0647 = 4A                  A:FB X:02 Y:C5 P:67 SP:FB
	// EEB7  C9 4A     CMP #$4A                        A:4A X:02 Y:C5 P:65 SP:FB
	// EEB9  F0 02     BEQ $EEBD                       A:4A X:02 Y:C5 P:67 SP:FB
	// EEBD  C8        INY                             A:4A X:02 Y:C5 P:67 SP:FB
	// EEBE  A9 29     LDA #$29                        A:4A X:02 Y:C6 P:E5 SP:FB
	// EEC0  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:C6 P:65 SP:FB
	// EEC3  20 8C FA  JSR $FA8C                       A:29 X:02 Y:C6 P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:02 Y:C6 P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:02 Y:C6 P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:02 Y:C6 P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:02 Y:C6 P:A4 SP:F9
	// EEC6  0F 47 06 *SLO $0647 = 29                  A:C3 X:02 Y:C6 P:A4 SP:FB
	// EEC9  EA        NOP                             A:D3 X:02 Y:C6 P:A4 SP:FB
	// EECA  EA        NOP                             A:D3 X:02 Y:C6 P:A4 SP:FB
	// EECB  EA        NOP                             A:D3 X:02 Y:C6 P:A4 SP:FB
	// EECC  EA        NOP                             A:D3 X:02 Y:C6 P:A4 SP:FB
	// EECD  20 91 FA  JSR $FA91                       A:D3 X:02 Y:C6 P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:02 Y:C6 P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:02 Y:C6 P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:02 Y:C6 P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:02 Y:C6 P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:02 Y:C6 P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:02 Y:C6 P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:02 Y:C6 P:27 SP:F9
	// EED0  AD 47 06  LDA $0647 = 52                  A:D3 X:02 Y:C6 P:27 SP:FB
	// EED3  C9 52     CMP #$52                        A:52 X:02 Y:C6 P:25 SP:FB
	// EED5  F0 02     BEQ $EED9                       A:52 X:02 Y:C6 P:27 SP:FB
	// EED9  C8        INY                             A:52 X:02 Y:C6 P:27 SP:FB
	// EEDA  A9 37     LDA #$37                        A:52 X:02 Y:C7 P:A5 SP:FB
	// EEDC  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:C7 P:25 SP:FB
	// EEDF  20 9E FA  JSR $FA9E                       A:37 X:02 Y:C7 P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:02 Y:C7 P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:02 Y:C7 P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:02 Y:C7 P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:02 Y:C7 P:65 SP:F9
	// EEE2  0F 47 06 *SLO $0647 = 37                  A:10 X:02 Y:C7 P:65 SP:FB
	// EEE5  EA        NOP                             A:7E X:02 Y:C7 P:64 SP:FB
	// EEE6  EA        NOP                             A:7E X:02 Y:C7 P:64 SP:FB
	// EEE7  EA        NOP                             A:7E X:02 Y:C7 P:64 SP:FB
	// EEE8  EA        NOP                             A:7E X:02 Y:C7 P:64 SP:FB
	// EEE9  20 A4 FA  JSR $FAA4                       A:7E X:02 Y:C7 P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:02 Y:C7 P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:02 Y:C7 P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:02 Y:C7 P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:02 Y:C7 P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:02 Y:C7 P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:02 Y:C7 P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:02 Y:C7 P:67 SP:F9
	// EEEC  AD 47 06  LDA $0647 = 6E                  A:7E X:02 Y:C7 P:67 SP:FB
	// EEEF  C9 6E     CMP #$6E                        A:6E X:02 Y:C7 P:65 SP:FB
	// EEF1  F0 02     BEQ $EEF5                       A:6E X:02 Y:C7 P:67 SP:FB
	// EEF5  A9 A5     LDA #$A5                        A:6E X:02 Y:C7 P:67 SP:FB
	// EEF7  8D 47 06  STA $0647 = 6E                  A:A5 X:02 Y:C7 P:E5 SP:FB
	// EEFA  A9 48     LDA #$48                        A:A5 X:02 Y:C7 P:E5 SP:FB
	// EEFC  85 45     STA $45 = 48                    A:48 X:02 Y:C7 P:65 SP:FB
	// EEFE  A9 05     LDA #$05                        A:48 X:02 Y:C7 P:65 SP:FB
	// EF00  85 46     STA $46 = 05                    A:05 X:02 Y:C7 P:65 SP:FB
	// EF02  A0 FF     LDY #$FF                        A:05 X:02 Y:C7 P:65 SP:FB
	// EF04  20 7B FA  JSR $FA7B                       A:05 X:02 Y:FF P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:E5 SP:F9
	// FA7D  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:05 X:02 Y:FF P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:02 Y:FF P:E4 SP:F9
	// EF07  13 45    *SLO ($45),Y = 0548 @ 0647 = A5  A:B3 X:02 Y:FF P:E4 SP:FB
	// EF09  EA        NOP                             A:FB X:02 Y:FF P:E5 SP:FB
	// EF0A  EA        NOP                             A:FB X:02 Y:FF P:E5 SP:FB
	// EF0B  08        PHP                             A:FB X:02 Y:FF P:E5 SP:FB
	// EF0C  48        PHA                             A:FB X:02 Y:FF P:E5 SP:FA
	// EF0D  A0 C8     LDY #$C8                        A:FB X:02 Y:FF P:E5 SP:F9
	// EF0F  68        PLA                             A:FB X:02 Y:C8 P:E5 SP:F9
	// EF10  28        PLP                             A:FB X:02 Y:C8 P:E5 SP:FA
	// EF11  20 81 FA  JSR $FA81                       A:FB X:02 Y:C8 P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:02 Y:C8 P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:02 Y:C8 P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:02 Y:C8 P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:02 Y:C8 P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:02 Y:C8 P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:02 Y:C8 P:67 SP:F9
	// EF14  AD 47 06  LDA $0647 = 4A                  A:FB X:02 Y:C8 P:67 SP:FB
	// EF17  C9 4A     CMP #$4A                        A:4A X:02 Y:C8 P:65 SP:FB
	// EF19  F0 02     BEQ $EF1D                       A:4A X:02 Y:C8 P:67 SP:FB
	// EF1D  A0 FF     LDY #$FF                        A:4A X:02 Y:C8 P:67 SP:FB
	// EF1F  A9 29     LDA #$29                        A:4A X:02 Y:FF P:E5 SP:FB
	// EF21  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:FF P:65 SP:FB
	// EF24  20 8C FA  JSR $FA8C                       A:29 X:02 Y:FF P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:02 Y:FF P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:02 Y:FF P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:02 Y:FF P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:02 Y:FF P:A4 SP:F9
	// EF27  13 45    *SLO ($45),Y = 0548 @ 0647 = 29  A:C3 X:02 Y:FF P:A4 SP:FB
	// EF29  EA        NOP                             A:D3 X:02 Y:FF P:A4 SP:FB
	// EF2A  EA        NOP                             A:D3 X:02 Y:FF P:A4 SP:FB
	// EF2B  08        PHP                             A:D3 X:02 Y:FF P:A4 SP:FB
	// EF2C  48        PHA                             A:D3 X:02 Y:FF P:A4 SP:FA
	// EF2D  A0 C9     LDY #$C9                        A:D3 X:02 Y:FF P:A4 SP:F9
	// EF2F  68        PLA                             A:D3 X:02 Y:C9 P:A4 SP:F9
	// EF30  28        PLP                             A:D3 X:02 Y:C9 P:A4 SP:FA
	// EF31  20 91 FA  JSR $FA91                       A:D3 X:02 Y:C9 P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:02 Y:C9 P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:02 Y:C9 P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:02 Y:C9 P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:02 Y:C9 P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:02 Y:C9 P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:02 Y:C9 P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:02 Y:C9 P:27 SP:F9
	// EF34  AD 47 06  LDA $0647 = 52                  A:D3 X:02 Y:C9 P:27 SP:FB
	// EF37  C9 52     CMP #$52                        A:52 X:02 Y:C9 P:25 SP:FB
	// EF39  F0 02     BEQ $EF3D                       A:52 X:02 Y:C9 P:27 SP:FB
	// EF3D  A0 FF     LDY #$FF                        A:52 X:02 Y:C9 P:27 SP:FB
	// EF3F  A9 37     LDA #$37                        A:52 X:02 Y:FF P:A5 SP:FB
	// EF41  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:FF P:25 SP:FB
	// EF44  20 9E FA  JSR $FA9E                       A:37 X:02 Y:FF P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:02 Y:FF P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:02 Y:FF P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:02 Y:FF P:65 SP:F9
	// EF47  13 45    *SLO ($45),Y = 0548 @ 0647 = 37  A:10 X:02 Y:FF P:65 SP:FB
	// EF49  EA        NOP                             A:7E X:02 Y:FF P:64 SP:FB
	// EF4A  EA        NOP                             A:7E X:02 Y:FF P:64 SP:FB
	// EF4B  08        PHP                             A:7E X:02 Y:FF P:64 SP:FB
	// EF4C  48        PHA                             A:7E X:02 Y:FF P:64 SP:FA
	// EF4D  A0 CA     LDY #$CA                        A:7E X:02 Y:FF P:64 SP:F9
	// EF4F  68        PLA                             A:7E X:02 Y:CA P:E4 SP:F9
	// EF50  28        PLP                             A:7E X:02 Y:CA P:64 SP:FA
	// EF51  20 A4 FA  JSR $FAA4                       A:7E X:02 Y:CA P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:02 Y:CA P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:02 Y:CA P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:02 Y:CA P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:02 Y:CA P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:02 Y:CA P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:02 Y:CA P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:02 Y:CA P:67 SP:F9
	// EF54  AD 47 06  LDA $0647 = 6E                  A:7E X:02 Y:CA P:67 SP:FB
	// EF57  C9 6E     CMP #$6E                        A:6E X:02 Y:CA P:65 SP:FB
	// EF59  F0 02     BEQ $EF5D                       A:6E X:02 Y:CA P:67 SP:FB
	// EF5D  A0 CB     LDY #$CB                        A:6E X:02 Y:CA P:67 SP:FB
	// EF5F  A2 FF     LDX #$FF                        A:6E X:02 Y:CB P:E5 SP:FB
	// EF61  A9 A5     LDA #$A5                        A:6E X:FF Y:CB P:E5 SP:FB
	// EF63  85 47     STA $47 = 6E                    A:A5 X:FF Y:CB P:E5 SP:FB
	// EF65  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:CB P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:CB P:E5 SP:F9
	// FA7D  18        CLC                             A:A5 X:FF Y:CB P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:CB P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:FF Y:CB P:E4 SP:F9
	// EF68  17 48    *SLO $48,X @ 47 = A5             A:B3 X:FF Y:CB P:E4 SP:FB
	// EF6A  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB
	// EF6B  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB
	// EF6C  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB
	// EF6D  EA        NOP                             A:FB X:FF Y:CB P:E5 SP:FB
	// EF6E  20 81 FA  JSR $FA81                       A:FB X:FF Y:CB P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:CB P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:CB P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:CB P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:FF Y:CB P:67 SP:F9
	// EF71  A5 47     LDA $47 = 4A                    A:FB X:FF Y:CB P:67 SP:FB
	// EF73  C9 4A     CMP #$4A                        A:4A X:FF Y:CB P:65 SP:FB
	// EF75  F0 02     BEQ $EF79                       A:4A X:FF Y:CB P:67 SP:FB
	// EF79  C8        INY                             A:4A X:FF Y:CB P:67 SP:FB
	// EF7A  A9 29     LDA #$29                        A:4A X:FF Y:CC P:E5 SP:FB
	// EF7C  85 47     STA $47 = 4A                    A:29 X:FF Y:CC P:65 SP:FB
	// EF7E  20 8C FA  JSR $FA8C                       A:29 X:FF Y:CC P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:FF Y:CC P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:FF Y:CC P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:CC P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:FF Y:CC P:A4 SP:F9
	// EF81  17 48    *SLO $48,X @ 47 = 29             A:C3 X:FF Y:CC P:A4 SP:FB
	// EF83  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB
	// EF84  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB
	// EF85  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB
	// EF86  EA        NOP                             A:D3 X:FF Y:CC P:A4 SP:FB
	// EF87  20 91 FA  JSR $FA91                       A:D3 X:FF Y:CC P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:CC P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:CC P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:CC P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:FF Y:CC P:27 SP:F9
	// EF8A  A5 47     LDA $47 = 52                    A:D3 X:FF Y:CC P:27 SP:FB
	// EF8C  C9 52     CMP #$52                        A:52 X:FF Y:CC P:25 SP:FB
	// EF8E  F0 02     BEQ $EF92                       A:52 X:FF Y:CC P:27 SP:FB
	// EF92  C8        INY                             A:52 X:FF Y:CC P:27 SP:FB
	// EF93  A9 37     LDA #$37                        A:52 X:FF Y:CD P:A5 SP:FB
	// EF95  85 47     STA $47 = 52                    A:37 X:FF Y:CD P:25 SP:FB
	// EF97  20 9E FA  JSR $FA9E                       A:37 X:FF Y:CD P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:CD P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:FF Y:CD P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:FF Y:CD P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:FF Y:CD P:65 SP:F9
	// EF9A  17 48    *SLO $48,X @ 47 = 37             A:10 X:FF Y:CD P:65 SP:FB
	// EF9C  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB
	// EF9D  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB
	// EF9E  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB
	// EF9F  EA        NOP                             A:7E X:FF Y:CD P:64 SP:FB
	// EFA0  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:CD P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:CD P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:CD P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:CD P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:CD P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:CD P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:CD P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:FF Y:CD P:67 SP:F9
	// EFA3  A5 47     LDA $47 = 6E                    A:7E X:FF Y:CD P:67 SP:FB
	// EFA5  C9 6E     CMP #$6E                        A:6E X:FF Y:CD P:65 SP:FB
	// EFA7  F0 02     BEQ $EFAB                       A:6E X:FF Y:CD P:67 SP:FB
	// EFAB  A9 A5     LDA #$A5                        A:6E X:FF Y:CD P:67 SP:FB
	// EFAD  8D 47 06  STA $0647 = 6E                  A:A5 X:FF Y:CD P:E5 SP:FB
	// EFB0  A0 FF     LDY #$FF                        A:A5 X:FF Y:CD P:E5 SP:FB
	// EFB2  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:FF P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9
	// FA7D  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9
	// EFB5  1B 48 05 *SLO $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB
	// EFB8  EA        NOP                             A:FB X:FF Y:FF P:E5 SP:FB
	// EFB9  EA        NOP                             A:FB X:FF Y:FF P:E5 SP:FB
	// EFBA  08        PHP                             A:FB X:FF Y:FF P:E5 SP:FB
	// EFBB  48        PHA                             A:FB X:FF Y:FF P:E5 SP:FA
	// EFBC  A0 CE     LDY #$CE                        A:FB X:FF Y:FF P:E5 SP:F9
	// EFBE  68        PLA                             A:FB X:FF Y:CE P:E5 SP:F9
	// EFBF  28        PLP                             A:FB X:FF Y:CE P:E5 SP:FA
	// EFC0  20 81 FA  JSR $FA81                       A:FB X:FF Y:CE P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:CE P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:CE P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:CE P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:FF Y:CE P:67 SP:F9
	// EFC3  AD 47 06  LDA $0647 = 4A                  A:FB X:FF Y:CE P:67 SP:FB
	// EFC6  C9 4A     CMP #$4A                        A:4A X:FF Y:CE P:65 SP:FB
	// EFC8  F0 02     BEQ $EFCC                       A:4A X:FF Y:CE P:67 SP:FB
	// EFCC  A0 FF     LDY #$FF                        A:4A X:FF Y:CE P:67 SP:FB
	// EFCE  A9 29     LDA #$29                        A:4A X:FF Y:FF P:E5 SP:FB
	// EFD0  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:FF P:65 SP:FB
	// EFD3  20 8C FA  JSR $FA8C                       A:29 X:FF Y:FF P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:FF P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:FF Y:FF P:A4 SP:F9
	// EFD6  1B 48 05 *SLO $0548,Y @ 0647 = 29         A:C3 X:FF Y:FF P:A4 SP:FB
	// EFD9  EA        NOP                             A:D3 X:FF Y:FF P:A4 SP:FB
	// EFDA  EA        NOP                             A:D3 X:FF Y:FF P:A4 SP:FB
	// EFDB  08        PHP                             A:D3 X:FF Y:FF P:A4 SP:FB
	// EFDC  48        PHA                             A:D3 X:FF Y:FF P:A4 SP:FA
	// EFDD  A0 CF     LDY #$CF                        A:D3 X:FF Y:FF P:A4 SP:F9
	// EFDF  68        PLA                             A:D3 X:FF Y:CF P:A4 SP:F9
	// EFE0  28        PLP                             A:D3 X:FF Y:CF P:A4 SP:FA
	// EFE1  20 91 FA  JSR $FA91                       A:D3 X:FF Y:CF P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:CF P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:CF P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:CF P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:FF Y:CF P:27 SP:F9
	// EFE4  AD 47 06  LDA $0647 = 52                  A:D3 X:FF Y:CF P:27 SP:FB
	// EFE7  C9 52     CMP #$52                        A:52 X:FF Y:CF P:25 SP:FB
	// EFE9  F0 02     BEQ $EFED                       A:52 X:FF Y:CF P:27 SP:FB
	// EFED  A0 FF     LDY #$FF                        A:52 X:FF Y:CF P:27 SP:FB
	// EFEF  A9 37     LDA #$37                        A:52 X:FF Y:FF P:A5 SP:FB
	// EFF1  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:FF P:25 SP:FB
	// EFF4  20 9E FA  JSR $FA9E                       A:37 X:FF Y:FF P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:FF Y:FF P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:FF Y:FF P:65 SP:F9
	// EFF7  1B 48 05 *SLO $0548,Y @ 0647 = 37         A:10 X:FF Y:FF P:65 SP:FB
	// EFFA  EA        NOP                             A:7E X:FF Y:FF P:64 SP:FB
	// EFFB  EA        NOP                             A:7E X:FF Y:FF P:64 SP:FB
	// EFFC  08        PHP                             A:7E X:FF Y:FF P:64 SP:FB
	// EFFD  48        PHA                             A:7E X:FF Y:FF P:64 SP:FA
	// EFFE  A0 D0     LDY #$D0                        A:7E X:FF Y:FF P:64 SP:F9
	// F000  68        PLA                             A:7E X:FF Y:D0 P:E4 SP:F9
	// F001  28        PLP                             A:7E X:FF Y:D0 P:64 SP:FA
	// F002  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:D0 P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:D0 P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:D0 P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:D0 P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:FF Y:D0 P:67 SP:F9
	// F005  AD 47 06  LDA $0647 = 6E                  A:7E X:FF Y:D0 P:67 SP:FB
	// F008  C9 6E     CMP #$6E                        A:6E X:FF Y:D0 P:65 SP:FB
	// F00A  F0 02     BEQ $F00E                       A:6E X:FF Y:D0 P:67 SP:FB
	// F00E  A0 D1     LDY #$D1                        A:6E X:FF Y:D0 P:67 SP:FB
	// F010  A2 FF     LDX #$FF                        A:6E X:FF Y:D1 P:E5 SP:FB
	// F012  A9 A5     LDA #$A5                        A:6E X:FF Y:D1 P:E5 SP:FB
	// F014  8D 47 06  STA $0647 = 6E                  A:A5 X:FF Y:D1 P:E5 SP:FB
	// F017  20 7B FA  JSR $FA7B                       A:A5 X:FF Y:D1 P:E5 SP:FB
	// FA7B  24 01     BIT $01 = FF                    A:A5 X:FF Y:D1 P:E5 SP:F9
	// FA7D  18        CLC                             A:A5 X:FF Y:D1 P:E5 SP:F9
	// FA7E  A9 B3     LDA #$B3                        A:A5 X:FF Y:D1 P:E4 SP:F9
	// FA80  60        RTS                             A:B3 X:FF Y:D1 P:E4 SP:F9
	// F01A  1F 48 05 *SLO $0548,X @ 0647 = A5         A:B3 X:FF Y:D1 P:E4 SP:FB
	// F01D  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB
	// F01E  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB
	// F01F  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB
	// F020  EA        NOP                             A:FB X:FF Y:D1 P:E5 SP:FB
	// F021  20 81 FA  JSR $FA81                       A:FB X:FF Y:D1 P:E5 SP:FB
	// FA81  50 63     BVC $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9
	// FA83  90 61     BCC $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9
	// FA85  10 5F     BPL $FAE6                       A:FB X:FF Y:D1 P:E5 SP:F9
	// FA87  C9 FB     CMP #$FB                        A:FB X:FF Y:D1 P:E5 SP:F9
	// FA89  D0 5B     BNE $FAE6                       A:FB X:FF Y:D1 P:67 SP:F9
	// FA8B  60        RTS                             A:FB X:FF Y:D1 P:67 SP:F9
	// F024  AD 47 06  LDA $0647 = 4A                  A:FB X:FF Y:D1 P:67 SP:FB
	// F027  C9 4A     CMP #$4A                        A:4A X:FF Y:D1 P:65 SP:FB
	// F029  F0 02     BEQ $F02D                       A:4A X:FF Y:D1 P:67 SP:FB
	// F02D  C8        INY                             A:4A X:FF Y:D1 P:67 SP:FB
	// F02E  A9 29     LDA #$29                        A:4A X:FF Y:D2 P:E5 SP:FB
	// F030  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:D2 P:65 SP:FB
	// F033  20 8C FA  JSR $FA8C                       A:29 X:FF Y:D2 P:65 SP:FB
	// FA8C  B8        CLV                             A:29 X:FF Y:D2 P:65 SP:F9
	// FA8D  18        CLC                             A:29 X:FF Y:D2 P:25 SP:F9
	// FA8E  A9 C3     LDA #$C3                        A:29 X:FF Y:D2 P:24 SP:F9
	// FA90  60        RTS                             A:C3 X:FF Y:D2 P:A4 SP:F9
	// F036  1F 48 05 *SLO $0548,X @ 0647 = 29         A:C3 X:FF Y:D2 P:A4 SP:FB
	// F039  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB
	// F03A  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB
	// F03B  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB
	// F03C  EA        NOP                             A:D3 X:FF Y:D2 P:A4 SP:FB
	// F03D  20 91 FA  JSR $FA91                       A:D3 X:FF Y:D2 P:A4 SP:FB
	// FA91  70 53     BVS $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9
	// FA93  F0 51     BEQ $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9
	// FA95  10 4F     BPL $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9
	// FA97  B0 4D     BCS $FAE6                       A:D3 X:FF Y:D2 P:A4 SP:F9
	// FA99  C9 D3     CMP #$D3                        A:D3 X:FF Y:D2 P:A4 SP:F9
	// FA9B  D0 49     BNE $FAE6                       A:D3 X:FF Y:D2 P:27 SP:F9
	// FA9D  60        RTS                             A:D3 X:FF Y:D2 P:27 SP:F9
	// F040  AD 47 06  LDA $0647 = 52                  A:D3 X:FF Y:D2 P:27 SP:FB
	// F043  C9 52     CMP #$52                        A:52 X:FF Y:D2 P:25 SP:FB
	// F045  F0 02     BEQ $F049                       A:52 X:FF Y:D2 P:27 SP:FB
	// F049  C8        INY                             A:52 X:FF Y:D2 P:27 SP:FB
	// F04A  A9 37     LDA #$37                        A:52 X:FF Y:D3 P:A5 SP:FB
	// F04C  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:D3 P:25 SP:FB
	// F04F  20 9E FA  JSR $FA9E                       A:37 X:FF Y:D3 P:25 SP:FB
	// FA9E  24 01     BIT $01 = FF                    A:37 X:FF Y:D3 P:25 SP:F9
	// FAA0  38        SEC                             A:37 X:FF Y:D3 P:E5 SP:F9
	// FAA1  A9 10     LDA #$10                        A:37 X:FF Y:D3 P:E5 SP:F9
	// FAA3  60        RTS                             A:10 X:FF Y:D3 P:65 SP:F9
	// F052  1F 48 05 *SLO $0548,X @ 0647 = 37         A:10 X:FF Y:D3 P:65 SP:FB
	// F055  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB
	// F056  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB
	// F057  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB
	// F058  EA        NOP                             A:7E X:FF Y:D3 P:64 SP:FB
	// F059  20 A4 FA  JSR $FAA4                       A:7E X:FF Y:D3 P:64 SP:FB
	// FAA4  50 40     BVC $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9
	// FAA6  F0 3E     BEQ $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9
	// FAA8  30 3C     BMI $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9
	// FAAA  B0 3A     BCS $FAE6                       A:7E X:FF Y:D3 P:64 SP:F9
	// FAAC  C9 7E     CMP #$7E                        A:7E X:FF Y:D3 P:64 SP:F9
	// FAAE  D0 36     BNE $FAE6                       A:7E X:FF Y:D3 P:67 SP:F9
	// FAB0  60        RTS                             A:7E X:FF Y:D3 P:67 SP:F9
	// F05C  AD 47 06  LDA $0647 = 6E                  A:7E X:FF Y:D3 P:67 SP:FB
	// F05F  C9 6E     CMP #$6E                        A:6E X:FF Y:D3 P:65 SP:FB
	// F061  F0 02     BEQ $F065                       A:6E X:FF Y:D3 P:67 SP:FB
	// F065  60        RTS                             A:6E X:FF Y:D3 P:67 SP:FB
	// C644  20 66 F0  JSR $F066                       A:6E X:FF Y:D3 P:67 SP:FD
	// F066  A9 FF     LDA #$FF                        A:6E X:FF Y:D3 P:67 SP:FB
	// F068  85 01     STA $01 = FF                    A:FF X:FF Y:D3 P:E5 SP:FB
	// F06A  A0 D4     LDY #$D4                        A:FF X:FF Y:D3 P:E5 SP:FB
	// F06C  A2 02     LDX #$02                        A:FF X:FF Y:D4 P:E5 SP:FB
	// F06E  A9 47     LDA #$47                        A:FF X:02 Y:D4 P:65 SP:FB
	// F070  85 47     STA $47 = 6E                    A:47 X:02 Y:D4 P:65 SP:FB
	// F072  A9 06     LDA #$06                        A:47 X:02 Y:D4 P:65 SP:FB
	// F074  85 48     STA $48 = 06                    A:06 X:02 Y:D4 P:65 SP:FB
	// F076  A9 A5     LDA #$A5                        A:06 X:02 Y:D4 P:65 SP:FB
	// F078  8D 47 06  STA $0647 = 6E                  A:A5 X:02 Y:D4 P:E5 SP:FB
	// F07B  20 53 FB  JSR $FB53                       A:A5 X:02 Y:D4 P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:02 Y:D4 P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:02 Y:D4 P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:02 Y:D4 P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:02 Y:D4 P:E4 SP:F9
	// F07E  23 45    *RLA ($45,X) @ 47 = 0647 = A5    A:B3 X:02 Y:D4 P:E4 SP:FB
	// F080  EA        NOP                             A:02 X:02 Y:D4 P:65 SP:FB
	// F081  EA        NOP                             A:02 X:02 Y:D4 P:65 SP:FB
	// F082  EA        NOP                             A:02 X:02 Y:D4 P:65 SP:FB
	// F083  EA        NOP                             A:02 X:02 Y:D4 P:65 SP:FB
	// F084  20 59 FB  JSR $FB59                       A:02 X:02 Y:D4 P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:02 Y:D4 P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:02 Y:D4 P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:02 Y:D4 P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:02 Y:D4 P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:02 Y:D4 P:67 SP:F9
	// FB63  60        RTS                             A:02 X:02 Y:D4 P:67 SP:F9
	// F087  AD 47 06  LDA $0647 = 4A                  A:02 X:02 Y:D4 P:67 SP:FB
	// F08A  C9 4A     CMP #$4A                        A:4A X:02 Y:D4 P:65 SP:FB
	// F08C  F0 02     BEQ $F090                       A:4A X:02 Y:D4 P:67 SP:FB
	// F090  C8        INY                             A:4A X:02 Y:D4 P:67 SP:FB
	// F091  A9 29     LDA #$29                        A:4A X:02 Y:D5 P:E5 SP:FB
	// F093  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:D5 P:65 SP:FB
	// F096  20 64 FB  JSR $FB64                       A:29 X:02 Y:D5 P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:02 Y:D5 P:65 SP:F9
	// FB65  18        CLC                             A:29 X:02 Y:D5 P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:02 Y:D5 P:24 SP:F9
	// FB68  60        RTS                             A:42 X:02 Y:D5 P:24 SP:F9
	// F099  23 45    *RLA ($45,X) @ 47 = 0647 = 29    A:42 X:02 Y:D5 P:24 SP:FB
	// F09B  EA        NOP                             A:42 X:02 Y:D5 P:24 SP:FB
	// F09C  EA        NOP                             A:42 X:02 Y:D5 P:24 SP:FB
	// F09D  EA        NOP                             A:42 X:02 Y:D5 P:24 SP:FB
	// F09E  EA        NOP                             A:42 X:02 Y:D5 P:24 SP:FB
	// F09F  20 69 FB  JSR $FB69                       A:42 X:02 Y:D5 P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:02 Y:D5 P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:02 Y:D5 P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:02 Y:D5 P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:02 Y:D5 P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:02 Y:D5 P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:02 Y:D5 P:27 SP:F9
	// FB77  60        RTS                             A:42 X:02 Y:D5 P:27 SP:F9
	// F0A2  AD 47 06  LDA $0647 = 52                  A:42 X:02 Y:D5 P:27 SP:FB
	// F0A5  C9 52     CMP #$52                        A:52 X:02 Y:D5 P:25 SP:FB
	// F0A7  F0 02     BEQ $F0AB                       A:52 X:02 Y:D5 P:27 SP:FB
	// F0AB  C8        INY                             A:52 X:02 Y:D5 P:27 SP:FB
	// F0AC  A9 37     LDA #$37                        A:52 X:02 Y:D6 P:A5 SP:FB
	// F0AE  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:D6 P:25 SP:FB
	// F0B1  20 68 FA  JSR $FA68                       A:37 X:02 Y:D6 P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:02 Y:D6 P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:02 Y:D6 P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:02 Y:D6 P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:02 Y:D6 P:65 SP:F9
	// F0B4  23 45    *RLA ($45,X) @ 47 = 0647 = 37    A:75 X:02 Y:D6 P:65 SP:FB
	// F0B6  EA        NOP                             A:65 X:02 Y:D6 P:64 SP:FB
	// F0B7  EA        NOP                             A:65 X:02 Y:D6 P:64 SP:FB
	// F0B8  EA        NOP                             A:65 X:02 Y:D6 P:64 SP:FB
	// F0B9  EA        NOP                             A:65 X:02 Y:D6 P:64 SP:FB
	// F0BA  20 6E FA  JSR $FA6E                       A:65 X:02 Y:D6 P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:02 Y:D6 P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:02 Y:D6 P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:02 Y:D6 P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:02 Y:D6 P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:02 Y:D6 P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:02 Y:D6 P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:02 Y:D6 P:67 SP:F9
	// F0BD  AD 47 06  LDA $0647 = 6F                  A:65 X:02 Y:D6 P:67 SP:FB
	// F0C0  C9 6F     CMP #$6F                        A:6F X:02 Y:D6 P:65 SP:FB
	// F0C2  F0 02     BEQ $F0C6                       A:6F X:02 Y:D6 P:67 SP:FB
	// F0C6  C8        INY                             A:6F X:02 Y:D6 P:67 SP:FB
	// F0C7  A9 A5     LDA #$A5                        A:6F X:02 Y:D7 P:E5 SP:FB
	// F0C9  85 47     STA $47 = 47                    A:A5 X:02 Y:D7 P:E5 SP:FB
	// F0CB  20 53 FB  JSR $FB53                       A:A5 X:02 Y:D7 P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:02 Y:D7 P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:02 Y:D7 P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:02 Y:D7 P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:02 Y:D7 P:E4 SP:F9
	// F0CE  27 47    *RLA $47 = A5                    A:B3 X:02 Y:D7 P:E4 SP:FB
	// F0D0  EA        NOP                             A:02 X:02 Y:D7 P:65 SP:FB
	// F0D1  EA        NOP                             A:02 X:02 Y:D7 P:65 SP:FB
	// F0D2  EA        NOP                             A:02 X:02 Y:D7 P:65 SP:FB
	// F0D3  EA        NOP                             A:02 X:02 Y:D7 P:65 SP:FB
	// F0D4  20 59 FB  JSR $FB59                       A:02 X:02 Y:D7 P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:02 Y:D7 P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:02 Y:D7 P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:02 Y:D7 P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:02 Y:D7 P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:02 Y:D7 P:67 SP:F9
	// FB63  60        RTS                             A:02 X:02 Y:D7 P:67 SP:F9
	// F0D7  A5 47     LDA $47 = 4A                    A:02 X:02 Y:D7 P:67 SP:FB
	// F0D9  C9 4A     CMP #$4A                        A:4A X:02 Y:D7 P:65 SP:FB
	// F0DB  F0 02     BEQ $F0DF                       A:4A X:02 Y:D7 P:67 SP:FB
	// F0DF  C8        INY                             A:4A X:02 Y:D7 P:67 SP:FB
	// F0E0  A9 29     LDA #$29                        A:4A X:02 Y:D8 P:E5 SP:FB
	// F0E2  85 47     STA $47 = 4A                    A:29 X:02 Y:D8 P:65 SP:FB
	// F0E4  20 64 FB  JSR $FB64                       A:29 X:02 Y:D8 P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:02 Y:D8 P:65 SP:F9
	// FB65  18        CLC                             A:29 X:02 Y:D8 P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:02 Y:D8 P:24 SP:F9
	// FB68  60        RTS                             A:42 X:02 Y:D8 P:24 SP:F9
	// F0E7  27 47    *RLA $47 = 29                    A:42 X:02 Y:D8 P:24 SP:FB
	// F0E9  EA        NOP                             A:42 X:02 Y:D8 P:24 SP:FB
	// F0EA  EA        NOP                             A:42 X:02 Y:D8 P:24 SP:FB
	// F0EB  EA        NOP                             A:42 X:02 Y:D8 P:24 SP:FB
	// F0EC  EA        NOP                             A:42 X:02 Y:D8 P:24 SP:FB
	// F0ED  20 69 FB  JSR $FB69                       A:42 X:02 Y:D8 P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:02 Y:D8 P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:02 Y:D8 P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:02 Y:D8 P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:02 Y:D8 P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:02 Y:D8 P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:02 Y:D8 P:27 SP:F9
	// FB77  60        RTS                             A:42 X:02 Y:D8 P:27 SP:F9
	// F0F0  A5 47     LDA $47 = 52                    A:42 X:02 Y:D8 P:27 SP:FB
	// F0F2  C9 52     CMP #$52                        A:52 X:02 Y:D8 P:25 SP:FB
	// F0F4  F0 02     BEQ $F0F8                       A:52 X:02 Y:D8 P:27 SP:FB
	// F0F8  C8        INY                             A:52 X:02 Y:D8 P:27 SP:FB
	// F0F9  A9 37     LDA #$37                        A:52 X:02 Y:D9 P:A5 SP:FB
	// F0FB  85 47     STA $47 = 52                    A:37 X:02 Y:D9 P:25 SP:FB
	// F0FD  20 68 FA  JSR $FA68                       A:37 X:02 Y:D9 P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:02 Y:D9 P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:02 Y:D9 P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:02 Y:D9 P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:02 Y:D9 P:65 SP:F9
	// F100  27 47    *RLA $47 = 37                    A:75 X:02 Y:D9 P:65 SP:FB
	// F102  EA        NOP                             A:65 X:02 Y:D9 P:64 SP:FB
	// F103  EA        NOP                             A:65 X:02 Y:D9 P:64 SP:FB
	// F104  EA        NOP                             A:65 X:02 Y:D9 P:64 SP:FB
	// F105  EA        NOP                             A:65 X:02 Y:D9 P:64 SP:FB
	// F106  20 6E FA  JSR $FA6E                       A:65 X:02 Y:D9 P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:02 Y:D9 P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:02 Y:D9 P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:02 Y:D9 P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:02 Y:D9 P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:02 Y:D9 P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:02 Y:D9 P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:02 Y:D9 P:67 SP:F9
	// F109  A5 47     LDA $47 = 6F                    A:65 X:02 Y:D9 P:67 SP:FB
	// F10B  C9 6F     CMP #$6F                        A:6F X:02 Y:D9 P:65 SP:FB
	// F10D  F0 02     BEQ $F111                       A:6F X:02 Y:D9 P:67 SP:FB
	// F111  C8        INY                             A:6F X:02 Y:D9 P:67 SP:FB
	// F112  A9 A5     LDA #$A5                        A:6F X:02 Y:DA P:E5 SP:FB
	// F114  8D 47 06  STA $0647 = 6F                  A:A5 X:02 Y:DA P:E5 SP:FB
	// F117  20 53 FB  JSR $FB53                       A:A5 X:02 Y:DA P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:02 Y:DA P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:02 Y:DA P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:02 Y:DA P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:02 Y:DA P:E4 SP:F9
	// F11A  2F 47 06 *RLA $0647 = A5                  A:B3 X:02 Y:DA P:E4 SP:FB
	// F11D  EA        NOP                             A:02 X:02 Y:DA P:65 SP:FB
	// F11E  EA        NOP                             A:02 X:02 Y:DA P:65 SP:FB
	// F11F  EA        NOP                             A:02 X:02 Y:DA P:65 SP:FB
	// F120  EA        NOP                             A:02 X:02 Y:DA P:65 SP:FB
	// F121  20 59 FB  JSR $FB59                       A:02 X:02 Y:DA P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:02 Y:DA P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:02 Y:DA P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:02 Y:DA P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:02 Y:DA P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:02 Y:DA P:67 SP:F9
	// FB63  60        RTS                             A:02 X:02 Y:DA P:67 SP:F9
	// F124  AD 47 06  LDA $0647 = 4A                  A:02 X:02 Y:DA P:67 SP:FB
	// F127  C9 4A     CMP #$4A                        A:4A X:02 Y:DA P:65 SP:FB
	// F129  F0 02     BEQ $F12D                       A:4A X:02 Y:DA P:67 SP:FB
	// F12D  C8        INY                             A:4A X:02 Y:DA P:67 SP:FB
	// F12E  A9 29     LDA #$29                        A:4A X:02 Y:DB P:E5 SP:FB
	// F130  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:DB P:65 SP:FB
	// F133  20 64 FB  JSR $FB64                       A:29 X:02 Y:DB P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:02 Y:DB P:65 SP:F9
	// FB65  18        CLC                             A:29 X:02 Y:DB P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:02 Y:DB P:24 SP:F9
	// FB68  60        RTS                             A:42 X:02 Y:DB P:24 SP:F9
	// F136  2F 47 06 *RLA $0647 = 29                  A:42 X:02 Y:DB P:24 SP:FB
	// F139  EA        NOP                             A:42 X:02 Y:DB P:24 SP:FB
	// F13A  EA        NOP                             A:42 X:02 Y:DB P:24 SP:FB
	// F13B  EA        NOP                             A:42 X:02 Y:DB P:24 SP:FB
	// F13C  EA        NOP                             A:42 X:02 Y:DB P:24 SP:FB
	// F13D  20 69 FB  JSR $FB69                       A:42 X:02 Y:DB P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:02 Y:DB P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:02 Y:DB P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:02 Y:DB P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:02 Y:DB P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:02 Y:DB P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:02 Y:DB P:27 SP:F9
	// FB77  60        RTS                             A:42 X:02 Y:DB P:27 SP:F9
	// F140  AD 47 06  LDA $0647 = 52                  A:42 X:02 Y:DB P:27 SP:FB
	// F143  C9 52     CMP #$52                        A:52 X:02 Y:DB P:25 SP:FB
	// F145  F0 02     BEQ $F149                       A:52 X:02 Y:DB P:27 SP:FB
	// F149  C8        INY                             A:52 X:02 Y:DB P:27 SP:FB
	// F14A  A9 37     LDA #$37                        A:52 X:02 Y:DC P:A5 SP:FB
	// F14C  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:DC P:25 SP:FB
	// F14F  20 68 FA  JSR $FA68                       A:37 X:02 Y:DC P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:02 Y:DC P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:02 Y:DC P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:02 Y:DC P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:02 Y:DC P:65 SP:F9
	// F152  2F 47 06 *RLA $0647 = 37                  A:75 X:02 Y:DC P:65 SP:FB
	// F155  EA        NOP                             A:65 X:02 Y:DC P:64 SP:FB
	// F156  EA        NOP                             A:65 X:02 Y:DC P:64 SP:FB
	// F157  EA        NOP                             A:65 X:02 Y:DC P:64 SP:FB
	// F158  EA        NOP                             A:65 X:02 Y:DC P:64 SP:FB
	// F159  20 6E FA  JSR $FA6E                       A:65 X:02 Y:DC P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:02 Y:DC P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:02 Y:DC P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:02 Y:DC P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:02 Y:DC P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:02 Y:DC P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:02 Y:DC P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:02 Y:DC P:67 SP:F9
	// F15C  AD 47 06  LDA $0647 = 6F                  A:65 X:02 Y:DC P:67 SP:FB
	// F15F  C9 6F     CMP #$6F                        A:6F X:02 Y:DC P:65 SP:FB
	// F161  F0 02     BEQ $F165                       A:6F X:02 Y:DC P:67 SP:FB
	// F165  A9 A5     LDA #$A5                        A:6F X:02 Y:DC P:67 SP:FB
	// F167  8D 47 06  STA $0647 = 6F                  A:A5 X:02 Y:DC P:E5 SP:FB
	// F16A  A9 48     LDA #$48                        A:A5 X:02 Y:DC P:E5 SP:FB
	// F16C  85 45     STA $45 = 48                    A:48 X:02 Y:DC P:65 SP:FB
	// F16E  A9 05     LDA #$05                        A:48 X:02 Y:DC P:65 SP:FB
	// F170  85 46     STA $46 = 05                    A:05 X:02 Y:DC P:65 SP:FB
	// F172  A0 FF     LDY #$FF                        A:05 X:02 Y:DC P:65 SP:FB
	// F174  20 53 FB  JSR $FB53                       A:05 X:02 Y:FF P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:E5 SP:F9
	// FB55  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:05 X:02 Y:FF P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:02 Y:FF P:E4 SP:F9
	// F177  33 45    *RLA ($45),Y = 0548 @ 0647 = A5  A:B3 X:02 Y:FF P:E4 SP:FB
	// F179  EA        NOP                             A:02 X:02 Y:FF P:65 SP:FB
	// F17A  EA        NOP                             A:02 X:02 Y:FF P:65 SP:FB
	// F17B  08        PHP                             A:02 X:02 Y:FF P:65 SP:FB
	// F17C  48        PHA                             A:02 X:02 Y:FF P:65 SP:FA
	// F17D  A0 DD     LDY #$DD                        A:02 X:02 Y:FF P:65 SP:F9
	// F17F  68        PLA                             A:02 X:02 Y:DD P:E5 SP:F9
	// F180  28        PLP                             A:02 X:02 Y:DD P:65 SP:FA
	// F181  20 59 FB  JSR $FB59                       A:02 X:02 Y:DD P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:02 Y:DD P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:02 Y:DD P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:02 Y:DD P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:02 Y:DD P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:02 Y:DD P:67 SP:F9
	// FB63  60        RTS                             A:02 X:02 Y:DD P:67 SP:F9
	// F184  AD 47 06  LDA $0647 = 4A                  A:02 X:02 Y:DD P:67 SP:FB
	// F187  C9 4A     CMP #$4A                        A:4A X:02 Y:DD P:65 SP:FB
	// F189  F0 02     BEQ $F18D                       A:4A X:02 Y:DD P:67 SP:FB
	// F18D  A0 FF     LDY #$FF                        A:4A X:02 Y:DD P:67 SP:FB
	// F18F  A9 29     LDA #$29                        A:4A X:02 Y:FF P:E5 SP:FB
	// F191  8D 47 06  STA $0647 = 4A                  A:29 X:02 Y:FF P:65 SP:FB
	// F194  20 64 FB  JSR $FB64                       A:29 X:02 Y:FF P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:02 Y:FF P:65 SP:F9
	// FB65  18        CLC                             A:29 X:02 Y:FF P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:02 Y:FF P:24 SP:F9
	// FB68  60        RTS                             A:42 X:02 Y:FF P:24 SP:F9
	// F197  33 45    *RLA ($45),Y = 0548 @ 0647 = 29  A:42 X:02 Y:FF P:24 SP:FB
	// F199  EA        NOP                             A:42 X:02 Y:FF P:24 SP:FB
	// F19A  EA        NOP                             A:42 X:02 Y:FF P:24 SP:FB
	// F19B  08        PHP                             A:42 X:02 Y:FF P:24 SP:FB
	// F19C  48        PHA                             A:42 X:02 Y:FF P:24 SP:FA
	// F19D  A0 DE     LDY #$DE                        A:42 X:02 Y:FF P:24 SP:F9
	// F19F  68        PLA                             A:42 X:02 Y:DE P:A4 SP:F9
	// F1A0  28        PLP                             A:42 X:02 Y:DE P:24 SP:FA
	// F1A1  20 69 FB  JSR $FB69                       A:42 X:02 Y:DE P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:02 Y:DE P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:02 Y:DE P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:02 Y:DE P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:02 Y:DE P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:02 Y:DE P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:02 Y:DE P:27 SP:F9
	// FB77  60        RTS                             A:42 X:02 Y:DE P:27 SP:F9
	// F1A4  AD 47 06  LDA $0647 = 52                  A:42 X:02 Y:DE P:27 SP:FB
	// F1A7  C9 52     CMP #$52                        A:52 X:02 Y:DE P:25 SP:FB
	// F1A9  F0 02     BEQ $F1AD                       A:52 X:02 Y:DE P:27 SP:FB
	// F1AD  A0 FF     LDY #$FF                        A:52 X:02 Y:DE P:27 SP:FB
	// F1AF  A9 37     LDA #$37                        A:52 X:02 Y:FF P:A5 SP:FB
	// F1B1  8D 47 06  STA $0647 = 52                  A:37 X:02 Y:FF P:25 SP:FB
	// F1B4  20 68 FA  JSR $FA68                       A:37 X:02 Y:FF P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:02 Y:FF P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:02 Y:FF P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:02 Y:FF P:65 SP:F9
	// F1B7  33 45    *RLA ($45),Y = 0548 @ 0647 = 37  A:75 X:02 Y:FF P:65 SP:FB
	// F1B9  EA        NOP                             A:65 X:02 Y:FF P:64 SP:FB
	// F1BA  EA        NOP                             A:65 X:02 Y:FF P:64 SP:FB
	// F1BB  08        PHP                             A:65 X:02 Y:FF P:64 SP:FB
	// F1BC  48        PHA                             A:65 X:02 Y:FF P:64 SP:FA
	// F1BD  A0 DF     LDY #$DF                        A:65 X:02 Y:FF P:64 SP:F9
	// F1BF  68        PLA                             A:65 X:02 Y:DF P:E4 SP:F9
	// F1C0  28        PLP                             A:65 X:02 Y:DF P:64 SP:FA
	// F1C1  20 6E FA  JSR $FA6E                       A:65 X:02 Y:DF P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:02 Y:DF P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:02 Y:DF P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:02 Y:DF P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:02 Y:DF P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:02 Y:DF P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:02 Y:DF P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:02 Y:DF P:67 SP:F9
	// F1C4  AD 47 06  LDA $0647 = 6F                  A:65 X:02 Y:DF P:67 SP:FB
	// F1C7  C9 6F     CMP #$6F                        A:6F X:02 Y:DF P:65 SP:FB
	// F1C9  F0 02     BEQ $F1CD                       A:6F X:02 Y:DF P:67 SP:FB
	// F1CD  A0 E0     LDY #$E0                        A:6F X:02 Y:DF P:67 SP:FB
	// F1CF  A2 FF     LDX #$FF                        A:6F X:02 Y:E0 P:E5 SP:FB
	// F1D1  A9 A5     LDA #$A5                        A:6F X:FF Y:E0 P:E5 SP:FB
	// F1D3  85 47     STA $47 = 6F                    A:A5 X:FF Y:E0 P:E5 SP:FB
	// F1D5  20 53 FB  JSR $FB53                       A:A5 X:FF Y:E0 P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:E0 P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:FF Y:E0 P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:E0 P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:FF Y:E0 P:E4 SP:F9
	// F1D8  37 48    *RLA $48,X @ 47 = A5             A:B3 X:FF Y:E0 P:E4 SP:FB
	// F1DA  EA        NOP                             A:02 X:FF Y:E0 P:65 SP:FB
	// F1DB  EA        NOP                             A:02 X:FF Y:E0 P:65 SP:FB
	// F1DC  EA        NOP                             A:02 X:FF Y:E0 P:65 SP:FB
	// F1DD  EA        NOP                             A:02 X:FF Y:E0 P:65 SP:FB
	// F1DE  20 59 FB  JSR $FB59                       A:02 X:FF Y:E0 P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:FF Y:E0 P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:FF Y:E0 P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:FF Y:E0 P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:FF Y:E0 P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:FF Y:E0 P:67 SP:F9
	// FB63  60        RTS                             A:02 X:FF Y:E0 P:67 SP:F9
	// F1E1  A5 47     LDA $47 = 4A                    A:02 X:FF Y:E0 P:67 SP:FB
	// F1E3  C9 4A     CMP #$4A                        A:4A X:FF Y:E0 P:65 SP:FB
	// F1E5  F0 02     BEQ $F1E9                       A:4A X:FF Y:E0 P:67 SP:FB
	// F1E9  C8        INY                             A:4A X:FF Y:E0 P:67 SP:FB
	// F1EA  A9 29     LDA #$29                        A:4A X:FF Y:E1 P:E5 SP:FB
	// F1EC  85 47     STA $47 = 4A                    A:29 X:FF Y:E1 P:65 SP:FB
	// F1EE  20 64 FB  JSR $FB64                       A:29 X:FF Y:E1 P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:FF Y:E1 P:65 SP:F9
	// FB65  18        CLC                             A:29 X:FF Y:E1 P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:FF Y:E1 P:24 SP:F9
	// FB68  60        RTS                             A:42 X:FF Y:E1 P:24 SP:F9
	// F1F1  37 48    *RLA $48,X @ 47 = 29             A:42 X:FF Y:E1 P:24 SP:FB
	// F1F3  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB
	// F1F4  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB
	// F1F5  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB
	// F1F6  EA        NOP                             A:42 X:FF Y:E1 P:24 SP:FB
	// F1F7  20 69 FB  JSR $FB69                       A:42 X:FF Y:E1 P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E1 P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E1 P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E1 P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E1 P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:FF Y:E1 P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E1 P:27 SP:F9
	// FB77  60        RTS                             A:42 X:FF Y:E1 P:27 SP:F9
	// F1FA  A5 47     LDA $47 = 52                    A:42 X:FF Y:E1 P:27 SP:FB
	// F1FC  C9 52     CMP #$52                        A:52 X:FF Y:E1 P:25 SP:FB
	// F1FE  F0 02     BEQ $F202                       A:52 X:FF Y:E1 P:27 SP:FB
	// F202  C8        INY                             A:52 X:FF Y:E1 P:27 SP:FB
	// F203  A9 37     LDA #$37                        A:52 X:FF Y:E2 P:A5 SP:FB
	// F205  85 47     STA $47 = 52                    A:37 X:FF Y:E2 P:25 SP:FB
	// F207  20 68 FA  JSR $FA68                       A:37 X:FF Y:E2 P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:E2 P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:FF Y:E2 P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:FF Y:E2 P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:FF Y:E2 P:65 SP:F9
	// F20A  37 48    *RLA $48,X @ 47 = 37             A:75 X:FF Y:E2 P:65 SP:FB
	// F20C  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB
	// F20D  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB
	// F20E  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB
	// F20F  EA        NOP                             A:65 X:FF Y:E2 P:64 SP:FB
	// F210  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E2 P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E2 P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:FF Y:E2 P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E2 P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:FF Y:E2 P:67 SP:F9
	// F213  A5 47     LDA $47 = 6F                    A:65 X:FF Y:E2 P:67 SP:FB
	// F215  C9 6F     CMP #$6F                        A:6F X:FF Y:E2 P:65 SP:FB
	// F217  F0 02     BEQ $F21B                       A:6F X:FF Y:E2 P:67 SP:FB
	// F21B  A9 A5     LDA #$A5                        A:6F X:FF Y:E2 P:67 SP:FB
	// F21D  8D 47 06  STA $0647 = 6F                  A:A5 X:FF Y:E2 P:E5 SP:FB
	// F220  A0 FF     LDY #$FF                        A:A5 X:FF Y:E2 P:E5 SP:FB
	// F222  20 53 FB  JSR $FB53                       A:A5 X:FF Y:FF P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9
	// F225  3B 48 05 *RLA $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB
	// F228  EA        NOP                             A:02 X:FF Y:FF P:65 SP:FB
	// F229  EA        NOP                             A:02 X:FF Y:FF P:65 SP:FB
	// F22A  08        PHP                             A:02 X:FF Y:FF P:65 SP:FB
	// F22B  48        PHA                             A:02 X:FF Y:FF P:65 SP:FA
	// F22C  A0 E3     LDY #$E3                        A:02 X:FF Y:FF P:65 SP:F9
	// F22E  68        PLA                             A:02 X:FF Y:E3 P:E5 SP:F9
	// F22F  28        PLP                             A:02 X:FF Y:E3 P:65 SP:FA
	// F230  20 59 FB  JSR $FB59                       A:02 X:FF Y:E3 P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:FF Y:E3 P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:FF Y:E3 P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:FF Y:E3 P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:FF Y:E3 P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:FF Y:E3 P:67 SP:F9
	// FB63  60        RTS                             A:02 X:FF Y:E3 P:67 SP:F9
	// F233  AD 47 06  LDA $0647 = 4A                  A:02 X:FF Y:E3 P:67 SP:FB
	// F236  C9 4A     CMP #$4A                        A:4A X:FF Y:E3 P:65 SP:FB
	// F238  F0 02     BEQ $F23C                       A:4A X:FF Y:E3 P:67 SP:FB
	// F23C  A0 FF     LDY #$FF                        A:4A X:FF Y:E3 P:67 SP:FB
	// F23E  A9 29     LDA #$29                        A:4A X:FF Y:FF P:E5 SP:FB
	// F240  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:FF P:65 SP:FB
	// F243  20 64 FB  JSR $FB64                       A:29 X:FF Y:FF P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9
	// FB65  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9
	// FB68  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9
	// F246  3B 48 05 *RLA $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB
	// F249  EA        NOP                             A:42 X:FF Y:FF P:24 SP:FB
	// F24A  EA        NOP                             A:42 X:FF Y:FF P:24 SP:FB
	// F24B  08        PHP                             A:42 X:FF Y:FF P:24 SP:FB
	// F24C  48        PHA                             A:42 X:FF Y:FF P:24 SP:FA
	// F24D  A0 E4     LDY #$E4                        A:42 X:FF Y:FF P:24 SP:F9
	// F24F  68        PLA                             A:42 X:FF Y:E4 P:A4 SP:F9
	// F250  28        PLP                             A:42 X:FF Y:E4 P:24 SP:FA
	// F251  20 69 FB  JSR $FB69                       A:42 X:FF Y:E4 P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E4 P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E4 P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E4 P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E4 P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:FF Y:E4 P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E4 P:27 SP:F9
	// FB77  60        RTS                             A:42 X:FF Y:E4 P:27 SP:F9
	// F254  AD 47 06  LDA $0647 = 52                  A:42 X:FF Y:E4 P:27 SP:FB
	// F257  C9 52     CMP #$52                        A:52 X:FF Y:E4 P:25 SP:FB
	// F259  F0 02     BEQ $F25D                       A:52 X:FF Y:E4 P:27 SP:FB
	// F25D  A0 FF     LDY #$FF                        A:52 X:FF Y:E4 P:27 SP:FB
	// F25F  A9 37     LDA #$37                        A:52 X:FF Y:FF P:A5 SP:FB
	// F261  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:FF P:25 SP:FB
	// F264  20 68 FA  JSR $FA68                       A:37 X:FF Y:FF P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9
	// F267  3B 48 05 *RLA $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB
	// F26A  EA        NOP                             A:65 X:FF Y:FF P:64 SP:FB
	// F26B  EA        NOP                             A:65 X:FF Y:FF P:64 SP:FB
	// F26C  08        PHP                             A:65 X:FF Y:FF P:64 SP:FB
	// F26D  48        PHA                             A:65 X:FF Y:FF P:64 SP:FA
	// F26E  A0 E5     LDY #$E5                        A:65 X:FF Y:FF P:64 SP:F9
	// F270  68        PLA                             A:65 X:FF Y:E5 P:E4 SP:F9
	// F271  28        PLP                             A:65 X:FF Y:E5 P:64 SP:FA
	// F272  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E5 P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E5 P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:FF Y:E5 P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E5 P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:FF Y:E5 P:67 SP:F9
	// F275  AD 47 06  LDA $0647 = 6F                  A:65 X:FF Y:E5 P:67 SP:FB
	// F278  C9 6F     CMP #$6F                        A:6F X:FF Y:E5 P:65 SP:FB
	// F27A  F0 02     BEQ $F27E                       A:6F X:FF Y:E5 P:67 SP:FB
	// F27E  A0 E6     LDY #$E6                        A:6F X:FF Y:E5 P:67 SP:FB
	// F280  A2 FF     LDX #$FF                        A:6F X:FF Y:E6 P:E5 SP:FB
	// F282  A9 A5     LDA #$A5                        A:6F X:FF Y:E6 P:E5 SP:FB
	// F284  8D 47 06  STA $0647 = 6F                  A:A5 X:FF Y:E6 P:E5 SP:FB
	// F287  20 53 FB  JSR $FB53                       A:A5 X:FF Y:E6 P:E5 SP:FB
	// FB53  24 01     BIT $01 = FF                    A:A5 X:FF Y:E6 P:E5 SP:F9
	// FB55  18        CLC                             A:A5 X:FF Y:E6 P:E5 SP:F9
	// FB56  A9 B3     LDA #$B3                        A:A5 X:FF Y:E6 P:E4 SP:F9
	// FB58  60        RTS                             A:B3 X:FF Y:E6 P:E4 SP:F9
	// F28A  3F 48 05 *RLA $0548,X @ 0647 = A5         A:B3 X:FF Y:E6 P:E4 SP:FB
	// F28D  EA        NOP                             A:02 X:FF Y:E6 P:65 SP:FB
	// F28E  EA        NOP                             A:02 X:FF Y:E6 P:65 SP:FB
	// F28F  EA        NOP                             A:02 X:FF Y:E6 P:65 SP:FB
	// F290  EA        NOP                             A:02 X:FF Y:E6 P:65 SP:FB
	// F291  20 59 FB  JSR $FB59                       A:02 X:FF Y:E6 P:65 SP:FB
	// FB59  50 1A     BVC $FB75                       A:02 X:FF Y:E6 P:65 SP:F9
	// FB5B  90 18     BCC $FB75                       A:02 X:FF Y:E6 P:65 SP:F9
	// FB5D  30 16     BMI $FB75                       A:02 X:FF Y:E6 P:65 SP:F9
	// FB5F  C9 02     CMP #$02                        A:02 X:FF Y:E6 P:65 SP:F9
	// FB61  D0 12     BNE $FB75                       A:02 X:FF Y:E6 P:67 SP:F9
	// FB63  60        RTS                             A:02 X:FF Y:E6 P:67 SP:F9
	// F294  AD 47 06  LDA $0647 = 4A                  A:02 X:FF Y:E6 P:67 SP:FB
	// F297  C9 4A     CMP #$4A                        A:4A X:FF Y:E6 P:65 SP:FB
	// F299  F0 02     BEQ $F29D                       A:4A X:FF Y:E6 P:67 SP:FB
	// F29D  C8        INY                             A:4A X:FF Y:E6 P:67 SP:FB
	// F29E  A9 29     LDA #$29                        A:4A X:FF Y:E7 P:E5 SP:FB
	// F2A0  8D 47 06  STA $0647 = 4A                  A:29 X:FF Y:E7 P:65 SP:FB
	// F2A3  20 64 FB  JSR $FB64                       A:29 X:FF Y:E7 P:65 SP:FB
	// FB64  B8        CLV                             A:29 X:FF Y:E7 P:65 SP:F9
	// FB65  18        CLC                             A:29 X:FF Y:E7 P:25 SP:F9
	// FB66  A9 42     LDA #$42                        A:29 X:FF Y:E7 P:24 SP:F9
	// FB68  60        RTS                             A:42 X:FF Y:E7 P:24 SP:F9
	// F2A6  3F 48 05 *RLA $0548,X @ 0647 = 29         A:42 X:FF Y:E7 P:24 SP:FB
	// F2A9  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB
	// F2AA  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB
	// F2AB  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB
	// F2AC  EA        NOP                             A:42 X:FF Y:E7 P:24 SP:FB
	// F2AD  20 69 FB  JSR $FB69                       A:42 X:FF Y:E7 P:24 SP:FB
	// FB69  70 0A     BVS $FB75                       A:42 X:FF Y:E7 P:24 SP:F9
	// FB6B  F0 08     BEQ $FB75                       A:42 X:FF Y:E7 P:24 SP:F9
	// FB6D  30 06     BMI $FB75                       A:42 X:FF Y:E7 P:24 SP:F9
	// FB6F  B0 04     BCS $FB75                       A:42 X:FF Y:E7 P:24 SP:F9
	// FB71  C9 42     CMP #$42                        A:42 X:FF Y:E7 P:24 SP:F9
	// FB73  F0 02     BEQ $FB77                       A:42 X:FF Y:E7 P:27 SP:F9
	// FB77  60        RTS                             A:42 X:FF Y:E7 P:27 SP:F9
	// F2B0  AD 47 06  LDA $0647 = 52                  A:42 X:FF Y:E7 P:27 SP:FB
	// F2B3  C9 52     CMP #$52                        A:52 X:FF Y:E7 P:25 SP:FB
	// F2B5  F0 02     BEQ $F2B9                       A:52 X:FF Y:E7 P:27 SP:FB
	// F2B9  C8        INY                             A:52 X:FF Y:E7 P:27 SP:FB
	// F2BA  A9 37     LDA #$37                        A:52 X:FF Y:E8 P:A5 SP:FB
	// F2BC  8D 47 06  STA $0647 = 52                  A:37 X:FF Y:E8 P:25 SP:FB
	// F2BF  20 68 FA  JSR $FA68                       A:37 X:FF Y:E8 P:25 SP:FB
	// FA68  24 01     BIT $01 = FF                    A:37 X:FF Y:E8 P:25 SP:F9
	// FA6A  38        SEC                             A:37 X:FF Y:E8 P:E5 SP:F9
	// FA6B  A9 75     LDA #$75                        A:37 X:FF Y:E8 P:E5 SP:F9
	// FA6D  60        RTS                             A:75 X:FF Y:E8 P:65 SP:F9
	// F2C2  3F 48 05 *RLA $0548,X @ 0647 = 37         A:75 X:FF Y:E8 P:65 SP:FB
	// F2C5  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB
	// F2C6  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB
	// F2C7  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB
	// F2C8  EA        NOP                             A:65 X:FF Y:E8 P:64 SP:FB
	// F2C9  20 6E FA  JSR $FA6E                       A:65 X:FF Y:E8 P:64 SP:FB
	// FA6E  50 76     BVC $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9
	// FA70  F0 74     BEQ $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9
	// FA72  30 72     BMI $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9
	// FA74  B0 70     BCS $FAE6                       A:65 X:FF Y:E8 P:64 SP:F9
	// FA76  C9 65     CMP #$65                        A:65 X:FF Y:E8 P:64 SP:F9
	// FA78  D0 6C     BNE $FAE6                       A:65 X:FF Y:E8 P:67 SP:F9
	// FA7A  60        RTS                             A:65 X:FF Y:E8 P:67 SP:F9
	// F2CC  AD 47 06  LDA $0647 = 6F                  A:65 X:FF Y:E8 P:67 SP:FB
	// F2CF  C9 6F     CMP #$6F                        A:6F X:FF Y:E8 P:65 SP:FB
	// F2D1  F0 02     BEQ $F2D5                       A:6F X:FF Y:E8 P:67 SP:FB
	// F2D5  60        RTS                             A:6F X:FF Y:E8 P:67 SP:FB
	// C647  20 D6 F2  JSR $F2D6                       A:6F X:FF Y:E8 P:67 SP:FD
	// F2D6  A9 FF     LDA #$FF                        A:6F X:FF Y:E8 P:67 SP:FB
	// F2D8  85 01     STA $01 = FF                    A:FF X:FF Y:E8 P:E5 SP:FB
	// F2DA  A0 E9     LDY #$E9                        A:FF X:FF Y:E8 P:E5 SP:FB
	// F2DC  A2 02     LDX #$02                        A:FF X:FF Y:E9 P:E5 SP:FB
	// F2DE  A9 47     LDA #$47                        A:FF X:02 Y:E9 P:65 SP:FB
	// F2E0  85 47     STA $47 = 6F                    A:47 X:02 Y:E9 P:65 SP:FB
	// F2E2  A9 06     LDA #$06                        A:47 X:02 Y:E9 P:65 SP:FB
	// F2E4  85 48     STA $48 = 06                    A:06 X:02 Y:E9 P:65 SP:FB
	// F2E6  A9 A5     LDA #$A5                        A:06 X:02 Y:E9 P:65 SP:FB
	// F2E8  8D 47 06  STA $0647 = 6F                  A:A5 X:02 Y:E9 P:E5 SP:FB
	// F2EB  20 1D FB  JSR $FB1D                       A:A5 X:02 Y:E9 P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:02 Y:E9 P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:02 Y:E9 P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:02 Y:E9 P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:02 Y:E9 P:E4 SP:F9
	// F2EE  43 45    *SRE ($45,X) @ 47 = 0647 = A5    A:B3 X:02 Y:E9 P:E4 SP:FB
	// F2F0  EA        NOP                             A:E1 X:02 Y:E9 P:E5 SP:FB
	// F2F1  EA        NOP                             A:E1 X:02 Y:E9 P:E5 SP:FB
	// F2F2  EA        NOP                             A:E1 X:02 Y:E9 P:E5 SP:FB
	// F2F3  EA        NOP                             A:E1 X:02 Y:E9 P:E5 SP:FB
	// F2F4  20 23 FB  JSR $FB23                       A:E1 X:02 Y:E9 P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:02 Y:E9 P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:02 Y:E9 P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:02 Y:E9 P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:02 Y:E9 P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:02 Y:E9 P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:02 Y:E9 P:67 SP:F9
	// F2F7  AD 47 06  LDA $0647 = 52                  A:E1 X:02 Y:E9 P:67 SP:FB
	// F2FA  C9 52     CMP #$52                        A:52 X:02 Y:E9 P:65 SP:FB
	// F2FC  F0 02     BEQ $F300                       A:52 X:02 Y:E9 P:67 SP:FB
	// F300  C8        INY                             A:52 X:02 Y:E9 P:67 SP:FB
	// F301  A9 29     LDA #$29                        A:52 X:02 Y:EA P:E5 SP:FB
	// F303  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:EA P:65 SP:FB
	// F306  20 2E FB  JSR $FB2E                       A:29 X:02 Y:EA P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:02 Y:EA P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:02 Y:EA P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:02 Y:EA P:24 SP:F9
	// FB32  60        RTS                             A:42 X:02 Y:EA P:24 SP:F9
	// F309  43 45    *SRE ($45,X) @ 47 = 0647 = 29    A:42 X:02 Y:EA P:24 SP:FB
	// F30B  EA        NOP                             A:56 X:02 Y:EA P:25 SP:FB
	// F30C  EA        NOP                             A:56 X:02 Y:EA P:25 SP:FB
	// F30D  EA        NOP                             A:56 X:02 Y:EA P:25 SP:FB
	// F30E  EA        NOP                             A:56 X:02 Y:EA P:25 SP:FB
	// F30F  20 33 FB  JSR $FB33                       A:56 X:02 Y:EA P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:02 Y:EA P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:02 Y:EA P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:02 Y:EA P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:02 Y:EA P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:02 Y:EA P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:02 Y:EA P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:02 Y:EA P:27 SP:F9
	// F312  AD 47 06  LDA $0647 = 14                  A:56 X:02 Y:EA P:27 SP:FB
	// F315  C9 14     CMP #$14                        A:14 X:02 Y:EA P:25 SP:FB
	// F317  F0 02     BEQ $F31B                       A:14 X:02 Y:EA P:27 SP:FB
	// F31B  C8        INY                             A:14 X:02 Y:EA P:27 SP:FB
	// F31C  A9 37     LDA #$37                        A:14 X:02 Y:EB P:A5 SP:FB
	// F31E  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:EB P:25 SP:FB
	// F321  20 40 FB  JSR $FB40                       A:37 X:02 Y:EB P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:02 Y:EB P:25 SP:F9
	// FB42  38        SEC                             A:37 X:02 Y:EB P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:02 Y:EB P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:02 Y:EB P:65 SP:F9
	// F324  43 45    *SRE ($45,X) @ 47 = 0647 = 37    A:75 X:02 Y:EB P:65 SP:FB
	// F326  EA        NOP                             A:6E X:02 Y:EB P:65 SP:FB
	// F327  EA        NOP                             A:6E X:02 Y:EB P:65 SP:FB
	// F328  EA        NOP                             A:6E X:02 Y:EB P:65 SP:FB
	// F329  EA        NOP                             A:6E X:02 Y:EB P:65 SP:FB
	// F32A  20 46 FB  JSR $FB46                       A:6E X:02 Y:EB P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:02 Y:EB P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:02 Y:EB P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:02 Y:EB P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:02 Y:EB P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:02 Y:EB P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:02 Y:EB P:67 SP:F9
	// FB52  60        RTS                             A:6E X:02 Y:EB P:67 SP:F9
	// F32D  AD 47 06  LDA $0647 = 1B                  A:6E X:02 Y:EB P:67 SP:FB
	// F330  C9 1B     CMP #$1B                        A:1B X:02 Y:EB P:65 SP:FB
	// F332  F0 02     BEQ $F336                       A:1B X:02 Y:EB P:67 SP:FB
	// F336  C8        INY                             A:1B X:02 Y:EB P:67 SP:FB
	// F337  A9 A5     LDA #$A5                        A:1B X:02 Y:EC P:E5 SP:FB
	// F339  85 47     STA $47 = 47                    A:A5 X:02 Y:EC P:E5 SP:FB
	// F33B  20 1D FB  JSR $FB1D                       A:A5 X:02 Y:EC P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:02 Y:EC P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:02 Y:EC P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:02 Y:EC P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:02 Y:EC P:E4 SP:F9
	// F33E  47 47    *SRE $47 = A5                    A:B3 X:02 Y:EC P:E4 SP:FB
	// F340  EA        NOP                             A:E1 X:02 Y:EC P:E5 SP:FB
	// F341  EA        NOP                             A:E1 X:02 Y:EC P:E5 SP:FB
	// F342  EA        NOP                             A:E1 X:02 Y:EC P:E5 SP:FB
	// F343  EA        NOP                             A:E1 X:02 Y:EC P:E5 SP:FB
	// F344  20 23 FB  JSR $FB23                       A:E1 X:02 Y:EC P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:02 Y:EC P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:02 Y:EC P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:02 Y:EC P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:02 Y:EC P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:02 Y:EC P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:02 Y:EC P:67 SP:F9
	// F347  A5 47     LDA $47 = 52                    A:E1 X:02 Y:EC P:67 SP:FB
	// F349  C9 52     CMP #$52                        A:52 X:02 Y:EC P:65 SP:FB
	// F34B  F0 02     BEQ $F34F                       A:52 X:02 Y:EC P:67 SP:FB
	// F34F  C8        INY                             A:52 X:02 Y:EC P:67 SP:FB
	// F350  A9 29     LDA #$29                        A:52 X:02 Y:ED P:E5 SP:FB
	// F352  85 47     STA $47 = 52                    A:29 X:02 Y:ED P:65 SP:FB
	// F354  20 2E FB  JSR $FB2E                       A:29 X:02 Y:ED P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:02 Y:ED P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:02 Y:ED P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:02 Y:ED P:24 SP:F9
	// FB32  60        RTS                             A:42 X:02 Y:ED P:24 SP:F9
	// F357  47 47    *SRE $47 = 29                    A:42 X:02 Y:ED P:24 SP:FB
	// F359  EA        NOP                             A:56 X:02 Y:ED P:25 SP:FB
	// F35A  EA        NOP                             A:56 X:02 Y:ED P:25 SP:FB
	// F35B  EA        NOP                             A:56 X:02 Y:ED P:25 SP:FB
	// F35C  EA        NOP                             A:56 X:02 Y:ED P:25 SP:FB
	// F35D  20 33 FB  JSR $FB33                       A:56 X:02 Y:ED P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:02 Y:ED P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:02 Y:ED P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:02 Y:ED P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:02 Y:ED P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:02 Y:ED P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:02 Y:ED P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:02 Y:ED P:27 SP:F9
	// F360  A5 47     LDA $47 = 14                    A:56 X:02 Y:ED P:27 SP:FB
	// F362  C9 14     CMP #$14                        A:14 X:02 Y:ED P:25 SP:FB
	// F364  F0 02     BEQ $F368                       A:14 X:02 Y:ED P:27 SP:FB
	// F368  C8        INY                             A:14 X:02 Y:ED P:27 SP:FB
	// F369  A9 37     LDA #$37                        A:14 X:02 Y:EE P:A5 SP:FB
	// F36B  85 47     STA $47 = 14                    A:37 X:02 Y:EE P:25 SP:FB
	// F36D  20 40 FB  JSR $FB40                       A:37 X:02 Y:EE P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:02 Y:EE P:25 SP:F9
	// FB42  38        SEC                             A:37 X:02 Y:EE P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:02 Y:EE P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:02 Y:EE P:65 SP:F9
	// F370  47 47    *SRE $47 = 37                    A:75 X:02 Y:EE P:65 SP:FB
	// F372  EA        NOP                             A:6E X:02 Y:EE P:65 SP:FB
	// F373  EA        NOP                             A:6E X:02 Y:EE P:65 SP:FB
	// F374  EA        NOP                             A:6E X:02 Y:EE P:65 SP:FB
	// F375  EA        NOP                             A:6E X:02 Y:EE P:65 SP:FB
	// F376  20 46 FB  JSR $FB46                       A:6E X:02 Y:EE P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:02 Y:EE P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:02 Y:EE P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:02 Y:EE P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:02 Y:EE P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:02 Y:EE P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:02 Y:EE P:67 SP:F9
	// FB52  60        RTS                             A:6E X:02 Y:EE P:67 SP:F9
	// F379  A5 47     LDA $47 = 1B                    A:6E X:02 Y:EE P:67 SP:FB
	// F37B  C9 1B     CMP #$1B                        A:1B X:02 Y:EE P:65 SP:FB
	// F37D  F0 02     BEQ $F381                       A:1B X:02 Y:EE P:67 SP:FB
	// F381  C8        INY                             A:1B X:02 Y:EE P:67 SP:FB
	// F382  A9 A5     LDA #$A5                        A:1B X:02 Y:EF P:E5 SP:FB
	// F384  8D 47 06  STA $0647 = 1B                  A:A5 X:02 Y:EF P:E5 SP:FB
	// F387  20 1D FB  JSR $FB1D                       A:A5 X:02 Y:EF P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:02 Y:EF P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:02 Y:EF P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:02 Y:EF P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:02 Y:EF P:E4 SP:F9
	// F38A  4F 47 06 *SRE $0647 = A5                  A:B3 X:02 Y:EF P:E4 SP:FB
	// F38D  EA        NOP                             A:E1 X:02 Y:EF P:E5 SP:FB
	// F38E  EA        NOP                             A:E1 X:02 Y:EF P:E5 SP:FB
	// F38F  EA        NOP                             A:E1 X:02 Y:EF P:E5 SP:FB
	// F390  EA        NOP                             A:E1 X:02 Y:EF P:E5 SP:FB
	// F391  20 23 FB  JSR $FB23                       A:E1 X:02 Y:EF P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:02 Y:EF P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:02 Y:EF P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:02 Y:EF P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:02 Y:EF P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:02 Y:EF P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:02 Y:EF P:67 SP:F9
	// F394  AD 47 06  LDA $0647 = 52                  A:E1 X:02 Y:EF P:67 SP:FB
	// F397  C9 52     CMP #$52                        A:52 X:02 Y:EF P:65 SP:FB
	// F399  F0 02     BEQ $F39D                       A:52 X:02 Y:EF P:67 SP:FB
	// F39D  C8        INY                             A:52 X:02 Y:EF P:67 SP:FB
	// F39E  A9 29     LDA #$29                        A:52 X:02 Y:F0 P:E5 SP:FB
	// F3A0  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:F0 P:65 SP:FB
	// F3A3  20 2E FB  JSR $FB2E                       A:29 X:02 Y:F0 P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:02 Y:F0 P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:02 Y:F0 P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:02 Y:F0 P:24 SP:F9
	// FB32  60        RTS                             A:42 X:02 Y:F0 P:24 SP:F9
	// F3A6  4F 47 06 *SRE $0647 = 29                  A:42 X:02 Y:F0 P:24 SP:FB
	// F3A9  EA        NOP                             A:56 X:02 Y:F0 P:25 SP:FB
	// F3AA  EA        NOP                             A:56 X:02 Y:F0 P:25 SP:FB
	// F3AB  EA        NOP                             A:56 X:02 Y:F0 P:25 SP:FB
	// F3AC  EA        NOP                             A:56 X:02 Y:F0 P:25 SP:FB
	// F3AD  20 33 FB  JSR $FB33                       A:56 X:02 Y:F0 P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:02 Y:F0 P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:02 Y:F0 P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:02 Y:F0 P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:02 Y:F0 P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:02 Y:F0 P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:02 Y:F0 P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:02 Y:F0 P:27 SP:F9
	// F3B0  AD 47 06  LDA $0647 = 14                  A:56 X:02 Y:F0 P:27 SP:FB
	// F3B3  C9 14     CMP #$14                        A:14 X:02 Y:F0 P:25 SP:FB
	// F3B5  F0 02     BEQ $F3B9                       A:14 X:02 Y:F0 P:27 SP:FB
	// F3B9  C8        INY                             A:14 X:02 Y:F0 P:27 SP:FB
	// F3BA  A9 37     LDA #$37                        A:14 X:02 Y:F1 P:A5 SP:FB
	// F3BC  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:F1 P:25 SP:FB
	// F3BF  20 40 FB  JSR $FB40                       A:37 X:02 Y:F1 P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:02 Y:F1 P:25 SP:F9
	// FB42  38        SEC                             A:37 X:02 Y:F1 P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:02 Y:F1 P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:02 Y:F1 P:65 SP:F9
	// F3C2  4F 47 06 *SRE $0647 = 37                  A:75 X:02 Y:F1 P:65 SP:FB
	// F3C5  EA        NOP                             A:6E X:02 Y:F1 P:65 SP:FB
	// F3C6  EA        NOP                             A:6E X:02 Y:F1 P:65 SP:FB
	// F3C7  EA        NOP                             A:6E X:02 Y:F1 P:65 SP:FB
	// F3C8  EA        NOP                             A:6E X:02 Y:F1 P:65 SP:FB
	// F3C9  20 46 FB  JSR $FB46                       A:6E X:02 Y:F1 P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:02 Y:F1 P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:02 Y:F1 P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:02 Y:F1 P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:02 Y:F1 P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:02 Y:F1 P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:02 Y:F1 P:67 SP:F9
	// FB52  60        RTS                             A:6E X:02 Y:F1 P:67 SP:F9
	// F3CC  AD 47 06  LDA $0647 = 1B                  A:6E X:02 Y:F1 P:67 SP:FB
	// F3CF  C9 1B     CMP #$1B                        A:1B X:02 Y:F1 P:65 SP:FB
	// F3D1  F0 02     BEQ $F3D5                       A:1B X:02 Y:F1 P:67 SP:FB
	// F3D5  A9 A5     LDA #$A5                        A:1B X:02 Y:F1 P:67 SP:FB
	// F3D7  8D 47 06  STA $0647 = 1B                  A:A5 X:02 Y:F1 P:E5 SP:FB
	// F3DA  A9 48     LDA #$48                        A:A5 X:02 Y:F1 P:E5 SP:FB
	// F3DC  85 45     STA $45 = 48                    A:48 X:02 Y:F1 P:65 SP:FB
	// F3DE  A9 05     LDA #$05                        A:48 X:02 Y:F1 P:65 SP:FB
	// F3E0  85 46     STA $46 = 05                    A:05 X:02 Y:F1 P:65 SP:FB
	// F3E2  A0 FF     LDY #$FF                        A:05 X:02 Y:F1 P:65 SP:FB
	// F3E4  20 1D FB  JSR $FB1D                       A:05 X:02 Y:FF P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:E5 SP:F9
	// FB1F  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:05 X:02 Y:FF P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:02 Y:FF P:E4 SP:F9
	// F3E7  53 45    *SRE ($45),Y = 0548 @ 0647 = A5  A:B3 X:02 Y:FF P:E4 SP:FB
	// F3E9  EA        NOP                             A:E1 X:02 Y:FF P:E5 SP:FB
	// F3EA  EA        NOP                             A:E1 X:02 Y:FF P:E5 SP:FB
	// F3EB  08        PHP                             A:E1 X:02 Y:FF P:E5 SP:FB
	// F3EC  48        PHA                             A:E1 X:02 Y:FF P:E5 SP:FA
	// F3ED  A0 F2     LDY #$F2                        A:E1 X:02 Y:FF P:E5 SP:F9
	// F3EF  68        PLA                             A:E1 X:02 Y:F2 P:E5 SP:F9
	// F3F0  28        PLP                             A:E1 X:02 Y:F2 P:E5 SP:FA
	// F3F1  20 23 FB  JSR $FB23                       A:E1 X:02 Y:F2 P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:02 Y:F2 P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:02 Y:F2 P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:02 Y:F2 P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:02 Y:F2 P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:02 Y:F2 P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:02 Y:F2 P:67 SP:F9
	// F3F4  AD 47 06  LDA $0647 = 52                  A:E1 X:02 Y:F2 P:67 SP:FB
	// F3F7  C9 52     CMP #$52                        A:52 X:02 Y:F2 P:65 SP:FB
	// F3F9  F0 02     BEQ $F3FD                       A:52 X:02 Y:F2 P:67 SP:FB
	// F3FD  A0 FF     LDY #$FF                        A:52 X:02 Y:F2 P:67 SP:FB
	// F3FF  A9 29     LDA #$29                        A:52 X:02 Y:FF P:E5 SP:FB
	// F401  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:FF P:65 SP:FB
	// F404  20 2E FB  JSR $FB2E                       A:29 X:02 Y:FF P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:02 Y:FF P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:02 Y:FF P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:02 Y:FF P:24 SP:F9
	// FB32  60        RTS                             A:42 X:02 Y:FF P:24 SP:F9
	// F407  53 45    *SRE ($45),Y = 0548 @ 0647 = 29  A:42 X:02 Y:FF P:24 SP:FB
	// F409  EA        NOP                             A:56 X:02 Y:FF P:25 SP:FB
	// F40A  EA        NOP                             A:56 X:02 Y:FF P:25 SP:FB
	// F40B  08        PHP                             A:56 X:02 Y:FF P:25 SP:FB
	// F40C  48        PHA                             A:56 X:02 Y:FF P:25 SP:FA
	// F40D  A0 F3     LDY #$F3                        A:56 X:02 Y:FF P:25 SP:F9
	// F40F  68        PLA                             A:56 X:02 Y:F3 P:A5 SP:F9
	// F410  28        PLP                             A:56 X:02 Y:F3 P:25 SP:FA
	// F411  20 33 FB  JSR $FB33                       A:56 X:02 Y:F3 P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:02 Y:F3 P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:02 Y:F3 P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:02 Y:F3 P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:02 Y:F3 P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:02 Y:F3 P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:02 Y:F3 P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:02 Y:F3 P:27 SP:F9
	// F414  AD 47 06  LDA $0647 = 14                  A:56 X:02 Y:F3 P:27 SP:FB
	// F417  C9 14     CMP #$14                        A:14 X:02 Y:F3 P:25 SP:FB
	// F419  F0 02     BEQ $F41D                       A:14 X:02 Y:F3 P:27 SP:FB
	// F41D  A0 FF     LDY #$FF                        A:14 X:02 Y:F3 P:27 SP:FB
	// F41F  A9 37     LDA #$37                        A:14 X:02 Y:FF P:A5 SP:FB
	// F421  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:FF P:25 SP:FB
	// F424  20 40 FB  JSR $FB40                       A:37 X:02 Y:FF P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FB42  38        SEC                             A:37 X:02 Y:FF P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:02 Y:FF P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:02 Y:FF P:65 SP:F9
	// F427  53 45    *SRE ($45),Y = 0548 @ 0647 = 37  A:75 X:02 Y:FF P:65 SP:FB
	// F429  EA        NOP                             A:6E X:02 Y:FF P:65 SP:FB
	// F42A  EA        NOP                             A:6E X:02 Y:FF P:65 SP:FB
	// F42B  08        PHP                             A:6E X:02 Y:FF P:65 SP:FB
	// F42C  48        PHA                             A:6E X:02 Y:FF P:65 SP:FA
	// F42D  A0 F4     LDY #$F4                        A:6E X:02 Y:FF P:65 SP:F9
	// F42F  68        PLA                             A:6E X:02 Y:F4 P:E5 SP:F9
	// F430  28        PLP                             A:6E X:02 Y:F4 P:65 SP:FA
	// F431  20 46 FB  JSR $FB46                       A:6E X:02 Y:F4 P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:02 Y:F4 P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:02 Y:F4 P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:02 Y:F4 P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:02 Y:F4 P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:02 Y:F4 P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:02 Y:F4 P:67 SP:F9
	// FB52  60        RTS                             A:6E X:02 Y:F4 P:67 SP:F9
	// F434  AD 47 06  LDA $0647 = 1B                  A:6E X:02 Y:F4 P:67 SP:FB
	// F437  C9 1B     CMP #$1B                        A:1B X:02 Y:F4 P:65 SP:FB
	// F439  F0 02     BEQ $F43D                       A:1B X:02 Y:F4 P:67 SP:FB
	// F43D  A0 F5     LDY #$F5                        A:1B X:02 Y:F4 P:67 SP:FB
	// F43F  A2 FF     LDX #$FF                        A:1B X:02 Y:F5 P:E5 SP:FB
	// F441  A9 A5     LDA #$A5                        A:1B X:FF Y:F5 P:E5 SP:FB
	// F443  85 47     STA $47 = 1B                    A:A5 X:FF Y:F5 P:E5 SP:FB
	// F445  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:F5 P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:F5 P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:FF Y:F5 P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:F5 P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:FF Y:F5 P:E4 SP:F9
	// F448  57 48    *SRE $48,X @ 47 = A5             A:B3 X:FF Y:F5 P:E4 SP:FB
	// F44A  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB
	// F44B  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB
	// F44C  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB
	// F44D  EA        NOP                             A:E1 X:FF Y:F5 P:E5 SP:FB
	// F44E  20 23 FB  JSR $FB23                       A:E1 X:FF Y:F5 P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:F5 P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:F5 P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:F5 P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:FF Y:F5 P:67 SP:F9
	// F451  A5 47     LDA $47 = 52                    A:E1 X:FF Y:F5 P:67 SP:FB
	// F453  C9 52     CMP #$52                        A:52 X:FF Y:F5 P:65 SP:FB
	// F455  F0 02     BEQ $F459                       A:52 X:FF Y:F5 P:67 SP:FB
	// F459  C8        INY                             A:52 X:FF Y:F5 P:67 SP:FB
	// F45A  A9 29     LDA #$29                        A:52 X:FF Y:F6 P:E5 SP:FB
	// F45C  85 47     STA $47 = 52                    A:29 X:FF Y:F6 P:65 SP:FB
	// F45E  20 2E FB  JSR $FB2E                       A:29 X:FF Y:F6 P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:FF Y:F6 P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:FF Y:F6 P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:FF Y:F6 P:24 SP:F9
	// FB32  60        RTS                             A:42 X:FF Y:F6 P:24 SP:F9
	// F461  57 48    *SRE $48,X @ 47 = 29             A:42 X:FF Y:F6 P:24 SP:FB
	// F463  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB
	// F464  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB
	// F465  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB
	// F466  EA        NOP                             A:56 X:FF Y:F6 P:25 SP:FB
	// F467  20 33 FB  JSR $FB33                       A:56 X:FF Y:F6 P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:FF Y:F6 P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:F6 P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:FF Y:F6 P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:FF Y:F6 P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:FF Y:F6 P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:F6 P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:FF Y:F6 P:27 SP:F9
	// F46A  A5 47     LDA $47 = 14                    A:56 X:FF Y:F6 P:27 SP:FB
	// F46C  C9 14     CMP #$14                        A:14 X:FF Y:F6 P:25 SP:FB
	// F46E  F0 02     BEQ $F472                       A:14 X:FF Y:F6 P:27 SP:FB
	// F472  C8        INY                             A:14 X:FF Y:F6 P:27 SP:FB
	// F473  A9 37     LDA #$37                        A:14 X:FF Y:F7 P:A5 SP:FB
	// F475  85 47     STA $47 = 14                    A:37 X:FF Y:F7 P:25 SP:FB
	// F477  20 40 FB  JSR $FB40                       A:37 X:FF Y:F7 P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:F7 P:25 SP:F9
	// FB42  38        SEC                             A:37 X:FF Y:F7 P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:FF Y:F7 P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:FF Y:F7 P:65 SP:F9
	// F47A  57 48    *SRE $48,X @ 47 = 37             A:75 X:FF Y:F7 P:65 SP:FB
	// F47C  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB
	// F47D  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB
	// F47E  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB
	// F47F  EA        NOP                             A:6E X:FF Y:F7 P:65 SP:FB
	// F480  20 46 FB  JSR $FB46                       A:6E X:FF Y:F7 P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:FF Y:F7 P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:F7 P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:FF Y:F7 P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:FF Y:F7 P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:F7 P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:FF Y:F7 P:67 SP:F9
	// FB52  60        RTS                             A:6E X:FF Y:F7 P:67 SP:F9
	// F483  A5 47     LDA $47 = 1B                    A:6E X:FF Y:F7 P:67 SP:FB
	// F485  C9 1B     CMP #$1B                        A:1B X:FF Y:F7 P:65 SP:FB
	// F487  F0 02     BEQ $F48B                       A:1B X:FF Y:F7 P:67 SP:FB
	// F48B  A9 A5     LDA #$A5                        A:1B X:FF Y:F7 P:67 SP:FB
	// F48D  8D 47 06  STA $0647 = 1B                  A:A5 X:FF Y:F7 P:E5 SP:FB
	// F490  A0 FF     LDY #$FF                        A:A5 X:FF Y:F7 P:E5 SP:FB
	// F492  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:FF P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:FF P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:FF Y:FF P:E4 SP:F9
	// F495  5B 48 05 *SRE $0548,Y @ 0647 = A5         A:B3 X:FF Y:FF P:E4 SP:FB
	// F498  EA        NOP                             A:E1 X:FF Y:FF P:E5 SP:FB
	// F499  EA        NOP                             A:E1 X:FF Y:FF P:E5 SP:FB
	// F49A  08        PHP                             A:E1 X:FF Y:FF P:E5 SP:FB
	// F49B  48        PHA                             A:E1 X:FF Y:FF P:E5 SP:FA
	// F49C  A0 F8     LDY #$F8                        A:E1 X:FF Y:FF P:E5 SP:F9
	// F49E  68        PLA                             A:E1 X:FF Y:F8 P:E5 SP:F9
	// F49F  28        PLP                             A:E1 X:FF Y:F8 P:E5 SP:FA
	// F4A0  20 23 FB  JSR $FB23                       A:E1 X:FF Y:F8 P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:F8 P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:F8 P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:F8 P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:FF Y:F8 P:67 SP:F9
	// F4A3  AD 47 06  LDA $0647 = 52                  A:E1 X:FF Y:F8 P:67 SP:FB
	// F4A6  C9 52     CMP #$52                        A:52 X:FF Y:F8 P:65 SP:FB
	// F4A8  F0 02     BEQ $F4AC                       A:52 X:FF Y:F8 P:67 SP:FB
	// F4AC  A0 FF     LDY #$FF                        A:52 X:FF Y:F8 P:67 SP:FB
	// F4AE  A9 29     LDA #$29                        A:52 X:FF Y:FF P:E5 SP:FB
	// F4B0  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FF P:65 SP:FB
	// F4B3  20 2E FB  JSR $FB2E                       A:29 X:FF Y:FF P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:FF Y:FF P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9
	// FB32  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9
	// F4B6  5B 48 05 *SRE $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB
	// F4B9  EA        NOP                             A:56 X:FF Y:FF P:25 SP:FB
	// F4BA  EA        NOP                             A:56 X:FF Y:FF P:25 SP:FB
	// F4BB  08        PHP                             A:56 X:FF Y:FF P:25 SP:FB
	// F4BC  48        PHA                             A:56 X:FF Y:FF P:25 SP:FA
	// F4BD  A0 F9     LDY #$F9                        A:56 X:FF Y:FF P:25 SP:F9
	// F4BF  68        PLA                             A:56 X:FF Y:F9 P:A5 SP:F9
	// F4C0  28        PLP                             A:56 X:FF Y:F9 P:25 SP:FA
	// F4C1  20 33 FB  JSR $FB33                       A:56 X:FF Y:F9 P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:FF Y:F9 P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:F9 P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:FF Y:F9 P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:FF Y:F9 P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:FF Y:F9 P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:F9 P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:FF Y:F9 P:27 SP:F9
	// F4C4  AD 47 06  LDA $0647 = 14                  A:56 X:FF Y:F9 P:27 SP:FB
	// F4C7  C9 14     CMP #$14                        A:14 X:FF Y:F9 P:25 SP:FB
	// F4C9  F0 02     BEQ $F4CD                       A:14 X:FF Y:F9 P:27 SP:FB
	// F4CD  A0 FF     LDY #$FF                        A:14 X:FF Y:F9 P:27 SP:FB
	// F4CF  A9 37     LDA #$37                        A:14 X:FF Y:FF P:A5 SP:FB
	// F4D1  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FF P:25 SP:FB
	// F4D4  20 40 FB  JSR $FB40                       A:37 X:FF Y:FF P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FB42  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9
	// F4D7  5B 48 05 *SRE $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB
	// F4DA  EA        NOP                             A:6E X:FF Y:FF P:65 SP:FB
	// F4DB  EA        NOP                             A:6E X:FF Y:FF P:65 SP:FB
	// F4DC  08        PHP                             A:6E X:FF Y:FF P:65 SP:FB
	// F4DD  48        PHA                             A:6E X:FF Y:FF P:65 SP:FA
	// F4DE  A0 FA     LDY #$FA                        A:6E X:FF Y:FF P:65 SP:F9
	// F4E0  68        PLA                             A:6E X:FF Y:FA P:E5 SP:F9
	// F4E1  28        PLP                             A:6E X:FF Y:FA P:65 SP:FA
	// F4E2  20 46 FB  JSR $FB46                       A:6E X:FF Y:FA P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:FF Y:FA P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:FA P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:FF Y:FA P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:FF Y:FA P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:FA P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:FF Y:FA P:67 SP:F9
	// FB52  60        RTS                             A:6E X:FF Y:FA P:67 SP:F9
	// F4E5  AD 47 06  LDA $0647 = 1B                  A:6E X:FF Y:FA P:67 SP:FB
	// F4E8  C9 1B     CMP #$1B                        A:1B X:FF Y:FA P:65 SP:FB
	// F4EA  F0 02     BEQ $F4EE                       A:1B X:FF Y:FA P:67 SP:FB
	// F4EE  A0 FB     LDY #$FB                        A:1B X:FF Y:FA P:67 SP:FB
	// F4F0  A2 FF     LDX #$FF                        A:1B X:FF Y:FB P:E5 SP:FB
	// F4F2  A9 A5     LDA #$A5                        A:1B X:FF Y:FB P:E5 SP:FB
	// F4F4  8D 47 06  STA $0647 = 1B                  A:A5 X:FF Y:FB P:E5 SP:FB
	// F4F7  20 1D FB  JSR $FB1D                       A:A5 X:FF Y:FB P:E5 SP:FB
	// FB1D  24 01     BIT $01 = FF                    A:A5 X:FF Y:FB P:E5 SP:F9
	// FB1F  18        CLC                             A:A5 X:FF Y:FB P:E5 SP:F9
	// FB20  A9 B3     LDA #$B3                        A:A5 X:FF Y:FB P:E4 SP:F9
	// FB22  60        RTS                             A:B3 X:FF Y:FB P:E4 SP:F9
	// F4FA  5F 48 05 *SRE $0548,X @ 0647 = A5         A:B3 X:FF Y:FB P:E4 SP:FB
	// F4FD  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB
	// F4FE  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB
	// F4FF  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB
	// F500  EA        NOP                             A:E1 X:FF Y:FB P:E5 SP:FB
	// F501  20 23 FB  JSR $FB23                       A:E1 X:FF Y:FB P:E5 SP:FB
	// FB23  50 50     BVC $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9
	// FB25  90 4E     BCC $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9
	// FB27  10 4C     BPL $FB75                       A:E1 X:FF Y:FB P:E5 SP:F9
	// FB29  C9 E1     CMP #$E1                        A:E1 X:FF Y:FB P:E5 SP:F9
	// FB2B  D0 48     BNE $FB75                       A:E1 X:FF Y:FB P:67 SP:F9
	// FB2D  60        RTS                             A:E1 X:FF Y:FB P:67 SP:F9
	// F504  AD 47 06  LDA $0647 = 52                  A:E1 X:FF Y:FB P:67 SP:FB
	// F507  C9 52     CMP #$52                        A:52 X:FF Y:FB P:65 SP:FB
	// F509  F0 02     BEQ $F50D                       A:52 X:FF Y:FB P:67 SP:FB
	// F50D  C8        INY                             A:52 X:FF Y:FB P:67 SP:FB
	// F50E  A9 29     LDA #$29                        A:52 X:FF Y:FC P:E5 SP:FB
	// F510  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FC P:65 SP:FB
	// F513  20 2E FB  JSR $FB2E                       A:29 X:FF Y:FC P:65 SP:FB
	// FB2E  B8        CLV                             A:29 X:FF Y:FC P:65 SP:F9
	// FB2F  18        CLC                             A:29 X:FF Y:FC P:25 SP:F9
	// FB30  A9 42     LDA #$42                        A:29 X:FF Y:FC P:24 SP:F9
	// FB32  60        RTS                             A:42 X:FF Y:FC P:24 SP:F9
	// F516  5F 48 05 *SRE $0548,X @ 0647 = 29         A:42 X:FF Y:FC P:24 SP:FB
	// F519  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB
	// F51A  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB
	// F51B  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB
	// F51C  EA        NOP                             A:56 X:FF Y:FC P:25 SP:FB
	// F51D  20 33 FB  JSR $FB33                       A:56 X:FF Y:FC P:25 SP:FB
	// FB33  70 40     BVS $FB75                       A:56 X:FF Y:FC P:25 SP:F9
	// FB35  F0 3E     BEQ $FB75                       A:56 X:FF Y:FC P:25 SP:F9
	// FB37  30 3C     BMI $FB75                       A:56 X:FF Y:FC P:25 SP:F9
	// FB39  90 3A     BCC $FB75                       A:56 X:FF Y:FC P:25 SP:F9
	// FB3B  C9 56     CMP #$56                        A:56 X:FF Y:FC P:25 SP:F9
	// FB3D  D0 36     BNE $FB75                       A:56 X:FF Y:FC P:27 SP:F9
	// FB3F  60        RTS                             A:56 X:FF Y:FC P:27 SP:F9
	// F520  AD 47 06  LDA $0647 = 14                  A:56 X:FF Y:FC P:27 SP:FB
	// F523  C9 14     CMP #$14                        A:14 X:FF Y:FC P:25 SP:FB
	// F525  F0 02     BEQ $F529                       A:14 X:FF Y:FC P:27 SP:FB
	// F529  C8        INY                             A:14 X:FF Y:FC P:27 SP:FB
	// F52A  A9 37     LDA #$37                        A:14 X:FF Y:FD P:A5 SP:FB
	// F52C  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FD P:25 SP:FB
	// F52F  20 40 FB  JSR $FB40                       A:37 X:FF Y:FD P:25 SP:FB
	// FB40  24 01     BIT $01 = FF                    A:37 X:FF Y:FD P:25 SP:F9
	// FB42  38        SEC                             A:37 X:FF Y:FD P:E5 SP:F9
	// FB43  A9 75     LDA #$75                        A:37 X:FF Y:FD P:E5 SP:F9
	// FB45  60        RTS                             A:75 X:FF Y:FD P:65 SP:F9
	// F532  5F 48 05 *SRE $0548,X @ 0647 = 37         A:75 X:FF Y:FD P:65 SP:FB
	// F535  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB
	// F536  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB
	// F537  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB
	// F538  EA        NOP                             A:6E X:FF Y:FD P:65 SP:FB
	// F539  20 46 FB  JSR $FB46                       A:6E X:FF Y:FD P:65 SP:FB
	// FB46  50 2D     BVC $FB75                       A:6E X:FF Y:FD P:65 SP:F9
	// FB48  F0 2B     BEQ $FB75                       A:6E X:FF Y:FD P:65 SP:F9
	// FB4A  30 29     BMI $FB75                       A:6E X:FF Y:FD P:65 SP:F9
	// FB4C  90 27     BCC $FB75                       A:6E X:FF Y:FD P:65 SP:F9
	// FB4E  C9 6E     CMP #$6E                        A:6E X:FF Y:FD P:65 SP:F9
	// FB50  D0 23     BNE $FB75                       A:6E X:FF Y:FD P:67 SP:F9
	// FB52  60        RTS                             A:6E X:FF Y:FD P:67 SP:F9
	// F53C  AD 47 06  LDA $0647 = 1B                  A:6E X:FF Y:FD P:67 SP:FB
	// F53F  C9 1B     CMP #$1B                        A:1B X:FF Y:FD P:65 SP:FB
	// F541  F0 02     BEQ $F545                       A:1B X:FF Y:FD P:67 SP:FB
	// F545  60        RTS                             A:1B X:FF Y:FD P:67 SP:FB
	// C64A  A5 00     LDA $00 = 00                    A:1B X:FF Y:FD P:67 SP:FD
	// C64C  85 11     STA $11 = 00                    A:00 X:FF Y:FD P:67 SP:FD
	// C64E  A9 00     LDA #$00                        A:00 X:FF Y:FD P:67 SP:FD
	// C650  85 00     STA $00 = 00                    A:00 X:FF Y:FD P:67 SP:FD
	// C652  20 46 F5  JSR $F546                       A:00 X:FF Y:FD P:67 SP:FD
	// F546  A9 FF     LDA #$FF                        A:00 X:FF Y:FD P:67 SP:FB
	// F548  85 01     STA $01 = FF                    A:FF X:FF Y:FD P:E5 SP:FB
	// F54A  A0 01     LDY #$01                        A:FF X:FF Y:FD P:E5 SP:FB
	// F54C  A2 02     LDX #$02                        A:FF X:FF Y:01 P:65 SP:FB
	// F54E  A9 47     LDA #$47                        A:FF X:02 Y:01 P:65 SP:FB
	// F550  85 47     STA $47 = 1B                    A:47 X:02 Y:01 P:65 SP:FB
	// F552  A9 06     LDA #$06                        A:47 X:02 Y:01 P:65 SP:FB
	// F554  85 48     STA $48 = 06                    A:06 X:02 Y:01 P:65 SP:FB
	// F556  A9 A5     LDA #$A5                        A:06 X:02 Y:01 P:65 SP:FB
	// F558  8D 47 06  STA $0647 = 1B                  A:A5 X:02 Y:01 P:E5 SP:FB
	// F55B  20 E9 FA  JSR $FAE9                       A:A5 X:02 Y:01 P:E5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:02 Y:01 P:E5 SP:F9
	// FAEB  18        CLC                             A:A5 X:02 Y:01 P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:02 Y:01 P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:02 Y:01 P:E4 SP:F9
	// F55E  63 45    *RRA ($45,X) @ 47 = 0647 = A5    A:B2 X:02 Y:01 P:E4 SP:FB
	// F560  EA        NOP                             A:05 X:02 Y:01 P:25 SP:FB
	// F561  EA        NOP                             A:05 X:02 Y:01 P:25 SP:FB
	// F562  EA        NOP                             A:05 X:02 Y:01 P:25 SP:FB
	// F563  EA        NOP                             A:05 X:02 Y:01 P:25 SP:FB
	// F564  20 EF FA  JSR $FAEF                       A:05 X:02 Y:01 P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:02 Y:01 P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:02 Y:01 P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:02 Y:01 P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:02 Y:01 P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:02 Y:01 P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:02 Y:01 P:27 SP:F9
	// F567  AD 47 06  LDA $0647 = 52                  A:05 X:02 Y:01 P:27 SP:FB
	// F56A  C9 52     CMP #$52                        A:52 X:02 Y:01 P:25 SP:FB
	// F56C  F0 02     BEQ $F570                       A:52 X:02 Y:01 P:27 SP:FB
	// F570  C8        INY                             A:52 X:02 Y:01 P:27 SP:FB
	// F571  A9 29     LDA #$29                        A:52 X:02 Y:02 P:25 SP:FB
	// F573  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:02 P:25 SP:FB
	// F576  20 FA FA  JSR $FAFA                       A:29 X:02 Y:02 P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:02 Y:02 P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:02 Y:02 P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:02 Y:02 P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:02 Y:02 P:24 SP:F9
	// F579  63 45    *RRA ($45,X) @ 47 = 0647 = 29    A:42 X:02 Y:02 P:24 SP:FB
	// F57B  EA        NOP                             A:57 X:02 Y:02 P:24 SP:FB
	// F57C  EA        NOP                             A:57 X:02 Y:02 P:24 SP:FB
	// F57D  EA        NOP                             A:57 X:02 Y:02 P:24 SP:FB
	// F57E  EA        NOP                             A:57 X:02 Y:02 P:24 SP:FB
	// F57F  20 FF FA  JSR $FAFF                       A:57 X:02 Y:02 P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:02 Y:02 P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:02 Y:02 P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:02 Y:02 P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:02 Y:02 P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:02 Y:02 P:27 SP:F9
	// FB09  60        RTS                             A:57 X:02 Y:02 P:27 SP:F9
	// F582  AD 47 06  LDA $0647 = 14                  A:57 X:02 Y:02 P:27 SP:FB
	// F585  C9 14     CMP #$14                        A:14 X:02 Y:02 P:25 SP:FB
	// F587  F0 02     BEQ $F58B                       A:14 X:02 Y:02 P:27 SP:FB
	// F58B  C8        INY                             A:14 X:02 Y:02 P:27 SP:FB
	// F58C  A9 37     LDA #$37                        A:14 X:02 Y:03 P:25 SP:FB
	// F58E  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:03 P:25 SP:FB
	// F591  20 0A FB  JSR $FB0A                       A:37 X:02 Y:03 P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:02 Y:03 P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:02 Y:03 P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:02 Y:03 P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:02 Y:03 P:65 SP:F9
	// F594  63 45    *RRA ($45,X) @ 47 = 0647 = 37    A:75 X:02 Y:03 P:65 SP:FB
	// F596  EA        NOP                             A:11 X:02 Y:03 P:25 SP:FB
	// F597  EA        NOP                             A:11 X:02 Y:03 P:25 SP:FB
	// F598  EA        NOP                             A:11 X:02 Y:03 P:25 SP:FB
	// F599  EA        NOP                             A:11 X:02 Y:03 P:25 SP:FB
	// F59A  20 10 FB  JSR $FB10                       A:11 X:02 Y:03 P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:02 Y:03 P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:02 Y:03 P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:02 Y:03 P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:02 Y:03 P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:02 Y:03 P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:02 Y:03 P:27 SP:F9
	// F59D  AD 47 06  LDA $0647 = 9B                  A:11 X:02 Y:03 P:27 SP:FB
	// F5A0  C9 9B     CMP #$9B                        A:9B X:02 Y:03 P:A5 SP:FB
	// F5A2  F0 02     BEQ $F5A6                       A:9B X:02 Y:03 P:27 SP:FB
	// F5A6  C8        INY                             A:9B X:02 Y:03 P:27 SP:FB
	// F5A7  A9 A5     LDA #$A5                        A:9B X:02 Y:04 P:25 SP:FB
	// F5A9  85 47     STA $47 = 47                    A:A5 X:02 Y:04 P:A5 SP:FB
	// F5AB  20 E9 FA  JSR $FAE9                       A:A5 X:02 Y:04 P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:02 Y:04 P:A5 SP:F9
	// FAEB  18        CLC                             A:A5 X:02 Y:04 P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:02 Y:04 P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:02 Y:04 P:E4 SP:F9
	// F5AE  67 47    *RRA $47 = A5                    A:B2 X:02 Y:04 P:E4 SP:FB
	// F5B0  EA        NOP                             A:05 X:02 Y:04 P:25 SP:FB
	// F5B1  EA        NOP                             A:05 X:02 Y:04 P:25 SP:FB
	// F5B2  EA        NOP                             A:05 X:02 Y:04 P:25 SP:FB
	// F5B3  EA        NOP                             A:05 X:02 Y:04 P:25 SP:FB
	// F5B4  20 EF FA  JSR $FAEF                       A:05 X:02 Y:04 P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:02 Y:04 P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:02 Y:04 P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:02 Y:04 P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:02 Y:04 P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:02 Y:04 P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:02 Y:04 P:27 SP:F9
	// F5B7  A5 47     LDA $47 = 52                    A:05 X:02 Y:04 P:27 SP:FB
	// F5B9  C9 52     CMP #$52                        A:52 X:02 Y:04 P:25 SP:FB
	// F5BB  F0 02     BEQ $F5BF                       A:52 X:02 Y:04 P:27 SP:FB
	// F5BF  C8        INY                             A:52 X:02 Y:04 P:27 SP:FB
	// F5C0  A9 29     LDA #$29                        A:52 X:02 Y:05 P:25 SP:FB
	// F5C2  85 47     STA $47 = 52                    A:29 X:02 Y:05 P:25 SP:FB
	// F5C4  20 FA FA  JSR $FAFA                       A:29 X:02 Y:05 P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:02 Y:05 P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:02 Y:05 P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:02 Y:05 P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:02 Y:05 P:24 SP:F9
	// F5C7  67 47    *RRA $47 = 29                    A:42 X:02 Y:05 P:24 SP:FB
	// F5C9  EA        NOP                             A:57 X:02 Y:05 P:24 SP:FB
	// F5CA  EA        NOP                             A:57 X:02 Y:05 P:24 SP:FB
	// F5CB  EA        NOP                             A:57 X:02 Y:05 P:24 SP:FB
	// F5CC  EA        NOP                             A:57 X:02 Y:05 P:24 SP:FB
	// F5CD  20 FF FA  JSR $FAFF                       A:57 X:02 Y:05 P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:02 Y:05 P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:02 Y:05 P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:02 Y:05 P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:02 Y:05 P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:02 Y:05 P:27 SP:F9
	// FB09  60        RTS                             A:57 X:02 Y:05 P:27 SP:F9
	// F5D0  A5 47     LDA $47 = 14                    A:57 X:02 Y:05 P:27 SP:FB
	// F5D2  C9 14     CMP #$14                        A:14 X:02 Y:05 P:25 SP:FB
	// F5D4  F0 02     BEQ $F5D8                       A:14 X:02 Y:05 P:27 SP:FB
	// F5D8  C8        INY                             A:14 X:02 Y:05 P:27 SP:FB
	// F5D9  A9 37     LDA #$37                        A:14 X:02 Y:06 P:25 SP:FB
	// F5DB  85 47     STA $47 = 14                    A:37 X:02 Y:06 P:25 SP:FB
	// F5DD  20 0A FB  JSR $FB0A                       A:37 X:02 Y:06 P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:02 Y:06 P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:02 Y:06 P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:02 Y:06 P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:02 Y:06 P:65 SP:F9
	// F5E0  67 47    *RRA $47 = 37                    A:75 X:02 Y:06 P:65 SP:FB
	// F5E2  EA        NOP                             A:11 X:02 Y:06 P:25 SP:FB
	// F5E3  EA        NOP                             A:11 X:02 Y:06 P:25 SP:FB
	// F5E4  EA        NOP                             A:11 X:02 Y:06 P:25 SP:FB
	// F5E5  EA        NOP                             A:11 X:02 Y:06 P:25 SP:FB
	// F5E6  20 10 FB  JSR $FB10                       A:11 X:02 Y:06 P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:02 Y:06 P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:02 Y:06 P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:02 Y:06 P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:02 Y:06 P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:02 Y:06 P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:02 Y:06 P:27 SP:F9
	// F5E9  A5 47     LDA $47 = 9B                    A:11 X:02 Y:06 P:27 SP:FB
	// F5EB  C9 9B     CMP #$9B                        A:9B X:02 Y:06 P:A5 SP:FB
	// F5ED  F0 02     BEQ $F5F1                       A:9B X:02 Y:06 P:27 SP:FB
	// F5F1  C8        INY                             A:9B X:02 Y:06 P:27 SP:FB
	// F5F2  A9 A5     LDA #$A5                        A:9B X:02 Y:07 P:25 SP:FB
	// F5F4  8D 47 06  STA $0647 = 9B                  A:A5 X:02 Y:07 P:A5 SP:FB
	// F5F7  20 E9 FA  JSR $FAE9                       A:A5 X:02 Y:07 P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:02 Y:07 P:A5 SP:F9
	// FAEB  18        CLC                             A:A5 X:02 Y:07 P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:02 Y:07 P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:02 Y:07 P:E4 SP:F9
	// F5FA  6F 47 06 *RRA $0647 = A5                  A:B2 X:02 Y:07 P:E4 SP:FB
	// F5FD  EA        NOP                             A:05 X:02 Y:07 P:25 SP:FB
	// F5FE  EA        NOP                             A:05 X:02 Y:07 P:25 SP:FB
	// F5FF  EA        NOP                             A:05 X:02 Y:07 P:25 SP:FB
	// F600  EA        NOP                             A:05 X:02 Y:07 P:25 SP:FB
	// F601  20 EF FA  JSR $FAEF                       A:05 X:02 Y:07 P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:02 Y:07 P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:02 Y:07 P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:02 Y:07 P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:02 Y:07 P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:02 Y:07 P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:02 Y:07 P:27 SP:F9
	// F604  AD 47 06  LDA $0647 = 52                  A:05 X:02 Y:07 P:27 SP:FB
	// F607  C9 52     CMP #$52                        A:52 X:02 Y:07 P:25 SP:FB
	// F609  F0 02     BEQ $F60D                       A:52 X:02 Y:07 P:27 SP:FB
	// F60D  C8        INY                             A:52 X:02 Y:07 P:27 SP:FB
	// F60E  A9 29     LDA #$29                        A:52 X:02 Y:08 P:25 SP:FB
	// F610  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:08 P:25 SP:FB
	// F613  20 FA FA  JSR $FAFA                       A:29 X:02 Y:08 P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:02 Y:08 P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:02 Y:08 P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:02 Y:08 P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:02 Y:08 P:24 SP:F9
	// F616  6F 47 06 *RRA $0647 = 29                  A:42 X:02 Y:08 P:24 SP:FB
	// F619  EA        NOP                             A:57 X:02 Y:08 P:24 SP:FB
	// F61A  EA        NOP                             A:57 X:02 Y:08 P:24 SP:FB
	// F61B  EA        NOP                             A:57 X:02 Y:08 P:24 SP:FB
	// F61C  EA        NOP                             A:57 X:02 Y:08 P:24 SP:FB
	// F61D  20 FF FA  JSR $FAFF                       A:57 X:02 Y:08 P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:02 Y:08 P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:02 Y:08 P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:02 Y:08 P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:02 Y:08 P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:02 Y:08 P:27 SP:F9
	// FB09  60        RTS                             A:57 X:02 Y:08 P:27 SP:F9
	// F620  AD 47 06  LDA $0647 = 14                  A:57 X:02 Y:08 P:27 SP:FB
	// F623  C9 14     CMP #$14                        A:14 X:02 Y:08 P:25 SP:FB
	// F625  F0 02     BEQ $F629                       A:14 X:02 Y:08 P:27 SP:FB
	// F629  C8        INY                             A:14 X:02 Y:08 P:27 SP:FB
	// F62A  A9 37     LDA #$37                        A:14 X:02 Y:09 P:25 SP:FB
	// F62C  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:09 P:25 SP:FB
	// F62F  20 0A FB  JSR $FB0A                       A:37 X:02 Y:09 P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:02 Y:09 P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:02 Y:09 P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:02 Y:09 P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:02 Y:09 P:65 SP:F9
	// F632  6F 47 06 *RRA $0647 = 37                  A:75 X:02 Y:09 P:65 SP:FB
	// F635  EA        NOP                             A:11 X:02 Y:09 P:25 SP:FB
	// F636  EA        NOP                             A:11 X:02 Y:09 P:25 SP:FB
	// F637  EA        NOP                             A:11 X:02 Y:09 P:25 SP:FB
	// F638  EA        NOP                             A:11 X:02 Y:09 P:25 SP:FB
	// F639  20 10 FB  JSR $FB10                       A:11 X:02 Y:09 P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:02 Y:09 P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:02 Y:09 P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:02 Y:09 P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:02 Y:09 P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:02 Y:09 P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:02 Y:09 P:27 SP:F9
	// F63C  AD 47 06  LDA $0647 = 9B                  A:11 X:02 Y:09 P:27 SP:FB
	// F63F  C9 9B     CMP #$9B                        A:9B X:02 Y:09 P:A5 SP:FB
	// F641  F0 02     BEQ $F645                       A:9B X:02 Y:09 P:27 SP:FB
	// F645  A9 A5     LDA #$A5                        A:9B X:02 Y:09 P:27 SP:FB
	// F647  8D 47 06  STA $0647 = 9B                  A:A5 X:02 Y:09 P:A5 SP:FB
	// F64A  A9 48     LDA #$48                        A:A5 X:02 Y:09 P:A5 SP:FB
	// F64C  85 45     STA $45 = 48                    A:48 X:02 Y:09 P:25 SP:FB
	// F64E  A9 05     LDA #$05                        A:48 X:02 Y:09 P:25 SP:FB
	// F650  85 46     STA $46 = 05                    A:05 X:02 Y:09 P:25 SP:FB
	// F652  A0 FF     LDY #$FF                        A:05 X:02 Y:09 P:25 SP:FB
	// F654  20 E9 FA  JSR $FAE9                       A:05 X:02 Y:FF P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:05 X:02 Y:FF P:A5 SP:F9
	// FAEB  18        CLC                             A:05 X:02 Y:FF P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:05 X:02 Y:FF P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:02 Y:FF P:E4 SP:F9
	// F657  73 45    *RRA ($45),Y = 0548 @ 0647 = A5  A:B2 X:02 Y:FF P:E4 SP:FB
	// F659  EA        NOP                             A:05 X:02 Y:FF P:25 SP:FB
	// F65A  EA        NOP                             A:05 X:02 Y:FF P:25 SP:FB
	// F65B  08        PHP                             A:05 X:02 Y:FF P:25 SP:FB
	// F65C  48        PHA                             A:05 X:02 Y:FF P:25 SP:FA
	// F65D  A0 0A     LDY #$0A                        A:05 X:02 Y:FF P:25 SP:F9
	// F65F  68        PLA                             A:05 X:02 Y:0A P:25 SP:F9
	// F660  28        PLP                             A:05 X:02 Y:0A P:25 SP:FA
	// F661  20 EF FA  JSR $FAEF                       A:05 X:02 Y:0A P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:02 Y:0A P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:02 Y:0A P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:02 Y:0A P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:02 Y:0A P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:02 Y:0A P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:02 Y:0A P:27 SP:F9
	// F664  AD 47 06  LDA $0647 = 52                  A:05 X:02 Y:0A P:27 SP:FB
	// F667  C9 52     CMP #$52                        A:52 X:02 Y:0A P:25 SP:FB
	// F669  F0 02     BEQ $F66D                       A:52 X:02 Y:0A P:27 SP:FB
	// F66D  A0 FF     LDY #$FF                        A:52 X:02 Y:0A P:27 SP:FB
	// F66F  A9 29     LDA #$29                        A:52 X:02 Y:FF P:A5 SP:FB
	// F671  8D 47 06  STA $0647 = 52                  A:29 X:02 Y:FF P:25 SP:FB
	// F674  20 FA FA  JSR $FAFA                       A:29 X:02 Y:FF P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:02 Y:FF P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:02 Y:FF P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:02 Y:FF P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:02 Y:FF P:24 SP:F9
	// F677  73 45    *RRA ($45),Y = 0548 @ 0647 = 29  A:42 X:02 Y:FF P:24 SP:FB
	// F679  EA        NOP                             A:57 X:02 Y:FF P:24 SP:FB
	// F67A  EA        NOP                             A:57 X:02 Y:FF P:24 SP:FB
	// F67B  08        PHP                             A:57 X:02 Y:FF P:24 SP:FB
	// F67C  48        PHA                             A:57 X:02 Y:FF P:24 SP:FA
	// F67D  A0 0B     LDY #$0B                        A:57 X:02 Y:FF P:24 SP:F9
	// F67F  68        PLA                             A:57 X:02 Y:0B P:24 SP:F9
	// F680  28        PLP                             A:57 X:02 Y:0B P:24 SP:FA
	// F681  20 FF FA  JSR $FAFF                       A:57 X:02 Y:0B P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:02 Y:0B P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:02 Y:0B P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:02 Y:0B P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:02 Y:0B P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:02 Y:0B P:27 SP:F9
	// FB09  60        RTS                             A:57 X:02 Y:0B P:27 SP:F9
	// F684  AD 47 06  LDA $0647 = 14                  A:57 X:02 Y:0B P:27 SP:FB
	// F687  C9 14     CMP #$14                        A:14 X:02 Y:0B P:25 SP:FB
	// F689  F0 02     BEQ $F68D                       A:14 X:02 Y:0B P:27 SP:FB
	// F68D  A0 FF     LDY #$FF                        A:14 X:02 Y:0B P:27 SP:FB
	// F68F  A9 37     LDA #$37                        A:14 X:02 Y:FF P:A5 SP:FB
	// F691  8D 47 06  STA $0647 = 14                  A:37 X:02 Y:FF P:25 SP:FB
	// F694  20 0A FB  JSR $FB0A                       A:37 X:02 Y:FF P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:02 Y:FF P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:02 Y:FF P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:02 Y:FF P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:02 Y:FF P:65 SP:F9
	// F697  73 45    *RRA ($45),Y = 0548 @ 0647 = 37  A:75 X:02 Y:FF P:65 SP:FB
	// F699  EA        NOP                             A:11 X:02 Y:FF P:25 SP:FB
	// F69A  EA        NOP                             A:11 X:02 Y:FF P:25 SP:FB
	// F69B  08        PHP                             A:11 X:02 Y:FF P:25 SP:FB
	// F69C  48        PHA                             A:11 X:02 Y:FF P:25 SP:FA
	// F69D  A0 0C     LDY #$0C                        A:11 X:02 Y:FF P:25 SP:F9
	// F69F  68        PLA                             A:11 X:02 Y:0C P:25 SP:F9
	// F6A0  28        PLP                             A:11 X:02 Y:0C P:25 SP:FA
	// F6A1  20 10 FB  JSR $FB10                       A:11 X:02 Y:0C P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:02 Y:0C P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:02 Y:0C P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:02 Y:0C P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:02 Y:0C P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:02 Y:0C P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:02 Y:0C P:27 SP:F9
	// F6A4  AD 47 06  LDA $0647 = 9B                  A:11 X:02 Y:0C P:27 SP:FB
	// F6A7  C9 9B     CMP #$9B                        A:9B X:02 Y:0C P:A5 SP:FB
	// F6A9  F0 02     BEQ $F6AD                       A:9B X:02 Y:0C P:27 SP:FB
	// F6AD  A0 0D     LDY #$0D                        A:9B X:02 Y:0C P:27 SP:FB
	// F6AF  A2 FF     LDX #$FF                        A:9B X:02 Y:0D P:25 SP:FB
	// F6B1  A9 A5     LDA #$A5                        A:9B X:FF Y:0D P:A5 SP:FB
	// F6B3  85 47     STA $47 = 9B                    A:A5 X:FF Y:0D P:A5 SP:FB
	// F6B5  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:0D P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:0D P:A5 SP:F9
	// FAEB  18        CLC                             A:A5 X:FF Y:0D P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:0D P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:FF Y:0D P:E4 SP:F9
	// F6B8  77 48    *RRA $48,X @ 47 = A5             A:B2 X:FF Y:0D P:E4 SP:FB
	// F6BA  EA        NOP                             A:05 X:FF Y:0D P:25 SP:FB
	// F6BB  EA        NOP                             A:05 X:FF Y:0D P:25 SP:FB
	// F6BC  EA        NOP                             A:05 X:FF Y:0D P:25 SP:FB
	// F6BD  EA        NOP                             A:05 X:FF Y:0D P:25 SP:FB
	// F6BE  20 EF FA  JSR $FAEF                       A:05 X:FF Y:0D P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:FF Y:0D P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:FF Y:0D P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:FF Y:0D P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:FF Y:0D P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:FF Y:0D P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:FF Y:0D P:27 SP:F9
	// F6C1  A5 47     LDA $47 = 52                    A:05 X:FF Y:0D P:27 SP:FB
	// F6C3  C9 52     CMP #$52                        A:52 X:FF Y:0D P:25 SP:FB
	// F6C5  F0 02     BEQ $F6C9                       A:52 X:FF Y:0D P:27 SP:FB
	// F6C9  C8        INY                             A:52 X:FF Y:0D P:27 SP:FB
	// F6CA  A9 29     LDA #$29                        A:52 X:FF Y:0E P:25 SP:FB
	// F6CC  85 47     STA $47 = 52                    A:29 X:FF Y:0E P:25 SP:FB
	// F6CE  20 FA FA  JSR $FAFA                       A:29 X:FF Y:0E P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:FF Y:0E P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:FF Y:0E P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:FF Y:0E P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:FF Y:0E P:24 SP:F9
	// F6D1  77 48    *RRA $48,X @ 47 = 29             A:42 X:FF Y:0E P:24 SP:FB
	// F6D3  EA        NOP                             A:57 X:FF Y:0E P:24 SP:FB
	// F6D4  EA        NOP                             A:57 X:FF Y:0E P:24 SP:FB
	// F6D5  EA        NOP                             A:57 X:FF Y:0E P:24 SP:FB
	// F6D6  EA        NOP                             A:57 X:FF Y:0E P:24 SP:FB
	// F6D7  20 FF FA  JSR $FAFF                       A:57 X:FF Y:0E P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:0E P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:FF Y:0E P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:0E P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:FF Y:0E P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:0E P:27 SP:F9
	// FB09  60        RTS                             A:57 X:FF Y:0E P:27 SP:F9
	// F6DA  A5 47     LDA $47 = 14                    A:57 X:FF Y:0E P:27 SP:FB
	// F6DC  C9 14     CMP #$14                        A:14 X:FF Y:0E P:25 SP:FB
	// F6DE  F0 02     BEQ $F6E2                       A:14 X:FF Y:0E P:27 SP:FB
	// F6E2  C8        INY                             A:14 X:FF Y:0E P:27 SP:FB
	// F6E3  A9 37     LDA #$37                        A:14 X:FF Y:0F P:25 SP:FB
	// F6E5  85 47     STA $47 = 14                    A:37 X:FF Y:0F P:25 SP:FB
	// F6E7  20 0A FB  JSR $FB0A                       A:37 X:FF Y:0F P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:0F P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:FF Y:0F P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:FF Y:0F P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:FF Y:0F P:65 SP:F9
	// F6EA  77 48    *RRA $48,X @ 47 = 37             A:75 X:FF Y:0F P:65 SP:FB
	// F6EC  EA        NOP                             A:11 X:FF Y:0F P:25 SP:FB
	// F6ED  EA        NOP                             A:11 X:FF Y:0F P:25 SP:FB
	// F6EE  EA        NOP                             A:11 X:FF Y:0F P:25 SP:FB
	// F6EF  EA        NOP                             A:11 X:FF Y:0F P:25 SP:FB
	// F6F0  20 10 FB  JSR $FB10                       A:11 X:FF Y:0F P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:FF Y:0F P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:FF Y:0F P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:FF Y:0F P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:FF Y:0F P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:0F P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:FF Y:0F P:27 SP:F9
	// F6F3  A5 47     LDA $47 = 9B                    A:11 X:FF Y:0F P:27 SP:FB
	// F6F5  C9 9B     CMP #$9B                        A:9B X:FF Y:0F P:A5 SP:FB
	// F6F7  F0 02     BEQ $F6FB                       A:9B X:FF Y:0F P:27 SP:FB
	// F6FB  A9 A5     LDA #$A5                        A:9B X:FF Y:0F P:27 SP:FB
	// F6FD  8D 47 06  STA $0647 = 9B                  A:A5 X:FF Y:0F P:A5 SP:FB
	// F700  A0 FF     LDY #$FF                        A:A5 X:FF Y:0F P:A5 SP:FB
	// F702  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:FF P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:FF P:A5 SP:F9
	// FAEB  18        CLC                             A:A5 X:FF Y:FF P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:FF P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:FF Y:FF P:E4 SP:F9
	// F705  7B 48 05 *RRA $0548,Y @ 0647 = A5         A:B2 X:FF Y:FF P:E4 SP:FB
	// F708  EA        NOP                             A:05 X:FF Y:FF P:25 SP:FB
	// F709  EA        NOP                             A:05 X:FF Y:FF P:25 SP:FB
	// F70A  08        PHP                             A:05 X:FF Y:FF P:25 SP:FB
	// F70B  48        PHA                             A:05 X:FF Y:FF P:25 SP:FA
	// F70C  A0 10     LDY #$10                        A:05 X:FF Y:FF P:25 SP:F9
	// F70E  68        PLA                             A:05 X:FF Y:10 P:25 SP:F9
	// F70F  28        PLP                             A:05 X:FF Y:10 P:25 SP:FA
	// F710  20 EF FA  JSR $FAEF                       A:05 X:FF Y:10 P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:FF Y:10 P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:FF Y:10 P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:FF Y:10 P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:FF Y:10 P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:FF Y:10 P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:FF Y:10 P:27 SP:F9
	// F713  AD 47 06  LDA $0647 = 52                  A:05 X:FF Y:10 P:27 SP:FB
	// F716  C9 52     CMP #$52                        A:52 X:FF Y:10 P:25 SP:FB
	// F718  F0 02     BEQ $F71C                       A:52 X:FF Y:10 P:27 SP:FB
	// F71C  A0 FF     LDY #$FF                        A:52 X:FF Y:10 P:27 SP:FB
	// F71E  A9 29     LDA #$29                        A:52 X:FF Y:FF P:A5 SP:FB
	// F720  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:FF P:25 SP:FB
	// F723  20 FA FA  JSR $FAFA                       A:29 X:FF Y:FF P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:FF Y:FF P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:FF Y:FF P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:FF Y:FF P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:FF Y:FF P:24 SP:F9
	// F726  7B 48 05 *RRA $0548,Y @ 0647 = 29         A:42 X:FF Y:FF P:24 SP:FB
	// F729  EA        NOP                             A:57 X:FF Y:FF P:24 SP:FB
	// F72A  EA        NOP                             A:57 X:FF Y:FF P:24 SP:FB
	// F72B  08        PHP                             A:57 X:FF Y:FF P:24 SP:FB
	// F72C  48        PHA                             A:57 X:FF Y:FF P:24 SP:FA
	// F72D  A0 11     LDY #$11                        A:57 X:FF Y:FF P:24 SP:F9
	// F72F  68        PLA                             A:57 X:FF Y:11 P:24 SP:F9
	// F730  28        PLP                             A:57 X:FF Y:11 P:24 SP:FA
	// F731  20 FF FA  JSR $FAFF                       A:57 X:FF Y:11 P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:11 P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:FF Y:11 P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:11 P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:FF Y:11 P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:11 P:27 SP:F9
	// FB09  60        RTS                             A:57 X:FF Y:11 P:27 SP:F9
	// F734  AD 47 06  LDA $0647 = 14                  A:57 X:FF Y:11 P:27 SP:FB
	// F737  C9 14     CMP #$14                        A:14 X:FF Y:11 P:25 SP:FB
	// F739  F0 02     BEQ $F73D                       A:14 X:FF Y:11 P:27 SP:FB
	// F73D  A0 FF     LDY #$FF                        A:14 X:FF Y:11 P:27 SP:FB
	// F73F  A9 37     LDA #$37                        A:14 X:FF Y:FF P:A5 SP:FB
	// F741  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:FF P:25 SP:FB
	// F744  20 0A FB  JSR $FB0A                       A:37 X:FF Y:FF P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:FF P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:FF Y:FF P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:FF Y:FF P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:FF Y:FF P:65 SP:F9
	// F747  7B 48 05 *RRA $0548,Y @ 0647 = 37         A:75 X:FF Y:FF P:65 SP:FB
	// F74A  EA        NOP                             A:11 X:FF Y:FF P:25 SP:FB
	// F74B  EA        NOP                             A:11 X:FF Y:FF P:25 SP:FB
	// F74C  08        PHP                             A:11 X:FF Y:FF P:25 SP:FB
	// F74D  48        PHA                             A:11 X:FF Y:FF P:25 SP:FA
	// F74E  A0 12     LDY #$12                        A:11 X:FF Y:FF P:25 SP:F9
	// F750  68        PLA                             A:11 X:FF Y:12 P:25 SP:F9
	// F751  28        PLP                             A:11 X:FF Y:12 P:25 SP:FA
	// F752  20 10 FB  JSR $FB10                       A:11 X:FF Y:12 P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:FF Y:12 P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:FF Y:12 P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:FF Y:12 P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:FF Y:12 P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:12 P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:FF Y:12 P:27 SP:F9
	// F755  AD 47 06  LDA $0647 = 9B                  A:11 X:FF Y:12 P:27 SP:FB
	// F758  C9 9B     CMP #$9B                        A:9B X:FF Y:12 P:A5 SP:FB
	// F75A  F0 02     BEQ $F75E                       A:9B X:FF Y:12 P:27 SP:FB
	// F75E  A0 13     LDY #$13                        A:9B X:FF Y:12 P:27 SP:FB
	// F760  A2 FF     LDX #$FF                        A:9B X:FF Y:13 P:25 SP:FB
	// F762  A9 A5     LDA #$A5                        A:9B X:FF Y:13 P:A5 SP:FB
	// F764  8D 47 06  STA $0647 = 9B                  A:A5 X:FF Y:13 P:A5 SP:FB
	// F767  20 E9 FA  JSR $FAE9                       A:A5 X:FF Y:13 P:A5 SP:FB
	// FAE9  24 01     BIT $01 = FF                    A:A5 X:FF Y:13 P:A5 SP:F9
	// FAEB  18        CLC                             A:A5 X:FF Y:13 P:E5 SP:F9
	// FAEC  A9 B2     LDA #$B2                        A:A5 X:FF Y:13 P:E4 SP:F9
	// FAEE  60        RTS                             A:B2 X:FF Y:13 P:E4 SP:F9
	// F76A  7F 48 05 *RRA $0548,X @ 0647 = A5         A:B2 X:FF Y:13 P:E4 SP:FB
	// F76D  EA        NOP                             A:05 X:FF Y:13 P:25 SP:FB
	// F76E  EA        NOP                             A:05 X:FF Y:13 P:25 SP:FB
	// F76F  EA        NOP                             A:05 X:FF Y:13 P:25 SP:FB
	// F770  EA        NOP                             A:05 X:FF Y:13 P:25 SP:FB
	// F771  20 EF FA  JSR $FAEF                       A:05 X:FF Y:13 P:25 SP:FB
	// FAEF  70 2A     BVS $FB1B                       A:05 X:FF Y:13 P:25 SP:F9
	// FAF1  90 28     BCC $FB1B                       A:05 X:FF Y:13 P:25 SP:F9
	// FAF3  30 26     BMI $FB1B                       A:05 X:FF Y:13 P:25 SP:F9
	// FAF5  C9 05     CMP #$05                        A:05 X:FF Y:13 P:25 SP:F9
	// FAF7  D0 22     BNE $FB1B                       A:05 X:FF Y:13 P:27 SP:F9
	// FAF9  60        RTS                             A:05 X:FF Y:13 P:27 SP:F9
	// F774  AD 47 06  LDA $0647 = 52                  A:05 X:FF Y:13 P:27 SP:FB
	// F777  C9 52     CMP #$52                        A:52 X:FF Y:13 P:25 SP:FB
	// F779  F0 02     BEQ $F77D                       A:52 X:FF Y:13 P:27 SP:FB
	// F77D  C8        INY                             A:52 X:FF Y:13 P:27 SP:FB
	// F77E  A9 29     LDA #$29                        A:52 X:FF Y:14 P:25 SP:FB
	// F780  8D 47 06  STA $0647 = 52                  A:29 X:FF Y:14 P:25 SP:FB
	// F783  20 FA FA  JSR $FAFA                       A:29 X:FF Y:14 P:25 SP:FB
	// FAFA  B8        CLV                             A:29 X:FF Y:14 P:25 SP:F9
	// FAFB  18        CLC                             A:29 X:FF Y:14 P:25 SP:F9
	// FAFC  A9 42     LDA #$42                        A:29 X:FF Y:14 P:24 SP:F9
	// FAFE  60        RTS                             A:42 X:FF Y:14 P:24 SP:F9
	// F786  7F 48 05 *RRA $0548,X @ 0647 = 29         A:42 X:FF Y:14 P:24 SP:FB
	// F789  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB
	// F78A  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB
	// F78B  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB
	// F78C  EA        NOP                             A:57 X:FF Y:14 P:24 SP:FB
	// F78D  20 FF FA  JSR $FAFF                       A:57 X:FF Y:14 P:24 SP:FB
	// FAFF  70 1A     BVS $FB1B                       A:57 X:FF Y:14 P:24 SP:F9
	// FB01  30 18     BMI $FB1B                       A:57 X:FF Y:14 P:24 SP:F9
	// FB03  B0 16     BCS $FB1B                       A:57 X:FF Y:14 P:24 SP:F9
	// FB05  C9 57     CMP #$57                        A:57 X:FF Y:14 P:24 SP:F9
	// FB07  D0 12     BNE $FB1B                       A:57 X:FF Y:14 P:27 SP:F9
	// FB09  60        RTS                             A:57 X:FF Y:14 P:27 SP:F9
	// F790  AD 47 06  LDA $0647 = 14                  A:57 X:FF Y:14 P:27 SP:FB
	// F793  C9 14     CMP #$14                        A:14 X:FF Y:14 P:25 SP:FB
	// F795  F0 02     BEQ $F799                       A:14 X:FF Y:14 P:27 SP:FB
	// F799  C8        INY                             A:14 X:FF Y:14 P:27 SP:FB
	// F79A  A9 37     LDA #$37                        A:14 X:FF Y:15 P:25 SP:FB
	// F79C  8D 47 06  STA $0647 = 14                  A:37 X:FF Y:15 P:25 SP:FB
	// F79F  20 0A FB  JSR $FB0A                       A:37 X:FF Y:15 P:25 SP:FB
	// FB0A  24 01     BIT $01 = FF                    A:37 X:FF Y:15 P:25 SP:F9
	// FB0C  38        SEC                             A:37 X:FF Y:15 P:E5 SP:F9
	// FB0D  A9 75     LDA #$75                        A:37 X:FF Y:15 P:E5 SP:F9
	// FB0F  60        RTS                             A:75 X:FF Y:15 P:65 SP:F9
	// F7A2  7F 48 05 *RRA $0548,X @ 0647 = 37         A:75 X:FF Y:15 P:65 SP:FB
	// F7A5  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB
	// F7A6  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB
	// F7A7  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB
	// F7A8  EA        NOP                             A:11 X:FF Y:15 P:25 SP:FB
	// F7A9  20 10 FB  JSR $FB10                       A:11 X:FF Y:15 P:25 SP:FB
	// FB10  70 09     BVS $FB1B                       A:11 X:FF Y:15 P:25 SP:F9
	// FB12  30 07     BMI $FB1B                       A:11 X:FF Y:15 P:25 SP:F9
	// FB14  90 05     BCC $FB1B                       A:11 X:FF Y:15 P:25 SP:F9
	// FB16  C9 11     CMP #$11                        A:11 X:FF Y:15 P:25 SP:F9
	// FB18  D0 01     BNE $FB1B                       A:11 X:FF Y:15 P:27 SP:F9
	// FB1A  60        RTS                             A:11 X:FF Y:15 P:27 SP:F9
	// F7AC  AD 47 06  LDA $0647 = 9B                  A:11 X:FF Y:15 P:27 SP:FB
	// F7AF  C9 9B     CMP #$9B                        A:9B X:FF Y:15 P:A5 SP:FB
	// F7B1  F0 02     BEQ $F7B5                       A:9B X:FF Y:15 P:27 SP:FB
	// F7B5  60        RTS                             A:9B X:FF Y:15 P:27 SP:FB
	// C655  A5 00     LDA $00 = 00                    A:9B X:FF Y:15 P:27 SP:FD
	// C657  05 10     ORA $10 = 00                    A:00 X:FF Y:15 P:27 SP:FD
	// C659  05 11     ORA $11 = 00                    A:00 X:FF Y:15 P:27 SP:FD
	// C65B  F0 0E     BEQ $C66B                       A:00 X:FF Y:15 P:27 SP:FD
	// C66B  20 89 C6  JSR $C689                       A:00 X:FF Y:15 P:27 SP:FD
	// C689  A9 02     LDA #$02                        A:00 X:FF Y:15 P:27 SP:FB
	// C68B  8D 15 40  STA $4015 = FF                  A:02 X:FF Y:15 P:25 SP:FB
	// C68E  A9 3F     LDA #$3F                        A:02 X:FF Y:15 P:25 SP:FB
	// C690  8D 04 40  STA $4004 = FF                  A:3F X:FF Y:15 P:25 SP:FB
	// C693  A9 9A     LDA #$9A                        A:3F X:FF Y:15 P:25 SP:FB
	// C695  8D 05 40  STA $4005 = FF                  A:9A X:FF Y:15 P:A5 SP:FB
	// C698  A9 FF     LDA #$FF                        A:9A X:FF Y:15 P:A5 SP:FB
	// C69A  8D 06 40  STA $4006 = FF                  A:FF X:FF Y:15 P:A5 SP:FB
	// C69D  A9 00     LDA #$00                        A:FF X:FF Y:15 P:A5 SP:FB
	// C69F  8D 07 40  STA $4007 = FF                  A:00 X:FF Y:15 P:27 SP:FB
	// C6A2  60        RTS                             A:00 X:FF Y:15 P:27 SP:FB
	// C66E  60        RTS                             A:00 X:FF Y:15 P:27 SP:FD
	// 0001  FF 00 00 *ISB $0000,X @ 00FF = 46         A:00 X:FF Y:15 P:27 SP:FF
	// 0004  00        BRK                             A:B9 X:FF Y:15 P:A4 SP:FF
}
