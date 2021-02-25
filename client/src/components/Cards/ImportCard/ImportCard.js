import react from 'react';

const ImportCard = () => {
    return (
        <div className="card">
            <i className="fa fa-upload fa-2x text-green" aria-hidden="true"></i>
            <div className="card_inner">
                <p className="text-primary-p">Import contacts</p>
            </div>
        </div>
    );
}

export default ImportCard;