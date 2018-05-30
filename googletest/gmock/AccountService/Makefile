.PHONY : all clean run

TARGET=account
SRCS=$(wildcard *.cpp)
OBJS=$(patsubst %.cpp,%.o,$(SRCS))

INCLUDE_FLAGS=-I.

CXXFLAGS+=-fPIC -O2 $(INCLUDE_FLAGS)

all : $(TARGET)

$(TARGET) : $(OBJS)
	$(CXX) -o $@ $^

clean :
	-rm -rf *.o $(TARGET)

run : all
	./$(TARGET)
