.PHONY : all clean run

TARGET=configureTest
SRCS=$(wildcard *.cpp)
OBJS=$(patsubst %.cpp,%.o,$(SRCS))

INCLUDE_FLAGS=-I. -I$(GTEST_DIR)/include

CXXFLAGS+=-fPIC -O2 $(INCLUDE_FLAGS)

all : $(TARGET)

$(TARGET) : $(OBJS)
	$(CXX) -o $@ $^ -L. -lgtest

clean :
	-rm -rf *.o $(TARGET)

run : all
	./$(TARGET)
