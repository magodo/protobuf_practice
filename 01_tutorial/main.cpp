/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Thu 18 Jan 2018 10:34:52 PM CST
 Description: 
 ************************************************************************/

#include <iostream>
#include <fstream>
#include <string>
#include "addressbook.pb.h"

using namespace std;

void listPeople(const tutorial::AddressBook& address_book)
{
    for (int i = 0; i < address_book.people_size(); ++i)
    {
        auto people = address_book.people(i);
        cout << "-----------------------" << endl;
        cout << "Name : " << people.name() << endl;
        cout << "ID   : " << people.id() << endl;
        if (people.has_email())
            cout << "Email: " << people.email() << endl;

        for (int j = 0; j < people.phones_size(); ++j)
        {
            auto phone = people.phones(i);
            cout << "Phone: " << phone.number();
            auto type = phone.type();
            switch (type)
            {
                case tutorial::Person::MOBILE:
                    cout << " (mobile)" << endl;
                    break;
                case tutorial::Person::HOME:
                    cout << " (home)" << endl;
                    break;
                case tutorial::Person::WORK:
                    cout << " (work)" << endl;
                    break;
            }
        }
    }
    cout << "-----------------------" << endl;
}

bool addPerson(tutorial::Person *person)
{
    cout << "Enter name: ";
    getline(cin, *person->mutable_name());

    cout << "Enter id: ";
    int32_t id;
    cin >> id;
    if (!cin)
    {
        cerr << "Invalid ID entered!" << endl;
        cin.clear();
        cin.ignore(numeric_limits<streamsize>::max(), '\n');
        return false;
    }
    person->set_id(id);
    cin.ignore(256, '\n');

    cout << "Enter email(blank for none): ";
    string email;
    getline(cin, email);
    if (!email.empty())
        person->set_email(email);

    while (1)
    {
        cout << "Enter phone number(blank to end): ";
        string number;
        getline(cin, number);
        if (number.empty())
            break;

        auto phone = person->add_phones();
        phone->set_number(number);

        cout << "Phone number type (mobile/home/work): ";
        string type;
        getline(cin, type);
        if (type == "mobile")
            phone->set_type(tutorial::Person::MOBILE);
        else if (type == "home")
            phone->set_type(tutorial::Person::HOME);
        else if (type == "work")
            phone->set_type(tutorial::Person::WORK);
        else
            cerr << "Unknown type!" << endl;
    }
    return true;
}

int main(int argc, char *argv[])
{
    /* check version between header and lib */
    GOOGLE_PROTOBUF_VERIFY_VERSION;

    if (argc != 2)
    {
        cerr << "usage: ./" << argv[0] << " <input_file>" << endl;
        return -1;
    }

    tutorial::AddressBook address_book;

    /* read from file */
    {
        fstream f(argv[1], ios_base::in | ios_base::binary);
        if (!f) // same as f.fail(), i.e. f.failbit or badbit is set.
        {
            cerr << "File not exists, creating one..." << endl;
            f.open(argv[1], ios_base::out);
        }
        else if (!address_book.ParseFromIstream(&f))
        {
            cerr << "Failed to parse address book!" << endl;
            return -1;
        }
    }

    while (true)
    {
        cout << "[1] Add one person\n"
             << "[2] List people\n"
             << "[q] Quit\n"
             << "What to do: ";

        string choice;
        getline(cin, choice);

        if (choice == "q" || choice == "Q")
            break;
        else if (choice == "1")
        {
            /* add an address */
            bool ret = addPerson(address_book.add_people());
            
            if (ret)
            {
                /* write back to file */
                fstream f(argv[1], ios_base::out | ios_base::trunc | ios_base::binary);
                if (!address_book.SerializeToOstream(&f))
                {
                    cerr << "Failed to write back to address book!" << endl;
                    return -1;
                }
            }
            else
            {
                /* remove the last added repeated field */
                address_book.GetReflection()->RemoveLast(&address_book, address_book.GetDescriptor()->FindFieldByName("people"));
            }
        }
        else if (choice == "2")
        {
            listPeople(address_book);
        }
    }

    // Optional:  Delete all global objects allocated by libprotobuf.
    //            (useful for libs which will be loaded and unloaded
    //             multiple times by a same process.)
    google::protobuf::ShutdownProtobufLibrary();

    return 0;
}
