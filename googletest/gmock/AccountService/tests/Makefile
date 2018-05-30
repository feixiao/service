.PHONY : all clean run

TARGET=accountTest
SRCS=$(wildcard *.cpp) $(shell ls ../*.cpp | grep -v Main)
OBJS=$(patsubst %.cpp,%.o,$(SRCS))

INCLUDE_FLAGS=-I. -I$(GTEST_DIR)/include -I$(GMOCK_DIR)/include

CXXFLAGS+=-fPIC -O2 $(INCLUDE_FLAGS)

all : $(TARGET)

$(TARGET) : $(OBJS)
	$(CXX) -o $@ $^ -L. -lgtest -lgmock

clean :
	-rm -rf *.o $(TARGET)

run : all
	./$(TARGET)
