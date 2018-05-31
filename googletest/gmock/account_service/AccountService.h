// AccountService.h
// the class to be tested
#pragma once

#include <string>

#include "Account.h"
#include "AccountManager.h"

class AccountService
{
private:
	AccountManager* pAccountManager;

public:
	AccountService();

	void setAccountManager(AccountManager* pManager);

	void transfer(const std::string& senderId, const std::string& beneficiaryId, long amount);
};
