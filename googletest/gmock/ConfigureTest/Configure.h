// Configure.h
#pragma once

#include <string>
#include <vector>

class Configure
{
private:
	std::vector<std::string> vItems;

public:
	int addItem(std::string str);

	std::string getItem(int index);

	int getSize();
};
