# 工具链
CC = g++
LD = ld
AR = ar
RM = rm -f
STRIP = strip

#MACROS	 = _LINUX_ 
#ARCMACROS = $(patsubst %,-D%,$(MACROS))

# 头文件路径
INC += -I./

# 第三方库
#EXTLIBS += -L../output/lib -lRtcMessage 

# 编译选项
CFLAGS += $(ARCMACROS) -Wall -O2 -Os -g  

# 系统库
LDFLAGS +=  -lgtest -pthread -lgtest_main  
  


# 需要编译的文件（静态库）
SRCPATH = .
SRC_1 += $(SRCPATH)/sample1.cc
SRC_1 += $(SRCPATH)/sample1_unittest.cc
S1_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_1)))
TARGET_1=./sample1_unittest 

SRC_2 += $(SRCPATH)/sample2.cc
SRC_2 += $(SRCPATH)/sample2_unittest.cc
S2_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_2)))
TARGET_2=./sample2_unittest 

SRC_3 += $(SRCPATH)/sample3_unittest.cc
S3_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_3)))
TARGET_3=./sample3_unittest 

SRC_4 += $(SRCPATH)/sample4.cc
SRC_4 += $(SRCPATH)/sample4_unittest.cc
S4_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_4)))
TARGET_4=./sample4_unittest 

SRC_5 += $(SRCPATH)/sample1.cc
SRC_5 += $(SRCPATH)/sample5_unittest.cc
S5_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_5)))
TARGET_5=./sample5_unittest 

SRC_6 += $(SRCPATH)/sample6_unittest.cc
S6_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_6)))
TARGET_6=./sample6_unittest 

SRC_7 += $(SRCPATH)/sample7_unittest.cc
S7_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_7)))
TARGET_7=./sample7_unittest 

SRC_8 += $(SRCPATH)/sample8_unittest.cc
S8_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_8)))
TARGET_8=./sample8_unittest 

SRC_9 += $(SRCPATH)/sample9_unittest.cc
S9_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_9)))
TARGET_9=./sample9_unittest 

SRC_10 += $(SRCPATH)/sample10_unittest.cc
S10_OBJS = $(patsubst %.c,%.o,$(patsubst %.cc,%.o,$(SRC_10)))
TARGET_10=./sample10_unittest 

all:$(TARGET_1) $(TARGET_2) $(TARGET_3) $(TARGET_4) $(TARGET_5) $(TARGET_6) $(TARGET_7) $(TARGET_8) $(TARGET_9) $(TARGET_10)

$(TARGET_1):$(S1_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_2):$(S2_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_3):$(S3_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_4):$(S4_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_5):$(S5_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_6):$(S6_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_7):$(S7_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_8):$(S8_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_9):$(S9_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)

$(TARGET_10):$(S10_OBJS) 
	$(CC) -o $@ $^ $(EXTLIBS) $(LDFLAGS)
 
OBJS += $(TARGET_1) $(S1_OBJS) $(TARGET_2) $(S2_OBJS) $(TARGET_3) $(S3_OBJS) $(TARGET_4) $(S4_OBJS) 
OBJS += $(TARGET_5) $(S5_OBJS) $(TARGET_6) $(S6_OBJS) $(TARGET_7) $(S7_OBJS) $(TARGET_8) $(S8_OBJS) 
OBJS += $(TARGET_9) $(S9_OBJS) $(TARGET_10) $(S10_OBJS) 

.c.o:
	$(CC) -c $(CFLAGS) $^ -o $@
	
clean:
	$(RM) $(OBJS)