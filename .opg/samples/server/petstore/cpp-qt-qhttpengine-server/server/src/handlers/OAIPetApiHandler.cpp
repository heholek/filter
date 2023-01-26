/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QVariantMap>
#include <QDebug>

#include "OAIPetApiHandler.h"
#include "OAIPetApiRequest.h"

namespace OpenAPI {

OAIPetApiHandler::OAIPetApiHandler(){

}

OAIPetApiHandler::~OAIPetApiHandler(){

}

void OAIPetApiHandler::addPet(OAIPet body) {
    Q_UNUSED(body);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        
        reqObj->addPetResponse();
    }
}
void OAIPetApiHandler::deletePet(qint64 pet_id, QString api_key) {
    Q_UNUSED(pet_id);
    Q_UNUSED(api_key);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        
        reqObj->deletePetResponse();
    }
}
void OAIPetApiHandler::findPetsByStatus(QList<QString> status) {
    Q_UNUSED(status);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        QList<OAIPet> res;
        reqObj->findPetsByStatusResponse(res);
    }
}
void OAIPetApiHandler::findPetsByTags(QList<QString> tags) {
    Q_UNUSED(tags);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        QList<OAIPet> res;
        reqObj->findPetsByTagsResponse(res);
    }
}
void OAIPetApiHandler::getPetById(qint64 pet_id) {
    Q_UNUSED(pet_id);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        OAIPet res;
        reqObj->getPetByIdResponse(res);
    }
}
void OAIPetApiHandler::updatePet(OAIPet body) {
    Q_UNUSED(body);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        
        reqObj->updatePetResponse();
    }
}
void OAIPetApiHandler::updatePetWithForm(qint64 pet_id, QString name, QString status) {
    Q_UNUSED(pet_id);
    Q_UNUSED(name);
    Q_UNUSED(status);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        
        reqObj->updatePetWithFormResponse();
    }
}
void OAIPetApiHandler::uploadFile(qint64 pet_id, QString additional_metadata, OAIHttpFileElement file) {
    Q_UNUSED(pet_id);
    Q_UNUSED(additional_metadata);
    Q_UNUSED(file);
    auto reqObj = qobject_cast<OAIPetApiRequest*>(sender());
    if( reqObj != nullptr )
    {
        OAIApiResponse res;
        reqObj->uploadFileResponse(res);
    }
}


}
